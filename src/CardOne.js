import { set } from 'ol/transform';
import React from 'react'
import { useState } from 'react';
import { CardTwo } from './CardTwo';
import { PostCard } from './PostCard';

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
      <PostCard/>
      <h1>CardOne Text 1: {counter} </h1>
      <h2>CardOne Text 2 {counter*10} </h2>
      <button className='rounded-lg text-pretty m-4 bg-gray-600' onClick={addValue}>Add Value {counter} </button>
      <br/>
      <button  className='rounded-lg text-pretty m-4 bg-gray-600' onClick={removeValue}>Remove Value {counter}</button>
      <p>CardOne Footer : {counter}</p>
      <br/>
      <div>
        <section>
          <CardTwo counter={counter}/>
          <PostCard/>
        </section>

      </div>
    </>
    )
}



  