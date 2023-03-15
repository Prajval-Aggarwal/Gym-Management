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
	fmt.Println("id is", id)
	dateStr := time.Now().Truncate(time.Hour)
	sub.StartDate = dateStr.Format("02 Jan 2006")

	sub.EndDate = dateStr.AddDate(0, 0, int(sub.Duration*30)).Format("02 Jan 2006")

	sub.User_Id = id
	db.DB.Create(&sub)
	json.NewEncoder(w).Encode(&sub)

}
