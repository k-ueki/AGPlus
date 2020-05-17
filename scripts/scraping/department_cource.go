package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"

	"github.com/k-ueki/AGPlus/server/common"
)

type (
	Faculty struct {
		Name         string
		Type         int
		CampusID     int
		FacultyID    int
		DepartmentID int
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
	faculty := []Faculty{}
	course := make(map[int][]string)

	columnCount := 0
	selection.Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			selection = s.Find("ul > li")
			selection.Each(func(j int, ss *goquery.Selection) {
				//dep[j].ID = j + 3
				//dep[j].Name = ss.Find("p.title").Text()

				//faculty[j].FacultyID = j + 3
				//faculty[j].Name = ss.Find("p.title").Text()
				name := ss.Find("p.title").Text()
				columnCount = j + 3

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

				campusID := 0
				doc.Find("div.table_typeG > div.tr").Each(func(ii int, ssss *goquery.Selection) {
					if ii == 2 {
						campus := ssss.Find("a.link").Text()
						if campus == "青山キャンパス" {
							//dep[j].Campus = 1
							//faculty[j].CampusID = 1
							campusID = 1
						} else {
							//dep[j].Campus = 2
							//faculty[j].CampusID = 2
							campusID = 2
						}
					}
				})

				doc.Find("div.box > ul > li.box_item").Each(func(jj int, sssss *goquery.Selection) {
					if sssss.Find("h5").Text() != "" {
						course[j+3] = append(course[j+3], sssss.Find("h5").Text())
					}
				})

				f := Faculty{
					Name:      name,
					Type:      2,
					CampusID:  campusID,
					FacultyID: j + 3,
				}
				faculty = append(faculty, f)

			})
		}
	})

	for i, v := range course {
		for _, vv := range v {
			columnCount++

			f := Faculty{}
			if i <= 9 {
				f = Faculty{
					Name:         vv,
					Type:         3,
					CampusID:     1,
					FacultyID:    i,
					DepartmentID: columnCount,
				}
			} else {
				f = Faculty{
					Name:         vv,
					Type:         3,
					CampusID:     2,
					FacultyID:    i,
					DepartmentID: columnCount,
				}
			}
			faculty = append(faculty, f)
			//db.LogMode(true).Table("course").Save(course)
		}
	}

	fmt.Println(faculty)

	for _, v := range faculty {
		db.LogMode(true).Table("faculty").Save(&v)
	}
}
