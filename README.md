# konnect
CREATE DATABASE konnect;
CREATE TYPE user_role AS ENUM ('customer', 'partner');
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";