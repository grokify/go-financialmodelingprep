# IncomeStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedDate** | Pointer to **string** |  | [optional] 
**CalendarYear** | Pointer to **string** |  | [optional] 
**Cik** | Pointer to **string** |  | [optional] 
**CostAndExpenses** | Pointer to **float32** |  | [optional] 
**CostOfRevenue** | Pointer to **float32** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**DepreciationAndAmortization** | Pointer to **float32** |  | [optional] 
**Ebitda** | Pointer to **float32** |  | [optional] 
**Ebitdaratio** | Pointer to **float32** |  | [optional] 
**Eps** | Pointer to **float32** |  | [optional] 
**Epsdiluted** | Pointer to **float32** |  | [optional] 
**FillingDate** | Pointer to **string** |  | [optional] 
**FinalLink** | Pointer to **string** |  | [optional] 
**GeneralAndAdministrativeExpenses** | Pointer to **float32** |  | [optional] 
**GrossProfit** | Pointer to **float32** |  | [optional] 
**GrossProfitRatio** | Pointer to **float32** |  | [optional] 
**IncomeBeforeTax** | Pointer to **float32** |  | [optional] 
**IncomeBeforeTaxRatio** | Pointer to **float32** |  | [optional] 
**IncomeTaxExpense** | Pointer to **float32** |  | [optional] 
**InterestExpense** | Pointer to **float32** |  | [optional] 
**InterestIncome** | Pointer to **float32** |  | [optional] 
**Interestincome** | Pointer to **float32** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**NetIncome** | Pointer to **float32** |  | [optional] 
**NetIncomeRatio** | Pointer to **float32** |  | [optional] 
**OperatingExpenses** | Pointer to **float32** |  | [optional] 
**OperatingIncome** | Pointer to **float32** |  | [optional] 
**OperatingIncomeRatio** | Pointer to **float32** |  | [optional] 
**OtherExpenses** | Pointer to **float32** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**ReportedCurrency** | Pointer to **string** |  | [optional] 
**ResearchAndDevelopmentExpenses** | Pointer to **float32** |  | [optional] 
**Revenue** | Pointer to **float32** |  | [optional] 
**SellingAndMarketingExpenses** | Pointer to **float32** |  | [optional] 
**SellingGeneralAndAdministrativeExpenses** | Pointer to **float32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TotalOtherIncomeExpensesNet** | Pointer to **float32** |  | [optional] 
**WeightedAverageShsOut** | Pointer to **float32** |  | [optional] 
**WeightedAverageShsOutDil** | Pointer to **float32** |  | [optional] 

## Methods

### NewIncomeStatement

`func NewIncomeStatement() *IncomeStatement`

NewIncomeStatement instantiates a new IncomeStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeStatementWithDefaults

`func NewIncomeStatementWithDefaults() *IncomeStatement`

NewIncomeStatementWithDefaults instantiates a new IncomeStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedDate

`func (o *IncomeStatement) GetAcceptedDate() string`

GetAcceptedDate returns the AcceptedDate field if non-nil, zero value otherwise.

### GetAcceptedDateOk

`func (o *IncomeStatement) GetAcceptedDateOk() (*string, bool)`

GetAcceptedDateOk returns a tuple with the AcceptedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedDate

`func (o *IncomeStatement) SetAcceptedDate(v string)`

SetAcceptedDate sets AcceptedDate field to given value.

### HasAcceptedDate

`func (o *IncomeStatement) HasAcceptedDate() bool`

HasAcceptedDate returns a boolean if a field has been set.

### GetCalendarYear

`func (o *IncomeStatement) GetCalendarYear() string`

GetCalendarYear returns the CalendarYear field if non-nil, zero value otherwise.

### GetCalendarYearOk

`func (o *IncomeStatement) GetCalendarYearOk() (*string, bool)`

GetCalendarYearOk returns a tuple with the CalendarYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarYear

`func (o *IncomeStatement) SetCalendarYear(v string)`

SetCalendarYear sets CalendarYear field to given value.

### HasCalendarYear

