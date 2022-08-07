CREATE TABLE public.desc
(
    id     SERIAL UNIQUE PRIMARY KEY,
    text   TEXT NOT NULL,
    img    TEXT NOT NULL
);

INSERT INTO public.desc(text, img)
VALUES ( '<b>Созданием</b> сайтов занимаюсь более <b>трёх</b> лет', 'uploaded/e9b2fe8c-0aad-11ed-8284-7412b3c0b125.png'),
       ( 'Работал с <b>множеством</b> фронтенд-фреймворков <br/> <b>Angular</b>, <b>Next</b>, <b>Vue</b>', 'uploaded/d489e051-0aad-11ed-8284-7412b3c0b125.png'),
       ( 'На работе, <b>кроме</b> серверной и клиентской логики, также <b>практиковал верстку</b>', 'uploaded/dbae8256-0aad-11ed-8284-7412b3c0b125.png'),
       ( 'На бекенде использую <b>Golang</b>. <br/> Работал с <b>разными</b> видами БД', 'uploaded/e3269f11-0aad-11ed-8284-7412b3c0b125.png')

