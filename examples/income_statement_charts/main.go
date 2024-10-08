package main

import (
	"fmt"
	"os"

	fmp "github.com/grokify/go-financialmodelingprep"
	"github.com/grokify/gocharts/v2/charts/google/linechart"
	"github.com/grokify/gocharts/v2/charts/wchart"
	"github.com/grokify/gocharts/v2/data/timeseries"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
)

func main() {
	file := "fmp_income-statement_aapl_ann_2021.json"
	symbol := "AAPL"

	incs, err := fmp.IncomeStatementsReadFile(file)
	logutil.FatalErr(err)

	fmtutil.MustPrintJSON(incs)

	tss, err := incs.TimeSeriesSetAnnual()
	logutil.FatalErr(err)

	fmtutil.MustPrintJSON(tss)

	ts, ok := tss.Series["Revenue"]
	if ok {
		fmtutil.MustPrintJSON(ts)
		yoy, err := ts.TimeSeriesYearYOY("")
		logutil.FatalErr(err)
		fmt.Printf("YOY")
		fmtutil.MustPrintJSON(yoy)
		panic("Z")
	}

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
	logutil.FatalErr(err)

	lcm, err := linechart.LineChartMaterialFromTimeSeriesSet(*tss, "Year")
	logutil.FatalErr(err)

	lcm.Title = fmt.Sprintf("%s Income Statement", symbol)
	lcm.Width = 1200
	logutil.FatalErr(err)

	pageHTML := linechart.LineChartMaterialPage(lcm)
	fmt.Println(pageHTML)

	err = os.WriteFile(fmt.Sprintf("_financials_%s.html", symbol), []byte(pageHTML), 0600)
	logutil.FatalErr(err)

	if 1 == 0 {
		c, err := incs.ChartAnnual()
		logutil.FatalErr(err)

		err = wchart.WritePNGFile(fmt.Sprintf("_%s.png", symbol), c)
		logutil.FatalErr(err)
	}

	fmt.Println("DONE")
}
