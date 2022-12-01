package sdk

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

type ChainAndCoinReq struct {
	AppId     string `json:"app_id"`
	CompanyId uint64 `json:"company_id"`
}

type CoinSymbolData struct {
	CoinName  string `json:"coin_name"`
	CoinType  int64  `json:"coin_type"`
	Support   uint8  `json:"support"`
	ChainType int64  `json:"chain_type"`
	Contract  string `json:"contract"`
	Protocol  string `json:"protocol"`
	Precision int32  `json:"precision"`
	Main      int8   `json:"main"`
}
type CoinSymbolRes struct {
	Code string           `json:"code"`
	Msg  string           `json:"msg"`
	Data []CoinSymbolData `json:"data"`
}

//获取网络和币的配置表
func (c *Client) ChainAndCoin(req *ChainAndCoinReq) ([]CoinSymbolData, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/chain_coin")
	bd, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return nil, err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return nil, err
	}
	//fmt.Println("res:", rs)
	var res CoinSymbolRes
	//res.Data = []CoinSymbolRes{}
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return nil, errors.New(res.Msg)
	}
	return res.Data, nil
}

//转账
type TransferReq struct {
	Order     string `json:"order" binding:"required"`
	Uid       string `json:"uid" binding:"required"`
	ChainType int64  `json:"chain_type"`
	CoinType  int64  `json:"coin_type"`
	Amount    string `json:"amount" binding:"required"`
	ToAddr    string `json:"to_addr" binding:"required"`
	Memo      string `json:"memo"`
	OrderType int8   `json:"order_type"`
}

func (c *Client) Transfer(req *TransferReq) error {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/transfer")
	bd, err := json.Marshal(req)
	if err != nil {
		return err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return err
	}
	var res JSONResponse
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return errors.New(res.Msg)
	}
	return nil
}

//收银台获取地址,请求结果
type PaymentReq struct {
	Uid       string          `json:"uid" binding:"required"`
	ChainType int64           `json:"chain_type"`
	CoinType  int64           `json:"coin_type"`
	Amount    decimal.Decimal `json:"amount" binding:"required"`
	OrderNo   string          `json:"order_no" binding:"required"`
}

//收银台返回结构
type PaymentRes struct {
	OrderNo string `json:"order_no"`
	PlatNo  string `json:"plat_no" `
	PayAddr string `json:"pay_addr"`
}

//收银台获取地址
func (c *Client) PaymentAddr(req *PaymentReq) (*PaymentRes, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/depos")
	bd, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return nil, err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return nil, err
	}
	var res JSONResponse
	res.Data = &PaymentRes{}
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return nil, errors.New(res.Msg)
	}
	return res.Data.(*PaymentRes), nil
}

//收银台订单查询
type CheckPaymentReq struct {
	PlatNo string `json:"plat_no" binding:"required"`
}

type CheckPaymentRes struct {
	OrderNo    string `json:"order_no"`
	PlatNo     string `json:"plat_no" `
	Amount     string `json:"amount"`
	RealAmount string `json:"real_amount"`
	Status     int    `json:"status"`
}

func (c *Client) CheckPayment(req *CheckPaymentReq) (*CheckPaymentRes, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/checkPayment")
	bd, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return nil, err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return nil, err
	}
	//var res CheckPaymentRes
	var res JSONResponse
	res.Data = &CheckPaymentRes{}
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return nil, errors.New(res.Msg)
	}
	return res.Data.(*CheckPaymentRes), nil

}

//获取用户钱包地址
type GetUserAddrReq struct {
	Uid       string `json:"uid" binding:"required"`
	ChainType int64  `json:"chain_type"`
	CoinType  int64  `json:"coin_type"`
}

type GetUserAddrRes struct {
	ChainType int64  `json:"chain"`
	Addr      string `json:"addr"`
	Memo      string `json:"memo"`
}

func (c *Client) GetUserAddr(req *GetUserAddrReq) (*GetUserAddrRes, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/getAddress")
	bd, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return nil, err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return nil, err
	}
	var res JSONResponse
	res.Data = &GetUserAddrRes{}
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return nil, errors.New(res.Msg)
	}
	return res.Data.(*GetUserAddrRes), nil
}

//获取热钱包地址和矿工费地址
type GetHotAddrReq struct {
	ChainType int64 `json:"chain_type"`
}
type GetHotAddrRes struct {
	ChainType int64  `json:"chain"`
	HotAddr   string `json:"hotAddr"`
	FeeAddr   string `json:"feeAddr"`
}

func (c *Client) GetHotAddr(req *GetHotAddrReq) (*GetHotAddrRes, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/getHotAddress")
	bd, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return nil, err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return nil, err
	}
	var res JSONResponse
	res.Data = &GetHotAddrRes{}
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return nil, errors.New(res.Msg)
	}
	return res.Data.(*GetHotAddrRes), nil
}

