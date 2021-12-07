package handler

import (
	"database/sql"
	"fmt"
	"go-mod-sandbox/internal/cassandra"
	"go-mod-sandbox/internal/libs/gzip"
	"go-mod-sandbox/internal/rdb"
	"net/http"

	"github.com/gocql/gocql"
)

const (
	//typeApplicationJson = "application/json"
	typeTextPlain = "text/plain"
	encodingGzip  = "gzip"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewCassandraHandler() (Handler, error) {

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "test"
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return cassandra.Handler{
		Session: session,
	}, nil
}

func NewPostgresHandler(db *sql.DB) (Handler, error) {
	return rdb.PostgresHandler{Db: db}, nil
}

// GzipFunc　は関数定義を返す関数
func GzipFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.FormValue("k")
		header := map[string]string{
			//"Content-Type": typeApplicationJson,
			"Content-Type":     typeTextPlain,
			"Content-Encoding": encodingGzip,
		}
		for k, v := range header {
			w.Header().Add(k, v)
		}
		w.WriteHeader(403)
		ret, err := gzip.Write(v)
		if err != nil {
			fmt.Println("error1")
			fmt.Println(err.Error())
		}
		unzipped, err := gzip.Read(ret)
		if err != nil {
			fmt.Println("error2")
			fmt.Println(err.Error())
		}
		fmt.Println(unzipped)
		_, _ = w.Write([]byte(ret))
	}
}
