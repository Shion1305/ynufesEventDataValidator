package model

type EventDataBuilder struct {
	EventTitle       string `field:"☆出展名"`
	EventDescription string `field:"企画説明文(字数制限なし)"`
	EventGenreText   string `field:"企画のジャンル"`
	OrgName          string `field:"企画団体名（前述と同じ）\nプレビューサイトをご確認の上、\n改行を行う場合は改行した状態で記入をお願いします。"`
	OrgDescription   string `field:"団体説明文(任意)"`
	SnsTwitter       string `field:"団体のTwitterアカウント(任意)"`
	SnsFacebook      string `field:"団体のFacebookアカウント(任意)"`
	SnsInstagram     string `field:"団体のInstagramアカウント(任意)"`
	SnsWebsite       string `field:"団体のWebページ(任意)"`
}
