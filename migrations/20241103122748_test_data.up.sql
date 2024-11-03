INSERT INTO public.users (username,password_hash,balance) VALUES
	 ('test','$2a$10$BZajwLmJ2neqjCKk.qobnuBoVU1YuLcoZCU9MYhOTsUwhbslQ7rGu',0),
	 ('test2','$2a$10$XWoZTLH5Wv1GFBdnq7T2xu4cPnMz/CEdwfGoo/bu9AwADqBqx41OW',0);

INSERT INTO public.assets (name,description,price) VALUES
	 ('sword','sword',10),
	 ('sword2','sword',100);

INSERT INTO public.access_assets (user_id, asset_id) VALUES
	 (1, 2);
