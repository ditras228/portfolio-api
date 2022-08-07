CREATE TABLE public.info
(
    id            SERIAL UNIQUE PRIMARY KEY ,
    name          TEXT NOT NULL,
    job           TEXT NOT NULL,
    experience    TEXT NOT NULL,
    img           TEXT NOT NULL,
    telegramTitle TEXT NOT NULL,
    telegramLink  TEXT NOT NULL,
    githubTitle   TEXT NOT NULL,
    githubLink    TEXT NOT NULL
);

INSERT INTO public.info( name, job, experience, img, telegramTitle, telegramLink, githubTitle, githubLink)
VALUES ('Дмитрий Дружинин', 'Middle Fullstack',
        '<ul> <li> коммерческой опыт <br/> разработки в команде </li> <li> исправлял и разрабатывал<br/> функционал на Angular </li> <li> разрабатывал API с <br/>использованием GraphQl </li> <li> разрабатывал серверную<br/> логику на Golang </li> </ul>',
        'uploaded/4bd26626-0a9c-11ed-8284-7412b3c0b125.png',
        '+7 (911) 878-03-02',
        'https://t.me/Ditras',
        'ditras228',
        'https://github.com/ditras228');


CREATE TABLE public.work
(
    id          SERIAL UNIQUE PRIMARY KEY ,
    name        TEXT NOT NULL,
    description TEXT NOT NULL,
    github      TEXT NOT NULL,
    demo        TEXT NOT NULL,
    figma       TEXT
);

INSERT INTO public.work( name, description, github, demo)

VALUES ('Облачное хранилище',
        '<ul> <li> Загрузка, поиск, скачивание<br/> файлов, либо папок </li>  <li> Рекурсивная загрузка папок на сервер, скачивание с помощью ZIP-архива </li> <li> Система drag & drop </li> <li> Система авторизации </li> <li> Рассылка писем на email </li> </ul>',
        'https://github.com/ditras228/cloud-disk',
        'http://localhost:4201'),
       ('Музыкальная платформа',
        '<ul> <li> Загрузка, комментирование треков </li> <li> Группировка по плейлистам </li> <li> Авторизация с помощью OAuth </li> <li> Server Side Rendering </li> <li> Переключение дневной и ночной тем </li> </ul>',
        'https://github.com/ditras228/cloud-disk',
        'http://localhost:5432'),
       ( 'Портфолио',
        '<ul> <li> Презентация работ, раздел «обо мне» </li> <li> Админка: CRUD всех данных, предоставленных на сайте </li>  <li> Оригинальный UI/UX дизайн, Mobile first </li>  <li> Свитч языков, редактирование переводов</li><</ul>',
        'https://github.com/ditras228/MERNPortfolio',
        'http://localhost:4200')