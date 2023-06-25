\c aquafarm;

CREATE TABLE public.farms (
  farm_id SERIAL PRIMARY KEY,
  farm_name VARCHAR,
  description VARCHAR,
  status int2 DEFAULT 1
);

CREATE TABLE public.ponds (
  pond_id SERIAL PRIMARY KEY,
  farm_id int4 NOT NULL,
  pond_name VARCHAR,
  description VARCHAR,
  status int2 DEFAULT 1
);