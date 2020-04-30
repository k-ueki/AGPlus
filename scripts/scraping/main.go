package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Class struct {
	Name        string `db:"name"`
	Semester    string `db:"semester"`
	Credits     int    `db:"credits"`
	Teacher     string `db:"teacher"`
	Description string `db:"description"`
	Year        int    `db:"year"`
	DayAt       string `db:"day_at"`
	Campus      string `db:"campus"`
	// Prerequisite string
	// Evaluation string
}

var endpoint = "http://syllabus.aoyama.ac.jp/"

func gormConnect() *gorm.DB {
	CONNECT := "root:@tcp(localhost:3306)/agplus"
	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func setQueryFactors(page string) [][2]string {
	etarget := "ctl00%24CPH1%24rptPagerT%24ctl03%24lnkPage"

	factor := make([][2]string, 30)
	factor[0] = [2]string{"__EVENTTARGET", etarget}
	factor[1] = [2]string{"__EVENTARGUMENT", ""}
	factor[2] = [2]string{"__VIEWSTATEGENERATOR", "309A73F1"}
	factor[3] = [2]string{"YR", "2019"}
	factor[4] = [2]string{"BU", "BU1"}
	factor[5] = [2]string{"KW", ""}
	factor[6] = [2]string{"KM", ""}
	factor[7] = [2]string{"KI", ""}
	factor[8] = [2]string{"CP1", "on"}
	// factor[8] = [2]string{"CP4", "on"}
	factor[9] = [2]string{"YB1", "on"}
	factor[10] = [2]string{"YB1", "on"}
	factor[11] = [2]string{"YB2", "on"}
	factor[12] = [2]string{"YB3", "on"}
	factor[13] = [2]string{"YB4", "on"}
	factor[14] = [2]string{"YB5", "on"}
	factor[15] = [2]string{"YB6", "on"}
	factor[16] = [2]string{"JG1", "on"}
	factor[17] = [2]string{"JG2", "on"}
	factor[18] = [2]string{"JG3", "on"}
	factor[19] = [2]string{"JG4", "on"}
	factor[20] = [2]string{"JG5", "on"}
	factor[21] = [2]string{"JG6", "on"}
	factor[22] = [2]string{"JG7", "on"}
	factor[23] = [2]string{"GB1B_0", ""}
	factor[24] = [2]string{"GKB", ""}
	factor[25] = [2]string{"DL", "ja"}
	factor[26] = [2]string{"ST", ""}
	factor[27] = [2]string{"PG", page}
	factor[28] = [2]string{"PC", "395"}
	// factor[29] = [2]string{"PI", page_i}

	return factor

}

func main() {
	db := gormConnect()
	defer db.Close()

	db.Exec("SET foreign_key_checks = 0")
	db.Exec("TRUNCATE TABLE class")
	db.Exec("SET foreign_key_checks = 1")

	_, err := url.Parse(endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}

	count := 0

	for i := 1; i <= 395; i++ {
		url := ""

		factor := setQueryFactors(strconv.Itoa(i))

		for i, v := range factor {
			if i == 0 {
				url += "?"
			} else {
				url += "&"
			}
			url += fmt.Sprintf("%s=%s", v[0], v[1])
		}

		resp, err := http.Get(endpoint + url)
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

		selection := doc.Find("table.result > tbody > tr")

		// href, _ := selection.Attr("href")
		hrefs := [20]string{}
		dayats := [20]string{}
		campus := [20]string{}
		temp2 := [10]string{}

		selection.Each(func(index int, s *goquery.Selection) {
			sel := s.Find("td.col2 > span#CPH1_gvw_kensaku_lblJigen_0 > span")
			sel.Each(func(i int, se *goquery.Selection) {
				temp2[i] = se.Text()
			})

			tmp, _ := s.Find("td.col8 > a").Attr("href")
			hrefs[index] = tmp
			dayats[index] = temp2[1]
			campus[index] = temp2[0]
		})

		/////詳細のスクレイピング
		for i, href := range hrefs {
			if href == "" {
				break
			}
			res, err := http.Get(endpoint + href)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			doc2, err := goquery.NewDocumentFromReader(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			selection2 := doc2.Find("table.editTable > tbody > tr > td")

			temp := [10]string{}
			selection2.Each(func(index int, s *goquery.Selection) {
				temp[index] = s.Text()
			})

			credits, _ := strconv.Atoi(strings.TrimSpace(temp[4]))
			year, _ := strconv.Atoi(temp[0])
			class := &Class{
				Name:        strings.TrimSpace(temp[1]),
				Semester:    strings.TrimSpace(temp[3]),
				Credits:     credits,
				Teacher:     strings.TrimSpace(temp[5]),
				Description: strings.TrimSpace(temp[7]),
				Year:        year,
				DayAt:       dayats[i],
				Campus:      campus[i],
			}
			count++

			if count%100 == 0 {
				fmt.Println(count)
			}

			db.Table("class").Create(class)
		}
	}
}
