package db

const create string = `
  CREATE TABLE IF NOT EXISTS entries (
  	id INTEGER NOT NULL PRIMARY KEY,
  	urge INTEGER NOT NULL,
  	need INTEGER NOT NULL,
  	time DATETIME NOT NULL
  );`
