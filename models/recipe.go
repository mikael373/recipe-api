package models

type Recipe struct {
	Titulo       string   `json:"titulo"`
	Ingredientes []string `json:"ingredientes"`
	ReceitaURL   string   `json:"receita-url"`
	ReceitaGif   string   `json:"receita-gif"`
}
