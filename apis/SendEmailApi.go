package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Wahbi8/PM_Golang/DTO"
	"github.com/Wahbi8/PM_Golang/logger"
)

func SendEmailApi(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var emailInfo dto.EmailInfo

	err := json.NewDecoder(r.Body).Decode(&emailInfo)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	logger.Log.Info().Interface("data", emailInfo).Msg("Received email request")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Data received successfully!")
}
