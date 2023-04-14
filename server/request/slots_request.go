package request


type UpdateSlotRequest struct {
	UserId string `json:"userId" validate:"required"`
	SlotId int `json:"slotId" validate:"required"`
}