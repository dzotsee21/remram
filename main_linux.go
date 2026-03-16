//go:build linux

package main

import (
	"fmt"
	"net/http"

	"github.com/bendahl/uinput"
	"github.com/dzotsee21/remram/internal"
	"github.com/dzotsee21/remram/utils"
	"github.com/kbinani/screenshot"
)

func setupPlatform() {
	fmt.Println("USING LINUX")
	utils.CheckUinput()
	bounds := screenshot.GetDisplayBounds(0)
	internal.LinuxClient.X = int32(bounds.Dx())
	internal.LinuxClient.Y = int32(bounds.Dy())

	vtp, _ := uinput.CreateTouchPad("/dev/uinput", []byte("VirtualTouchpad"), 0, internal.LinuxClient.X, 0, internal.LinuxClient.Y)
	vk, _ := uinput.CreateKeyboard("/dev/uinput", []byte("VirtualKeyboard"))
	vm, _ := uinput.CreateMouse("/dev/uinput", []byte("VirtualMouse"))

	internal.LinuxClient.VirtualTouchpad = vtp
	internal.LinuxClient.VirtualKeyboard = vk
	internal.LinuxClient.VirtualMouse = vm

	http.HandleFunc("/ws", internal.LinuxClient.UinputWsHandler)
}
