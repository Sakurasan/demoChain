package main

import (
	"blockChain/core"
	"encoding/json"
	"net/http"
)

var bc = core.NewBlockChain()

func main() {

	// bc := core.NewBlockChain()
	// bc.SendData("爱")
	// bc.SendData("你")

	http.HandleFunc("/", readHandler)
	http.HandleFunc("/w", writeHandler)
	http.ListenAndServe(":8080", nil)
	// bc.Print()

}

func readHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		jsonBlockChain, _ := json.Marshal(bc)
		w.Write(jsonBlockChain)
	}
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("data")
	bc.SendData(data)
	readHandler(w, r)
}
