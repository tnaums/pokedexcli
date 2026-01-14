package dex

type Stats struct {
	Hp int
	Attack int
	Defense int
	SpecialA int
	SpecialD int
	Speed int
}

type  Pokemon struct {
	Name string
	Height int
	Weight int
	Stats Stats
	Embodiments []string
}

func NewDex() map[string]Pokemon {
	return make(map[string]Pokemon)
}
