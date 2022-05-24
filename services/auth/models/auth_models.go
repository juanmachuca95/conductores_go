package auth

type Register struct {
	Name       string `json:"name"`
	DNI        int64  `json:"dni"`
	Cedula     string `json:"cedula"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
