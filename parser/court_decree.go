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
	DecisionType    *DecisionType
}

type DecisionType struct {
	Name string
	ID   string
}

func FindCourtDecree(content string) (c *CourtDecree) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
	s := doc.Find("MessageInfo CourtDecision").First()
	if s == nil || s.Length() == 0 {
		return
	}

	sCourt := s.Find("CourtDecree")
	date := removeNewLine(s.Find("DecisionDate").First().Text())
	c = &CourtDecree{
		ID:              removeSpace(sCourt.Find("CourtId").First().Text()),
		Name:            removeNewLine(sCourt.Find("CourtName").First().Text()),
		Number:          removeNewLine(sCourt.Find("FileNumber").First().Text()),
		DecisionDateRaw: date,
		DecisionDate:    DateTime(date, time.DateOnly),
	}

	if c.DecisionDateRaw != "" {
		c.DecisionDate = DateTime(c.DecisionDateRaw, time.DateOnly)
	}

	sDecisionType := s.Find("DecisionType").First()
	if sDecisionType.Length() > 0 {
		c.DecisionType = &DecisionType{
			Name: removeNewLine(sDecisionType.AttrOr("name", "")),
			ID:   removeSpace(sDecisionType.AttrOr("id", "")),
		}
	}
	return
}
