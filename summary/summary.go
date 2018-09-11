package summary

import (
	"fmt"

	"github.com/pruser/stock-daily-summary/email"
	"github.com/pruser/stock-daily-summary/stock"
)

type SummaryGenerator struct {
	source   stock.DataSource
	sender   email.Sender
	settings *email.MessageSettings
}

func NewGenerator(source stock.DataSource, sender email.Sender, settings *email.MessageSettings) *SummaryGenerator {
	return &SummaryGenerator{source: source, sender: sender, settings: settings}
}

func (g SummaryGenerator) Generate(symbols []string) error {
	values, err := g.source.Get(symbols)
	if err != nil {
		return err
	}

	emailContent := ""
	for _, v := range values {
		emailContent += fmt.Sprintf("%s\n", v)
	}

	err = g.sender.Send(*g.settings, emailContent)
	if err != nil {
		return err
	}
	return nil
}
