-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books (
  id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  title varchar(255) NOT NULL UNIQUE,
  description text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
