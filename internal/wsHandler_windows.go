//go:build windows

package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-vgo/robotgo"
)

type WindowsDisplay struct {
	X     int32 `json:"x"`
	Y     int32 `json:"y"`
	Shift bool
}

var WindowsClient WindowsDisplay

func (d *WindowsDisplay) RobotGoWsHandler(w http.ResponseWriter, r *http.Request) {
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

		var inputInfo InputInfo
		err = json.Unmarshal(p, &inputInfo)
		if err != nil {
			log.Println("unmarshal error: ", err)
			continue
		}

		switch inputInfo.Type {
		case 0: // Move Mouse Position
			robotgo.Move(int(inputInfo.X), int(inputInfo.Y))

		case 1: // Left Click
			robotgo.Click("left", false)

		case 2: // Right Click
			robotgo.Click("right", false)

		case 3: // Key Press
			robotKey := KeyToRobotGo[inputInfo.Key]

			if robotKey == "shift" && !WindowsClient.Shift {
				WindowsClient.Shift = true
				robotgo.KeyDown("shift")
			} else {
				robotgo.KeyTap(robotKey)

				if WindowsClient.Shift {
					WindowsClient.Shift = false
					robotgo.KeyUp("shift")
				}
			}

		case 4: // Left Press
			robotgo.MouseDown("left")

		case 5: // Left Release
			robotgo.MouseUp("left")

		case 6: // Scroll
			var direction string

			if inputInfo.Horizontal {
				if inputInfo.Delta >= 1 {
					direction = "left"
				} else {
					direction = "right"
				}
			} else {
				if inputInfo.Delta >= 1 {
					direction = "up"
				} else {
					direction = "down"
				}
			}
			robotgo.ScrollDir(int(inputInfo.Delta), direction)

		default:
			fmt.Printf("Unknown action type: %d\n", inputInfo.Type)
		}
	}
}
