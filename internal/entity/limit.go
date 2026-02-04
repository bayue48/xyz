package entity

type CustomerLimit struct {
	ID          int64 `gorm:"column:id;primaryKey;autoIncrement"`
	CustomerID  int64 `gorm:"column:customer_id;not null;index:uq_customer_tenor,unique"`
	TenorMonth  int   `gorm:"column:tenor_month;not null;index:uq_customer_tenor,unique"`
	LimitAmount int64 `gorm:"column:limit_amount;not null"`
	UsedAmount  int64 `gorm:"column:used_amount;not null"`
}

func (cl *CustomerLimit) TableName() string {
	return "customer_limits"
}
