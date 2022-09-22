package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"ynufesEventDataValidator/pkg/domain/model"
)

func main() {
	var builders []model.EventDataBuilder
	in, err := os.Open("EventRawData.csv")
	if err != nil {
		panic(err)
	}
	if err := gocsv.UnmarshalFile(in, &builders); err != nil {
		panic(err)
	}
	eventDataSet := model.NewMultiEventData(builders)
	//checkPatches(eventDataSet)
	model.ValidateTwitter(eventDataSet)
	processIcons(eventDataSet)
	exportCSV(eventDataSet)
	exportJson(eventDataSet)
}

func processIcons(data []*model.EventData) {
	drive := model.InitGD()
	for _, e := range data {
		//model.PrintError(e)
		model.ProcessGD(drive, *e)
	}
}

func exportJson(data []*model.EventData) error {
	out, err := os.Create("source.json")
	if err != nil {
		return fmt.Errorf("error opening source: %w\n", err)
	}
	var exports []model.ExportEventData
	for _, d := range data {
		exports = append(exports, d.Export())
	}
	byteData, _ := json.Marshal(exports)
	_, err = out.Write(byteData)
	return err
}

func checkPatches(data []*model.EventData) {
	re := regexp.MustCompile("^patch-\\d{2}\\.json$")
	filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println("パッチを読み込めませんでした。")
			return err
		}
		if !info.IsDir() && re.MatchString(info.Name()) {
			fmt.Printf("Patch detected: %s\n", path)
			for _, patch := range model.ReadPatches(path) {
				err := patch.ApplyPatch(data)
				if err != nil {
					fmt.Println("パッチの適用に失敗しました")
				}
			}
		}
		return nil
	})
}

func exportCSV(data []*model.EventData) {
	var checks []*model.CheckEventData
	for _, d := range data {
		checks = append(checks, d.ExportCheck())
	}

	file, err := os.Create("out.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	gocsv.MarshalFile(&checks, file)
}
