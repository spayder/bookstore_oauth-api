package cassandra

import (
	"githab.com/spayder/bookstore_oauth-api/src/utils/config"
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster(config.Env("DB_HOST"))
	cluster.Keyspace = config.Env("DB_NAME")
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
