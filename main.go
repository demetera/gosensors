package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

const temp_file string = "ftemp.txt"
const fan_file string = "ffan.txt"
//const fan_file string = "/sys/class/hwmon/hwmon1/fan1_input"

type Data struct {
	Temp int	`json:"temp"`
	Fan int		`json:"fan"`
}

var hwData Data

func update (w http.ResponseWriter, h *http.Request) {
	hwData.Temp = ReadValue(temp_file, 1)
	hwData.Fan = ReadValue(fan_file, 2)
	out, _ := json.MarshalIndent(hwData, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func hp(w http.ResponseWriter, h *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
	log.Println(h.RemoteAddr)
}

func main() {
	log.Println("Starting server on port 8080")
	http.HandleFunc("/", hp)
	http.HandleFunc("/update", update)
	http.ListenAndServe(":8080", nil)
}
