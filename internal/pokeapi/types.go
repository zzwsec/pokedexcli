package pokeapi

type LocationAreasResponse struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []LocationAreaItem `json:"results"`
}

type LocationAreaItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
