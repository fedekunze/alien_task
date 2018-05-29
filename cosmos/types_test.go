package cosmos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

// Test City

func TestNewCity(t *testing.T) {
	city := NewCity("Foo")
	assert.Equal(t, city.Name(), "Foo")
	assert.NotNil(t, city.roads)
	assert.NotNil(t, city.aliens)
	assert.False(t, city.IsDestroyed())
}

func TestAddRoad(t *testing.T) {
	city := NewCity("Foo")
	otherCity := NewCity("Bar")
	westDir, err := StrToDir("west")
	assert.Nil(t, err)
	road := NewRoad(city, westDir, otherCity)
	err = city.AddRoad(road)
	assert.Nil(t, err)
	rroad, err := city.GetRoad(3)
	assert.Nil(t, err)
	assert.Equal(t, road, rroad)
	rroad, err = city.GetRoad(5) // invalid road
	assert.Nil(t, rroad)
	assert.Error(t, err)
	assert.Equal(t, city.roads, city.GetRoads())
}

func TestAddAlien(t *testing.T) {
	city := NewCity("Foo")
	alien1 := NewAlien(1, city)
	err := city.AddAlien(alien1)
	assert.Nil(t, err)
	alien, err := city.aliens.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, alien1, alien)
	alien2 := NewAlien(2, city)
	err = alien2.Kill()
	assert.Nil(t, err)
	err = city.AddAlien(alien2)
	assert.Error(t, err)

}

func TestRemoveAlien(t *testing.T) {
	city := NewCity("Foo")
	alien1 := NewAlien(1, city)
	err := city.AddAlien(alien1)
	assert.Nil(t, err)
	err = city.RemoveAlien(1)
	assert.Nil(t, err)
}

func TestHasFight(t *testing.T) {
	city := NewCity("Foo")
	alien1 := NewAlien(1, city)
	alien2 := NewAlien(2, city)
	err := city.AddAlien(alien1)
	assert.Nil(t, err)
	err = city.AddAlien(alien2)
	assert.Nil(t, err)
	assert.True(t, city.HasFight())
}

// Test Roads

func TestDestroyRoads(t *testing.T) {
	city := NewCity("Foo")
	otherCity := NewCity("Bar")
	anotherCity := NewCity("NYC")
	northDir, err := StrToDir("north")
	assert.Nil(t, err)
	southDir, err := StrToDir("south")
	assert.Nil(t, err)
	roads := InitRoads()
	road := NewRoad(city, northDir, otherCity)
	roadSouth := NewRoad(city, southDir, anotherCity)
	roads, err = roads.AddRoad(roadSouth)
	assert.Nil(t, err)
	assert.Equal(t, roadSouth, roads[1])
	roads, err = roads.AddRoad(road)
	assert.Nil(t, err)
	assert.Equal(t, road, roads[0])
	roads, err = roads.Destroy(northDir)
	assert.Nil(t, err)
	assert.Nil(t, roads[0])
	assert.NotNil(t, roads[1])
	// err = roads.DestroyAll()
	// assert.Nil(t, err)
}

// Test Road
// TODO all directions
func TestRoadDirections(t *testing.T) {
	city := NewCity("Foo")
	otherCity := NewCity("Bar")
	westDir, err := StrToDir("west")
	assert.Nil(t, err)
	eastDir, err := StrToDir("east")
	assert.Nil(t, err)
	EWroad := NewRoad(city, westDir, otherCity)
	assert.Equal(t, city, EWroad.Origin())
	assert.Equal(t, westDir, EWroad.GetDirection())
	assert.Equal(t, otherCity, EWroad.Destination())
	WEroad := NewRoad(otherCity, eastDir, city)
	edir := EWroad.OppositeDirection()
	assert.Equal(t, eastDir, edir)
	wdir := WEroad.OppositeDirection()
	assert.Equal(t, westDir, wdir)
}

func TestDestroy(t *testing.T) {
	city := NewCity("Foo")
	otherCity := NewCity("Bar")
	westDir, err := StrToDir("west")
	assert.Nil(t, err)
	road := NewRoad(city, westDir, otherCity)
	err = road.Destroy()
	assert.Nil(t, err)
	err = road.Destroy()
	assert.Error(t, err)
}

// Test Aliens

func TestSet(t *testing.T) {
	aliens := InitAliens()
	city := NewCity("Foo")
	alien := NewAlien(1, city)
	aliens.Set(alien.ID(), alien)
	alien1, err := aliens.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, alien, alien1)
}

func TestRemove(t *testing.T) {
	aliens := InitAliens()
	city := NewCity("Foo")
	alien := NewAlien(1, city)
	aliens.Set(alien.ID(), alien)
	aliens.remove(alien.ID())
	alien1, err := aliens.Get(1)
	assert.NotNil(t, err)
	assert.Nil(t, alien1)
}

// Test Alien

func TestAlienNewAlien(t *testing.T) {
	city := NewCity("Foo")
	alien := NewAlien(1, city)
	assert.Equal(t, alien.id, 1)
}

func TestAlienGetPosition(t *testing.T) {
	city := NewCity("Foo")
	alien := NewAlien(1, city)
	assert.Equal(t, city, alien.GetPosition())
}

func TestAliensetPosition(t *testing.T) {
	city := NewCity("Foo")
	alien := NewAlien(1, city)
	assert.Error(t, alien.setPosition(city))
	city = NewCity("Bar")
	assert.Nil(t, alien.setPosition(city))
}

func TestAlienIsAlive(t *testing.T) {
	city := NewCity("Foo")
	alien := NewAlien(1, city)
	assert.Equal(t, alien.alive, alien.IsAlive())
	assert.True(t, alien.alive)
	alien.Kill()
	assert.Equal(t, alien.alive, alien.IsAlive())
	assert.False(t, alien.alive)
}
