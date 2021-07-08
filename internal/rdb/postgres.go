package rdb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// import driver
	_ "github.com/lib/pq"
)

type PostgresHandler struct {
	Db *sql.DB
}

func (p PostgresHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := p.getSandboxTable()
	for _, v := range data {
		fmt.Println(v.Name)
	}

	jsonValue, _ := json.Marshal(data)
	fmt.Println(string(jsonValue))
	_, _ = w.Write(jsonValue)
}

func (p PostgresHandler) getSandboxTable() []Data {
	cmd := "SELECT * FROM sandbox_table"
	rows, err := p.Db.Query(cmd)
	if err != nil {
		fmt.Println("失敗: ", err.Error())
		return []Data{}
	}
	defer rows.Close()

	var data []Data
	for rows.Next() {
		var d Data
		err := rows.Scan(&d.Id, &d.Name, &d.Category)
		if err != nil {
			log.Fatalln("取得失敗", err)
		}
		data = append(data, d)
	}
	return data
}

type Data struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Response struct {
	Data []Data `json:"data"`
}
