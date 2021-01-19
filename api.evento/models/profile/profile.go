package models

// User struct, utilizada durante gets
type User struct {
	ID             uint   `json:"-"`
	UUID           string `json:"uuid"`
	Nickname       string `json:"nickname" form:"nickname"` //max 15
	Name           string `json:"name" form:"name"`
	Lastname       string `json:"lastname" form:"lastname"`
	DataNascimento string `json:"datanascimento" form:"datanascimento" ` //maximo 8 digitos
	Deleted        bool   `json:"deleted" form:"deleted"`
}
