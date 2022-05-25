package auth

var GetUser = func() string {
	return "SELECT id, name, email, password FROM users WHERE email = ?;"
}

var GetUserById = func() string {
	return "SELECT id, name, email, rol FROM users WHERE id = ?;"
}

var InsertUser = func() string {
	return "INSERT INTO users (name, email, password, rol, created_at, updated_at) VALUES (?,?,?,?, NOW(), NOW());"
}
