package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/models"
	"net/http"
	"time"
)

func CreateSubsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var sub mod.Subscription
	json.NewDecoder(r.Body).Decode(&sub)
	id := r.URL.Query().Get("id")
	// fmt.Println("id is", id)
	dateStr := time.Now().Truncate(time.Hour)
	sub.StartDate = dateStr.Format("02 Jan 2006")

	sub.EndDate = dateStr.AddDate(0, 0, int(sub.Duration*30)).Format("02 Jan 2006")

	sub.User_Id = id
	db.DB.Create(&sub)
	// json.NewEncoder(w).Encode(&sub)

	//show bill according to the subscription chosen and duration given
	var subcription_type mod.SubsType
	db.DB.Where("subs_name=?",sub.Subs_Name).First(&subcription_type)
	
	
	

	//bill amount if duration is 6 months or 12 months
	var billamount float64
	if (sub.Duration==6 ){
		//10% discount
		billamount=(subcription_type.Price*sub.Duration)*0.9
		fmt.Fprintln(w,"10% Discount applied")

	}else if(sub.Duration==12){
		//20% discount
		billamount=(subcription_type.Price*sub.Duration)*0.8
		fmt.Fprintln(w,"20% Discount applied")


	}else{
	billamount=subcription_type.Price*sub.Duration
	}
	fmt.Fprint(w,"\n\n")
	fmt.Fprint(w,"BILL OF SUBSCRIPTION\n\n")

	

	

	fmt.Fprintf(w," Subscription:%s \n Duration:%d \n Bill Amount:%d ",sub.Subs_Name,int(sub.Duration),int(billamount))

}
