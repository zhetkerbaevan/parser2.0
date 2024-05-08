package parser

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/parser2.0/internal/model"
)

func checkForError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getHtml(url string) *http.Response { // server response
	res, err := http.Get(url)
	checkForError(err)

	if res.StatusCode > 400 {
		fmt.Println("Status Code: ", res.StatusCode)
	}

	return res
}

func GetData() []model.Automobile {
	var autos []model.Automobile
	url := "https://kolesa.kz/cars/almaty/"
	res := getHtml(url)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body) //converting html code to tree of objects, that will allow us to search throw it
	checkForError(err)

	doc.Find("div.a-list>div.a-list__item").Each(func(index int, item *goquery.Selection) {
		// id
		data_id, _ := item.Find("div").Attr("data-id")

		if len(data_id) > 0 {
			id, err := strconv.Atoi(data_id)
			checkForError(err)
			//name of model - получено
			h5 := item.Find("h5")
			entity := strings.TrimSpace(h5.Text())
			splittedModel := strings.Split(entity, "    ")[0]

			//price - получено
			span := item.Find("span.a-card__price")
			prWithSign := strings.TrimSpace(span.Text())
			pr := removeSpace(prWithSign)
			pr = strings.Trim(pr, "₸")
			price, err := strconv.Atoi(strings.Trim(pr, "₸"))
			checkForError(err)

			//year
			p := item.Find("p")
			desc := p.Text()
			splittedYear := strings.Split(desc, "\n")
			var year int
			for _, elem := range splittedYear {
				elem = strings.TrimSpace(elem)
				if len(elem) > 0 && checkDigit(elem) {
					year, err = strconv.Atoi(elem[:4])
					checkForError(err)
				}
			}
			var auto = model.Automobile{ID: id, Model: splittedModel, Price: price, Year: year}
			autos = append(autos, auto)
		}
	})
	return autos

}

func removeSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}

func checkDigit(desc string) bool {
	runes := []rune(desc)
	for _, elem := range runes {
		if unicode.IsNumber(elem) {
			return true
		}
	}
	return false
}
