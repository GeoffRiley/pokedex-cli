package pokeapi

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func LocationStartURL() *string {
	locBase := baseURL + "/location-area"
	return &locBase
}
