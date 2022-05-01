import React, { useEffect } from 'react';
import {
  VStack, Heading, Button, Image,
} from 'native-base';
import * as MediaLibrary from 'expo-media-library';
import { NavigationProp, RouteProp, ParamListBase } from '@react-navigation/native';
import { ImagePickerResult } from 'expo-image-picker';
import { ImageInfo } from 'expo-image-picker/build/ImagePicker.types';
import { fetchClasses } from '../../src/database/get';
import { ClassDAO, PermissionsGranted } from '../../src/types/type';
import { createPhoto } from '../../src/database/create';
import { generateAlbumNameForClass } from '../../src/helpers/generators';

interface Props {
  route: RouteProp<{ params: { result: ImagePickerResult & ImageInfo } }, 'params'>
  navigation: NavigationProp<ParamListBase>
}

const NewSavePage = ({ route, navigation }: Props) => {
  const [yourClasses, setYourClasses] = React.useState([] as JSX.Element[]);
  const [fetchedClasses, setFetchedClasses] = React.useState([] as ClassDAO[]);

  const { result } = route.params;

  if (!result) {
    navigation.navigate('Home');
  }

  useEffect(() => {
    if (result.cancelled) {
      navigation.navigate('Home');
    }
  }, []);

  const saveFileAndProgress = async (selectedClass: ClassDAO) => {
    // Save file...
    const { status } = await MediaLibrary.requestPermissionsAsync();
    const ok = status === PermissionsGranted;
    if (ok) {
      const albumName = generateAlbumNameForClass(selectedClass.name);
      const cachedAsset = await MediaLibrary.createAssetAsync(result.uri);
      const album = await MediaLibrary.getAlbumAsync(albumName);

      if (album) {
        await MediaLibrary.addAssetsToAlbumAsync([cachedAsset], album, false);
      } else {
        await MediaLibrary.createAlbumAsync(albumName, cachedAsset);
      }

      await createPhoto(cachedAsset.uri, selectedClass);
      navigation.navigate('ExistingClass', {
        selectedClass,
      });
    }
  };

  const generateExistingClassesButtons = (classes: ClassDAO[]) => {
    let buttons: JSX.Element[] = [];
    buttons = classes.map((value: ClassDAO) => (
      <Button
        key={value.id}
        colorScheme="secondary"
        onPress={() => saveFileAndProgress(value)}
      >
        {value.name}
      </Button>
    ));
    setYourClasses(buttons);
  };

  React.useEffect(() => {
    fetchClasses(setFetchedClasses);
  }, []);

  React.useEffect(() => {
    generateExistingClassesButtons(fetchedClasses);
  }, [fetchedClasses]);

  const displayImage = (
    <Image
      source={{
        uri: result.uri,
      }}
      alt="temp class image"
      size="xl"
    />
  );

  return (
    <VStack space={4} alignItems="center" style={{ top: 100 }}>
      <Heading textAlign="center">
        Where to Save New Photo?
      </Heading>
      {displayImage}
      {yourClasses}
      <Button colorScheme="primary" onPress={() => navigation.navigate('New')}>Back</Button>
    </VStack>
  );
};

export default NewSavePage;
