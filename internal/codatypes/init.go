package codatypes

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"time"
)

var DatesLocation *time.Location

var EnPrinter *message.Printer

func init() {
	DatesLocation = time.FixedZone("UTC+3", 3*60*60)
	EnPrinter = message.NewPrinter(language.AmericanEnglish)
}
