package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	// PostOpenAiを呼び出します。
	result := PostOpenAi("白菜に関する情報やアドバイスを提供するアフィリエイトブログ記事を書いてください。マークダウン形式で書いてください。関連する写真の画像も適宜挟んでください。読者に価値ある情報を提供し、関連するアフィリエイト商品を自然に紹介してください。")

	fmt.Println(result)
	// department := "野菜"
	// titles := makeTitles(department)

	// // titlesのを１つずつ処理します。
	// for _, title := range titles {
	// 	// blogPromptを宣言
	// 	var blogPrompt string

	// 	// ブログを作成するためのプロンプトを作成します。
	// 	blogPrompt = makeBlogPrompt(title)

	// 	// apiを呼び出してマークダウンのブログコンテンツを作成します。
	// 	blogContent := PostOpenAi(blogPrompt)

	// 	// 物理ファイルに書き出します。
	// 	writeFile(blogContent)
	// }

	//

	// gojuon := []string{"あ", "い", "う", "え", "お",
	// 	"か", "き", "く", "け", "こ",
	// 	"さ", "し", "す", "せ", "そ",
	// 	"た", "ち", "つ", "て", "と",
	// 	"な", "に", "ぬ", "ね", "の",
	// 	"は", "ひ", "ふ", "へ", "ほ",
	// 	"ま", "み", "む", "め", "も",
	// 	"や", "ゆ", "よ",
	// 	"ら", "り", "る", "れ", "ろ",
	// 	"わ", "を", "ん"}

	// for _, kana := range gojuon {
	// 	fmt.Println(kana)
	// 	time.Sleep(1 * time.Second)
	// }
}

// func makeBlogPrompt(title string) string {
// 	// blogPromptを宣言します。
// 	var blogPrompt string
// 	blogPrompt = "に関する情報やアドバイスを提供するアフィリエイトブログ記事を書いてください。マークダウン形式で書いてください。読者に価値ある情報を提供し、関連するアフィリエイト商品を自然に紹介してください。"
// 	// titleとblogPromptを結合します。
// 	blogPrompt = title + blogPrompt
// 	// blogPromptを返します。
// 	return blogPrompt
// }

// .envを呼び出します。
func loadEnv() {
	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

}

// func makeTitles(department string) []string {
// 	// promptを宣言します。
// 	var prompt string
// 	// departmentに属するリストをcsv文字列で取得します。
// 	prompt = "のリストをcsv形式で出力してください。"
// 	titles := PostOpenAi(prompt)
// 	// csv文字列をリスト形式に変換します。
// 	list := csvToList(titles)
// 	// listを返します。
// 	return list
// }

// 物理ファイルに書き出します。
func WriteFile(blogContent string, timestamp string) {
	// ファイル名を作成します。
	filename := timestamp + ".md"

	// ファイルを作成します。
	file, err := os.Create(filename)

	// もしエラーがあれば、エラーを出力します。
	if err != nil {
		fmt.Println(err)
	}

	// ファイルにblogContentを書き込みます。
	file.WriteString(blogContent)

	// ファイルを閉じます。
	defer file.Close()

}

func GetTimeStamp() string {
	// msまでのunixタイムスタンプを取得します。
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	// timestampをstringに変換します。
	timestampString := strconv.FormatInt(timestamp, 10)
	// timestampStringを返します。
	return timestampString
}
