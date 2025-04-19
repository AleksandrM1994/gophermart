package repository

type OrderStatus string

const (
	OrderStatusUnknown    OrderStatus = "UNKNOWN"
	OrderStatusRegistered OrderStatus = "REGISTERED"
	OrderStatusInvalid    OrderStatus = "INVALID"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusProcessed  OrderStatus = "PROCESSED"
)

func (t OrderStatus) ToString() string {
	return string(t)
}
