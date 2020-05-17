package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"

	"github.com/k-ueki/AGPlus/server/common"
)

type (
	Dep struct {
		ID     int
		Name   string
		Campus int
	}

	Course struct {
		DepID int
		Name  string
	}
)

var endPoint = "https://www.aoyama.ac.jp/faculty/"
var courseBase = "https://www.aoyama.ac.jp"

func main() {
	db, err := common.NewDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	db.Exec("SET foreign_key_checks = 0")
	db.Exec("TRUNCATE TABLE department")
	db.Exec("TRUNCATE TABLE course")
	db.Exec("SET foreign_key_checks = 1")

	_, err = url.Parse(endPoint)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Get(endPoint)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	selection := doc.Find("div.subjectbox")
	dep := [11]Dep{}
	course := make(map[int][]string)

	selection.Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			selection = s.Find("ul > li")
			selection.Each(func(j int, ss *goquery.Selection) {
				dep[j].ID = j + 1
				dep[j].Name = ss.Find("p.title").Text()

				courseURL, _ := ss.Find("a").Attr("href")

				resp, err := http.Get(courseBase + courseURL)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer resp.Body.Close()

				doc, err := goquery.NewDocumentFromReader(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				doc.Find("div.table_typeG > div.tr").Each(func(ii int, ssss *goquery.Selection) {
					if ii == 2 {
						campus := ssss.Find("a.link").Text()
						if campus == "青山キャンパス" {
							dep[j].Campus = 1
						} else {
							dep[j].Campus = 2
						}
					}
				})

				doc.Find("div.box > ul > li.box_item").Each(func(jj int, sssss *goquery.Selection) {
					if sssss.Find("h5").Text() != "" {
						course[j+1] = append(course[j+1], sssss.Find("h5").Text())
					}
				})
			})
		}
	})

	fmt.Println(dep)
	for _, v := range dep {
		db.Table("department").Save(&v)
	}

	fmt.Println(course)
	for i, v := range course {
		for _, vv := range v {
			course := Course{
				DepID: i,
				Name:  vv,
			}
			db.LogMode(true).Table("course").Save(course)
		}
	}

}
