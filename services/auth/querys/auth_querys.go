package auth

var GetUser = func() string {
	return "SELECT id, name, email, password FROM users WHERE email = ?;"
}
