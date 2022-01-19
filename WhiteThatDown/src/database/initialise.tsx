import * as SQLite from 'expo-sqlite';
import { GetDB } from './get';

const createTables = () => {
  const db = GetDB();

  console.log('creating classes');
  db.transaction((tx: SQLite.SQLTransaction) => {
    tx.executeSql(
      'CREATE TABLE IF NOT EXISTS classes (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, created TEXT, active INTEGER);',
      undefined,
      undefined,
      (_, err: SQLite.SQLError) => {
        throw new Error(err.message);
      },
    );
  });

  console.log('creating photos');
  db.transaction((tx: SQLite.SQLTransaction) => {
    tx.executeSql(
      'CREATE TABLE IF NOT EXISTS photos (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, location TEXT, created TEXT, class_id INTEGER);',
      undefined,
      undefined,
      (_, err: SQLite.SQLError) => {
        throw new Error(err.message);
      },
    );
  });
};

export default createTables;
