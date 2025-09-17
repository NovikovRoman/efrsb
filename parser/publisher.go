package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PublisherCompany struct {
	Name string
	Ogrn string
	Inn  string
}

type PublisherPerson struct {
	Lastname              string
	Firstname             string
	Middlename            string
	Inn                   string
	Snils                 string
	CorrespondenceAddress string
	Email                 string
}

type Sro struct {
	Name    string
	Ogrn    string
	Inn     string
	Address string
}

type Publisher struct {
	Company *PublisherCompany
	Person  *PublisherPerson
	Sro     *Sro
}

func FindPublisher(content string) *Publisher {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
	s := doc.Find("Publisher").First()
	if s.Length() == 0 {
		return nil
	}

	p := &Publisher{
		Person: _publisherPerson(s),
	}
	if p.Person == nil {
		p.Company = _publisherCompany(s)
	}

	elSro := s.Find("Sro").First()
	if elSro.Length() == 0 {
		return p
	}

	p.Sro = &Sro{
		Name:    elSro.Find("Name").First().Text(),
		Ogrn:    elSro.Find("Ogrn").First().Text(),
		Inn:     elSro.Find("Inn").First().Text(),
		Address: elSro.Find("Address").First().Text(),
	}
	return p
}

func _publisherPerson(s *goquery.Selection) *PublisherPerson {
	elFio := s.Find("Fio")
	if elFio.First().Length() == 0 {
		return nil
	}

	return &PublisherPerson{
		Lastname:              elFio.Find("LastName").First().Text(),
		Firstname:             elFio.Find("FirstName").First().Text(),
		Middlename:            elFio.Find("MiddleName").First().Text(),
		Inn:                   s.Find("Inn").First().Text(),
		Snils:                 s.Find("Snils").First().Text(),
		CorrespondenceAddress: s.Find("CorrespondenceAddress").First().Text(),
		Email:                 s.Find("Email").First().Text(),
	}
}

func _publisherCompany(s *goquery.Selection) *PublisherCompany {
	elOgrn := s.Find("Ogrn")
	if elOgrn.First().Length() == 0 {
		return nil
	}
	return &PublisherCompany{
		Name: s.Find("Name").First().Text(),
		Ogrn: elOgrn.Text(),
		Inn:  s.Find("Inn").First().Text(),
	}
}
