CREATE TABLE "customers" (
     "id" uuid PRIMARY KEY,
     "name" varchar NOT NULL,
     "email" varchar UNIQUE,
     "phone" varchar,
     "address" text,
     "created_at" timestamp DEFAULT (now()),
     "updated_at" timestamp
);