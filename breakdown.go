package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func readHtml(filename string) (string, error) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func ParseHTML(text string) (data []string) {
	tkn := html.NewTokenizer(strings.NewReader(text))

	var vals []string
	var isDt, isDd bool
	for {
		tt := tkn.Next()

		switch {
		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:

			t := tkn.Token()

			if t.Data == "dt" {
				isDt = t.Data == "dt"
			}
			if t.Data == "dd" {
				isDd = t.Data == "dd"
			}
		case tt == html.TextToken:
			t := tkn.Token()

			if isDt || isDd {
				//add dt/dd to vals with commas in between
				vals = append(vals, t.Data)
			}
			isDt = false
			isDd = false
		}
	}
}

func breakdown() {
	filename := "templates/players.html"
	text, err := readHtml(filename)
	if err != nil {
		panic(err)
	}
	data := ParseHTML(text)
	fmt.Println(data)

}
