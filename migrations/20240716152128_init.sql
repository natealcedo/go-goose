-- +goose Up
-- +goose StatementBegin
create table public.test_tables (
    id uuid primary key default gen_random_uuid(),
    name text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
  drop table public.test_tables;
-- +goose StatementEnd
