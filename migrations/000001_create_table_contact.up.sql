CREATE TABLE public.contact
(
   id bigserial NOT NULL,
   full_name character varying(100) NOT NULL DEFAULT '',
   PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.contact
   OWNER to admin;