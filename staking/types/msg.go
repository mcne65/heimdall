package types

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hmCommon "github.com/maticnetwork/heimdall/common"
	"github.com/maticnetwork/heimdall/helper"
	"github.com/maticnetwork/heimdall/types"
	hmTypes "github.com/maticnetwork/heimdall/types"
)

var cdc = codec.New()

//
// Validator Join
//

var _ sdk.Msg = &MsgValidatorJoin{}

type MsgValidatorJoin struct {
	From         hmTypes.HeimdallAddress `json:"from"`
	ID           hmTypes.ValidatorID     `json:"id"`
	SignerPubKey hmTypes.PubKey          `json:"pub_key"`
	TxHash       hmTypes.HeimdallHash    `json:"tx_hash"`
	LogIndex     uint64                  `json:"log_index"`
}

// NewMsgValidatorJoin creates new validator-join
func NewMsgValidatorJoin(
	from hmTypes.HeimdallAddress,
	id uint64,
	pubkey hmTypes.PubKey,
	txhash hmTypes.HeimdallHash,
	logIndex uint64,
) MsgValidatorJoin {

	return MsgValidatorJoin{
		From:         from,
		ID:           hmTypes.NewValidatorID(id),
		SignerPubKey: pubkey,
		TxHash:       txhash,
		LogIndex:     logIndex,
	}
}

func (msg MsgValidatorJoin) Type() string {
	return "validator-join"
}

func (msg MsgValidatorJoin) Route() string {
	return RouterKey
}

func (msg MsgValidatorJoin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{hmTypes.HeimdallAddressToAccAddress(msg.From)}
}

func (msg MsgValidatorJoin) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgValidatorJoin) ValidateBasic() sdk.Error {
	if msg.ID <= 0 {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid validator ID %v", msg.ID)
	}

	if bytes.Equal(msg.SignerPubKey.Bytes(), helper.ZeroPubKey.Bytes()) {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid pub key %v", msg.SignerPubKey.String())
	}

	if msg.From.Empty() {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid proposer %v", msg.From.String())
	}

	return nil
}

// GetTxHash Returns tx hash
func (msg MsgValidatorJoin) GetTxHash() types.HeimdallHash {
	return msg.TxHash
}

// GetLogIndex Returns log index
func (msg MsgValidatorJoin) GetLogIndex() uint64 {
	return msg.LogIndex
}

//
// Stake update
//

//
// validator exit
//

var _ sdk.Msg = &MsgStakeUpdate{}

// MsgStakeUpdate represents stake update
type MsgStakeUpdate struct {
	From     hmTypes.HeimdallAddress `json:"from"`
	ID       hmTypes.ValidatorID     `json:"id"`
	TxHash   hmTypes.HeimdallHash    `json:"tx_hash"`
	LogIndex uint64                  `json:"log_index"`
}

// NewMsgStakeUpdate represents stake update
func NewMsgStakeUpdate(from hmTypes.HeimdallAddress, id uint64, txhash hmTypes.HeimdallHash, logIndex uint64) MsgStakeUpdate {
	return MsgStakeUpdate{
		From:     from,
		ID:       hmTypes.NewValidatorID(id),
		TxHash:   txhash,
		LogIndex: logIndex,
	}
}

func (msg MsgStakeUpdate) Type() string {
	return "validator-stake-update"
}

func (msg MsgStakeUpdate) Route() string {
	return RouterKey
}

func (msg MsgStakeUpdate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{hmTypes.HeimdallAddressToAccAddress(msg.From)}
}

func (msg MsgStakeUpdate) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgStakeUpdate) ValidateBasic() sdk.Error {
	if msg.ID <= 0 {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid validator ID %v", msg.ID)
	}

	if msg.From.Empty() {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid proposer %v", msg.From.String())
	}

	return nil
}

// GetTxHash Returns tx hash
func (msg MsgStakeUpdate) GetTxHash() types.HeimdallHash {
	return msg.TxHash
}

// GetLogIndex Returns log index
func (msg MsgStakeUpdate) GetLogIndex() uint64 {
	return msg.LogIndex
}

//
// validator update
//
var _ sdk.Msg = &MsgSignerUpdate{}

// MsgSignerUpdate signer update struct
// TODO add old signer sig check
type MsgSignerUpdate struct {
	From            hmTypes.HeimdallAddress `json:"from"`
	ID              hmTypes.ValidatorID     `json:"id"`
	NewSignerPubKey hmTypes.PubKey          `json:"pubKey"`
	TxHash          hmTypes.HeimdallHash    `json:"tx_hash"`
	LogIndex        uint64                  `json:"log_index"`
}

func NewMsgSignerUpdate(
	from hmTypes.HeimdallAddress,
	id uint64,
	pubKey hmTypes.PubKey,
	txhash hmTypes.HeimdallHash,
	logIndex uint64,
) MsgSignerUpdate {
	return MsgSignerUpdate{
		From:            from,
		ID:              hmTypes.NewValidatorID(id),
		NewSignerPubKey: pubKey,
		TxHash:          txhash,
		LogIndex:        logIndex,
	}
}

func (msg MsgSignerUpdate) Type() string {
	return "signer-update"
}

func (msg MsgSignerUpdate) Route() string {
	return RouterKey
}

func (msg MsgSignerUpdate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{hmTypes.HeimdallAddressToAccAddress(msg.From)}
}

func (msg MsgSignerUpdate) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgSignerUpdate) ValidateBasic() sdk.Error {
	if msg.ID <= 0 {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid validator ID %v", msg.ID)
	}

	if msg.From.Empty() {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid proposer %v", msg.From.String())
	}

	if bytes.Equal(msg.NewSignerPubKey.Bytes(), helper.ZeroPubKey.Bytes()) {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid pub key %v", msg.NewSignerPubKey.String())
	}

	return nil
}

// GetTxHash Returns tx hash
func (msg MsgSignerUpdate) GetTxHash() types.HeimdallHash {
	return msg.TxHash
}

// GetLogIndex Returns log index
func (msg MsgSignerUpdate) GetLogIndex() uint64 {
	return msg.LogIndex
}

//
// validator exit
//

var _ sdk.Msg = &MsgValidatorExit{}

type MsgValidatorExit struct {
	From     hmTypes.HeimdallAddress `json:"from"`
	ID       hmTypes.ValidatorID     `json:"id"`
	TxHash   hmTypes.HeimdallHash    `json:"tx_hash"`
	LogIndex uint64                  `json:"log_index"`
}

func NewMsgValidatorExit(from hmTypes.HeimdallAddress, id uint64, txhash hmTypes.HeimdallHash, logIndex uint64) MsgValidatorExit {
	return MsgValidatorExit{
		From:     from,
		ID:       hmTypes.NewValidatorID(id),
		TxHash:   txhash,
		LogIndex: logIndex,
	}
}

func (msg MsgValidatorExit) Type() string {
	return "validator-exit"
}

func (msg MsgValidatorExit) Route() string {
	return RouterKey
}

func (msg MsgValidatorExit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{hmTypes.HeimdallAddressToAccAddress(msg.From)}
}

func (msg MsgValidatorExit) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgValidatorExit) ValidateBasic() sdk.Error {
	if msg.ID <= 0 {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid validator ID %v", msg.ID)
	}

	if msg.From.Empty() {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid proposer %v", msg.From.String())
	}

	return nil
}

// GetTxHash Returns tx hash
func (msg MsgValidatorExit) GetTxHash() types.HeimdallHash {
	return msg.TxHash
}

// GetLogIndex Returns log index
func (msg MsgValidatorExit) GetLogIndex() uint64 {
	return msg.LogIndex
}
