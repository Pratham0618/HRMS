import React from 'react'
import PopoverDemo from './Notification';

export default function Header(){
    return(
        <div className='bg-slate-600 h-14 p-4 text-white flex justify-between items-center'>
            Employee Leave Management System
            <PopoverDemo />
        </div>
    )
}