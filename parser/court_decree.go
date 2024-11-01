package parser

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type CourtDecree struct {
	ID              string
	Name            string
	Number          string
	DecisionDateRaw string
	DecisionDate    time.Time
}

func FindCourtDecree(content string) (c *CourtDecree) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
	s := doc.Find("MessageInfo CourtDecision CourtDecree").First()
	if s == nil || s.Length() == 0 {
		return
	}

	date := removeNewLine(s.Find("DecisionDate").First().Text())
	c = &CourtDecree{
		ID:              removeSpace(s.Find("CourtId").First().Text()),
		Name:            removeNewLine(s.Find("CourtName").First().Text()),
		Number:          removeNewLine(s.Find("FileNumber").First().Text()),
		DecisionDateRaw: date,
		DecisionDate:    DateTime(date, time.DateOnly),
	}

	if c.DecisionDateRaw != "" {
		c.DecisionDate = DateTime(c.DecisionDateRaw, time.DateOnly)
	}
	return
}
