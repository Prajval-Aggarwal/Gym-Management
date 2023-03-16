package cont

import (
	"encoding/json"
	db "gym-api/Database"
	mod "gym-api/models"
	"net/http"
)

func GetPrices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var memberhips []mod.SubsType
	db.DB.Find(&memberhips)
	json.NewEncoder(w).Encode(&memberhips)
}

func PriceUpdateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var memShip mod.SubsType

	err := json.NewDecoder(r.Body).Decode(&memShip)
	if err != nil {
		panic(err)
	}

	db.DB.Model(&mod.SubsType{}).Where("subs_name =?", memShip.Subs_Name).Updates(&memShip)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Price updated successfully"))

}
