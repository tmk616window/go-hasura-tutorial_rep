
CREATE DATABASE test_db;

\c test_db

CREATE TABLE "users" (
  "id" serial,
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY(id)
);

CREATE TABLE "priorities" (
  "id" serial,
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE TABLE "labels" (
  "id" serial,
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE TABLE "statuses" (
  "id" serial,
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE TABLE "todos" (
  "id" serial,
  "title" text NOT NULL,
  "description" text,
  "user_id" integer,
  "status_id" integer,
  "priority_id" integer,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "finished_at" timestamptz,
  PRIMARY KEY ("id"),
  FOREIGN KEY ("user_id") REFERENCES users ("id"),
  FOREIGN KEY ("status_id") REFERENCES statuses ("id"),
  FOREIGN KEY ("priority_id") REFERENCES priorities ("id")
);

CREATE TABLE "todos_labels_relation" (
  "id" serial,
  "todo_id" integer,
  "label_id" integer,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  FOREIGN KEY ("todo_id") REFERENCES todos ("id"),
  FOREIGN KEY ("label_id") REFERENCES labels ("id"),
  PRIMARY KEY ("id")
);

insert into users (name) values ('test');
insert into labels 
  (name) 
values 
  ('label1'),
  ('label2'),
  ('label3'),
  ('label4'),
  ('label5'),
  ('label6'),
  ('label7'),
  ('label8'),
  ('label9');
insert into priorities 
  (name) 
values 
  ('高'),
  ('中'),
  ('低');
insert into statuses 
  (name) 
values 
  ('未完'),
  ('実行中'),
  ('完了');
insert into todos (text, status, user_id, status_id, priority_id) values ('test', 'status', 1, 1, 1);
insert into todos_labels_relation 
  (todo_id, label_id) 
values 
  (1, 1),
  (1, 2);
