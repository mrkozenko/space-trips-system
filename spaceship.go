package space_trips_system

import "time"

type Spaceship struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CountSeats int    `json:"countSeats"`
}

type SpaceTrips struct {
	Id           int       `json:"id"`
	FromPlanet   string    `json:"fromPlanet"`
	ToPlanet     string    `json:"toPlanet"`
	Price        float64   `json:"price"`
	DispatchDate time.Time `json:"dispatchDate"`
	SpaceshipId  int       `json:"spaceshipId"`
}
