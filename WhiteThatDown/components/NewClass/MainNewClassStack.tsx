import React from 'react';
import {
  VStack,
  Heading,
  Input,
  FormControl,
  Button,
} from 'native-base';
import { NavigationProp, ParamListBase } from '@react-navigation/native';
import { createClass } from '../../src/database/create';

interface Props {
  navigation: NavigationProp<ParamListBase>
}

const MainNewClassStack = ({ navigation }: Props) => {
  const [newClassName, setNewClassName] = React.useState('');
  const [isFormInvalid, setIsFormInvalid] = React.useState(false);

  const handleSubmit = () => {
    try {
      setIsFormInvalid(false);
      createClass(newClassName);
      navigation.navigate('Existing');
    } catch (err) {
      setIsFormInvalid(true);
    }
  };

  return (
    <VStack space={4} alignItems="center" style={{ top: 100 }}>
      <Heading textAlign="center">
        Create a new class
      </Heading>
      <FormControl
        w={{
          base: '75%',
          md: '25%',
        }}
        isInvalid={isFormInvalid}
      >
        <FormControl.Label>Class Name</FormControl.Label>
        <Input
          size="lg"
          value={newClassName}
          onChangeText={(text: string) => setNewClassName(text)}
        />
        <FormControl.ErrorMessage>Error creating the class.</FormControl.ErrorMessage>
      </FormControl>
      <Button onPress={() => handleSubmit()}>Submit</Button>
    </VStack>
  );
};

export default MainNewClassStack;
