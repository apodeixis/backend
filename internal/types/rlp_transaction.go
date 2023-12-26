package types

import (
	"database/sql/driver"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type RLPTransaction types.Transaction

func (t RLPTransaction) Value() (driver.Value, error) {
	encoded, err := rlp.EncodeToBytes((*types.Transaction)(&t))
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

func (t *RLPTransaction) Scan(src interface{}) error {
	var encoded []byte
	switch s := src.(type) {
	case string:
		encoded = []byte(s)
	case []byte:
		encoded = s
	default:
		return fmt.Errorf("invalid RLPTransaction value: %v", src)
	}
	var tx types.Transaction
	if err := rlp.DecodeBytes(encoded, &tx); err != nil {
		return err
	}
	*t = RLPTransaction(tx)
	return nil
}
