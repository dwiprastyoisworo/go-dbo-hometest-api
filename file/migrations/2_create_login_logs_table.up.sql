CREATE TABLE "login_logs" (
      "id" uuid PRIMARY KEY,
      "user_id" uuid,
      "ip_address" varchar,
      "user_agent" text,
      "login_time" timestamp DEFAULT (now())
);

ALTER TABLE "login_logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");