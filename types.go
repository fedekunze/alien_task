package main

import (
	"reflect"
)

// ========== City ==========

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
		roads:     InitRoads(),
		aliens:    InitAliens(),
		destroyed: false,
	}
}

// GetRoad returns a pointer to the road in the desired direction
func (city City) GetRoad(direction Direction) *Road {
	switch direction {
	case North:
		return city.roads[0]
	case South:
		return city.roads[1]
	case East:
		return city.roads[2]
	case West:
		return city.roads[3]
	default:
		return nil
	}
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

// AddAlien adds an alient to the mapping
func (city City) AddAlien(alien *Alien) bool {
	if !alien.IsAlive() {
		return false
	}

	return true
}

// RemoveAlien removes an alien from the set
func (city City) RemoveAlien(i int) bool {
	return city.aliens.remove(i)
}

// ========== Roads ==========

// Roads is a set of 4 roads, one for each direction
type Roads [4]*Road

// InitRoads initializes an empty road array of size 4
func InitRoads() Roads {
	return [4]*Road{}
}

// Destroy a single road in a given direction
func (roads Roads) Destroy(dir Direction) bool {
	switch dir {
	case North:
		roads[0] = nil
		return true
	case South:
		roads[1] = nil
		return true
	case East:
		roads[2] = nil
		return true
	case West:
		roads[3] = nil
		return true
	default:
		return false
	}
}

// DestroyAll destroys all the roads of a city
func (roads Roads) DestroyAll() bool {
	for _, road := range roads {
		var isDestroyed = road.Destroy()
		if !isDestroyed {
			return false
		}
		road = nil
	}
	return true
}

// ========== Direction ==========

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

// ========== Road ==========

// Road struct definition
type Road struct {
	origin      *City
	direction   Direction
	destination *City
}

// NewRoad creates a new Road
func NewRoad(cityA *City, dir Direction, cityB *City) Road {
	return Road{
		origin:      cityA,
		direction:   dir,
		destination: cityB,
	}
}

// Origin city
func (road Road) Origin() *City {
	return road.origin
}

// Destination city
func (road Road) Destination() *City {
	return road.destination
}

// GetDirection gets the current direction of the road (i.e. from origin to destination)
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

// OppositeDirection gets the opossite direction of the road (i.e. from destination to origin)
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
	road.origin = nil
	road.destination = nil
	return true
}

// ========== Aliens ==========

// Aliens is a set of aliens
// NOTE: I used a map for constant addition and deletion of values
type Aliens map[int]*Alien

// InitAliens initializes an empty aliens map
func InitAliens() Aliens {
	return map[int]*Alien{}
}

// AddAlien adds a new alien to the map
func (aliens Aliens) AddAlien(i int, alien *Alien) bool {
	if aliens.Exists(i) {
		return false
	}
	aliens[i] = alien
	return true
}

// Get alien value from the map
func (aliens Aliens) Get(i int) *Alien {
	alien, ok := aliens[i]
	if !ok {
		return nil
	}
	return alien
}

// Exists checks if a given alien is on the map
func (aliens Aliens) Exists(i int) bool {
	_, exists := aliens[i]
	return exists
}

// KillAll kills every alien on the mapping
func (aliens Aliens) KillAll() bool {
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

// ----- Unexported functions -----

// set value of alien
func (aliens Aliens) set(i int, alien *Alien) bool {
	aliens[i] = alien
	return true
}

// remove the pointer to the alien in the mapping and removes the entry
func (aliens Aliens) remove(i int) bool {
	if !aliens.Exists(i) {
		return false
	}
	aliens.set(i, nil)
	delete(aliens, i)
	return true
}

// ========== Alien ==========

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

// GetPosition of the city
func (alien Alien) GetPosition() City {
	return alien.position
}

// GetPosition of the city
func (alien Alien) setPosition(city City) bool {
	alien.position = city
	return true
}

// Kill turns the alien into a dead alien
func (alien Alien) Kill() bool {
	alien.alive = false
	return true
}

// IsAlive gets the current status of the alien
func (alien Alien) IsAlive() bool {
	return alien.alive
}
