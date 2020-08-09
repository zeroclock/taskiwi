package config

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/zeroclock/taskiwi/model"
)

type Config struct {
	IData *InputData
	CData *[]model.ClockData
}

type InputData struct {
	FilePath    string
	FileContent [][]string
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

	// skip header line
	if _, err := reader.Read(); err != nil {
		log.Println(err)
	}

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
	}

	idata := InputData{
		FilePath:    path,
		FileContent: csvData,
	}

	var cdata model.ClockData
	var cdatas []model.ClockData

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