`func (o *IncomeStatement) HasCalendarYear() bool`

HasCalendarYear returns a boolean if a field has been set.

### GetCik

`func (o *IncomeStatement) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *IncomeStatement) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *IncomeStatement) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *IncomeStatement) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetCostAndExpenses

`func (o *IncomeStatement) GetCostAndExpenses() float32`

GetCostAndExpenses returns the CostAndExpenses field if non-nil, zero value otherwise.

### GetCostAndExpensesOk

`func (o *IncomeStatement) GetCostAndExpensesOk() (*float32, bool)`

GetCostAndExpensesOk returns a tuple with the CostAndExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAndExpenses

`func (o *IncomeStatement) SetCostAndExpenses(v float32)`

SetCostAndExpenses sets CostAndExpenses field to given value.

### HasCostAndExpenses

`func (o *IncomeStatement) HasCostAndExpenses() bool`

HasCostAndExpenses returns a boolean if a field has been set.

### GetCostOfRevenue

`func (o *IncomeStatement) GetCostOfRevenue() float32`

GetCostOfRevenue returns the CostOfRevenue field if non-nil, zero value otherwise.

### GetCostOfRevenueOk

`func (o *IncomeStatement) GetCostOfRevenueOk() (*float32, bool)`

GetCostOfRevenueOk returns a tuple with the CostOfRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOfRevenue

`func (o *IncomeStatement) SetCostOfRevenue(v float32)`

SetCostOfRevenue sets CostOfRevenue field to given value.

### HasCostOfRevenue

`func (o *IncomeStatement) HasCostOfRevenue() bool`

HasCostOfRevenue returns a boolean if a field has been set.

### GetDate

`func (o *IncomeStatement) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *IncomeStatement) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *IncomeStatement) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *IncomeStatement) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDepreciationAndAmortization

`func (o *IncomeStatement) GetDepreciationAndAmortization() float32`

GetDepreciationAndAmortization returns the DepreciationAndAmortization field if non-nil, zero value otherwise.

### GetDepreciationAndAmortizationOk

`func (o *IncomeStatement) GetDepreciationAndAmortizationOk() (*float32, bool)`

GetDepreciationAndAmortizationOk returns a tuple with the DepreciationAndAmortization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationAndAmortization

`func (o *IncomeStatement) SetDepreciationAndAmortization(v float32)`

SetDepreciationAndAmortization sets DepreciationAndAmortization field to given value.

### HasDepreciationAndAmortization

`func (o *IncomeStatement) HasDepreciationAndAmortization() bool`

HasDepreciationAndAmortization returns a boolean if a field has been set.

### GetEbitda

`func (o *IncomeStatement) GetEbitda() float32`

GetEbitda returns the Ebitda field if non-nil, zero value otherwise.

### GetEbitdaOk

`func (o *IncomeStatement) GetEbitdaOk() (*float32, bool)`

GetEbitdaOk returns a tuple with the Ebitda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbitda

`func (o *IncomeStatement) SetEbitda(v float32)`

SetEbitda sets Ebitda field to given value.

### HasEbitda

`func (o *IncomeStatement) HasEbitda() bool`

HasEbitda returns a boolean if a field has been set.

### GetEbitdaratio

`func (o *IncomeStatement) GetEbitdaratio() float32`

GetEbitdaratio returns the Ebitdaratio field if non-nil, zero value otherwise.

### GetEbitdaratioOk

`func (o *IncomeStatement) GetEbitdaratioOk() (*float32, bool)`

GetEbitdaratioOk returns a tuple with the Ebitdaratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbitdaratio

`func (o *IncomeStatement) SetEbitdaratio(v float32)`

SetEbitdaratio sets Ebitdaratio field to given value.

### HasEbitdaratio

`func (o *IncomeStatement) HasEbitdaratio() bool`

HasEbitdaratio returns a boolean if a field has been set.

### GetEps

`func (o *IncomeStatement) GetEps() float32`

GetEps returns the Eps field if non-nil, zero value otherwise.

