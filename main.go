package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Post struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	// 引数からファイルパスを取得
	filePath := flag.String("file", "", "Path to the JSON file")
	flag.Parse()

	if *filePath == "" {
		log.Fatal("Please specify the path to the JSON file using the -file flag")
	}

	// JSONデータをファイルから読み込む
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// JSONデータを読み込む
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// JSONをスライスにデコード
	var posts []Post
	err = json.Unmarshal(bytes, &posts)
	if err != nil {
		log.Fatal(err)
	}

	totalMinutes := 0
	windowStart := posts[0].CreatedAt
	inWindow := false

	// 投稿を順に処理してSNSの利用時間を計算する
	for i := 1; i < len(posts); i++ {
		if posts[i].CreatedAt.Sub(windowStart) <= 10*time.Minute {
			if !inWindow {
				inWindow = true
				totalMinutes += 10
			}
		} else {
			windowStart = posts[i].CreatedAt
			inWindow = false
		}
	}

	fmt.Printf("Total time spent: %d minutes\n", totalMinutes)
}
