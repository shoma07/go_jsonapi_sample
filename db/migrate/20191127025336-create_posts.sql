-- +migrate Up
CREATE TABLE posts (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  title VARCHAR NOT NULL,
  body TEXT NOT NULL,
  published BOOLEAN NOT NULL DEFAULT 'f',
  created_at TIMESTAMP NOT NULL DEFAULT (datetime(CURRENT_TIMESTAMP, 'utc')),
  updated_at TIMESTAMP NOT NULL DEFAULT (datetime(CURRENT_TIMESTAMP, 'utc'))
);

-- +migrate Down
DROP TABLE posts;
