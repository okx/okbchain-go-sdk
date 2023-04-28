package exposed

import (
	"github.com/okx/okbchain-go-sdk/module/auth/types"
	gosdktypes "github.com/okx/okbchain-go-sdk/types"
)

// Auth shows the expected behavior for inner auth client
type Auth interface {
	gosdktypes.Module
	AuthQuery
}

// AuthQuery shows the expected query behavior for inner auth client
type AuthQuery interface {
	QueryAccount(accAddrStr string) (types.Account, error)
}
