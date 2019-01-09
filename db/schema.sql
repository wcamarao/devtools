create table "entity_sample" (
  "id" text primary key,
  "is_active" boolean not null,
  "created_at" timestamptz not null,
  "updated_at" timestamptz not null
);
