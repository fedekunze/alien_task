package main

// Move moves the alien from city A to city B if there's a path between them
func Move(alien Alien, direction Direction) bool {
	// get destination city from the destination value
	var currentCity = alien.GetPosition()
	var road = currentCity.GetRoad(direction)
	var destination = road.Destination()
	if destination == nil {
		return false // return false if the road is destroyed
	}
	// remove the alien from origin City
	// currentCity.RemoveAlien(&alien)

	// Add the alien to destination's aliens
	destination.AddAlien(&alien)

	// update this alien position
	alien.setPosition(*destination)
	return true
}

// Fight destroys all the roads of the city and its aliens and
// sets the state to destroyed
func Fight(city City) bool {
	city.roads.Destroy()  // destroy all roads
	city.aliens.KillAll() // kill all aliens in the city
	city.destroyed = true // set state of city to destroyed
	return true
}
