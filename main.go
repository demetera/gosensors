package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// const temp_file string = "ftemp.txt"
// const fan_file string = "ffan.txt"

const temp_file string = "/sys/class/hwmon/hwmon2/temp3_input"
const fan_file string = "/sys/class/hwmon/hwmon2/fan1_input"

type Data struct {
	Temp int `json:"temp"`
	Fan  int `json:"fan"`
}

var hwData Data

func update(w http.ResponseWriter, h *http.Request) {
	hwData.Temp = ReadValue(temp_file, 1)
	hwData.Fan = ReadValue(fan_file, 2)
	out, _ := json.MarshalIndent(hwData, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	log.Println("Refreshed: ", h.RemoteAddr)
}

func hp(w http.ResponseWriter, h *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
	log.Println("Connected:", h.RemoteAddr)
}

func test(w http.ResponseWriter, h *http.Request) {
	t := template.New("mytemp")
	t.Execute(w, nil)
	log.Println("Test:", h.RemoteAddr)
	q := h.URL.Query()
	_ = q
	if q["a"][0] == "123" {
		log.Println("GOT IT")
	}
}

func doNothing(w http.ResponseWriter, r *http.Request) {}

func main() {
	log.Println("Starting server on port 8081")
	m := http.NewServeMux()
	m.HandleFunc("/update", update)
	m.HandleFunc("/test", test)
	m.HandleFunc("/", hp)
	m.HandleFunc("/favicon.ico", doNothing)
	http.ListenAndServe(":8081", m)
}
