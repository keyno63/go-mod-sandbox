package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-mod-sandbox/internal/app/controller"
	"go-mod-sandbox/internal/app/domain/model"
	"go-mod-sandbox/internal/app/repository"
	"go-mod-sandbox/internal/app/service"
	libHandler "go-mod-sandbox/internal/libs/handler"
	"go-mod-sandbox/internal/parser"
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/go-ini/ini.v1"
)

// 実行関数
func main() {
	// Http Handler の設定
	mux := http.NewServeMux()
	//LoadConfig()
	//LoggingSettings(Config.LogFile)

	// db 初期化
	db, err := sql.Open("postgres", "postgres://postgres:pass@127.0.01:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln("接続失敗", err)
	} else {
		// Health Check
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`OK`))
		})

		handler := NewHandler(db)
		mux.Handle("/api/go-app/handle", handler)
		mux.HandleFunc("/api/go-app/handle2", func(w http.ResponseWriter, r *http.Request) {
			handler.ServeHTTP(w, r)
		})

		pHandler, _ := libHandler.NewPostgresHandler(db)
		mux.Handle("/api/go-app/psql", pHandler)
		mux.HandleFunc("/api/gzip", libHandler.GzipFunc())
	}
	defer db.Close()

	cHandler, err := libHandler.NewCassandraHandler()
	if err == nil {
		mux.Handle("/api/go-app/cass", cHandler)
	}

	// server の起動設定
	server := http.Server{
		//Addr:    ":8180",
		Addr:    ":5000",
		Handler: mux,
	}

	// server の起動
	_ = server.ListenAndServe()
}

type handler struct {
	app *App
}

func NewHandler(db *sql.DB) libHandler.Handler {
	// DI
	r := repository.NewUserRepositoryImpl(db)
	s := service.NewUserServiceImpl(r)
	c := controller.NewUserController(s)
	app := App{c}

	return handler{&app}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)

	userID := r.FormValue("user_id")
	user, err := h.app.ExecGetUser(userID)
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

	if parser.Parseable("") {
		fmt.Println("parseable")
	}

	fmt.Printf("write ret: %d\n", ret)
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

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config ConfigList

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
