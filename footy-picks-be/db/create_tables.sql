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
  passcode varchar(10),
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

/* TODO: Remove mock data */
/* players */
INSERT INTO player (player_id, email, player_name, player_login, created, games) 
  VALUES (1, 'mattaquiles@gmail.com', 'Matthew', 'fakelogin', '2021-12-28', '{1,2}');
INSERT INTO player (player_id, email, player_name, player_login, created, games) 
  VALUES (2, 'player2@gmail.com', 'Player 2', 'fakelogin', '2021-12-28', '{1,2}');
INSERT INTO player (player_id, email, player_name, player_login, created, games) 
  VALUES (3, 'plqyer3@gmail.com', 'Player 3', 'fakelogin', '2021-12-28', '{1,2}');
INSERT INTO player (player_id, email, player_name, player_login, created, games) 
  VALUES (4, 'player4@gmail.com', 'Player 4', 'fakelogin', '2021-12-28', '{1,2}');
INSERT INTO player (player_id, email, player_name, player_login, created, games) 
  VALUES (5, 'player5@gmail.com', 'Player 5', 'fakelogin', '2021-12-28', '{1,2}');
INSERT INTO player (player_id, email, player_name, player_login, created, games) 
  VALUES (6, 'player6@gmail.com', 'Player 6', 'fakelogin', '2021-12-28', '{1,2}');
INSERT INTO player (player_id, email, player_name, player_login, created, games) 
  VALUES (7, 'player7@gmail.com', 'Player 7', 'fakelogin', '2021-12-28', '{1,2}');
INSERT INTO player (player_id, email, player_name, player_login, created, games) 
  VALUES (8, 'player8@gmail.com', 'Player 8', 'fakelogin', '2021-12-28', '{1,2}');
INSERT INTO player (player_id, email, player_name, player_login, created, games) 
  VALUES (9, 'player9@gmail.com', 'Player 9', 'fakelogin', '2021-12-28', '{1,2}');
INSERT INTO player (player_id, email, player_name, player_login, created, games) 
  VALUES (10, 'player10@gmail.com', 'Player 10', 'fakelogin', '2021-12-28', '{1,2}');

/* survivor_games */


/* survivor_picks */
