package financialmodelingprep

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/grokify/gocharts/v2/charts/wchart/sts2wchart"
	"github.com/grokify/gocharts/v2/data/timeseries"
	"github.com/grokify/mogo/time/timeutil"
	"github.com/grokify/mogo/type/slicesutil"
	chart "github.com/wcharczuk/go-chart/v2"
)

const PeriodFY = "FY"

// https://financialmodelingprep.com/api/v3/income-statement/AAPL?limit=120&apikey=mykey

type IncomeStatementOpts struct {
	APIKEY string `url:"apikey"`
	Limit  uint   `url:"limit"`
	Period string `url:"period"`
}

type IncomeStatements []IncomeStatement

func IncomeStatementsReadFile(filename string) (IncomeStatements, error) {
	if f, err := os.Open(filename); err != nil {
		return IncomeStatements{}, err
	} else if res, err := IncomeStatementParse(f); err != nil {
		return IncomeStatements{}, err
	} else {
		return res, f.Close()
	}
}

func IncomeStatementParse(r io.Reader) (IncomeStatements, error) {
	if b, err := io.ReadAll(r); err != nil {
		return IncomeStatements{}, err
	} else {
		return IncomeStatementsParseBytes(b)
	}
}

func IncomeStatementsParseBytes(data []byte) (IncomeStatements, error) {
	incs := IncomeStatements{}
	err := json.Unmarshal(data, &incs)
	return incs, err
}

/*
// RevenueTimeSeries returns a time series of revnue. The periods of the income statements
// must be the same.
func (iss IncomeStatements) RevenueTimeSeries() (*timeseries.TimeSeries, error) {
	periods := iss.Periods(true)
	if len(periods) > 1 {
		return nil, errors.New("too many time periods")
	} else if len(periods) == 0 {
		return nil, errors.New("no time periods")
	}
	ts := timeseries.NewTimeSeries("income statements")
	ts.IsFloat = true
	if periods[0] == PeriodFY {
		ts.Interval = timeutil.IntervalYear
	} else {
		return nil, fmt.Errorf("interval not supported (%s)", periods[0])
	}
	for _, is := range iss {
		dt, err := time.Parse(timeutil.RFC3339FullDate, is.Date)
		if err != nil {
			return nil, err
		}
		ts.AddFloat64(dt, is.Revenue)
	}
	return &ts, nil
}
*/

func (iss IncomeStatements) Periods(dedupe bool) []string {
	var periods []string
	for _, is := range iss {
		periods = append(periods, is.Period)
	}
	if dedupe {
		return slicesutil.Dedupe(periods)
	}
	return periods
}

func (iss IncomeStatements) ChartAnnual() (chart.Chart, error) {
	tss, err := iss.TimeSeriesSetAnnual()
	if err != nil {
		return chart.Chart{}, err
	}
	opts := sts2wchart.DefaultLineChartOpts()
	opts.Interval = timeutil.IntervalYear
	opts.XAxisTickInterval = timeutil.IntervalYear
	opts.XAxisGridInterval = timeutil.IntervalYear

	return sts2wchart.TimeSeriesSetToLineChart(*tss, opts)
}

func (iss IncomeStatements) TimeSeriesSetAnnual() (*timeseries.TimeSeriesSet, error) {
	if len(iss) == 0 {
		return nil, errors.New("no income statements")
	}
	tss := timeseries.NewTimeSeriesSet(iss[0].Symbol)
	tss.Interval = timeutil.IntervalYear
	tss.IsFloat = false
	for _, is := range iss {
		yyyy, err := is.CalendarYearInt()
		if err != nil {
			return nil, err
		}
		dt := time.Date(yyyy, time.January, 1, 0, 0, 0, 0, time.UTC)
		tss.AddInt64(Revenue, dt, int64(is.Revenue))
		tss.AddInt64(CostOfRevenue, dt, int64(is.CostOfRevenue))
		tss.AddInt64(ResearchAndDevelopmentExpenses, dt, int64(is.ResearchAndDevelopmentExpenses))
		tss.AddInt64(GeneralAndAdministrativeExpenses, dt, int64(is.GeneralAndAdministrativeExpenses))
		tss.AddInt64(SellingAndMarketingExpenses, dt, int64(is.SellingAndMarketingExpenses))
		tss.AddInt64(OtherExpenses, dt, int64(is.OtherExpenses))
		tss.AddInt64(OperatingExpenses, dt, int64(is.OperatingExpenses))
		tss.AddInt64(InterestExpense, dt, int64(is.InterestExpense))
		tss.AddInt64(DepreciationAndAmortization, dt, int64(is.DepreciationAndAmortization))
	}

	tss.Inflate()
	tss.Order = Ordered()

	return &tss, nil
}

func (iss IncomeStatements) FinancialsMatrixAnnual() ([][]int, error) {
	matrix := [][]int{}
	tss, err := iss.TimeSeriesSetAnnual()
	if err != nil {
		return matrix, err
	}
	for _, dt := range tss.Times {
		if !timeutil.NewTimeMore(dt, time.Sunday).IsYearStart() {
			return matrix, errors.New("non-year start time")
		}
		rfc3339 := dt.Format(time.RFC3339)
		row := []int{dt.Year()}
		for _, seriesName := range tss.Order {
			item, err := tss.Item(seriesName, rfc3339)
			if err != nil {
				row = append(row, 0)
			} else {
				row = append(row, int(item.Int64()))
			}
		}
		matrix = append(matrix, row)
	}
	return matrix, nil
}

