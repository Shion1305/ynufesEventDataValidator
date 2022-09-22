package model

import (
	"errors"
	"github.com/gnue/go-disp_width"
	"net/http"
	"regexp"
)

type EventData struct {
	originOrg         string
	eventIdMD5        ID
	iconDataId        string
	eventOrgName      string
	eventTitle        string
	eventSummary      string
	eventTitleWeb     string
	eventDescription  string
	eventDescriptionP string
	eventGenre        EventGenre
	orgNameWeb        string
	orgDescription    string
	snsTwitter        verificationField
	snsFacebook       verificationField
	snsInstagram      verificationField
	snsWebsite        verificationField
	contactAddress    string
	originalBuilder   *EventDataBuilder
	ImgStatus         string
}

type verificationField struct {
	Value       string
	OriginValue string
	Verified    VerificationStatus
	Status      Status
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

func (e *verificationField) getSafeValue() string {
	if e.Status != NG && e.Verified != Error {
		return e.Value
	}
	return ""
}

func (e *verificationField) getCheckString() string {
	if e.OriginValue == "" {
		return "(設定なし)"
	}
	if e.Status != NG && e.Verified == Verified {
		return e.Value + " (確認済み)"
	}
	if e.Status != NG && e.Verified == Unverified {
		return e.Value + " (未確認)"
	}
	return "(エラー・設定なし 無効な入力: " + e.OriginValue + ")"

}

type EventField string

const (
	OriginOrg         EventField = "originOrg"
	EventOrgName      EventField = "eventOrgName"
	EventTitle        EventField = "eventTitle"
	EventDescription  EventField = "eventDescription"
	EventDescriptionP EventField = "eventDescriptionP"
	EventGenreF       EventField = "eventGenre"
	OrgNameWeb        EventField = "orgNameWeb"
	OrgDescription    EventField = "orgDescription"
	SnsTwitter        EventField = "snsTwitter"
	SnsFacebook       EventField = "snsFacebook"
	SnsInstagram      EventField = "snsInstagram"
	SnsWebsite        EventField = "snsWebsite"
	ContactAddress    EventField = "contactAddress"
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

func (genre EventGenre) getEventGenreId() int {
	switch genre {
	case Exhibition:
		return 1
	case Performance:
		return 2
	case GameSports:
		return 3
	case Dessert:
		return 4
	case NoodleTeppanyaki:
		return 5
	case FastFood:
		return 6
	case Drink:
		return 7
	case RiceDish:
		return 8
	default:
		return 0
	}
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
	newData.eventIdMD5 = genID(builder.OriginOrg)
	newData.originOrg = builder.OriginOrg
	newData.iconDataId = getIconId(builder.IconDataId)
	newData.eventOrgName = builder.EventOrgName
	newData.eventTitle = builder.EventTitle
	newData.eventSummary = builder.EventSummary
	newData.eventTitleWeb = builder.EventTitleWeb
	newData.orgNameWeb = builder.OrgNameWeb
	newData.eventDescription = builder.EventDescription
	newData.eventDescriptionP = builder.EventDescriptionP
	newData.eventGenre = EventGenre(builder.EventGenreText)
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
	resp.OriginValue = value
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
	case EventOrgName:
		e.eventOrgName = value
		break
	case OriginOrg:
		e.originOrg = value
		break
	case OrgDescription:
		e.orgDescription = value
		break
	case EventDescriptionP:
		e.eventDescriptionP = value
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
	e.validateSnsInstagram()
	e.validateSnsFacebook()
	e.validateSnsWebsite()
}

func validAsID(s string) string {
	re := regexp.MustCompile(`^@?((\w){1,15}) *$`)
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

func (e *EventData) ValidateDescriptionP() bool {
	//limit EventDescriptionP visual text length to 60(with half size character)
	return disp_width.Measure(e.eventDescriptionP) <= 60
}

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
	re := regexp.MustCompile("^@?([A-Z.a-z0-9]{5,})$")
	check := re.FindStringSubmatch(e.snsFacebook.Value)
	if check != nil && check[1] != "" {
		e.snsFacebook.setStatus(check[1], OK)
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
	re := regexp.MustCompile("^@?([A-Z.a-z0-9_]+)$")
	check := re.FindStringSubmatch(e.snsInstagram.Value)
	if check != nil && check[1] != "" {
		e.snsInstagram.setStatus(check[1], OK)
		return
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

type CheckEventData struct {
	OriginOrg         string `csv:"OriginOrg"`
	ContactAddress    string `csv:"ContactAddress"`
	Url               string `csv:"Url"`
	EventOrgName      string `csv:"eventOrgName"`
	EventTitle        string `csv:"eventTitle"`
	EventSummary      string `csv:"eventSummary"`
	EventTitleWeb     string `csv:"eventTitleWeb"`
	EventDescription  string `csv:"eventDescription"`
	EventDescriptionP string `csv:"eventDescriptionP"`
	EventGenreText    string `csv:"eventGenreText"`
	OrgNameWeb        string `csv:"orgNameWeb"`
	OrgDescription    string `csv:"orgDescription"`
	SnsTwitter        string `csv:"snsTwitter"`
	SnsFacebook       string `csv:"snsFacebook"`
	SnsInstagram      string `csv:"snsInstagram"`
	SnsWebsite        string `csv:"snsWebsite"`
	ImageComment      string `csv:"imageComment"`
}

func (e *EventData) ExportCheck() *CheckEventData {
	return &CheckEventData{
		OriginOrg:         e.originOrg,
		ContactAddress:    e.contactAddress,
		Url:               "https://tokiwa22.ynu-fes.yokohama/preview/event-detail/" + string(e.eventIdMD5),
		EventOrgName:      e.eventOrgName,
		EventTitle:        e.eventTitle,
		EventSummary:      e.eventSummary,
		EventTitleWeb:     e.eventTitleWeb,
		EventDescription:  e.eventDescription,
		EventDescriptionP: e.eventDescriptionP,
		EventGenreText:    string(e.eventGenre),
		OrgNameWeb:        e.orgNameWeb,
		OrgDescription:    e.orgDescription,
		SnsTwitter:        e.snsTwitter.getCheckString(),
		SnsFacebook:       e.snsFacebook.getCheckString(),
		SnsInstagram:      e.snsInstagram.getCheckString(),
		SnsWebsite:        e.snsWebsite.getCheckString(),
		ImageComment:      e.ImgStatus,
	}
}
