-- public.address definition

-- Drop table

-- DROP TABLE public.address;

CREATE TABLE public.address (
	id bigserial NOT NULL,
	"name" varchar NULL,
	country varchar NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT address_pk PRIMARY KEY (id)
);