package models

type Personalidade struct {
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade
