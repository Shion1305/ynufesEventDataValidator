package model

type EventDataBuilder struct {
	OriginOrg        string `csv:"☆母団体名"`
	EventTitle       string `csv:"☆出展名"`
	EventDescription string `csv:"企画説明文(字数制限なし)"`
	EventGenreText   string `csv:"企画のジャンル"`
	OrgName          string `csv:"企画団体名（前述と同じ）\nプレビューサイトをご確認の上、\n改行を行う場合は改行した状態で記入をお願いします。"`
	OrgDescription   string `csv:"団体説明文(任意)"`
	SnsTwitter       string `csv:"団体のTwitterアカウント(任意)"`
	SnsFacebook      string `csv:"団体のFacebookアカウント(任意)"`
	SnsInstagram     string `csv:"団体のInstagramアカウント(任意)"`
	SnsWebsite       string `csv:"団体のWebページ(任意)"`
	ContactAddress   string `csv:"☆連絡用メールアドレス"`
}
