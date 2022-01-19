import React from 'react';
import { VStack, Heading, Button } from 'native-base';
import { NavigationProp, ParamListBase } from '@react-navigation/native';
import { fetchClasses } from '../../src/database/get';
import { ClassDAO } from '../../src/types/type';

interface Props {
  navigation: NavigationProp<ParamListBase>
}

const ExistingStack = ({ navigation }: Props) => {
  const [yourClasses, setYourClasses] = React.useState([] as JSX.Element[]);
  const [fetchedClasses, setFetchedClasses] = React.useState([] as ClassDAO[]);

  const generateExistingClassesButtons = (classes: ClassDAO[]) => {
    let buttons: JSX.Element[] = [];
    buttons = classes.map((value: ClassDAO) => (
      <Button
        key={value.id}
        colorScheme="primary"
        onPress={() => navigation.navigate('ExistingClass', {
          selectedClass: value,
        })}
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

  return (
    <VStack space={4} alignItems="center" style={{ top: 100 }}>
      <Heading textAlign="center">
        Your Classes
      </Heading>
      {yourClasses}
    </VStack>
  );
};

export default ExistingStack;
