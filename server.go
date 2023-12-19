package main

import (
	"encoding/json"
	"net/http"
)

var jsonData = []byte(`
{
  "data": [
    ["Weeks", "Completed", "Inprogress", "Failed", "Rejected"],
    ["Mon", 600, 250, 210, 140],
    ["Tue", 470, 260, 250, 120],
    ["Wed", 660, 320, 230, 100],
    ["Thu", 730, 280, 235, 105],
    ["Fri", 930, 340, 265, 109],
    ["Sat", 530, 440, 378, 210],
    ["Sun", 830, 240, 150, 95]
  ]
}
`)

func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/api/data/sms", handleAPIRequest)

	port := "8081"
	addr := "localhost:" + port
	println("Listening on http://" + addr + "/api/data/sms")

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
