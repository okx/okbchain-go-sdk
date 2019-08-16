package okclient

import (
	"fmt"
	"testing"
	"time"
)

const (
	addr   = "okchain1mm43akh88a3qendlmlzjldf8lkeynq68r8l6ts"
	rpcUrl = "localhost:26657"
)

func TestGetAccountInfoByAddr(t *testing.T) {
	okCli := NewClient(rpcUrl)
	acc, err := okCli.GetAccountInfoByAddr(addr)
	assertNotEqual(t, err, nil)
	fmt.Println(acc)
}

func TestGetTokensInfoByAddr(t *testing.T) {
	okCli := NewClient(rpcUrl)
	tokensInfo, err := okCli.GetTokensInfoByAddr(addr)
	assertNotEqual(t, err, nil)
	fmt.Println(tokensInfo)
}

func TestGetTokenInfoByAddr(t *testing.T) {
	okCli := NewClient(rpcUrl)
	tokenInfo, err := okCli.GetTokenInfoByAddr(addr, "okb")
	assertNotEqual(t, err, nil)
	fmt.Println(tokenInfo)
}

func TestGetTokensInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	tokensInfo, err := okCli.GetTokensInfo()
	assertNotEqual(t, err, nil)
	for _, t := range tokensInfo {
		fmt.Println(t)
	}
}

func TestGetTokenInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	tokenInfo, err := okCli.GetTokenInfo("okb")
	assertNotEqual(t, err, nil)
	fmt.Println(tokenInfo)
}

func TestGetProductsInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	productsList, err := okCli.GetProductsInfo()
	assertNotEqual(t, err, nil)
	for _, p := range productsList {
		fmt.Println(p)
	}
}

func TestGetDepthbookInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	depthbook, err := okCli.GetDepthbookInfo("xxb_okb")
	assertNotEqual(t, err, nil)
	for _, ask := range depthbook.Asks {
		fmt.Println(ask)
	}
	for _, bid := range depthbook.Bids {
		fmt.Println(bid)
	}

}

func TestGetCandlesInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	candles, err := okCli.GetCandlesInfo("xxb_okb", 60, 100)
	assertNotEqual(t, err, nil)
	for _, line := range candles {
		fmt.Println(line)
	}
}

func TestGetTickersInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	tickers, err := okCli.GetTickersInfo(10)
	assertNotEqual(t, err, nil)
	for _, ticker := range tickers {
		fmt.Println(ticker)
	}
}

func TestGetRecentTxRecord(t *testing.T) {
	okCli := NewClient(rpcUrl)
	records, err := okCli.GetRecentTxRecord("xxb_okb", 0, int(time.Now().Unix()), 0, 10)
	assertNotEqual(t, err, nil)
	for _, record := range records {
		fmt.Println(record)
	}
}




func assertNotEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("test failed: %s", a)
	}
}

func assertEqual(t *testing.T, a, b interface{}) {
	if a == b {
		t.Errorf("test failed: %s", a)
	}
}
