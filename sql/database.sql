--
-- PostgreSQL database dump
--

-- Dumped from database version 13.2 (Debian 13.2-1.pgdg100+1)
-- Dumped by pg_dump version 13.2 (Debian 13.2-1.pgdg100+1)

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
-- Name: symbols; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.symbols (
    symbol character varying(16) NOT NULL,
    currency character varying(3) NOT NULL
);


ALTER TABLE public.symbols OWNER TO postgres;

--
-- Name: values; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."values" (
    symbol character varying(16) NOT NULL,
    value real NOT NULL,
    "timestamp" timestamp without time zone NOT NULL
);


ALTER TABLE public."values" OWNER TO postgres;

--
-- Name: symbols stocks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.symbols
    ADD CONSTRAINT stocks_pkey PRIMARY KEY (symbol);


--
-- Name: values values_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."values"
    ADD CONSTRAINT values_pkey PRIMARY KEY (symbol);


--
-- PostgreSQL database dump complete
--

