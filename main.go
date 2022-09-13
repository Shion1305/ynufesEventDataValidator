package main

import (
	"github.com/gocarina/gocsv"
	"os"
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
	for _, d := range eventDataSet {
		d.Validate()
	}
}
