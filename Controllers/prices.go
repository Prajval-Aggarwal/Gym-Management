package Controllers

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/Models"
	"net/http"
)

func GetPricesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var memberhips []mod.SubsType
	db.DB.Find(&memberhips)

	//sort the database entries by price
	query:="SELECT * FROM subs_types ORDER BY price ASC"
	db.DB.Raw(query).Scan(&memberhips)
	json.NewEncoder(w).Encode(&memberhips)
}


func PriceUpdateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var membership mod.SubsType

	err := json.NewDecoder(r.Body).Decode(&membership)
	if err != nil {
		panic(err)
	}

	result:=db.DB.Model(&mod.SubsType{}).Where("subs_name =?", membership.Subs_Name).Updates(&membership)
	if result.Error!=nil{
		fmt.Println("error in DB")
	}else if result.RowsAffected == 0 {//if the subs_name is not in record then create new record
		db.DB.Create(&membership)
		fmt.Fprint(w,"New subscription type added")

	}else{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(" Old Price updated successfully"))
	}

}




