package main

import (
	"encoding/csv"
	"strings"
)

func csvToList(csvString string) [][]string {

	// CSVリーダーの作成
	reader := csv.NewReader(strings.NewReader(csvString))

	// デリミタの設定
	reader.Comma = ','

	// レコードを格納するスライスの初期化
	var records [][]string

	// CSV文字列を1行ずつ読み込む
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		// レコードをスライスとしてrecordsに追加する
		records = append(records, record)
	}

	// recordsを返却
	return records
}
