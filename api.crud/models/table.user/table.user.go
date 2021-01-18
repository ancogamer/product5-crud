package tableuser

import "github.com/ancogamer/product5/api.crud/models/table"

// TableUser estrutura relação Mesa X Usuario
type TableUser struct {
	TableUUID string `json:"table_uuid"`
	UserUUID  string `json:"user_uuid"`
	Permition int    `json:"permition"`
}

// UpdateTableUser utilizada pelo dono da mesa ou então pelos moderados para atualizar algum dado
type UpdateTableUser struct {
	table.Table
	UserUUID string
}
