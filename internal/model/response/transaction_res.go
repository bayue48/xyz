package response

import "xyz-multifinance/internal/entity"

type TransactionRes struct {
	ID             int64  `json:"id"`
	ContractNumber string `json:"contract_number"`
	CustomerID     int64  `json:"customer_id"`
	OTR            int64  `json:"otr"`
	AdminFee       int64  `json:"admin_fee"`
	Installment    int64  `json:"installment_amount"`
	Interest       int64  `json:"interest_amount"`
	AssetName      string `json:"asset_name"`
	TenorMonth     int    `json:"tenor_month"`
	CreatedAt      string `json:"created_at"`
}

func ToTransactionRes(entity *entity.Transaction) *TransactionRes {
	return &TransactionRes{
		ID:             entity.ID,
		ContractNumber: entity.ContractNumber,
		CustomerID:     entity.CustomerID,
		OTR:            entity.OTR,
		AdminFee:       entity.AdminFee,
		Installment:    entity.Installment,
		Interest:       entity.Interest,
		AssetName:      entity.AssetName,
		TenorMonth:     entity.TenorMonth,
		CreatedAt:      entity.CreatedAt,
	}
}

func ToTransactionResList(entities *[]entity.Transaction) []TransactionRes {
	var res []TransactionRes
	for _, entity := range *entities {
		res = append(res, *ToTransactionRes(&entity))
	}
	return res
}
