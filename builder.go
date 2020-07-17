package main

import (
	"github.com/SlyMarbo/rss"
	"os"
	"strings"
	"text/template"
)

func main() {
	feed, err := rss.Fetch("https://pablo.tools/index.xml")
	if err != nil {
		// handle error.
	}

	// ... Some time later ...

	// fmt.Println(feed.Nickname)
	// fmt.Println(feed.Title)

	posts := make(map[string][]*rss.Item)

	for _, v := range feed.Items {
		if strings.HasPrefix(v.Link, "https://pablo.tools/posts") {
			cat := strings.Title(strings.Split(v.Link, "/")[4])
			posts[cat] = append(v, posts[cat])
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
