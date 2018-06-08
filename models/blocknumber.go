package models

import (
	"fmt"

	"github.com/SmartMeshFoundation/SmartRaiden/log"
)

const bucketBlockNumber = "bucketBlockNumber"
const keyBlockNumber = "blocknumber"

//GetLatestBlockNumber lastest block number
func (m *ModelDB) GetLatestBlockNumber() int64 {
	var number int64
	err := m.db.Get(bucketBlockNumber, keyBlockNumber, &number)
	if err != nil {
		log.Error(fmt.Sprintf("models GetLatestBlockNumber err=%s", err))
	}
	return number
}

//SaveLatestBlockNumber block numer has been processed
func (m *ModelDB) SaveLatestBlockNumber(blockNumber int64) {
	err := m.db.Set(bucketBlockNumber, keyBlockNumber, blockNumber)
	if err != nil {
		log.Error(fmt.Sprintf("models SaveLatestBlockNumber err=%s", err))
	}
}