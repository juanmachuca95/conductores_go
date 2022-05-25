package conductores

var GetConductores = func() string {
	return "SELECT c.id, u.name, u.email, c.users_id, c.matricula, c.vehiculo, c.created_at, c.updated_at FROM conductores as c INNER JOIN users as u ON u.id = c.users_id LIMIT ? OFFSET ?;"
}
