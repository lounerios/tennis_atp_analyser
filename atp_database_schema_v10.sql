CREATE TABLE player (
        id INTEGER PRIMARY KEY,
        firstname VARCHAR(128) NOT NULL,
        lastname VARCHAR(128) NOT NULL,
        status CHAR(1) NOT NULL,
        birthdate CHAR(8) NOT NULL,
        country CHAR(3) NOT NULL
    );
CREATE TABLE atp_ranking(
        date CHAR(8) NOT NULL,
        number INTEGER NOT NULL,
        player_id INTEGER NOT NULL,
        points INTEGER NOT NULl
    );
