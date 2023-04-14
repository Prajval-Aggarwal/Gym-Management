package request

type CreateSubRequest struct {
	UserId   string `json:"userId" validate:"required"`
	SubsName string `json:"subsName" validate:"required"`
	Duration int64  `json:"duration" validate:"required"`
	SlotId   int    `json:"slotId" validate:"required"`
}
type UpdateSubRequest struct {
	UserId   string `json:"userId" validate:"required"`
	SubsName string `json:"subsName" validate:"required"`
}
type EndSubRequest struct {
	UserId string `json:"userId" validate:"required"`
}
