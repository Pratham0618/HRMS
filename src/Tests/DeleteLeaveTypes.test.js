import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import TablesLeaveTypes from '../components/shared/LeaveTypes-Table';


beforeAll(() => {
  window.matchMedia = window.matchMedia || function() {
    return {
      matches: false,
      addListener: function() {},
      removeListener: function() {}
    };
  };
});

describe.skip('LeaveTypesDelete Component', () => {
  it('opens warning on Delete Leave Types button click', async () => {

  render(<TablesLeaveTypes />);

  const deleteButton = screen.getAllByTestId('deletebutton');
  fireEvent.click(deleteButton[0]);

  await waitFor(() => {
    const dialogTitle = screen.getByText('Are you sure to delete this row?');
    expect(dialogTitle).toBeInTheDocument();
  });
  });
});