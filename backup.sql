--
-- PostgreSQL database dump
--

-- Dumped from database version 15.6
-- Dumped by pg_dump version 15.6

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: role; Type: TYPE; Schema: public; Owner: diploma
--

CREATE TYPE public.role AS ENUM (
    'admin',
    'user'
);


ALTER TYPE public.role OWNER TO diploma;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: authors; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.authors (
    id integer NOT NULL,
    name character varying(100) NOT NULL
);


ALTER TABLE public.authors OWNER TO diploma;

--
-- Name: authors_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.authors_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.authors_id_seq OWNER TO diploma;

--
-- Name: authors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.authors_id_seq OWNED BY public.authors.id;


--
-- Name: book_comments; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.book_comments (
    id integer NOT NULL,
    book_id integer NOT NULL,
    user_id integer NOT NULL,
    comment text NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.book_comments OWNER TO diploma;

--
-- Name: book_comments_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.book_comments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.book_comments_id_seq OWNER TO diploma;

--
-- Name: book_comments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.book_comments_id_seq OWNED BY public.book_comments.id;


--
-- Name: book_ratings; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.book_ratings (
    id integer NOT NULL,
    book_id integer NOT NULL,
    user_id integer NOT NULL,
    rating integer NOT NULL
);


ALTER TABLE public.book_ratings OWNER TO diploma;

--
-- Name: book_ratings_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.book_ratings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.book_ratings_id_seq OWNER TO diploma;

--
-- Name: book_ratings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.book_ratings_id_seq OWNED BY public.book_ratings.id;


--
-- Name: books; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.books (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    pub_date date NOT NULL,
    edition character varying(50) NOT NULL,
    language character varying(50) NOT NULL,
    rating numeric(5,2) NOT NULL,
    image character varying(255) DEFAULT ''::character varying,
    description text DEFAULT ''::text
);


ALTER TABLE public.books OWNER TO diploma;

--
-- Name: books_authors; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.books_authors (
    book_id integer NOT NULL,
    author_id integer NOT NULL
);


ALTER TABLE public.books_authors OWNER TO diploma;

--
-- Name: books_categories; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.books_categories (
    book_id integer NOT NULL,
    category_id integer NOT NULL
);


ALTER TABLE public.books_categories OWNER TO diploma;

--
-- Name: books_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.books_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.books_id_seq OWNER TO diploma;

--
-- Name: books_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;


--
-- Name: categories; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(100) NOT NULL
);


ALTER TABLE public.categories OWNER TO diploma;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.categories_id_seq OWNER TO diploma;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: friends; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.friends (
    id integer NOT NULL,
    user_id integer NOT NULL,
    friend_id integer NOT NULL,
    status character varying(50) NOT NULL
);


ALTER TABLE public.friends OWNER TO diploma;

--
-- Name: friends_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.friends_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.friends_id_seq OWNER TO diploma;

--
-- Name: friends_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.friends_id_seq OWNED BY public.friends.id;


--
-- Name: goose_db_version; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goose_db_version OWNER TO diploma;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.goose_db_version_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.goose_db_version_id_seq OWNER TO diploma;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.goose_db_version_id_seq OWNED BY public.goose_db_version.id;


--
-- Name: liked_books; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.liked_books (
    id integer NOT NULL,
    user_id integer NOT NULL,
    book_id integer NOT NULL
);


ALTER TABLE public.liked_books OWNER TO diploma;

--
-- Name: liked_books_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.liked_books_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.liked_books_id_seq OWNER TO diploma;

--
-- Name: liked_books_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.liked_books_id_seq OWNED BY public.liked_books.id;


--
-- Name: question_results; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.question_results (
    id integer NOT NULL,
    quiz_result_id integer NOT NULL,
    quiestion_id integer NOT NULL,
    user_answer character varying(255) NOT NULL,
    is_correct boolean NOT NULL
);


ALTER TABLE public.question_results OWNER TO diploma;

--
-- Name: question_results_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.question_results_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.question_results_id_seq OWNER TO diploma;

--
-- Name: question_results_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.question_results_id_seq OWNED BY public.question_results.id;


--
-- Name: questions; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.questions (
    id integer NOT NULL,
    quiz_id integer NOT NULL,
    question text NOT NULL,
    options jsonb NOT NULL,
    answer character varying(255) NOT NULL
);


ALTER TABLE public.questions OWNER TO diploma;

--
-- Name: questions_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.questions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.questions_id_seq OWNER TO diploma;

--
-- Name: questions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.questions_id_seq OWNED BY public.questions.id;


