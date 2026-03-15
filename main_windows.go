//go:build windows

package main

import (
	"fmt"
	"net/http"

	"github.com/dzotsee21/remram/internal"
	"github.com/kbinani/screenshot"
)

func setupPlatform() {
	fmt.Println("USING WINDOWS")
	bounds := screenshot.GetDisplayBounds(0)
	internal.WindowsClient.X = int32(bounds.Dx())
	internal.WindowsClient.Y = int32(bounds.Dy())

	http.HandleFunc("/ws", internal.WindowsClient.RobotGoWsHandler)
}
