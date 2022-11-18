package db

const createTableQuery string = `
  CREATE TABLE IF NOT EXISTS entries (
  	id INTEGER NOT NULL PRIMARY KEY,
  	urge INTEGER NOT NULL,
  	need INTEGER NOT NULL,
  	time DATETIME NOT NULL
  );`

const createTrackQuery string = `INSERT INTO entries (urge, need, time) VALUES (?, ?, ?);`
