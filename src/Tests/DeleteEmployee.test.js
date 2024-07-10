import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import TablesEmp from '../components/shared/Employee-Table';


beforeAll(() => {
  window.matchMedia = window.matchMedia || function() {
    return {
      matches: false,
      addListener: function() {},
      removeListener: function() {}
    };
  };
});

describe.skip('EmployeeDelete Component', () => {
  it('opens warning on Delete Employee button click', async () => {

  render(<TablesEmp />);

  const deleteButton = screen.getAllByTestId('deletebutton');
  fireEvent.click(deleteButton[0]);

  await waitFor(() => {
    const dialogTitle = screen.getByText('Are you sure to delete this row?');
    expect(dialogTitle).toBeInTheDocument();
  });
  });
});