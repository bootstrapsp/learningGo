package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	* reviewing Rune
	* method 1
		Using Rune to fetch the int encoding value for the letter assigned to rune()
		Letter picked up is coming from the index chosen

	* method 2
		Using rune() to check the encoding value for the letter

	* method 3
		Uses for loop to print out the int value and priting out it's string value
*/
func startingOutWithRune() {
	// method 1
	word := "Hello"
	letter := rune(word[3])

	fmt.Println("Rune returning the ecoding value for letter assigned to word parameter", letter)

	// method 2
	blah := rune('H')
	fmt.Println(" Rune value for H", blah)
}

// method 3
func creatingBucket() {

	for i := 65; i <= 122; i++ {
		fmt.Println("int Value ", i, " - ", "Converting int to string", string(i), " - ", "putting 'em in bucket", i%12)
	}

}

// func for putting things in bucket

func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}

// function fetching remote data using HTTP.GET
func hashMapCreation() {

	// response and error is returned by method HTTP.GET
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	// checking for error
	if err != nil {
		log.Fatal(err)
	}

	// working with the response returned above

	// ReadAll is returning slice of byte, variable named below is just to reflect that, though it can be named anything

	byteSlice, err := ioutil.ReadAll(res.Body)

	// closing the response

	res.Body.Close()

	// checking for error
	if err != nil {
		log.Fatal(err)
	}
	// using "%s" to convert the byteslice response received above into string
	fmt.Printf("%s", byteSlice)

	// printing out the unformatted byte slices received above for fun
	fmt.Println(byteSlice)
}

// main method
func main() {
	startingOutWithRune()
	creatingBucket()
	n := hashBucket("Demo", 12)

	fmt.Println(n)

	hashMapCreation()
}