const (
	Revenue                          = "Revenue"
	CostOfRevenue                    = "Cost of Revenue"
	ResearchAndDevelopmentExpenses   = "Research and Development Expenses"
	GeneralAndAdministrativeExpenses = "General and Administrative Expenses"
	SellingAndMarketingExpenses      = "Selling and Marketing Expenses"
	OtherExpenses                    = "Other Expenses"
	OperatingExpenses                = "Operating Expenses"
	InterestExpense                  = "InterestExpense"
	DepreciationAndAmortization      = "Depreciation and Amortization"
)

func Ordered() []string {
	return []string{
		Revenue,
		CostOfRevenue,
		ResearchAndDevelopmentExpenses,
		GeneralAndAdministrativeExpenses,
		SellingAndMarketingExpenses,
		OtherExpenses,
		OperatingExpenses,
		InterestExpense,
		DepreciationAndAmortization}
}

type IncomeStatement struct {
	CalendarYear                     string  `json:"calendarYear"`
	Date                             string  `json:"date"`
	FillingDate                      string  `json:"fillingDate"`
	Period                           string  `json:"period'`
	Symbol                           string  `json:"symbol"`
	Revenue                          float64 `json:"revenue"`
	CostOfRevenue                    float64 `json:"costOfRevenue"`
	GrossProfit                      float64 `json:"grossProfit"`
	ResearchAndDevelopmentExpenses   float64 `json:"researchAndDevelopmentExpenses"`
	GeneralAndAdministrativeExpenses float64 `json:"generalAndAdministrativeExpenses"`
	SellingAndMarketingExpenses      float64 `json:"sellingAndMarketingExpenses"`
	OtherExpenses                    float64 `json:"otherExpenses"`
	OperatingExpenses                float64 `json:"operatingExpenses"`
	CostAndExpenses                  float64 `json:"costAndExpenses"`
	InterestIncome                   float64 `json:"interestIncome"`
	InterestExpense                  float64 `json:"interestExpense"`
	DepreciationAndAmortization      float64 `json:"depreciationAndAmortization"`
	EBITDA                           float64 `json:"ebitda"`
	EBITDAratio                      float64 `json:"ebitdaratio"`
	OperatingIncome                  float64 `json:"operatingIncome"`
	OperatingIncomeRatio             float64 `json:"operatingIncomeRatio"`
	TotalOtherIncomeExpensesNet      float64 `json:"totalOtherIncomeExpensesNet"`
	IncomeBeforeTax                  float64 `json:"incomeBeforeTax"`
	IncomeBeforeTaxRatio             float64 `json:"incomeBeforeTaxRatio"`
	IncomeTaxExpense                 float64 `json:"incomeTaxExpense"`
	NetIncome                        float64 `json:"netIncome"`
}

func (is *IncomeStatement) CalendarYearInt() (int, error) {
	return strconv.Atoi(is.CalendarYear)
}

/*
  {
    "date": "2021-12-31",
    "symbol": "JPM",
    "reportedCurrency": "USD",
    "cik": "0000019617",
    "fillingDate": "2022-02-22",
    "acceptedDate": "2022-02-22 16:56:31",
    "calendarYear": "2021",
    "period": "FY",
    "revenue": 130898000000,
    "costOfRevenue": 0,
    "grossProfit": 130898000000,
    "grossProfitRatio": 1,
    "researchAndDevelopmentExpenses": 0,
    "generalAndAdministrativeExpenses": 19755000000,
    "sellingAndMarketingExpenses": 3036000000,
    "sellingGeneralAndAdministrativeExpenses": 65874000000,
    "otherExpenses": 0,
    "operatingExpenses": 71336000000,
    "costAndExpenses": 71336000000,
    "interestIncome": 57864000000,
    "interestExpense": 5553000000,
    "depreciationAndAmortization": -59988000000,
    "ebitda": -2257000000,
    "ebitdaratio": -0.017242433,
    "operatingIncome": 59988000000,
    "operatingIncomeRatio": 0.4582804932,
    "totalOtherIncomeExpensesNet": -426000000,
    "incomeBeforeTax": 59562000000,
    "incomeBeforeTaxRatio": 0.4550260508,
    "incomeTaxExpense": 11228000000,
    "netIncome": 46503000000,
    "netIncomeRatio": 0.3552613485,
    "eps": 15.39,
    "epsdiluted": 15.36,
    "weightedAverageShsOut": 3027539062,
    "weightedAverageShsOutDil": 3026600000,
    "link": "https://www.sec.gov/Archives/edgar/data/19617/000001961722000272/0000019617-22-000272-index.htm",
    "finalLink": "https://www.sec.gov/Archives/edgar/data/19617/000001961722000272/jpm-20211231.htm"
  },
*/
