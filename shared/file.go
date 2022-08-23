package shared

import (
	"log"
	"os"
)

func openFile(file string) (*os.File, error ){
	
	csvFile, err := os.Open(file)
	if err != nil{
		return nil, err
	}
	return csvFile, nil

}

func writeFile(path, texto string)  {

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	_, err = file.WriteString(texto)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	
}