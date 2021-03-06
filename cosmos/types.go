package cosmos

import (
	"fmt"
	"strconv"
	"strings"
)

// ========== Cosmos ==========

// Map represents the overall map structure with cities and aliens
type Map struct {
	cities       map[string]*City // Array of cities in the map
	Aliens       Aliens           // Map of aliens in the map.
	CitiesIDName map[int]string
}

// CreateMap creates a new Galaxy
func CreateMap() *Map {
	return &Map{
		cities:       make(map[string]*City),
		Aliens:       InitAliens(),
		CitiesIDName: make(map[int]string),
	}
}

// GetCity Gets a city from cities mapping
func (m Map) GetCity(name string) (*City, error) {
	var city, ok = m.cities[name]
	if !ok {
		return nil, fmt.Errorf("Couldn't find city %v", name)
	}
	return city, nil
}

// SetCity Sets a city to mapping
func (m *Map) SetCity(city *City) {
	m.cities[city.Name()] = city
}

// CitiesLen return the total amount of cities in the map
func (m Map) CitiesLen() int {
	return len(m.cities)
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
func NewCity(name string) *City {
	return &City{
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

// GetRoads returns all roads from the city
func (city City) GetRoads() Roads {
	return city.roads
}

// GetRoad returns a pointer to the road in the desired direction
func (city City) GetRoad(i int) (*Road, error) {
	if i < 4 {
		return city.roads[i], nil
	}
	return nil, fmt.Errorf("Invalid direction")
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
func (city *City) AddAlien(alien *Alien) error {
	if !alien.IsAlive() {
		return fmt.Errorf("Alien is not alive")
	}
	city.aliens.Set(alien.ID(), alien)
	return nil
}

// RemoveAlien removes an alien from the set of aliens
func (city *City) RemoveAlien(id int) error {
	err := city.aliens.remove(id)
	return err
}

// AddRoad adds a new road to the city
func (city *City) AddRoad(road *Road) error {
	roads, err := city.roads.AddRoad(road)
	if err == nil {
		city.roads = roads
	}
	return err
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
func (roads Roads) AddRoad(road *Road) (Roads, error) {
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
		return roads, fmt.Errorf("Invalid direction: %v", dir)
	}
	return roads, nil
}

// AvailableRoads filters all roads that are not destroyed from a set
func (roads Roads) AvailableRoads() int {
	avaliable := 0
	for i := 0; i < 4; i++ {
		if roads[i] != nil && roads[i].IsAvailable() {
			avaliable++
		}
	}
	return avaliable
}

// Destroy a single road in a given direction
func (roads Roads) Destroy(dir Direction) (Roads, error) {
	switch dir {
	case North:
		err := roads[0].Destroy()
		return roads, err
	case South:
		err := roads[1].Destroy()
		return roads, err
	case East:
		err := roads[2].Destroy()
		return roads, err
	case West:
		err := roads[3].Destroy()
		return roads, err
	default:
		return roads, fmt.Errorf("Invalid direction")
	}
}

// DestroyAll destroys all the roads of a city
func (roads Roads) DestroyAll() error {
	for i := 0; i < 4; i++ {
		if roads[i] == nil {
			continue
		}
		var err = roads[i].Destroy()
		if err != nil {
			return err
		}
		// roads[i] = nil
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

func (dir Direction) Value() (string, error) {
	switch dir {
	case North:
		return "north", nil
	case South:
		return "south", nil
	case East:
		return "east", nil
	case West:
		return "west", nil
	case Destroyed:
		return "", nil
	default:
		return "", fmt.Errorf("%v is not a valid direction", dir)
	}
}

// IntValue gets the corresponding integer value in the array of roads
func (dir Direction) IntValue() int {
	switch dir {
	case North:
		return 0
	case South:
		return 1
	case East:
		return 2
	case West:
		return 3
	default:
		return -1
	}
}

// StrToDir converts string to Direction type
func StrToDir(str string) (Direction, error) {
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	if str == "north" || str == " north" {
		return North, nil
	} else if str == "south" || str == " south" {
		return South, nil
	} else if str == "east" || str == " east" {
		return East, nil
	} else if str == "west" || str == " west" {
		return West, nil
	} else {
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
func NewRoad(cityA *City, dir Direction, cityB *City) *Road {
	return &Road{
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
func (road *Road) Destroy() error {
	if !road.IsAvailable() {
		return fmt.Errorf("Road is currently destroyed")
	}
	road.available = false
	return nil
}

// IsAvailable checks if the road is not destroyed
func (road Road) IsAvailable() bool {
	return road.available
}

// ========== Aliens ==========

// Aliens is a set of aliens
// NOTE: I used a map for constant addition and deletion of values
type Aliens map[int]*Alien

// InitAliens initializes an empty aliens map
// NOTE: the benefit of using init over make(map[int]*Alien) is that it returns
// a type Aliens struct
func InitAliens() Aliens {
	return make(map[int]*Alien)
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
	_, ok := aliens[i]
	if !ok {
		return false
	}
	return true
}

// Kill a given alien from the map
func (aliens Aliens) Kill(id int) (Aliens, error) {
	var alien, err = aliens.Get(id)
	if err != nil {
		return InitAliens(), err
	}
	err = alien.Kill()
	if err != nil {
		return InitAliens(), err
	}
	err = aliens.remove(id)
	if err != nil {
		return InitAliens(), err
	}
	return aliens, nil
}

// KillAll kills every alien on the mapping
func (aliens Aliens) KillAll() error {
	var err error
	for i := 0; i < aliens.Len(); i++ {
		aliens, err = aliens.Kill(i)
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

// Set value of alien
func (aliens Aliens) Set(i int, alien *Alien) {
	aliens[i] = alien
}

// remove the pointer to the alien in the mapping and removes the entry
func (aliens Aliens) remove(i int) error {
	if !aliens.Exists(i) {
		return fmt.Errorf("Alien %v not found", i)
	}
	delete(aliens, i)
	return nil
}

// ========== Alien ==========

// Alien structure
type Alien struct {
	id       int
	position *City // current City where the alien is standing
	alive    bool
}

// NewAlien creates a new alien
func NewAlien(i int, city *City) *Alien {
	return &Alien{
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
func (alien Alien) GetPosition() *City {
	return alien.position
}

// GetPosition of the city
func (alien *Alien) setPosition(city *City) error {
	var name = city.Name()
	if alien.position.Name() == name {
		return fmt.Errorf("Alien " + strconv.Itoa(alien.ID()) + " is already in city " + name)
	}
	alien.position = city
	return nil
}

// Kill turns the alien into a dead alien
func (alien *Alien) Kill() error {
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
