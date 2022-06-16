-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE "task" (
	id uuid NOT NULL,
  title character varying(255),
  is_finished boolean,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
	PRIMARY KEY (id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS "task"
