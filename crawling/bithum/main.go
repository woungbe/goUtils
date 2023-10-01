package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	StrEmun := []string{"입출금 일시 중지", "출금 일시 중지", "투자유의종목 지정", "투자유의종목 해제", "입출금 일시 중지 안내", "거래 유의 안내"}
	url := "https://cafe.bithumb.com/view/boards/43?keyword=&noticeCategory=" // 크롤링할 웹 페이지의 URL을 여기에 넣어주세요

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// doc.Find("a").Each(func(index int, element *goquery.Selection) {
	// 	linkText := element.Text()
	// 	// class="one-line"
	// 	linkURL, _ := element.Attr("href")
	// 	fmt.Printf("링크 텍스트: %s\n링크 URL: %s\n", linkText, linkURL)
	// })

	doc.Find(".col-20.col-md-3").Each(func(index int, element *goquery.Selection) {
		// fmt.Println(" : ", e.Text())

		// linkText

		// linkText.Find("tr").Each(func(index2 int, element2 *goquery.Selection) {
		// 	if index2 == 0 {
		// 		fmt.Println("element :", element2)
		// 	}
		// })

		// element.Children().Find("td").Each(func(i int, e *goquery.Selection) {
		// 	if i == 0 {
		// 		linkText := e.Text()
		// 		fmt.Printf("텍스트: %s\n", linkText)
		// 	}
		// })

		// linkURL, _ := element.Attr("href")
		b := false
		element.Find("td").Each(func(index2 int, element2 *goquery.Selection) {
			if index2 == 0 {
				// fmt.Println(element2.Text())
				if element2.Text() != "■" {
					// fmt.Println("출력하자!!")
					b = true
				}
			}

			if index2 == 1 && b {
				for _, v := range StrEmun {
					if strings.Contains(element2.Text(), v) {
						fmt.Println("제목 : ", strings.Trim(element2.Text(), " "))
						break
					}
				}
			}
		})

		// html, err := element.Html()
		// if err != nil {
		// 	fmt.Println("에러났다꿍~")
		// }
		// fmt.Println("html : ", html)

		// linkText := element.Text()
		// fmt.Printf("링크 텍스트: %s\n", linkText)

	})

}
