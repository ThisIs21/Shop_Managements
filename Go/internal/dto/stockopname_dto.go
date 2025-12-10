package dto

import "time"

type OpnameItemReq struct {
	ProductID   uint `json:"product_id" validate:"required"`
	QtyPhysical int  `json:"qty_fisik" validate:"gte=0"`
}

type CreateOpnameReq struct {
	Date  string      `json:"date" time_format:"2006-01-02" validate:"required"`
	Note  string         `json:"note"`
	Items []OpnameItemReq `json:"items" validate:"required,dive"`
}

type CreateOpnameReqService struct {
    Date  time.Time
    Note  string
    Items []OpnameItemReq
}

//ctrl z till createopnameservice DISSAPPEAR