package main

import (
	"fmt"
	"io/ioutil"

	"github.com/fedekunze/alien_task/cosmos"
)

// Init initializes the battle of aliens according to the provided arguments
// in the CLI
func Init(totalAliens int) {
	var gal = cosmos.CreateGalaxy()

	// var totalCities =
	// // select a random city from the set to add one
	// for index := 0; index < totalAliens; index++ {
	// 	var randCity = rand.Intn(totalCities)
	// 	// cities[randCity].
	// }
}

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

// ConcatRoads concatenates roads into the desired output format for printing results
func ConcatRoads(road *cosmos.Road, line string) string {
	var direction = road.GetDirection()
	var destination = road.Destination().Name()
	return line + " " + direction + "=" + destination
}

// PrettyPrint prints the state of the cosmos
func PrettyPrint(galaxy cosmos.Galaxy) {
	for _, city := range galaxy.cities {
		var newline = city.Name()
		for dir := 0; dir < 4; dir++ {
			var road = city.roads[dir]
			newline = ConcatRoads(road, newline)
		}
		fmt.Println(newline)
	}
}
