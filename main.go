package main

import (
	"log"
	"net/http"

	"github.com/dzotsee21/remram/internal"
	"github.com/dzotsee21/remram/utils"

	"github.com/bendahl/uinput"
	"github.com/kbinani/screenshot"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}

func main() {
	utils.CheckUinput()
	localIP := utils.GetLocalIP()

	bounds := screenshot.GetDisplayBounds(0)
	internal.Client.X = int32(bounds.Dx())
	internal.Client.Y = int32(bounds.Dy())

	vtp, err := uinput.CreateTouchPad("/dev/uinput", []byte("VirtualTouchpad"), 0, internal.Client.X, 0, internal.Client.Y)
	if err != nil {
		log.Fatal(err)
	}
	defer vtp.Close()

	vk, err := uinput.CreateKeyboard("/dev/uinput", []byte("VirtualKeyboard"))
	if err != nil {
		log.Fatal(err)
	}
	defer vk.Close()

	vm, err := uinput.CreateMouse("/dev/uinput", []byte("VirtualMouse"))
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Close()

	internal.Client.VirtualTouchpad = vtp
	internal.Client.VirtualKeyboard = vk
	internal.Client.VirtualMouse = vm

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", internal.Client.WsHandler)

	log.Printf("server starting at http://%s:8080\n", localIP)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
