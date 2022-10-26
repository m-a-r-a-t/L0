package order

type GetOrderQuery struct {
	id string `json:"id,omitempty" validate:"required" schema:"id,required"`
}

type Order struct {
	OrderUid          string `json:"order_uid,omitempty" validate:"required"`
	TrackNumber       string `json:"track_number,omitempty" validate:"required"`
	Entry             string `json:"entry,omitempty" validate:"required"`
	Delivery          `json:"delivery,omitempty" validate:"required"`
	Payment           `json:"payment,omitempty" validate:"required"`
	Items             `json:"items,omitempty" validate:"required"`
	Locale            string `json:"locale,omitempty" validate:"required"`
	InternalSignature string `json:"internal_signature,omitempty" validate:"required"`
	CustomerId        string `json:"customer_id,omitempty" validate:"required"`
	DeliveryService   string `json:"delivery_service,omitempty" validate:"required"`
	Shardkey          string `json:"shardkey,omitempty" validate:"required"`
	SmID              int64  `json:"sm_id,omitempty" validate:"required"`
	DateCreated       string `json:"date_created,omitempty" validate:"required"`
	OofShard          string `json:"oof_shard,omitempty" validate:"required"`
}

type Delivery struct {
	Name    string `json:"name,omitempty" validate:"required"`
	Phone   string `json:"phone,omitempty" validate:"required"`
	Zip     string `json:"zip,omitempty" validate:"required"`
	City    string `json:"city,omitempty" validate:"required"`
	Address string `json:"address,omitempty" validate:"required"`
	Region  string `json:"region,omitempty" validate:"required"`
	Email   string `json:"email,omitempty" validate:"required"`
}

type Payment struct {
	Transaction  string `json:"transaction,omitempty" validate:"required"`
	RequestId    string `json:"request_id,omitempty" validate:"required"`
	Currency     string `json:"currency,omitempty" validate:"required"`
	Provider     string `json:"provider,omitempty" validate:"required"`
	Amount       int64  `json:"amount,omitempty" validate:"required"`
	PaymentDt    int64  `json:"payment_dt,omitempty" validate:"required"`
	Bank         string `json:"bank,omitempty" validate:"required"`
	DeliveryCost int64  `json:"delivery_cost,omitempty" validate:"required"`
	GoodsTotal   int64  `json:"goods_total,omitempty" validate:"required"`
	CustomFee    int64  `json:"custom_fee,omitempty" validate:"required"`
}

type Items struct {
	ChrtId      int64  `json:"chrt_id,omitempty" validate:"required"`
	TrackNumber string `json:"track_number,omitempty" validate:"required"`
	Price       int64  `json:"price,omitempty" validate:"required"`
	Rid         string `json:"rid,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"required"`
	Sale        int64  `json:"sale,omitempty" validate:"required"`
	Size        string `json:"size,omitempty" validate:"required"`
	TotalPrice  int64  `json:"total_price,omitempty" validate:"required"`
	NmId        int64  `json:"nm_id,omitempty" validate:"required"`
	Brand       string `json:"brand,omitempty" validate:"required"`
	Status      int64  `json:"status,omitempty" validate:"required"`
}
