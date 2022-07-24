CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA tenant_one;
CREATE SCHEMA tenant_two;

CREATE TABLE "tenant_one.User"
(
    "Id"        uuid PRIMARY KEY     DEFAULT uuid_generate_v4(),
    "Password"  varchar     NOT NULL,
    "UserName"  varchar     NOT NULL,
    "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
    "UpdatedAt" timestamptz          DEFAULT (null)
);

CREATE TABLE "tenant_two.User"
(
    "Id"        uuid PRIMARY KEY     DEFAULT uuid_generate_v4(),
    "Password"  varchar     NOT NULL,
    "UserName"  varchar     NOT NULL,
    "CreatedAt" timestamptz NOT NULL DEFAULT (now()),
    "UpdatedAt" timestamptz          DEFAULT (null)
);
