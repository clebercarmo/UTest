package shared

import "encoding/csv"

func CargaDados()  {

	csvFile, err := openFile("data/DEINFO_AB_FEIRASLIVRES_2014.csv")
	if err != nil {
        SaveLog("error", err.Error())
	}
	SaveLog("info", err.Error())
	defer csvFile.Close()
    
    csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        SaveLog("error", err.Error())
    }    
    for _, line := range csvLines {
        emp := empData{
            Name: line[0],
            Age: line[1],
			City: line[2],
        }
        
    }
	
}