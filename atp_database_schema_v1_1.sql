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
  points INTEGER NOT NULL
);

CREATE TABLE tournament (
  id INTEGER PRIMARY_KEY,
  name VARCHAR(128)  NOT NULL,
  surface VARCHAR(16) NOT NULL,
  draw_size INTEGER NOT NULL,
  level CHAR(1) NOT NULL,
  date CHAR(8) NOT NULL
);

CREATE INDEX plr_country on player (country);
CREATE INDEX plr_status on player (status);
CREATE INDEX rnk_date on atp_ranking(date);
CREATE INDEX rnk_number on atp_ranking(number);
CREATE INDEX rnk_player_id on atp_ranking(player_id);
