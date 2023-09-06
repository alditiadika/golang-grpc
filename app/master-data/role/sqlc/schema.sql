
CREATE TABLE public.master_role (
	id bigserial NOT NULL,
	"role_code" varchar NOT NULL,
	role_name varchar not NULL,
	is_active bool NOT NULL DEFAULT true,
	created_by int8 NOT NULL,
	created_date timestamp NOT NULL DEFAULT now(),
	modified_date timestamp NOT NULL DEFAULT now(),
	modified_by int8 NOT NULL,
	CONSTRAINT master_role_pk PRIMARY KEY (id)
);