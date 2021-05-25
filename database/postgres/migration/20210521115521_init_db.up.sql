CREATE TABLE IF NOT EXISTS "todo" (
  "id" bigserial PRIMARY KEY,
  "name" text NOT NULL,
  "complete" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
