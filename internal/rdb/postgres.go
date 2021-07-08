package rdb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	// import driver
	_ "github.com/lib/pq"
)

type PostgresHandler struct {
	Db *sql.DB
}

func (p PostgresHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	queryType := r.FormValue("type")
	switch strings.ToLower(queryType) {
	case "insert":
		id, _ := strconv.Atoi(r.FormValue("id"))
		data := Data{
			Id:       id,
			Name:     r.FormValue("name"),
			Category: r.FormValue("c"),
		}
		goChan := p.asyncWriteSandboxTable(data)
		defer close(goChan)
		_, _ = w.Write([]byte("ok"))
		// ここで待ってしまうので注意
		<-goChan
	case "search":
		p.search(w)
	default:
		p.search(w)
	}
}

func (p PostgresHandler) search(w http.ResponseWriter) {
	data := p.getSandboxTable()
	for _, v := range data {
		// for debug
		fmt.Println(v.Name)
	}

	jsonValue, _ := json.Marshal(data)
	fmt.Println(string(jsonValue))
	_, _ = w.Write(jsonValue)
}

func (p PostgresHandler) asyncWriteSandboxTable(d Data) chan sql.Result {
	goChan := make(chan sql.Result)
	go func() {
		// 非同期の確認のために wait 入れている
		time.Sleep(10 * time.Second)
		r, err := p.syncWriteSandboxTable(d)
		if err != nil {
			fmt.Println("failed syncWriteSandboxTable: ", err)
		}
		goChan <- *r
	}()
	return goChan
}

func (p PostgresHandler) syncWriteSandboxTable(d Data) (*sql.Result, error) {
	ins, err := p.Db.Prepare("INSERT INTO sandbox_table VALUES($1,$2,$3)")
	if err != nil {
		log.Fatal("failed Prepare:", err)
		return nil, err
	}
	r, err := ins.Exec(d.Id, d.Name, d.Category)
	if err != nil {
		log.Fatal("failed Exec:", err)
		return nil, err
	}
	fmt.Println(r)
	return &r, nil
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
