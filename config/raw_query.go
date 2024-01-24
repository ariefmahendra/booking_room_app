package config

const (
	RawQueryInsertEmployee     = `INSERT INTO mst_employee (name, email, password, division, position, role, contact) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, email, division, position, role, contact, created_at`
	RawQueryGetEmployeeById    = `SELECT id, name, email, division, position, role, contact, created_at, updated_at, deleted_at FROM mst_employee WHERE id = $1`
	RawQueryGetEmployeeByEmail = `SELECT id, name, email, division, position, role, contact, created_at, updated_at, deleted_at FROM mst_employee WHERE email = $1`
	RawQueryGetEmployees       = `SELECT id, name, email, division, position, role, contact, created_at, updated_at, deleted_at FROM mst_employee LIMIT $1 OFFSET $2`
	RawDeleteEmployeeById      = `UPDATE mst_employee SET deleted_at = (CURRENT_TIMESTAMP) WHERE id = $1`
	RawUpdateEmployeeById      = `UPDATE mst_employee SET name = $1, email = $2, password = $3, division = $4, position = $5, role = $6, contact = $7, updated_at = (CURRENT_TIMESTAMP) WHERE id = $8 RETURNING id, name, email, division, position, role, contact, created_at, updated_at`
	RAWQueryPaging             = `SELECT COUNT (*) FROM mst_employee`
)
