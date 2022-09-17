package main

import (
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
	checkPatches(eventDataSet)
	model.ValidateTwitter(eventDataSet)
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
