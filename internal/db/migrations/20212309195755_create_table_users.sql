-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS users CASCADE;

create table users(
      user_id SERIAL PRIMARY KEY,
      name varchar(255) NOT NULL,
      email varchar(50) UNIQUE NOT NULL,
      password CHAR(82) NOT NULL,
      created_at timestamp WITHOUT TIME ZONE,
      updated_at timestamp WITHOUT TIME ZONE,
      deleted_at timestamp WITHOUT TIME ZONE
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table users;