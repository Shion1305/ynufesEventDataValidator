package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Patch struct {
	EventTitle string     `json:"eventTitle"`
	FieldName  EventField `json:"patchField"`
	FieldValue string     `json:"patchValue"`
}

func ReadPatches(file string) []*Patch {
	source, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
		return nil
	}
	var patches []*Patch
	err = json.Unmarshal(source, &patches)
	if err != nil {
		panic(err)
		return nil
	}
	return patches
}

func (p *Patch) ApplyPatch(data []*EventData) error {
	var hits []*EventData
	for _, d := range data {
		if d.eventTitle == d.eventTitle {
			hits = append(hits, d)
		}
	}
	if len(hits) > 1 {
		return errors.New("候補が複数あります。")
	}
	if len(hits) == 0 {
		fmt.Printf("Not Applied: %s, %s -> %s", p.EventTitle, p.FieldName, p.FieldValue)
		return nil
	}
	err := hits[0].UpdateField(p.FieldName, p.FieldValue)
	if err != nil {
		return err
	}
	return nil
}
