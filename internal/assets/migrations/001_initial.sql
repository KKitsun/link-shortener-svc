-- +migrate Up

CREATE TABLE link
(
	id SERIAL PRIMARY KEY,
	alias TEXT NOT NULL UNIQUE,
	url TEXT NOT NULL
);
CREATE INDEX link_index ON link(alias);

-- +migrate Down
DROP TABLE link;
