-- +goose Up
-- +goose StatementBegin
create table test_table (
    id serial primary key,
    name text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
  drop table test_table;
-- +goose StatementEnd
