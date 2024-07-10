import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import '@testing-library/jest-dom/extend-expect';
import DepartmentAdd from '../components/Add-Department';

describe.skip('DepartmentAdd Component', () => {
  it('opens dialog on Add Department button click', () => {

  render(<DepartmentAdd />);

  const addButton = screen.getByTestId('adddepartment');
  fireEvent.click(addButton);

  const dialogTitle = screen.getByTestId('adddepartmenttitle');
  expect(dialogTitle).toBeInTheDocument();
  });

  it('closes dialog on Close button click', () => {
    render(<DepartmentAdd />);

    const addButton = screen.getByTestId('adddepartment');
    fireEvent.click(addButton);

    const dialogTitle = screen.getByTestId('adddepartmenttitle');

    const closeButton = screen.getByTestId('close');
    fireEvent.click(closeButton);

    expect(dialogTitle).not.toBeInTheDocument();
  });

  it('displays error messages for invalid form input', () => {
    render(<DepartmentAdd />);
    const addButton = screen.getByTestId('adddepartment');
    fireEvent.click(addButton);

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);
    
    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Please fill in all asterix fields.');
  });

  it('enables Save button on valid form input', async () => {
    render(<DepartmentAdd />);
    
    const addButton = screen.getByTestId('adddepartment');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('adddepartmenttitle');
    expect(dialogTitle).toBeInTheDocument();

    const departmentnameInput = screen.getByTestId('departmentnamelabel');
    userEvent.type(departmentnameInput, 'Accounts');

    const codeInput = screen.getByTestId('codelabel');
    userEvent.type(codeInput, '106');

    const saveButton = screen.getByTestId('save');
    expect(saveButton).not.toHaveClass('cursor-not-allowed');
  });

  it('displays error messages for invalid department name input', async () => {
    render(<DepartmentAdd />);
    
    const addButton = screen.getByTestId('adddepartment');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('adddepartmenttitle');
    expect(dialogTitle).toBeInTheDocument();

    const departmentnameInput = screen.getByTestId('departmentnamelabel');
    userEvent.type(departmentnameInput, 'Accounts2');

    const codeInput = screen.getByTestId('codelabel');
    userEvent.type(codeInput, '106');

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);

    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Department Name should be in text format');
  });

  it('displays error messages for invalid code input', async () => {
    render(<DepartmentAdd />);
    
    const addButton = screen.getByTestId('adddepartment');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('adddepartmenttitle');
    expect(dialogTitle).toBeInTheDocument();

    const departmentnameInput = screen.getByTestId('departmentnamelabel');
    userEvent.type(departmentnameInput, 'Accounts');

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
