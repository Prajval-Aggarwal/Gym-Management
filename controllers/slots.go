package cont

import (
	"encoding/json"
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
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("We are updating the slot")
	fmt.Println("Please provide the user id for which you want to update the slot...")
	// if r.Method == "POST" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)
	var sub mod.Subscription
	db.DB.Where("user_id =?", id).First(&sub)
	oldSlotId := sub.Slot_id
	fmt.Println("old slot id is: ", oldSlotId)
	sub1 := make(map[string]int64)

	json.NewDecoder(r.Body).Decode(&sub1)

	newSlotid := sub1["slot_id"]
	sub.Slot_id = newSlotid
	var slot mod.Slot
	db.DB.Where("id =?", oldSlotId).Find(&slot)
	slot.Available_space += 1
	db.DB.Where("id =?", oldSlotId).Updates(&slot)
	var slot1 mod.Slot
	db.DB.Where("id =?", newSlotid).Find(&slot1)
	slot1.Available_space = slot1.Available_space - 1
	db.DB.Where("id =?", newSlotid).Updates(&slot1)
	db.DB.Where("user_id =?", id).Updates(&sub)
	json.NewEncoder(w).Encode(&sub)
	

}
