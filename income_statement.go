package financialmodelingprep

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/grokify/gocharts/v2/charts/wchart/sts2wchart"
	"github.com/grokify/gocharts/v2/data/timeseries"
	"github.com/grokify/mogo/time/timeutil"
	chart "github.com/wcharczuk/go-chart/v2"
)

const (
	APIEndpoint        = "https://financialmodelingprep.com"
	IncomeStatementURL = "api/v3/income-statement"
)

// https://financialmodelingprep.com/api/v3/income-statement/AAPL?limit=120&apikey=mykey

type IncomeStatementOpts struct {
	APIKEY string `url:"apikey"`
	Limit  uint   `url:"limit"`
	Period string `url:"period"`
}

type IncomeStatements []IncomeStatement

func ParseIncomeStatementsFile(filename string) (IncomeStatements, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return IncomeStatements{}, err
	}
	return ParseIncomeStatements(data)
}

func ParseIncomeStatements(data []byte) (IncomeStatements, error) {
	incs := IncomeStatements{}
	err := json.Unmarshal(data, &incs)
	return incs, err
}

func (iss IncomeStatements) ChartAnnual() (chart.Chart, error) {
	tss, err := iss.TimeSeriesSetAnnual()
	if err != nil {
		return chart.Chart{}, err
	}
	opts := sts2wchart.DefaultLineChartOpts()
	opts.Interval = timeutil.Year
	opts.XAxisTickInterval = timeutil.Year
	opts.XAxisGridInterval = timeutil.Year

	return sts2wchart.TimeSeriesSetToLineChart(*tss, opts)
}

func (iss IncomeStatements) TimeSeriesSetAnnual() (*timeseries.TimeSeriesSet, error) {
	if len(iss) == 0 {
		return nil, errors.New("no income statements")
	}
	tss := timeseries.NewTimeSeriesSet(iss[0].Symbol)
	tss.Interval = timeutil.Year
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
		if !timeutil.IsYearStart(dt) {
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
	Symbol                           string  `json:"symbol"`
	FillingDate                      string  `json:"fillingDate"`
	CalendarYear                     string  `json:"calendarYear"`
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
