package main

import (
	"os"
	"strings"
	"text/template"

	"github.com/SlyMarbo/rss"
)

func main() {
	feed, err := rss.Fetch("https://pablo.tools/atom.xml")
	if err != nil {
		panic(err)
	}

	posts := make(map[string][]*rss.Item)

	for _, v := range feed.Items {
		if strings.HasPrefix(v.Link, "https://pablo.tools/blog") {

			cat := strings.Title(strings.Split(v.Link, "/")[4])
			posts[cat] = append(posts[cat], v)
		}
	}

	t, err := template.ParseFiles("template")

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, posts)
	if err != nil {
		panic(err)
	}

}
