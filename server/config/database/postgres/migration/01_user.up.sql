CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "tb_employees" (
  "registry" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" TEXT NOT NULL,
  "email" TEXT UNIQUE NOT NULL,
  "sector" TEXT NOT NULL,
  "unit" TEXT NOT NULL,
  "administrator" BOOLEAN NOT NULL DEFAULT FALSE
);