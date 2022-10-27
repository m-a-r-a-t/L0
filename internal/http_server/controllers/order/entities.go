package order

type GetOrderQuery struct {
	id string `json:"id,omitempty" validate:"required" schema:"id,required"`
}
