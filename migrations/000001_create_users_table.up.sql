CREATE TABLE IF NOT EXISTS public.users (
	id serial4 NOT NULL,
	username text NOT NULL,
	password_hash text NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_unique UNIQUE (username)
);
