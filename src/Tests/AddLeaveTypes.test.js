import React, { useState } from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import '@testing-library/jest-dom/extend-expect';
import LeaveTypesAdd from '../components/Add-LeaveTypes';

describe.skip('LeaveTypesAdd Component', () => {
  it('opens dialog on Add Leave Types button click', () => {

  render(<LeaveTypesAdd />);

  const addButton = screen.getByTestId('addleavetypes');
  fireEvent.click(addButton);

  const dialogTitle = screen.getByTestId('addleavetypestitle');
  expect(dialogTitle).toBeInTheDocument();
  });

  it('closes dialog on Close button click', () => {
    render(<LeaveTypesAdd />);

    const addButton = screen.getByTestId('addleavetypes');
    fireEvent.click(addButton);

    const dialogTitle = screen.getByTestId('addleavetypestitle');

    const closeButton = screen.getByTestId('close');
    fireEvent.click(closeButton);

    expect(dialogTitle).not.toBeInTheDocument();
  });

  it('displays error messages for invalid form input', () => {
    render(<LeaveTypesAdd />);
    const addButton = screen.getByTestId('addleavetypes');
    fireEvent.click(addButton);

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);
    
    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Please fill in all asterix fields.');
  });

  it('enables Save button on valid form input', async () => {
    render(<LeaveTypesAdd />);
    
    const addButton = screen.getByTestId('addleavetypes');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('addleavetypestitle');
    expect(dialogTitle).toBeInTheDocument();

    const leavetypenameInput = screen.getByTestId('leavetypenamelabel');
    userEvent.type(leavetypenameInput, 'Government Holiday');

    const codeInput = screen.getByTestId('codelabel');
    userEvent.type(codeInput, '104');

    const saveButton = screen.getByTestId('save');
    expect(saveButton).not.toHaveClass('cursor-not-allowed');
  });

  it('displays error messages for invalid leave type name input', async () => {
    render(<LeaveTypesAdd />);
    
    const addButton = screen.getByTestId('addleavetypes');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('addleavetypestitle');
    expect(dialogTitle).toBeInTheDocument();

    const leavetypenameInput = screen.getByTestId('leavetypenamelabel');
    userEvent.type(leavetypenameInput, 'Government Holiday 2');

    const codeInput = screen.getByTestId('codelabel');
    userEvent.type(codeInput, '104');

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);

    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Leave Type Name should be in text format');
  });

  it('displays error messages for invalid code input', async () => {
    render(<LeaveTypesAdd />);
    
    const addButton = screen.getByTestId('addleavetypes');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('addleavetypestitle');
    expect(dialogTitle).toBeInTheDocument();

    const leavetypenameInput = screen.getByTestId('leavetypenamelabel');
    userEvent.type(leavetypenameInput, 'Government Holiday');

    const codeInput = screen.getByTestId('codelabel');
    userEvent.type(codeInput, 'abc');

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);

    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Code should be numeric');
  });

});
