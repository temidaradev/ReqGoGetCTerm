// go:build cgo

package main

import "C"
import (
	"io/ioutil"
	"net/http"
)

//export FetchData
func FetchData() *C.char {
	resp, err := http.Get("https://api.example.com/data")
	if err != nil {
		return C.CString("Error fetching data")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return C.CString("Error reading response")
	}

	return C.CString(string(body))
}

func main() {}
