package utils

import (
	"fmt"
	"log"
	"net"
	"os"
)

func CheckUinput() {
	f, err := os.OpenFile("/dev/uinput", os.O_WRONLY, 0o660)
	if err != nil {
		if os.IsPermission(err) {
			log.Fatal("Error: Missing permissions for /dev/uinput. Run ./setup.sh and relog.")
		}
		log.Fatalf("Error: /dev/uinput not found or unreachable: %v", err)
	}
	f.Close()
}

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("Couldn't find ipv4.")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
