package transaction_dao

import "time"

type DealTx struct {
	TxID      uint
	RoomID    uint
	ProductID uint
	TxAmount  uint
	TxTime    time.Time
	UserID    uint
}
