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
	contactAddress   string
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
	OK      Status = "OK"
	Changed Status = "フォーマットの変更がされました"
	Warning Status = "確認が必要な変更がされました"
	NG      Status = "不正な値です"
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
	newData.contactAddress = builder.ContactAddress
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
	//_, s1 := e.validateEventTitle()
	//_, s1 := e.validateEventDescription()
	//_, s1 := e.validateEventGenreText()
	//_, s1 := e.validateOrgName()
	//_, s1 := e.validateOrgDescription()
	//_, s1 := e.validateSnsTwitter()
	//if s1 == NG {
	//	fmt.Println(e.snsTwitter)
	//}
	//_, s1 := e.validateSnsInstagram()
	//_, s1 := e.validateSnsFacebook()
	_, s2 := e.validateSnsWebsite()
	if s2 == NG {
		fmt.Println(e.snsWebsite)
	}
}

func validAsID(s string) string {
	re := regexp.MustCompile(`^@?([A-Za-z0-9_]+) *$`)
	id := re.FindStringSubmatch(s)
	if id == nil {
		return ""
	}
	return id[0]
}

func (e *EventData) validateEventTitle() (string, Status) {

	return "", OK
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
	if e.snsTwitter == "" {
		return "", OK
	}
	if id := validAsID(e.snsTwitter); id != "" {
		return id, OK
	}
	re := regexp.MustCompile("^https://twitter.com/([a-zA-Z0-9_]+)")
	if id := re.FindStringSubmatch(e.snsTwitter); id != nil {
		return id[0], Changed
	}
	return "", NG
}

func (e *EventData) validateSnsFacebook() (string, Status) {
	if e.snsFacebook == "" {
		return "", OK
	}
	if id := validAsID(e.snsFacebook); id != "" {
		return id, OK
	}
	return "", NG
}

func (e *EventData) validateSnsInstagram() (string, Status) {
	if e.snsInstagram == "" {
		return "", OK
	}
	if id := validAsID(e.snsInstagram); id != "" {
		return id, OK
	}
	return "", NG
}

//func (e *EventData) validateRequestInstagram() {
//	resp, err := http.Get("https://www.instagram.com/usernamafawefwaefaefawefawe")
//	fmt.Println(err)
//	fmt.Println(resp)
//}

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
