CREATE TABLE public.tag
(
    id    SERIAL UNIQUE PRIMARY KEY,
    title VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO public.tag(title)
VALUES ('Angular'),
       ('GraphQL'),
       ('React'),
       ('NextJS'),
       ('Express'),
       ('NestJS'),
       ('NodeJS'),
       ('Golang'),
       ('MongoDB'),
       ('PostgresQL');

CREATE TABLE public.worktag
(
    id     SERIAL UNIQUE PRIMARY KEY,
    workId INT NOT NULL REFERENCES public.work (id),
    tagId  INT NOT NULL REFERENCES public.tag (id)
);

INSERT INTO public.worktag(workId, tagId)
VALUES (1, 3),
       (1, 5),
       (1, 7),
       (1, 9),

       (2, 4),
       (2, 6),
       (2, 7),
       (2, 9),

       (3, 1),
       (3, 2),
       (3, 8),
       (3, 10);
