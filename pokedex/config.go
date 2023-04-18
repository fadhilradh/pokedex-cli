package pokedex

type Config struct {
	NextLocURL *string
	PrevLocURL *string
}

const BaseUrl = "https://pokeapi.co/api/v2/location"

var Cfg = &Config{}
