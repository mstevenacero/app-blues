package users

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	users "github.com/mstevenacero/application-core-blues/components/users"
	models "github.com/mstevenacero/application-core-blues/models"
)

func GetUsersRoute(w http.ResponseWriter, r *http.Request) {
	response, err := users.GetUsersMongo()
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AddUsersRoute(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if len(bodyBytes) == 0 {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	fmt.Println("Raw body:", string(bodyBytes))
	var user models.UserJson
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Decoded body: %+v\n", user)
	response, err := users.AddUsersMongo(user)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
