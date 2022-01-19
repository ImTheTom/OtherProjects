import React from 'react';
import {
  Button, Center, Heading, Stack,
} from 'native-base';
import { NavigationProp, ParamListBase } from '@react-navigation/native';

interface Props {
  navigation: NavigationProp<ParamListBase>
}

const HomeStack = ({ navigation }: Props) => (
  <Center flex={1} p="4">
    <Heading size="xl">White That Down</Heading>
    <Stack mb="2.5" mt="1.5" direction="column" space="md">
      <Button colorScheme="primary" onPress={() => navigation.navigate('New')}>Take A New Picture</Button>
      <Button colorScheme="secondary" onPress={() => navigation.navigate('Existing')}>View Previous Pictures</Button>
      <Button colorScheme="tertiary" onPress={() => navigation.navigate('NewClass')}>Create A New Class</Button>
    </Stack>
  </Center>
);

export default HomeStack;
