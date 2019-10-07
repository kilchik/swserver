package store

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Storage interface {
	GetNewSrvEvents(ctx context.Context, uid int64, revFrom int64) (events []*Event, err error)
	AddSrvEvent(ctx context.Context, uid int64, event *Event) error
}

type StorageImpl struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *StorageImpl {
	return &StorageImpl{db}
}

const (
	SrvEventTypeAdd = iota
)

type Event struct {
	Type int32  `json:"type",db:"type"`
	Data []byte `json:"data",db:"data"`
}

func (s *StorageImpl) AddSrvEvent(ctx context.Context, uid int64, event *Event) error {
	if err := s.withTransaction(ctx, func(tx *sqlx.Tx) error {
		var maxRev int64
		if err := tx.GetContext(ctx, &maxRev, `
SELECT MAX(rev) FROM srv_events
WHERE uid=$1`, uid); err != nil {
			return errors.Wrapf(err, "get last user revision")
		}

		if _, err := tx.ExecContext(ctx, `
INSERT INTO srv_events(uid, type, data, rev) VALUES ($1, $2, $3, $4)`, uid, event.Type, event.Data, maxRev+1); err != nil {
			return errors.Wrapf(err, "insert new event")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "add srv event: run transaction")
	}

	return nil
}

func (s *StorageImpl) GetNewSrvEvents(ctx context.Context, uid int64, revFrom int64) (events []*Event, err error) {
	err = s.db.SelectContext(ctx, &events, `
SELECT type, data FROM srv_events 
WHERE uid=$1 AND rev>=$2 
ORDER BY created_at
LIMIT 10;`, uid, revFrom)
	return
}

func (s *StorageImpl) withTransaction(ctx context.Context, f func(tx *sqlx.Tx) error) (err error) {
	var trx *sqlx.Tx
	trx, err = s.db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "begin transaction")
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recover after panic: %v", r)
		}
		for {
			if err != nil {
				if errRollback := trx.Rollback(); errRollback != nil {
					err = errors.Wrapf(errRollback, fmt.Sprintf("rollback after error: %v", err))
				}
				return
			}
			if err := trx.Commit(); err != nil {
				err = errors.Wrapf(err, "commit transaction")
			} else {
				return
			}
		}
	}()
	err = f(trx)
	return err
}
