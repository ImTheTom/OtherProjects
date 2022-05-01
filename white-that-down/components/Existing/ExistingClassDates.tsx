import React from 'react';
import { ScrollView } from 'react-native';
import { VStack, Heading, Image } from 'native-base';
import * as MediaLibrary from 'expo-media-library';
import { RouteProp } from '@react-navigation/native';
import { ClassDAO } from '../../src/types/type';
import { generateAlbumNameForClass } from '../../src/helpers/generators';

interface Props {
  route: RouteProp<{ params: { selectedClass: ClassDAO } }, 'params'>,
}

const ExistingClassDates = ({ route }: Props) => {
  const [pictures, setPicture] = React.useState([] as JSX.Element[]);

  const { selectedClass } = route.params;

  const albumName = generateAlbumNameForClass(selectedClass.name);

  const generateExistingPictures = async () => {
    let buttons: JSX.Element[] = [];
    const album = await MediaLibrary.getAlbumAsync(albumName);
    if (!album) {
      return;
    }

    const pagedAsset = await MediaLibrary.getAssetsAsync({
      album,
    });

    if (pagedAsset.totalCount === 0) {
      return;
    }

    buttons = pagedAsset.assets.map((ass) => (
      <Image
        key={ass.id}
        source={{
          uri: ass.uri,
        }}
        alt="class image"
        size="xl"
      />
    ));

    setPicture(buttons);
  };

  React.useEffect(() => {
    generateExistingPictures();
  }, []);

  return (
    <ScrollView style={{ top: 100 }}>
      <Heading textAlign="center">
        Your pictures for&nbsp;
        {selectedClass.name}
      </Heading>
      <VStack space={4} alignItems="center">
        {pictures}
      </VStack>
    </ScrollView>
  );
};

export default ExistingClassDates;
