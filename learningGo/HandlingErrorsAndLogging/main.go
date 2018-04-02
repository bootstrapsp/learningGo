package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("generalLog.log")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {

	_, err := os.Open("no-File.txt")
	if err != nil {
		log.Println("err happened", err)
	}

}
