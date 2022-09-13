package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

type EventData struct {
	eventTitle       string `field:"☆出展名"`
	eventDescription string `field:"企画説明文(字数制限なし)"`
	eventGenreText   string `field:"企画のジャンル"`
	orgName          string `field:"企画団体名（前述と同じ）\nプレビューサイトをご確認の上、\n改行を行う場合は改行した状態で記入をお願いします。"`
	orgDescription   string `field:"団体説明文(任意)"`
	snsTwitter       string `field:"団体のTwitterアカウント(任意)"`
	snsFacebook      string `field:"団体のFacebookアカウント(任意)"`
	snsInstagram     string `field:"団体のInstagramアカウント(任意)"`
	snsWebsite       string `field:"団体のWebページ(任意)"`
}

func main() {
	f, err := excelize.OpenFile("EventRawData.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet")
	if err != nil {
		fmt.Println(err)
		return
	}

	//for i := 0; i < len(rows[0]); i++ {
	//
	//}
	fmt.Println(len(rows))
	//
	rawData := []*EventData{}
	for i := 1; i < len(rows); i++ {
		var newData EventData
		newData.eventTitle = rows[i][5]
		newData.eventDescription = rows[i][9]
		newData.eventGenreText = rows[i][16]
		newData.orgName = rows[i][7]
		newData.orgDescription = rows[i][10]
		newData.snsTwitter = rows[i][11]
		newData.snsFacebook = rows[i][13]
		newData.snsInstagram = rows[i][12]
		newData.snsWebsite = rows[i][14]
		fmt.Println(newData)
		rawData = append(rawData, &newData)
	}
	for _, data := range rawData {
		fmt.Println(data.orgName)
	}
}
