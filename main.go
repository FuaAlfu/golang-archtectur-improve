package main

import(
	"fmt"
	"encoding/base64"
	_"log"
	_"encoding/json"
	_"net/http"
)

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
}