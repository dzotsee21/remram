package main

import (
	"log"
	"net/http"

	"github.com/dzotsee21/remram/utils"
)

func main() {
	localIP := utils.GetLocalIP()

	setupPlatform()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	log.Printf("server starting at http://%s:8080\n", localIP)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
