package main

import (
	"log"
	"net/http"
)

func main() {
	const port = ":5500"
	const dir = "./requests"

	fs := http.FileServer(http.Dir(dir))

	log.Println("📄 LINSITrack API Docs server")
	log.Println("📂 Serving directory:", dir)
	log.Println("🌐 URL: http://localhost" + port)
	log.Println("⛔ Press Ctrl+C to stop")

	err := http.ListenAndServe(port, fs)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
