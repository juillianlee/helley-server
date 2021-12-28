package routes

import (
	"app-helley/src/websocket"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateChatRoutes(db *mongo.Database) []Route {
	chatWebsocket := websocket.NewChatWebsocket()
	return makeChatRoutes(chatWebsocket)
}

func makeChatRoutes(handler websocket.ChatWebsocket) []Route {
	return []Route{
		{
			Path:                   "/chat",
			HandleFunc:             handler.Chat,
			RequiredAuthentication: false,
		},
		{
			Path:                   "/c",
			HandleFunc:             handler.ChatView,
			Method:                 http.MethodGet,
			RequiredAuthentication: false,
		},
	}
}
