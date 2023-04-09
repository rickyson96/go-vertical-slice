-- Write your migrate up statements here

CREATE TABLE users(
       id SERIAL PRIMARY KEY,
       name TEXT NOT NULL
);

---- create above / drop below ----

DROP TABLE users;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
