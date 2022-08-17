CREATE TABLE public.pegawai (
	id int4 NOT NULL,
	nama varchar(100) NULL,
	alamat text NULL,
	telepon varchar(15) NULL,
	CONSTRAINT pegawai_id PRIMARY KEY (id)
);