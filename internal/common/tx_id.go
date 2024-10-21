package common

import (
	"fmt"
	"github.com/mitchellh/hashstructure/v2"
)

func SetIDForTransaction(tx *CointrackingTx) error {
	hash, err := hashstructure.Hash(tx, hashstructure.FormatV2, nil)
	if err != nil {
		return err
	}
	tx.ID = fmt.Sprintf("%x", hash)
	return nil
}
