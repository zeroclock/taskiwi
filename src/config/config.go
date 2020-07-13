package config

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type Config struct {
	IData *InputData
	CData *[]ClockData
}

type InputData struct {
	FilePath    string
	FileContent [][]string
}

type ClockData struct {
	Task     string   `json:"task"`
	Parents  string   `json:"parents"`
	Category string   `json:"category"`
	Start    string   `json:"start"`
	End      string   `json:"end"`
	Effort   string   `json:"effort"`
	Ishabit  string   `json:"ishabit"`
	Tags     []string `json:"tags"`
}

var GlobalConf *Config

func InitConfig(path string) *Config {
	csvFile, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
	}

	idata := InputData{
		FilePath:    path,
		FileContent: csvData,
	}

	var cdata ClockData
	var cdatas []ClockData

	for _, each := range csvData {
		cdata.Task = each[0]
		cdata.Parents = each[1]
		cdata.Category = each[2]
		cdata.Start = each[3]
		cdata.End = each[4]
		cdata.Effort = each[5]
		cdata.Ishabit = each[6]
		cdata.Tags = strings.Split(each[7], ":")

		cdatas = append(cdatas, cdata)
	}

	return &Config{
		IData: &idata,
		CData: &cdatas,
	}
}
