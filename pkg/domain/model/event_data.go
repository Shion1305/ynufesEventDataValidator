package model

import (
	"fmt"
	"net/http"
	"regexp"
)

type EventData struct {
	eventTitle       string
	eventDescription string
	eventGenreText   string
	orgName          string
	orgDescription   string
	snsTwitter       string
	snsFacebook      string
	snsInstagram     string
	snsWebsite       string
}

type ValidationErrors struct {
	err ValidationError
}
type ValidationError struct {
	Field   string
	Message string
}

type Status string

const (
	OK       Status = "OK"
	Changed  Status = "フォーマットの変更がされました"
	Selected Status = "複数の候補から選択しました"
	NG       Status = "不正な値です"
)

func NewEventData(builder EventDataBuilder) *EventData {
	var newData EventData
	newData.eventTitle = builder.EventTitle
	newData.eventDescription = builder.EventDescription
	newData.eventGenreText = builder.EventGenreText
	newData.orgName = builder.OrgName
	newData.orgDescription = builder.OrgDescription
	newData.snsTwitter = builder.SnsTwitter
	newData.snsFacebook = builder.SnsFacebook
	newData.snsInstagram = builder.SnsInstagram
	newData.snsWebsite = builder.SnsWebsite
	return &newData
}

func NewMultiEventData(builders []EventDataBuilder) []*EventData {
	var data []*EventData
	for _, builder := range builders {
		data = append(data, NewEventData(builder))
	}
	return data
}

func (e *EventData) Validate() {

}

func isPlainString(s string) bool {
	res, _ := regexp.MatchString("^[A-Za-z0-9]+$", s)
	return res
}

func (e *EventData) validateEventTitle() ValidationError {
	return ValidationError{}
}

func (e *EventData) validateEventDescription() (string, Status) {
	return "", OK
}

func (e *EventData) validateEventGenreText() (string, Status) {
	return "", OK
}

func (e *EventData) validateOrgName() (string, Status) {
	return "", OK
}

func (e *EventData) validateOrgDescription() (string, Status) {
	return "", OK
}

func (e *EventData) validateSnsTwitter() (string, Status) {
	if isPlainString(e.snsTwitter) {
		return "", OK
	}
	return "", Changed
}

func (e *EventData) validateSnsFacebook() (string, Status) {
	if isPlainString(e.snsFacebook) {
		return "", OK
	}
	return "", OK
}

func (e *EventData) validateSnsInstagram() (string, Status) {
	if isPlainString(e.snsInstagram) {
		return "", OK
	}
	return "", OK
}

func (e *EventData) validateSnsWebsite() (string, Status) {
	if e.snsWebsite == "" {
		return "", OK
	}
	re := regexp.MustCompile("^(https?://[/.a-z0-9_-]+)$")
	if match := re.FindStringSubmatch(e.snsWebsite); match != nil {
		if accessTest(match[0]) {
			return match[0], OK
		} else {
			return "", NG
		}
	}
	return "", NG
}

func accessTest(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		print(err)
		return false
	}
	if resp.StatusCode > 300 {
		return false
	}
	return true
}
