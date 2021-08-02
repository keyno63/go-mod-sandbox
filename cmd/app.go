package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"go-mod2/internal/app/controller"
	"go-mod2/internal/app/model"
	"go-mod2/internal/app/repository"
	"go-mod2/internal/app/service"
	"go-mod2/internal/cassandra"
	"go-mod2/internal/rdb"
	"log"
	"net/http"
)

// 実行関数
func main() {
	// Http Handler の設定
	mux := http.NewServeMux()

	// db 初期化
	db, err := sql.Open("postgres", "postgres://postgres:pass@127.0.01:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln("接続失敗", err)
	} else {
		handler := NewHandler(db)
		mux.Handle("/api/go-app/handle", handler)
		mux.HandleFunc("/api/go-app/handle2", func(w http.ResponseWriter, r *http.Request) {
			handler.ServeHTTP(w, r)
		})

		pHandler, _ := NewPostgresHandler(db)
		mux.Handle("/api/go-app/psql", pHandler)
	}
	defer db.Close()

	cHandler, err := NewCassandraHandler()
	if err == nil {
		mux.Handle("/api/go-app/cass", cHandler)
	}

	// server の起動設定
	server := http.Server{
		Addr:    ":8180",
		Handler: mux,
	}

	// server の起動
	_ = server.ListenAndServe()
}

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

type handler struct {
	app *App
}

func NewHandler(db *sql.DB) Handler {
	// DI
	r := repository.NewUserRepositoryImpl(db)
	s := service.NewUserServiceImpl(r)
	c := controller.NewUserController(s)
	app := App{c}

	return handler{&app}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	userId := r.FormValue("user_id")
	user, err := h.app.ExecGetUser(userId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		//_, _ = fmt.Fprintf(w, "sample handle")
		_, err := w.Write([]byte("{\"res\": \"NG\""))
		if err != nil {
			fmt.Printf("occured error: %s\n", err.Error())
		}
		return
	}

	jsonUser, _ := json.Marshal(user)

	printValue := fmt.Sprintf("param value is = [%s]", user)
	fmt.Println(printValue)

	w.Header().Set("Content-Type", "application/json")
	//_, _ = fmt.Fprintf(w, "sample handle")
	ret, err := w.Write(jsonUser)
	if err != nil {
		fmt.Printf("occured error: %s \n", err.Error())
	}

	fmt.Printf("write ret: %d\n",ret)
}

/**
App.
handler と Controller が 1:1 になっているので、この層は不要？
TODO: 要不要の検討, 修正.
*/

type App struct {
	userController controller.UserController
}

func (app App) ExecGetUser(id string) (*model.UserAccount, error) {
	return app.userController.GetUser(id)
}
