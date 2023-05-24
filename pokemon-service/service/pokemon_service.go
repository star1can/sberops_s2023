package service

const (
	URL string = "http://pokeapi.mesh-external.svc.cluster.local/api/v2/pokemon/"
)

type PokemonService interface {
	GetPokemons() ([]Pokemon, error)
}

func GetPokemonService() PokemonService {
	return &pokemonServiceImpl{}
}

type Pokemon struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	BaseExperience int64  `json:"base_experience"`
	Height         int64  `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int64  `json:"order"`
	Weight         int64  `json:"weight"`
}
