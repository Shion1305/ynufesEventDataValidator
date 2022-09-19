function myFunction() {
    const c = SpreadsheetApp.getActive().getSheetByName("out-1");
    const sheet = c.getDataRange().getValues();
    for (var i = 1; i < sheet.length; i++) {
        let recipient = sheet[i][1];
        let subject = "[22常盤祭] 企画団体情報のご確認のお願い";
        let body = createBody(sheet[i]);
        GmailApp.createDraft(recipient,subject,body)
    }

}


function createBody(data) {
    console.log(data)
    var url=String(data[2])
    console.log(url)
    var content = data[0] + " " + "常盤祭 企画担当者様へ" + "\n" +
        "\n" +
        "お世話になっております。\n" +
        "大学祭実行委員会編集部、技術部です。先日は企画団体情報フォームのご提出のほどありがとうございました。\n" +
        "ご提出いただきました企画団体情報について、ご確認をお願いしたいです。"
        + "\n【共通項目】"
        + "\n出展名: " + data[3]
        + "\n企画内容: " + data[4]
        + "\n企画のジャンル: " + data[7]
        + "\n【Web項目】"
        + "\n団体名: " + data[8]
        + "\n企画説明文(Web 字数制限無し): " + data[5]
        + "\n団体説明文: " + data[9]
        + "\nTwitterアカウント(IDのみ): " + data[10]
        + "\nFacebookアカウント(IDのみ): " + data[11]
        + "\nInstagramアカウント(IDのみ): " + data[12]
        + "\nWebページ(任意項目): " + data[13]
        + "\n【パンフ項目】"
        + "\n企画説明文(パンフレット 全角30文字まで): " + data[6]
        + "\n"
        + "\n"
        + "また、以下のURLよりWEBページでの表示の確認ができます。必ずご確認くださいますようお願いいたします。\n";
    content=content+url;

    content=content+
        "\n" +
        "\n" +
        "また、ご提出いただいた情報を訂正したい場合は、以下のGoogleフォームより訂正した企画団体情報をご提出ください。提出期限は、【9/23(金)24:00】までとさせていただきます。\n" +
        "https://forms.gle/mxrgWkAmRRmZCdfaA\n" +
        "\n" +
        "今後ともどうぞよろしくお願いいたします。\n" +
        "\n" +
        "22大学祭実行委員会\n" +
        "編集部\n" +
        "技術部"
    console.log(content)
    return content
}
