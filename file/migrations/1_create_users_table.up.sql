CREATE TABLE "users" (
     "id" uuid PRIMARY KEY,
     "username" varchar UNIQUE NOT NULL,
     "email" varchar UNIQUE NOT NULL,
     "password_hash" varchar NOT NULL,
     "role" varchar,
     "created_at" timestamp DEFAULT (now()),
     "updated_at" timestamp
);