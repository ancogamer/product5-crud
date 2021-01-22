package table

// Table struct de mesa aonde está amarrado os usuarios
type Table struct {
	TabEvUUID            string `json:"-"` // uuid dos eventos
	DataCreate           string `json:"-"` // data de criação do evento
	DataUpdate           string `json:"-"` // data de atualização do evento
	UserCreateUUID       string `json:"user_create"`
	ID                   int    `json:"-"`                                   // 	ID da tabela
	UUID                 string `json:"uuid"`                                // 	UUID da tabela
	Name                 string `json:"name" form:"name"`                    //	Nome
	Description          string `json:"description" form:"description" `     //	Descrição
	NumberOfParticipants int    `json:"participants" form:"participants"`    //	Numero atual de participantes
	MaxOfParticipants    int    `json:"maxparticipants" form:"participants"` //	Numero máximo de participantes
	Privacy              int    `json:"privacy" form:"privacy"`              //	Privacidade
	Modalidade           int    `json:"modalidade"`                          // Presencial ou remoto
	HoraMesaInit         string `json:"hora_mesa_init"`
	HoraMesaEnd          string `json:"hora_mesa_end"`
}
