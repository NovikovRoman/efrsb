package parser

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type MonetaryObligation struct {
	CreditorName     string
	CreditorRegion   string
	CreditorLocation string
	Content          string
	Basis            string
	TotalSum         float64
	DebtSum          float64
	PenaltySum       float64
}

type MonetaryObligations struct {
	Entrepreneurship    []MonetaryObligation
	NonEntrepreneurship []MonetaryObligation
}

func FindMonetaryObligations(content string) (m MonetaryObligations) {
	m = MonetaryObligations{
		Entrepreneurship:    []MonetaryObligation{},
		NonEntrepreneurship: []MonetaryObligation{},
	}
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
	s := doc.Find("MessageInfo StartOfExtrajudicialBankruptcy").First()
	if s == nil || s.Length() == 0 {
		return
	}

	m.Entrepreneurship = findMonetaryObligations(s, "CreditorsFromEntrepreneurship")
	m.NonEntrepreneurship = findMonetaryObligations(s, "CreditorsNonFromEntrepreneurship")
	return
}

func findMonetaryObligations(s *goquery.Selection, category string) (m []MonetaryObligation) {
	m = []MonetaryObligation{}
	s.Find(category + " MonetaryObligations MonetaryObligation").Each(func(i int, s *goquery.Selection) {
		total, _ := strconv.ParseFloat(removeSpace(s.Find("TotalSum").Text()), 64)
		debtSum, _ := strconv.ParseFloat(removeSpace(s.Find("DebtSum").Text()), 64)
		penaltySum, _ := strconv.ParseFloat(removeSpace(s.Find("PenaltySum").Text()), 64)

		m = append(m, MonetaryObligation{
			CreditorName:     removeNewLine(s.Find("CreditorName").Text()),
			CreditorRegion:   removeNewLine(s.Find("CreditorRegion").Text()),
			CreditorLocation: removeNewLine(s.Find("CreditorLocation").Text()),
			Content:          removeNewLine(s.Find("Content").Text()),
			Basis:            removeNewLine(s.Find("Basis").Text()),
			TotalSum:         total,
			DebtSum:          debtSum,
			PenaltySum:       penaltySum,
		})
	})
	return
}
