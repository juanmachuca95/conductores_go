package auth

var GetUser = func() string {
	return "SELECT id, name, dni, username, email FROM users WHERE email = ?;"
}
