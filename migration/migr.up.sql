CREATE TABLE [IF NOT EXISTS] "customers" (
    "id" integer PRIMARY KEY,
    "first_name" varying character(50),
    "last_name" varying character(50),
    "phone_number" varying character(20),
    "password" varying character(100)
);

CREATE TABLE [IF NOT EXISTS] "employees" (
    "id" integer PRIMARY KEY,
    "first_name" varying character(50),
    "last_name" varying character(50),
    "position" varying character(100),
    "email" varying character(100),
    "phone_number" varying character(20),
    "login" varying character(100),
    "password" varying character(100)
);

CREATE TABLE [IF NOT EXISTS] "categories" (
    "id" integer PRIMARY KEY,
    "name" varying character(100)
);

CREATE TABLE [IF NOT EXISTS] "products" (
    "id" integer PRIMARY KEY,
    "name" varying character(100),
    "price" double precision,
    "category_id" integer
);

CREATE TABLE [IF NOT EXISTS] "korzina" (
    "customer_id" integer,
    "product_id" integer,
    "count" integer
);

CREATE TABLE [IF NOT EXISTS] "purchases" (
    "purchase_id" integer PRIMARY KEY,
    "customer_id" integer,
    "product_id" integer,
    "date" timestamp without time zone
);

CREATE TABLE [IF NOT EXISTS] "store" (
    "product_id" integer,
    "count" integer,
    "open_flg" boolean
);

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories"("id");
ALTER TABLE "korzina" ADD FOREIGN KEY ("customer_id") REFERENCES "customers"("id");
ALTER TABLE "korzina" ADD FOREIGN KEY ("product_id") REFERENCES "products"("id");
ALTER TABLE "purchases" ADD FOREIGN KEY ("customer_id") REFERENCES "customers"("id");
ALTER TABLE "purchases" ADD FOREIGN KEY ("product_id") REFERENCES "products"("id");
ALTER TABLE "store" ADD FOREIGN KEY ("product_id") REFERENCES "products"("id");

