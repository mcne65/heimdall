package types

import (
	"bytes"
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v2"
)

const (
	// AddrLen defines a valid address length
	AddrLen = 20
	// MainPrefix defines the  prefix main prefix
	MainPrefix = "matic"

	// PrefixAccount is the prefix for account keys
	PrefixAccount = "acc"
	// PrefixValidator is the prefix for validator keys
	PrefixValidator = "val"
	// PrefixConsensus is the prefix for consensus keys
	PrefixConsensus = "cons"
	// PrefixPublic is the prefix for public keys
	PrefixPublic = "pub"
	// PrefixOperator is the prefix for operator keys
	PrefixOperator = "oper"

	// PrefixAddress is the prefix for addresses
	PrefixAddress = "addr"

	// PrefixAccAddr defines the  prefix of an account's address
	PrefixAccAddr = MainPrefix
	// PrefixAccPub defines the  prefix of an account's public key
	PrefixAccPub = MainPrefix + PrefixPublic
	// PrefixValAddr defines the  prefix of a validator's operator address
	PrefixValAddr = MainPrefix + PrefixValidator + PrefixOperator
	// PrefixValPub defines the  prefix of a validator's operator public key
	PrefixValPub = MainPrefix + PrefixValidator + PrefixOperator + PrefixPublic
	// PrefixConsAddr defines the  prefix of a consensus node address
	PrefixConsAddr = MainPrefix + PrefixValidator + PrefixConsensus
	// PrefixConsPub defines the  prefix of a consensus node public key
	PrefixConsPub = MainPrefix + PrefixValidator + PrefixConsensus + PrefixPublic
)

// Ensure that different address types implement the interface
var _ sdk.Address = HeimdallAddress{}
var _ yaml.Marshaler = HeimdallAddress{}

// HeimdallAddress represents heimdall address
type HeimdallAddress common.Address

// ZeroHeimdallAddress represents zero address
var ZeroHeimdallAddress = HeimdallAddress{}

// EthAddress get eth address
func (aa HeimdallAddress) EthAddress() common.Address {
	return common.Address(aa)
}

// Equals returns boolean for whether two AccAddresses are Equal
func (aa HeimdallAddress) Equals(aa2 sdk.Address) bool {
	if aa.Empty() && aa2.Empty() {
		return true
	}

	return bytes.Equal(aa.Bytes(), aa2.Bytes())
}

// Empty returns boolean for whether an AccAddress is empty
func (aa HeimdallAddress) Empty() bool {
	aa2 := HeimdallAddress{}
	return bytes.Equal(aa.Bytes(), aa2.Bytes())
}

// Marshal returns the raw address bytes. It is needed for protobuf
// compatibility.
func (aa HeimdallAddress) Marshal() ([]byte, error) {
	return aa.Bytes(), nil
}

// Unmarshal sets the address to the given data. It is needed for protobuf
// compatibility.
func (aa *HeimdallAddress) Unmarshal(data []byte) error {
	*aa = HeimdallAddress(common.BytesToAddress(data))
	return nil
}

// MarshalJSON marshals to JSON using Bech32.
func (aa HeimdallAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(aa.String())
}

// MarshalYAML marshals to YAML using Bech32.
func (aa HeimdallAddress) MarshalYAML() (interface{}, error) {
	return aa.String(), nil
}

// UnmarshalJSON unmarshals from JSON assuming Bech32 encoding.
func (aa *HeimdallAddress) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*aa = HexToHeimdallAddress(s)
	return nil
}

// UnmarshalYAML unmarshals from JSON assuming Bech32 encoding.
func (aa *HeimdallAddress) UnmarshalYAML(data []byte) error {
	var s string
	err := yaml.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*aa = HexToHeimdallAddress(s)
	return nil
}

// Bytes returns the raw address bytes.
func (aa HeimdallAddress) Bytes() []byte {
	return aa[:]
}

// String implements the Stringer interface.
func (aa HeimdallAddress) String() string {
	if aa.Empty() {
		return ""
	}

	return common.Address(aa).Hex()
}

// Format implements the fmt.Formatter interface.
// nolint: errcheck
func (aa HeimdallAddress) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(aa.String()))
	case 'p':
		s.Write([]byte(fmt.Sprintf("%p", aa)))
	default:
		s.Write([]byte(fmt.Sprintf("%X", aa.Bytes())))
	}
}

// BytesToHeimdallAddress returns Address with value b.
func BytesToHeimdallAddress(b []byte) HeimdallAddress {
	return HeimdallAddress(common.BytesToAddress(b))
}

// HexToHeimdallAddress returns Address with value b.
func HexToHeimdallAddress(b string) HeimdallAddress {
	return HeimdallAddress(common.HexToAddress(b))
}

// AccAddressToHeimdallAddress returns Address with value b.
func AccAddressToHeimdallAddress(b sdk.AccAddress) HeimdallAddress {
	return BytesToHeimdallAddress(b[:])
}
