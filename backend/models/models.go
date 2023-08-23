package models

type Item interface {
	Id() Id
	Name() string
}

type Id int

type ItemMeta struct {
	id   Id
	Name string
}

func (i ItemMeta) Id() Id {
	return i.id
}

type DeckItem struct {
	ItemMeta
	Decks map[Id]int
	Cards map[Id]int
}

type CardItem struct {
	ItemMeta
	Contents string
}

type Card struct {
	Item CardItem
	Face bool
}
