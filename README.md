# GoCrawler

GoCrawler is a golang tool for scraping data of websites using the golang library Goquery.

## Features

- Function to get all links related to a DOM element.
- Function to obtain the metadata of a web.
- Function to extract some element of the DOM.

## Using Goquery

```
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
```

## Credits
- [Leonel Vega](https://twitter.com/leonel_py)

## License

[GNU](https://www.gnu.org/licenses/gpl-3.0.en.html)
