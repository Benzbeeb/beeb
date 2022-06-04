package types

import (
	"encoding/base64"
	fmt "fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type ProposalType string

const (
	ProposalTypeSendIbcData ProposalType = "SendIbcData"
)

// DisableAllProposals contains no wasm gov types.
var DisableAllProposals []ProposalType

// EnableAllProposals contains all wasm gov types as keys.
var EnableAllProposals = []ProposalType{
	ProposalTypeSendIbcData,
}

// ConvertToProposals maps each key to a ProposalType and returns a typed list.
// If any string is not a valid type (in this file), then return an error
func ConvertToProposals(keys []string) ([]ProposalType, error) {
	valid := make(map[string]bool, len(EnableAllProposals))
	for _, key := range EnableAllProposals {
		valid[string(key)] = true
	}

	proposals := make([]ProposalType, len(keys))
	for i, key := range keys {
		if _, ok := valid[key]; !ok {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "'%s' is not a valid ProposalType", key)
		}
		proposals[i] = ProposalType(key)
	}
	return proposals, nil
}

func init() { // register new content types with the sdk
	govtypes.RegisterProposalType(string(ProposalTypeSendIbcData))
	govtypes.RegisterProposalTypeCodec(&SendIbcDataProposal{}, "beeb/SendIbcDataProposal")

}

// ProposalRoute returns the routing key of a parameter change proposal.
func (p SendIbcDataProposal) ProposalRoute() string { return RouterKey }

// GetTitle returns the title of the proposal
func (p *SendIbcDataProposal) GetChanel() string { return p.Channel }

// GetDescription returns the human readable description of the proposal
func (p SendIbcDataProposal) GetData() []byte { return p.Data }

// ProposalType returns the type
func (p SendIbcDataProposal) ProposalType() string { return string(ProposalTypeSendIbcData) }

// ValidateBasic validates the proposal
func (p SendIbcDataProposal) ValidateBasic() error {
	// TODO: add this
	return nil
}

// String implements the Stringer interface.
func (p SendIbcDataProposal) String() string {
	return fmt.Sprintf(`Send Ibc Data Proposal:
  Channel:       			%s
  Timeout Height:			%d
  Timeout Timestamp:      	%d
  Data:    					%X
`, p.Channel, p.TimeoutHeight, p.TimeoutTimestamp, p.Data)
}

// MarshalYAML pretty prints the wasm byte code
func (p SendIbcDataProposal) MarshalYAML() (interface{}, error) {
	return struct {
		Channel          string `yaml:"channel"`
		TimeoutHeight    uint64 `yaml:"timeout_height"`
		TimeoutTimestamp uint64 `yaml:"timeout_timestamp"`
		Data             string `yaml:"data"`
	}{
		Channel:          p.Channel,
		TimeoutHeight:    p.TimeoutHeight,
		TimeoutTimestamp: p.TimeoutTimestamp,
		Data:             base64.StdEncoding.EncodeToString(p.Data),
	}, nil
}
