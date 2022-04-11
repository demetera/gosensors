package main

import (
	"html/template"
	"log"
	"net/http"
)

const temp_file string = "ex.txt"
//const fan_file string = "/sys/class/hwmon/hwmon1/fan1_input"

type Data struct {
	Temp int
}

func hp(w http.ResponseWriter, h *http.Request) {
	// Temper := ReadValue(temp_file, 1)
	mData := Data{34}	
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, mData)
	log.Println(h.RemoteAddr)

}

func main() {
	log.Println("Starting server on port 8080")
	http.HandleFunc("/", hp)
	http.ListenAndServe(":8080", nil)
}
