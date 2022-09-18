package model

import "fmt"

func PrintError(event *EventData) {
	hasError := false
	if event.snsTwitter.Status == NG {
		printHeader(&hasError)
		printOriginValue(event.originalBuilder.SnsTwitter)
		fmt.Println("TwitterIDではありません。")
	} else if event.snsTwitter.Verified == Error {
		printHeader(&hasError)
		printOriginValue(event.originalBuilder.SnsTwitter)
		fmt.Println("TwitterIDが存在しませんでした。")
	}
	if event.snsInstagram.Status == NG {
		printHeader(&hasError)
		printOriginValue(event.originalBuilder.SnsInstagram)
		fmt.Println("InstagramIDではありません。")
	} else if event.snsInstagram.Verified == Error {
		printHeader(&hasError)
		printOriginValue(event.originalBuilder.SnsInstagram)
		fmt.Println("Instagramアカウントが存在しませんでした。")
	}
	if event.snsFacebook.Status == NG {
		printHeader(&hasError)
		printOriginValue(event.originalBuilder.SnsFacebook)
		fmt.Println("FacebookIDではありません。")
	} else if event.snsFacebook.Verified == Error {
		printHeader(&hasError)
		printOriginValue(event.originalBuilder.SnsFacebook)
		fmt.Println("Facebookアカウントが存在しませんでした。")
	}
	if event.snsWebsite.Status == NG {
		printHeader(&hasError)
		printOriginValue(event.originalBuilder.SnsWebsite)
		fmt.Println("URLのフォーマットではありません。")
	} else if event.snsWebsite.Verified == Error {
		printHeader(&hasError)
		printOriginValue(event.originalBuilder.SnsWebsite)
		fmt.Println("ウェブサイトに正しくアクセスできませんでした。")
	}
}

func printHeader(hasError *bool) {
	if !*hasError {
		fmt.Printf("%30sの%30sについて以下のエラーがあります。\n")
		*hasError = true
	}
}

func printOriginValue(value string) {
	fmt.Printf("入力値: %s\n", value)
}
