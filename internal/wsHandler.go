package internal

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type InputInfo struct {
	Type       int8  `json:"type"` // 0 -> moveTo; 1 -> LeftClick; 2 -> RightClick; 3 -> KeyPress; 4 -> LeftPress; 5 -> leftRelease; 6 -> Scroll
	X          int32 `json:"x"`
	Y          int32 `json:"y"`
	Horizontal bool  `json:"horizontal"`
	Delta      int32 `json:"delta"`
	Key        rune  `json:"key"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
