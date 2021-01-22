package address

// Address de uma mesa
type Address struct {
	TableUUID      string `json:"table_uuid"`
	DataCreate     string `json:"-"` // data de criação do evento
	DataUpdate     string `json:"-"` // data de atualização do evento
	UserCreateUUID string `json:"user_create"`
	ID             string `json:"-"`
	UUID           string `json:"endereco_uuid"`
	City           string `json:"city"`
	District       string `json:"district"`
	Street         string `json:"street"`
	Number         int    `json:"number"`
	CEP            string `json:"cep"`
}
