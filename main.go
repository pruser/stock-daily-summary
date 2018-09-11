package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pruser/stock-daily-summary/email"
	"github.com/pruser/stock-daily-summary/stock"
	"github.com/pruser/stock-daily-summary/summary"
)

func main() {
	Sender := os.Getenv("SUMMARY_SENDER")
	Recipient := os.Getenv("SUMMARY_RECIPIENT")
	URL := os.Getenv("SUMMARY_URL")
	SelectedRecords := os.Getenv("SUMMARY_RECORDS")
	Subject := os.Getenv("SUMMARY_SUBJECT")

	recordIDs := strings.Split(SelectedRecords, ",")
	for k, v := range recordIDs {
		recordIDs[k] = strings.TrimSpace(v)
	}

	parser := func(data []string) (stock.DailyData, error) {
		if len(data) < 7 {
			return stock.DailyData{}, fmt.Errorf("formatting error")
		}
		return stock.DailyData{Symbol: data[0], Open: data[2], Min: data[4], Max: data[3], Price: data[5], Volume: data[6]}, nil
	}

	source := stock.NewCSVDataSource(URL, parser)
	// aws credentials not set to enforce policy usage
	sender, err := email.NewSESEmailSender(nil)
	if err != nil {
		panic(err)
	}

	generator := summary.NewGenerator(source, sender, email.NewMessageSettings(Subject, Sender, Recipient))

	lambda.Start(func() error {
		return generator.Generate(recordIDs)
	})
}
