import React from 'react';
import { render, fireEvent, screen, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import TablesEmp from '../components/shared/Employee-Table';

export async function fetchData() {
  try {
    const response = await fetch('https://my-json-server.typicode.com/pk2601/employee-dashboard/posts');
    if (!response.ok) {
      throw new Error('Failed to fetch data');
    }
    return await response.json();
  } catch (error) {
    console.error('Error fetching data:', error);
    return [];
  }
}


beforeAll(() => {
  window.matchMedia = window.matchMedia || function() {
    return {
      matches: false,
      addListener: function() {},
      removeListener: function() {}
    };
  };
});

const mockDataa = [
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

describe('TablesEmp component', () => {

  let mockData = [];

  beforeAll(async () => {
    mockData = await fetchData();
  });

  beforeEach(() => {
    render(<TablesEmp dataemp={mockData} />);
  });

  it('renders table columns correctly', () => {
    const columns = screen.getAllByRole('columnheader');
    expect(columns).toHaveLength(9);
    expect(screen.getByTestId('id')).toBeInTheDocument();
    expect(screen.getByTestId('name')).toBeInTheDocument();
    expect(screen.getByTestId('department')).toBeInTheDocument();
    expect(screen.getByTestId('manager')).toBeInTheDocument();
    expect(screen.getByTestId('email')).toBeInTheDocument();
    expect(screen.getByTestId('phone')).toBeInTheDocument();
    expect(screen.getByTestId('address')).toBeInTheDocument();
    expect(screen.getByTestId('dob')).toBeInTheDocument();
  });

  it('renders initial rows from mock data', () => {
    const rows = screen.getAllByRole('row');
    expect(rows.slice(1)).toHaveLength(mockData.length);
    expect(screen.getByText('John Brown')).toBeInTheDocument();
  });

  it('filters table rows based on search input', async () => {
    const searchInput = screen.getByPlaceholderText(/search/i);
    fireEvent.change(searchInput, { target: { value: 'John' } });
    const filteredRow = await screen.findByText('John Brown');
    expect(filteredRow).toBeInTheDocument();
  });

  it('deletes a row on confirmation', async () => {
    const deleteButton = screen.getAllByTestId('deletebutton');
    fireEvent.click(deleteButton[0]);
    const confirmButton = await screen.findByText('Yes');
    fireEvent.click(confirmButton);
    await waitFor(() => {
      const deletedRow = screen.queryByText('John Brown');
      expect(deletedRow).not.toBeInTheDocument();
    });
  });
});
