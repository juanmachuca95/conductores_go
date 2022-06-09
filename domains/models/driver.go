package models

type Driver struct {
	Id         int64  `json:"id"`
	Users_id   int64  `json:"users_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Matricula  string `json:"matricula"`
	Vehiculo   string `json:"vehiculo"`
	Created_at string `json:"created_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
}
