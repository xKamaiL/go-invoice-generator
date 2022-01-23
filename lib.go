package generator

import "github.com/phpdave11/gofpdf"

var trFunc func(string) string

func encodeString(str string) string {
	if trFunc == nil {
		pdf := gofpdf.New("P", "mm", "A4", "")
		trFunc = pdf.UnicodeTranslatorFromDescriptor("")
	}

	return trFunc(str)
}

func (d *Document) typeAsString() string {
	if d.Type == Invoice {
		return d.Options.TextTypeInvoice
	}

	if d.Type == Quotation {
		return d.Options.TextTypeQuotation
	}

	return d.Options.TextTypeDeliveryNote
}
