package cont

import (
	"fmt"
	db "gym-api/Database"
	mod "gym-api/models"
	ut "gym-api/utils"
	"net/http"
	"time"
)

func SlotDistribution() {
	sTime, _ := time.Parse("15:04", ut.Start_time)
	eTime, _ := time.Parse("15:04", ut.End_time)
	diff := eTime.Sub(sTime)
	noOfSlots := int64(diff.Hours() / ut.Slot_length)
	//fmt.Println("number of slots is:", noOfSlots)
	ssTime := sTime
	var slot mod.Slot
	//fmt.Println("sstime is", ssTime)
	var i int64
	for i = 1; i <= noOfSlots; i++ {

		slot.ID = i
		seTime := ssTime.Add(time.Hour * 2)

		slot.Start_time = ssTime.Format("15:04")
		slot.End_time = seTime.Format("15:04")
		ssTime = seTime

		db.DB.Create(&slot)
	}

}

// slot update handler
func SlotUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We are updating the slot")
	fmt.Println("Please provide the user id for which you want to update the slot...")
	if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)
	w.Header().Set("Content-Type", "application/")
	var sub mod.Subscription
	db.DB.Where("user_id=?",sub.User_Id).Updates(&sub)

	//	TODO add the person to the new slot and delete a person from previous slot
	
}
