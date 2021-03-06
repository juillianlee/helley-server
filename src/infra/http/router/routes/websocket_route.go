package routes

import (
	"fmt"
	"helley/src/infra/config"
	hWebsocket "helley/src/infra/http/websocket"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var pool = hWebsocket.NewPool()

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

func executeWebsocket(c echo.Context) error {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Fprintf(c.Response(), "%+v\n", err)
	}

	client := &hWebsocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
	return nil
}

func handleFile(c echo.Context) error {
	return c.File("websockets.html")
}

func NewWebsocketRoutes() []config.Route {
	go pool.Start()

	return makeWebSocketRoutes()
}

func makeWebSocketRoutes() []config.Route {
	return []config.Route{
		{
			Path:                   "/",
			Method:                 http.MethodGet,
			HandleFunc:             handleFile,
			RequiredAuthentication: false,
		},
		{
			Path:                   "/ws",
			Method:                 http.MethodGet,
			HandleFunc:             executeWebsocket,
			RequiredAuthentication: false,
		},
	}
}
