CREATE TABLE IF NOT EXISTS public.users (
	id serial4 NOT NULL,
	username text NOT NULL,
	password_hash text NOT NULL,
	balance numeric NOT NULL DEFAULT 0,
	CONSTRAINT users_check CHECK ((balance >= (0)::numeric)),
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_unique UNIQUE (username)
);
