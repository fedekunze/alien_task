package cosmos

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

// TODO

var totalAliens = 5

func TestFight(t *testing.T) {
	city := NewCity("Foo")
	alien1 := NewAlien(1, city)
	alien2 := NewAlien(2, city)
	city.AddAlien(alien1)
	city.AddAlien(alien2)
	err := fight(4, city, 1)
	assert.Error(t, err)
	err = fight(1, city, 2)
	assert.Nil(t, err)
}

func TestRemovePaths(t *testing.T) {
	city := NewCity("Foo")
	otherCity := NewCity("Bar")
	westDir, err := StrToDir("west")
	assert.Nil(t, err)
	eastDir, err := StrToDir("east")
	assert.Nil(t, err)
	roadEast := NewRoad(city, eastDir, otherCity)
	roadWest := NewRoad(otherCity, westDir, city)
	err = city.AddRoad(roadEast)
	assert.Nil(t, err)
	err = otherCity.AddRoad(roadWest)
	assert.Nil(t, err)
	assert.NotNil(t, city)
	err = removePaths(city)
	assert.Nil(t, err)
}

func TestMove(t *testing.T) {
	city := NewCity("Foo")
	otherCity := NewCity("Bar")
	westDir, err := StrToDir("west")
	eastDir, err := StrToDir("east")
	roadEast := NewRoad(city, eastDir, otherCity)
	roadWest := NewRoad(otherCity, westDir, city)
	err = city.AddRoad(roadEast)
	assert.Nil(t, err)
	err = otherCity.AddRoad(roadWest)
	alien := NewAlien(1, city)
	city.AddAlien(alien)
	pos, err := move(alien, 2)
	assert.Nil(t, err)
	assert.Equal(t, otherCity, pos)
	pos, err = move(alien, 3)
	assert.Nil(t, err)
	assert.Equal(t, city, pos)
}

func TestSimulate(t *testing.T) {
	m := CreateMap()
	westDir, err := StrToDir("west")
	assert.Nil(t, err)
	assert.Equal(t, West, westDir)
	eastDir, err := StrToDir("east")
	assert.Nil(t, err)
	assert.Equal(t, East, eastDir)
	for i := 0; i < totalAliens; i++ {
		city := NewCity("City" + strconv.Itoa(i))
		m.CitiesIDName[i] = "City" + strconv.Itoa(i)

		alien := NewAlien(i, city)
		err := city.AddAlien(alien)
		assert.Nil(t, err)
		// adds roads to the next and previous cities
		if i > 0 {
			prevCity, err := m.GetCity(m.CitiesIDName[i-1])
			assert.Nil(t, err)
			roadEast := NewRoad(prevCity, eastDir, city)
			err = prevCity.AddRoad(roadEast)
			assert.Nil(t, err)
			assert.NotNil(t, prevCity.roads[2])
			roadWest := NewRoad(city, westDir, prevCity)
			err = city.AddRoad(roadWest)
			assert.Nil(t, err)
			assert.NotNil(t, city.roads[3])
		}
		if i == totalAliens-1 {
			zeroCity, err := m.GetCity(m.CitiesIDName[0])
			assert.Nil(t, err)
			roadEast := NewRoad(city, eastDir, zeroCity)
			err = city.AddRoad(roadEast)
			assert.Nil(t, err)
			_, err = city.GetRoad(2)
			assert.Nil(t, err)
			assert.NotNil(t, city.roads[3])
			roadWest := NewRoad(zeroCity, westDir, city)
			err = zeroCity.AddRoad(roadWest)
			assert.Nil(t, err)
			assert.NotNil(t, zeroCity.roads[3])
			_, err = city.GetRoad(3)
			assert.Nil(t, err)
		}
		alien.setPosition(city)
		m.SetCity(city)
		m.Aliens.Set(i, alien)
	}
	err = Simulate(m, totalAliens)
	assert.Nil(t, err)
}
