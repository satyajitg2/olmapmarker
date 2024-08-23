import React from 'react'
import MapComponent from './MapComponent.js';
export function CardTwo({counter}) {
    console.log("props data ",counter)
    return (
        <>
            <div className='box-content h-400 w-400 p-2 border-4 bg-black rounded-3xl'>
                <MapComponent/>
                <div className= 'bg-yellow-700 rounded-3xl'>
                    <h1 className='justify-center text-white p-4'> This is CardTWO</h1>
                    <h2 className='justify-center text-white p-4'>
                        CardTwo has props data set from CardOne counter
                        <p className='bg-amber-400 text-black text-2xl w-9 font-bold'>{counter}</p>
                    </h2>

                </div>
            
            </div>
        </>
    )
}