--
-- Name: quiz_comments; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.quiz_comments (
    id integer NOT NULL,
    quiz_id integer NOT NULL,
    user_id integer NOT NULL,
    comment text NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.quiz_comments OWNER TO diploma;

--
-- Name: quiz_comments_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.quiz_comments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.quiz_comments_id_seq OWNER TO diploma;

--
-- Name: quiz_comments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.quiz_comments_id_seq OWNED BY public.quiz_comments.id;


--
-- Name: quiz_ratings; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.quiz_ratings (
    id integer NOT NULL,
    quiz_id integer NOT NULL,
    user_id integer NOT NULL,
    rating integer NOT NULL
);


ALTER TABLE public.quiz_ratings OWNER TO diploma;

--
-- Name: quiz_ratings_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.quiz_ratings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.quiz_ratings_id_seq OWNER TO diploma;

--
-- Name: quiz_ratings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.quiz_ratings_id_seq OWNED BY public.quiz_ratings.id;


--
-- Name: quiz_results; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.quiz_results (
    id integer NOT NULL,
    quiz_id integer NOT NULL,
    user_id integer NOT NULL,
    coorect integer NOT NULL,
    total integer NOT NULL
);


ALTER TABLE public.quiz_results OWNER TO diploma;

--
-- Name: quiz_results_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.quiz_results_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.quiz_results_id_seq OWNER TO diploma;

--
-- Name: quiz_results_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.quiz_results_id_seq OWNED BY public.quiz_results.id;


--
-- Name: quizzes; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.quizzes (
    id integer NOT NULL,
    user_id integer NOT NULL,
    book_id integer NOT NULL,
    title character varying(100) NOT NULL,
    rating numeric(5,2) NOT NULL,
    created_at timestamp with time zone DEFAULT now()
);


ALTER TABLE public.quizzes OWNER TO diploma;

--
-- Name: quizzes_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.quizzes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.quizzes_id_seq OWNER TO diploma;

--
-- Name: quizzes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.quizzes_id_seq OWNED BY public.quizzes.id;


--
-- Name: share_requests; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.share_requests (
    id integer NOT NULL,
    sender_id integer NOT NULL,
    receiver_id integer NOT NULL,
    sender_book_id integer NOT NULL,
    receiver_book_id integer NOT NULL,
    sender_status character varying(100) NOT NULL,
    receiver_status character varying(100) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.share_requests OWNER TO diploma;

--
-- Name: share_requests_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.share_requests_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.share_requests_id_seq OWNER TO diploma;

--
-- Name: share_requests_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.share_requests_id_seq OWNED BY public.share_requests.id;


--
-- Name: stock_books; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.stock_books (
    id integer NOT NULL,
    user_id integer NOT NULL,
    book_id integer NOT NULL
);


ALTER TABLE public.stock_books OWNER TO diploma;

--
-- Name: stock_books_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.stock_books_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.stock_books_id_seq OWNER TO diploma;

--
-- Name: stock_books_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.stock_books_id_seq OWNED BY public.stock_books.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: diploma
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(255) NOT NULL,
    role public.role DEFAULT 'user'::public.role NOT NULL,
    city character varying(100)
);


ALTER TABLE public.users OWNER TO diploma;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: diploma
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO diploma;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: diploma
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: authors id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.authors ALTER COLUMN id SET DEFAULT nextval('public.authors_id_seq'::regclass);


--
-- Name: book_comments id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.book_comments ALTER COLUMN id SET DEFAULT nextval('public.book_comments_id_seq'::regclass);


--
-- Name: book_ratings id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.book_ratings ALTER COLUMN id SET DEFAULT nextval('public.book_ratings_id_seq'::regclass);


--
-- Name: books id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: friends id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.friends ALTER COLUMN id SET DEFAULT nextval('public.friends_id_seq'::regclass);


--
-- Name: goose_db_version id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.goose_db_version ALTER COLUMN id SET DEFAULT nextval('public.goose_db_version_id_seq'::regclass);


--
-- Name: liked_books id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.liked_books ALTER COLUMN id SET DEFAULT nextval('public.liked_books_id_seq'::regclass);


--
-- Name: question_results id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.question_results ALTER COLUMN id SET DEFAULT nextval('public.question_results_id_seq'::regclass);


--
-- Name: questions id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.questions ALTER COLUMN id SET DEFAULT nextval('public.questions_id_seq'::regclass);


--
-- Name: quiz_comments id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.quiz_comments ALTER COLUMN id SET DEFAULT nextval('public.quiz_comments_id_seq'::regclass);


--
-- Name: quiz_ratings id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.quiz_ratings ALTER COLUMN id SET DEFAULT nextval('public.quiz_ratings_id_seq'::regclass);


--
-- Name: quiz_results id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.quiz_results ALTER COLUMN id SET DEFAULT nextval('public.quiz_results_id_seq'::regclass);


--
-- Name: quizzes id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.quizzes ALTER COLUMN id SET DEFAULT nextval('public.quizzes_id_seq'::regclass);


--
-- Name: share_requests id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.share_requests ALTER COLUMN id SET DEFAULT nextval('public.share_requests_id_seq'::regclass);


--
-- Name: stock_books id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.stock_books ALTER COLUMN id SET DEFAULT nextval('public.stock_books_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: authors; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.authors (id, name) FROM stdin;
121	Николас Спаркс
122	Анна Джейн
123	Фрэнк Герберт
124	Николас Спаркс
125	Агата Кристи
126	Анна Джейн
127	Анна Джейн
128	Адам Сильвера
129	Янагихара Ханья
130	Нора Сакавич
131	Эрих Мария Ремарк
132	Николас Спаркс
133	Нора Сакавич
134	Али Хейзелвуд
135	Х. Д. Карлтон
136	Катерина Сильванова
137	Фрэнк Герберт
138	Шарлотта Бронте
139	Джейн Остин
140	Мосян Тунсю
141	Агата Кристи
142	Нора Сакавич
143	Агата Кристи
144	Луиза Мэй Олкотт
145	Анна Джейн
146	Оскар Уайльд
147	Анна Джейн
148	Эрих Мария Ремарк
149	Тилли Коул
150	Чак Паланик
151	Мосян Тунсю
152	Агата Кристи
153	Али Хейзелвуд
154	Халед Хоссейни
155	Анна Джейн
156	Міржақып Дулатұлы
157	Эмили Джейн Бронте
158	Эрих Мария Ремарк
159	Фрэнсис Скотт Фицджеральд
160	Агата Кристи
161	Харпер Ли
162	Эрих Мария Ремарк
163	Агустина Бастеррика
164	Энн Бронте
165	Холли Блэк
166	Айн Рэнд
167	Юкио Мисима
168	Хлоя Уолш
169	Мосян Тунсю
170	Уильям Сомерсет Моэм
171	Фрэнсис Элиза Ходжсон Бернетт
172	Екатерина Звонцова
173	Ли Бардуго
174	Ребекка Яррос
175	Анна Джейн
176	Мосян Тунсю
177	Элис Уокер
178	Колин Маккалоу
179	Николас Спаркс
180	Федор Михайлович Достоевский
181	Элиф Шафак
182	Альбер Камю
183	Джек Лондон
184	Михаил Афанасьевич Булгаков
185	Мариам Петросян
186	Герман Гессе
187	Холли Блэк
188	Агата Кристи
189	Фрэнк Герберт
190	Фредрик Бакман
191	Юлия Вереск
192	Анна Джейн
193	Джордж Оруэлл
194	Халед Хоссейни
195	Кристина Старк
196	Ана Шерри
197	Дэниел Киз
198	Агата Кристи
199	Кристина Старк
200	Тартт Донна
201	Мосян Тунсю
202	Виктор Мари Гюго
203	Николас Спаркс
204	Фредрик Бакман
205	Агата Кристи
206	Харуки Мураками
207	Николас Спаркс
208	Михаил Афанасьевич Булгаков
209	Уильям Сомерсет Моэм
210	Фрэнк Герберт
211	Дэвид Левитан
212	Мосян Тунсю
213	Дана Делон
214	Агата Кристи
215	Фрэнк Герберт
216	Джеймс Дэшнер
217	Холли Блэк
218	Наруки Нагакава
219	Владимир Владимирович Набоков
220	Харуки Мураками
221	Анна Джейн
222	Агата Кристи
223	Халед Хоссейни
224	Агата Кристи
225	Джей Кристофф
226	Ли Мие
227	Ася Лавринович
228	Варис Дирие
229	Эрих Мария Ремарк
230	Бенджамин Алир Саэнс
231	Эмма Скотт
232	Пауло Коэльо
233	Ася Лавринович
234	Эрих Мария Ремарк
235	Ася Лавринович
236	Лоран Гунель
237	Артур Голден
238	Микита Франко
239	Уильям Сомерсет Моэм
240	Сара Дж. Маас
\.


--
-- Data for Name: book_comments; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.book_comments (id, book_id, user_id, comment, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: book_ratings; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.book_ratings (id, book_id, user_id, rating) FROM stdin;
1	3	8	3
\.


--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.books (id, title, pub_date, edition, language, rating, image, description) FROM stdin;
4	Влюблённая ведьма. Восхитительная ведьма. Комплект из 2 книг	2024-04-08	1	рус	0.00	https://s.f.kz/prod/3167/3166753_1000.jpg	Восхитительная ведьма. Книга первая\n\nПожелайте мне удачи. Я покоряю человека с железным сердцем, каменными плечами и ужасным характером. Я не знаю, чего мне хочется больше — придушить или обнять его. Он зануда, который старается все делать правильно, но совершает такие безбашенные поступки, на которые я никогда не решилась бы. Он университетский преподаватель, а я старшекурсница. Пока он не знает, что станет моим парнем. Но ведь это дело времени, правда? Я добьюсь его во что бы то ни стало.\n\nВлюбленная ведьма. Книга вторая\n\nСоскучились? Студентка Таня Ведьмина может влюбить в себя абсолютно любого мужчину. Даже если он — ее преподаватель в университете. Она покорила сердце Олега Владыко несмотря на его яростное сопротивление. Он притворился ее парнем и сам не заметил как влюбился! У него самые нежные губы и самые ласковые руки...а еще ужасный характер. Таню тянет к нему, но что-то не дает им быть вместе: гордость, чужая зависть, ложь, месть...Опасность грозит не просто их отношениям, но жизни Олега. И Таня должна спасти свою любовь. Любой ценой. Влюбленная ведьма способна на все — так гласит закон Тани Ведьминой.
5	Дюна	2024-04-09	1	рус	0.00	https://s.f.kz/prod/3141/3140696_1000.jpg	Фрэнк Герберт (1920-1986) успел написать много, но в истории остался прежде всего как автор эпопеи «Дюна». Возможно, самой прославленной фантастической саги двадцатого столетия, саги, переведенной на десятки языков и завоевавшей по всему миру миллионы поклонников.\n\nСамый авторитетный журнал научной фантастики «Локус» признал «Дюну», первый роман эпопеи о песчаной планете, лучшим научно-фантастическим романом всех времен и народов. В «Дюне» Фрэнку Герберту удалось совершить невозможное — создать своеобразную «хронику далекого будущего». И не было за всю историю мировой фантастики картины грядущего более яркой, более зримой, более мощной и оригинальной.
6	Дневник памяти	2024-04-10	1	рус	0.00	https://s.f.kz/prod/617/616016_550.jpg	Это — не «любовный роман», а роман о любви. О любви обычных мужчины и женщины — таких, как мы…\n\nПочему же эта книга стала абсолютным бестселлером во всем мире?\n\nПочему она трогает душу читателей самого разного возраста и интеллектуального уровня?\n\nКак Николасу Спарксу удалось повторить сенсационный успех «Истории любви» и «Неспящих в Сиэтле»?\n\nПочему фильм, снятый по роману «Дневник памяти», имел огромный успех во всем мире? Объяснить это невозможно.\n\nПрочитайте «Дневник памяти» — и поймете сами!
7	Десять негритят	2024-04-11	1	рус	0.00	https://s.f.kz/prod/1454/1453202_1000.jpg	Десять никак не связанных между собой людей в особняке на уединенном острове... Кто вызвал их сюда таинственным приглашением? Зачем кто-то убивает их, одного за другим, самыми невероятными способами? Почему все происходящее так тесно переплетено с веселым детским стишком?
8	Звезда Черного дракона	2024-04-12	1	рус	0.00	https://s.f.kz/prod/3029/3028989_1000.jpg	«Звезда Черного дракона» — это продолжение романа «Тайна Черного дракона» и заключительная часть цикла «Нежеланная невеста» Анны Джейн.\n\nВ детстве я мечтала попасть на императорскую свадьбу, а теперь мне суждено сыграть на ней роль невесты. К моей свадьбе готовится вся империя. Чем ближе день торжественного ритуала, тем чернее сгущаются тучи в Небесном дворце...\nС севера на нас надвигается древнее зло, и говорят, что лишь брак двух драконов сможет спасти мир. Но какой будет цена за победу над злом?\n\n#любовь #предназначение #магия — все это в романе «Звезда Черного дракона» Анны Джейн.
15	Лисья нора. Книга 1: роман	2024-04-19	1	рус	0.00	https://s.f.kz/prod/1701/1700069_1000.jpg	Сенсационная трилогия писательницы Норы Сакавич «Все ради игры» была впервые опубликована в интернете и молниеносно покорила читателей во всем мире. «Лисья нора» повествует о команде «Лисов» — игроков экси (вымышленный спорт), которые, будучи отбросами в жизни, пытаются подняться со дна турнирной таблицы и выиграть чемпионат страны. Главный герой, Нил Джостен, скрывается от своего темного прошлого, однако, став частью команды, вынужден сражаться не только с соперниками, но и с новоиспеченными товарищами, каждый из которых хранит свои секреты.
9	ЛюбовьНенависть. НенавистьЛюбовь. Комплект из 2 книг	2024-04-13	1	рус	0.00	https://s.f.kz/prod/3167/3166765_1000.jpg	#Любовьненависть. Книга первая\n\nВ детстве мы были неразлучны. Сладкая парочка #ДашаДаня сидели за одной партой, дрались, ставили друг другу подножки. Но мы всегда мирились, и он даже хотел на мне жениться. Потом мы повзрослели. Заклятый друг превратился в лучшего врага. Мы оба заигрались в ненависть и уже не можем остановиться. Но разве у #ЛюбвиНенависти бывает конец? Где бы я ни оказалась, судьба постоянно сталкивает меня с Даней, давая нам шанс все изменить.\n\n#Ненавистьлюбовь. Книга вторая\n\nМоя любовь обернулась ненавистью. Даня предал меня. У него есть другая, у меня — другой. Наша вселенная для двоих разрушилась в один день. Я решила быть счастливой ему назло. Теперь где бы он ни появлялся, я делаю все, чтобы напомнить, как подло он поступил со мной. Но кажется, Даня не забывал. Почему-то он по-прежнему рядом, когда это так нужно. #НенавистьЛюбовь пылает в наших сердцах. И я не знаю, что победит.
10	В конце они оба умрут	2024-04-14	1	рус	0.00	https://s.f.kz/prod/2077/2076336_1000.jpg	Однажды ночью сотрудники «Отдела Смерти» звонят Матео Торресу и Руфусу Эметерио, чтобы сообщить им плохие новости: сегодня они умрут. Матео и Руфус не знакомы, но оба по разным причинам ищут себе друга, с которым проведут Последний День. К счастью, специально для этого есть приложение «Последний друг», которое помогает им встретиться и вместе прожить целую жизнь за один день. Все это происходит в книге «В конце они оба умрут».
11	Маленькая жизнь	2024-04-15	1	рус	0.00	https://s.f.kz/prod/776/775860_1000.jpg	Американская писательница Ханья Янагихара создала необычный роман, где и о страшном, и о радостном говорится без лишнего надрыва и сентиментальности. Четверо друзей — талантливый архитектор Малкольм, начинающий актер Виллем, уверенный в собственной неповторимости художник Джей-Би и гениальный юрист и математик Джуд — пытаются добиться успеха в Нью-Йорке, но оказывается, что ни карьера, ни деньги, ни слава не могут справиться с прошлым, если оно сильнее жизни...
12	Свита короля	2024-04-16	1	рус	0.00	https://s.f.kz/prod/2054/2053554_1000.jpg	«Свита короля» — продолжение бестселлеров Норы Сакавич «Лисья нора» и «Король Воронов» и третья часть сенсационной трилогии «Все ради игры».\nВремя на исходе. Оказавшись в Университете Пальметто, Нил Джостен знал, что не доживет до конца года, но теперь, когда смерть не за горами, он больше чем прежде хочет жить. Дружба с Лисами была опрометчивой затеей, а поцелуй с одним из них — затеей немыслимой. Пока «Лисы» пытаются во что бы то ни стало выйти в финал чемпионата, Нил сражается за свою жизнь, ведь теперь ей угрожает не только Рико Морияма, но и мафиозный клан Балтиморского Мясника. Правда — единственный шанс Нила на спасение, однако она может привести к гибели всех, кто ему дорог...
13	Триумфальная арка	2024-04-17	1	рус	0.00	https://s.f.kz/prod/902/901367_1000.jpg	«Триумфальная арка» — пронзительная история любви всему наперекор, любви, приносящей боль, но и дарующей бесконечную радость.Место действия — Париж накануне Второй мировой войны. Герой — беженец из Германии, без документов, скрывающийся и от французов, и от нацистов, хирург, спасающий человеческие жизни. Героиня — итальянская актриса, окруженная поклонниками, вспыльчивая, как все артисты, прекрасная и неотразимая.И время, когда влюбленным довелось встретиться, и город, пронизанный ощущением надвигающейся катастрофы, становятся героями этого романа.«Триумфальная арка» была дважды экранизирована и по-прежнему покоряет читателей всего мира.
14	Лучшее во мне	2024-04-18	1	рус	0.00	https://s.f.kz/prod/2769/2768566_1000.jpg	Каждому хочется верить: настоящая любовь бессмертна.\nКаждому хочется надеяться: истинное чувство можно пронести сквозь годы и испытания...\nДоусон Коул и первая красавица школы Аманда полюбили друг друга, — однако жизнь развела их.\nПрошло много лет. Аманда стала женой другого, у нее семья, дом, дети...\nНо случай приводит ее в родной городок и дарит новую встречу с Доусоном.\nИх любовь вспыхивает вновь, — и Аманда, и Доусон понимают, что расставание было трагической ошибкой.\nНеужели, им представился шанс начать все сначала? Или у судьбы свои планы?
35	Гипотеза любви	2024-05-09	1	рус	0.00	https://s.f.kz/prod/2511/2510124_1000.jpg	Оливия — аспирантка-биолог, которая не верит в любовь. Пытаясь убедить свою лучшую подругу в том, что у нее есть отношения, она целует первого, кто попадается ей на глаза. К удивлению Оливии, симпатичный преподаватель Адам решает ей подыграть.
16	Шах и мат	2024-04-20	1	рус	0.00	https://s.f.kz/prod/3085/3084856_1000.jpg	Мэллори Гринлиф бросила шахматы, когда они разрушили ее семью. Но четыре года спустя она все же соглашается принять участие в благотворительном турнире и случайно обыгрывает Нолана Сойера — действующего чемпиона и настоящего «бэд боя» в мире шахмат.\n\nПроигрыш Нолана новичку шокирует всех, а громкая победа открывает Мэллори двери к денежному призу, который так ей необходим. Мэллори сомневается, стоит ли идти на сделку с совестью, но Сойер серьезно намерен отыграться и не желает выпускать загадочную соперницу из поля зрения.\n\nВступая в большую игру, Мэллори изо всех сил старается оградить свою семью от того, что когда-то ее разрушило, и заново не влюбиться в шахматы. Однако вскоре понимает, что не может перестать думать о Сойере, ведь он не только умен, но и чертовски привлекателен...
17	Преследуя Аделин	2024-04-21	1	рус	0.00	https://s.f.kz/prod/3007/3006578_1000.jpg	Манипулятор\nЯ могу манипулировать эмоциями любого, кто мне это позволит. Я заставлю вас страдать, заставлю вас плакать, заставлю вас смеяться и вздыхать. Но мои слова на него не действуют. Особенно тогда, когда я умоляю его уйти.\nОн всегда рядом, наблюдает и ждет. И я просто не могу отвести взгляд. Тогда, когда хочу, чтобы он был ближе ко мне.\n\nТень\nЯ не планировал влюбляться. Но теперь я не могу смотреть по сторонам. Я загипнотизирован ее улыбкой, ее глазами и тем, как она двигается. Как она раздевается...\nЯ буду продолжать наблюдать и ждать. Пока не смогу сделать ее своей женой. И как только она согласится, я никогда ее не отпущу. Даже тогда, когда она будет умолять меня об этом.
18	О чём молчит Ласточка	2024-04-22	1	рус	0.00	https://s.f.kz/prod/2486/2485955_1000.jpg	«Что бы ни случилось, не потеряйте друг друга. Что бы ни случилось, не потеряйте себя», — повторяли они в далекой юности. Не сбылось.\n\nНо спустя двадцать лет Володя и Юра встретились снова. Возможно ли построить будущее на руинах давно забытого прошлого? Или лучше позволить ему умереть, сделав по-настоящему ценным?..\n\n«О чем молчит Ласточка» — долгожданное продолжение бестселлера «Лето в пионерском галстуке».
19	Дюна	2024-04-23	1	рус	0.00	https://s.f.kz/prod/2517/2516982_1000.jpg	Фрэнк Герберт успел написать много, но в истории остался прежде всего как автор эпопеи «Дюна». Возможно, самой прославленной фантастической саги двадцатого столетия, саги, переведенной на десятки языков и завоевавшей по всему миру миллионы поклонников. Самый авторитетный журнал научной фантастики «Локус» признал «Дюну», первый роман эпопеи о песчаной планете, лучшим научно-фантастическим романом всех времен и народов. В «Дюне» Фрэнку Герберту удалось совершить невозможное — создать своеобразную «хронику далекого будущего». И не было за всю историю мировой фантастики картины грядущего более яркой, более зримой, более мощной и оригинальной.\n\nПесчаная планета Арракис, Дюна, — единственный на всю Вселенной источник «пряности». Тот, кто контролирует «пряность», контролирует саму Вселенную, ведь без неё немыслимы сами межзвёздные перелёты. Последние десятилетия правами на добычу пряности владел Великий дом Харконненов. Однако недавно падишах-император Шаддам IV, правитель всей известной человечеству части вселенной, отобрал у Харконненов концессию и передал её их злейшим врагам — дому Атрейдесов. Лето Атрейдес, мудрый и справедливый глава дома Атрейдесов, понимает, что коварные Харконнены пойдут на всё, чтобы вернуть себе главный источник дохода. Но он и не подозревает, что вся эта затея — лишь часть хитрого плана по уничтожению дома Атрейдесов. Замыслил этот план вовсе не Харконнен, и ловушка уже вот-вот готова захлопнуться.
20	Джейн Эйр	2024-04-24	1	рус	0.00	https://s.f.kz/prod/739/738171_550.jpg	Даже не читая этого романа, вы наверняка слышали про Джейн Эйр. Скорее всего, вы смотрели фильм по этому роману. Надо сказать, что экранизаций «Джейн Эйр» очень много, одна из них — с Шарлоттой Генсбур в главной роли (режиссер — Франко Дзефирелли). Джейн Эйр давным-давно переросла свою героиню и стала почти что именем нарицательным. Многие девушки, наверное, мечтали бы повторить ее судьбу. Да, детство у нее было настоящим кошмаром. Да и молодость выдалась тяжелой. Пусть Джейн не так уж хороша собой, и жизнь ее вовсе не балует, но в конечном итоге она обретет свое счастье. Полный загадок, порой трагических поворотов сюжета, этот роман был и остается символом веры, надежды и любви.
21	Гордость и предубеждение	2024-04-25	1	рус	0.00	https://s.f.kz/prod/546/545921_1000.jpg	«Гордость и предубеждение» – самый популярный женский роман в мире, провозглашенный интернет-пользователями Великобритании одной из лучших книг всех времен и народов.\nМистер Дарси – главный герой романа – стал для многих читательниц эталоном мужчины, благородный аристократ, который закрывает глаза на сословные предрассудки и женится по любви на женщине, стоящей гораздо ниже его по положению. На Элизабет Беннет, гордой, неприступной девушке, умной, начитанной и глубоко чувствующей.\nНесколько экранизаций, два сериала и армия поклонников по всему миру навеки вписали роман «Гордость и предубеждение» в летопись лучших историй о любви, побеждающей любые преграды.
22	Благословение небожителей. Том 6	2024-04-26	1	рус	0.00	https://s.f.kz/prod/3123/3122269_1000.jpg	Сказать по правде, в этом мире нет богов.\n\nДве тысячи лет небесные чертоги стояли незыблемо, но любое пиршество рано или поздно кончается. Когда боги узнают страшную тайну, они бегут из столицы бессмертных, охваченной пламенем войны. Чтобы спасти три мира от существа, что сильнее всех ныне живущих, Се Ляню с Хуа Чэном придётся заручиться поддержкой не только старых друзей, но и старых врагов.\n\nВсё решится на мосту, что прежде вёл в Небеса, а теперь обрывается над огненной пропастью. Смогут ли герои, даже объединив силы, одолеть противника? И какова будет цена победы?
23	Убийство в «Восточном экспрессе»	2024-04-27	1	рус	0.00	https://s.f.kz/prod/116/115788_1000.jpg	Находившийся в Стамбуле великий сыщик Эркюль Пуаро возвращается в Англию на знаменитом «Восточном экспрессе», в котором вместе с ним едут, кажется, представители всех возможных национальностей. Один из пассажиров, неприятный американец по фамилии Рэтчетт, предлагает Пуаро стать своим телохранителем, поскольку считает, что его должны убить.\nЗнаменитый бельгиец отмахивается от этой абсурдной просьбы. А на следующий день американца находят мертвым в своем купе, причем двери закрыты, а окно открыто. Пуаро немедленно берется за расследование — и выясняет, что купе полно всевозможных улик, указывающих... практически на всех пассажиров «Восточного экспресса». Вдобавок поезд наглухо застревает в снежных заносах в безлюдном месте. Пуаро необходимо найти убийцу до того, как экспресс продолжит свой путь...
24	Король Воронов	2024-04-28	1	рус	0.00	https://s.f.kz/prod/1845/1844555_1000.jpg	«Король Воронов» — продолжение бестселлера Норы Сакавич «Лисья нора» и вторая часть сенсационной трилогии «Все ради игры». Смерть товарища, потрясшая «Лисов» вскоре после начала игрового сезона, помогает команде наконец сплотиться. Проблема лишь в том, что Эндрю по-прежнему ни во что не ставит старшекурсников, а без этого победа над главным соперником «Лисов» — «Воронами» — невозможна. Единственным, кто может достучаться до Эндрю, становится Нил, только вот Эндрю никогда не делает ничего бесплатно, а Нил не доверяет никому, кроме себя. Едва контакт между ними налаживается, как на горизонте вновь появляется Рико, намеренный уничтожить новую жизнь Нила, а заодно и всю его команду.
25	Большая четверка	2024-04-29	1	рус	0.00	https://s.f.kz/prod/1797/1796162_1000.jpg	В дом Эркюля Пуаро вламывается изможденный человек в выпачканном костюме. Он едва живой. Несколько капель коньяка приводят его в чувство. Человек пытается что-то объяснить, но сил его хватает лишь на то, чтобы несколько раз написать на листе бумаги большую цифру «4»...
26	Маленькие женщины	2024-04-30	1	рус	0.00	https://s.f.kz/prod/2759/2758814_1000.jpg	Роман «Маленькие женщины» повествует о судьбах четырех сестер Марч — добросердечной красавицы Мэг, «сорванца в юбке» Джо, кроткой и нежной Бет и талантливой фантазерки Эми.\nГражданская война идет далеко на Юге, но ее грозные отголоски достигли дружной семьи небогатого провинциального пастора Марча — сам он ушел на фронт полковым священником, и жена с дочерями, как и множество других американок, день за днем живут в страхе трагических известий. Однако война не в силах помешать девочкам взрослеть и превращаться в юных девушек. Девушек с их обычными и трогательными девичьими мечтами и надеждами, радостями, горестями и, конечно, первой любовью...\nКнига издается в новом переводе.
27	Наследница черного дракона	2024-05-01	1	рус	0.00	https://s.f.kz/prod/2177/2176214_1000.jpg	Долгожданная новая книга от Анны Джейн, автора дилогии «Восхитительная ведьма», «Наследница черного дракона» в жанре романтического фэнтези.\nГоворят, что в праздничную ночь исполняются любые желания, стоит лишь правильно попросить духов! Я попросила настоящую любовь... Бойтесь своих желаний, ведь от любви до ненависти всего один шаг.\nВ Ночь зимнего свершения я стала невестой наследного принца Вечной империи. Мы учимся в одной академии магии и ненавидим друг друга, но связаны брачной клятвой, которую нельзя разорвать.\nИграть ли нам влюбленную пару или попытаться найти способ разрушить клятву? День свадьбы все ближе. А чувства... ярче.\nАнна Джейн «Наследница черного дракона» — это история любви, разгоревшейся как пламя из искры.
28	Портрет Дориана Грея	2024-05-02	1	рус	0.00	https://s.f.kz/prod/457/456662_1000.jpg	"Портрет Дориана Грея" — самое знаменитое произведение Оскара Уайльда, единственный его роман, вызвавший в свое время шквал негативных оценок и тем не менее имевший невероятный успех. Главный герой романа, красавец Дориан, — фигура двойственная, неоднозначная. Тонкий эстет и романтик становится безжалостным преступником. Попытка сохранить свою необычайную красоту и молодость оборачивается провалом. Вместо героя стареет его портрет — но это не может продолжаться вечно, и смерть Дориана расставляет все по своим местам. Роман Оскара Уайльда продолжает быть очень актуальным и сегодня — разве погоня за вечной молодостью порой не оборачивается потерей своего истинного лица?
29	По осколкам твоего сердца. Твоё сердце будет разбито. Комплект из 2 книг	2024-05-03	1	рус	0.00	https://s.f.kz/prod/3167/3166818_1000.jpg	Твое сердце будет разбито\n\nНовый роман Анны Джейн — история про хрупкую первую любовь между двумя одинокими подростками. Я вижу его каждый день в окне дома напротив. Он — самый крутой парень в моей новой школе. Красавчик, от которого все без ума. Опасный одиночка, который меня не замечает. Однажды он спас меня от одноклассниц, решивших превратить мою жизнь в ад. Он сказал всем, что я его девушка, и украл мой первый поцелуй при всех! Теперь никто не смеет меня тронуть, потому что никто не рискнет с ним связываться. Взамен я должна делать все, что он скажет. Я терпеть его не могу, но вынуждена подчиниться, если хочу закончить школу без проблем. Чье сердце будет разбито? Мое или его?\n\nВ новой романтической истории от автора бестселлеров Анны Джейн есть все, что так любят молодые читатели: школьные будни, противостояние характеров и первая любовь, без которой невозможно взросление.\n\nПо осколкам твоего сердца\n\nПродолжение дилогии, разбившей сердца — история про хрупкую первую любовь между двумя одинокими подростками. В новой романтической истории от автора бестселлеров Анны Джейн есть все, что так любят молодые читатели: школьные будни, противостояние характеров и первая любовь, без которой невозможно взросление.\n\nОн был моей первой любовью. Плохой парень, которого все боялись. Опасный красавчик, от которого были без ума все девчонки. Он спас меня от ада, когда весь класс издевался надо мной. Никто не смел меня больше трогать. Кто пойдет против самого крутого парня в школе? Я влюбилась в него. И время, проведенное рядом с ним, было лучшим. Но однажды его не стало. Мое сердце разбилось на осколки. Я решила измениться. Я хочу начать жизнь заново. Но почему мне кажется, что он все еще жив?
30	Три товарища	2024-05-04	1	рус	0.00	https://s.f.kz/prod/1195/1194133_1000.jpg	Самый красивый в двадцатом столетии роман о любви... Самый увлекательный в двадцатом столетии роман о дружбе... Самый трагический и пронзительный роман о человеческих отношениях за всю историю двадцатого столетия.
31	Тысяча поцелуев, которые невозможно забыть	2024-05-05	1	рус	0.00	https://s.f.kz/prod/920/919890_1000.jpg	Один поцелуй длится мгновение.Тысяча — целую жизнь.Парень и девушка.Любовь с первого взгляда.Любовь, которую ни время, ни расстояние не сможет разрушить.Любовь, которая продлится вечно.По крайней мере, так думают они.Когда семнадцатилетний Руне Кристиансен возвращается из родной Норвегии в маленький городок в штате Джорджия, туда, где он встретил свою Поппи, в голове у него только одна мысль. Почему девушка, которая завладела не только его сердцем, но и душой, которая обещала его ждать вечно, вдруг перестала отвечать на письма и звонки?Сердце Руне разбилось два года назад, когда Поппи по какой-то причине вычеркнула его из жизни. Но когда он узнает правду, то понимает, что самая большая боль еще впереди.
32	Бойцовский клуб	2024-05-06	1	рус	0.00	https://s.f.kz/prod/2305/2304663_1000.jpg	«Бойцовский клуб» — самый знаменитый роман Чака Паланика. Все помнят фильм режиссера Дэвида Финчера с Брэдом Питтом в главной роли? Он именно по этой книге. Это роман-вызов, роман, созданный всем назло и вопреки всему, в нем описывается поколение озлобившихся людей, потерявших представление о том, что можно и чего нельзя, где добро и зло, кто они сами и кто их окружает. Сам Паланик называет свой «Бойцовский клуб» новым «Великим Гэтсби». Какие же они — эти Гэтсби конца XX века?
33	Благословение небожителей. Том 5	2024-05-07	1	рус	0.00	https://s.f.kz/prod/3005/3004791_1000.jpg	Один человек. Всего один. На самом деле этого достаточно.\n\nПутешествие к Медной Печи становится всё более непредсказуемым. По дороге к вулкану Се Лянь узнаёт шокирующую правду о Хуа Чэне, но герои даже не успевают объясниться, ведь Поднебесной грозит страшная опасность — Безликий Бай вернулся! Се Ляню предстоит разгадать тайну личности демона и отразить новую волну поветрия ликов, но перед этим он столкнётся с призраками прошлого...\n\nЧто произошло с наследным принцем после первого низвержения? Как он всё потерял и опустился на самое дно? Какие страдания выпали на его долю? Кто заставил его блуждать во тьме, а кто помог обратиться к свету?
34	Убийство на поле для гольфа	2024-05-08	1	рус	0.00	https://s.f.kz/prod/1481/1480743_1000.jpg	Эркюль Пуаро получает письмо от южноамериканского миллионера с просьбой о помощи. Но Пуаро и Гастингс не успевают спасти несчастного месье Рено — его труп находят на поле для гольфа. Пуаро вынужден погрузиться в запутанную цепь событий, происходивших задолго до расследования, чтобы выйти на убийцу миллионера.
36	Бегущий за ветром	2024-05-10	1	рус	0.00	https://s.f.kz/prod/359/358989_1000.jpg	Ошеломляющий дебютный роман, который уже называют главным романом нового века, а его автора — живым классиком. Это поразительная история о детстве, дружбе, предательстве, чувстве вины и ее искуплении. Амира и Хасана разделяла пропасть. Один принадлежал к местной аристократии, другой — к презираемому меньшинству. У одного отец был красив и важен, у другого — хром и жалок. Их история разворачивается на фоне кабульской идиллии, которая вскоре сменится грозными бурями, что подхватит мальчиков и разнесет в разные стороны, как два воздушных змея. У каждого своя судьба и своя трагедия, но они по-прежнему связаны прочнейшими узами.\n\nРоман стал одним из самых ярких явлений в мировой литературе последних лет. Нежный, тонкий, ироничный, по-хорошему сентиментальный, он напоминает живописное полотно, которое можно разглядывать бесконечно.
37	Тайна черного дракона	2024-05-11	1	рус	0.00	https://s.f.kz/prod/2671/2670647_1000.jpg	Меня нарекли невестой наследного принца, которого я ненавидела. Я пыталась разорвать брачную клятву, но не заметила как влюбилась. Меня назвали драконом, последней из рода Черного дракона. Я мечтала о счастье, но обрела силу. Как принять, что любовь — это тьма, от которой звезды сияют еще ярче?
38	Бақытсыз Жамал	2024-05-12	1	рус	0.00	https://s.f.kz/prod/3103/3102908_1000.jpg	Көрнекті қоғам қайраткері, ағартушы, ақын, жазушы Міржақып Дулатовтың шығармашылығы қалың оқырманға кеңінен таныс. Қаламгер өзі өмір сүрген қоғамның өзекті мәселелерін, халықтың әлеуметтік жағдайы мен психологиялық күйін шебер суреттейді. Қай шығармасы да сол заманның шынайы келбетін көрсетеді. Бұл жинаққа «Бақытсыз Жамал» романы, әңгімелері мен өлеңдері еніп отыр. Кітап көпшілік оқырман қауымға арналған.
39	Грозовой перевал	2024-05-13	1	рус	0.00	https://s.f.kz/prod/1771/1770296_1000.jpg	«Грозовой перевал» Эмили Бронте — не просто золотая классика мировой литературы, но роман, перевернувший в свое время представления о романтической прозе. Проходят годы и десятилетия, но история бурной, страстной, трагической любви Хитклифа и Кэти по-прежнему не поддается ходу времени. «Грозовым перевалом» зачитывалось уже много поколений женщин — продолжают зачитываться и сейчас. Эта книга не стареет, как не стареет истинная любовь...
40	Жизнь взаймы, или У неба любимчиков нет	2024-05-14	1	рус	0.00	https://s.f.kz/prod/1243/1242031_1000.jpg	Ранее роман публиковался под названием «Жизнь взаймы» в сокращенном журнальном варианте 1959 года. В данном издании публикуется окончательный книжный вариант 1961 года.\nЭту жизнь герои отвоевывают у смерти! Когда терять уже нечего, когда один стоит на краю гибели, так и не узнав жизни, а другому она стала невыносима. И как всегда у Ремарка, только любовь и дружба остаются незыблемыми. Только в них можно найти точку опоры...\nВ 1977 году по книге был снят фильм с легендарным Аль Пачино.
41	Великий Гэтсби	2024-05-15	1	рус	0.00	https://s.f.kz/prod/514/513476_550.jpg	"Бурные" двадцатые годы прошлого столетия…\nВремя шикарных вечеринок, "сухого закона" и "легких" денег…\nЭти "новые американцы" уверены, что расцвет будет вечным, что достигнув вершин власти и богатства, они обретут и личное счастье…\nТаким был и Джей Гэтсби, ставший жертвой бессмысленной погони за пленительной мечтой об истинной и вечной любви, которой не суждено было сбыться…\nПеред вами - самый знаменитый роман Ф. С. Фицджеральда в новом переводе!
42	Смерть в облаках	2024-05-16	1	рус	0.00	https://s.f.kz/prod/722/721436_1000.jpg	На борту авиарейса Париж—Кройдон спит спокойном сном знаменитый сыщик Эркюль Пуаро, даже не подозревая, что происходит буквально рядом с ним. Проснувшись, он узнает, что его соседка по самолету, пожилая француженка мадам Жизель, обнаружена мертвой. Сначала ее смерть списывают на укус осы, но Пуаро выясняет, что женщина погибла от укола отравленным дротиком. По заключению врача, смерть наступила в результате отравления редчайшим ядом. На борту самолета находятся лишь десять пассажиров и два стюарда — и один из них убийца. Под подозрением находится и сам Пуаро — крайне непривычная роль для известного сыщика...
43	Убить пересмешника…	2024-05-17	1	рус	0.00	https://s.f.kz/prod/1083/1082113_1000.jpg	Харпер Ли — «гений одной книги», роман «Убить пересмешника» — ее единственное известное произведение. Но за эту книгу, переведенную едва ли не на все языки мира, писательница была удостоена Пулитцеровской премии. Книга была признана лучшим американским романом ХХ века по версии «Library Journal», а затем принесла автору высшую гражданскую награду США — медаль Свободы. Ее суммарный тираж только в Штатах составил более тридцати миллионов экземпляров! История маленького сонного городка на юге Америки, поведанная маленькой девочкой. История ее брата Джима, друга Дилла и ее отца — честного, принципиального адвоката Аттикуса Финча, одного из последних и лучших представителей старой «южной аристократии». История судебного процесса по делу чернокожего парня, обвиненного в насилии над белой девушкой. Но прежде всего — история переломной эпохи, когда ксенофобия, расизм, нетерпимость и ханжество, присущие американскому югу, постепенно уходят в прошлое. «Ветер перемен» только-только повеял над Америкой. Что он принесет?..
44	Ночь в Лиссабоне	2024-05-18	1	рус	0.00	https://s.f.kz/prod/460/459662_1000.jpg	"Ночь в Лиссабоне" (1962) - трагический, проникновенный роман Эриха Марии Ремарка о Второй мировой войне.\nЭто не только одна ночь в Лиссабоне, в которую и уместился весь этот рассказ. Это не просто случайная встреча двух отчаявшихся людей, один из которых тщетно пытается найти два билета на пароход до Америки, а другой - ищет собеседника, чтобы излить ему душу. Это настоящая исповедь отважного, смелого человека, на долю которого выпали немыслимые по тяжести испытания. Это история целого поколения людей, столкнувшихся с войной, попавших в тиски фашизма. Это еще и история любви, перед которой отступает даже смерть.
45	Особое мясо	2024-05-19	1	рус	0.00	https://s.f.kz/prod/1838/1837087_550.jpg	Внезапное появление смертоносного вируса, поражающего животных, стремительно меняет облик мира. Все они — от домашних питомцев до диких зверей — подлежат немедленному уничтожению с целью нераспространения заразы. Употреблять их мясо в пищу категорически запрещено. В этой чрезвычайной ситуации, грозящей массовым голодом, правительства разных стран приходят к радикальному решению: легализовать разведение, размножение, убой и переработку человеческой плоти. Узаконенный каннибализм разделяет общество на две группы: тех, кто ест, и тех, кого съедят.
46	Незнакомка из Уайлдфелл-Холла	2024-05-20	1	рус	0.00	https://s.f.kz/prod/2693/2692529_1000.jpg	Одно из лучших произведений «золотого века» английской литературы.\nМногократно экранизированный шедевр психологического реализма.\nРоман, который был впервые опубликован в 1848 году, — и произвел в Англии сенсацию, поскольку в нем, впервые в европейской литературе, со всей откровенностью и прямотой, задавался очень неудобный вопрос: должна ли женщина, ставшая женой домашнего тирана, покорно нести свой крест «ради сохранения семьи», во имя детей? А может, разорвать узы такого брака — не только ее право, но и долг именно как матери?\nВопрос, который, увы, для многих женщин во всем мире сохраняет актуальность и сейчас.
47	Жестокий принц	2024-05-21	1	рус	0.00	https://s.f.kz/prod/1093/1092818_1000.jpg	Наточи свой клинок, ожесточи сердце! Игра началась... Разумеется я хочу во всем походить на них. Они прекрасны, как мечи, выкованные в божественном огне. Они будут жить вечно. А принц Кардан прекраснее всех. Но как же я его ненавижу! Когда я смотрю на него, то задыхаюсь от ненависти! Джуд было семь, когда ее родителей убили, а девочку вместе с сестрами забрали к себе фейри. Прошло десять лет и все, чего она хочет, так это походить на прекрасных, но коварных созданий, которые ее воспитали. И это несмотря на то, что фейри презирают людей. А особенно принц Кардан, самый младший и самый жестокий из сыновей Верховного короля. Чтобы занять свое место при дворе, Джуд должна лгать, изворачиваться, драться, строить заговоры, и... победить Кардана. Но когда королевство фейри из-за неуемной жажды власти и кровопролития принцев оказывается на краю гибели, Джуд понимает, что это ее шанс. Она может подняться на недосягаемую ранее высоту, спасти сестер и собственную жизнь.
48	Атлант расправил плечи. В 3-х книгах	2024-05-22	1	рус	0.00	https://s.f.kz/prod/312/311141_1000.jpg	Клянусь своей жизнью и любовью к ней, что никогда не буду жить для кого-то другого и никогда не попрошу кого-то другого жить для меня.\nАйн Рэнд\n\nК власти в США приходят социалисты и правительство берет курс на «равные возможности», считая справедливым за счет талантливых и состоятельных сделать богатыми никчемных и бесталанных. Гонения на бизнес приводят к разрушению экономики, к тому же один за другим при загадочных обстоятельствах начинают исчезать талантливые люди и лучшие предприниматели. Главные герои романа стальной король Хэнк Риарден и вице-президент железнодорожной компании Дагни Таггерт тщетно пытаются противостоять трагическим событиям. Вместо всеобщего процветания общество погружается в апатию и хаос.\n\nЧасть 1. Непротиворечие. В первом томе читатели знакомятся с главными героями, гениальными предпринимателями, которым противостоят их антиподы — бездарные государственные чиновники. Повествование начинается с вопроса: «Кто такой Джон Голт?» — и на этот вопрос будут искать ответ герои романа и его читатели. Перевод с английского Ю. Соколова.\n\nЧасть 2. Или — или. Вторая часть романа — социальный прогноз. В ситуации, когда правительство берет курс на «равные возможности», считая справедливым за счет талантливых и состоятельных сделать богатыми никчемных и бесталанных, проигравшими оказываются все. Запрет на развитие производства и лоббирование интересов «нужных» людей разрушают общества. Динамика повествования задается сложным переплетением судеб героев, любовными коллизиями и загадкой, кто же такой Джон Голт. Перевод с английского В. Вебера.\n\nЧасть 3. А есть А. Третья часть романа развенчивает заблуждения мечтательных борцов за равенство и братство. Государственные чиновники, лицемерно призывающие граждан к самопожертвованию, но ограничивающие свободу предпринимательства, приводят страну к экономическому краху. Сюжет сплетается из финансовых и политических интриг, и одновременно звучит гимн новой этике: капиталистическая система ценностей не только социально оправданна, но и нравственна. Герой нового мира, гениальный изобретать Джон Голт, провозглашает принцип «нравственности разумного эгоизма» одной фразой: «Я никогда не буду жить ради другого человека и никогда не попрошу другого человека жить ради меня»
49	Жизнь на продажу	2024-05-23	1	рус	0.00	https://s.f.kz/prod/2280/2279890_1000.jpg	Юкио Мисима — самый знаменитый и читаемый в мире японский писатель. Прославился он в равной степени как своими произведениями во всех мыслимых жанрах (романы, пьесы, рассказы, эссе), так и экстравагантным стилем жизни и смерти (харакири после неудачной попытки монархического переворота).\n\nВ романе «Жизнь на продажу» молодой служащий рекламной фирмы Ханио Ямада после неудачной попытки самоубийства помещает в газете объявление: «Продам жизнь. Можете использовать меня по своему усмотрению. Конфиденциальность гарантирована». И кто только к нему не обращается! Среди его клиентов ревнивый муж, наследница-нимфоманка, разведслужба посольства, неспособная самостоятельно решить загадку отравленной моркови, и даже натуральный вампир. И вот, вместо того чтобы тихо-мирно свести счеты с жизнью, Ханио Ямада оказывается в центре заговора глобального масштаба...\n\n«Блестящий пример бескрайнего воображения Мисимы на пике формы. Парадоксальные идеи о природе бытия изложены с фирменной иронической усмешкой» (The Japan Times).
50	Удержать 13-го	2024-05-24	1	рус	0.00	https://s.f.kz/prod/3252/3251042_1000.jpg	Удержать сложнее, чем завоевать, и знают об этом не только спортсмены. Новой ученице удалось покорить главную знаменитость Томмен-колледжа — восходящую звезду регби Джонни Кавану. Но можно ли удержать рядом парня, для которого на первом месте всегда была карьера? Восстановившись после операции, Джонни вновь приступает к тренировкам. Его мечта — играть за национальную сборную. Он готовился к этому с детства, однако, кажется, все равно недооценил ответственность и нагрузки. К тому же на другой чаше весов — девушка с полуночно-синими глазами... Жизнь Шаннон Линч никогда не была спокойной, но теперь ситуация накалилась до предела. Конфликты в семье не прекращаются, а возвращение старшего брата, покинувшего родных несколько лет назад, лишь ухудшает положение. Шаннон ищет спасения у Джонни и все больше сближается с ним, однако она понимает: скоро важные игры, а значит, неизбежно расставание...\nВпервые на русском!
51	Благословение небожителей. Том 1	2024-05-25	1	рус	0.00	https://s.f.kz/prod/2280/2279264_1000.jpg	В незапамятные времена Се Лянь был наследным принцем государства Сяньлэ. Судьба одарила его всем: прекрасным ликом, чистыми помыслами и бесконечной любовью своих подданных. И если уж кому-то и было предначертано стать Божеством, то именно Его Высочеству.Однако удержаться на Небесах оказалось для него не так просто. Се Лянь возносился дважды и дважды был изгнан на землю. И вот спустя восемьсот лет скитаний Его Высочество вновь возвращается в Небесные чертоги. Получив своё первое задание в роли Божества, он сталкивается с таинственным и невероятно могущественным демоном, который, как оказалось, уже давно положил глаз на наследного принца...
52	Театр	2024-05-26	1	рус	0.00	https://s.f.kz/prod/388/387912_1000.jpg	Самый знаменитый роман Сомерсета Моэма. Тонкая, едко-ироничная история блистательной, умной актрисы, отмечающей «кризис середины жизни» романом с красивым молодым «хищником»? «Ярмарка тщеславия» бурных двадцатых? Или — неподвластная времени увлекательнейшая книга, в которой каждый читатель находит что-то лично для себя? «Весь мир — театр, и люди в нем — актеры!» Так было — и так будет всегда!
53	Таинственный сад	2024-05-27	1	рус	0.00	https://s.f.kz/prod/1619/1618768_1000.jpg	В центре романа «Таинственный сад» — десятилетняя Мери Леннокс, вернувшаяся из Индии в Англию после смерти родителей. В поместье ее дяди Арчибальда Крэвена Мери вынуждена привыкать к совершенно другой жизни, непохожей на ту, что была у нее до сих пор. Однажды девочка узнает о загадочном заброшенном саде, в который запрещено входить, и решает разыскать его. Вместе с новыми друзьями Мери предстоит проникнуть в тайны этого удивительного места, преображающего души людей.
54	Белые пешки	2024-05-28	1	рус	0.00	https://s.f.kz/prod/2873/2872033_1000.jpg	2007 год. Москву потрясает серия жестоких убийств. Следствие заходит в тупик, ведь дело обстоит намного серьезнее, чем кажется. Давний враг, подчинивший Хаос, бросает вызов самому Мирозданию, он вернулся, чтобы забрать свое. Ради победы над общим врагом Добро и Зло готовы пойти на невозможное — объединиться. Но пешками в этой шахматной партии на древней Великой Доске суждено стать невинным людям. Так судьбы восьми друзей оказываются в руках Небесной Канцелярии. Удастся ли им восстановить мировое равновесие? И кто окажется истинным победителем в этом неравном поединке?
55	Шестерка воронов	2024-05-29	1	рус	0.00	https://s.f.kz/prod/908/907470_1000.jpg	Каз Бреккер никогда не снимает черных перчаток. Но, если не хочешь стать ужином для акул, не спрашивай его, почему. Никому не известно, где его семья, откуда он пришел и почему остался в Кеттердаме. Зато он знает обо всех и все. Бреккер — правая рука главаря одной из самых влиятельных банд в городе. Казино, бордели, нелегальная торговля — его стихия. А еще шантаж, грабеж и, если понадобится, хладнокровное убийство. Но все это мелочи по сравнению с новым заказом.\n\nНа кону — баснословные деньги и… секрет, который может уничтожить одни народы и возвеличить другие. Какие именно — теперь зависит от Каза и его команды. Шестерых "воронов", которым нечего терять кроме надежды. Это дело объединит их. Лучшего стрелка банды Отбросов и новичка, который не умеет держать пистолет в руках. Соблазнительную чародейку, умеющую с помощью магии взрывать сердца, и безжалостного охотника на таких, как она. Юную гимнастку из самого известного публичного дома во всей Керчии и Каза Бреккера, способного без тени сомнения вырвать глаз предателю. Им предстоит один путь, но у каждого своя цель…
56	Четвертое крыло	2024-05-30	1	рус	0.00	https://s.f.kz/prod/3006/3005106_1000.jpg	Двадцатилетняя Вайолет Сорренгейл готовилась стать писцом и спокойно жить среди книг и пыльных документов. Но ее мать — прославленный генерал, и она не потерпит слабости ни в каком виде. Поэтому Вайолет вынуждена присоединиться к сотням молодых людей, стремящихся стать элитой Наварры — всадниками на драконах. Однако из военной академии Басгиат есть только два выхода: окончить ее или умереть. Смерть ходит по пятам за каждым кадетом, потому что драконы не выбирают слабаков. Они их сжигают. Сами кадеты тоже будут убивать, чтобы повысить свои шансы на успех. Некоторые готовы прикончить Вайолет только за то, что она дочь своей матери. Например, Ксейден Риорсон — сильный и безжалостный командир крыла в квадранте всадников. Тем временем война, которую ведет Наварра, становится все более тяжелой, и совсем скоро Вайолет придется вступить в бой.\n\nКнига содержит нецензурную лексику.\nКнига с черным срезом.
57	По осколкам твоего сердца	2024-05-31	1	рус	0.00	https://s.f.kz/prod/2956/2955107_1000.jpg	Дилогия, которая разбивает сердца. Он был моей первой любовью. Плохой парень, которого все боялись. Опасный красавчик, от которого были без ума все девчонки. Он спас меня от ада, когда весь класс издевался надо мной. Никто не смел меня больше трогать.\n\nКто пойдет против самого крутого парня в школе? Я влюбилась в него. И время, проведенное рядом с ним, было лучшим. Но однажды его не стало. Мое сердце разбилось на осколки. Я решила измениться. Я хочу начать жизнь заново. Но почему мне кажется, что он все еще жив?
58	Благословение небожителей. Том 2	2024-06-01	1	рус	0.00	https://s.f.kz/prod/2362/2361139_1000.jpg	Всё в этом мире имеет счёт: и удача, и невезение.\n\nБеды преследуют Се Ляня одна за другой: вот он ненароком устроил поджог, а вот столкнулся с призраками прошлого, о которых предпочёл бы забыть...\n\nВосемьсот лет назад он был любимцем толпы, а матушка, отец и советник возлагали на него большие надежды. Слава о его подвигах стремительно достигла Небес, и в год своего семнадцатилетия Се Лянь вознёсся. Однако радость была скоротечна: в государстве Сяньлэ настали тяжёлые времена. Видя страдания своего народа, принц был не в силах остаться в стороне. Но может ли бог выбирать чью-то сторону?\n\nПервый тираж первого тома цикла напечатан колоссальным тиражом 80 000 книг. А также:\n• китайский цикл романов, который имеет просто сумасшедшее количество фанатов в России и во всем мире;\n• поддержка со стороны лидеров фандома, имеющих десятки и сотни тысяч подписчиков;\n• самый успешный азиатский автор последних лет, работающий в жанре исторической драмы;\n• фэнтезийный мир Древнего Китая, харизматичные персонажи, остроумные диалоги и сюжет, в котором сочетаются комедия и душераздирающие твисты;\n• каждый том российского издания выполнен на высшем уровне: софт-тач ламинация, ляссе, серебряное тиснение, суперобложка с уф-лаком;\n• внутри книги — эксклюзивные иллюстрации российской художницы Антейку, чьи работы оценены фандомом серии во всем мире.
59	Цвет пурпурный	2024-06-02	1	рус	0.00	https://s.f.kz/prod/1879/1878281_1000.jpg	Унижения, боль, насилие, бесправие — такова была судьба темнокожей женщины Глубокого Юга в начале прошлого века. Такова судьба главной героини романа Сили. Ей приходилось играть роль покорной служанки жестокого отца, разлучившего ее с детьми и любимой сестрой, а потом забитой жены-рабыни сурового мужа...\nНо однажды в жизни Сили появляется наставница и настоящая подруга, которой она не безразлична. Вместе с ней Сили найдет путь к свободе и независимости. Сделав первый шаг и оставив прошлое позади, она поймет свое призвание в этом мире и окружит себя любимыми людьми...
60	Поющие в терновнике	2024-06-03	1	рус	0.00	https://s.f.kz/prod/373/372916_1000.jpg	«Поющие в терновнике».\n\nЛюбовный роман, поднятый на уровень настоящей литературы. Трогательная история взаимоотношений влюбленных, завораживающая читателя своей искренностью, чистотой и глубиной…
61	Выбор	2024-06-04	1	рус	0.00	https://s.f.kz/prod/949/948676_1000.jpg	Любовь? Серьезные отношения? Ответственность? Семья? Закоренелый холостяк Тревис Паркер уверен, что все это не для него. Ему вполне достаточно отличной работы и верных друзей. Он увлекается охотой, рыбалкой, занимается экстремальными видами спорта — и избегает серьезных отношений...\nОднако с появлением новой соседки Габи Холланд жизнь Тревиса постепенно меняется. Первоначальные неприязнь и отторжение оборачиваются незнакомым до этого времени чувством — Тревис влюбляется. Он счастлив и наслаждается жизнью, но безмятежность не вечна: любовь оказалась не только величайшей наградой, но и тяжелейшим испытанием...\nЭтот роман лег в основу одноименного фильма с Терезой Палмер и Бенджамином Уокером.
62	Белые ночи	2024-06-05	1	рус	0.00	https://s.f.kz/prod/1055/1054883_1000.jpg	В этот сборник вошли две ранние повести Достоевского — «Белые ночи» и «Неточка Незванова», которые считаются самыми поэтичными произведениями великого романиста.\n\n«Белые ночи» — одно из лучших произведений школы «сентиментального натурализма», по мнению критика Аполлона Григорьева. Это лирическая исповедь героя-мечтателя, одинокого и робкого человека, в жизни которого на какое-то время появляется девушка, а вместе с ней и надежда на более светлое будущее.\n\n«Неточка Незванова» — повесть, изначально задуманная автором как роман, где в основе сюжета лежит история жизни маленькой девочки. Неточка — тоже персонаж-мечтатель, она грезит о жизни в большом красивом особняке, который видит из окна каморки на чердаке. Но, очутившись в нем, Неточка сталкивается с действительностью, которая оказалась вовсе не так прекрасна...
63	Сорок правил любви	2024-06-06	1	рус	0.00	https://s.f.kz/prod/531/530949_550.jpg	Любовь — вода жизни. Влюбленные — огонь души. Вся вселенная начинает кружиться иначе, когда огонь влюбляется в воду.\nXIII век. В маленьком городке Конья, в городке, куда с запада не дошли крестоносцы после разграбления Константинополя и куда с востока не докатились орды Чингисхана, «несколько истинно верующих» нанимают убийцу по прозвищу Шакалья Голова для устранения Шамса Тебризи, странствующего дервиша, проповедующего «сорок правил религии любви». Ведь известно, чем больше человек говорит о любви, тем сильнее его ненавидят...\nНаши дни. США. Элла Рубинштейн, работающая в литературном агентстве, получает на рецензию рукопись «Сладостное богохульство», действие которой происходит в XIII веке. Роман настолько захватывает Эллу, что она начинает подозревать, что автора непостижимым образом вдохновил герой романа Шамс из Тебриза. И вот любовь к автору книги врывается в ее сердце, полностью переворачивая привычную и такую милую ей жизнь…
64	Посторонний	2024-06-07	1	рус	0.00	https://s.f.kz/prod/2089/2088054_1000.jpg	«Посторонний» — дебютная работа молодого писателя, своеобразный творческий манифест. Понятие абсолютной свободы — основной постулат этого манифеста. Героя этой повести судят за убийство, которое он совершил по самой глупой из всех возможных причин. И это правда, которую герой не боится бросить в лицо своим судьям, пойти наперекор всему, забыть обо всех условностях и умереть во имя своих убеждений.
65	Мартин Иден	2024-06-08	1	рус	0.00	https://s.f.kz/prod/485/484566_1000.jpg	"Мартин Иден" — самый известный роман Джека Лондона, впервые напечатанный в 1908-1909 гг. Во многом автобиографическая книга о человеке, который "сделал себя сам", выбравшись из самых низов, добился признания. Любовь к девушке из высшего общества побуждает героя заняться самообразованием. Он становится писателем, но все издательства отказывают ему в публикации. И как это часто бывает в жизни, пройдя сквозь лишения и унижения, получив отказ от любимой девушки, он наконец становится знаменитым. Но ни слава, ни деньги, ни успех, ни даже возвращение его возлюбленной не могут уберечь Мартина от разочарования в этой насквозь фальшивой жизни.
66	Мастер и Маргарита	2024-06-09	1	рус	0.00	https://s.f.kz/prod/296/295466_1000.jpg	«Мастер и Маргарита» М. А. Булгакова — самое удивительное и загадочное произведение ХХ века. Опубликованный в середине 1960-х, этот роман поразил читателей необычностью замысла, красочностью и фантастичностью действия, объединяющего героев разных эпох и культур. Автор создал «роман в романе», где сплетены воедино религиозно-историческая мистерия, восходящая к легенде о распятом Христе, московская «буффонада» и сверхъестественные сцены с персонажами, воплощающими некую темную силу, которая однако «вечно хочет зла и вечно совершает благо».\n\n«Есть в этой книге какая-то безрасчетность, какая-то предсмертная ослепительность большого таланта...» — писал Константин Симонов в своем предисловии к первой публикации романа, открывшей всему миру большого художника, подлинного Мастера слова.
67	Дом, в котором...	2024-06-10	1	рус	0.00	https://s.f.kz/prod/134/133736_1000.jpg	На окраине города, среди стандартных новостроек, стоит Серый Дом, в котором живут Сфинкс, Слепой, Лорд, Табаки, Македонский, Черный и многие другие. Неизвестно, действительно ли Лорд происходит из благородного рода драконов, но вот Слепой действительно слеп, а Сфинкс — мудр. Табаки, конечно, не шакал, хотя и любит поживиться чужим добром. Для каждого в Доме есть своя кличка, и один день в нем порой вмещает столько, сколько нам, в Наружности, не прожить и за целую жизнь. Каждого Дом принимает или отвергает. Дом хранит уйму тайн, и банальные «скелеты в шкафах» — лишь самый понятный угол того незримого мира, куда нет хода из Наружности, где перестают действовать привычные законы пространства-времени.\n\nДом — это нечто гораздо большее, чем интернат для детей, от которых отказались родители. Дом — это их отдельная вселенная.
68	Демиан	2024-06-11	1	рус	0.00	https://s.f.kz/prod/450/449431_1000.jpg	«Демиан» — философский роман, мрачный и мистический. Можно считать его и автобиографичным — об этом Гессе заявляет в предисловии. Знаковое произведение, оказавшее огромное влияние на дальнейшее творчество писателя, а великий Томас Манн сравнивал эту книгу со «Страданиями юного Вертера». Это история взросления и становления юноши, который шаг за шагом все дальше отходит от лицемерных норм общественной морали и все яснее открывает для себя глубинное, темное «я» — свободное, неподвластное царящему вокруг добродетельному фарисейству. В этом ему помогает таинственный друг Демиан — носитель «печати Каина», не то дьявол, не то загадочное божество, не то просто порождение воображения героя…
69	Злой король	2024-06-12	1	рус	0.00	https://s.f.kz/prod/1330/1329174_1000.jpg	Долгожданное продолжение международного бестселлера и бестселлера New York Times «Жестокий принц». Одна из самых ожидаемых книг 2019 года! Неповторимая история о смертной девушке, которая благодаря своему упорству, хитрости и коварству, поднялась на трон мира фейри! Джуд связала Кардана обещанием подчиняться ей, обещанием, которое продлится ровно год и один день. Теперь она главная фигура за троном, которая дергает за ниточки и умело манипулирует королем. Но Джуд ввязалась в опасную игру фейри, не имея ни друзей, ни союзников. Подстегиваемая амбициями и целью выжить во что бы то ни стало, она плетет интриги и наносит молниеносные удары. Однако когда Джуд выясняет, что среди тех, кому она безоговорочно доверяла, появился предатель, а ее близким грозит опасность, ей приходится предпринять важный шаг и возможно даже изменить все правила игры. Тем более, что Кардан оказался вовсе не таким слабым и безвольным королем, как думали все обитатели Фейриленда...
70	Тайна «Голубого поезда»	2024-06-13	1	рус	0.00	https://s.f.kz/prod/1236/1235215_1000.jpg	В роскошном экспрессе «Голубой поезд», следующем из Лондона на Французскую Ривьеру, произошла трагедия. Задушена в собственном купе дочь известного американского миллионера Рут Кеттеринг, а все ее драгоценности, в том числе изумительный рубин, исчезли.\nВ том же поезде ехал отдыхать на море и великий сыщик Эркюль Пуаро, который, разумеется, не мог оставаться в стороне от расследования. Все вокруг него только и говорят, что о загадочном похитителе драгоценностей по кличке Маркиз, и полностью уверены, что это преступление — его рук дело. Однако Пуаро получил информацию о том, что незадолго до того, как несчастную женщину обнаружили мертвой, из ее купе выходил муж Рут, Дерек...
71	Дети Дюны	2024-06-14	1	рус	0.00	https://s.f.kz/prod/1629/1628275_1000.jpg	«Дети Дюны» — третья книга культовой эпопеи Герберта. На политическом ландшафте Арракиса появляются новые игроки — наделенные сверхспособностями и даром предвидения дети Пола Атрейдеса, близнецы Лето и Гханима. Дети, лишенные детства. Дети, которым придется проявить мудрость и стойкость, чтобы выжить самим и спасти Вселенную...
72	Медвежий угол	2024-06-15	1	рус	0.00	https://s.f.kz/prod/1295/1294543_1000.jpg	Захолустный Бьорнстад — Медвежий город — затерян в северной шведской глуши: дальше только непроходимые леса. Когда-то здесь кипела жизнь, а теперь царят безработица и безысходность.\n\nПоследняя надежда жителей — местный юниорский хоккейный клуб, когда-то занявший второе место в чемпионате страны. Хоккей в Бьорнстаде — не просто спорт: вокруг него кипят нешуточные страсти, на нем завязаны все интересы, от него зависит, как сложатся судьбы.\n\nДень победы в матче четвертьфинала стал самым счастливым и для города, и для руководства клуба, и для команды, и для ее семнадцатилетнего капитана Кевина Эрдаля. Но для пятнадцатилетней Маи Эриксон и ее родителей это был страшный день, перевернувший всю их жизнь...\n\nПеред каждым жителем города встала необходимость сделать моральный выбор, ответить на вопрос: какую цену ты готов заплатить за победу?
73	Тот самый	2024-06-16	1	рус	0.00	https://s.f.kz/prod/1791/1790941_1000.jpg	Матвей прячется от внешнего мира в книгах. После семейной трагедии он еще больше замыкается в себе, а в густом воздухе старого дома, где он живет с мамой и сестрой, постоянно вспыхивают ссоры и недопонимания. Одна случайная встреча жарким летним днем навсегда меняет представления Матвея о себе, о любви и о мире.
74	По осколкам твоего сердца. Твоё сердце будет разбито. Комплект из 2 книг	2024-06-17	1	рус	0.00	https://s.f.kz/prod/3167/3166792_1000.jpg	Твое сердце будет разбито\n\nНовый роман Анны Джейн — история про хрупкую первую любовь между двумя одинокими подростками. Я вижу его каждый день в окне дома напротив. Он — самый крутой парень в моей новой школе. Красавчик, от которого все без ума. Опасный одиночка, который меня не замечает. Однажды он спас меня от одноклассниц, решивших превратить мою жизнь в ад. Он сказал всем, что я его девушка, и украл мой первый поцелуй при всех! Теперь никто не смеет меня тронуть, потому что никто не рискнет с ним связываться. Взамен я должна делать все, что он скажет. Я терпеть его не могу, но вынуждена подчиниться, если хочу закончить школу без проблем. Чье сердце будет разбито? Мое или его?\n\nВ новой романтической истории от автора бестселлеров Анны Джейн есть все, что так любят молодые читатели: школьные будни, противостояние характеров и первая любовь, без которой невозможно взросление.\n\nПо осколкам твоего сердца\n\nПродолжение дилогии, разбившей сердца — история про хрупкую первую любовь между двумя одинокими подростками. В новой романтической истории от автора бестселлеров Анны Джейн есть все, что так любят молодые читатели: школьные будни, противостояние характеров и первая любовь, без которой невозможно взросление.\n\nОн был моей первой любовью. Плохой парень, которого все боялись. Опасный красавчик, от которого были без ума все девчонки. Он спас меня от ада, когда весь класс издевался надо мной. Никто не смел меня больше трогать. Кто пойдет против самого крутого парня в школе? Я влюбилась в него. И время, проведенное рядом с ним, было лучшим. Но однажды его не стало. Мое сердце разбилось на осколки. Я решила измениться. Я хочу начать жизнь заново. Но почему мне кажется, что он все еще жив?
75	1984	2024-06-18	1	рус	0.00	https://s.f.kz/prod/2409/2408258_1000.jpg	Своеобразный антипод второй великой антиутопии XX века — «О дивный новый мир» Олдоса Хаксли. Что, в сущности, страшнее: доведенное до абсурда «общество потребления» — или доведенное до абсолюта «общество идеи»?\nПо Оруэллу, нет и не может быть ничего ужаснее тотальной несвободы...\n\nКаждый день Уинстон Смит переписывает историю в соответствии с новой линией Министерства Правды. С каждой ложью, которую он переносит на бумагу, Уинстон всё больше ненавидит Партию, которая не интересуется ничем кроме власти, и которая не терпит инакомыслия. Но чем больше Уинстон старается думать иначе, тем сложнее ему становится избежать ареста, ведь Большой Брат всегда следит за тобой...
76	Тысяча сияющих солнц	2024-06-19	1	рус	0.00	https://s.f.kz/prod/359/358994_1000.jpg	Премия «Выбор читателя» 2007 года В США и Великобритании. Абсолютный мировой бестселлер 2007 года.\n\nВ центре романа — две женщины, которые оказались жертвами потрясений, разрушивших мирный Афганистан. Мариам — незаконная дочь богатого бизнесмена, с детства познавшая, что такое несчастье, с ранних лет ощутившая собственную обреченность. Лейла — напротив, любимая дочка в дружной семье, мечтающая об интересной и прекрасной жизни. Они живут в разных мирах, которым не суждено было бы пересечься, если бы не огненный шквал войны. Отныне Лейла и Мариам связаны самыми тесными узами и сами не знают, кто они — враги, подруги или сестры. Но в одиночку им не выжить в обезумевшем мире, не выстоять перед средневековым деспотизмом и жестокостью, затопившими улицы и дома некогда уютного города.\n\nРоман рассказывает о хитросплетениях женских судеб и о том, как Большая Война влияет на обыкновенный семейный быт. Писатель достиг того, чего с трудом могли бы достичь политологи, журналисты и новостные сводки.\n\nРоман Халеда Хоссейни невообразимо трагичен и неотразимо прекрасен, как ветхозаветная история. Читатели, которых подкупил его первый роман «Бегущий за ветром», полюбят и «Тысячу сияющих солнц».
77	Гончие Лилит	2024-06-20	1	рус	0.00	https://s.f.kz/prod/1189/1188311_1000.jpg	Однажды в дублинском кафе, где работает Скай Полански, появляется загадочная незнакомка по имени Лилит, которая предлагает Скай работу секретарем в своей клинике в Бостоне. Скай не из тех, кто готов к переменам, но, устав от неудач в личной жизни, она соглашается. Вскоре выясняется, что под прикрытием клиники скрывается весьма изощренный бизнес, способный дать все: красивую жизнь, роскошь и адреналин. Но праздник длится недолго: дьявол уже приготовил для Скай свои страшные дары.\n«Гончие Лилит» — это романтический триллер о реальности, которая бывает страшнее вымысла, и о любви, которая может стать последним шансом на спасение.
78	Хрупкое равновесие. Комплект из 3 книг	2024-06-21	1	рус	0.00	https://s.f.kz/prod/3167/3166199_1000.jpg	Легендарная трилогия Аны Шерри «Хрупкое равновесие»: комплект из трех книг!\n\nДиана Оливер работает парамедиком на «Скорой». Днем она спасает людям жизни, а по вечерам ходит с друзьями по барам или стреляет в тире. Диана была уверена в своем счастливом и благополучном будущем, пока судьба не свела ее с главой итальянской мафии. Девушка еще не догадывается, но совсем скоро ей предстоит стать частью криминального мира, ведь ее талант попадать точно в цель не остался незамеченным.\nСтефано Висконти — глава клана «Morte Nera». Он привык все держать под контролем. Каждая его сделка — успешная, каждый враг — заклятый. Но для проведения крупных операций ему нужен снайпер, человек с идеальной меткостью, подчиняющийся любым его приказам, преданный ему. И на эту роль идеально подходит хрупкая девушка-парамедик, способная справиться не только с собственными чувствами, но и с любым оружием...
79	Цветы для Элджернона	2024-06-22	1	рус	0.00	https://s.f.kz/prod/1799/1798906_1000.jpg	«Цветы для Элджернона» Дэниела Киза входят в программу обязательного чтения в американских школах. Это единственная история в жанре научной фантастики, автор которой был дважды награжден, сначала за рассказ, а потом за роман с одним и тем же названием, героем, сюжетом.\nТридцатитрехлетний Чарли Гордон — умственно отсталый. При этом у него есть работа, друзья и непреодолимое желание учиться. Он соглашается принять участие в опасном научном эксперименте в надежде стать умным...\nЭта фантастическая история обладает поразительной психологической силой и заставляет задуматься над общечеловеческими вопросами нравственности: имеем ли мы право ставить друг над другом эксперименты, к каким результатам это может привести и какую цену мы готовы заплатить за то, чтобы стать «самым умным». А одиноким?\nНа вопросы, которые поднимали еще М. Булгаков в «Собачьем сердце» и Дж. Лондон в «Мартине Идене», Дэниел Киз дает свой однозначный ответ.
80	Загадочное происшествие в Стайлзе	2024-06-23	1	рус	0.00	https://s.f.kz/prod/1405/1404112_1000.jpg	«Загадочное происшествие в Стайлзе» — это самый первый роман Агаты Кристи. В этой книге, вышедшей в 1920 году, читатель в первый раз встречается с самым знаменитым сыщиком XX столетия — усатым бельгийцем Эркюлем Пуаро, а также с его другом и помощником Гастингсом. Именно здесь Пуаро впервые предоставлена возможность продемонстрировать свои дедуктивные способности и раскрыть загадочное преступление (отравление миссис Инглторп, хозяйки поместья Стайлз), опираясь на всем известные факты.
81	Стигмалион	2024-06-24	1	рус	0.00	https://s.f.kz/prod/1062/1061833_1000.jpg	Меня зовут Долорес Макбрайд, и я с рождения страдаю от очень редкой формы аллергии: прикосновения к другим людям вызывают у меня сильнейшие ожоги. Я не могу поцеловать парня, обнять родителей, выйти из дому, не надев перчатки. Я неприкасаемая. Я словно живу в заколдованном замке, который держит меня в плену и наказывает ожогами и шрамами за каждую попытку «побега». Даже придумала имя для своей тюрьмы: Стигмалион.\n\nМеня уже не приводит в отчаяние мысль, что я всю жизнь буду пленницей своего диагноза — и пленницей умру. Я не тешу себя мечтами, что от моей болезни изобретут лекарство, и не рассчитываю, что встречу человека, не оставляющего на мне ожогов...\n\nНо до чего же это живучее чувство — надежда. А вдруг я все-таки совершу побег из Стигмалиона? Вдруг и я смогу однажды познать все это: прикосновения, объятия, поцелуи, безумство, свободу, любовь?
82	Щегол	2024-06-25	1	рус	0.00	https://s.f.kz/prod/482/481939_1000.jpg	Новая книга известнейшей американской писательницы.\nДонна Тартт вошла в список журнала Time "100 самых влиятельных людей года" 2014.\nПрава на экранизацию "Щегла" были приобретены кинокомпанией Warner Brothers.\nОчнувшись после взрыва в музее, Тео Декер получает от умирающего старика кольцо и редкую картину с наказом вынести их из музея. Тео будет швырять по разным домам и семьям – от нью-йоркских меценатов до старика-краснодеревщика, от дома в Лас-Вегасе до гостиничного номера в Амстердаме, а украденная картина станет и тем проклятьем, что утянет его на самое дно, и той соломинкой, которая поможет его выбраться к свету.
83	Благословение небожителей. Том 3	2024-06-26	1	рус	0.00	https://s.f.kz/prod/2648/2647263_1000.jpg	Для меня едино, окружён ты сиянием или повержен в грязь. Значение имеешь только ты сам.\nБлизится Праздник середины осени, а вместе с ним — роскошный пир и Состязание фонарей в небесных чертогах. Отправляясь на торжество, Се Лянь не подозревает, что окажется в центре внимания. По возвращении в мир смертных он принимается за работу, ведь теперь нужно прокормить сразу три лишних рта! К несчастью, на задании, которое получил принц, он вновь сталкивается с озлобленным духом младенца, однако Хуа Чэн успевает прийти на выручку.\nПосле всей неразберихи герои решают отдохнуть — но не тут-то было! На пороге храма Водяных Каштанов возникает Ши Цинсюань в сопровождении Мин И. Оказывается, Повелителя Ветра уже довольно давно преследует одна тварь: божок-пустозвон. Се Лянь соглашается помочь и расправиться с ним, и тут дело принимает совсем дурной оборот...
84	Человек, который смеется	2024-06-27	1	рус	0.00	https://s.f.kz/prod/568/567893_1000.jpg	«Человек, который смеется» — один из наиболее известных романов Виктора Гюго.\n\nВ центре повествования Гуинплен, в раннем детстве похищенный бандитами, до неузнаваемости обезобразившими его лицо. Судьба преподнесла ему немало испытаний, но душу не искалечила — преодолев нелегкий путь от ярмарочного актера до лорда и члена парламента, он смог остаться честным и благородным человеком…
85	Последняя песня	2024-06-28	1	рус	0.00	https://s.f.kz/prod/2424/2423552_1000.jpg	Роман Николаса Спаркса, который лег в основу одноименного фильма.\nРазвод труден для супругов, но еще тяжелее его переживают дети. Особенно подростки...\nРонни Миллер винит в распаде семьи своего отца, и когда приходится уехать к нему в маленький приморский городок на все лето, ей это кажется ссылкой... пока Ронни не встречает там новую подругу, а потом и парня, которому предназначено стать ее первой настоящей любовью.\nЕго зовут Уилл. Он красив и благороден, защищает животных и успешен в спорте. О таком можно лишь мечтать. Но летние романы не длятся вечно, и расставание совсем близко...
86	Тревожные люди	2024-06-29	1	рус	0.00	https://s.f.kz/prod/1891/1890055_1000.jpg	В маленьком шведском городке накануне Нового года вооруженный пистолетом человек в маске после неудачной попытки ограбить банк захватывает восемь заложников во время показа покупателям выставленной на продажу квартиры.У подъезда тут же собирается толпа жадных до сенсаций репортеров, полиция блокирует все подступы к дому и готовится штурмовать квартиру... атмосфера накаляется. Не выдерживая нарастающего напряжения, заложники делятся друг с другом своими самыми сокровенными тайнами...Вскоре грабитель начинает склоняться к тому, что, возможно, лучше добровольно отдать себя в руки полиции, чем продолжать оставаться в замкнутом пространстве со всеми этими невыносимыми людьми...
87	Труп в библиотеке	2024-06-30	1	рус	0.00	https://s.f.kz/prod/1455/1454323_1000.jpg	Все было прекрасно, когда Долли Бантри проснулась у себя дома в тихой деревушке Сент-Мэри-Мид. Прекрасно — до тех пор, пока в комнате, где располагалась семейная библиотека, не было обнаружено тело молодой девушки. По словам врача, она была задушена вчерашним вечером до полуночи. При этом никто в доме не знал эту привлекательную блондинку. Так кто же она, и кто мог убить ее? Подозрение сразу пало на мужа Долли, полковника в отставке, имеющего репутацию волокиты. Тот клянется, что в жизни не видел эту девушку. Но как же она тогда оказалась в его библиотеке? Тогда Долли зовет на помощь свою давнюю подругу мисс Марпл, чтобы та нашла настоящего убийцу — или открыла страшную правду о муже...
88	Слушай песню ветра. Пинбол 1973	2024-07-01	1	рус	0.00	https://s.f.kz/prod/2708/2707398_1000.jpg	История началась 8 августа 1970 года и закончилась через 18 дней — то есть 26 августа того же года. Вся история длится девятнадцать полных дней, уверяет нас герой-рассказчик, но не автор. Если посчитать внимательно, ты увидишь, что в девятнадцать дней эта история не укладывается, хоть тресни. Поверишь или проверишь? «За лето мы с Крысой выпили 25-метровый бассейн пива и покрыли пол «Джейз-бара» пятисантиметровым слоем арахисовой шелухи. Если бы мы этого не делали, то просто бы не выжили, такое скучное было лето«.\n\nВсе, что главному герою близко и дорого, кончается. Утрачивается, умирает и переходит в иной мир. Главный вопрос — как ко всему этому относиться? «Пинбол» начинается в сентябре 1973 года. Это вход. Но главное — понять, где выход. Выход должен быть. Обязательно.
89	Тихая гавань	2024-07-02	1	рус	0.00	https://s.f.kz/prod/2054/2053048_1000.jpg	Кэти — женщина, много лет страдавшая от жестокости мужа.\nПолиция не могла ей помочь — ведь именно там служил человек, превративший ее жизнь в ад...\nИ вот однажды терпение Кэти лопнуло. Потеряв надежду на спасение, она совершила отчаянный побег — и обрела «тихую гавань» в маленьком спокойном южном городке.\nНо готова ли Кэти к новым отношениям? Способна ли вновь поверить мужчине, понять его и полюбить? Даже если речь идет о таком обаятельном человеке, как молодой вдовец Алекс Уитли, который видит в Кэти не только возлюбленную и подругу, но и мать для своих детей.\nАлекс и Кэти идут по тонкому льду неизвестности, — а между тем муж Кэти уже начал ее поиски...
90	Мастер и Маргарита	2024-07-03	1	рус	0.00	https://s.f.kz/prod/2939/2938721_1000.jpg	Самый известный и неоднозначный роман XX века в современном оформлении.\n\nВ Москву прибывает дьявол со своей свитой — так начинается череда невероятных происшествий. В это время талантливый мастер томится в доме скорби, а его возлюбленная Маргарита пытается его спасти. Творчество или ремесло, благополучие или страсть, справедливость или карьера, свет или тьма... Каждый из героев делает собственный выбор, который становится вкладом в вечную борьбу добра и зла.
91	На вилле	2024-07-04	1	рус	0.00	https://s.f.kz/prod/133/132796_1000.jpg	Италия, тридцатые годы.\nГде-то приходят к власти фашисты...\nГде-то гибнут невинные люди...\nА на роскошной итальянской вилле надежно отгородилась от мира лекомысленная компания эстетствующих «светских львов и львиц» и их богемных приятелей.\nОни развлекаются изысканными разговорами и любовными интрижками.\nОни делают вид, что их маленькая искусственная реальность — единственно возможная.\nНо жизнь вторгается в красивую игру. Вторгается неожиданно и жестоко...
92	Дюна	2024-07-05	1	рус	0.00	https://s.f.kz/prod/1609/1608205_1000.jpg	Роман «Дюна», первая книга прославленной саги, знакомит читателя с Арракисом — миром суровых пустынь, исполинских песчаных червей, отважных фрименов и таинственной специи. Безграничная фантазия автора создала яркую, почти осязаемую вселенную, в которой есть враждующие Великие Дома, могущественная Космическая Гильдия, загадочный Орден Бинэ Гессерит и неуловимые ассасины. По мотивам «Дюны» снял свой гениальный фильм Дэвид Линч, а в 2020 году поклонников произведения ждет новая экранизация Дени Вильнёва, главные роли в которой исполнят Стеллан Скарсгард, Тимоти Шаламе, Зендая, Джейсон Момоа и другие.
93	Парень встретил парня	2024-07-06	1	рус	0.00	https://s.f.kz/prod/2563/2562968_1000.jpg	Однажды в книжном магазине Пол знакомится с Ноем, который совсем недавно переехал в город. Вскоре они проводят все свободное время вместе — пока Пол не совершает ошибку. А тут еще его лучшая подруга Джони отдаляется и не отвечает на звонки, а друг Тони страдает из-за плохих отношений с родителями. И кстати, подготовка к выпускному тоже идет не по плану. Но Пол не готов сдаваться — и сделает все, что можно, ради своих друзей.
94	Благословение небожителей. Том 4	2024-07-07	1	рус	0.00	https://s.f.kz/prod/2802/2801289_1000.jpg	Двум тиграм не ужиться на одной горе.\n\nРаз в сотню лет разгорается огонь в Медной Печи, и тысячи демонов откликаются на её зов. Они спешат к горе Тунлу — и каждый надеется выжить в смертельной схватке за право стать новым князем демонов. Предотвратить надвигающуюся угрозу по силам лишь небожителям.\n\nНо в столице бессмертных царит хаос: пленённая нечисть вырвалась на свободу, а боги войны заняты своими делами. Чтобы не допустить рождения очередного «непревзойдённого», Се Лянь отправляется к вулкану в сопровождении Хуа Чэна, вынужденного скрывать свой истинный облик под личиной маленького мальчика...
95	Артур, Луи и Адель	2024-07-08	1	рус	0.00	https://s.f.kz/prod/1604/1603507_1000.jpg	«Пронзительная история любви и дружбы. Легкая, словно ветер. Терпкая, словно пепел», — так описала историю популярная писательница Анна Джейн.\nАртур, Луи и Адель познакомились однажды летом на юге Франции и с тех пор каждые каникулы проводили вместе. Беззаботные дни, звездные ночи, запах соленого моря, счастливые улыбки и... первая любовь.\nТолько Адель не помнит этого. Несчастный случай забрал у нее все до единого воспоминания.\nЧто останется от человека, если лишить его памяти? Адель не может жить дальше, пока не узнает правду. Что скрывает ее прошлое? Знает Артур, но он не готов рассказать. Молодой человек уверен: порой воспоминания страшнее неведения.\nАдель ничего не помнит, Артур помнит все, а Луи расскажет, как все было.
96	Вечеринка в Хэллоуин	2024-07-09	1	рус	0.00	https://s.f.kz/prod/1536/1535885_1000.jpg	Писательница Ариадна Оливер приглашена в дом подруги, где в самом разгаре приготовления к празднованию Хэллоуина — веселого карнавала для детишек и взрослых. Одна из гостей — девочка-подросток, известная тем, что обожает рассказывать завиральные истории о всяких тайнах. Вот и теперь она поразила общество рассказом о том, что когда-то видела самое настоящее убийство! Никто не поверил ей. И вдруг в тот же вечер ее нашли... утопленной в ведре с водой и яблоками! Чертовски странно. Кому понадобилась смерть девочки? Возможно, она действительно видела нечто, что представляло опасность для кого-то из присутствовавших на вечеринке? В любом случае, Эркюль Пуаро, который взялся помочь миссис Оливер, своей старой знакомой, встал перед непростой задачей сорвать карнавальную маску с убийцы и явить обществу его лицо...
97	Дюна: Дюна. Мессия Дюны. Дети Дюны	2024-07-10	1	рус	0.00	https://s.f.kz/prod/1868/1867498_1000.jpg	Фрэнк Герберт (1920–1986) успел написать много, но в истории остался прежде всего как автор эпопеи «Дюна» — возможно, самой прославленной саги в научной фантастике, саги, переведенной на десятки языков и завоевавшей по всему миру миллионы поклонников. Самый авторитетный журнал научной фантастики «Локус» признал «Дюну», первый роман эпопеи о песчаной планете, лучшим научно-фантастическим романом всех времен и народов. В «Дюне» Фрэнку Герберту удалось совершить невозможное — создать своеобразную «хронику далекого будущего». И не было за всю историю мировой фантастики картины грядущего более яркой, более зримой, более мощной и оригинальной.\nВ сборник включены первые три романа из знаменитого цикла Фрэнка Герберта: «Дюна», «Мессия Дюны» и «Дети Дюны».
109	Нелюбовь сероглазого короля	2024-07-22	1	рус	0.00	https://s.f.kz/prod/1810/1809826_1000.jpg	Даша Севастьянова думала, что единственное, о чем ей стоит беспокоиться в последний учебный год, — это предстоящие экзамены, но внезапно проблемы стали нарастать как снежный ком. Непонимание в семье, возникшая симпатия к другу детства, новенькая, которая решила занять ее место и стать самой популярной девчонкой в школе... А еще эта дурацкая необходимость изображать влюбленную пару вместе с ненавистным Робертом Кайзером! И все из-за того, что кое-кто не умеет держать язык за зубами. Кайзер думает, что обыграл ее? Что ж, вызов принят...
98	Бегущий в Лабиринте. Испытание огнем. Лекарство от смерти	2024-07-11	1	рус	0.00	https://s.f.kz/prod/480/479061_1000.jpg	Вчера они были обычными парнями... Сегодня они — пешки в чужой игре, похищенные неизвестно кем для участия в чудовищном эксперименте. Их память стерта. Их новый дом — гигантский комплекс, отгороженный от еще более огромного Лабиринта стенами, которые раздвигаются утром и замыкаются вечером. И никто еще из тех, кто остался в Лабиринте после наступления ночи, не вернулся... Они не сомневаются: если сумеют разгадать тайну Лабиринта, то вырвутся из заточения и вернутся домой. Итак, что же будет дальше? Кто рискнет жизнью ради других, и кто выживет в смертоносном испытании?.. Читайте знаменитую трилогию Дж. Дэшнера, разошедшуюся по миру тиражом более 2 млн экземпляров!
99	Королева ничего	2024-07-12	1	рус	0.00	https://s.f.kz/prod/1581/1580135_1000.jpg	Долгожданный финал культовой серии, которая разошлась по миру тиражом свыше 10 миллионов экземпляров! Власть легче завоевать, чем удержать. Джуд, Верховная королева Эльфхейма, лишена власти, но не сломлена. Даже находясь в изгнании, она не оставляет надежды вернуться ко Двору фейри. Отринув свои чувства к Кардану, Джуд спешит на помощь сестре, которая попала в беду в Фейриленде. Но знакомого Эльфхейма больше нет. Война на пороге. И Джуд должна собрать все свои силы, чтобы спасти то, что осталось. И даже, возможно, стать настоящей королевой монстров, если потребуется...
100	Она и ее кот	2024-07-13	1	рус	0.00	https://s.f.kz/prod/3037/3036269_1000.jpg	Однажды девушка случайно встречает кота.\nОсобенный момент, который меняет в их жизни все.\nУ нее много проблем: переживания о личной жизни, работа, ежедневная рутина и бесконечное чувство одиночества. Теперь за ней присматривает кот, который ловит каждый ее вздох, видит ее слезы и каждый день рассказывает другим котам, как сильно ее любит, а ей — как прошел его день. Но она не понимает его, а он не знает, как ей помочь...\nИ все же они тянутся друг к другу.
101	Лолита	2024-07-14	1	рус	0.00	https://s.f.kz/prod/2097/2096800_1000.jpg	В детстве Гумберт Гумберт пережил взаимную, но невероятно болезненную влюбленность в девочку Анабель Ли. Родителям в конце концов пришлось их разлучить, и Анабель в скором времени умерла. Это событие, по мнению самого Гумберта, стало причиной, по которой он испытывает страсть к несовершеннолетним девочкам, которых называет «нимфетками».\n1947 год. Тридцатисемилетний преподаватель французской литературы Гумберт Гумберт снимает дом на северо-востоке США, в Новой Англии. У хозяйки дома Шарлотты Гейз есть двенадцатилетняя дочь Долорес, которую Гумберт ласково называет Лолитой. Вспыхнувшая в Гумберте страсть заставляет его искать любой предлог, чтобы оказаться наедине с Лолитой. Так начинается история, которая изменит литературу XX века.\nРоман был экранизирован дважды: в 1962 году Стэнли Кубриком с Джеймсом Мэйсоном, Шелли Уинтерс и Сью Лайон в главных ролях и в 1997 году Эдрианом Лайном с Джереми Айронсом, Доминик Суэйн и Мелани Гриффит.
102	Охота на овец	2024-07-15	1	рус	0.00	https://s.f.kz/prod/1848/1847386_1000.jpg	«О ее смерти мне сообщил по телефону старый приятель...»\nТак начинается «Охота на Овец» — пожалуй, самое странное путешествие по закоулкам современного мира и человеческого сознания, придуманное легендой современной литературы японским писателем Харуки Мураками. Этот роман стал абсолютным мировым бестселлером: «охота на овец» в наших душах не закончится никогда!\nИтак, позвоните в полицию, спросите адрес и номер телефона семьи, затем позвоните семье и узнайте дату и время похорон. А после, в назначенный день, садитесь в пригородную электричку от станции Васэда. И надейтесь, что охота будет удачной...
103	Влюбленная ведьма. Книга 2	2024-07-16	1	рус	0.00	https://s.f.kz/prod/1807/1806690_1000.jpg	Соскучились? Я — Таня Ведьмина и могу влюбить в себя абсолютно любого мужчину. Даже если он — мой преподаватель в университете. Я покорила сердце Олега Владыко несмотря на его яростное сопротивление. Он притворился моим парнем и сам не заметил как влюбился! У него самые нежные губы и самые ласковые руки...а еще ужасный характер. Меня тянет к нему, но что-то не дает нам быть вместе: гордость, чужая зависть, ложь, месть...Опасность грозит не просто нашим отношениям, но жизни Олега. И я должна спасти его и нашу любовь. Любой ценой. Влюбленная ведьма способна на все — так гласит закон Тани Ведьминой.
104	Смерть на Ниле	2024-07-17	1	рус	0.00	https://s.f.kz/prod/1288/1287490_1000.jpg	На роскошном пароходе «Карнак», плывущем по Нилу, убита молодая миллионерша, недавно вышедшая замуж и, как выяснилось, имевшая множество врагов среди пассажиров. Любой мог убить самоуверенную и нагловатую девушку, укравшую жениха у лучшей подруги. Но ни один из вероятных подозреваемых не совершал этого преступления... К счастью, на пароходе находится великий сыщик Эркюль Пуаро, который знает все общество, представленное в круизе, еще по Лондону, и в курсе возможных мотивов каждого из присутствующих. И, конечно, первое, о чем задумывается бельгиец, — это о «любовном треугольнике», состоявшем из убитой, ее свежеиспеченного мужа и очень темпераментной женщины, которую тот бросил ради миллионерши...
105	И эхо летит по горам	2024-07-18	1	рус	0.00	https://s.f.kz/prod/359/358993_1000.jpg	Долгожданный новый роман автора книг «Бегущий за ветром» и «Тысяча сияющих солнц». Главный мировой бестселлер 2013 года! Книга будет издана в 80 странах. 1952 год, звездная ночь в пустыне, отец рассказывает афганскую притчу сыну и дочери. Они устроились на ночлег в горах, на пути в Кабул. Черное небо, сияющие звезды и камни да колючки на многие мили вокруг. Затаив дыхание, Абдулла и совсем еще маленькая Пари слушают историю о том, как одного мальчика похитил ужасный Дэйв и бедняге предстоит самая страшная судьба на свете. Но жизнь не раскрашена в черно-белые тона — даже в сказках… Наутро отец и дети продолжают путь в Кабул, и этот день станет развилкой их судеб. Они расстанутся и, возможно, навсегда. Разлука брата и сестры даст начало сразу нескольким сплетающимся и расплетающимся историям. И в центре этой паутины жизни — Пари, нареченная этим именем вовсе не в честь французской столицы, а потому что так зовут на фарси фей. Пять поколений, немало стран и городов будут вовлечены в притчу жизни Пари-феи, — жизни, которая разворачивается через войны, рождения, смерти, любви, предательства и надежды. И всем героям предстоит не раз принимать решения, большие и малые, которые, сплетясь, образуют одно главное решение их жизней. Новый роман Халеда Хоссейни, прозрачный, пронзительный, многоголосый, — о том, что любое решение, принятое за другого человека, — добра ради или зла — имеет цену, и судьба непременно выставит за него счет. Это роман о силе дешевых слов и дорогих поступков, о коварстве жизненного предназначения, о неизбежности воздаяния, о шумном малодушии и безмолвной преданности. Это новый роман Халеда Хоссейни.
106	Карты на столе	2024-07-19	1	рус	0.00	https://s.f.kz/prod/1426/1425140_1000.jpg	Не зря говорится: «Не буди лихо, пока оно тихо»... Но мистер Шайтана, собирая в своем доме друзей и знакомых на партию в бридж, забыл об этой расхожей истине. Весь вечер он говорил об убийствах и убийцах, настаивая на том, что большинство подобных преступлений так и не были раскрыты. Вот и напросился, что называется...\nЧерез некоторое время его находят в гостиной с ножевой раной в груди. По счастью, среди гостей присутствуют сразу четверо выдающихся сыщиков — Эркюль Пуаро, Ариадна Оливер, суперинтендант Баттл и полковник Рейс. Скопище блестящих умов... Однако и им приходится напрячь все свои аналитические способности, чтобы изобличить убийцу. Ведь разгадка кроется в ходе карточной партии, и найдет ее тот, кто лучше всех играет в бридж...
107	Танцующая с бурей	2024-07-20	1	рус	0.00	https://s.f.kz/prod/3304/3303223_1000.jpg	В мощной империи, где правят сёгуны, а старинные законы и древние искусства соседствуют с умными машинами и новыми технологиями, юная Юкико и ее отец, доблестный воин Масару, по капризу правителя отправляются на поиски арашиторы, страшного мистического чудовища из народных легенд. Говорят, что арашиторы давно вымерли и участь охотников печальна: они или вернутся к сёгуну ни с чем, рискуя быть казненными, или попадут в лапы безжалостного монстра, которого никому не удавалось пленить.
108	«Магазин снов» мистера Талергута	2024-07-21	1	рус	0.00	https://s.f.kz/prod/2204/2203815_1000.jpg	Когда люди засыпают, их души устремляются в загадочный город, где в воздухе порхают сказочные лепрекрылы, а по улицам носятся трудолюбивые ноктилуки. Центром города является таинственный «Магазин снов» мистера Талергута, в котором можно найти сны на любой вкус, но за соответствующую цену.\nПенни — новая сотрудница магазина, и теперь каждый ее день полон сюрпризов, необычных знакомств и необыкновенных приключений. На ее глазах благодаря снам люди влюбляются, заряжаются энергией, обретают уверенность в себе, находят вдохновение и даже заглядывают в будущее.\n\nДебютный роман южнокорейской писательницы Ли Мие — это причудливая смесь фантазии и реальности, оригинальный и остроумный взгляд на притягательный мир сновидений.
110	Цветок пустыни. Реальная история супермодели Варис Дирие	2024-07-23	1	рус	0.00	https://s.f.kz/prod/2380/2379950_1000.jpg	«Варис родилась в Сомали. В 5 лет она подверглась процедуре женского обрезания. А в 13 ушла из дома, чтобы не выходить замуж за старика. Сбежав в Лондон, она не знала ни языка, ни местной культуры, жила без документов и работала уборщицей в ресторане, пока судьбоносная встреча с именитым фотографом не изменила ее жизнь. Сегодня Варис Дирие — известная топ-модель, писательница и специальный посол ООН за права женщин. Ее история — прекрасное вдохновение для каждого! Из Сомалийской пустыни она добралась до мира высокой моды. Она боролась с несправедливостью и вышла из этой борьбы победительницей», — Элтон Джон.
111	На Западном фронте без перемен	2024-07-24	1	рус	0.00	https://s.f.kz/prod/499/498053_1000.jpg	Говоря о Первой мировой войне, всегда вспоминают одно произведение Эриха Марии Ремарка.\n«На Западном фронте без перемен».\nЭто рассказ о немецких мальчишках, которые под действием патриотической пропаганды идут на войну, не зная о том, что впереди их ждет не слава героев, а инвалидность и смерть.\nКаждый день войны уносит жизни чьих-то отцов, сыновей, а газеты тем временем бесстрастно сообщают: "На Западном фронте без перемен...".\nЭта книга — не обвинение, не исповедь.\nЭто попытка рассказать о поколении, которое погубила война, о тех, кто стал ее жертвой, даже если сумел спастись от снарядов и укрыться от пули.
112	Аристотель и Данте открывают тайны Вселенной	2024-07-25	1	рус	0.00	https://s.f.kz/prod/1445/1444213_1000.jpg	Аристотель — замкнутый подросток, брат которого сидит в тюрьме, а отец до сих пор не может забыть войну. Данте — умный и начитанный парень с отличным чувством юмора и необычным взглядом на мир. Однажды встретившись, Аристотель и Данте понимают, что совсем друг на друга не похожи, однако их общение быстро перерастает в настоящую дружбу. Благодаря этой дружбе они находят ответы на сложные вопросы, которые раньше казались им непостижимыми загадками Вселенной, и наконец осознают, кто они на самом деле.
113	Среди тысячи слов	2024-07-26	1	рус	0.00	https://s.f.kz/prod/1868/1867795_1000.jpg	УИЛЛОУ\nЕе душа помнила все. Знала, что такое одиночество в огромном городе.\nИ каково видеть лишь темноту, когда кругом обжигающий свет.\nОна сбегала от мира на страницы книг.\nНа прослушивании кричала. Это крик был с ней внутри. Каждую минуту.\nЕй нужна была эта роль, чтобы изгнать своих демонов и обрести спокойствие. А затем — просто исчезнуть, не оставив следа.\nНевинная надежда. Которая разлетится на миллион чертовых осколков.\n\nАЙЗЕК\nОн был гладким клинком. Резал взглядом.\nОн словно пришел из другого мира. До него нельзя было дотронуться.\nОн играл так, что на глазах у всех выступали слезы. А боль растворялась.\nСцена стала для него сродни очищению: столько гнева и сожалений.\nОн мечтал обрести свой собственный голос и уехать прочь из этого города.\nЕго талант — это все, что у него было. Пока не появилась она.
114	Алхимик	2024-07-27	1	рус	0.00	https://s.f.kz/prod/486/485537_1000.jpg	"Алхимик" — самый известный роман бразильского писателя Пауло Коэльо, любимая книга миллионов людей во всем мире. В юности люди не боятся мечтать, все кажется им возможным. Но проходит время, и таинственная сила принимается им внушать, что их желания неосуществимы. "Добиться воплощения Своей Судьбы — вот единственная подлинная обязанность человека..." — утверждает Пауло Коэльо.\nЭтот, ставший культовым, роман-притча способен изменить жизнь своих читателей.
115	Загадай любовь	2024-07-28	1	рус	0.00	https://s.f.kz/prod/2191/2190480_1000.jpg	Неразделенная любовь - это проблема, и еще какая! Наташа Зуева знает об этом не понаслышке, ведь она безответно влюблена. Накануне новогодних праздников может случиться настоящее чудо. Наташа вместе с классом проведет зимние каникулы на прекрасном горнолыжном курорте. Какая девушка откажется встретить Новый год с парнем своей мечты? И неважно, что предмет воздыханий Наташи совершенно не подозревает о ее чувствах.Казалось, никто и ничто не сможет испортить ей праздник. Разве только ненавистный Тимур Макеев, самодовольный парень и главный прогульщик в школе.Ася Лавринович - один из самых популярных молодежных авторов. Суммарный тираж ее книг составляет более 100 000 экземпляров.
116	Скажи мне, что ты меня любишь...	2024-07-29	1	рус	0.00	https://s.f.kz/prod/2348/2347191_1000.jpg	История романа самого прославленного певца «потерянного поколения» Эриха Марии Ремарка и самой знаменитой «фам фаталь» мирового кинематографа Марлен Дитрих, поведанная ими самими — в письмах, которые они писали друг другу. Их отношения не были простыми. В них вспышки страсти и нежности слишком часто сменялись непониманием, ревностью, недоверием и даже враждой. Их отношения были и необходимы, и мучительны как для Дитрих, так и для Ремарка. Они должны были закончиться плохо — и закончились плохо. Но и сейчас письма великого писателя и гениальной актрисы трогают до глубины души...
117	Там, где живет любовь	2024-07-30	1	рус	0.00	https://s.f.kz/prod/1589/1588480_1000.jpg	Вере Азаровой хватало неприятностей, но их стало больше, когда мечта всех девчонок Марк Василевский обратил на нее внимание. Кто-то начал присылать Вере анонимные сообщения с угрозами. Возможно, это самая красивая девчонка школы, которая также положила глаз на Марка?\nВот только Вера не боится завистников и отправляется в летний лагерь, чтобы быть рядом с парнем своей мечты. Но сможет ли она добиться своей цели, когда вокруг столько соперниц и недоброжелателей?
118	День, когда я научился жить	2024-07-31	1	рус	0.00	https://s.f.kz/prod/2042/2041963_1000.jpg	Лоран Гунель входит в пятерку самых популярных беллетристов Франции. Его новая книга называется «День, когда я научился жить». Представьте, что в один прекрасный день цыганка, посмотрев на вашу ладонь, испуганно замолкает. И лишь после настойчивых просьб вы добиваетесь от нее ответа. Звучит он довольно страшно: «Ты скоро умрешь». Для Джонатана, героя романа Гунеля, мир совершенно меняется. И хотя никаких болезней врачи у него не находят, он решает провести несколько недель среди гор и холмов на берегу океана. Впервые он никуда не спешит. В поисках ответа на вечные вопросы о смысле жизни Джонатан заново знакомится с самим собой. Встречи, странствия, приключения и самопознание в корне меняют и его видение жизни, и саму жизнь...\n\nЭта книга проливает новый свет на наше существование и отношения с другими людьми, кажется, будто кто-то отворил окно и в затхлую атмосферу ворвался свежий ветер.
119	Мемуары гейши	2024-08-01	1	рус	0.00	https://s.f.kz/prod/560/559162_1000.jpg	Искусство гейши — не обольщать мужчин, но покорять их. Ее профессия — развлекать и очаровывать. Но вступать в связь с клиентом строго запрещено, а полюбить кого-то из них считается позором. Перед вами — история Саюри, девушки из простой крестьянской семьи, ставшей королевой гейш Киото. История вражды и соперничества, изощренных женских интриг и великой, пронесенной через десятилетия любви, ради которой Саюри дерзнула нарушить закон гейш…
120	Тетрадь в клеточку	2024-08-02	1	рус	0.00	https://s.f.kz/prod/1863/1862391_550.jpg	«Привет, Тетрадь в клеточку» — так начинается каждая запись в дневнике Ильи, который он начал вести после переезда. В новом городе Илья очень хочет найти друзей, но с ним разговаривают только девочка-мигрантка и одноклассник, про которого ходят странные слухи. Илья очень хочет казаться обычным, но боится микробов и постоянно моет руки. А еще он очень хочет забыть о страшном Дне S, но тот постоянно возвращается к нему в воспоминаниях.
121	Узорный покров	2024-08-03	1	рус	0.00	https://s.f.kz/prod/682/681977_1000.jpg	«Узорный покров» (1925) — полная трагизма история любви, разворачивающаяся в небольшом городке в Китае, куда приезжают бороться с эпидемией холеры молодой английский бактериолог с женой.\n\nКнига легла в основу голливудского фильма «Разрисованная вуаль», главные роли в котором исполнили великолепные Наоми Уоттс и Эдвард Нортон.
122	Королевство гнева и тумана	2024-08-04	1	рус	0.00	https://s.f.kz/prod/1911/1910459_1000.jpg	Фейра уже не та простая смертная девушка, какой была на землях людей. Здесь, в Притиании, она обрела бессмертие, развила магические способности, ее возлюбленный — верховный правитель Двора весны, и скоро состоится их свадьба. Но когда-то она заключила договор с правителем Двора ночи и обязана неделю каждого месяца проводить в соседних владениях, о которых идет недобрая слава. Между тем владыка Сонного королевства, давний враг Притиании и мира людей, готовит вторжение на их земли. В его руках мощный артефакт, давший когда-то жизнь всему миру и способный оживлять мертвых. Противостоять его силе может лишь Книга Дуновений, и Фейра делает все возможное и невозможное, чтобы заполучить Книгу.
3	Спеши любить	2024-04-07	1	рус	3.00	https://s.f.kz/prod/786/785876_1000.jpg	Тихий городок Бофор.\nКаждый год Лэндон Картер приезжает сюда, чтобы вспомнить историю своей первой любви...\n\nИсторию страсти и нежности, много лет назад связавшей его, парня из богатой семьи, и Джейми Салливан, скромную дочь местного пастора.\n\nИстория радости и грусти, счастья и боли.\n\nИсторию чувства, которое человеку доводится испытать лишь раз в жизни — и запомнить навсегда...
\.


--
-- Data for Name: books_authors; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.books_authors (book_id, author_id) FROM stdin;
3	122
4	123
5	124
6	125
7	126
8	127
9	128
10	129
11	130
12	131
13	132
14	133
15	134
16	135
17	136
18	137
19	138
20	139
21	140
22	141
23	142
24	143
25	144
26	145
27	146
28	147
29	148
30	149
31	150
32	151
33	152
34	153
35	154
36	155
37	156
38	157
39	158
40	159
41	160
42	161
43	162
44	163
45	164
46	165
47	166
48	167
49	168
50	169
51	170
52	171
53	172
54	173
55	174
56	175
57	176
58	177
59	178
60	179
61	180
62	181
63	182
64	183
65	184
66	185
67	186
68	187
69	188
70	189
71	190
72	191
73	192
74	193
75	194
76	195
77	196
78	197
79	198
80	199
81	200
82	201
83	202
84	203
85	204
86	205
87	206
88	207
89	208
90	209
91	210
92	211
93	212
94	213
95	214
96	215
97	216
98	217
99	218
100	219
101	220
102	221
103	222
104	223
105	224
106	225
107	226
108	227
109	228
110	229
111	230
112	231
113	232
114	233
115	234
116	235
117	236
118	237
119	238
120	239
121	240
\.


--
-- Data for Name: books_categories; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.books_categories (book_id, category_id) FROM stdin;
3	2
4	4
5	5
6	3
7	4
8	1
9	1
10	3
11	1
12	4
13	3
14	1
15	4
16	1
17	2
18	1
19	4
20	3
21	5
22	2
23	5
24	1
25	3
26	8
27	4
28	1
29	4
30	1
31	3
32	5
33	1
34	1
35	3
36	4
37	1
38	1
39	4
40	5
41	1
42	4
43	1
44	1
45	3
46	1
47	1
48	4
49	3
50	1
51	1
52	3
53	3
54	3
55	4
56	3
57	1
58	4
59	4
60	10
61	4
62	1
63	7
64	1
65	1
66	6
67	3
68	5
69	2
70	1
71	1
72	4
73	2
74	1
75	5
76	4
77	1
78	5
79	4
80	5
81	3
82	7
83	1
84	1
85	5
86	1
87	4
88	1
89	1
90	2
91	1
92	3
93	4
94	5
95	3
96	2
97	2
98	1
99	1
100	1
101	4
102	5
103	1
104	5
105	3
106	3
107	4
108	12
109	9
110	1
111	1
112	7
113	1
114	12
115	4
116	1
117	1
118	1
119	4
120	3
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.categories (id, name) FROM stdin;
1	Проза
2	Фантастика
3	Фэнтези
4	Любовные романы
5	Детективы
6	Мистика
7	Психологические романы
8	Триллеры
9	Исторические романы
10	Повести, рассказы
11	Биографии людей
12	Мемуары
13	Манга, комиксы и артбуки (издания для взрослых)
\.


--
-- Data for Name: friends; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.friends (id, user_id, friend_id, status) FROM stdin;
1	9	9	sent
3	8	1	sent
\.


--
-- Data for Name: goose_db_version; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.goose_db_version (id, version_id, is_applied, tstamp) FROM stdin;
1	0	t	2024-03-30 23:34:10.05031
2	20240301010615	t	2024-03-30 23:34:10.060291
3	20240310180400	t	2024-03-30 23:34:10.068239
4	20240328074047	t	2024-03-30 23:34:10.07537
5	20240331214754	t	2024-03-31 18:56:03.630311
6	20240331215120	t	2024-03-31 18:56:03.636307
7	20240402194315	t	2024-04-04 19:40:18.323807
8	20240402194542	t	2024-04-04 19:40:18.342171
9	20240403042942	t	2024-04-04 19:40:18.356654
10	20240403044147	t	2024-04-04 19:40:18.361343
11	20240404222214	t	2024-04-04 19:40:18.368044
12	20240406123816	t	2024-04-07 18:02:32.149546
13	20240406174040	t	2024-04-07 18:02:32.166807
14	20240407124437	t	2024-04-07 18:02:32.171291
15	20240418235057	t	2024-04-18 20:56:28.569889
16	20240419022301	t	2024-04-25 02:41:37.86107
17	20240425045429	t	2024-04-25 02:41:37.875098
18	20240425052711	t	2024-04-25 02:41:37.882133
19	20240502001319	t	2024-05-01 19:58:26.231312
\.


--
-- Data for Name: liked_books; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.liked_books (id, user_id, book_id) FROM stdin;
7	9	16
8	9	15
9	9	113
\.


--
-- Data for Name: question_results; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.question_results (id, quiz_result_id, quiestion_id, user_answer, is_correct) FROM stdin;
1	1	3	Darina	f
2	1	4	Collegue	f
3	2	3	Diwka	f
4	2	4	Collegue	f
5	3	3	Diwka	f
6	3	4		f
7	4	3		f
8	4	4	Collegue	f
9	5	3		f
10	5	4	Collegue	f
11	6	3		f
12	6	4	Ulan	f
13	7	3		f
14	7	4	Ulan	f
15	7	4	Ulan	f
16	8	3	Diana	t
17	8	4	IDK	t
18	9	3	Diana	t
19	9	4	IDK	t
20	9	4	IDK	t
21	10	5	2	t
22	10	5	2	t
23	10	5	2	t
\.


--
-- Data for Name: questions; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.questions (id, quiz_id, question, options, answer) FROM stdin;
1	7	fdfaf	["cas", "cas", "cas"]	sdcasd
2	8	heth	["gf", "gd", "gd"]	hdffd
3	9	Who are u?	["Diwka", "Di", "Darina"]	Diana
4	9	Where is uTo?	["Me also", "Collegue", "Ulan"]	IDK
5	10	What is that	["sdcsd", "fsdfsdf", "fsdfsds"]	2
\.


--
-- Data for Name: quiz_comments; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.quiz_comments (id, quiz_id, user_id, comment, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: quiz_ratings; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.quiz_ratings (id, quiz_id, user_id, rating) FROM stdin;
1	1	8	4
\.


--
-- Data for Name: quiz_results; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.quiz_results (id, quiz_id, user_id, coorect, total) FROM stdin;
1	9	9	0	2
2	9	9	0	2
3	9	9	0	2
4	9	9	0	2
5	9	9	0	2
6	9	9	0	2
7	9	9	0	3
8	9	9	2	2
9	9	9	3	3
10	10	9	3	3
\.


--
-- Data for Name: quizzes; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.quizzes (id, user_id, book_id, title, rating, created_at) FROM stdin;
2	9	3	jana	0.00	2024-04-18 20:56:28.569889+00
3	9	3	jana	0.00	2024-04-18 20:56:28.569889+00
4	9	7	Exchange	0.00	2024-04-18 20:56:28.569889+00
5	9	7	Exchange	0.00	2024-04-18 20:56:28.569889+00
6	9	9	love	0.00	2024-04-18 20:56:28.569889+00
7	9	10	okay	0.00	2024-04-18 20:56:28.569889+00
8	9	10	Exchange	0.00	2024-04-18 20:56:28.569889+00
9	9	6	Look at this?	0.00	2024-04-19 12:23:00.572418+00
1	9	5	string	4.00	2024-04-18 20:56:28.569889+00
10	9	5	Dune	0.00	2024-05-03 06:46:47.671083+00
\.


--
-- Data for Name: share_requests; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.share_requests (id, sender_id, receiver_id, sender_book_id, receiver_book_id, sender_status, receiver_status, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: stock_books; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.stock_books (id, user_id, book_id) FROM stdin;
2	8	4
3	9	6
4	8	3
6	8	5
7	9	10
8	9	7
9	9	20
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: diploma
--

COPY public.users (id, username, email, password, role, city) FROM stdin;
1	admin	admin@mail.com	password	admin	Almaty
2	+7778	grewgweg	12345678	user	\N
3	+7	grewgweg	12345678	user	\N
4	+7	grewgweg	12345678	user	\N
5	+7	grewgweg	12345678	user	\N
6	Diwka	diwka	123	user	\N
7	diwkaa	diwka2	123	user	\N
9	as	as	as	user	\N
10	qwe	mani	pp	user	\N
11	uTo	uto	uto	user	\N
8	dhucs	ddd	dhucs	user	\N
\.


--
-- Name: authors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.authors_id_seq', 240, true);


--
-- Name: book_comments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.book_comments_id_seq', 1, false);


--
-- Name: book_ratings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.book_ratings_id_seq', 2, true);


--
-- Name: books_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.books_id_seq', 122, true);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.categories_id_seq', 13, true);


--
-- Name: friends_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.friends_id_seq', 3, true);


--
-- Name: goose_db_version_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.goose_db_version_id_seq', 19, true);


--
-- Name: liked_books_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.liked_books_id_seq', 9, true);


--
-- Name: question_results_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.question_results_id_seq', 23, true);


--
-- Name: questions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.questions_id_seq', 5, true);


--
-- Name: quiz_comments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.quiz_comments_id_seq', 1, false);


--
-- Name: quiz_ratings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.quiz_ratings_id_seq', 1, true);


--
-- Name: quiz_results_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.quiz_results_id_seq', 10, true);


--
-- Name: quizzes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.quizzes_id_seq', 10, true);


--
-- Name: share_requests_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.share_requests_id_seq', 1, false);


--
-- Name: stock_books_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.stock_books_id_seq', 10, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: diploma
--

SELECT pg_catalog.setval('public.users_id_seq', 11, true);


--
-- Name: authors authors_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.authors
    ADD CONSTRAINT authors_pkey PRIMARY KEY (id);


--
-- Name: book_comments book_comments_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.book_comments
    ADD CONSTRAINT book_comments_pkey PRIMARY KEY (id);


--
-- Name: book_ratings book_ratings_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.book_ratings
    ADD CONSTRAINT book_ratings_pkey PRIMARY KEY (id);


--
-- Name: books_authors books_authors_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.books_authors
    ADD CONSTRAINT books_authors_pkey PRIMARY KEY (book_id, author_id);


--
-- Name: books_categories books_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.books_categories
    ADD CONSTRAINT books_categories_pkey PRIMARY KEY (book_id, category_id);


--
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: friends friends_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.friends
    ADD CONSTRAINT friends_pkey PRIMARY KEY (id);


--
-- Name: goose_db_version goose_db_version_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);


--
-- Name: liked_books liked_books_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.liked_books
    ADD CONSTRAINT liked_books_pkey PRIMARY KEY (id);


--
-- Name: question_results question_results_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.question_results
    ADD CONSTRAINT question_results_pkey PRIMARY KEY (id);


--
-- Name: questions questions_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.questions
    ADD CONSTRAINT questions_pkey PRIMARY KEY (id);


--
-- Name: quiz_comments quiz_comments_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.quiz_comments
    ADD CONSTRAINT quiz_comments_pkey PRIMARY KEY (id);


--
-- Name: quiz_ratings quiz_ratings_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.quiz_ratings
    ADD CONSTRAINT quiz_ratings_pkey PRIMARY KEY (id);


--
-- Name: quiz_results quiz_results_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.quiz_results
    ADD CONSTRAINT quiz_results_pkey PRIMARY KEY (id);


--
-- Name: quizzes quizzes_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.quizzes
    ADD CONSTRAINT quizzes_pkey PRIMARY KEY (id);


--
-- Name: share_requests share_requests_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.share_requests
    ADD CONSTRAINT share_requests_pkey PRIMARY KEY (id);


--
-- Name: stock_books stock_books_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.stock_books
    ADD CONSTRAINT stock_books_pkey PRIMARY KEY (id);


--
-- Name: book_ratings unique_book_user_rating; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.book_ratings
    ADD CONSTRAINT unique_book_user_rating UNIQUE (book_id, user_id);


--
-- Name: friends unique_friendship; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.friends
    ADD CONSTRAINT unique_friendship UNIQUE (user_id, friend_id);


--
-- Name: quiz_ratings unique_quiz_user_rating; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.quiz_ratings
    ADD CONSTRAINT unique_quiz_user_rating UNIQUE (quiz_id, user_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: books_authors books_authors_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.books_authors
    ADD CONSTRAINT books_authors_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.authors(id);


--
-- Name: books_categories books_categories_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: diploma
--

ALTER TABLE ONLY public.books_categories
    ADD CONSTRAINT books_categories_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id);


--
-- PostgreSQL database dump complete
--

