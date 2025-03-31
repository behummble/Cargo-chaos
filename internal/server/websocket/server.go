package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/behummble/Cargo-chaos/internal/config"
	"github.com/gorilla/websocket"
)

const packageName = "SERVER:"
var id int

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type WebServer struct {
	server *http.Server
	connections map[int]*websocket.Conn
	log *slog.Logger
}

func New(config *config.Config, log *slog.Logger) *WebServer {
	server := getServer(config)
	connections := make(map[int]*websocket.Conn, config.Server.ConnNumber)
	id = 1
	return &WebServer{
		server: server,
		connections: connections,
		log: log,
	}
}

func(webServer *WebServer) Register() {
	http.HandleFunc("/init", webServer.initHandler)
}

func(webServer *WebServer) Run() {
	webServer.log.Error(packageName, webServer.server.ListenAndServe())
}

func getServer(cfg *config.Config) *http.Server {
	return &http.Server{
		Addr: fmt.Sprintf("%s:%s", cfg.Server.Addres, cfg.Server.Port),
		Handler: http.NewServeMux(),
	}
}

func(webServer *WebServer) initHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	
    if err != nil {
        webServer.log.Error(packageName, err)
        return
    }
	webServer.connections[id] = conn
	id++
	
}