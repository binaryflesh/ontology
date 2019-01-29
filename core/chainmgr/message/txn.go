/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package message

import (
	"fmt"

	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"github.com/ontio/ontology/errors"
)

type TxResult struct {
	Err  errors.ErrCode `json:"err"`
	Hash common.Uint256 `json:"hash"`
	Desc string         `json:"desc"`
}

type TxRequest struct {
	Tx         *types.Transaction
	TxResultCh chan *TxResult
}

func (this *TxRequest) MarshalJSON() ([]byte, error) {
	sink := common.NewZeroCopySink(nil)
	if this.Tx != nil {
		if err := this.Tx.Serialization(sink); err != nil {
			return nil, fmt.Errorf("TxRequest marshal: %s", err)
		}
	}

	return sink.Bytes(), nil
}

func (this *TxRequest) UnmarshalJSON(data []byte) error {
	tx := &types.Transaction{}
	if err := tx.Deserialization(common.NewZeroCopySource(data)); err != nil {
		return fmt.Errorf("TxRequest unmarshal: %s", err)
	}

	this.Tx = tx
	return nil
}