### GetEpsOk

`func (o *IncomeStatement) GetEpsOk() (*float32, bool)`

GetEpsOk returns a tuple with the Eps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEps

`func (o *IncomeStatement) SetEps(v float32)`

SetEps sets Eps field to given value.

### HasEps

`func (o *IncomeStatement) HasEps() bool`

HasEps returns a boolean if a field has been set.

### GetEpsdiluted

`func (o *IncomeStatement) GetEpsdiluted() float32`

GetEpsdiluted returns the Epsdiluted field if non-nil, zero value otherwise.

### GetEpsdilutedOk

`func (o *IncomeStatement) GetEpsdilutedOk() (*float32, bool)`

GetEpsdilutedOk returns a tuple with the Epsdiluted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsdiluted

`func (o *IncomeStatement) SetEpsdiluted(v float32)`

SetEpsdiluted sets Epsdiluted field to given value.

### HasEpsdiluted

`func (o *IncomeStatement) HasEpsdiluted() bool`

HasEpsdiluted returns a boolean if a field has been set.

### GetFillingDate

`func (o *IncomeStatement) GetFillingDate() string`

GetFillingDate returns the FillingDate field if non-nil, zero value otherwise.

### GetFillingDateOk

`func (o *IncomeStatement) GetFillingDateOk() (*string, bool)`

GetFillingDateOk returns a tuple with the FillingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillingDate

`func (o *IncomeStatement) SetFillingDate(v string)`

SetFillingDate sets FillingDate field to given value.

### HasFillingDate

`func (o *IncomeStatement) HasFillingDate() bool`

HasFillingDate returns a boolean if a field has been set.

### GetFinalLink

`func (o *IncomeStatement) GetFinalLink() string`

GetFinalLink returns the FinalLink field if non-nil, zero value otherwise.

### GetFinalLinkOk

`func (o *IncomeStatement) GetFinalLinkOk() (*string, bool)`

GetFinalLinkOk returns a tuple with the FinalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalLink

`func (o *IncomeStatement) SetFinalLink(v string)`

SetFinalLink sets FinalLink field to given value.

### HasFinalLink

`func (o *IncomeStatement) HasFinalLink() bool`

HasFinalLink returns a boolean if a field has been set.

### GetGeneralAndAdministrativeExpenses

`func (o *IncomeStatement) GetGeneralAndAdministrativeExpenses() float32`

GetGeneralAndAdministrativeExpenses returns the GeneralAndAdministrativeExpenses field if non-nil, zero value otherwise.

### GetGeneralAndAdministrativeExpensesOk

`func (o *IncomeStatement) GetGeneralAndAdministrativeExpensesOk() (*float32, bool)`

GetGeneralAndAdministrativeExpensesOk returns a tuple with the GeneralAndAdministrativeExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralAndAdministrativeExpenses

`func (o *IncomeStatement) SetGeneralAndAdministrativeExpenses(v float32)`

SetGeneralAndAdministrativeExpenses sets GeneralAndAdministrativeExpenses field to given value.

### HasGeneralAndAdministrativeExpenses

`func (o *IncomeStatement) HasGeneralAndAdministrativeExpenses() bool`

HasGeneralAndAdministrativeExpenses returns a boolean if a field has been set.

### GetGrossProfit

`func (o *IncomeStatement) GetGrossProfit() float32`

GetGrossProfit returns the GrossProfit field if non-nil, zero value otherwise.

### GetGrossProfitOk

`func (o *IncomeStatement) GetGrossProfitOk() (*float32, bool)`

GetGrossProfitOk returns a tuple with the GrossProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossProfit

`func (o *IncomeStatement) SetGrossProfit(v float32)`

SetGrossProfit sets GrossProfit field to given value.

### HasGrossProfit

`func (o *IncomeStatement) HasGrossProfit() bool`

HasGrossProfit returns a boolean if a field has been set.

### GetGrossProfitRatio

`func (o *IncomeStatement) GetGrossProfitRatio() float32`

