package pokeapi

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type MapResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Maps struct {
	Count    int         `json:"count"`
	Next     *string     `json:"next"`
	Previous *string     `json:"previous"`
	Results  []MapResult `json:"results"`
}
