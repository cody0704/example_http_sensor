package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	// "github.com/shirou/gopsutil/mem"
)

func main() {
	http.HandleFunc("/test", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// vm, _ := mem.VirtualMemory()

	var res map[string]interface{}
	res = make(map[string]interface{})

	res["test"] = "Hello"

	resJson, _ := json.Marshal(res)

	fmt.Fprintf(w, "%s", resJson)
}
