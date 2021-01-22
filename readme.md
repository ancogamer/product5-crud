Bom, aqui vamos começar as discuções, sobre a ideia de produto que surgiu no grupo "Os programadores".
<br>A ideia principal, é um web app, para agrupar as reuniões e eventos sobre programação. 
<br>Fique a vontade para contribuir, seja com codigo, ou com pitaco em telas, etc <3 


Estrutura base: 

-   EVENTO ->  Mesas <- Usuarios 

Por padrão, todos os eventos possuem 1 mesa. 
Vale lembrar que a privacidade de TODAS as mesas, variam entre :
- publico
- privado

 A privacidade do evento é definida pelas mesas, se tiver pelo menos 1 mesa publica, o evento é publico.

 Por hora a estrutura do evento (event já esta fechada), sendo a tabela : 

(
    
    CREATE TABLE  events (
	    ev_id serial not null,
	    ev_uuid uuid default uuid_generate_v4(),
	    ev_data_create date not null default now,
	    ev_date_update date not null default now,
	    ev_name character varying(300) not null,
	    ev_user_create_uuid uuid not null,
	    ev_data_init date not null,
	    ev_data_end date not null,
	    ev_hour_init time without timezone not null,
	    ev_hour_end time without timezone not null,
	    ev_qtd_max integer not null,
	    ev_qtd_now integer,
	    ev_qtd_min integer not null default 0
	    CONSTRAINT  ev_pkey PRIMARY KEY(ev_uuid)
    )

Se tiver alguma melhoria, ou ideia, por favor, envie como Issue ou então PR, explicando qual é <3

 
