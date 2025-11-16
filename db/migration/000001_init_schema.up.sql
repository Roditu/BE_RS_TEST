CREATE TABLE "person" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "ambition" varchar
);

CREATE TABLE "user" (
  "user_id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "exp" bigint NOT NULL DEFAULT 0,
  "level" bigint NOT NULL DEFAULT 1
);

CREATE TABLE "task" (
  "task_id" bigserial PRIMARY KEY,
  "todo" varchar NOT NULL,
  "exp" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "status" varchar NOT NULL DEFAULT 'UNFINISHED'
);

ALTER TABLE "task" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");
