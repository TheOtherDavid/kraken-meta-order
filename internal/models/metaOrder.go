package models

import (
	"time"
)

type MetaOrder struct {
	MetaOrderId        int                `json:"metaOrderId"`
	MetaOrderType      string             `json:"metaOrderType"`
	Status             string             `json:"status"`
	Exchange           string             `json:"exchange"`
	StopLossTakeProfit StopLossTakeProfit `json:"stopLossTakeProfit"`
	CreateDateTime     time.Time          `json:"createDateTime"`
	CreateUserName     string             `json:"createUserName"`
	LastUpdateDateTime time.Time          `json:"lastUpdateDateTime"`
	LastUpdateUserName string             `json:"lastUpdateUserName"`
}

type StopLossTakeProfit struct {
	StopLossPrice   float32 `json:"stopLossPrice"`
	TakeProfitPrice float32 `json:"takeProfitPrice"`
}
