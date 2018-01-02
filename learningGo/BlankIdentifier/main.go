package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// "_" is the blank identifier which tells Go to dump the error
	// it tells the compiler something in the code is not being used.
	res, _ := http.Get("https://wanderersbook.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
