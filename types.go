package app

import (
	"reflect"
)

// City definition
type City struct {
	name      string
	roads     Roads
	aliens    Aliens
	destroyed bool
}

// NewCity creates a new city
func NewCity(name string) City {
	return City{
		name:      name,
		roads:     nil,
		aliens:    nil,
		destroyed: false,
	}
}

// Destroy destroys all the roads of the city and its aliens and
// sets the state to destroyed
func (city City) Destroy() bool {
	city.roads.Destroy()  // destroy all roads
	city.aliens.Kill()    // kill all aliens in the city
	city.destroyed = true // set state of city to destroyed
	return true
}

// IsDestroyed returns the current status of the city
func (city City) IsDestroyed() bool {
	return city.destroyed
}

// CountAliens returns the total amount of aliens for the given city
func (city City) CountAliens() int {
	return city.aliens.Len()
}

// HasFight checks if there's a fight in the current move
// A fight happens if there's more than 2 aliens in the same city
func (city City) HasFight() bool {
	var totalAliens = city.CountAliens()
	return totalAliens > 1
}

// Roads is a set of roads
type Roads []Road

// Destroy destroys all the roads of a city
func (roads Roads) Destroy() bool {
	for _, road := range roads {
		var isDestroyed = road.Destroy()
		if !isDestroyed {
			return false
		}
	}
	return true
}

// Direction nolint is one of [N, S, E, W]
type Direction string

const (
	// nolint
	North     Direction = "North"
	South     Direction = "South"
	East      Direction = "East"
	West      Direction = "West"
	Destroyed Direction = ""
)

// Road struct definition
type Road struct {
	cityA     City
	direction Direction
	cityB     City
}

// GetDirection gets the current direction of the road (i.e. from cityA to cityB)
func (road Road) GetDirection() string {
	switch road.direction {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"
	case Destroyed:
		return ""
	default:
		errMsg := "Unrecognized Direction type: " + reflect.TypeOf(road.direction).Name()
		return errMsg
	}
}

// OppositeDirection gets the opossite direction of the road (i.e. from cityB to cityA)
func (road Road) OppositeDirection() Direction {
	switch road.direction {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		return Destroyed
	}
}

// Destroy destroys the road
func (road Road) Destroy() bool {
	road.direction = Destroyed
	return true
}

// Aliens is a set of aliens
type Aliens []Alien

// Kill kills every alien
func (aliens Aliens) Kill() bool {
	for _, alien := range aliens {
		var isDead = alien.Kill()
		if !isDead {
			return false
		}
	}
	return true
}

// Len returns the total amount of aliens
func (aliens Aliens) Len() int {
	return len(aliens)
}

// Alien structure
type Alien struct {
	position City // current City where the alien is standing
	alive    bool
}

// NewAlien creates a new alien
func NewAlien(city City) Alien {
	return Alien{
		position: city,
		alive:    true,
	}
}

// TODO
// Move moves the alien from city A to city B if there's a path between them
func (alien Alien) Move(direction Direction) bool {
	// remove the alien from cityA
	// Add the alien to cityB's aliens
	// update this alien position
	return true
}

// Kill turns the alien into a dead alien
func (alien Alien) Kill() bool {
	alien.alive = false
	return !alien.alive
}

// IsAlive gets the current status of the alien
func (alien Alien) IsAlive() bool {
	return alien.alive
}
