//go:build cgo

package main

import "C"

import (
	"io"
	"net/http"
)

//export FetchData
func FetchData() *C.char {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT")
	if err != nil {
		return C.CString("Error fetching data")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return C.CString("Error reading response")
	}

	return C.CString(string(body))
}

func main() {}
