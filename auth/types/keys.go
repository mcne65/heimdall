package types

import (
	"github.com/maticnetwork/heimdall/types"
)

const (
	// ModuleName is the name of the module
	ModuleName = "auth"

	// StoreKey is the store key string for bor
	StoreKey = ModuleName

	// RouterKey is the message route for bor
	RouterKey = ModuleName

	// QuerierRoute is the querier route for bor
	QuerierRoute = ModuleName

	// DefaultParamspace default name for parameter store
	DefaultParamspace = ModuleName

	// FeeStoreKey is a string representation of the store key for fees
	FeeStoreKey = "fee"

	// FeeCollectorName the root string for the fee collector account address
	FeeCollectorName = "fee_collector"
)

var (
	// AddressStoreKeyPrefix prefix for account-by-address store
	AddressStoreKeyPrefix = []byte{0x01}

	// GlobalAccountNumberKey param key for global account number
	GlobalAccountNumberKey = []byte("globalAccountNumber")
)

// AddressStoreKey turn an address to key used to get it from the account store
func AddressStoreKey(addr types.HeimdallAddress) []byte {
	return append(AddressStoreKeyPrefix, addr.Bytes()...)
}
