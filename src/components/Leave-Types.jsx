import React from 'react'
import LeaveTypesAdd from './Add-LeaveTypes'
import TablesLeaveTypes from './shared/LeaveTypes-Table'

export default function LeaveTypes(){
    return (
        <div>
            <p style={{ marginBottom: '20px', fontSize: '1.5rem', color: 'slate-800' }}>Leave Types</p>
            <div style={{ marginBottom: '15px' }}><LeaveTypesAdd  /></div>
            <TablesLeaveTypes />
        </div>
    )
}