//go:build linux

package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bendahl/uinput"
)

type Display struct {
	X               int32 `json:"x"`
	Y               int32 `json:"y"`
	Shift           bool
	VirtualTouchpad uinput.TouchPad
	VirtualMouse    uinput.Mouse
	VirtualKeyboard uinput.Keyboard
}

var LinuxClient Display

func (d *Display) UinputWsHandler(w http.ResponseWriter, r *http.Request) {
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
			log.Fatal(err)
		}

		switch inputInfo.Type {
		case 0: // Move Mouse Position
			fmt.Printf("%d-%d\n", inputInfo.X, inputInfo.Y)

			err = d.VirtualTouchpad.MoveTo(inputInfo.X, inputInfo.Y)
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
		case 3: // Key Press
			uinputKey := KeyToUinput[inputInfo.Key]

			// shift
			if uinputKey == uinput.KeyLeftshift && !LinuxClient.Shift {
				LinuxClient.Shift = true

				err = d.VirtualKeyboard.KeyDown(uinputKey)
				if err != nil {
					log.Fatal(err)
				}
			} else { // everything else
				err = d.VirtualKeyboard.KeyPress(uinputKey)
				if err != nil {
					log.Fatal(err)
				}

				if LinuxClient.Shift {
					LinuxClient.Shift = false
					err = d.VirtualKeyboard.KeyUp(uinput.KeyLeftshift)
					if err != nil {
						log.Fatal(err)
					}
				}
			}

		case 4: // Left Press
			err = d.VirtualTouchpad.LeftPress()
			if err != nil {
				log.Fatal(err)
			}
		case 5: // Left Release
			err = d.VirtualTouchpad.LeftRelease()
			if err != nil {
				log.Fatal(err)
			}
		case 6: // Scroll
			err = d.VirtualMouse.Wheel(inputInfo.Horizontal, inputInfo.Delta)
			if err != nil {
				log.Fatal(err)
			}

		default:
			fmt.Printf("Unknown mouse action type: %d\n", inputInfo.Type)
		}
	}
}
