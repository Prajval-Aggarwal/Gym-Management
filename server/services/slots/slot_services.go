package slots

import (
	"gym/server/db"
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"
	constants "gym/server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func SlotServices(context *gin.Context, slotRequest request.UpdateSlotRequest) {
	var user model.User
	err := db.FindById(&user, slotRequest.UserId, "user_id")
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	var subscription model.Subscription
	err = db.FindById(&subscription, user.User_Id, "user_id")
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	oldSlotId := subscription.Slot_id

	subscription.Slot_id = slotRequest.SlotId

	var slot model.Slot

	err = db.FindById(&slot, oldSlotId, "slot_id")
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	slot.Available_space += 1
	result := db.UpdateRecord(&slot, oldSlotId, "slot_id")
	if result.Error != nil {
		response.ErrorResponse(context, 500, result.Error.Error())
		return
	}

	var slot1 model.Slot
	err = db.FindById(&slot1, slotRequest.SlotId, "slot_id")
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	slot1.Available_space = slot1.Available_space - 1

	result = db.UpdateRecord(&slot1, slotRequest.SlotId, "slot_id")
	if result.Error != nil {
		response.ErrorResponse(context, 500, result.Error.Error())
		return
	}

	result = db.UpdateRecord(&subscription, slotRequest.UserId, "user_id")
	if result.Error != nil {
		response.ErrorResponse(context, 500, result.Error.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"Slots updated successfully",
		subscription,
		context,
	)
}
func SlotDistribution() {
	startTime, _ := time.Parse("15:04", constants.START_TIME)
	endTime, _ := time.Parse("15:04", constants.END_TIME)
	diff := endTime.Sub(startTime)
	noOfSlots := int(diff.Hours() / constants.SLOT_LEN)

	slotStartTime := startTime
	var slot model.Slot
	for i := 1; i <= noOfSlots; i++ {
		slot.SlotId = i
		seTime := slotStartTime.Add(time.Hour * 2)
		slot.Start_time = slotStartTime.Format("15:04")
		slot.End_time = seTime.Format("15:04")
		slotStartTime = seTime

		err := db.CreateRecord(&slot)
		if err != nil {
			panic(err)
		}

	}
}
