package model

// ExportEventData WEB用のExporter 必要のないデータは含まない
type ExportEventData struct {
	EventIdMD5       string `json:"event_id"`
	EventTitle       string `json:"event_title"`
	EventSummary     string `json:"event_summary"`
	EventDescription string `json:"event_description"`
	EventGenreId     int    `json:"event_genre_id"`
	OrgNameWeb       string `json:"org_name"`
	OrgDescription   string `json:"org_description"`
	SnsTwitter       string `json:"sns_twitter"`
	SnsFacebook      string `json:"sns_facebook"`
	SnsInstagram     string `json:"sns_instagram"`
	SnsWebsite       string `json:"sns_website"`
}

func (e *EventData) Export() ExportEventData {
	return ExportEventData{
		EventIdMD5:       string(e.eventIdMD5),
		EventTitle:       e.eventTitle,
		EventDescription: e.eventDescription,
		EventGenreId:     e.eventGenre.getEventGenreId(),
		OrgNameWeb:       e.orgNameWeb,
		OrgDescription:   e.orgDescription,
		SnsTwitter:       e.snsTwitter.getSafeValue(),
		SnsFacebook:      e.snsFacebook.getSafeValue(),
		SnsInstagram:     e.snsInstagram.getSafeValue(),
		SnsWebsite:       e.snsWebsite.getSafeValue(),
	}
}
