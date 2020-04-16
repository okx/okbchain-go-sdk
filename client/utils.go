package client

import (
	"errors"
	"github.com/okex/okchain-go-sdk/common"
)





func checkParamsGetDealsInfo(addr, product, side string, start, end, page, perPage int) (perPageRet int, err error) {
	return checkParamsGetOpenClosedOrders(addr, product, side, start, end, page, perPage)
}

func checkParamsGetTransactionsInfo(addr string, type_, start, end, page, perPage int) (perPageRet int, err error) {
	if !common.IsValidAccaddr(addr) {
		return 0, errors.New("invalid address input")
	}

	if type_ < 0 {
		return 0, errors.New("'type_' cannot be negative")

	}

	perPageRet, err = common.CheckParamsPaging(start, end, page, perPage)
	return
}