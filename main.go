package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bendahl/uinput"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	vm, err := uinput.CreateMouse("/dev/uinput", []byte("VirtualMouse"))
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Close()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade err: ", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error: ", err)
			break
		}

		if string(p) == "s" {
			fmt.Println("HELLO")

			vm.Move(100, 100)
		}
		fmt.Printf("received message: %s\n", p)

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			log.Println("write error: ", err)
			break
		}

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
