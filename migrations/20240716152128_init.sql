-- +goose Up
-- +goose StatementBegin
create table public.test_tables (
    id serial primary key,
    name text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
  drop table public.test_tables;
-- +goose StatementEnd
