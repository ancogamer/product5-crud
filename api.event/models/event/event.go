package evento

// Event estrutura basica do evento
type Event struct {
	UUID       string `json:"event"`
	DataCreate string `json:"-"`
	DataUpdate string `json:"-"`
	Name       string `json:"name"`
	UserCreate string `json:"user_create"`
	DataInit   string `json:"data_init"`
	DataEnd    string `json:"data_end"`
	HourInit   string `json:"hour_init"`
	HourEnd    string `json:"hour_end"`
	QtdMin     int    `json:"qtd_min"`
	QtdMax     int    `json:"qtd_max"`
	QtdNow     int    `json:"qtd_now"`
}

/*
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE  events (
	ev_id serial not null,
	ev_uuid uuid default uuid_generate_v4(),
	ev_data_create date not null default now(),
	ev_date_update date not null default now(),
	ev_hora_create time without time zone not null default now(),
	ev_hora_update time without time zone not null default now(),
	ev_name character varying(300) not null,
	ev_user_create_uuid uuid not null,
	ev_data_init date not null,
	ev_data_end date not null,
	ev_hour_init time without time zone not null,
	ev_hour_end time without time zone not null,
	ev_qtd_max integer not null,
	ev_qtd_now integer,
	ev_qtd_min integer not null default 0,

	CONSTRAINT  ev_pkey PRIMARY KEY(ev_uuid)
)
*/
