import React from 'react';
import { NativeBaseProvider } from 'native-base';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import HomeStack from './components/Home/HomePage';
import ExistingStack from './components/Existing/ExistingPage';
import NewPage from './components/New/NewPicture';
import NewSavePage from './components/New/NewSavePage';
import createTables from './src/database/initialise';
import MainNewClassStack from './components/NewClass/MainNewClassStack';
import ExistingClassDates from './components/Existing/ExistingClassDates';

const Stack = createNativeStackNavigator();

createTables();

const App = () => (
  <NativeBaseProvider>
    <NavigationContainer>
      <Stack.Navigator
        screenOptions={{
          headerShown: false,
        }}
      >
        <Stack.Screen
          name="Home"
          component={HomeStack}
        />
        <Stack.Screen
          name="New"
          component={NewPage}
        />
        <Stack.Screen
          name="NewSave"
          component={NewSavePage}
        />
        <Stack.Screen
          name="Existing"
          component={ExistingStack}
        />
        <Stack.Screen
          name="ExistingClass"
          component={ExistingClassDates}
        />
        <Stack.Screen
          name="NewClass"
          component={MainNewClassStack}
        />
      </Stack.Navigator>
    </NavigationContainer>
  </NativeBaseProvider>
);

export default App;
