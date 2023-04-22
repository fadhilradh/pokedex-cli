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

type LocDetail struct {
	EncounterMethodRates []struct {
		EncounterMethod Result `json:"encounter_method"`
		VersionDetails  []struct {
			Rate    int    `json:"rate"`
			Version Result `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int    `json:"game_index"`
	ID        int    `json:"id"`
	Location  Result `json:"location"`
	Name      string `json:"name"`
	Names     []struct {
		Language Result `json:"language"`
		Name     string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        Result `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int    `json:"chance"`
				ConditionValues []any  `json:"condition_values"`
				MaxLevel        int    `json:"max_level"`
				Method          Result `json:"method"`
				MinLevel        int    `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int    `json:"max_chance"`
			Version   Result `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
