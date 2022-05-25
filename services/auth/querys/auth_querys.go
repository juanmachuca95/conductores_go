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

var GetRolesUser = func() string {
	return "SELECT r.id, r.role FROM roles AS r INNER JOIN rolesusers AS ro ON r.id = ro.roles_id WHERE ro.users_id = ?;"
}
