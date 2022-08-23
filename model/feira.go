package model

import "gorm.io/gorm"

type Feira struct {
	gorm.Model
	Id         int16
	Long       string
	Lat        string
	SetCens    string
	AreaP      string
	CodDist    string
	Distrito   string
	CodSubPref string
	SubPrere   string
	Regiao5    string
	Regiao8    string
	NomeFreira string
	Registo    string
	Logradouro string
	Numero     string
	Bairro     string
	Referencia string
}
