import React from 'react'

type FooterProps = {
  guesses: string[]
  currentGuess: number
  SetNewGuess: (text: string[]) => void;
  SetCurrentGuess: (value: number) => void;
}

function Footer({guesses, currentGuess, SetNewGuess, SetCurrentGuess}: FooterProps) {
  const [guess, setguess] = React.useState('');

  const submitGuess = () => {
    guesses[currentGuess] = guess
    SetNewGuess(guesses);
    SetCurrentGuess(currentGuess++)
    setguess('');
  }

	return (
    <footer className='footer'>
      <input type="submit" value="Play" />
      <input
        type="text"
        name="guess"
        value={guess}
        onChange={(e) => setguess(e.target.value)}
        placeholder="Name of the song"
      />
      <input type="submit" value="Submit" onClick={submitGuess}/>
      <input type="submit" value="Skip" />
    </footer>
	)
}

export default Footer