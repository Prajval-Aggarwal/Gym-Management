package cont

import (
	"fmt"
	db "gym-api/Database"
	mod "gym-api/models"
	cons "gym-api/utils"
	"time"
)

func SlotDistribution() {
	sTime, _ := time.Parse("15:04", cons.Start_time)
	eTime, _ := time.Parse("15:04", cons.End_time)
	diff := eTime.Sub(sTime)
	noOfSlots := int(diff.Hours() / cons.SlotLen)
	//fmt.Println("number of slots is:", noOfSlots)
	ssTime := sTime
	var slot mod.Slot
	//fmt.Println("sstime is", ssTime)
	for i := 1; i <= noOfSlots; i++ {
		slot.ID = i
		seTime := ssTime.Add(time.Hour * 2)
		fmt.Println("setie is", seTime)
		slot.Start_time = ssTime.Format("15:04")
		slot.End_time = seTime.Format("15:04")
		ssTime = seTime
		fmt.Println("slot is:", slot)
		db.DB.Create(&slot)
	}

}
