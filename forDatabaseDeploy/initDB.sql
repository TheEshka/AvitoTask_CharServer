CREATE DATABASE chat_db;

\c chat_db;

CREATE TABLE IF NOT EXISTS  users
(
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS  chats
(
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS  chat_users
(
  chat_id INTEGER NOT NULL REFERENCES chats (id),
  user_id INTEGER NOT NULL REFERENCES users (id),
  PRIMARY KEY(chat_id, user_id)
);

CREATE TABLE IF NOT EXISTS  messages
(
  id SERIAL PRIMARY KEY,
  chat_id INTEGER NOT NULL,
  author_id INTEGER NOT NULL,
  text VARCHAR(500) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (chat_id, author_id) REFERENCES chat_users (chat_id, user_id)
);