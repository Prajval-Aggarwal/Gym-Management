package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/Models"
	cons "gym-api/Utils"
	"net/http"
	"time"
)

func SlotDistribution() {
	startTime, _ := time.Parse("15:04", cons.Start_time)
	endTime, _ := time.Parse("15:04", cons.End_time)
	diff := endTime.Sub(startTime)
	noOfSlots := int(diff.Hours() / cons.SlotLen)

	//fmt.Println("number of slots is:", noOfSlots)
	slotStartTime := startTime
	var slot mod.Slot
	for i := 1; i <= noOfSlots; i++ {
		slot.ID = i
		seTime := slotStartTime.Add(time.Hour * 2)
		slot.Start_time = slotStartTime.Format("15:04")
		slot.End_time = seTime.Format("15:04")
		slotStartTime = seTime
		//fmt.Println("slot is:", slot)

		db.DB.Create(&slot)
	}

}

// slot update handler
func SlotUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)
	var user mod.User
	db.DB.Where("user_id = ?", id).Find(&user)
	fmt.Println("user: ", user)
	if user.User_Id == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User with id %s not found", id)
		return
	}
	var subscription mod.Subscription
	db.DB.Where("user_id =?", id).First(&subscription)
	oldSlotId := subscription.Slot_id
	fmt.Println("old slot id is: ", oldSlotId)
	sub1 := make(map[string]int64)

	json.NewDecoder(r.Body).Decode(&sub1)

	newSlotid := sub1["slot_id"]
	subscription.Slot_id = newSlotid
	var slot mod.Slot
	db.DB.Where("id =?", oldSlotId).Find(&slot)
	slot.Available_space += 1
	db.DB.Where("id =?", oldSlotId).Updates(&slot)
	var slot1 mod.Slot
	db.DB.Where("id =?", newSlotid).Find(&slot1)
	slot1.Available_space = slot1.Available_space - 1
	db.DB.Where("id =?", newSlotid).Updates(&slot1)
	db.DB.Where("user_id =?", id).Updates(&subscription)
	json.NewEncoder(w).Encode(&subscription)

}
