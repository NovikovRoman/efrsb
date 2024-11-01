package parser

import (
	"regexp"
	"strings"
	"time"
)

var (
	reSpace = regexp.MustCompile(`\s+`)
)

func removeNewLine(s string) string {
	return strings.TrimSpace(reSpace.ReplaceAllString(s, " "))
}

func removeSpace(s string) string {
	return reSpace.ReplaceAllString(s, "")
}

func DateTime(datetime, layout string) (t time.Time) {
	if datetime != "" {
		t, _ = time.Parse(layout, datetime)
	}
	return
}
