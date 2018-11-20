package model

//Cidade representa a cidadee estado do Brasil
type Cidade struct {
	Nome   string `json:"nome"`
	Estado string `json:"estado"`
}
