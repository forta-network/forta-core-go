package manifest

import "github.com/forta-network/forta-core-go/utils"

// AgentManifest represents the dev-provided properties of an agent
type AgentManifest struct {
	From            *string                       `json:"from"`
	Name            *string                       `json:"name"`
	Description     *string                       `json:"description"`
	LongDescription *string                       `json:"longDescription"`
	AgentID         *string                       `json:"agentId"`
	AgentIDHash     *string                       `json:"agentIdHash"`
	Version         *string                       `json:"version"`
	Timestamp       *string                       `json:"timestamp"`
	ImageReference  *string                       `json:"imageReference"`
	Repository      *string                       `json:"repository"`
	PromoUrl        *string                       `json:"promoUrl"`
	LicenseUrl      *string                       `json:"licenseUrl"`
	Documentation   *string                       `json:"documentation"`
	ChainIDs        []int64                       `json:"chainIds"`
	ChainSettings   map[string]AgentChainSettings `json:"chainSettings"`
	External        bool                          `json:"external"`
	ProtocolVersion int                           `json:"protocolVersion"`
}

// AgentChainSettings is the per-chain configuration of a bot.
type AgentChainSettings struct {
	Shards uint `json:"shards"`
	Target uint `json:"target"`
}

// SignedAgentManifest is the contents of an agent manifest
type SignedAgentManifest struct {
	Manifest  *AgentManifest `json:"manifest"`
	Signature string         `json:"signature"`
}

type ValidationError struct {
	msg string
}

func (m *ValidationError) Error() string {
	return m.msg
}

func validationErr(msg string) *ValidationError {
	return &ValidationError{msg: msg}
}

func (sm *SignedAgentManifest) Validate() error {
	if sm.Manifest == nil {
		return validationErr("manifest is not present")
	}
	// validate images of non-external bots
	if !sm.Manifest.External {
		if sm.Manifest.ImageReference == nil {
			return validationErr("manifest.imageReference is not present")
		}
		if _, ok := utils.ValidateImageRef("", *sm.Manifest.ImageReference); !ok {
			return validationErr("manifest.imageReference is not valid")
		}
	}
	return nil
}