GetGrossProfitRatio returns the GrossProfitRatio field if non-nil, zero value otherwise.

### GetGrossProfitRatioOk

`func (o *IncomeStatement) GetGrossProfitRatioOk() (*float32, bool)`

GetGrossProfitRatioOk returns a tuple with the GrossProfitRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossProfitRatio

`func (o *IncomeStatement) SetGrossProfitRatio(v float32)`

SetGrossProfitRatio sets GrossProfitRatio field to given value.

### HasGrossProfitRatio

`func (o *IncomeStatement) HasGrossProfitRatio() bool`

HasGrossProfitRatio returns a boolean if a field has been set.

### GetIncomeBeforeTax

`func (o *IncomeStatement) GetIncomeBeforeTax() float32`

GetIncomeBeforeTax returns the IncomeBeforeTax field if non-nil, zero value otherwise.

### GetIncomeBeforeTaxOk

`func (o *IncomeStatement) GetIncomeBeforeTaxOk() (*float32, bool)`

GetIncomeBeforeTaxOk returns a tuple with the IncomeBeforeTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeBeforeTax

`func (o *IncomeStatement) SetIncomeBeforeTax(v float32)`

SetIncomeBeforeTax sets IncomeBeforeTax field to given value.

### HasIncomeBeforeTax

`func (o *IncomeStatement) HasIncomeBeforeTax() bool`

HasIncomeBeforeTax returns a boolean if a field has been set.

### GetIncomeBeforeTaxRatio

`func (o *IncomeStatement) GetIncomeBeforeTaxRatio() float32`

GetIncomeBeforeTaxRatio returns the IncomeBeforeTaxRatio field if non-nil, zero value otherwise.

### GetIncomeBeforeTaxRatioOk

`func (o *IncomeStatement) GetIncomeBeforeTaxRatioOk() (*float32, bool)`

GetIncomeBeforeTaxRatioOk returns a tuple with the IncomeBeforeTaxRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeBeforeTaxRatio

`func (o *IncomeStatement) SetIncomeBeforeTaxRatio(v float32)`

SetIncomeBeforeTaxRatio sets IncomeBeforeTaxRatio field to given value.

### HasIncomeBeforeTaxRatio

`func (o *IncomeStatement) HasIncomeBeforeTaxRatio() bool`

HasIncomeBeforeTaxRatio returns a boolean if a field has been set.

### GetIncomeTaxExpense

`func (o *IncomeStatement) GetIncomeTaxExpense() float32`

GetIncomeTaxExpense returns the IncomeTaxExpense field if non-nil, zero value otherwise.

### GetIncomeTaxExpenseOk

`func (o *IncomeStatement) GetIncomeTaxExpenseOk() (*float32, bool)`

GetIncomeTaxExpenseOk returns a tuple with the IncomeTaxExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeTaxExpense

`func (o *IncomeStatement) SetIncomeTaxExpense(v float32)`

SetIncomeTaxExpense sets IncomeTaxExpense field to given value.

### HasIncomeTaxExpense

`func (o *IncomeStatement) HasIncomeTaxExpense() bool`

HasIncomeTaxExpense returns a boolean if a field has been set.

### GetInterestExpense

`func (o *IncomeStatement) GetInterestExpense() float32`

GetInterestExpense returns the InterestExpense field if non-nil, zero value otherwise.

### GetInterestExpenseOk

`func (o *IncomeStatement) GetInterestExpenseOk() (*float32, bool)`

GetInterestExpenseOk returns a tuple with the InterestExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestExpense

`func (o *IncomeStatement) SetInterestExpense(v float32)`

SetInterestExpense sets InterestExpense field to given value.

### HasInterestExpense

`func (o *IncomeStatement) HasInterestExpense() bool`

HasInterestExpense returns a boolean if a field has been set.

### GetInterestIncome

`func (o *IncomeStatement) GetInterestIncome() float32`

GetInterestIncome returns the InterestIncome field if non-nil, zero value otherwise.

### GetInterestIncomeOk

`func (o *IncomeStatement) GetInterestIncomeOk() (*float32, bool)`

