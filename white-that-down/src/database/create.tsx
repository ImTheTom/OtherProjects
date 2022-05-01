import * as SQLite from 'expo-sqlite';
import { generateNameForClass } from '../helpers/generators';
import { ClassDAO } from '../types/type';
import { GetDB } from './get';

const createClass = (className: string) => {
  const funDB = GetDB();
  const todayDate = new Date().toISOString().slice(0, 10);

  funDB.transaction((tx) => {
    tx.executeSql(
      'INSERT INTO classes (name, created, active) values (?, ?, ?)',
      [className, todayDate, 1],
      undefined,
      (_, err: SQLite.SQLError) => {
        throw new Error(err.message);
      },
    );
  });
};

const createPhoto = async (photoLocation: string, forClass: ClassDAO) => {
  const funDB = GetDB();
  const todayDate = new Date().toISOString().slice(0, 10);
  const photoName = generateNameForClass(forClass.name);

  funDB.transaction((tx) => {
    tx.executeSql(
      'INSERT INTO photos (name, location, created, class_id) values (?, ?, ?, ?)',
      [photoName, photoLocation, todayDate, forClass.id],
      undefined,
      (_, err: SQLite.SQLError) => {
        throw new Error(err.message);
      },
    );
  });
};

export { createClass, createPhoto };
