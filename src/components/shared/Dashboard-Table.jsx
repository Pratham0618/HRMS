import React from 'react';
import { Table } from 'antd';

const columns_dashboard = [
  {
    title: 'Name',
    dataIndex: 'name',
    key: 'name',
    sorter: (a, b) => a.name.localeCompare(b.name),
  },
  {
    title: 'Age',
    dataIndex: 'age',
    key: 'age',
    sorter: (a, b) => a.age - b.age,
  },
  {
    title: 'Address',
    dataIndex: 'address',
    key: 'address',
    sorter: (a, b) => a.address.localeCompare(b.address),
  },
];

const data_dashboard = [
  {
    key: '1',
    name: 'John Brown',
    age: 32,
    address: 'New York No. 1 Lake Park',
  },
  {
    key: '2',
    name: 'Jim Green',
    age: 42,
    address: 'London No. 1 Lake Park',
  },
  {
    key: '3',
    name: 'Joe Black',
    age: 32,
    address: 'Sydney No. 1 Lake Park',
  },
  {
    key: '4',
    name: 'Jane Doe',
    age: 29,
    address: 'San Francisco No. 2 Lake Park',
  },
  {
    key: '5',
    name: 'Michael Johnson',
    age: 45,
    address: 'Los Angeles No. 3 Lake Park',
  },
  {
    key: '6',
    name: 'Michael Johnson',
    age: 45,
    address: 'Los Angeles No. 3 Lake Park',
  },
  {
    key: '7',
    name: 'Michael Johnson',
    age: 45,
    address: 'Los Angeles No. 3 Lake Park',
  },
  {
    key: '8',
    name: 'Michael Johnson',
    age: 45,
    address: 'Los Angeles No. 3 Lake Park',
  },
  {
    key: '9',
    name: 'Michael Johnson',
    age: 45,
    address: 'Los Angeles No. 3 Lake Park',
  },
  {
    key: '10',
    name: 'Michael Johnson',
    age: 45,
    address: 'Los Angeles No. 3 Lake Park',
  },
  {
    key: '11',
    name: 'Michael Johnson',
    age: 45,
    address: 'Los Angeles No. 3 Lake Park',
  },
  {
    key: '12',
    name: 'Michael Johnson',
    age: 45,
    address: 'Los Angeles No. 3 Lake Park',
  },
  {
    key: '13',
    name: 'Michael Johnson',
    age: 45,
    address: 'Los Angeles No. 3 Lake Park',
  },
  {
    key: '14',
    name: 'Michael Johnson',
    age: 45,
    address: 'Los Angeles No. 3 Lake Park',
  },
  {
    key: '15',
    name: 'Michael Johnson',
    age: 45,
    address: 'Los Angeles No. 3 Lake Park',
  },
];

const TablesDashboard = () => <Table columns={columns_dashboard} dataSource={data_dashboard} />;

export default TablesDashboard;