GetInterestIncomeOk returns a tuple with the InterestIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestIncome

`func (o *IncomeStatement) SetInterestIncome(v float32)`

SetInterestIncome sets InterestIncome field to given value.

### HasInterestIncome

`func (o *IncomeStatement) HasInterestIncome() bool`

HasInterestIncome returns a boolean if a field has been set.

### GetInterestincome

`func (o *IncomeStatement) GetInterestincome() float32`

GetInterestincome returns the Interestincome field if non-nil, zero value otherwise.

### GetInterestincomeOk

`func (o *IncomeStatement) GetInterestincomeOk() (*float32, bool)`

GetInterestincomeOk returns a tuple with the Interestincome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestincome

`func (o *IncomeStatement) SetInterestincome(v float32)`

SetInterestincome sets Interestincome field to given value.

### HasInterestincome

`func (o *IncomeStatement) HasInterestincome() bool`

HasInterestincome returns a boolean if a field has been set.

### GetLink

`func (o *IncomeStatement) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *IncomeStatement) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *IncomeStatement) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *IncomeStatement) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetNetIncome

`func (o *IncomeStatement) GetNetIncome() float32`

GetNetIncome returns the NetIncome field if non-nil, zero value otherwise.

### GetNetIncomeOk

`func (o *IncomeStatement) GetNetIncomeOk() (*float32, bool)`

GetNetIncomeOk returns a tuple with the NetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncome

`func (o *IncomeStatement) SetNetIncome(v float32)`

SetNetIncome sets NetIncome field to given value.

### HasNetIncome

`func (o *IncomeStatement) HasNetIncome() bool`

HasNetIncome returns a boolean if a field has been set.

### GetNetIncomeRatio

`func (o *IncomeStatement) GetNetIncomeRatio() float32`

GetNetIncomeRatio returns the NetIncomeRatio field if non-nil, zero value otherwise.

### GetNetIncomeRatioOk

`func (o *IncomeStatement) GetNetIncomeRatioOk() (*float32, bool)`

GetNetIncomeRatioOk returns a tuple with the NetIncomeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncomeRatio

`func (o *IncomeStatement) SetNetIncomeRatio(v float32)`

SetNetIncomeRatio sets NetIncomeRatio field to given value.

### HasNetIncomeRatio

`func (o *IncomeStatement) HasNetIncomeRatio() bool`

HasNetIncomeRatio returns a boolean if a field has been set.

### GetOperatingExpenses

`func (o *IncomeStatement) GetOperatingExpenses() float32`

GetOperatingExpenses returns the OperatingExpenses field if non-nil, zero value otherwise.

### GetOperatingExpensesOk

`func (o *IncomeStatement) GetOperatingExpensesOk() (*float32, bool)`

GetOperatingExpensesOk returns a tuple with the OperatingExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingExpenses

`func (o *IncomeStatement) SetOperatingExpenses(v float32)`

SetOperatingExpenses sets OperatingExpenses field to given value.

### HasOperatingExpenses

`func (o *IncomeStatement) HasOperatingExpenses() bool`

HasOperatingExpenses returns a boolean if a field has been set.

### GetOperatingIncome

`func (o *IncomeStatement) GetOperatingIncome() float32`

GetOperatingIncome returns the OperatingIncome field if non-nil, zero value otherwise.

### GetOperatingIncomeOk

`func (o *IncomeStatement) GetOperatingIncomeOk() (*float32, bool)`

GetOperatingIncomeOk returns a tuple with the OperatingIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingIncome

`func (o *IncomeStatement) SetOperatingIncome(v float32)`

SetOperatingIncome sets OperatingIncome field to given value.

### HasOperatingIncome

`func (o *IncomeStatement) HasOperatingIncome() bool`

HasOperatingIncome returns a boolean if a field has been set.

### GetOperatingIncomeRatio

`func (o *IncomeStatement) GetOperatingIncomeRatio() float32`

GetOperatingIncomeRatio returns the OperatingIncomeRatio field if non-nil, zero value otherwise.

### GetOperatingIncomeRatioOk

`func (o *IncomeStatement) GetOperatingIncomeRatioOk() (*float32, bool)`

GetOperatingIncomeRatioOk returns a tuple with the OperatingIncomeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingIncomeRatio

`func (o *IncomeStatement) SetOperatingIncomeRatio(v float32)`

SetOperatingIncomeRatio sets OperatingIncomeRatio field to given value.

### HasOperatingIncomeRatio

`func (o *IncomeStatement) HasOperatingIncomeRatio() bool`

HasOperatingIncomeRatio returns a boolean if a field has been set.

### GetOtherExpenses

`func (o *IncomeStatement) GetOtherExpenses() float32`

GetOtherExpenses returns the OtherExpenses field if non-nil, zero value otherwise.

### GetOtherExpensesOk

`func (o *IncomeStatement) GetOtherExpensesOk() (*float32, bool)`

GetOtherExpensesOk returns a tuple with the OtherExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherExpenses

`func (o *IncomeStatement) SetOtherExpenses(v float32)`

SetOtherExpenses sets OtherExpenses field to given value.

### HasOtherExpenses

`func (o *IncomeStatement) HasOtherExpenses() bool`

HasOtherExpenses returns a boolean if a field has been set.

### GetPeriod

`func (o *IncomeStatement) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *IncomeStatement) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *IncomeStatement) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *IncomeStatement) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetReportedCurrency

