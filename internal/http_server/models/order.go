package models

type Order struct {
	OrderUid          *string `json:"order_uid" validate:"required"`
	TrackNumber       *string `json:"track_number" validate:"required"`
	Entry             *string `json:"entry" validate:"required"`
	Delivery          `json:"delivery" validate:"required"`
	Payment           `json:"payment" validate:"required"`
	Items             []*Item `json:"items" validate:"required"`
	Locale            *string `json:"locale" validate:"required"`
	InternalSignature *string `json:"internal_signature" validate:"required"`
	CustomerId        *string `json:"customer_id" validate:"required"`
	DeliveryService   *string `json:"delivery_service" validate:"required"`
	Shardkey          *string `json:"shardkey" validate:"required"`
	SmID              *int64  `json:"sm_id" validate:"required"`
	DateCreated       *string `json:"date_created" validate:"required"`
	OofShard          *string `json:"oof_shard" validate:"required"`
}

type Delivery struct {
	Name    *string `json:"name" validate:"required"`
	Phone   *string `json:"phone" validate:"required"`
	Zip     *string `json:"zip" validate:"required"`
	City    *string `json:"city" validate:"required"`
	Address *string `json:"address" validate:"required"`
	Region  *string `json:"region" validate:"required"`
	Email   *string `json:"email" validate:"required"`
}

type Payment struct {
	Transaction  *string `json:"transaction" validate:"required"`
	RequestId    *string `json:"request_id" validate:"required"`
	Currency     *string `json:"currency" validate:"required"`
	Provider     *string `json:"provider" validate:"required"`
	Amount       *int64  `json:"amount" validate:"required"`
	PaymentDt    *int64  `json:"payment_dt" validate:"required"`
	Bank         *string `json:"bank" validate:"required"`
	DeliveryCost *int64  `json:"delivery_cost" validate:"required"`
	GoodsTotal   *int64  `json:"goods_total" validate:"required"`
	CustomFee    *int64  `json:"custom_fee" validate:"required"`
}

type Item struct {
	ChrtId      *int64  `json:"chrt_id" validate:"required"`
	TrackNumber *string `json:"track_number" validate:"required"`
	Price       *int64  `json:"price" validate:"required"`
	Rid         *string `json:"rid" validate:"required"`
	Name        *string `json:"name" validate:"required"`
	Sale        *int64  `json:"sale" validate:"required"`
	Size        *string `json:"size" validate:"required"`
	TotalPrice  *int64  `json:"total_price" validate:"required"`
	NmId        *int64  `json:"nm_id" validate:"required"`
	Brand       *string `json:"brand" validate:"required"`
	Status      *int64  `json:"status" validate:"required"`
}
