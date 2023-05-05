package models

import (
	"time"

	"gorm.io/gorm"
)

type Recette struct {
	gorm.Model
	Instructions []Instruction
	Name         string `json:"name"`
	Link         string `json:"link"`
	Image        string `json:"image"`
	Ingredients  []Ingredient
	CreatedAt    time.Time
}
