package parser

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type MonetaryObligation struct {
	CreditorName   string
	CreditorRegion string
	Content        string
	Basis          string
	TotalSum       float64
	DebtSum        float64
	PenaltySum     float64
}

func FindMonetaryObligations(content string) (m []MonetaryObligation) {
	m = make([]MonetaryObligation, 0)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
	query := "MessageInfo StartOfExtrajudicialBankruptcy CreditorsNonFromEntrepreneurship MonetaryObligations MonetaryObligation"
	doc.Find(query).Each(func(i int, s *goquery.Selection) {
		total, _ := strconv.ParseFloat(removeSpace(s.Find("TotalSum").Text()), 64)
		debtSum, _ := strconv.ParseFloat(removeSpace(s.Find("DebtSum").Text()), 64)
		penaltySum, _ := strconv.ParseFloat(removeSpace(s.Find("PenaltySum").Text()), 64)

		m = append(m, MonetaryObligation{
			CreditorName:   removeNewLine(s.Find("CreditorName").Text()),
			CreditorRegion: removeNewLine(s.Find("CreditorRegion").Text()),
			Content:        removeNewLine(s.Find("Content").Text()),
			Basis:          removeNewLine(s.Find("Basis").Text()),
			TotalSum:       total,
			DebtSum:        debtSum,
			PenaltySum:     penaltySum,
		})
	})

	return
}
