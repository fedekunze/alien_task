package cosmos

import (
	"fmt"
	"math/rand"
)

// Simulate simulates a battle of aliens
func Simulate(m Map, destroyedCities int) error {
	var round = 0
	for destroyedCities > 0 {
		fmt.Println("––––––––––– Turn %v –––––––––––", round)
		// iterate over aliens
		for i, alien := range m.aliens {
			// check if alien is alive
			if alien.IsAlive() {
				var currentCity = alien.GetPosition()
				// select a valid direction to move from alien current city
				var availableRoads = currentCity.roads.AvailableRoads()
				var dir = rand.Intn(len(availableRoads))
				var direction = availableRoads[dir].GetDirection()
				// move
				var dest, err = Move(alien, direction)
				if err != nil {
					return err
				}
				// check if fight
				if dest.HasFight() {
					Fight(i, dest)
					destroyedCities--
				}
			}
		}
	}
	return nil
}

// Move moves the alien from origin to a random destination if there's a path between them
func Move(alien *Alien, direction Direction) (*City, error) {
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
	} else if currentCity.Name() == destination.Name() {
		return nil, fmt.Errorf("Loopy path from and to " + currentCity.Name())
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
	alien.setPosition(*destination)
	return destination, nil
}

// RemovePaths removes all the paths from the neighbour cities
func RemovePaths(city *City) bool {
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
func Fight(alienID int, city *City) error {
	// TODO Print fight between alienID and aliens in city
	// Remove paths
	RemovePaths(city)
	var err = city.roads.DestroyAll() // destroy all roads
	if err != nil {
		return err
	}
	city.aliens.KillAll() // kill all aliens in the city
	city.destroyed = true // set state of city to destroyed
	// TODO return error if city was already destroyed or if aliens where destroyed
	return nil
}