`func (o *IncomeStatement) GetReportedCurrency() string`

GetReportedCurrency returns the ReportedCurrency field if non-nil, zero value otherwise.

### GetReportedCurrencyOk

`func (o *IncomeStatement) GetReportedCurrencyOk() (*string, bool)`

GetReportedCurrencyOk returns a tuple with the ReportedCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedCurrency

`func (o *IncomeStatement) SetReportedCurrency(v string)`

SetReportedCurrency sets ReportedCurrency field to given value.

### HasReportedCurrency

`func (o *IncomeStatement) HasReportedCurrency() bool`

HasReportedCurrency returns a boolean if a field has been set.

### GetResearchAndDevelopmentExpenses

`func (o *IncomeStatement) GetResearchAndDevelopmentExpenses() float32`

GetResearchAndDevelopmentExpenses returns the ResearchAndDevelopmentExpenses field if non-nil, zero value otherwise.

### GetResearchAndDevelopmentExpensesOk

`func (o *IncomeStatement) GetResearchAndDevelopmentExpensesOk() (*float32, bool)`

GetResearchAndDevelopmentExpensesOk returns a tuple with the ResearchAndDevelopmentExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearchAndDevelopmentExpenses

`func (o *IncomeStatement) SetResearchAndDevelopmentExpenses(v float32)`

SetResearchAndDevelopmentExpenses sets ResearchAndDevelopmentExpenses field to given value.

### HasResearchAndDevelopmentExpenses

`func (o *IncomeStatement) HasResearchAndDevelopmentExpenses() bool`

HasResearchAndDevelopmentExpenses returns a boolean if a field has been set.

### GetRevenue

