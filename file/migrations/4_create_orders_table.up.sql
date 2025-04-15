CREATE TABLE "orders" (
      "id" uuid PRIMARY KEY,
      "customer_id" uuid,
      "order_date" timestamp DEFAULT (now()),
      "status" varchar,
      "total_amount" decimal,
      "created_at" timestamp,
      "updated_at" timestamp
);

ALTER TABLE "orders" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");