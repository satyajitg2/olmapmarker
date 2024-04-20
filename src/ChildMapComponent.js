import React from 'react'

const ChildMapComponent = ({location, dataLoc}) => {
  return (
    <>
      <div>ChildMapComponent</div>
      <section>
        <h1 className='bg-yellow-200'>ChildMapComponent Item 1: {location} and Data {dataLoc} </h1>
      </section>
    </>
  )
}

export default ChildMapComponent