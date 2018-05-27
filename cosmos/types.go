package cosmos

import (
	"fmt"
	"strconv"
	"strings"
)

// ========== Cosmos ==========

// Map represents the overall map structure with cities and aliens
type Map struct {
	cities []*City // Array of cities in the map
	aliens Aliens  // Map of aliens in the map.
}

// CreateMap creates a new Galaxy
func CreateMap() *Map {
	return &Map{
		cities: []*City{},
		aliens: InitAliens(),
	}
}

// GetCity Gets a city from cities mapping
func (m *Map) GetCity(i int) (*City, error) {
	if i < len(m.cities) {
		return m.cities[i], nil
	}
	return nil, fmt.Errorf("Couldn't find city with id %v", i)
}

// SetCity Sets a city to mapping
func (m *Map) SetCity(city *City) {
	m.cities = append(m.cities, city)
}

// ========== City ==========

// City struct definition
type City struct {
	name      string // name of the city
	roads     Roads  // map of road structs
	aliens    Aliens // map of alien structs
	destroyed bool   // boolean to check if the city is destroyed
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

// Name returns the name of the city
func (city City) Name() string {
	return city.name
}

// GetRoad returns a pointer to the road in the desired direction
func (city City) GetRoad(direction Direction) (*Road, error) {
	switch direction {
	case North:
		return city.roads[0], nil
	case South:
		return city.roads[1], nil
	case East:
		return city.roads[2], nil
	case West:
		return city.roads[3], nil
	default:
		return nil, fmt.Errorf("Invalid direction: %v", direction)
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
func (city City) AddAlien(alien *Alien) error {
	if !alien.IsAlive() {
		return fmt.Errorf("Alien is not alive")
	}
	var id = alien.ID()
	var res = city.aliens.Set(id, alien)
	return res
}

// RemoveAlien removes an alien from the set of aliens
func (city City) RemoveAlien(id int) error {
	return city.aliens.remove(id)
}

// ========== Roads ==========

// Roads is a set of 4 roads, one for each direction:
// North, South, East, West
type Roads [4]*Road

// InitRoads initializes an empty road array of size 4
func InitRoads() Roads {
	return [4]*Road{}
}

// AddRoad adds a road to its corresponding position in the array
func (roads Roads) AddRoad(road *Road) error {
	var dir = road.GetDirection()
	if dir == North {
		roads[0] = road
	} else if dir == South {
		roads[1] = road
	} else if dir == East {
		roads[2] = road
	} else if dir == West {
		roads[3] = road
	} else {
		return fmt.Errorf("Invalid direction: %v", dir)
	}
	return nil
}

// AvailableRoads filters all roads that are not destroyed from a set
func (roads Roads) AvailableRoads() []*Road {
	var availableRoads = []*Road{}
	for _, road := range roads {
		if road.Available() {
			availableRoads = append(availableRoads, road)
		}
	}
	return availableRoads
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
func (roads Roads) DestroyAll() error {
	for _, road := range roads {
		var err = road.Destroy()
		if err != nil {
			return err
		}
		road = nil
	}
	return nil
}

// ========== Direction ==========

// Direction is one of [N, S, E, W]
type Direction string

const (
	// North direction
	North Direction = "north"
	// South direction
	South Direction = "south"
	// East direction
	East Direction = "east"
	// West direction
	West Direction = "west"
	// Destroyed city is evaluated as an empty string
	Destroyed Direction = ""
)

// StrToDir converts string to Direction type
func StrToDir(str string) (Direction, error) {
	str = strings.ToLower(str)
	switch str {
	case "north":
		return North, nil
	case "south":
		return South, nil
	case "east":
		return East, nil
	case "west":
		return West, nil
	case "":
		return Destroyed, nil
	default:
		return "", fmt.Errorf("String %v is not a valid direction", str)
	}
}

// ========== Road ==========

// Road struct definition
type Road struct {
	origin      *City     // origin city
	direction   Direction // direction is one of the (N, S, E, W) directions
	destination *City     // destination city
	available   bool      // available road to move
}

// NewRoad creates a new Road instance
func NewRoad(cityA *City, dir Direction, cityB *City) Road {
	return Road{
		origin:      cityA,
		direction:   dir,
		destination: cityB,
		available:   true,
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
func (road Road) GetDirection() Direction {
	return road.direction
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
func (road Road) Destroy() error {
	if !road.Available() {
		return fmt.Errorf("Road is currently destroyed")
	}
	road.available = false
	return nil
}

// Available checks if the road is not destroyed
func (road Road) Available() bool {
	if road.available {
		return true
	}
	return false
}

// ========== Aliens ==========

// Aliens is a set of aliens
// NOTE: I used a map for constant addition and deletion of values
type Aliens []*Alien

// InitAliens initializes an empty aliens map
// NOTE: the benefit of using init over make(map[int]*Alien) is that it returns
// a type Aliens struct
func InitAliens() Aliens {
	return []*Alien{}
}

// Set adds a new alien to the map
func (aliens Aliens) Set(i int, alien *Alien) error {
	if aliens.Exists(i) {
		return fmt.Errorf("Alien %v already exists", i)
	}
	aliens[i] = alien
	return nil
}

// Get alien value from the map
func (aliens Aliens) Get(i int) (*Alien, error) {
	if aliens.Exists(i) {
		return aliens[i], nil
	}
	return nil, fmt.Errorf("Couldn't find alien with id %v", i)
}

// Exists checks if a given alien is on the map
func (aliens Aliens) Exists(i int) bool {
	if i < len(aliens) {
		return true
	}
	return false
}

// Kill a given alien from the map
func (aliens Aliens) Kill(id int) error {
	var alien, err = aliens.Get(id)
	if err != nil {
		return err
	}
	err = alien.Kill()
	if err != nil {
		return err
	}
	err = aliens.remove(id)
	if err != nil {
		return err
	}
	return nil
}

// KillAll kills every alien on the mapping
func (aliens Aliens) KillAll() error {
	for i := 0; i < aliens.Len(); i++ {
		var err = aliens.Kill(i)
		if err != nil {
			return err
		}
	}
	return nil
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
func (aliens Aliens) remove(i int) error {
	if !aliens.Exists(i) {
		return fmt.Errorf("Alien %v not found", i)
	}
	aliens.set(i, nil) // set the value of the mapping to that Alien to nil
	return nil
}

// ========== Alien ==========

// Alien structure
type Alien struct {
	id       int
	position City // current City where the alien is standing
	alive    bool
}

// NewAlien creates a new alien
func NewAlien(i int, city City) Alien {
	return Alien{
		id:       i,
		position: city,
		alive:    true,
	}
}

// ID returns the id value of an alien
func (alien Alien) ID() int {
	return alien.id
}

// GetPosition of the city
func (alien Alien) GetPosition() City {
	return alien.position
}

// GetPosition of the city
func (alien Alien) setPosition(city City) error {
	var name = city.Name()
	if alien.position.Name() == name {
		return fmt.Errorf("Alien " + strconv.Itoa(alien.ID()) + " is already in city " + name)
	}
	alien.position = city
	return nil
}

// Kill turns the alien into a dead alien
func (alien Alien) Kill() error {
	if !alien.IsAlive() {
		return fmt.Errorf("Alien %v is already dead", alien.ID())
	}
	alien.alive = false
	return nil
}

// IsAlive gets the current status of the alien
func (alien Alien) IsAlive() bool {
	return alien.alive
}
