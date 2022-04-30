package coingeckoapi

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type PriceResponse struct {
	Bitcoin struct {
		Usd float64 `json:"usd"`
	} `json:"bitcoin"`
}

// comma-separated if querying more than 1
// baseID is from coins/list endpoint
// quoteCurrency ex => usd
func (b *Client) SimplePrice(baseID, quoteCurrency string) (*PriceResponse, error) {
	type opt struct {
		Base     string `url:"ids"`
		Currency string `url:"vs_currencies"`
	}
	input := opt{
		Base:     baseID,
		Currency: strings.ToLower(quoteCurrency),
	}
	res, err := b.do("spot", http.MethodGet, "simple/price", input, false, false)
	if err != nil {
		return nil, err
	}
	result := PriceResponse{}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// baseID is from coins/list endpoint
// quoteCurrency ex => usd
func (b *Client) PriceFromData(baseID, quoteCurrency string) (decimal.Decimal, error) {
	type opt struct {
		Localization  string `json:"localization"`
		Tickers       bool   `url:"tickers"`
		MarketData    bool   `json:"market_data"`
		CommunityData bool   `json:"community_data"`
		DeveloperData bool   `json:"developer_data"`
		SparkLine     bool   `json:"sparkline"`
	}
	input := opt{
		Localization:  "false",
		Tickers:       false,
		MarketData:    false,
		CommunityData: false,
		DeveloperData: false,
		SparkLine:     false,
	}
	url := fmt.Sprintf("coins/%s", baseID)
	res, err := b.do("spot", http.MethodGet, url, input, false, false)
	if err != nil {
		return decimal.Zero, err
	}
	result := PriceFromDataResponse{}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return decimal.Zero, err
	}
	var out decimal.Decimal
	switch {
	case strings.EqualFold(quoteCurrency, "usd"):
		out = decimal.NewFromFloat(result.MarketData.CurrentPrice.Usd)
	default:
		return decimal.Zero, errors.New("un-support quoteCurrency")
	}

	return out, nil
}

