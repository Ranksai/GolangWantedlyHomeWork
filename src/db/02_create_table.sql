--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

--
-- Name: set_update_time(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION set_update_time() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
  begin
    new.updated_at := 'now';
    return new;
  end;
$$;


ALTER FUNCTION public.set_update_time() OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: wantedlyusers; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE wantedlyhomework.wantedlyusers (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.wantedlyusers OWNER TO postgres;

--
-- Name: wantedlyusers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE wantedlyusers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wantedlyusers_id_seq OWNER TO postgres;

--
-- Name: wantedlyusers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE wantedlyusers_id_seq OWNED BY wantedlyusers.id;

--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY wantedlyusers ALTER COLUMN id SET DEFAULT nextval('wantedlyusers_id_seq'::regclass);

--
-- Data for Name: wantedlyusers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY wantedlyusers (id, name, email, created_at, updated_at) FROM stdin;
10	test	test@test	2017-11-11 17:19:32.324046	2017-11-11 17:19:32.324046
2	test	hogehogehoghogehogeohgeoe@example.com	2017-11-11 17:34:18.915196	2017-11-11 17:34:18.915196
4	ssss	hogeh@example.com	2017-11-11 17:39:35.760281	2017-11-11 17:39:35.760281
\.


--
-- Name: wantedlyusers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('wantedlyusers_id_seq', 4, true);


--
-- Name: wantedlyusers_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY wantedlyusers
    ADD CONSTRAINT wantedlyusers_email_key UNIQUE (email);


--
-- Name: wantedlyusers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY wantedlyusers
    ADD CONSTRAINT wantedlyusers_pkey PRIMARY KEY (id);


--
-- Name: update_tri; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER update_tri BEFORE UPDATE ON wantedlyusers FOR EACH ROW EXECUTE PROCEDURE set_update_time();


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

