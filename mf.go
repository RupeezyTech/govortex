package govortex

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// URIMfHoldings         string = "/mf/holdings"
// URIFundDetails        string = "/mf/fund/details"
// URIMfFunds            string = "/mf/fund"
// URIFundMainCategories string = "/mf/fund/mainCategories"
// URIFundCategories     string = "/mf/fund/categories"

func (v *VortexApi) MfHoldings(ctx context.Context) (*MfHoldingsResponse, error) {
	var resp MfHoldingsResponse
	_, err := v.doJson(ctx, "GET", URIMfHoldings, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) FundDetails(ctx context.Context, isin string) (*MfDetailResponse, error) {
	var resp MfDetailResponse

	u := &url.URL{
		Path: URIFundDetails,
	}
	params := url.Values{}
	params.Add("isin", isin)
	u.RawQuery = params.Encode()
	_, err := v.doJson(ctx, "GET", u.String(), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
func (v *VortexApi) FundNavs(ctx context.Context, isin string) (*NAVResponse, error) {
	var resp NAVResponse

	u := &url.URL{
		Path: URIFundNavs,
	}
	params := url.Values{}
	params.Add("isin", isin)
	u.RawQuery = params.Encode()
	_, err := v.doJson(ctx, "GET", u.String(), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) GetPortfolioImpact(ctx context.Context, isin string, amount float64) (*PortfolioImpactResponse, error) {
	var resp PortfolioImpactResponse
	u := &url.URL{
		Path: URIPortfolioImpact,
	}
	params := url.Values{}
	params.Add("isin", isin)
	params.Add("amount", fmt.Sprintf("%f", amount))
	u.RawQuery = params.Encode()
	_, err := v.doJson(ctx, "GET", u.String(), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) GetFundOverlap(ctx context.Context, isin string, amount float64) (*FundOverlapResponse, error) {
	var resp FundOverlapResponse
	u := &url.URL{
		Path: URIPortfolioOverlap,
	}
	params := url.Values{}
	params.Add("isin", isin)
	params.Add("amount", fmt.Sprintf("%f", amount))
	u.RawQuery = params.Encode()
	_, err := v.doJson(ctx, "GET", u.String(), nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) SIPBook(ctx context.Context) (*SipsResponse, error) {
	var resp SipsResponse
	_, err := v.doJson(ctx, "GET", URIMFSips, nil, nil, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (v *VortexApi) MFOrderBook(ctx context.Context) {}

type HoldingConfig struct {
	OverAllXirr             float64    `json:"OverAllXirr"`
	MarketValueAmount       float64    `json:"MarketValueAmount"`
	InvestedAmount          float64    `json:"InvestedAmount"`
	Returns                 float64    `json:"Returns"`
	ReturnsPercentage       float64    `json:"ReturnsPercentage"`
	LastImportedAt          *time.Time `json:"LastImportedAt"`
	FundsImportStatus       string     `json:"FundsImportStatus"`
	OneDayReturns           float64    `json:"OneDayReturns"`
	OneDayReturnsPercentage float64    `json:"OneDayReturnsPercentage"`
	NavAsOn                 time.Time  `json:"NavAsOn"`
}

type MfHoldingsResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Data    MfHoldings `json:"data"`
}

type MfHoldings struct {
	OverallConfig HoldingConfig               `json:"OverallConfig"`
	IsinWiseData  map[string]*HoldingResponse `json:"IsinWiseData"`
}

type HoldingResponse struct {
	Name                    string                `json:"Name"`
	MarketValueAmount       float64               `json:"MarketValueAmount"`
	InvestedAmount          float64               `json:"InvestedAmount"`
	Returns                 float64               `json:"Returns"`
	ReturnsPercentage       float64               `json:"ReturnsPercentage"`
	OneDayReturns           float64               `json:"OneDayReturns"`
	OneDayReturnsPercentage float64               `json:"OneDayReturnsPercentage"`
	Folios                  []HoldingReportScheme `json:"Folios"`
	BseScheme               BseSchemeDetail       `json:"BseScheme"`
	CmotsScheme             SchemeMaster          `json:"CmotsScheme"`
	IsinXirr                float64               `json:"IsinXirr"`
}

type HoldingReportScheme struct {
	Xirr                    float64    `json:"xirr"`
	ClientCode              string     `json:"client_code"`
	DpamId                  int        `json:"dpam_id"`
	FolioNumber             string     `json:"folio_number"`
	FolioXirr               float64    `json:"folio_xirr"`
	Isin                    string     `json:"isin"`
	CoCode                  int        `json:"co_code"`
	Name                    string     `json:"name"`
	Type                    string     `json:"type"`
	Units                   float64    `json:"units"`
	RedeemableUnits         float64    `json:"redeemable_units"`
	MarketValueAmount       float64    `json:"market_value_amount"`
	RedeemableAmount        float64    `json:"redeemable_amount"`
	InvestedAmount          float64    `json:"invested_amount"`
	NavValue                float64    `json:"nav_value"`
	BuyPrice                float64    `json:"buy_price"`
	AsOn                    *time.Time `json:"as_on"`
	CreatedAt               time.Time  `json:"created_at"`
	UpdatedAt               time.Time  `json:"updated_at"`
	HoldType                string     `json:"hold_type"`
	IsDemat                 string     `json:"is_demat"`
	OneDayReturns           float64    `json:"one_day_returns"`
	OneDayReturnsPercentage float64    `json:"one_day_returns_percentage"`
}

type BseSchemeDetail struct {
	UniqueNo                     int       `json:"UniqueNo"`
	SchemeCode                   string    `json:"SchemeCode"`
	AMCSchemeCode                string    `json:"AMCSchemeCode"`
	RTASchemeCode                string    `json:"RTASchemeCode"`
	ISIN                         string    `gorm:"primaryKey" json:"ISIN"`
	AMCCode                      string    `json:"AMCCode"`
	SchemeType                   string    `json:"SchemeType"`
	SchemePlan                   string    `json:"SchemePlan"`
	SchemeName                   string    `json:"SchemeName"`
	PurchaseAllowed              string    `json:"PurchaseAllowed"`
	PurchaseTransactionMode      string    `json:"PurchaseTransactionMode"`
	MinimumPurchaseAmount        float64   `json:"MinimumPurchaseAmount"`
	AdditionalPurchaseAmount     float64   `json:"AdditionalPurchaseAmount"`
	MaximumPurchaseAmount        float64   `json:"MaximumPurchaseAmount"`
	PurchaseAmountMultiplier     float64   `json:"PurchaseAmountMultiplier"`
	PurchaseCutoffTime           string    `json:"PurchaseCutoffTime"`
	RedemptionAllowed            string    `json:"RedemptionAllowed"`
	RedemptionTransactionMode    string    `json:"RedemptionTransactionMode"`
	MinimumRedemptionQty         float64   `json:"MinimumRedemptionQty"`
	RedemptionQtyMultiplier      float64   `json:"RedemptionQtyMultiplier"`
	MaximumRedemptionQty         float64   `json:"MaximumRedemptionQty"`
	RedemptionAmountMinimum      float64   `json:"RedemptionAmountMinimum"`
	RedemptionAmountMaximum      float64   `json:"RedemptionAmountMaximum"`
	RedemptionAmountMultiple     float64   `json:"RedemptionAmountMultiple"`
	RedemptionCutoffTime         string    `json:"RedemptionCutoffTime"`
	RTAAgentCode                 string    `json:"RTAAgentCode"`
	AMCActiveFlag                string    `json:"AMCActiveFlag"`
	DividendReinvestmentFlag     string    `json:"DividendReinvestmentFlag"`
	SIPFlag                      string    `json:"SIPFlag"`
	STPFlag                      string    `json:"STPFlag"`
	SWPFlag                      string    `json:"SWPFlag"`
	SwitchFlag                   string    `json:"SwitchFlag"`
	SettlementType               string    `json:"SettlementType"`
	AMCInd                       string    `json:"AMCInd"`
	FaceValue                    float64   `json:"FaceValue"`
	StartDate                    time.Time `json:"StartDate"`
	EndDate                      time.Time `json:"EndDate"`
	ExitLoadFlag                 string    `json:"ExitLoadFlag"`
	ExitLoad                     float64   `json:"ExitLoad"`
	LockInPeriodFlag             string    `json:"LockInPeriodFlag"`
	LockInPeriod                 int       `json:"LockInPeriod"`
	ChannelPartnerCode           string    `json:"ChannelPartnerCode"`
	ReOpeningDate                time.Time `json:"ReOpeningDate"`
	SIPTransactionMode           string    `json:"SIPTransactionMode"`
	SIPFrequency                 string    `json:"SIPFrequency"`
	SIPDates                     string    `json:"SIPDates"`
	SIPMinimumGap                int       `json:"SIPMinimumGap"`
	SIPMaximumGap                int       `json:"SIPMaximumGap"`
	SIPInstallmentGap            int       `json:"SIPInstallmentGap"`
	SIPStatus                    string    `json:"SIPStatus"`
	SIPMinimumInstallmentAmount  float64   `json:"SIPMinimumInstallmentAmount"`
	SIPMaximumInstallmentAmount  float64   `json:"SIPMaximumInstallmentAmount"`
	SIPMultiplierAmount          float64   `json:"SIPMultiplierAmount"`
	SIPMinimumInstallmentNumbers int       `json:"SIPMinimumInstallmentNumbers"`
	SIPMaximumInstallmentNumbers int       `json:"SIPMaximumInstallmentNumbers"`
	SchemeISIN                   string    `json:"SchemeISIN"`
	PauseFlag                    string    `json:"PauseFlag"`
	PauseMinimumInstallments     int       `json:"PauseMinimumInstallments"`
	PauseMaximumInstallments     int       `json:"PauseMaximumInstallments"`
	PauseModificationCount       int       `json:"PauseModificationCount"`
	L0Flag                       bool      `json:"L0Flag"`
}

type SchemeMaster struct {
	Isin                     string    `json:"isin"`
	Mfschcode                int       `json:"Mfschcode"`
	Mfcocode                 int       `json:"Mfcocode"`
	Amficode                 int       `json:"amficode"`
	CategoryCode             int       `json:"ClassCode"`
	Schname                  string    `json:"sch_name"`
	Navrs                    float64   `json:"navrs"`
	Navdate                  string    `json:"Navdate"`
	RtCode                   string    `json:"rtcode"`
	IsinReinvestment         string    `json:"isin_reinvestment"`
	FundManager              string    `json:"FundManager"`
	LaunchDateString         string    `json:"LaunchDate"`
	LaunchDate               time.Time `json:"launch_date"`
	InceptionDateString      string    `json:"InceptionDate"`
	InceptionDate            time.Time `json:"inception_date"`
	MinInvestment            float64   `json:"MinInvestment"`
	IncrementalInvestment    float64   `json:"IncrementalInvestment"`
	MinInvestmentSIP         float64   `json:"MinInvestment_SIP"`
	Frequency                string    `json:"frequency"`
	SchemeAUM                float64   `json:"SchemeAUM"`
	EntryLoad                string    `json:"EntrytLoad"`
	ExitLoad                 string    `json:"ExitLoad"`
	FundType                 string    `json:"FundType"`
	InvestmentType           string    `json:"InvestmentType"`
	MCAPCategory             string    `json:"MCAPCategory"`
	BMCode                   int       `json:"BM_code"`
	BMCodefloat              float64   `json:"BMCode"`
	BenchmarkName            string    `json:"BenchmarkName"`
	RiskometerValue          string    `json:"riskometervalue"`
	SchemeInvestmentType     string    `json:"SchemeInvestmentType"`
	SchemeType               string    `json:"SchemeType"`
	GroupCode                string    `json:"groupcode"`
	GroupName                string    `json:"groupname"`
	MaturityDate             string    `json:"maturitydate"`
	SubscriptionAvailability string    `json:"subscriptionavailability"`

	// 	Scheme Profile
	MaturityPeriod string  `json:"MaturityPeriod"`
	ExpRatio       float64 `json:"EXPRATIO"`
	TaxBName       string  `json:"taxbname"`
	TAXB           float64 `json:"TAXB"`
	SIPExitLoad    float64 `json:"SIP_Exitload"`

	//  Scheme ratios
	RatiosAsOnDate string  `json:"Date"`
	BETA           float64 `json:"BETA"`
	SD             float64 `json:"SD"`
	TREYNOR        float64 `json:"TREYNOR"`
	ALPHA          float64 `json:"ALPHA"`
	SHARPE         float64 `json:"SHARPE"`
	RSQUARE        float64 `json:"RSQUARE"`
	PTRatio        float64 `json:"PTRatio"`
	PE             float64 `json:"PE"`

	IsAllowed         bool
	SchemeReturns     SchemeReturns
	SchemeCAGR        SchemeReturns
	StdDeviation      SchemeStdDeviation
	CanonicalUrl      string `json:"canonical_url"`
	IsNfo             bool
	SubCategory_Theme string `json:"SubCategory_Theme"`
	Category          string `json:"Category"`
	FundRating        *int   `json:"fund_rating"`
}

type FundOfferScheme struct {
	Mfcocode   float64 `json:"Mfcocode"`
	Mfschcode  float64 `json:"Mfschcode"`
	Name       string  `json:"Name"`
	SchName    string  `json:"sch_name"`
	ISIN       string  `json:"isin"`
	OfferPrice float64 `json:"OFFERPRICE"`
	ClassCode  int     `json:"ClassCode"`
}
type SchemeReturns struct {
	Isin          string
	OneDay        float64
	OneWeek       float64
	OneMonth      float64
	ThreeMonth    float64
	SixMonth      float64
	OneYear       float64
	ThreeYear     float64
	FiveYear      float64
	TenYear       float64
	ThreeYearCagr float64
	Date          time.Time
}

type SchemeStdDeviation struct {
	Isin      string
	Date      time.Time
	OneYear   float64
	ThreeYear float64
	FiveYear  float64
	TenYear   float64
}

type SchemeDetailsResponse struct {
	AssetAllocation   []AssetAllocation   `json:"assetAllocation"`
	SectorAllocation  []SectorAllocation  `json:"sectorAllocation"`
	HoldingAllocation []HoldingAllocation `json:"holdingAllocation"`
	CMOTSDetails      SchemeMaster        `json:"cmotsDetails"`
	BseSchemeDetails  BseSchemeDetail     `json:"bseSchemeDetails"`
}

type MfDetailResponse struct {
	Status string                `json:"status"`
	Data   SchemeDetailsResponse `json:"data"`
}

type AssetAllocation struct {
	Isin                string  `json:"Isin"`
	AssetCodeFloat      float64 `json:"AssetCode"`
	AssetCode           int     `json:"asset_code"`
	AssetName           string  `json:"AssetName"`
	HoldingCurrentMonth float64 `json:"Holding_CurrentMonth"`
	HoldingPrevMonth    float64 `json:"Holding_PrevMonth"`
	CurrentMonth        string  `gorm:"-" json:"CurrentMonth"`
}

type SectorAllocation struct {
	Isin           string    `json:"Isin"`
	Value          float64   `json:"VALUE"`
	PercentHolding float64   `json:"PERC_HOLD"`
	Sector         string    `json:"SECTOR"`
	AsOnDate       time.Time `json:"AsOnDate"`
}

type HoldingAllocation struct {
	Isin          string    `json:"Isin"`
	CompanyCode   int       `json:"co_code"`
	CompanyName   string    `json:"co_name"`
	InvDateString string    `json:"invdate"`
	InvDate       time.Time `json:"InvDate"`
	PercHolding   float64   `json:"Perc_hold"`
	MktValue      float64   `json:"MktValue"`
	TotalShares   float64   `json:"TotalShares"`
	AssetType     string    `json:"AssetType"`
	Rating        string    `json:"rating"`
}

type NAVResponse struct {
	Status string       `json:"status"`
	Data   []TimeSeries `json:"data"`
}
type TimeSeries struct {
	Date  time.Time `json:"date"`
	Value float64   `json:"value"`
}

type PortfolioImpactResponse struct {
	Status string          `json:"status"`
	Data   PortfolioImpact `json:"data"`
}

type PortfolioImpact struct {
	New      PortfolioStats `json:"New"`
	Original PortfolioStats `json:"Original"`
}

type PortfolioStats struct {
	Returns float64 `json:"Returns"`
	Risk    float64 `json:"Risk"`
}

type FundOverlapResponse struct {
	Status string      `json:"status"`
	Data   FundOverlap `json:"data"`
}

type FundOverlap struct {
	Value    float64
	Overlaps []Overlaps
}

type Overlaps struct {
	Isin     string
	Holdings []Allocation
}

type Allocation struct {
	Name  string
	Value float64
}

type SipsResponse struct {
	Status string `json:"status"`
	Data   []Sip  `json:"data"`
}

type Sip struct {
	ID                      int            `json:"ID"`
	CreatedAt               time.Time      `json:"CreatedAt"`
	UpdatedAt               time.Time      `json:"UpdatedAt"`
	DeletedAt               *time.Time     `json:"DeletedAt,omitempty"`
	TransactionCode         string         `json:"transactionCode"`
	Isin                    string         `json:"isin"`
	Amount                  float64        `json:"amount"`
	ClientCode              string         `json:"clientCode"`
	SchemeCode              string         `json:"schemeCode"`
	Remarks                 string         `json:"remarks"`
	StartDate               time.Time      `json:"startDate"`
	StartDay                int            `json:"startDay"`
	ModifiedStartDate       *time.Time     `json:"modifiedStartDate,omitempty"`
	MandateId               string         `json:"mandateId"`
	State                   string         `json:"state"`
	Frequency               string         `json:"frequency"`
	StepUp                  bool           `json:"stepUp"`
	StepUpFrequency         string         `json:"stepUpFrequency"`
	StepUpPercentage        float64        `json:"stepUpPercentage"`
	StepUpAmount            float64        `json:"stepUpAmount"`
	IsAmcSip                bool           `json:"isAmcSip"`
	InitialInvestmentAmount float64        `json:"initialInvestmentAmount"`
	BuySellType             string         `json:"buySellType"`
	NextInstallmentDate     *time.Time     `json:"nextInstallmentDate,omitempty"`
	MfLabId                 int            `json:"mfLabId"`
	PauseDate               *time.Time     `json:"pauseDate,omitempty"`
	PauseNoOfInstallments   int            `json:"pauseNoOfInstallments"`
	SwitchNoOfInstallments  int            `json:"switchNoOfInstallments"`
	InstallmentNumber       int            `json:"installmentNumber"`
	LastStepUpDate          *time.Time     `json:"lastStepUpDate,omitempty"`
	SchemeName              string         `json:"schemeName"`
	CoCode                  int            `json:"coCode"`
	SchemesDetails          SchemesDetails `json:"schemes_details"`
}

type SchemesDetails struct {
	BseSchemeDetails BseSchemeDetail `json:"bseSchemeDetails"`
	CmotsDetails     SchemeMaster    `json:"cmotsDetails"`
	NfoDetails       FundOfferScheme `json:"nfoDetails"`
}
