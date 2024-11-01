package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PersonCategory struct {
	Code        string
	Description string
}

func FindPersonCategory(content string) *PersonCategory {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
	s := doc.Find("MessageInfo StartOfExtrajudicialBankruptcy PersonCategory").First()
	if s == nil || s.Length() == 0 {
		return nil
	}

	return &PersonCategory{
		Code:        removeSpace(s.Find("Code").First().Text()),
		Description: removeNewLine(s.Find("Description").First().Text()),
	}
}
