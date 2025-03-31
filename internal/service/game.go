package service

import (
	//"github.com/gorilla/websocket"
)

type Game interface {
	initialization(data []byte, msgType int)
}