package cont

import (
	"net/http"
)

func APIdocsHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "https://app.swaggerhub.com/apis/PRAJVAL_1/gym-api/1.0.0", http.StatusMovedPermanently)

}
