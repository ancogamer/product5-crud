package evento

// Event estrutura basica do evento
type Event struct {
	DataCreate string `json:"-"`
	DataUpdate string `json:"-"`
	Name       string `json:"name"`
	UserCreate string `json:"user_create"`
	DataInit   string `json:"data_init"`
	DataEnd    string `json:"data_end"`
	HourInit   string `json:"hour_init"`
	HourEnd    string `json:"hour_end"`
	QtdMax     int    `json:"qtd_max"`
	QtdNow     int    `json:"qtd_now"`
}
