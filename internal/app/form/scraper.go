package form

import (
	"strings"

	"github.com/everitosan/sniim-scrapper/internal/app/utils"
	"github.com/gocolly/colly"
)

// Enums
type FormType int64

const (
	Form0Type FormType = iota
)

/*
* Form scraps an html and create a form struct with it's selects by category
* when finishes, every type represents a single input in the form
 */

type FormScraper struct {
	Inputs inputContainer // help to store all available options of a select
	Params FormParams     // Parametrs found in the form based in readme
}

func NewFormScraper() *FormScraper {
	return &FormScraper{
		Inputs: *NewInputContainer(),
	}
}

func (f *FormScraper) GetFormInputs(html *colly.HTMLElement, keyJoined string) {
	html.ForEach("table", func(_ int, table *colly.HTMLElement) {
		tableId := table.Attr("id")
		keys := strings.Split(keyJoined, utils.KeyCatalogueSeparator)

		switch tableId {
		case "tblDatos":
			f.Params = *NewFormParams(keys, Form0Type)
			From0Scraper(table, keys, f)
		case "tblFiltro":
			f.Params = *NewFormParams(keys, Form0Type)
			From0Scraper(table, keys, f)
		}
	})
}
