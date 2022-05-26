package auth

type Register struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Rol        string `json:"rol"`
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
	Rol        string `json:"rol,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
}

type UserToken struct {
	Token string `json:"_token"`
}
type RolesUser struct {
	Role string `json:"role"`
}
