// Package generator allows you to easily generate invoices, delivery notes and quotations in GoLang.
package generator

import (
	"github.com/creasty/defaults"
	"github.com/phpdave11/gofpdf"
)

// New return a new documents with provided types and defaults
func New(docType string, options *Options) (*Document, error) {
	if err := defaults.Set(options); err != nil {
		return nil, err
	}

	doc := &Document{
		Options: options,
		Type:    docType,
	}

	doc.pdf = gofpdf.New("P", "mm", "A4", "")

	return doc, nil
}
