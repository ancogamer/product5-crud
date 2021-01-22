package table

// Table struct de mesa aonde está amarrado os usuarios
type Table struct {
	UUID           string `json:"table"` // UUID da tabela
	TabEvUUID      string `json:"-"`     // uuid dos eventos
	ID             int    `json:"-"`     // ID da tabela
	DataCreate     string `json:"-"`     // data de criação do evento
	DataUpdate     string `json:"-"`     // data de atualização do evento
	UserCreateUUID string `json:"user_create"`
	Name           string `json:"name" form:"name"`                //	Nome
	Description    string `json:"description" form:"description" ` //	Descrição
	QtdUsers       int    `json:"users" form:"users"`              //	Numero atual de participantes
	QtdUsersMax    int    `json:"users_max" form:"users_max"`      //	Numero máximo de participantes
	Privacy        int    `json:"privacy" form:"privacy"`          //	Privacidade
	Modalidade     int    `json:"modalidade"`                      // Presencial ou remoto
	HoraMesaInit   string `json:"hora_mesa_init"`
	HoraMesaEnd    string `json:"hora_mesa_end"`
}

/*
CREATE TABLE tables(
	tab_id serial NOT NULL,
	tab_uuid uuid DEFAULT uuid_generate_v4(),
	tab_ev_uuid uuid NOT NULL,
	tab_data_create date NOT NULL DEFAULT now(),
	tab_date_update date NOT NULL DEFAULT now(),
	tab_hora_create time without time zone NOT NULL DEFAULT now(),
	tab_hora_update time without time zone NOT NULL DEFAULT now(),
	tab_user_create_uuid uuid NOT NULL,
	tab_name
	tab_description
	tab_qtdUser
	tab_MaxOfParticipants
	tab_Privacy
	tab_Modalidade
	tab_HoraMesaInit
	tab_HoraMesaEnd

)
*/
