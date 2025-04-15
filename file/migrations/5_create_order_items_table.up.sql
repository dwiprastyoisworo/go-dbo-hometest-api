CREATE TABLE "order_items" (
       "id" uuid PRIMARY KEY,
       "order_id" uuid,
       "product_name" varchar,
       "quantity" int,
       "price" decimal,
       "subtotal" decimal
);

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");