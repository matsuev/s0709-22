ALTER TABLE IF EXISTS public.contact
    ADD COLUMN father_name character varying(50) NOT NULL DEFAULT '';

ALTER TABLE IF EXISTS public.contact
    ADD COLUMN mother_name character varying(50) NOT NULL DEFAULT '';