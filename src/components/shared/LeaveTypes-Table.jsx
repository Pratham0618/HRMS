import React, { useState } from 'react';
import { Table, Input, Select, Button, Popconfirm } from 'antd';
import { MinusCircleFilled } from '@ant-design/icons';

const { Search } = Input;
const { Option } = Select;

const data_leavetypes = [
  {
    key: '1',
    leavetypeid: 1,
    leavetypecode: 101,
    leavetypename: 'Annual Leave',
  },
  {
    key: '2',
    leavetypeid: 2,
    leavetypecode: 102,
    leavetypename: 'Medical Leave',
  },
  {
    key: '3',
    leavetypeid: 3,
    leavetypecode: 103,
    leavetypename: 'Sick Leave',
  },
];

const TablesLeaveTypes = () => {
  const [searchText, setSearchText] = useState('');
  const [searchColumn, setSearchColumn] = useState('leavetypename');
  const [data, setData] = useState(data_leavetypes);

  const handleSearch = (selectedColumn, value) => {
    setSearchColumn(selectedColumn);
    setSearchText(value);
  };

  const handleDelete = (key) => {
    setData(data.filter(item => item.key !== key));
  };

  const columns_leavetypes = [
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
      dataIndex: 'leavetypeid',
      key: 'leavetypeid',
      sorter: (a, b) => a.leavetypeid - b.leavetypeid,
    },
    {
      title: 'Code',
      dataIndex: 'leavetypecode',
      key: 'leavetypecode',
      sorter: (a, b) => a.leavetypecode - b.leavetypecode,
    },
    {
      title: 'Leave Type Name',
      dataIndex: 'leavetypename',
      key: 'leavetypename',
      sorter: (a, b) => a.leavetypename.localeCompare(b.leavetypename),
    },
  ];

  const filteredData = data.filter(item =>
    item[searchColumn].toString().toLowerCase().includes(searchText.toLowerCase())
  );

  return (
    <div>
      <div style={{ marginBottom: 16, display: 'flex', alignItems: 'center' }}>
        <Select
          defaultValue="leavetypename"
          style={{ width: 200, marginRight: 8 }}
          onChange={value => setSearchColumn(value)}
        >
          <Option value="leavetypeid">Id</Option>
          <Option value="leavetypecode">Code</Option>
          <Option value="leavetypename">Leave Type Name</Option>
        </Select>
        <Search
          placeholder={`Search ${searchColumn}`}
          onSearch={value => handleSearch(searchColumn, value)}
          onChange={e => setSearchText(e.target.value)}
          style={{ flex: 1 }}
        />
      </div>
      <Table columns={columns_leavetypes} dataSource={filteredData} />
    </div>
  );
};

export default TablesLeaveTypes;
