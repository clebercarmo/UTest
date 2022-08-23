package shared

func SaveLog(tipo, texto string)  {
	writeFile("data/"+tipo, texto)
}