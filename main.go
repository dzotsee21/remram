package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"

	"github.com/bendahl/uinput"
	"github.com/gorilla/websocket"
)

type MousePos struct {
	X int32 `json:x`
	Y int32 `json:y`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	vtp, err := uinput.CreateTouchPad("/dev/uinput", []byte("VirtualTouchpad"), 0, 1919, 0, 1079) // should calculate device screen size.
	if err != nil {
		log.Fatal(err)
	}
	defer vtp.Close()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade err: ", err)
		return
	}
	defer conn.Close()

  res := map[string]interface{}{"type": "res", "x": 1919, "y": 1079}
  if err := conn.WriteJSON(res); err != nil {
      return
  }

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error: ", err)
			break
		}

		var pos MousePos 
    err = json.Unmarshal(p, &pos)
    if err != nil {
        log.Fatal(err)
    }

		fmt.Println(pos.X)
		fmt.Println("---")
		fmt.Println(pos.Y)

		vtp.MoveTo(pos.X, pos.Y)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", echoHandler)
	log.Println("server starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}

	// vk, err := uinput.CreateKeyboard("/dev/uinput", []byte("VirtualKeyboard"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// defer vk.Close()
	//
	// vm, err := uinput.CreateMouse("/dev/uinput", []byte("VirtualMouse"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// defer vm.Close()
	//
	// // vm.Move(100, 100)
	// // vm.LeftClick()
	//
	// vk.KeyPress(uinput.KeyA)
	// time.Sleep(100 * time.Millisecond)
	// vk.KeyPress(uinput.KeyI)
	// time.Sleep(100 * time.Millisecond)
}
