import React from 'react';
import {
  Text,
  View,
} from 'react-native';
import * as ImagePicker from 'expo-image-picker';
import { NavigationProp, ParamListBase } from '@react-navigation/native';
import { PermissionsGranted } from '../../src/types/type';

interface Props {
  navigation: NavigationProp<ParamListBase>
}

const NewPage = ({ navigation }: Props) => {
  const onHandlePermission = async () => {
    const { status } = await ImagePicker.requestCameraPermissionsAsync();
    const ok = status === PermissionsGranted;
    if (ok) {
      const result = await ImagePicker.launchCameraAsync();
      navigation.navigate('NewSave', {
        result,
      });
    }
  };

  React.useEffect(() => {
    const unsubscribe = navigation.addListener('focus', () => {
      onHandlePermission();
    });

    return unsubscribe;
  }, [navigation]);

  return (
    <View>
      <Text>Opening camera now...</Text>
    </View>
  );
};

export default NewPage;
