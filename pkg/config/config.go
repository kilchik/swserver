//go:generate congo

package config

type Desc struct {
	listenGrpcAddr  string
	listenHttpAddr  string
	kafkaBrokers    string
	databaseDsn     string
	enableDebugLogs bool
}
