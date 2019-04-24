
CREATE USER lilvue_admin PASSWORD 'test';

CREATE SCHEMA lilvue_schema;

GRANT ALL ON SCHEMA lilvue_schema TO lilvue_admin;
GRANT ALL ON ALL TABLES IN SCHEMA lilvue_schema TO lilvue_admin;

CREATE DATABASE lilvue
   WITH OWNER lilvue_admin 
   TEMPLATE template0
   ENCODING 'SQL_ASCII'
   TABLESPACE  pg_default
   LC_COLLATE  'C'
   LC_CTYPE  'C'
   CONNECTION LIMIT  -1;

