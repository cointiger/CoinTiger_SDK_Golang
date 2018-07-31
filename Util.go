package CoinTiger_SDK_Golang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const API_PATH = "https://api.cointiger.pro"
const Trading_Macro_v2 = API_PATH + "/exchange/trading/api/v2"
const Trading_Macro = API_PATH + "/exchange/trading/api"
const Market_Macro = API_PATH + "/exchange/trading/api/market"

/*获取系统时间*/
const GET_TIMESTAMP = Trading_Macro_v2 + "/timestampw"

/*获取支持的所有币种*/
const GET_CURRENCYS = Trading_Macro_v2 + "/currencys"

/*获取支持24小时行情*/
const GET_24HOURS = Market_Macro + "/detail"

/*获取深度盘口*/
const GET_DEPTH = Market_Macro + "/depth"

/*获取历史K线*/
const GET_KLine = Market_Macro + "/history/kline"

/*获取成交历史数据*/
const GET_TRADE = Market_Macro + "/history/trade"

/*获取24小时行情 ALL*/
const GET_24HOURS_ALL = "https://www.cointiger.pro/exchange/api/public/market/detail"

/*创建订单*/
const CREATE_ORDER = Trading_Macro_v2 + "/order"

/*撤销订单*/
const CANCEL_ORDER = Trading_Macro_v2 + "/order/batch_cancel"

/*获取当前委托*/
const GET_NOW_ORDER = Trading_Macro_v2 + "/order/orders"

/*获取当前用户订单,成交中和未成交*/
const GET_NOW_USER_ORDER = Trading_Macro + "/order/new"

/*获取当前用户订单,成交和已撤销*/
const GET_USER_HISTORY = Trading_Macro + "/order/history"

/*用户撤单(单个)*/
const DELETE_ORDER = Trading_Macro + "/order"

/*获取时间戳*/
func getTimeStamp() string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)[:13]
}

/*HTTP GET*/
func httpGet(url *url.URL) (string, error) {
	resp, err := http.Get(url.String())
	if err != nil {
		return "", err
	}
	nbytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(nbytes), nil
}

