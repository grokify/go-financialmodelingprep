package main

import (
	"fmt"
	"os"

	fmp "github.com/grokify/go-financialmodelingprep"
	"github.com/grokify/gocharts/v2/charts/google"
	"github.com/grokify/gocharts/v2/charts/wchart"
	"github.com/grokify/gocharts/v2/data/timeseries"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
)

func main() {
	file := "fmp_income-statement_aapl_ann_2021.json"
	symbol := "AAPL"

	incs, err := fmp.ParseIncomeStatementsFile(file)
	logutil.FatalErr(err)

	fmtutil.PrintJSON(incs)

	tss, err := incs.TimeSeriesSetAnnual()
	logutil.FatalErr(err)

	fmtutil.PrintJSON(tss)

	tbl, err := tss.Table(&timeseries.TimeSeriesSetTableOpts{
		TimeColumnTitle: "Year",
		// 	FuncFormatTime:  func(dt time.Time) string { return dt.Format("2006") },
	})
	logutil.FatalErr(err)
	err = tbl.WriteXLSX("_income-statement.xlsx", symbol)
	logutil.FatalErr(err)

	t2, err := tbl.Transpose()
	logutil.FatalErr(err)
	err = t2.WriteXLSX("_income-statement_transpose.xlsx", symbol)

	lcm, err := google.LineChartMaterialFromTimeSeriesSet(*tss, "Year")
	lcm.Title = fmt.Sprintf("%s Income Statement", symbol)
	lcm.Width = 1200
	logutil.FatalErr(err)

	pageHTML := google.LineChartMaterialPage(lcm)
	fmt.Println(pageHTML)

	err = os.WriteFile(fmt.Sprintf("_financials_%s.html", symbol), []byte(pageHTML), 0600)
	logutil.FatalErr(err)

	if 1 == 0 {
		c, err := incs.ChartAnnual()
		logutil.FatalErr(err)

		err = wchart.WritePNG(fmt.Sprintf("_%s.png", symbol), c)
		logutil.FatalErr(err)
	}

	fmt.Println("DONE")
}
