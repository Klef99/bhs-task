CREATE TABLE IF NOT EXISTS public.assets (
	id serial4 NOT NULL,
	"name" text NOT NULL,
	description text NULL,
	price numeric NOT NULL,
	owner_id serial4 NOT NULL,
	CONSTRAINT assets_check CHECK ((price >= (0)::numeric)),
	CONSTRAINT assets_pk PRIMARY KEY (id),
	CONSTRAINT assets_users_fk FOREIGN KEY (owner_id) REFERENCES public.users(id) ON DELETE CASCADE ON UPDATE CASCADE
);