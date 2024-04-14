import { set } from 'ol/transform';
import React from 'react'
import { useState } from 'react';
import { CardTwo } from './CardTwo';

export function CardOne() {

  const [counter, setCounter] = useState(15)

  const addValue = () => {
    console.log("Button clicked ", counter+Math.random())
    setCounter(counter+1)
  } 

  const removeValue = () => {
    setCounter(counter-1)
  }

  return (
    <>
      <h1>Text 1: {counter} </h1>
      <h2>Text {counter*10} </h2>
      <button className='rounded-lg text-pretty m-4 bg-gray-600' onClick={addValue}>Add Value {counter} </button>
      <br/>
      <button  className='rounded-lg text-pretty m-4 bg-gray-600' onClick={removeValue}>Remove Value {counter}</button>
      <p>Footer : {counter}</p>
      <br/>
      <div>
        <section>
          <CardTwo counter={counter}/>
        </section>

      </div>
    </>
    )
}



  