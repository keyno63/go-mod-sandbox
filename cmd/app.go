package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"go-mod2/internal/cassandra"
	"go-mod2/internal/rdb"
	"log"
	"net/http"
)

var db *sql.DB

// 実行関数
func main() {
	// Http Handler の設定
	mux := http.NewServeMux()
	handler := NewHandler()
	mux.Handle("/api/go-app/handle", handler)
	mux.HandleFunc("/api/go-app/handle2", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})

	cHandler, err := NewCassandraHandler()
	if err == nil {
		mux.Handle("/api/go-app/cass", cHandler)
	}

	// db 初期化
	db, err := sql.Open("postgres", "postgres://postgres:pass@127.0.01:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln("接続失敗", err)
	} else {
		pHandler, _ := NewPostgresHandler(db)
		mux.Handle("/api/go-app/psql", pHandler)
	}
	defer db.Close()

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

func NewHandler() Handler {
	// DI
	repository := UserRepositoryImpl{db}
	service := UserServiceImpl{repository}
	controller := UserControllerImpl{service}
	app := App{controller}

	return handler{&app}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	userId := r.FormValue("user_id")
	user := h.app.ExecGetUser(userId)

	jsonUser, _ := json.Marshal(user)

	printValue := fmt.Sprintf("param value is = [%s]", user)
	fmt.Println(printValue)

	w.Header().Set("Content-Type", "application/json")
	//_, _ = fmt.Fprintf(w, "sample handle")
	ret, err := w.Write(jsonUser)
	if err != nil {
		fmt.Printf("occured error: " + err.Error())
	}

	fmt.Println("write ret: " + string(rune(ret)))
}

/**
App.
handler と Controller が 1:1 になっているので、この層は不要？
TODO: 要不要の検討, 修正.
*/

type App struct {
	userController UserController
}

func (app App) ExecGetUser(id string) UserAccount {
	return app.userController.GetUser(id)
}

/**
Controller
*/
type UserController interface {
	GetUser(id string) UserAccount
}

type UserControllerImpl struct {
	userService UserService
}

func (uc UserControllerImpl) GetUser(id string) UserAccount {
	return uc.userService.GetUser(id)
}

/**
Service
*/
type UserService interface {
	GetUser(id string) UserAccount
}

type UserServiceImpl struct {
	userRepository UserRepository
}

func (us UserServiceImpl) GetUser(id string) UserAccount {
	return us.userRepository.GetUser(id)
}

/**
Repository
*/
type UserRepository interface {
	GetUser(id string) UserAccount
}

type UserRepositoryImpl struct {
	dbConnector *sql.DB
}

func (us UserRepositoryImpl) GetUser(id string) UserAccount {
	// 仮
	// TODO: DBとの接続の実装
	return UserAccount{
		Id:        id,
		FirstName: "first",
		LastName:  "last",
	}
}

/**
  data struct
*/

type UserAccount struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
