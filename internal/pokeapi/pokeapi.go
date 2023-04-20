package pokeapi

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Location struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}