/* HTTP POST*/
func httpPostForm(url *url.URL, values url.Values) (string, error) {
	//http.MethodDelete
	resp, err := http.PostForm(url.String(), values)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/* HTTP Delete*/
func httpDelete(url *url.URL) (string, error) {

	req, err := http.NewRequest(http.MethodDelete, url.String(), nil)
	fmt.Println(url.String())
	if err != nil {
		return "", err
	}

	resp, err := (&http.Client{}).Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*获取服务器时间*/
func GetTimeStamp() (string, error) {
	u, err := url.Parse(GET_TIMESTAMP)
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/*获取支持的币种*/
func GetCurrencys() (string, error) {
	u, err := url.Parse(GET_CURRENCYS)
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/*获取24小时行情*/
func Get24Hours(coin_type string) (string, error) {
	u, err := url.Parse(GET_24HOURS)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	u.RawQuery = q.Encode()
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/*获取24小时行情*/
func Get24HoursAll() (string, error) {
	u, err := url.Parse(GET_24HOURS_ALL)
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/**
*	获取深度
*   coin_type:交易对
*	depth_type:Depth 类型 :"step0", "step1", "step2"
**/
func GetDepth(coin_type, depth_type string) (string, error) {
	u, err := url.Parse(GET_DEPTH)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	q.Set("type", depth_type)
	u.RawQuery = q.Encode()
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/*获取历史K线*/
/*period: K线类型: 1min,5min,15min,30min,60min,1day,1week,1month*/
/*size: 获取数量: [1,2000]*/
func GetKLine(coin_type, period, size string) (string, error) {
	u, err := url.Parse(GET_KLine)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	q.Set("period", period)
	q.Set("size", size)
	u.RawQuery = q.Encode()
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/*获取历史K线*/
/*period: K线类型: 1min,5min,15min,30min,60min,1day,1week,1month*/
func GetKLineEasy(coin_type, period string) (string, error) {
	u, err := url.Parse(GET_KLine)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	q.Set("period", period)
	q.Set("size", "200")
	u.RawQuery = q.Encode()
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/*获取历史成交				*/
/*size: 获取数量: [1,2000]	*/
func GetTrade(coin_type, size string) (string, error) {

	u, err := url.Parse(GET_TRADE)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	q.Set("size", size)
	u.RawQuery = q.Encode()
	if err != nil {
		return "", err
	}
	str, err := httpGet(u)
	return str, err
}

/* 创建订单				*/
/* coin_type:	交易对*/
/* price:		价格*/
/* volume:		数量*/
/* side:		买卖方向 买BUY 卖SELL*/
/* strType:		1 ：限价交易，2：市价交易*/
func CreateOder(coin_type, price, volume, side, strType string) (string, error) {

	strTime := getTimeStamp()
	u, err := url.Parse(CREATE_ORDER)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("time", strTime)
	if err != nil {
		return "", err
	}
	vals := url.Values{
		"symbol": {coin_type},
		"price":  {price},
		"volume": {volume},
		"side":   {side},
		"type":   {strType},
		"time":   {strTime}}
	data := vals.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	strSign := Sign(API_SECRET, data+API_SECRET)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpPostForm(u, vals)
	return str, err
}

type Orders struct {
	m map[string][]string
}

/*添加一个订单*/
/*coin_type: 交易对*/
/*order_num: 订单号*/
func (o *Orders) AddOrder(coin_type, order_num string) {
	if o.m == nil {
		o.m = make(map[string][]string)
		o.m[coin_type] = []string{order_num}
	} else if o.m[coin_type] == nil {
		o.m[coin_type] = []string{order_num}
	} else {
		o.m[coin_type] = append(o.m[coin_type], order_num)
	}
}

/*批量结束订单*/
/*o : 订单结构体*/
func CancelOrders(o Orders) (string, error) {
	strTime := getTimeStamp()
	u, err := url.Parse(CANCEL_ORDER)
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("time", strTime)
	if err != nil {
		return "", err
	}
	bytes, _ := json.Marshal(o.m)
	vals := url.Values{
		"orderIdList": {string(bytes)},
		"time":        {strTime}}
	data := vals.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpPostForm(u, vals)
	return str, err
}

/*查询制定交易对的所有订单*/
/*coin_type 	: 		交易对*/
/*size		 	:		查询数量*/
func GetAllOrder(coin_type, size string) (string, error) {
	u, err := url.Parse(GET_TRADE)
	strTime := getTimeStamp()
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	q.Set("size", size)
	q.Set("time", strTime)
	q.Set("states", "new,canceled,expired,filled,part_filled,pending_cancel,")
	if err != nil {
		return "", err
	}
	vals := url.Values{
		"states": {"new,canceled,expired,filled,part_filled,pending_cancel,"},
		"symbol": {coin_type},
		"size":   {size},
		"time":   {strTime}}
	data := vals.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpGet(u)
	return str, err
}

/*查询制定交易筛选订单*/
/*coin_type 	: 		交易对*/
/*size		 	:		查询数量*/
/*states		:		查询状态*/
func GetOrders(coin_type, size, from, direct string, states, types []string) (string, error) {
	u, err := url.Parse(GET_TRADE)
	strTime := getTimeStamp()
	q := u.Query()
	q.Set("api_key", API_KEY)
	q.Set("symbol", coin_type)
	q.Set("size", size)
	q.Set("from", from)
	q.Set("direct", direct)
	q.Set("time", strTime)
	strStates := ""
	strTypes := ""
	for _, v := range states {
		strStates = strStates + v + ","
	}
	for _, v := range types {
		strTypes = strTypes + v + ","
	}
	q.Set("states", strStates)
	q.Set("types", strTypes)
	if err != nil {
		return "", err
	}
	vals := url.Values{
		"from":   {from},
		"direct": {from},
		"states": {strStates},
		"types":  {strTypes},
		"symbol": {coin_type},
		"size":   {size},
		"time":   {strTime}}
	data := vals.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpGet(u)
	return str, err
}

/*获取当前用户的委托-未成交&成交中*/
func GetApprove(coin_type, offset, limit string) (string, error) {
	u, err := url.Parse(GET_NOW_USER_ORDER)
	strTime := getTimeStamp()
	q := u.Query()
	q.Set("symbol", coin_type)
	q.Set("offset", offset)
	q.Set("limit", limit)
	q.Set("time", strTime)

	data := q.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)

	q.Set("api_key", API_KEY)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpGet(u)
	return str, err
}

/*获取用户的委托-成交&已撤销*/
func GetApproveHistory(coin_type, offset, limit string) (string, error) {
	u, err := url.Parse(GET_USER_HISTORY)
	strTime := getTimeStamp()
	q := u.Query()
	q.Set("symbol", coin_type)
	q.Set("offset", offset)
	q.Set("limit", limit)
	q.Set("time", strTime)

	data := q.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)

	q.Set("api_key", API_KEY)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpGet(u)
	return str, err
}

/*用户撤单-单个订单*/
func DeleteOrder(coin_type, order_id string) (string, error) {
	u, err := url.Parse(DELETE_ORDER)
	strTime := getTimeStamp()
	q := u.Query()
	q.Set("symbol", coin_type)
	q.Set("order_id", order_id)
	q.Set("time", strTime)

	data := q.Encode()
	data = strings.Replace(data, "&", "", -1)
	data = strings.Replace(data, "=", "", -1)
	data = strings.Replace(data, " ", "", -1)
	data, _ = url.QueryUnescape(data)
	strSign := Sign(API_SECRET, data+API_SECRET)

	q.Set("api_key", API_KEY)
	q.Set("sign", strSign)
	u.RawQuery = q.Encode()
	str, err := httpDelete(u)
	return str, err
}
