package main

import (
	"log"
	"os"

	"fmt"

	"bufio"

	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// 読み込みバッファのサイズ
const (
	BUFSIZE = 1024
)

func main() {
	// ファイルを読み込んでトークン取得
	file, err := os.Open(`./token.txt`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := bufio.NewScanner(file)
	token := ""
	for buf.Scan() {
		token = buf.Text()
	}
	if err := buf.Err(); err != nil {
		log.Fatal(err)
	}

	// go-github と oauth2 で GitHub の認証
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	options := github.ListOptions{Page: 1, PerPage: 50}
	events, _, err := client.Activity.ListEventsPerformedByUser(oauth2.NoContext, "rnitame", false, &options)
	if err != nil {
		log.Fatal(err)
	}

	// 自分が実行したイベント一覧表示
	jst, _ := time.LoadLocation("Asia/Tokyo")
	for _, value := range events {
		fmt.Println(value.CreatedAt.In(jst))
	}
}
