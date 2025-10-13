package service

const (
	STAGE_AWAITING_FRAME  = "awaiting_frame"
	STAGE_AWAITING_LENSES = "awaiting_lenses"
)

type OrderState struct {
	OrderId int64
	Stage   string
}
