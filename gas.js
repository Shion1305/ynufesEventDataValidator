function myFunction() {
    const c = SpreadsheetApp.getActive().getSheetByName("out");
    const sheet = c.getDataRange().getValues();
    for (var i = 1; i < sheet.length; i++) {
        let recipient = sheet[i][1];
        let subject = "[22常盤祭] 企画団体情報のご確認のお願い";
        let body = createBody(sheet[i]);
        GmailApp.createDraft(recipient, subject, body)
    }
}


function createBody(data) {
    var url = String(data[2])
    console.log(url)
    var content = "常盤祭に参加される企画団体の皆様\n" +
        "\n" +
        "お世話になっております。\nフォームにご回答いただきありがとうございました。\n\n"
        + "2次受付の企画情報と一部団体の企画情報の訂正を反映しました。"
        + "\nアイコン画像やSNS欄も含めて登録内容に問題ないか再度ご確認をお願いします。"
        + "\nまた一次受付団体につきまして、以前確認のメールを送付させて頂いておりましたが"
        + "\nご案内しましたプレビュー用リンクが変更となりました。"
        + "\n下記のURLよりご確認いただくようよろしくお願いします。\n";
    content = content + url;
    content = content
        + "\n\n【共通項目】"
        + "\n母団体名: " + data[0]
        + "\n企画団体名: " + data[3]
        + "\n出展名: " + data[4]
        + "\n企画内容: " + data[5]
        + "\n企画のジャンル: " + data[9]
        + "\nアイコン画像: " + data[16]
        + "\n【Web項目】"
        + "\n団体名: " + data[10]
        + "\n企画説明文(Web 字数制限無し): " + data[7]
        + "\n団体説明文: " + data[11]
        + "\nTwitterアカウント(IDのみ): " + data[12]
        + "\nFacebookアカウント(IDのみ): " + data[13]
        + "\nInstagramアカウント(IDのみ): " + data[14]
        + "\nWebページ(任意項目): " + data[15]
        + "\n【パンフ項目】"
        + "\n企画説明文(パンフレット 全角30文字まで): " + data[8]
        + "\n\n"
        + "また、ご提出いただいた情報を訂正したい場合は、以下のGoogleフォームより訂正した企画団体情報をご提出ください。提出期限は、【9/24(金)18:00】までとさせていただきます。\n"
        + "(訂正する項目が複数存在する場合は複数ご回答ください。)\n"
        + "\nhttps://docs.google.com/forms/d/e/1FAIpQLSfh-qVNMF3jPBVUmMC8w3ukkLQy49QAOLnW3NuAurqoOszsXQ/viewform?usp=sf_link"
        + "\n今後ともどうぞよろしくお願いいたします。\n"
        + "\n" +
        "22大学祭実行委員会\n" +
        "編集部\n技術部"
    return content
}
