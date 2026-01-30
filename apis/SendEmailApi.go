package apis

import(
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/Wahbi8/PM_Golang/DTO"
)

func SendEmailApi(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var emailInfo dto.EmailInfo

	err := json.NewDecoder(r.Body).Decode(&emailInfo)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest )
		return
	}

	fmt.Printf("Recieved the data from c#: %+v\n", emailInfo)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Data received successfully!")
}