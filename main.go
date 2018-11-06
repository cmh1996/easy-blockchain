package main

import (
	"blockchain/core"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var blockchain *core.BlockChain

func run() {
	http.HandleFunc("/blockchain/get", blockchainGetHandle)
	http.HandleFunc("/blockchain/write", blockchainWriteHandle)
	fmt.Println("the server is running at locahost:easy")
	http.ListenAndServe("localhost:easy", nil)
}

func blockchainGetHandle(w http.ResponseWriter, r *http.Request) {
	bytes, error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func blockchainWriteHandle(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandle(w, r)
}

func main() {
	blockchain = core.GenBlockChain("genesis data")
	run()
}
