package models

type User struct {
	ID       string `bson:"_id,omitempty"`
	Name     string `bson:"Name"`
	Email    string `bson:"Email"`
	Password string `bson:"Password"`
}

type UserJson struct {
	ID       string `json:"id,omitempty"` // Expose ID as 'id' in JSON, not '_id'
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"Password"` // Exclude the password from the JSON response for security
}
