package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bendahl/uinput"
	"github.com/gorilla/websocket"
	"github.com/kbinani/screenshot"
)

type Display struct {
	X               int32 `json:"x"`
	Y               int32 `json:"y"`
	VirtualTouchpad uinput.TouchPad
}

var display Display

type MouseInfo struct {
	Type int8  `json:"type"` // 0 -> moveTo; 1 -> LeftClick; 2 -> RightClick
	X    int32 `json:"x"`
	Y    int32 `json:"y"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (d *Display) echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade err: ", err)
		return
	}
	defer conn.Close()

	res := map[string]any{"type": "res", "x": d.X - 1, "y": d.Y - 1}
	if err := conn.WriteJSON(res); err != nil {
		return
	}

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error: ", err)
			break
		}

		var mouseInfo MouseInfo
		err = json.Unmarshal(p, &mouseInfo)
		if err != nil {
			log.Fatal(err)
		}

		switch mouseInfo.Type {
		case 0: // Move Mouse Position
			fmt.Println(mouseInfo.X)
			fmt.Println("---")
			fmt.Println(mouseInfo.Y)

			err = d.VirtualTouchpad.MoveTo(mouseInfo.X, mouseInfo.Y)
			if err != nil {
				log.Fatal(err)
			}

		case 1: // Left Click
			err = d.VirtualTouchpad.LeftClick()
			if err != nil {
				log.Fatal(err)
			}

		case 2: // Right Click
			err = d.VirtualTouchpad.RightClick()
			if err != nil {
				log.Fatal(err)
			}

		default:
			fmt.Printf("Unknown mouse action type: %d\n", mouseInfo.Type)
		}
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}

func main() {
	bounds := screenshot.GetDisplayBounds(0)
	display.X = int32(bounds.Dx())
	display.Y = int32(bounds.Dy())

	vtp, err := uinput.CreateTouchPad("/dev/uinput", []byte("VirtualTouchpad"), 0, display.X, 0, display.Y)
	if err != nil {
		log.Fatal(err)
	}
	display.VirtualTouchpad = vtp
	defer vtp.Close()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", display.echoHandler)
	log.Println("server starting on :8080")

	err = http.ListenAndServe(":8080", nil)
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
