CREATE TABLE public.pegawai (
	id int4 NOT NULL,
	nama varchar(100) NOT NULL,
	alamat text NULL,
	telepon varchar(15) NULL,
	CONSTRAINT pegawai_id PRIMARY KEY (id)
);


CREATE TABLE public.users (
	id int4 NOT NULL,
	username varchar(100) NOT NULL,
	password VARCHAR(100) NOT NULL,
	CONSTRAINT users_id PRIMARY KEY (id)
);