type PriceFromDataResponse struct {
	ID              string      `json:"id"`
	Symbol          string      `json:"symbol"`
	Name            string      `json:"name"`
	AssetPlatformID interface{} `json:"asset_platform_id"`
	Platforms       struct {
		string `json:""`
	} `json:"platforms"`
	BlockTimeInMinutes float64       `json:"block_time_in_minutes"`
	HashingAlgorithm   interface{}   `json:"hashing_algorithm"`
	Categories         []interface{} `json:"categories"`
	PublicNotice       interface{}   `json:"public_notice"`
	AdditionalNotices  []interface{} `json:"additional_notices"`
	Description        struct {
		En string `json:"en"`
	} `json:"description"`
	Links struct {
		Homepage                    []string    `json:"homepage"`
		BlockchainSite              []string    `json:"blockchain_site"`
		OfficialForumURL            []string    `json:"official_forum_url"`
		ChatURL                     []string    `json:"chat_url"`
		AnnouncementURL             []string    `json:"announcement_url"`
		TwitterScreenName           string      `json:"twitter_screen_name"`
		FacebookUsername            string      `json:"facebook_username"`
		BitcointalkThreadIdentifier interface{} `json:"bitcointalk_thread_identifier"`
		TelegramChannelIdentifier   string      `json:"telegram_channel_identifier"`
		SubredditURL                interface{} `json:"subreddit_url"`
		ReposURL                    struct {
			Github    []string      `json:"github"`
			Bitbucket []interface{} `json:"bitbucket"`
		} `json:"repos_url"`
	} `json:"links"`
	Image struct {
		Thumb string `json:"thumb"`
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"image"`
	CountryOrigin                string      `json:"country_origin"`
	GenesisDate                  interface{} `json:"genesis_date"`
	SentimentVotesUpPercentage   float64     `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64     `json:"sentiment_votes_down_percentage"`
	MarketCapRank                interface{} `json:"market_cap_rank"`
	CoingeckoRank                float64     `json:"coingecko_rank"`
	CoingeckoScore               float64     `json:"coingecko_score"`
	DeveloperScore               float64     `json:"developer_score"`
	CommunityScore               float64     `json:"community_score"`
	LiquidityScore               float64     `json:"liquidity_score"`
	PublicInterestScore          float64     `json:"public_interest_score"`
	MarketData                   struct {
		CurrentPrice struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"current_price"`
		TotalValueLocked interface{} `json:"total_value_locked"`
		McapToTvlRatio   interface{} `json:"mcap_to_tvl_ratio"`
		FdvToTvlRatio    interface{} `json:"fdv_to_tvl_ratio"`
		Roi              interface{} `json:"roi"`
		Ath              struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"ath"`
		AthChangePercentage struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"ath_change_percentage"`
		AthDate struct {
			Aed  time.Time `json:"aed"`
			Ars  time.Time `json:"ars"`
			Aud  time.Time `json:"aud"`
			Bch  time.Time `json:"bch"`
			Bdt  time.Time `json:"bdt"`
			Bhd  time.Time `json:"bhd"`
			Bmd  time.Time `json:"bmd"`
			Bnb  time.Time `json:"bnb"`
			Brl  time.Time `json:"brl"`
			Btc  time.Time `json:"btc"`
			Cad  time.Time `json:"cad"`
			Chf  time.Time `json:"chf"`
			Clp  time.Time `json:"clp"`
			Cny  time.Time `json:"cny"`
			Czk  time.Time `json:"czk"`
			Dkk  time.Time `json:"dkk"`
			Dot  time.Time `json:"dot"`
			Eos  time.Time `json:"eos"`
			Eth  time.Time `json:"eth"`
			Eur  time.Time `json:"eur"`
			Gbp  time.Time `json:"gbp"`
			Hkd  time.Time `json:"hkd"`
			Huf  time.Time `json:"huf"`
			Idr  time.Time `json:"idr"`
			Ils  time.Time `json:"ils"`
			Inr  time.Time `json:"inr"`
			Jpy  time.Time `json:"jpy"`
			Krw  time.Time `json:"krw"`
			Kwd  time.Time `json:"kwd"`
			Lkr  time.Time `json:"lkr"`
			Ltc  time.Time `json:"ltc"`
			Mmk  time.Time `json:"mmk"`
			Mxn  time.Time `json:"mxn"`
			Myr  time.Time `json:"myr"`
			Ngn  time.Time `json:"ngn"`
			Nok  time.Time `json:"nok"`
			Nzd  time.Time `json:"nzd"`
			Php  time.Time `json:"php"`
			Pkr  time.Time `json:"pkr"`
			Pln  time.Time `json:"pln"`
			Rub  time.Time `json:"rub"`
			Sar  time.Time `json:"sar"`
			Sek  time.Time `json:"sek"`
			Sgd  time.Time `json:"sgd"`
			Thb  time.Time `json:"thb"`
			Try  time.Time `json:"try"`
			Twd  time.Time `json:"twd"`
			Uah  time.Time `json:"uah"`
			Usd  time.Time `json:"usd"`
			Vef  time.Time `json:"vef"`
			Vnd  time.Time `json:"vnd"`
			Xag  time.Time `json:"xag"`
			Xau  time.Time `json:"xau"`
			Xdr  time.Time `json:"xdr"`
			Xlm  time.Time `json:"xlm"`
			Xrp  time.Time `json:"xrp"`
			Yfi  time.Time `json:"yfi"`
			Zar  time.Time `json:"zar"`
			Bits time.Time `json:"bits"`
			Link time.Time `json:"link"`
			Sats time.Time `json:"sats"`
		} `json:"ath_date"`
		Atl struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"atl"`
		AtlChangePercentage struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"atl_change_percentage"`
		AtlDate struct {
			Aed  time.Time `json:"aed"`
			Ars  time.Time `json:"ars"`
			Aud  time.Time `json:"aud"`
			Bch  time.Time `json:"bch"`
			Bdt  time.Time `json:"bdt"`
			Bhd  time.Time `json:"bhd"`
			Bmd  time.Time `json:"bmd"`
			Bnb  time.Time `json:"bnb"`
			Brl  time.Time `json:"brl"`
			Btc  time.Time `json:"btc"`
			Cad  time.Time `json:"cad"`
			Chf  time.Time `json:"chf"`
			Clp  time.Time `json:"clp"`
			Cny  time.Time `json:"cny"`
			Czk  time.Time `json:"czk"`
			Dkk  time.Time `json:"dkk"`
			Dot  time.Time `json:"dot"`
			Eos  time.Time `json:"eos"`
			Eth  time.Time `json:"eth"`
			Eur  time.Time `json:"eur"`
			Gbp  time.Time `json:"gbp"`
			Hkd  time.Time `json:"hkd"`
			Huf  time.Time `json:"huf"`
			Idr  time.Time `json:"idr"`
			Ils  time.Time `json:"ils"`
			Inr  time.Time `json:"inr"`
			Jpy  time.Time `json:"jpy"`
			Krw  time.Time `json:"krw"`
			Kwd  time.Time `json:"kwd"`
			Lkr  time.Time `json:"lkr"`
			Ltc  time.Time `json:"ltc"`
			Mmk  time.Time `json:"mmk"`
			Mxn  time.Time `json:"mxn"`
			Myr  time.Time `json:"myr"`
			Ngn  time.Time `json:"ngn"`
			Nok  time.Time `json:"nok"`
			Nzd  time.Time `json:"nzd"`
			Php  time.Time `json:"php"`
			Pkr  time.Time `json:"pkr"`
			Pln  time.Time `json:"pln"`
			Rub  time.Time `json:"rub"`
			Sar  time.Time `json:"sar"`
			Sek  time.Time `json:"sek"`
			Sgd  time.Time `json:"sgd"`
			Thb  time.Time `json:"thb"`
			Try  time.Time `json:"try"`
			Twd  time.Time `json:"twd"`
			Uah  time.Time `json:"uah"`
			Usd  time.Time `json:"usd"`
			Vef  time.Time `json:"vef"`
			Vnd  time.Time `json:"vnd"`
			Xag  time.Time `json:"xag"`
			Xau  time.Time `json:"xau"`
			Xdr  time.Time `json:"xdr"`
			Xlm  time.Time `json:"xlm"`
			Xrp  time.Time `json:"xrp"`
			Yfi  time.Time `json:"yfi"`
			Zar  time.Time `json:"zar"`
			Bits time.Time `json:"bits"`
			Link time.Time `json:"link"`
			Sats time.Time `json:"sats"`
		} `json:"atl_date"`
		MarketCap struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"market_cap"`
		MarketCapRank         interface{} `json:"market_cap_rank"`
		FullyDilutedValuation struct {
		} `json:"fully_diluted_valuation"`
		TotalVolume struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  int64   `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  int64   `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  int64   `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  int64   `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  int64   `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  int64   `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats int64   `json:"sats"`
		} `json:"total_volume"`
		High24H struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"high_24h"`
		Low24H struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"low_24h"`
		PriceChange24H               float64 `json:"price_change_24h"`
		PriceChangePercentage24H     float64 `json:"price_change_percentage_24h"`
		PriceChangePercentage7D      float64 `json:"price_change_percentage_7d"`
		PriceChangePercentage14D     float64 `json:"price_change_percentage_14d"`
		PriceChangePercentage30D     float64 `json:"price_change_percentage_30d"`
		PriceChangePercentage60D     float64 `json:"price_change_percentage_60d"`
		PriceChangePercentage200D    float64 `json:"price_change_percentage_200d"`
		PriceChangePercentage1Y      float64 `json:"price_change_percentage_1y"`
		MarketCapChange24H           float64 `json:"market_cap_change_24h"`
		MarketCapChangePercentage24H float64 `json:"market_cap_change_percentage_24h"`
		PriceChange24HInCurrency     struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"price_change_24h_in_currency"`
		PriceChangePercentage1HInCurrency struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"price_change_percentage_1h_in_currency"`
		PriceChangePercentage24HInCurrency struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"price_change_percentage_24h_in_currency"`
		PriceChangePercentage7DInCurrency struct {
		} `json:"price_change_percentage_7d_in_currency"`
		PriceChangePercentage14DInCurrency struct {
		} `json:"price_change_percentage_14d_in_currency"`
		PriceChangePercentage30DInCurrency struct {
		} `json:"price_change_percentage_30d_in_currency"`
		PriceChangePercentage60DInCurrency struct {
		} `json:"price_change_percentage_60d_in_currency"`
		PriceChangePercentage200DInCurrency struct {
		} `json:"price_change_percentage_200d_in_currency"`
		PriceChangePercentage1YInCurrency struct {
		} `json:"price_change_percentage_1y_in_currency"`
		MarketCapChange24HInCurrency struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"market_cap_change_24h_in_currency"`
		MarketCapChangePercentage24HInCurrency struct {
			Aed  float64 `json:"aed"`
			Ars  float64 `json:"ars"`
			Aud  float64 `json:"aud"`
			Bch  float64 `json:"bch"`
			Bdt  float64 `json:"bdt"`
			Bhd  float64 `json:"bhd"`
			Bmd  float64 `json:"bmd"`
			Bnb  float64 `json:"bnb"`
			Brl  float64 `json:"brl"`
			Btc  float64 `json:"btc"`
			Cad  float64 `json:"cad"`
			Chf  float64 `json:"chf"`
			Clp  float64 `json:"clp"`
			Cny  float64 `json:"cny"`
			Czk  float64 `json:"czk"`
			Dkk  float64 `json:"dkk"`
			Dot  float64 `json:"dot"`
			Eos  float64 `json:"eos"`
			Eth  float64 `json:"eth"`
			Eur  float64 `json:"eur"`
			Gbp  float64 `json:"gbp"`
			Hkd  float64 `json:"hkd"`
			Huf  float64 `json:"huf"`
			Idr  float64 `json:"idr"`
			Ils  float64 `json:"ils"`
			Inr  float64 `json:"inr"`
			Jpy  float64 `json:"jpy"`
			Krw  float64 `json:"krw"`
			Kwd  float64 `json:"kwd"`
			Lkr  float64 `json:"lkr"`
			Ltc  float64 `json:"ltc"`
			Mmk  float64 `json:"mmk"`
			Mxn  float64 `json:"mxn"`
			Myr  float64 `json:"myr"`
			Ngn  float64 `json:"ngn"`
			Nok  float64 `json:"nok"`
			Nzd  float64 `json:"nzd"`
			Php  float64 `json:"php"`
			Pkr  float64 `json:"pkr"`
			Pln  float64 `json:"pln"`
			Rub  float64 `json:"rub"`
			Sar  float64 `json:"sar"`
			Sek  float64 `json:"sek"`
			Sgd  float64 `json:"sgd"`
			Thb  float64 `json:"thb"`
			Try  float64 `json:"try"`
			Twd  float64 `json:"twd"`
			Uah  float64 `json:"uah"`
			Usd  float64 `json:"usd"`
			Vef  float64 `json:"vef"`
			Vnd  float64 `json:"vnd"`
			Xag  float64 `json:"xag"`
			Xau  float64 `json:"xau"`
			Xdr  float64 `json:"xdr"`
			Xlm  float64 `json:"xlm"`
			Xrp  float64 `json:"xrp"`
			Yfi  float64 `json:"yfi"`
			Zar  float64 `json:"zar"`
			Bits float64 `json:"bits"`
			Link float64 `json:"link"`
			Sats float64 `json:"sats"`
		} `json:"market_cap_change_percentage_24h_in_currency"`
		TotalSupply       interface{} `json:"total_supply"`
		MaxSupply         interface{} `json:"max_supply"`
		CirculatingSupply float64     `json:"circulating_supply"`
		LastUpdated       time.Time   `json:"last_updated"`
	} `json:"market_data"`
	PublicInterestStats struct {
		AlexaRank   interface{} `json:"alexa_rank"`
		BingMatches interface{} `json:"bing_matches"`
	} `json:"public_interest_stats"`
	StatusUpdates []interface{} `json:"status_updates"`
	LastUpdated   time.Time     `json:"last_updated"`
}
