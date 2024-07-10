import React, { useState } from 'react';
import { Table, Input, Select, Button, Popconfirm } from 'antd';
import { MinusCircleFilled } from '@ant-design/icons';

const { Search } = Input;
const { Option } = Select;

const data_emp = [
  {
    key: '1',
    employeeid: 1,
    name: 'John Brown',
    departmentid: 101,
    managerid: 201,
    email: 'john.brown@example.com',
    phone: '123-456-7890',
    address: 'New York No. 1 Lake Park',
    dob: '1990-01-01',
  },
  {
    key: '2',
    employeeid: 2,
    name: 'Jim Green',
    departmentid: 102,
    managerid: 202,
    email: 'jim.green@example.com',
    phone: '098-765-4321',
    address: 'London No. 1 Lake Park',
    dob: '1982-05-12',
  },
  {
    key: '3',
    employeeid: 3,
    name: 'Joe Black',
    departmentid: 103,
    managerid: 203,
    email: 'joe.black@example.com',
    phone: '111-222-3333',
    address: 'Sydney No. 1 Lake Park',
    dob: '1990-03-15',
  },
  {
    key: '4',
    employeeid: 4,
    name: 'Jane Doe',
    departmentid: 104,
    managerid: 204,
    email: 'jane.doe@example.com',
    phone: '444-555-6666',
    address: 'San Francisco No. 2 Lake Park',
    dob: '1993-09-21',
  },
  {
    key: '5',
    employeeid: 5,
    name: 'Michael Johnson',
    departmentid: 105,
    managerid: 205,
    email: 'michael.johnson@example.com',
    phone: '777-888-9999',
    address: 'Los Angeles No. 3 Lake Park',
    dob: '1978-11-30',
  },
];


const TablesEmp = ({dataemp = data_emp}) => {
  const [searchText, setSearchText] = useState('');
  const [searchColumn, setSearchColumn] = useState('name');
  const [data, setData] = useState(dataemp);

  const handleSearch = (selectedColumn, value) => {
    setSearchColumn(selectedColumn);
    setSearchText(value);
  };

  const handleDelete = (key) => {
    setData(data.filter(item => item.key !== key));
  };

  const columns_emp = [
    {
      title: '',
      key: 'delete',
      width: 30,
      render: (record) => (
        <Popconfirm
          title="Are you sure to delete this row?"
          onConfirm={() => handleDelete(record.key)}
          okText="Yes"
          cancelText="No"
        >
          <Button
            type="primary"
            danger
            shape="circle"
            icon={<MinusCircleFilled />}
            size="small"
            data-testid='deletebutton'
          />
        </Popconfirm>
      ),
    },
    {
      title: <div  data-testid='id'>'Id'</div>,
      dataIndex: 'employeeid',
      key: 'employeeid',
      sorter: (a, b) => a.employeeid - b.employeeid,
    },
    {
      title: <div  data-testid='name'>'Name'</div>,
      dataIndex: 'name',
      key: 'name',
      sorter: (a, b) => a.name.localeCompare(b.name),
    },
    {
      title: <div  data-testid='department'>'Department'</div>,
      dataIndex: 'departmentid',
      key: 'departmentid',
      sorter: (a, b) => a.departmentid - b.departmentid,
    },
    {
      title: <div  data-testid='manager'>'Manager'</div>,
      dataIndex: 'managerid',
      key: 'managerid',
      sorter: (a, b) => a.managerid - b.managerid,
    },
    {
      title: <div  data-testid='email'>'Email'</div>,
      dataIndex: 'email',
      key: 'email',
      sorter: (a, b) => a.email.localeCompare(b.email),
    },
    {
      title: <div  data-testid='phone'>'Phone'</div>,
      dataIndex: 'phone',
      key: 'phone',
      sorter: (a, b) => a.phone.localeCompare(b.phone),
    },
    {
      title: <div  data-testid='address'>'Address'</div>,
      dataIndex: 'address',
      key: 'address',
      sorter: (a, b) => a.address.localeCompare(b.address),
    },
    {
      title: <div  data-testid='dob'>'DOB'</div>,
      dataIndex: 'dob',
      key: 'dob',
      sorter: (a, b) => a.dob.localeCompare(b.dob),
    },
  ];

  const filteredData = data.filter(item => 
    item[searchColumn].toString().toLowerCase().includes(searchText.toLowerCase())
  );

  return (
    <div>
      <div style={{ marginBottom: 16, display: 'flex', alignItems: 'center' }}>
        <Select
          defaultValue="name"
          style={{ width: 120, marginRight: 8 }}
          onChange={value => setSearchColumn(value)}
        >
          <Option value="employeeid">Id</Option>
          <Option value="name">Name</Option>
          <Option value="departmentid">Department</Option>
          <Option value="managerid">Manager</Option>
          <Option value="email">Email</Option>
          <Option value="phone">Phone</Option>
          <Option value="address">Address</Option>
          <Option value="dob">DOB</Option>
        </Select>
        <Search
          placeholder={`Search ${searchColumn}`}
          onSearch={value => handleSearch(searchColumn, value)}
          onChange={e => setSearchText(e.target.value)}
          style={{ flex:1 }}
        />
      </div>
      <Table columns={columns_emp} dataSource={filteredData} />
    </div>
  );
};

export default TablesEmp;
