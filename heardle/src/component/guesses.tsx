import React from 'react'

type GuessesProps = {
  guesses: string[]
}

function Guesses({guesses}: GuessesProps) {
  console.log(guesses)

  const rows = guesses.map(
    value => (
      <div key={Math.random().toString(36).substr(2, 9)} className="guess">
        {value}
      </div>
    )
  )

  return (
    <main className='content'>
      {rows}
    </main>
  )
}

export default Guesses
