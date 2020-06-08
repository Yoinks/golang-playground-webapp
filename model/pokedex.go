package model

//Pokedex used for json serialization and served by web controller
type Pokedex struct {
	ID    int
	Name  string
	Type  []string
	Stats PokeDexStats
}
