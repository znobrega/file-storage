-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS files;

create table files(
       file_id UUID PRIMARY KEY UNIQUE,
       name varchar(255) NOT NULL,
       path varchar(255) NOT NULL UNIQUE,
       directory varchar(255) NOT NULL,
       extension varchar(10) NOT NULL,
       file_size varchar (100),
       user_id int NOT NULL,
       created_at timestamp WITHOUT TIME ZONE,
       updated_at timestamp WITHOUT TIME ZONE,
       deleted_at timestamp WITHOUT TIME ZONE,
       FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table files;