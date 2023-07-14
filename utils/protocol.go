package utils

import (
	"github.com/forta-network/forta-core-go/protocol"
	log "github.com/sirupsen/logrus"
)

// GetChainIDsForAlert is a utility method for collecting mentioned chainIDs in an alert
func GetChainIDsForAlert(f *protocol.SignedAlert) ([]uint64, error) {
	var result []uint64
	// avoid crashes for bad data
	if f.Alert != nil && f.Alert.Finding != nil && f.Alert.Finding.Source != nil {
		for _, c := range f.Alert.Finding.Source.Chains {
			if c.ChainId > 0 {
				result = append(result, c.ChainId)
			}
		}
		for _, t := range f.Alert.Finding.Source.Transactions {
			if t.ChainId > 0 {
				result = append(result, t.ChainId)
			}
		}
		for _, b := range f.Alert.Finding.Source.Blocks {
			if b.ChainId > 0 {
				result = append(result, b.ChainId)
			}
		}
	}

	// this is the scanner's chainID, last in priority
	chainID, err := HexToBigInt(f.ChainId)
	if err != nil {
		log.WithError(err).Error("error parsing signedAlert chainID")
		return nil, err
	}

	result = append(result, chainID.Uint64())
	return UniqUInt64(result), nil
}
