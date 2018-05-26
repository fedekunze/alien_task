package main

import (
	"fmt"
	"io/ioutil"

	"github.com/fedekunze/alien_task/cosmos"
)

// Init initializes the battle of aliens according to the provided arguments
// in the CLI
func Init(totalAliens int) {
	var m = cosmos.CreateMap()

	// filestring = ReadMap()
	// parse map

	// var totalCities = len(cities)
	// var city = NewCity(name)
	// // select a random city from the set to add one
	// for index := 0; index < totalAliens; index++ {
	// TODO
	// 	var randCity = rand.Intn(totalCities)
	//  cities[randCity]
	// var alien = NewAlien(index, city)
	// m.aliens.Set(index, alien)
	// m.SetCity(city)

	// 	//
	var err = cosmos.Simulate(m, totalAliens)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All aliens where destroyed. Printing :")
		fmt.Println()
		PrettyPrint(m)
	}
}

// ReadMap reads a map from a .txt file
func ReadMap(file string) error {
	// TODO
	// 1) full path to file
	// 2) get filename and check if file is txt
	var b, err = ioutil.ReadFile(file) // just pass the file name
	if err != nil {
		return err
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
func PrettyPrint(m cosmos.Map) {
	for _, city := range m.cities {
		var newline = city.Name()
		for dir := 0; dir < 4; dir++ {
			var road = city.roads[dir]
			newline = ConcatRoads(road, newline)
		}
		fmt.Println(newline)
	}
}
