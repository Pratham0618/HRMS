import React from 'react'
import TablesDept from './shared/Department-Table'
import DepartmentAdd from './Add-Department'

export default function Departmentsection(){
    return (
        <div>
            <p style={{ marginBottom: '20px', fontSize: '1.5rem', color: 'slate-800' }}>Department List</p>
            <div style={{ marginBottom: '15px' }}><DepartmentAdd  /></div>
            <TablesDept />
        </div>
    )
}