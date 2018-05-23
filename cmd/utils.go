package main

import (
	"fmt"
	"io/ioutil"
)

// ReadMap reads a map from a .txt file
func ReadMap(file string) {
	// TODO
	// 1) full path to file
	// 2) get filename and check if file is txt
	b, err := ioutil.ReadFile(file) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	fmt.Println(str) // print the content as a 'string'
}
