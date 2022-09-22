package model

type EventDataBuilder struct {
	OriginOrg         string `csv:"母団体名"`
	IconDataId        string `csv:"アイコン画像(正方形)"`
	EventOrgName      string `csv:"企画団体名"`
	EventTitle        string `csv:"出展名"`
	EventSummary      string `csv:"企画内容"`
	EventTitleWeb     string `csv:"出展名(Web)"`
	OrgNameWeb        string `csv:"団体名(Web)"`
	EventDescription  string `csv:"企画説明文(字数制限なし)"`
	EventDescriptionP string `csv:"企画説明文(全角30字まで)"`
	EventGenreText    string `csv:"企画のジャンル"`
	OrgDescription    string `csv:"団体説明文(任意)"`
	SnsTwitter        string `csv:"団体のTwitterアカウント(任意)"`
	SnsFacebook       string `csv:"団体のFacebookアカウント(任意)"`
	SnsInstagram      string `csv:"団体のInstagramアカウント(任意)"`
	SnsWebsite        string `csv:"団体のWebページ(任意)"`
	ContactAddress    string `csv:"連絡先メールアドレス"`
}
