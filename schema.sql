--
-- PostgreSQL database dump
--

-- Dumped from database version 12.8 (Ubuntu 12.8-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.8 (Ubuntu 12.8-0ubuntu0.20.04.1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: fixed_property; Type: TABLE; Schema: public; Owner: rtc
--

CREATE TABLE public.fixed_property (
    id integer NOT NULL,
    inventory text,
    serial text,
    name text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    warehouse integer DEFAULT 1,
    state integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.fixed_property OWNER TO rtc;

--
-- Name: TABLE fixed_property; Type: COMMENT; Schema: public; Owner: rtc
--

COMMENT ON TABLE public.fixed_property IS 'Основные средства';


--
-- Name: COLUMN fixed_property.inventory; Type: COMMENT; Schema: public; Owner: rtc
--

COMMENT ON COLUMN public.fixed_property.inventory IS 'Инвентарный номер';


--
-- Name: COLUMN fixed_property.serial; Type: COMMENT; Schema: public; Owner: rtc
--

COMMENT ON COLUMN public.fixed_property.serial IS 'Серийный номер';


--
-- Name: COLUMN fixed_property.name; Type: COMMENT; Schema: public; Owner: rtc
--

COMMENT ON COLUMN public.fixed_property.name IS 'Название';


--
-- Name: Fixed_property_id_seq; Type: SEQUENCE; Schema: public; Owner: rtc
--

CREATE SEQUENCE public."Fixed_property_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Fixed_property_id_seq" OWNER TO rtc;

--
-- Name: Fixed_property_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rtc
--

ALTER SEQUENCE public."Fixed_property_id_seq" OWNED BY public.fixed_property.id;


--
-- Name: actions; Type: TABLE; Schema: public; Owner: rtc
--

CREATE TABLE public.actions (
    id integer NOT NULL,
    action text NOT NULL,
    state text NOT NULL
);


ALTER TABLE public.actions OWNER TO rtc;

--
-- Name: actions_id_seq; Type: SEQUENCE; Schema: public; Owner: rtc
--

CREATE SEQUENCE public.actions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.actions_id_seq OWNER TO rtc;

--
-- Name: actions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rtc
--

ALTER SEQUENCE public.actions_id_seq OWNED BY public.actions.id;


--
-- Name: groups; Type: TABLE; Schema: public; Owner: rtc
--

CREATE TABLE public.groups (
    id integer NOT NULL,
    name text NOT NULL,
    description text,
    username integer DEFAULT 1 NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.groups OWNER TO rtc;

--
-- Name: groups_id_seq; Type: SEQUENCE; Schema: public; Owner: rtc
--

CREATE SEQUENCE public.groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.groups_id_seq OWNER TO rtc;

--
-- Name: groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rtc
--

ALTER SEQUENCE public.groups_id_seq OWNED BY public.groups.id;


--
-- Name: property_groups; Type: TABLE; Schema: public; Owner: rtc
--

CREATE TABLE public.property_groups (
    id integer NOT NULL,
    property_id integer NOT NULL,
    group_id integer NOT NULL
);


ALTER TABLE public.property_groups OWNER TO rtc;

--
-- Name: property_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: rtc
--

CREATE SEQUENCE public.property_groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.property_groups_id_seq OWNER TO rtc;

--
-- Name: property_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rtc
--

ALTER SEQUENCE public.property_groups_id_seq OWNED BY public.property_groups.id;


--
-- Name: property_history; Type: TABLE; Schema: public; Owner: rtc
--

CREATE TABLE public.property_history (
    id integer NOT NULL,
    action integer NOT NULL,
    note text,
    date timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    username integer NOT NULL,
    property_id integer NOT NULL
);


ALTER TABLE public.property_history OWNER TO rtc;

--
-- Name: TABLE property_history; Type: COMMENT; Schema: public; Owner: rtc
--

COMMENT ON TABLE public.property_history IS 'История изменений для основных средств';


--
-- Name: property_history_id_seq; Type: SEQUENCE; Schema: public; Owner: rtc
--

CREATE SEQUENCE public.property_history_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.property_history_id_seq OWNER TO rtc;

--
-- Name: property_history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rtc
--

ALTER SEQUENCE public.property_history_id_seq OWNED BY public.property_history.id;


--
-- Name: refresh_tokens; Type: TABLE; Schema: public; Owner: rtc
--

CREATE TABLE public.refresh_tokens (
    id integer NOT NULL,
    username text NOT NULL,
    token text NOT NULL
);


ALTER TABLE public.refresh_tokens OWNER TO rtc;

--
-- Name: refresh_tokens_id_seq; Type: SEQUENCE; Schema: public; Owner: rtc
--

CREATE SEQUENCE public.refresh_tokens_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.refresh_tokens_id_seq OWNER TO rtc;

--
-- Name: refresh_tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rtc
--

ALTER SEQUENCE public.refresh_tokens_id_seq OWNED BY public.refresh_tokens.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: rtc
--

CREATE TABLE public.roles (
    id integer NOT NULL,
    name text NOT NULL,
    description text
);


ALTER TABLE public.roles OWNER TO rtc;

--
-- Name: TABLE roles; Type: COMMENT; Schema: public; Owner: rtc
--

COMMENT ON TABLE public.roles IS 'Список ролей пользователей';


--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: rtc
--

CREATE SEQUENCE public.roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO rtc;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rtc
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: staff; Type: TABLE; Schema: public; Owner: rtc
--

CREATE TABLE public.staff (
    id integer NOT NULL,
    "table" text NOT NULL,
    name text NOT NULL,
    manager text NOT NULL,
    department text NOT NULL,
    job text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.staff OWNER TO rtc;

--
-- Name: staff_id_seq; Type: SEQUENCE; Schema: public; Owner: rtc
--

CREATE SEQUENCE public.staff_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.staff_id_seq OWNER TO rtc;

--
-- Name: staff_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rtc
--

ALTER SEQUENCE public.staff_id_seq OWNED BY public.staff.id;


--
-- Name: staff_property; Type: TABLE; Schema: public; Owner: rtc
--

CREATE TABLE public.staff_property (
    id integer NOT NULL,
    property_id integer NOT NULL,
    staff_id integer NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.staff_property OWNER TO rtc;

--
-- Name: staff_property_id_seq; Type: SEQUENCE; Schema: public; Owner: rtc
--

CREATE SEQUENCE public.staff_property_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.staff_property_id_seq OWNER TO rtc;

--
-- Name: staff_property_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rtc
--

ALTER SEQUENCE public.staff_property_id_seq OWNED BY public.staff_property.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: rtc
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username text NOT NULL,
    password text NOT NULL,
    description text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    active boolean DEFAULT false NOT NULL,
    role integer DEFAULT 1 NOT NULL,
    display_name text DEFAULT ''::text
);


ALTER TABLE public.users OWNER TO rtc;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: rtc
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO rtc;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rtc
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: warehouses; Type: TABLE; Schema: public; Owner: rtc
--

CREATE TABLE public.warehouses (
    id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.warehouses OWNER TO rtc;

--
-- Name: TABLE warehouses; Type: COMMENT; Schema: public; Owner: rtc
--

COMMENT ON TABLE public.warehouses IS 'Список складов, за которой закреплена техника';


--
-- Name: warehouses_id_seq; Type: SEQUENCE; Schema: public; Owner: rtc
--

CREATE SEQUENCE public.warehouses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.warehouses_id_seq OWNER TO rtc;

--
-- Name: warehouses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rtc
--

ALTER SEQUENCE public.warehouses_id_seq OWNED BY public.warehouses.id;


--
-- Name: actions id; Type: DEFAULT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.actions ALTER COLUMN id SET DEFAULT nextval('public.actions_id_seq'::regclass);


--
-- Name: fixed_property id; Type: DEFAULT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.fixed_property ALTER COLUMN id SET DEFAULT nextval('public."Fixed_property_id_seq"'::regclass);


--
-- Name: groups id; Type: DEFAULT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.groups ALTER COLUMN id SET DEFAULT nextval('public.groups_id_seq'::regclass);


--
-- Name: property_groups id; Type: DEFAULT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.property_groups ALTER COLUMN id SET DEFAULT nextval('public.property_groups_id_seq'::regclass);


--
-- Name: property_history id; Type: DEFAULT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.property_history ALTER COLUMN id SET DEFAULT nextval('public.property_history_id_seq'::regclass);


--
-- Name: refresh_tokens id; Type: DEFAULT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.refresh_tokens ALTER COLUMN id SET DEFAULT nextval('public.refresh_tokens_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: staff id; Type: DEFAULT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.staff ALTER COLUMN id SET DEFAULT nextval('public.staff_id_seq'::regclass);


--
-- Name: staff_property id; Type: DEFAULT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.staff_property ALTER COLUMN id SET DEFAULT nextval('public.staff_property_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: warehouses id; Type: DEFAULT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.warehouses ALTER COLUMN id SET DEFAULT nextval('public.warehouses_id_seq'::regclass);


--
-- Name: fixed_property Fixed_property_pkey; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.fixed_property
    ADD CONSTRAINT "Fixed_property_pkey" PRIMARY KEY (id);


--
-- Name: actions actions_action_key; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.actions
    ADD CONSTRAINT actions_action_key UNIQUE (action);


--
-- Name: actions actions_pkey; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.actions
    ADD CONSTRAINT actions_pkey PRIMARY KEY (id);


--
-- Name: fixed_property fixed_property_serial_key; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.fixed_property
    ADD CONSTRAINT fixed_property_serial_key UNIQUE (serial);


--
-- Name: groups groups_name_key; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_name_key UNIQUE (name);


--
-- Name: groups groups_pkey; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);


--
-- Name: property_groups property_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.property_groups
    ADD CONSTRAINT property_groups_pkey PRIMARY KEY (id);


--
-- Name: property_history property_history_pkey; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.property_history
    ADD CONSTRAINT property_history_pkey PRIMARY KEY (id);


--
-- Name: refresh_tokens refresh_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_pkey PRIMARY KEY (id);


--
-- Name: refresh_tokens refresh_tokens_token_key; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_token_key UNIQUE (token);


--
-- Name: refresh_tokens refresh_tokens_username_key; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.refresh_tokens
    ADD CONSTRAINT refresh_tokens_username_key UNIQUE (username);


--
-- Name: roles roles_name_key; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_name_key UNIQUE (name);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: staff staff_pkey; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.staff
    ADD CONSTRAINT staff_pkey PRIMARY KEY (id);


--
-- Name: staff_property staff_property_pkey; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.staff_property
    ADD CONSTRAINT staff_property_pkey PRIMARY KEY (id);


--
-- Name: staff staff_table_key; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.staff
    ADD CONSTRAINT staff_table_key UNIQUE ("table");


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: warehouses warehouses_pkey; Type: CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.warehouses
    ADD CONSTRAINT warehouses_pkey PRIMARY KEY (id);


--
-- Name: fixed_property fixed_property_state_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.fixed_property
    ADD CONSTRAINT fixed_property_state_fkey FOREIGN KEY (state) REFERENCES public.actions(id) ON DELETE SET DEFAULT;


--
-- Name: fixed_property fixed_property_warehouse_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.fixed_property
    ADD CONSTRAINT fixed_property_warehouse_fkey FOREIGN KEY (warehouse) REFERENCES public.warehouses(id);


--
-- Name: property_groups property_groups_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.property_groups
    ADD CONSTRAINT property_groups_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(id) ON DELETE CASCADE;


--
-- Name: property_groups property_groups_property_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.property_groups
    ADD CONSTRAINT property_groups_property_id_fkey FOREIGN KEY (property_id) REFERENCES public.fixed_property(id) ON DELETE CASCADE;


--
-- Name: property_history property_history_action_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.property_history
    ADD CONSTRAINT property_history_action_fkey FOREIGN KEY (action) REFERENCES public.actions(id);


--
-- Name: property_history property_history_property_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.property_history
    ADD CONSTRAINT property_history_property_id_fkey FOREIGN KEY (property_id) REFERENCES public.fixed_property(id) ON DELETE CASCADE;


--
-- Name: property_history property_history_username_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.property_history
    ADD CONSTRAINT property_history_username_fkey FOREIGN KEY (username) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: staff_property staff_property_property_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.staff_property
    ADD CONSTRAINT staff_property_property_id_fkey FOREIGN KEY (property_id) REFERENCES public.fixed_property(id) ON DELETE CASCADE;


--
-- Name: staff_property staff_property_staff_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.staff_property
    ADD CONSTRAINT staff_property_staff_id_fkey FOREIGN KEY (staff_id) REFERENCES public.staff(id) ON DELETE CASCADE;


--
-- Name: staff_property staff_property_username_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.staff_property
    ADD CONSTRAINT staff_property_username_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: users users_role_fkey; Type: FK CONSTRAINT; Schema: public; Owner: rtc
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_role_fkey FOREIGN KEY (role) REFERENCES public.roles(id);


--
-- PostgreSQL database dump complete
--

