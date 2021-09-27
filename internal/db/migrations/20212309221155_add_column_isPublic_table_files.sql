-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
alter table files add column is_public boolean default false;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
alter table files drop column is_public boolean default false;