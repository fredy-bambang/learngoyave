-- +goose Up
-- +goose StatementBegin
SELECT
  'up SQL query';

-- +goose StatementEnd
CREATE TABLE "user" (
  id uuid NOT NULL,
  email character varying(255) UNIQUE NOT NULL,
  password character varying(255) NOT NULL,
  role character varying(255) NOT NULL,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  PRIMARY KEY (id)
);

-- +goose Down
-- +goose StatementBegin
SELECT
  'down SQL query';

-- +goose StatementEnd
DROP TABLE IF EXISTS "user";