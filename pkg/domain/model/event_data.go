package model

import (
	"errors"
	"net/http"
	"regexp"
)

type EventData struct {
	originOrg        string
	eventTitle       string
	eventDescription string
	eventGenre       EventGenre
	orgName          string
	orgDescription   string
	snsTwitter       verificationField
	snsFacebook      verificationField
	snsInstagram     verificationField
	snsWebsite       verificationField
	contactAddress   string
	originalBuilder  *EventDataBuilder
}

type verificationField struct {
	Value    string
	Verified VerificationStatus
	Status   Status
}

type VerificationStatus string

const (
	Verified   VerificationStatus = "Verified"
	Unverified VerificationStatus = "Unverified"
	Error      VerificationStatus = "Error"
)

func (e *verificationField) setVerification(vStatus VerificationStatus) {
	e.Verified = vStatus
}

func (e *verificationField) setStatus(value string, status Status) {
	e.Status = status
	e.Value = value
}

type EventField string

const (
	EventTitle       EventField = "eventTitle"
	EventDescription EventField = "eventDescription"
	EventGenreF      EventField = "eventGenre"
	OrgName          EventField = "orgName"
	OrgDescription   EventField = "orgDescription"
	SnsTwitter       EventField = "snsTwitter"
	SnsFacebook      EventField = "snsFacebook"
	SnsInstagram     EventField = "snsInstagram"
	SnsWebsite       EventField = "snsWebsite"
	ContactAddress   EventField = "contactAddress"
)

type EventGenre string

const (
	Exhibition       EventGenre = "展示・体験・販売"
	Performance      EventGenre = "パフォーマンス"
	GameSports       EventGenre = "ゲーム・スポーツ"
	Dessert          EventGenre = "デザート"
	NoodleTeppanyaki EventGenre = "鉄板・麺類"
	FastFood         EventGenre = "ファストフード"
	Drink            EventGenre = "ドリンク"
	RiceDish         EventGenre = "ご飯もの"
)

type Status string

const (
	OK      Status = "OK"
	Changed Status = "フォーマットの変更がされました"
	Warning Status = "確認が必要な変更がされました"
	NG      Status = "不正な値です"
)

func NewEventData(builder EventDataBuilder) *EventData {
	var newData EventData
	newData.originOrg = builder.OriginOrg
	newData.eventTitle = builder.EventTitle
	newData.eventDescription = builder.EventDescription
	newData.eventGenre = EventGenre(builder.EventGenreText)
	newData.orgName = builder.OrgName
	newData.orgDescription = builder.OrgDescription
	newData.snsTwitter = initVerificationField(builder.SnsTwitter)
	newData.snsFacebook = initVerificationField(builder.SnsFacebook)
	newData.snsInstagram = initVerificationField(builder.SnsInstagram)
	newData.snsWebsite = initVerificationField(builder.SnsWebsite)
	newData.contactAddress = builder.ContactAddress
	newData.originalBuilder = &builder
	newData.validate()
	return &newData
}

func initVerificationField(value string) verificationField {
	var resp verificationField
	resp.Value = value
	resp.Verified = Unverified
	return resp
}

func NewMultiEventData(builders []EventDataBuilder) []*EventData {
	var data []*EventData
	for _, builder := range builders {
		data = append(data, NewEventData(builder))
	}
	return data
}

func (e *EventData) UpdateField(field EventField, value string) error {
	switch field {
	case EventTitle:
		e.eventTitle = value
		break
	case EventDescription:
		e.eventDescription = value
		break
	case EventGenreF:
		e.eventGenre = EventGenre(value)
		break
	case OrgName:
		e.orgName = value
		break
	case OrgDescription:
		e.orgDescription = value
		break
	case SnsTwitter:
		e.snsTwitter.Value = value
		break
	case SnsFacebook:
		e.snsFacebook.Value = value
		break
	case SnsInstagram:
		e.snsInstagram.Value = value
		break
	case SnsWebsite:
		e.snsWebsite.Value = value
		break
	case ContactAddress:
		e.contactAddress = value
		break
	default:
		return errors.New("unknown Field")
	}
	e.validate()
	return nil
}

