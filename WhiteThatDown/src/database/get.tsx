import * as SQLite from 'expo-sqlite';
import { ClassDAO } from '../types/type';

let db: SQLite.WebSQLDatabase | undefined;

const GetDB = (): SQLite.WebSQLDatabase => {
  if (!db) {
    db = SQLite.openDatabase('db.testDb');
  }
  return db;
};

const fetchClasses = (setClasses: any) => {
  const funDB = GetDB();
  funDB.transaction((tx) => {
    tx.executeSql(
      'SELECT * FROM classes',
      undefined,
      (_, { rows: { _array } }) => {
        setClasses(_array as ClassDAO[]);
      },
      (txObj, err: SQLite.SQLError) => {
        throw new Error(err.message);
      },
    );
  });
};

const fetchPictures = (selectedClassID: number, setPictures: any) => {
  const funDB = GetDB();
  funDB.transaction((tx) => {
    tx.executeSql(
      'SELECT * FROM photos',
      [],
      (_, { rows: { _array } }) => {
        setPictures(_array as ClassDAO[]);
      },
      (txObj, err: SQLite.SQLError) => {
        throw new Error(err.message);
      },
    );
  });
};

export { GetDB, fetchClasses, fetchPictures };
