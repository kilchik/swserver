module swatcher-server

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.0+incompatible
	github.com/jmoiron/sqlx v1.2.0
	github.com/kilchik/logo v1.0.1
	github.com/kilchik/swserver/pkg/swatcher-server v0.0.0-00010101000000-000000000000
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/shopspring/decimal v0.0.0-20190905144223-a36b5d85f337 // indirect
	golang.org/x/text v0.3.8 // indirect
	google.golang.org/genproto v0.0.0-20190911173649-1774047e7e51 // indirect
	google.golang.org/grpc v1.23.1
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace github.com/kilchik/swserver/pkg/swatcher-server => ./pkg/swatcher-server

go 1.12