func (e *EventData) validate() {
	//_, s1 := e.validateEventTitle()
	//_, s1 := e.validateEventDescription()
	//_, s1 := e.validateEventGenreText()
	//_, s1 := e.validateOrgName()
	//_, s1 := e.validateOrgDescription()
	e.ValidateSnsTwitter()
	//_, s1 := e.validateSnsInstagram()
	//_, s1 := e.validateSnsFacebook()
	//_, s2 := e.validateSnsWebsite()
	//if s2 == NG {
	//	fmt.Println(e.snsWebsite)
	//}
}

func validAsID(s string) string {
	re := regexp.MustCompile(`^@?([A-Za-z0-9]*_?[A-Za-z0-9]*) *$`)
	id := re.FindStringSubmatch(s)
	if id == nil {
		return ""
	}
	return id[1]
}

//func (e *EventData) validateEventTitle() (string, Status) {
//
//	return "", OK
//}
//
//func (e *EventData) validateEventDescription() (string, Status) {
//	return "", OK
//}
//
//func (e *EventData) validateEventGenreText() (string, Status) {
//	return "", OK
//}
//
//func (e *EventData) validateOrgName() (string, Status) {
//	return "", OK
//}
//
//func (e *EventData) validateOrgDescription() {
//	return "", OK
//}

func (e *EventData) ValidateSnsTwitter() {
	if e.snsTwitter.Value == "" {
		e.snsTwitter.setStatus("", OK)
		return
	}
	if id := validAsID(e.snsTwitter.Value); id != "" {
		e.snsTwitter.setStatus(id, OK)
		return
	}
	re := regexp.MustCompile("^https://twitter.com/([A-Za-z0-9]*_?[A-Za-z0-9]*)")
	if id := re.FindStringSubmatch(e.snsTwitter.Value); id != nil {
		if name := validAsID(id[1]); name != "" {
			e.snsTwitter.setStatus(name, Changed)
			return
		}
	}
	e.snsTwitter.setStatus(e.snsTwitter.Value, NG)
	return
}

func (e *EventData) validateSnsFacebook() {
	if e.snsFacebook.Value == "" {
		e.snsFacebook.setStatus("", OK)
	}
	if id := validAsID(e.snsFacebook.Value); id != "" {
		e.snsFacebook.setStatus(id, OK)
		return
	}
	e.snsFacebook.setStatus("", NG)
	return
}

func (e *EventData) validateSnsInstagram() {
	if e.snsInstagram.Value == "" {
		e.snsInstagram.setStatus("", OK)
		return
	}
	if id := validAsID(e.snsInstagram.Value); id != "" {
		e.snsInstagram.setStatus(id, OK)
	}
	e.snsInstagram.setStatus("", NG)
	return
}

//func (e *EventData) validateRequestInstagram() {
//	resp, err := http.Get("https://www.instagram.com/usernamafawefwaefaefawefawe")
//	fmt.Println(err)
//	fmt.Println(resp)
//}

func (e *EventData) validateSnsWebsite() {
	if e.snsWebsite.Value == "" {
		e.snsWebsite.setStatus("", OK)
		return
	}
	re := regexp.MustCompile("^(https?://[/.a-z0-9_-]+)$")
	if match := re.FindStringSubmatch(e.snsWebsite.Value); match != nil {
		if accessTest(match[0]) {
			e.snsWebsite.setStatus(match[0], OK)
			e.snsWebsite.setVerification(Verified)
		} else {
			e.snsWebsite.setStatus(match[0], OK)
			e.snsWebsite.setVerification(Error)
		}
		return
	}
	e.snsWebsite.setStatus(e.snsWebsite.Value, NG)
	return
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
