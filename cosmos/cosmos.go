package cosmos

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Simulate simulates a battle of aliens
func Simulate(m Map, aliensLeft int) error {
	var round = 0 // number of times all the aliens have moved in the map
	// Iterate over aliens until all of them are dead or
	// each​ ​alien​ ​has​ ​moved​ ​at​ ​least​ ​10,000​ ​times
	for aliensLeft > 0 || round < 10000 {
		fmt.Println("––––––––––– Round " + strconv.Itoa(round) + " –––––––––––")
		fmt.Println()
		for i, alien := range m.aliens {
			// check if alien is alive
			if alien.IsAlive() {
				var currentCity = alien.GetPosition()
				// select a valid direction to move from alien current city
				var availableRoads = currentCity.roads.AvailableRoads()
				var dir = rand.Intn(len(availableRoads))
				var direction = availableRoads[dir].GetDirection()
				// move
				var dest, err = move(alien, direction)
				if err != nil {
					return err
				}
				// check if there is more than one alien in the city to fight
				if dest.HasFight() {
					var aliensInCity = dest.aliens.Len()
					fight(i, dest)
					aliensLeft -= aliensInCity
				}
			}
		}
		round++
	}
	return nil
}

// Move moves the alien from origin to a random destination if there's a path between them
func move(alien *Alien, direction Direction) (*City, error) {
	// get destination city from the destination value
	var currentCity = alien.GetPosition()
	// TODO random select road
	var road, err = currentCity.GetRoad(direction)
	if err != nil {
		return nil, err
	}
	var destination = road.Destination()
	if !road.Available() || destination == nil {
		return nil, fmt.Errorf("Road to" + destination.Name() +
			" is already destroyed")
	}
	// remove the alien from origin City
	err = currentCity.RemoveAlien(alien.ID())
	if err != nil {
		return nil, err
	}
	// Add the alien to destination's aliens
	err = destination.AddAlien(alien)
	if err != nil {
		return nil, err
	}
	// update this alien position
	err = alien.setPosition(*destination)
	if err != nil {
		return nil, err
	}
	return destination, nil
}

// RemovePaths removes all the paths from the neighbour cities
func removePaths(city *City) bool {
	for _, road := range city.roads {
		var opositeDir = road.OppositeDirection()
		var destRoad = road.destination.roads
		var res = destRoad.Destroy(opositeDir)
		if !res {
			return false
		}
	}
	return true
}

// Fight destroys all the roads of the city and its aliens and
// sets the state to destroyed
func fight(alienID int, city *City) error {
	// Print fight between alienID and aliens in city
	for i, alien := range city.aliens {
		var msg = city.Name() + " ​has​ ​been​ ​destroyed​ ​by​ ​alien " + strconv.Itoa(alienID) +
			"​ ​and​ ​alien​ " + strconv.Itoa(alien.ID()) + "!"
		fmt.Println(msg)
		var err = city.aliens.Kill(i) // destroy each alien in the city
		if err != nil {
			return err
		}
	}
	var err = city.aliens.Kill(alienID) // destroy the moving alien
	if err != nil {
		return err
	}
	// Remove paths from neighbour cities
	removePaths(city)
	err = city.roads.DestroyAll() // destroy all roads
	if err != nil {
		return err
	}
	city.destroyed = true // set state of city to destroyed
	return nil
}
