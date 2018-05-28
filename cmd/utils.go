package cmd

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/fedekunze/alien_task/cosmos"
)

// Init initializes the battle of aliens according to the provided arguments
// in the CLI
func Init(filename string, totalAliens int) error {
	var m = cosmos.CreateMap()
	fmt.Println("Reading file...")
	fmt.Println()
	err := ReadMap(filename, m)
	if err != nil {
		return err
	}
	fmt.Println("Placing aliens in cities...")

	// // select a random city from the set to add one
	for index := 0; index < totalAliens; index++ {
		randCity := rand.Intn(m.CitiesLen())
		cityName := m.CitiesIDName[randCity]
		city, err := m.GetCity(cityName)
		if err != nil {
			return nil
		}
		alien := cosmos.NewAlien(index, *city)
		city.AddAlien(alien)
		m.Aliens.Set(index, alien)
	}

	// 	//
	fmt.Println("Running simulation...")
	err = cosmos.Simulate(m, totalAliens)
	if err != nil {
		return err
	}
	fmt.Println("Simulation ended")
	fmt.Println("All aliens where destroyed. Printing results:")
	fmt.Println()
	PrettyPrint(*m)
	return nil
}

// ParseLine parses each line from the file and creates a city
func ParseLine(line string, m *cosmos.Map) error {
	words := strings.Split(line, " ")
	city, err := m.GetCity(words[0])
	if err != nil {
		city := cosmos.NewCity(words[0])
		m.SetCity(city)
		m.CitiesIDName = append(m.CitiesIDName, words[0])
	}
	for i := 1; i < len(words); i++ {
		path := strings.Split(words[i], "=")
		dir, err := cosmos.StrToDir(path[0])
		if err != nil {
			return err
		}
		// check if city with name == path[1] exists
		destCity, err := m.GetCity(path[1])
		if err != nil {
			destCity = cosmos.NewCity(path[1])
			m.SetCity(destCity)
			m.CitiesIDName = append(m.CitiesIdName, path[1])
		}
		road := cosmos.NewRoad(*city, dir, destCity)
		err = city.AddRoad(*road)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadMap reads a map from a .txt file
func ReadMap(filename string, m *cosmos.Map) error {
	// Check if file is txt
	var extension = filepath.Ext(filename)
	if !strings.Contains(extension, ".txt") {
		return fmt.Errorf("File %v does not have .txt format", filename)
	}
	// Get filename from absolute path
	if filepath.IsAbs(filename) {
		_, filename = filepath.Split(filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close() // closes file on return

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// This is our buffer now
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println(line)
		err = ParseLine(line, m)
		if err != nil {
			return err
		}
	}
	fmt.Println()
	return nil
}

// ConcatRoads concatenates roads into the desired output format for printing results
func ConcatRoads(road *cosmos.Road, line string) string {
	var direction = road.GetDirection()
	var destination = road.Destination().Name()
	strValue, _ := direction.Value()
	return line + " " + strValue + "=" + destination
}

// PrettyPrint prints the state of the cosmos
func PrettyPrint(m cosmos.Map) {
	for i := 0; i < m.CitiesLen(); i++ {
		newline := m.CitiesIDName[i]
		city, _ := m.GetCity(newline)
		for dir := 0; dir < 4; dir++ {
			road, _ := city.GetRoad(dir)
			newline = ConcatRoads(road, newline)
		}
		fmt.Println(newline)
	}
}
