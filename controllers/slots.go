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
	noOfSlots := int64(diff.Hours() / ut.SlotLen)
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

	json.NewDecoder(r.Body).Decode(&sub)
//delete a person from previous slot
	
	var slot_to_erase mod.Slot

	db.DB.Where("ID=?",sub.Slot_id).First(&slot_to_erase)

	slot_to_erase.Available_space+=1

	//--------update db-----------
	db.DB.Where("user_id=?",id).Updates(&sub)  

//	add the person to the new slot 
	var slot_to_add mod.Slot

	db.DB.Where("ID=?",sub.Slot_id).First(&slot_to_add)

	slot_to_add.Available_space-=1

	fmt.Fprint(w,"Slot updated successfully!!")


}