import React from 'react'
import MapComponent from './MapComponent.js';
export function CardTwo({counter}) {
    console.log("props data ",counter)
    return (
        <>
            <div>
                <MapComponent/>
                <h1 className='bg-gray-800 justify-center text-white p-4 rounded-3xl'> This is CardTWO</h1>
                <h2>CardTwo has props data {counter}</h2>
            
            </div>
        </>
    )
}
