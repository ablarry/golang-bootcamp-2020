package repo

import "github.com/ablarry/golang-bootcamp-2020/pkg/model"

// PMock is a mock to test
var PMock = &model.Pokemon{
	ID:       1,
	Name:     "Bulbasaur",
	Type:     "Grass",
	Category: "Seed",
	Weakness: "Fire",
}

// PMock2 is a mock to test
var PMock2 = &model.Pokemon{
	ID:       2,
	Name:     "Ivysaur",
	Type:     "Grass",
	Category: "Category",
	Weakness: "Fire",
}

// PMocks list of mocks to test
var PMocks = []*model.Pokemon{PMock, PMock2}
