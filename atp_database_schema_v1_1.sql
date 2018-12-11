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
CREATE INDEX plr_country on player (country);
CREATE INDEX plr_status on player (status);
CREATE INDEX rnk_date on atp_ranking(date);
CREATE INDEX rnk_number on atp_ranking(number);
CREATE INDEX rnk_player_id on atp_ranking(player_id);
