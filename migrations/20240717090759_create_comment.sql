-- +goose Up
-- +goose StatementBegin
create table comments(
  id uuid default gen_random_uuid() primary key,
  post_id uuid not null references posts(id) on delete cascade,
  content text not null,
  created_at timestamptz  default now(),
  updated_at timestamptz  default now()
);

create index  btree on comments(post_id) ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table comments;
-- +goose StatementEnd
