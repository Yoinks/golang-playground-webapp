package model

//PokeDexStats used for json serialization and served by web controller
type PokeDexStats struct {
	HP        int
	Speed     int
	Attack    int
	Defense   int
	SPAttack  int
	SPDEFENSE int
}
