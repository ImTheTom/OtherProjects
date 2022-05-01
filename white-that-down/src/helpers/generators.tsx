import { AlbumName } from '../types/type';

const generateNameForClass = (className: string): string => {
  const todayDate = new Date().toISOString().slice(0, 10);
  return `${className}_${todayDate}_${Date.now()}`;
};

const generateAlbumNameForClass = (className: string): string => (`${AlbumName}_${className}`);

export {
  generateNameForClass,
  generateAlbumNameForClass,
};
