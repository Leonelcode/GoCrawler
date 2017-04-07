package main

import (
    // import standard libraries
    "fmt"
    "log"

    // import third party libraries
    "github.com/PuerkitoBio/goquery"
)

func postScrape() {
    doc, err := goquery.NewDocument("https://es.wikipedia.org/wiki/Vanuatu")
    if err != nil {
        log.Fatal(err)
    }

    // use CSS selector found with the browser inspector
    // for each, use index and item
    doc.Find(".mw-body-content").Each(func(index int, item *goquery.Selection) {
        title := item.Text()
        linkTag := item.Find("p")
        link, _ := linkTag.Attr("href")
        fmt.Printf("Post #%d: %s - %s\n", index, title, link)
    })
}

func main() {
    postScrape()
}