//获取gas 价格
type GetGasPriceReq struct {
	Uid       string `json:"uid" binding:"required"`
	ChainType int64  `json:"chain_type"`
	CoinType  int64  `json:"coin_type"`
}

func (c *Client) GetGasPrice(req *GetGasPriceReq) (string, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/suggestGasPrice")
	bd, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return "", err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return "", err
	}
	//fmt.Println("rs:", rs)
	var res JSONResponse
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return "", errors.New(res.Msg)
	}
	return res.Data.(string), nil
}

//获取钱包地址余额
type BalanceReq struct {
	Addr  string `json:"addr" binding:"required"`
	Chain int64  `json:"chain_type"`
	Coin  int64  `json:"coin_type"`
}

func (c *Client) GetBalance(req *BalanceReq) (string, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/balance")
	bd, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return "", err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return "", err
	}
	var res JSONResponse
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return "", errors.New(res.Msg)
	}
	return res.Data.(string), nil
}

//更新商户信息
type UpMchtReq struct {
	MerchantName string `json:"merchant_name"`
	CallbackURL  string `json:"callback_url" binding:"required"`
	IpWhites     string `json:"ip_whites"`
	IpWhiteOpen  uint8  `json:"ip_white_open"`
}

func (c *Client) UpdateMerchanntInfor(req *UpMchtReq) error {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/update/merchant")
	bd, err := json.Marshal(req)
	if err != nil {
		return err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return err
	}
	var res JSONResponse
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return errors.New(res.Msg)
	}
	return nil
}

// 获取商户归集余额
type CollectBalanceReq struct {
	MerchantId string `json:"merchant_id"`
	AppId      string `json:"app_id"`
}

type CollectBalanceData struct {
	Addr        string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	TotalAmount string `protobuf:"bytes,2,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
	FrozeAmount string `protobuf:"bytes,3,opt,name=frozeAmount,proto3" json:"frozeAmount,omitempty"`
	ValidAmount string `protobuf:"bytes,4,opt,name=validAmount,proto3" json:"validAmount,omitempty"`
	ChainAmount string `protobuf:"bytes,5,opt,name=chainAmount,proto3" json:"chainAmount,omitempty"`
	CoinType    int32  `protobuf:"varint,6,opt,name=coinType,proto3" json:"coinType,omitempty"`
	ChainType   int32  `protobuf:"varint,7,opt,name=chainType,proto3" json:"chainType,omitempty"`
}
type CollectBalanceRes struct {
	Code string               `json:"code"`
	Msg  string               `json:"msg"`
	Data []CollectBalanceData `json:"data"`
}

func (c *Client) CollectBalance(req *CollectBalanceReq) ([]CollectBalanceData, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/getMerchantCollectBalance")
	bd, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return nil, err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return nil, err
	}
	//fmt.Println("res:", rs)
	var res CollectBalanceRes
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return nil, errors.New(res.Msg)
	}
	return res.Data, nil
}

//获取用户余额
type UserBalanceReq struct {
	MerchantId string `json:"merchant_id"`
	AppId      string `json:"app_id"`
	Uid        string `json:"uid"`
}
type UserBalanceData struct {
	Addr        string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	TotalAmount string `protobuf:"bytes,2,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
	FrozeAmount string `protobuf:"bytes,3,opt,name=frozeAmount,proto3" json:"frozeAmount,omitempty"`
	ValidAmount string `protobuf:"bytes,4,opt,name=validAmount,proto3" json:"validAmount,omitempty"`
	ChainAmount string `protobuf:"bytes,5,opt,name=chainAmount,proto3" json:"chainAmount,omitempty"`
	CoinType    int32  `protobuf:"varint,6,opt,name=coinType,proto3" json:"coinType,omitempty"`
	ChainType   int32  `protobuf:"varint,7,opt,name=chainType,proto3" json:"chainType,omitempty"`
}
type UserBalanceRes struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data []UserBalanceData `json:"data"`
}

func (c *Client) UserBalance(req *UserBalanceReq) ([]UserBalanceData, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/getUserBalance")
	bd, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return nil, err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return nil, err
	}
	var res UserBalanceRes
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return nil, errors.New(res.Msg)
	}
	return res.Data, nil
}

//获取提币订单详情
type WithdrawOrderInfoReq struct {
	OrderNo string `json:"order_no"`
}

