package codatypes

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"time"
)

var DatesLocation = time.FixedZone("UTC+3", 3*60*60)
var EnPrinter = message.NewPrinter(language.AmericanEnglish)