`func (o *IncomeStatement) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *IncomeStatement) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *IncomeStatement) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *IncomeStatement) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetSellingAndMarketingExpenses

`func (o *IncomeStatement) GetSellingAndMarketingExpenses() float32`

GetSellingAndMarketingExpenses returns the SellingAndMarketingExpenses field if non-nil, zero value otherwise.

### GetSellingAndMarketingExpensesOk

`func (o *IncomeStatement) GetSellingAndMarketingExpensesOk() (*float32, bool)`

GetSellingAndMarketingExpensesOk returns a tuple with the SellingAndMarketingExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingAndMarketingExpenses

`func (o *IncomeStatement) SetSellingAndMarketingExpenses(v float32)`

SetSellingAndMarketingExpenses sets SellingAndMarketingExpenses field to given value.

### HasSellingAndMarketingExpenses

`func (o *IncomeStatement) HasSellingAndMarketingExpenses() bool`

HasSellingAndMarketingExpenses returns a boolean if a field has been set.

### GetSellingGeneralAndAdministrativeExpenses

`func (o *IncomeStatement) GetSellingGeneralAndAdministrativeExpenses() float32`

GetSellingGeneralAndAdministrativeExpenses returns the SellingGeneralAndAdministrativeExpenses field if non-nil, zero value otherwise.

### GetSellingGeneralAndAdministrativeExpensesOk

`func (o *IncomeStatement) GetSellingGeneralAndAdministrativeExpensesOk() (*float32, bool)`

GetSellingGeneralAndAdministrativeExpensesOk returns a tuple with the SellingGeneralAndAdministrativeExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingGeneralAndAdministrativeExpenses

`func (o *IncomeStatement) SetSellingGeneralAndAdministrativeExpenses(v float32)`

SetSellingGeneralAndAdministrativeExpenses sets SellingGeneralAndAdministrativeExpenses field to given value.

### HasSellingGeneralAndAdministrativeExpenses

`func (o *IncomeStatement) HasSellingGeneralAndAdministrativeExpenses() bool`

HasSellingGeneralAndAdministrativeExpenses returns a boolean if a field has been set.

### GetSymbol

`func (o *IncomeStatement) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *IncomeStatement) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *IncomeStatement) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *IncomeStatement) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTotalOtherIncomeExpensesNet

`func (o *IncomeStatement) GetTotalOtherIncomeExpensesNet() float32`

GetTotalOtherIncomeExpensesNet returns the TotalOtherIncomeExpensesNet field if non-nil, zero value otherwise.

### GetTotalOtherIncomeExpensesNetOk

`func (o *IncomeStatement) GetTotalOtherIncomeExpensesNetOk() (*float32, bool)`

GetTotalOtherIncomeExpensesNetOk returns a tuple with the TotalOtherIncomeExpensesNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOtherIncomeExpensesNet

`func (o *IncomeStatement) SetTotalOtherIncomeExpensesNet(v float32)`

SetTotalOtherIncomeExpensesNet sets TotalOtherIncomeExpensesNet field to given value.

### HasTotalOtherIncomeExpensesNet

`func (o *IncomeStatement) HasTotalOtherIncomeExpensesNet() bool`

HasTotalOtherIncomeExpensesNet returns a boolean if a field has been set.

### GetWeightedAverageShsOut

`func (o *IncomeStatement) GetWeightedAverageShsOut() float32`

GetWeightedAverageShsOut returns the WeightedAverageShsOut field if non-nil, zero value otherwise.

### GetWeightedAverageShsOutOk

`func (o *IncomeStatement) GetWeightedAverageShsOutOk() (*float32, bool)`

GetWeightedAverageShsOutOk returns a tuple with the WeightedAverageShsOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAverageShsOut

`func (o *IncomeStatement) SetWeightedAverageShsOut(v float32)`

SetWeightedAverageShsOut sets WeightedAverageShsOut field to given value.

### HasWeightedAverageShsOut

`func (o *IncomeStatement) HasWeightedAverageShsOut() bool`

HasWeightedAverageShsOut returns a boolean if a field has been set.

### GetWeightedAverageShsOutDil

`func (o *IncomeStatement) GetWeightedAverageShsOutDil() float32`

GetWeightedAverageShsOutDil returns the WeightedAverageShsOutDil field if non-nil, zero value otherwise.

### GetWeightedAverageShsOutDilOk

`func (o *IncomeStatement) GetWeightedAverageShsOutDilOk() (*float32, bool)`

GetWeightedAverageShsOutDilOk returns a tuple with the WeightedAverageShsOutDil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAverageShsOutDil

`func (o *IncomeStatement) SetWeightedAverageShsOutDil(v float32)`

SetWeightedAverageShsOutDil sets WeightedAverageShsOutDil field to given value.

### HasWeightedAverageShsOutDil

`func (o *IncomeStatement) HasWeightedAverageShsOutDil() bool`

HasWeightedAverageShsOutDil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


