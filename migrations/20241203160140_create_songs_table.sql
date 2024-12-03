-- +goose Up
CREATE TABLE songs (
                       songid SERIAL PRIMARY KEY,
                       "group" VARCHAR(255) NOT NULL,
                       song VARCHAR(255) NOT NULL,
                       release_date VARCHAR(50),
                       text TEXT,
                       link VARCHAR(255)
);

INSERT INTO songs ("group", song, release_date, text, link) VALUES
    ('testgroup1', 'testsong', '2000', 'test verdes\n\n test verdes2 \n\n test verdes3', 'http://example.com/'),
('testgroup2', 'testsong', '2000', 'some text', 'http://example.com/');


-- +goose Down
DELETE FROM users WHERE song = 'testsong' IF EXISTS;

DROP TABLE IF EXISTS songs;
