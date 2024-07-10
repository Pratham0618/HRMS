import React, { useState } from 'react';
import { Table, Input, Select, Button, Popconfirm } from 'antd';
import { MinusCircleFilled } from '@ant-design/icons';

const { Search } = Input;
const { Option } = Select;

const data_dept = [
  {
    key: '1',
    departmentid: 1,
    departmentcode: 101,
    departmentname: 'Human Resources',
  },
  {
    key: '2',
    departmentid: 2,
    departmentcode: 102,
    departmentname: 'Information Technology',
  },
  {
    key: '3',
    departmentid: 3,
    departmentcode: 103,
    departmentname: 'Finance',
  },
  {
    key: '4',
    departmentid: 4,
    departmentcode: 104,
    departmentname: 'Sales',
  },
  {
    key: '5',
    departmentid: 5,
    departmentcode: 105,
    departmentname: 'Marketing',
  },
];


const TablesDept = () => {
  const [searchText, setSearchText] = useState('');
  const [searchColumn, setSearchColumn] = useState('departmentname');
  const [data, setData] = useState(data_dept);

  const handleSearch = (selectedColumn, value) => {
    setSearchColumn(selectedColumn);
    setSearchText(value);
  };

  const handleDelete = (key) => {
    setData(data.filter(item => item.key !== key));
  };

  const columns_dept = [
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
      title: 'Id',
      dataIndex: 'departmentid',
      key: 'departmentid',
      sorter: (a, b) => a.departmentid - b.departmentid,
    },
    {
      title: 'Code',
      dataIndex: 'departmentcode',
      key: 'departmentcode',
      sorter: (a, b) => a.departmentcode - b.departmentcode,
    },
    {
      title: 'Department Name',
      dataIndex: 'departmentname',
      key: 'departmentname',
      sorter: (a, b) => a.departmentname.localeCompare(b.departmentname),
    },
  ];

  const filteredData = data.filter(item => 
    item[searchColumn].toString().toLowerCase().includes(searchText.toLowerCase())
  );

  return (
    <div>
      <div style={{ marginBottom: 16, display: 'flex', alignItems: 'center' }}>
        <Select
          defaultValue="departmentname"
          style={{ width: 200, marginRight: 8 }}
          onChange={value => setSearchColumn(value)}
        >
          <Option value="departmentid">Id</Option>
          <Option value="departmentcode">Code</Option>
          <Option value="departmentname"> Department Name</Option>
        </Select>
        <Search
          placeholder={`Search ${searchColumn}`}
          onSearch={value => handleSearch(searchColumn, value)}
          onChange={e => setSearchText(e.target.value)}
          style={{ flex:1 }}
        />
      </div>
      <Table columns={columns_dept} dataSource={filteredData} />
    </div>
  );
};

export default TablesDept;
