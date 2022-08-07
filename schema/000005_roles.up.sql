CREATE TABLE public.role(
    id SERIAL UNIQUE  PRIMARY KEY,
    name TEXT
);

INSERT INTO public.role(name)
VALUES  ('USER'),
        ('ADMIN');

ALTER TABLE public.user ADD COLUMN  role INT REFERENCES public.role("id") DEFAULT 1;

-- Присвоение рутовому юзеру роли админа
UPDATE public.user SET "role" = 2 WHERE "id" = 1
