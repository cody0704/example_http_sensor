package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	http.HandleFunc("/test", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Println(err)
	}

	var res map[string]interface{}
	res = make(map[string]interface{})

	res["CPU Counts"] = cpuInfo

	resJson, _ := json.Marshal(res)

	fmt.Fprintf(w, "%s", prtgFormat(resJson))
}

func prtgFormat(data []byte) string {
	return "0:GolangHttpCPU:" + string(data)
}
