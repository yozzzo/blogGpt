// testを書きます。
package main

import (
	"os"
	"testing"
)

// PostOpenAiをテストします。
func TestPostOpenAi(t *testing.T) {
	// テストケースを作成します。
	testCases := []struct {
		name   string
		prompt string
	}{
		{
			name:   "test",
			prompt: "こんにちは",
		},
	}
	// テストケースを１つずつ処理します。
	for _, tc := range testCases {
		// テストケースの名前を出力します。
		t.Run(tc.name, func(t *testing.T) {
			// テスト対象の関数を呼び出します。
			result := PostOpenAi(tc.prompt)
			// resultが文字列であることを確認します。
			if result == "" {
				t.Errorf("result is empty string")
			}
		})
	}
}

// WriteFileをテストします
func TestWriteFile(t *testing.T) {
	// テストケースを作成します。
	testCases := []struct {
		name      string
		content   string
		timestamp string
	}{
		{
			name:      "test1",
			content:   "中身はこうです。",
			timestamp: GetTimeStamp(),
		},
		{
			name:      "test2",
			content:   "中身はこうです。",
			timestamp: GetTimeStamp(),
		},
	}
	// テストケースを１つずつ処理します。
	for _, tc := range testCases {
		// テストケースの名前を出力します。
		t.Run(tc.name, func(t *testing.T) {
			// テスト対象の関数を呼び出します。
			WriteFile(tc.content, tc.timestamp)
			// １秒まつ
			// time.Sleep(1 * time.Second)
			// ファイルが存在するか確認します。
			if _, err := os.Stat(tc.timestamp + ".md"); os.IsNotExist(err) {
				// ファイルが存在しない場合はエラーを出力します。
				t.Errorf("ファイルが存在しません。")
			}
		})
	}
}
