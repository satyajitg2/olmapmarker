import React from 'react'
import MapComponent from './MapComponent.js';
export function CardTwo({counter}) {
    console.log("props data ",counter)
    return (
        <>
            <div>
                <h0 className='text-center bg-yellow-300 text-2xl p-2 rounded-2xl'>MAP COMPONENT</h0>
                <MapComponent/>
                <h1 className='bg-gray-800 justify-center text-white p-4 rounded-3xl'> This is CardTWO</h1>
                <h2>CardTwo has props data {counter}</h2>
            
            </div>
        </>
    )
}
