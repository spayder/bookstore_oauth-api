package cassandra

import (
	"githab.com/spayder/bookstore_oauth-api/src/utils/config"
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	cluster = gocql.NewCluster(config.Env("DB_HOST"))
	cluster.Keyspace = config.Env("DB_NAME")
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
