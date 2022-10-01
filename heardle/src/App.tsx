import React from 'react'
import Guesses from './component/guesses';
import Header from './component/header';
import Footer from './component/footer';

function App() {
  const [guesses, SetNewGuess] = React.useState(['','','','','','']);
  const [currentGuess, SetCurrentGuess] = React.useState(0);
  const [isOver, SetIsOver] = React.useState(false);

  if (!isOver) {
    if (currentGuess >= 5) {
      SetIsOver(true)
    }
  }

  return (
    <div className='App'>
      <Header />
      <Guesses guesses={guesses}/>
      <Footer guesses={guesses} currentGuess={currentGuess} SetNewGuess={SetNewGuess} SetCurrentGuess={SetCurrentGuess}/>
    </div>
  )
}

export default App
