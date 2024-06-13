package main

import "fmt"

func main () {
	var cnc = make(map[string]string)
	cnc["France"] = "Paris"
	cnc["Germany"] = "Berlin"
	cnc["Italy"] = "Rome"

	fmt.Println(cnc["France"])
	
}
