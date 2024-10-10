CREATE TABLE IF NOT EXISTS public.access_assets (
	user_id serial4 NOT NULL,
	asset_id serial4 NOT NULL,
	CONSTRAINT access_assets_assets_fk FOREIGN KEY (asset_id) REFERENCES public.assets(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT access_assets_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE ON UPDATE CASCADE
);