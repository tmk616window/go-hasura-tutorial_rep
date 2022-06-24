CREATE TABLE "public"."statuses" ("id" integer, "name" text, "created_at" timestamptz, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id") , UNIQUE ("id"));
