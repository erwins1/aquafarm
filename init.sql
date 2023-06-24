\c aquafarm;

CREATE TABLE public.farms (
  farm_id SERIAL PRIMARY KEY,
  farm_name VARCHAR
);

CREATE TABLE public.ponds (
  pond_id SERIAL PRIMARY KEY,
  farm_id int4 NOT NULL,
  farm_name VARCHAR
);