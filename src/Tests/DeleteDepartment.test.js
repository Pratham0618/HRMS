import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import TablesDept from '../components/shared/Department-Table';


beforeAll(() => {
  window.matchMedia = window.matchMedia || function() {
    return {
      matches: false,
      addListener: function() {},
      removeListener: function() {}
    };
  };
});

describe.skip('DepartmentDelete Component', () => {
  it('opens warning on Delete Department button click', async () => {

  render(<TablesDept />);

  const deleteButton = screen.getAllByTestId('deletebutton');
  fireEvent.click(deleteButton[0]);

  await waitFor(() => {
    const dialogTitle = screen.getByText('Are you sure to delete this row?');
    expect(dialogTitle).toBeInTheDocument();
  });
  });
});