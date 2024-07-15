package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	financialmodelingprep "github.com/grokify/go-financialmodelingprep"
	"github.com/grokify/mogo/encoding/jsonutil"
	"github.com/grokify/mogo/log/logutil"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	APIKey string `short:"k" long:"apikey" description:"API Key" required:"true"`
	Symbol string `short:"s" long:"symbol" description:"Stock Symbol" required:"true"`
	Period string `short:"p" long:"period" description:"Period as 'annual' or 'quarterly'" required:"true"`
	Limit  uint   `short:"l" long:"limit" description:"Limit"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}
	opts.Period = strings.ToLower(strings.TrimSpace(opts.Period))
	opts.Symbol = strings.ToLower(strings.TrimSpace(opts.Symbol))

	c := financialmodelingprep.NewClient(opts.APIKey)

	resp, err := c.GetIncomeStatement(
		opts.Symbol,
		opts.Period,
		opts.Limit)
	logutil.FatalErr(err)
	b, err := jsonutil.PrintReaderIndent(resp.Body, "", "  ")
	logutil.FatalErr(err)

	err = os.WriteFile(fmt.Sprintf(".income-statement_%s_%s.json", opts.Symbol, opts.Period), b, 0600)

	logutil.FatalErr(err)
	fmt.Println("DONE")
}
