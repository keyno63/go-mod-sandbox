package cassandra

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gocql/gocql"
)

type Handler struct {
	Session *gocql.Session
}

func (c Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var id, data, updateTime string
	// SELECT 文は Scan?
	if err := c.Session.Query("SELECT * FROM test_table;").
		Scan(&id, &data, &updateTime); err != nil {
		fmt.Println("failed")
		if errors.Is(err, gocql.ErrNotFound) {
			fmt.Println("ErrNotFound")
		}
		fmt.Println(err.Error())
	}
	v := fmt.Sprintf("write ret: id=[%s], data=[%s], time=[%s]", id, data, updateTime)
	fmt.Println(v)
	_, _ = w.Write([]byte(v))
}
