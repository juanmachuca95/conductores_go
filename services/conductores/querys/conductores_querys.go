package conductores

var GetConductores = func() string {
	return "SELECT c.id, u.name, u.email, c.users_id, c.matricula, c.vehiculo, c.created_at, c.updated_at FROM conductores as c INNER JOIN users as u ON u.id = c.users_id LIMIT ? OFFSET ?;"
}

var GetConductoresDisponibles = func() string {
	return "SELECT c.id, u.name, u.email, c.users_id, c.matricula, c.vehiculo, c.created_at, c.updated_at FROM conductores as c INNER JOIN users as u ON u.id = c.users_id INNER JOIN viajes v ON v.users_id = c.id WHERE trip_status=1;" // Status 1 han finalizado su viaje y estan disponibles
}

var InsertUser = func() string {
	return "INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?,?,?, NOW(), NOW());"
}

var InsertConductor = func() string {
	return "INSERT INTO conductores (users_id, matricula, vehiculo, created_at, updated_at) VALUES (?,?,?, NOW(), NOW());"
}
