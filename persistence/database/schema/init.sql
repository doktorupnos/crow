-- Stuff to initialize

-- This will only run on testing
CREATE DATABASE IF NOT EXISTS crow;

-- This will only run on testing
CREATE USER IF NOT EXISTS crow WITH
   LOGIN
   PASSWORD 'your_password'; -- Replace 'your_password' with the actual password

-- Rest content including production
GRANT ALL PRIVILEGES ON DATABASE crow TO crow;

-- Adjusting for remote connections (assuming you are connecting from any IP)
ALTER USER crow SET client_encoding TO 'utf8';
ALTER USER crow SET default_transaction_isolation TO 'read committed';
ALTER USER crow SET timezone TO 'UTC';

-- Grant necessary privileges for remote connections
GRANT CONNECT ON DATABASE crow TO crow;
GRANT USAGE ON SCHEMA public TO crow;
