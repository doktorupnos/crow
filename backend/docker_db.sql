--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

-- Started on 2023-11-30 21:25:37 EET

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

DROP DATABASE IF EXISTS crow;
--
-- TOC entry 4345 (class 1262 OID 17186)
-- Name: crow; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE crow WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.UTF-8';


ALTER DATABASE crow OWNER TO postgres;

\connect crow

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
-- TOC entry 216 (class 1259 OID 17197)
-- Name: posts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.posts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    body text NOT NULL,
    user_id uuid
);


ALTER TABLE public.posts OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 17187)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name text NOT NULL,
    password text NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 4339 (class 0 OID 17197)
-- Dependencies: 216
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.posts VALUES ('efb04cb9-1e74-4823-b0b8-f143b3824637', '2023-11-30 21:18:04.962478+02', '2023-11-30 21:18:04.962478+02', 'first post by zoumas', '5232d6a9-11fd-49b4-814d-33efbc0c3359');
INSERT INTO public.posts VALUES ('617c018b-f964-49c1-b960-ed06a5f52bf7', '2023-11-30 21:22:09.282925+02', '2023-11-30 21:22:09.282925+02', 'In Go''s silent dance, code whispers, Efficiency blooms, simplicity glisten.', '5232d6a9-11fd-49b4-814d-33efbc0c3359');
INSERT INTO public.posts VALUES ('831a31c6-9a3e-443c-b26e-fe52834026d4', '2023-11-30 21:23:39.126678+02', '2023-11-30 21:23:39.126678+02', 'Hohoho', '5232d6a9-11fd-49b4-814d-33efbc0c3359');


--
-- TOC entry 4338 (class 0 OID 17187)
-- Dependencies: 215
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES ('5232d6a9-11fd-49b4-814d-33efbc0c3359', '2023-11-30 21:17:13.572438+02', '2023-11-30 21:17:13.572438+02', 'zoumas', '$2a$10$Fld0WoOtOrSWkwGRdVvpC.rffAJVQ16Nk.AK3wLiXEMDmRPmT2AcG');


--
-- TOC entry 4193 (class 2606 OID 17204)
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- TOC entry 4189 (class 2606 OID 17196)
-- Name: users users_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_name_key UNIQUE (name);


--
-- TOC entry 4191 (class 2606 OID 17194)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4194 (class 2606 OID 17205)
-- Name: posts fk_posts_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT fk_posts_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


-- Completed on 2023-11-30 21:25:38 EET

--
-- PostgreSQL database dump complete
--

