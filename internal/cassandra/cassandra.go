package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"net/http"
)

type Handler struct {
	Session *gocql.Session
}

func (c Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var id, data, update_time string
	// SELECT 文は Scan?
	if err := c.Session.Query("SELECT * FROM test_table;").
		Scan(&id, &data, &update_time); err != nil {
		fmt.Println("failed")
		if err == gocql.ErrNotFound {
			fmt.Println("ErrNotFound")
		}
		fmt.Println(err.Error())
	}
	v := fmt.Sprintf("write ret: id=[%s], data=[%s], time=[%s]", id, data, update_time)
	fmt.Println(v)
	_, _ = w.Write([]byte(v))
}
