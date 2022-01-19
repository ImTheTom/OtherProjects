interface ClassDAO {
  active: number,
  created: string,
  id: number,
  name: string;
}

interface PictureDAO {
  class_id: number,
  created: string,
  location: string,
  id: number,
  name: string;
}

const AlbumName = 'WhiteThatDown';

const PermissionsGranted = 'granted';

export {
  ClassDAO,
  PictureDAO,
  AlbumName,
  PermissionsGranted,
};
