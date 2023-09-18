package entity

import (
	"context"
	"time"

	"github.com/PauloRVF/desafio_client_server_api/server/dto"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Exchange struct {
	Code       string
	Codein     string
	Name       string
	High       string
	Low        string
	VarBid     string
	PctChange  string
	Bid        string
	Ask        string
	Timestamp  string
	CreateDate string
	gorm.Model
}

func NewExchange(economia_api *dto.EconomiaApi) *Exchange {
	return &Exchange{
		Code:       economia_api.Usdbrl.Code,
		Codein:     economia_api.Usdbrl.Codein,
		Name:       economia_api.Usdbrl.Name,
		High:       economia_api.Usdbrl.High,
		Low:        economia_api.Usdbrl.Low,
		VarBid:     economia_api.Usdbrl.VarBid,
		PctChange:  economia_api.Usdbrl.PctChange,
		Bid:        economia_api.Usdbrl.Bid,
		Ask:        economia_api.Usdbrl.Ask,
		Timestamp:  economia_api.Usdbrl.Timestamp,
		CreateDate: economia_api.Usdbrl.CreateDate,
	}
}

func PersistExchange(exchange *Exchange) error {
	db, err := gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Exchange{})
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	return db.WithContext(ctx).Create(&exchange).Error
}
