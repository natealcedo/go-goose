-- +goose Up
-- +goose StatementBegin
create table posts(
  id uuid default gen_random_uuid() primary key,
  title text not null,
  content text not null,
  created_at timestamptz  default now(),
  updated_at timestamptz  default now()
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
drop table posts;
-- +goose StatementEnd
