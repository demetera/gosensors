package main

import (
	"fmt"
	"log"
	"net/http"
)

const temp_file string = "/sys/class/hwmon/hwmon1/temp3_input"
const fan_file string = "/sys/class/hwmon/hwmon1/fan1_input"

func hp(w http.ResponseWriter, h *http.Request) {
	log.Println(h.RemoteAddr)
	t := ReadValue(temp_file, 1)
	f := ReadValue(fan_file, 2)

	fmt.Fprintf(w, "<p>%d<br>%d</p>", t, f)
}

func main() {
	http.HandleFunc("/", hp)
	http.ListenAndServe(":8080", nil)
}
