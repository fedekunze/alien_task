package cosmos

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Simulate simulates a battle of aliens
func Simulate(m *Map, aliensLeft int) error {
	var round = 0 // number of times all the aliens have moved in the map

	// Iterate over aliens until all of them are dead or
	// each​ ​alien​ ​has​ ​moved​ ​at​ ​least​ ​10,000​ ​times
	rand.Seed(time.Now().Unix())
	for aliensLeft > 0 || round < 10000 {
		fmt.Println()
		fmt.Println("––––––––––– Round " + strconv.Itoa(round) + " –––––––––––")
		fmt.Println()
		for i, alien := range m.Aliens {
			// check if alien is alive
			if alien.IsAlive() {
				currentCity := alien.GetPosition()
				// select a valid direction to move from alien current city
				if &currentCity == nil {
					return fmt.Errorf("Alien hasn't been placed")
				}
				selectedRoad, err := currentCity.GetRoad(rand.Intn(4))
				for selectedRoad == nil {
					selectedRoad, err = currentCity.GetRoad(rand.Intn(4))

				}
				if !selectedRoad.IsAvailable() {
					return fmt.Errorf("City %v is destroyed", currentCity.Name())
				}
				direction := selectedRoad.GetDirection()
				if direction == Destroyed {
					return fmt.Errorf("City %v is destroyed", currentCity.Name())
				}
				intDir := direction.IntValue()
				// move
				dest, err := move(alien, intDir)
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
func move(alien *Alien, direction int) (*City, error) {
	// get destination city from the destination value
	var currentCity = alien.GetPosition()
	var road, err = currentCity.GetRoad(direction)
	if err != nil {
		return nil, err
	}
	var destination = road.Destination()
	if !road.IsAvailable() || destination == nil {
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
	err = alien.setPosition(destination)
	if err != nil {
		return nil, err
	}
	return destination, nil
}

// RemovePaths removes all the paths from the neighbour cities
func removePaths(city *City) error {
	for i := 0; i < 4; i++ {
		if city.roads[i] != nil {
			opositeDir := city.roads[i].OppositeDirection()
			destCity := city.roads[i].Destination()
			if destCity == nil {
				return fmt.Errorf("Destination city does not exist")
			}
			destRoads := destCity.GetRoads()
			var roads, res = destRoads.Destroy(opositeDir)
			if !res {
				return fmt.Errorf("Invalid direction")
			}
			destCity.roads = roads
		}
	}
	return nil
}

// Fight destroys all the roads of the city and its aliens and
// sets the state to destroyed
func fight(alienID int, city *City) error {
	_, Err := city.aliens.Get(alienID)
	if Err != nil {
		return Err
	}
	// Print fight between alienID and aliens in city
	for i, alien := range city.aliens {
		if alien.ID() == alienID {
			continue
		}
		var msg = city.Name() + " ​has​ ​been​ ​destroyed​ ​by​ ​alien " + strconv.Itoa(alienID) +
			"​ ​and​ ​alien​ " + strconv.Itoa(alien.ID()) + "!"
		fmt.Println(msg)
		aliens, err := city.aliens.Kill(i) // destroy each alien in the city
		if err != nil {
			return err
		}
		city.aliens = aliens
	}
	aliens, err := city.aliens.Kill(alienID) // destroy the moving alien
	if err != nil {
		return err
	}
	city.aliens = aliens
	// Remove paths from neighbour cities
	err = removePaths(city)
	if err != nil {
		return fmt.Errorf("Couldn't delete destination roads")
	}
	err = city.roads.DestroyAll() // destroy all roads
	if err != nil {
		return err
	}
	city.destroyed = true // set state of city to destroyed
	return nil
}