type WithdrawOrderInfoRes struct {
	Code      int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"` // 状态吗 200 为成功
	Msg       string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`    //  错误信息
	ChainType int32  `protobuf:"varint,3,opt,name=chainType,proto3" json:"chainType,omitempty"`
	CoinName  string `protobuf:"bytes,4,opt,name=coinName,proto3" json:"coinName,omitempty"`
	Addr      string `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr,omitempty"`
	Amount    string `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee       string `protobuf:"bytes,7,opt,name=fee,proto3" json:"fee,omitempty"`
	Time      string `protobuf:"bytes,8,opt,name=time,proto3" json:"time,omitempty"`
	TradeNo   string `protobuf:"bytes,9,opt,name=tradeNo,proto3" json:"tradeNo,omitempty"`
	TxId      string `protobuf:"bytes,10,opt,name=txId,proto3" json:"txId,omitempty"`
	Extra     string `protobuf:"bytes,11,opt,name=extra,proto3" json:"extra,omitempty"`
	Status    int32  `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
}

func (c *Client) WithdrawOrderInfo(req *WithdrawOrderInfoReq) (*WithdrawOrderInfoRes, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/getWithdrawOrderInfo")
	bd, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return nil, err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return nil, err
	}
	//var res CheckPaymentRes
	var res JSONResponse
	res.Data = &WithdrawOrderInfoRes{}
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return nil, errors.New(res.Msg)
	}
	return res.Data.(*WithdrawOrderInfoRes), nil
}

//获取提币订单列表
type WithdrawOrderListReq struct {
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	PageSize     int    `json:"page_size"`
	PageNum      int    `json:"page_num"`
	TransferType int32  `json:"transfer_type"`
}
type WithdrawOrderListData struct {
	Amount      string `json:"amount"`
	CoinType    int    `json:"coin_type"`
	CreatedAt   string `json:"created_at"`
	OrderStatus int    `json:"order_status"`
	RealAmount  string `json:"real_amount"`
	ToAddr      string `json:"to_addr"`
	TxId        string `json:"tx_id"`
}

type WithdrawOrderListRes struct {
	Code string                  `json:"code"`
	Msg  string                  `json:"msg"`
	Data []WithdrawOrderListData `json:"data"`
}

func (c *Client) WithdrawOrderList(req *WithdrawOrderListReq) ([]WithdrawOrderListData, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/findWithdrawOrderList")
	bd, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return nil, err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return nil, err
	}
	var res WithdrawOrderListRes
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return nil, errors.New(res.Msg)
	}
	return res.Data, nil
}

//获取存币订单列表
type DeposOrderListReq struct {
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	PageSize     int    `json:"page_size"`
	PageNum      int    `json:"page_num"`
	TransferType int32  `json:"transfer_type"`
}
type DeposOrderListData struct {
	Amount      string `json:"amount"`
	CoinType    int    `json:"coin_type"`
	CreatedAt   string `json:"created_at"`
	OrderStatus int    `json:"order_status"`
	RealAmount  string `json:"real_amount"`
	ToAddr      string `json:"to_addr"`
	TxId        string `json:"tx_id"`
}
type DeposOrderListRes struct {
	Code string               `json:"code"`
	Msg  string               `json:"msg"`
	Data []DeposOrderListData `json:"data"`
}

func (c *Client) DeposOrderList(req *DeposOrderListReq) ([]DeposOrderListData, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/findDeposOrderList")
	bd, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 加密数据
	data := map[string]interface{}{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return nil, err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, string(bd), true)
	if err != nil {
		return nil, err
	}
	//var res CheckPaymentRes
	var res DeposOrderListRes
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return nil, errors.New(res.Msg)
	}
	return res.Data, nil
}

// 查询支持的所有链类型和币类型
type CoinTypeRes struct {
	CoinName  string `json:"coin_name"`
	CoinType  int64  `json:"coin_type"`
	Support   int8   `json:"support"`
	ChainType int64  `json:"chain_type"`
	Contract  string `json:"contract"`
	Protocol  string `json:"protocol"`
	Precision int8   `json:"precision"`
	Main      int8   `json:"main"`
}
type ChainTypeRes struct {
	ChainName string `json:"chain_name"`
	ChainType int64  `json:"chain_type"`
}
type GetChainAndCoinListRes struct {
	CoinType  []CoinTypeRes  `json:"coin_type"`
	ChainType []ChainTypeRes `json:"chain_type"`
}

func (c *Client) ChainAndCoinList() (*GetChainAndCoinListRes, error) {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, "finance/getChainAndCoinList")
	// 加密数据
	data := map[string]interface{}{}
	token, err := c.SignHelper(c.AppID, c.Secret, 0, data)
	if err != nil {
		return nil, err
	}
	rs, err := c.DoPost(url, map[string]string{"App": c.AppID, "Access-Token": token}, "", true)
	if err != nil {
		return nil, err
	}
	var res JSONResponse
	res.Data = &GetChainAndCoinListRes{}
	json.Unmarshal([]byte(rs), &res)
	if res.Code != RS_OK {
		return nil, errors.New(res.Msg)
	}
	return res.Data.(*GetChainAndCoinListRes), nil
}
