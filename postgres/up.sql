CREATE TABLE public.baseball
(
    value real NOT NULL,
    CONSTRAINT baseball_pkey PRIMARY KEY (value)
);
CREATE TABLE public.football
(
    value real NOT NULL,
    CONSTRAINT footbal_pkey PRIMARY KEY (value)
);
CREATE TABLE public.soccer
(
    value real NOT NULL,
    CONSTRAINT soccer_pkey PRIMARY KEY (value)
);