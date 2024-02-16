CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "employees" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" VARCHAR(100),
  "email" VARCHAR(50),
  "username" VARCHAR(50),
  "password" VARCHAR(100),
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP
);

CREATE TABLE "products" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" VARCHAR(100),
  "price" BIGINT,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP
);

CREATE TABLE "customers" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" VARCHAR(100),
  "phone" VARCHAR(16),
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP
);

CREATE TABLE "bills" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "bill_date" DATE,
  "customer_id" UUID,
  "employee_id" UUID,
  "total_price" BIGINT,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP
);

CREATE TABLE "bill_details" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "bill_id" UUID,
  "product_id" UUID,
  "quantity" int,
  "sub_total" BIGINT,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" TIMESTAMP
);

ALTER TABLE "bills" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "bills" ADD FOREIGN KEY ("employee_id") REFERENCES "employees" ("id");

ALTER TABLE "bill_details" ADD FOREIGN KEY ("bill_id") REFERENCES "bills" ("id");

ALTER TABLE "bill_details" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");
