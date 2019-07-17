CREATE DATABASE chat_db;

\c chat_db;

CREATE TABLE IF NOT EXISTS  users
(
  id SERIAL PRIMARY KEY,
  username VARCHAR(255),
  created_at DATE
);

CREATE TABLE IF NOT EXISTS  chats
(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255),
  created_at DATE
);

CREATE TABLE IF NOT EXISTS  chat_users
(
  chat_id INTEGER REFERENCES chats (id),
  user_id INTEGER REFERENCES users (id),
  PRIMARY KEY(chat_id, user_id)
);

CREATE TABLE IF NOT EXISTS  messages
(
  id SERIAL PRIMARY KEY,
  chat_id INTEGER,
  author_id INTEGER,
  text VARCHAR(500),
  created_at DATE,
  FOREIGN KEY (chat_id, author_id) REFERENCES chat_users (chat_id, user_id)
);

INSERT INTO users (username, created_at) VALUES ('gagik', '2009-11-10');
INSERT INTO users (username, created_at) VALUES ('svetu', '2019-11-10');
INSERT INTO users (username, created_at) VALUES ('mixas', '2029-11-10');
INSERT INTO users (username, created_at) VALUES ('mixas', '1029-11-10');

INSERT INTO chats (name, created_at) VALUES ('chat1', '2000-01-10');
INSERT INTO chats (name, created_at) VALUES ('chat2', '2005-01-10');

INSERT INTO chat_users (chat_id, user_id) VALUES (1, 6);
INSERT INTO chat_users (chat_id, user_id) VALUES (1, 7);
INSERT INTO chat_users (chat_id, user_id) VALUES (1, 9);
INSERT INTO chat_users (chat_id, user_id) VALUES (2, 11);
INSERT INTO chat_users (chat_id, user_id) VALUES (2, 7);

INSERT INTO messages (chat_id, author_id, text) VALUES (2, 7, 'privet');
INSERT INTO messages (chat_id, author_id, text) VALUES (2, 11, 'darova');
INSERT INTO messages (chat_id, author_id, text) VALUES (1, 9, 'privet');
INSERT INTO messages (chat_id, author_id, text) VALUES (1, 7, 'pacaniiiiiiii');
INSERT INTO messages (chat_id, author_id, text) VALUES (1, 7, 'nu neeeet');
INSERT INTO messages (chat_id, author_id, text) VALUES (2, 11, 'kak dela');


SELECT * FROM chats INNER JOIN chat_users ON chats.id = chat_users.chat_id INNER JOIN users ON users.id = chat_users.user_id WHERE users.id = 6 ORDER BY chats.created_at;

SELECT (messages.id), (messages.chat_id), (messages.author_id), (messages.text), (messages.created_at) FROM messages INNER JOIN chats ON messages.chat_id = chats.id WHERE chats.id = 2 ORDER BY messages.created_at;