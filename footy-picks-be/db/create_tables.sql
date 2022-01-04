CREATE TABLE player (
  player_id serial PRIMARY KEY,
  email varchar(100) NOT NULL,
  player_name varchar(150) NOT NULL,
  player_login varchar(150) NOT NULL,
  created timestamp(6) NOT NULL,
  games integer[] NOT NULL
);

CREATE TABLE survivor_game (
  game_id serial PRIMARY KEY,
  game_name varchar(150) NOT NULL,
  passcode varchar(10) UNIQUE,
  league_id integer NOT NULL,
  league varchar(150) NOT NULL,
  ongoing boolean NOT NULL,
  beginning_round integer NOT NULL,
  created timestamp(6) NOT NULL,
  creator integer REFERENCES player,
  players integer[] NOT NULL
);

CREATE TABLE survivor_pick (
  game_round integer NOT NULL,
  pick varchar(10) NOT NULL,
  correct boolean NOT NULL,
  game integer REFERENCES survivor_game,
  player integer REFERENCES player,
  PRIMARY KEY (game, game_round, player)
);
