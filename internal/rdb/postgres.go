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
	case "json":
		id, _ := strconv.Atoi(r.FormValue("id"))
		var dto UserDto
		err := p.Db.QueryRow("SELECT id, updated_at FROM sandbox_json_table where id = $1", id).
			Scan(&dto.id, &dto.update)
		if err != nil {
			fmt.Errorf("error %s", err.Error())
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte("ng"))
			return
		}

		var ud UpdateAt
		_ = json.Unmarshal(dto.update, &ud)

		user := User{
			id:     dto.id,
			update: &ud,
		}
		fmt.Println(user)
		fmt.Println(user.update)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("ok"))

	case "insert":
		fmt.Println(time.Now())
		id, _ := strconv.Atoi(r.FormValue("id"))
		data := Data{
			Id:       id,
			Name:     r.FormValue("name"),
			Category: r.FormValue("c"),
		}
		//goChan := make(chan sql.Result)
		goChan := make(chan struct{})
		//defer close(goChan)
		p.asyncWriteSandboxTable(data, goChan)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("ok"))
		// ここで待ってしまうので注意
		<-goChan
		fmt.Println(time.Now())
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

func (p PostgresHandler) asyncWriteSandboxTable(d Data, channel chan<- struct{}) {
	//defer close(channel)
	//defer hoge2(channel)
	//if true {
	//	fmt.Println("nothing to do")
	//	time.Sleep(10 * time.Second)
	//	//go hoge(channel)
	//	return
	//}
	go func() {
		// 非同期の確認のために wait 入れている
		time.Sleep(10 * time.Second)
		_, err := p.syncWriteSandboxTable(d)
		if err != nil {
			fmt.Println("failed syncWriteSandboxTable: ", err)
		}
		channel <- struct{}{}
		close(channel)
	}()
}

func hoge(channel chan<- struct{}) {
	channel <- struct{}{}
	fmt.Println("empty add channel")
}

func hoge2(channel chan<- struct{}) {
	go func() {
		channel <- struct{}{}
		fmt.Println("empty add channel")
	}()
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

type UserDto struct {
	id int
	//Update *UpdateAt
	update json.RawMessage
}

type User struct {
	id     int
	update *UpdateAt
}

type UpdateAt struct {
	Nanos   int `json:"Nanos"`
	Seconds int `json:"Seconds"`
}
