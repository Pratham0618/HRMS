import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import { BrowserRouter as Router } from 'react-router-dom';
import '@testing-library/jest-dom/extend-expect';
import Sidebar from '../components/shared/Sidebar';

beforeAll(() => {
    window.matchMedia = window.matchMedia || function() {
      return {
        matches: false,
        addListener: function() {},
        removeListener: function() {}
      };
    };
  });

describe.skip('Logoutconfirmation Component', () => {
  it('opens warning on Logout button click', async () => {

    render(
        <Router>
          <Sidebar />
        </Router>
      );

  const logoutButton = screen.getByTestId('logout');
  fireEvent.click(logoutButton);

  await waitFor(() => {
    const dialogTitle = screen.getByText('Are you sure you want to logout?');
    expect(dialogTitle).toBeInTheDocument();
  });
  });
});