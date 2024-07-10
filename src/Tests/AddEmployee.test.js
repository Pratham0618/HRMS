import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import '@testing-library/jest-dom/extend-expect';
import EmployeeAdd from '../components/Add-Employee';

describe.skip('EmployeeAdd Component', () => {
  it('opens dialog on Add Employee button click', () => {

  render(<EmployeeAdd />);

  const addButton = screen.getByTestId('addemployee');
  fireEvent.click(addButton);

  const dialogTitle = screen.getByTestId('addemployeetitle');
  expect(dialogTitle).toBeInTheDocument();
  });

  it('closes dialog on Close button click', () => {
    render(<EmployeeAdd />);

    const addButton = screen.getByTestId('addemployee');
    fireEvent.click(addButton);

    const dialogTitle = screen.getByTestId('addemployeetitle');

    const closeButton = screen.getByTestId('close');
    fireEvent.click(closeButton);

    expect(dialogTitle).not.toBeInTheDocument();
  });

  it('displays error messages for invalid form input', () => {
    render(<EmployeeAdd />);
    const addButton = screen.getByTestId('addemployee');
    fireEvent.click(addButton);

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);
    
    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Please fill in all asterix fields.');
  });

  it('enables Save button on valid form input', async () => {
    render(<EmployeeAdd />);
    
    const addButton = screen.getByTestId('addemployee');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('addemployeetitle');
    expect(dialogTitle).toBeInTheDocument();

    const nameInput = screen.getByTestId('namelabel');
    userEvent.type(nameInput, 'James');

    const emailInput = screen.getByTestId('emaillabel');
    userEvent.type(emailInput, 'james@example.com');

    const phoneInput = screen.getByTestId('phonelabel');
    userEvent.type(phoneInput, '1234567890');

    const addressInput = screen.getByTestId('addresslabel');
    userEvent.type(addressInput, '123 Main St');

    const dobInput = screen.getByTestId('doblabel');
    userEvent.type(dobInput, '1990-01-01');

    const saveButton = screen.getByTestId('save');
    expect(saveButton).not.toHaveClass('cursor-not-allowed');
  });

  it('displays error messages for invalid name input', async () => {
    render(<EmployeeAdd />);
    
    const addButton = screen.getByTestId('addemployee');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('addemployeetitle');
    expect(dialogTitle).toBeInTheDocument();

    const nameInput = screen.getByTestId('namelabel');
    userEvent.type(nameInput, '123');

    const emailInput = screen.getByTestId('emaillabel');
    userEvent.type(emailInput, 'james@example.com');

    const phoneInput = screen.getByTestId('phonelabel');
    userEvent.type(phoneInput, '1234567890');

    const addressInput = screen.getByTestId('addresslabel');
    userEvent.type(addressInput, '123 Main St');

    const dobInput = screen.getByTestId('doblabel');
    userEvent.type(dobInput, '1990-01-01');

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);

    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Name should be in text format');
  });

  it('displays error messages for invalid managerid input', async () => {
    render(<EmployeeAdd />);
    
    const addButton = screen.getByTestId('addemployee');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('addemployeetitle');
    expect(dialogTitle).toBeInTheDocument();

    const nameInput = screen.getByTestId('namelabel');
    userEvent.type(nameInput, 'James');

    const managerInput = screen.getByTestId('manageridlabel');
    userEvent.type(managerInput, 'error');

    const emailInput = screen.getByTestId('emaillabel');
    userEvent.type(emailInput, 'james@example.com');

    const phoneInput = screen.getByTestId('phonelabel');
    userEvent.type(phoneInput, '1234567890');

    const addressInput = screen.getByTestId('addresslabel');
    userEvent.type(addressInput, '123 Main St');

    const dobInput = screen.getByTestId('doblabel');
    userEvent.type(dobInput, '1990-01-01');

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);

    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Manager Id should be numeric');
  });

  it('displays error messages for invalid email input', async () => {
    render(<EmployeeAdd />);
    
    const addButton = screen.getByTestId('addemployee');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('addemployeetitle');
    expect(dialogTitle).toBeInTheDocument();

    const nameInput = screen.getByTestId('namelabel');
    userEvent.type(nameInput, 'james');

    const emailInput = screen.getByTestId('emaillabel');
    userEvent.type(emailInput, 'jamesexample.com');

    const phoneInput = screen.getByTestId('phonelabel');
    userEvent.type(phoneInput, '1234567890');

    const addressInput = screen.getByTestId('addresslabel');
    userEvent.type(addressInput, '123 Main St');

    const dobInput = screen.getByTestId('doblabel');
    userEvent.type(dobInput, '1990-01-01');

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);

    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Email is of wrong type');
  });

  it('displays error messages for invalid less phone input', async () => {
    render(<EmployeeAdd />);
    
    const addButton = screen.getByTestId('addemployee');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('addemployeetitle');
    expect(dialogTitle).toBeInTheDocument();

    const nameInput = screen.getByTestId('namelabel');
    userEvent.type(nameInput, 'james');

    const emailInput = screen.getByTestId('emaillabel');
    userEvent.type(emailInput, 'james@example.com');

    const phoneInput = screen.getByTestId('phonelabel');
    userEvent.type(phoneInput, '12345670');

    const addressInput = screen.getByTestId('addresslabel');
    userEvent.type(addressInput, '123 Main St');

    const dobInput = screen.getByTestId('doblabel');
    userEvent.type(dobInput, '1990-01-01');

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);

    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Phone must contain 10 numbers');
  });

  it('displays error messages for invalid text phone input', async () => {
    render(<EmployeeAdd />);
    
    const addButton = screen.getByTestId('addemployee');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('addemployeetitle');
    expect(dialogTitle).toBeInTheDocument();

    const nameInput = screen.getByTestId('namelabel');
    userEvent.type(nameInput, 'james');

    const emailInput = screen.getByTestId('emaillabel');
    userEvent.type(emailInput, 'james@example.com');

    const phoneInput = screen.getByTestId('phonelabel');
    userEvent.type(phoneInput, '1234567as0');

    const addressInput = screen.getByTestId('addresslabel');
    userEvent.type(addressInput, '123 Main St');

    const dobInput = screen.getByTestId('doblabel');
    userEvent.type(dobInput, '1990-01-01');

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);

    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Phone must contain 10 numbers');
  });

  it('displays error messages for invalid address input', async () => {
    render(<EmployeeAdd />);
    
    const addButton = screen.getByTestId('addemployee');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('addemployeetitle');
    expect(dialogTitle).toBeInTheDocument();

    const nameInput = screen.getByTestId('namelabel');
    userEvent.type(nameInput, 'james');

    const emailInput = screen.getByTestId('emaillabel');
    userEvent.type(emailInput, 'james@example.com');

    const phoneInput = screen.getByTestId('phonelabel');
    userEvent.type(phoneInput, '1234567890');

    const addressInput = screen.getByTestId('addresslabel');
    userEvent.type(addressInput, '     123 Main St');

    const dobInput = screen.getByTestId('doblabel');
    userEvent.type(dobInput, '1990-01-01');

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);

    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Address is of wrong type');
  });

  it('displays error messages for invalid dob input', async () => {
    render(<EmployeeAdd />);
    
    const addButton = screen.getByTestId('addemployee');
    fireEvent.click(addButton);

    const dialogTitle = await screen.findByTestId('addemployeetitle');
    expect(dialogTitle).toBeInTheDocument();

    const nameInput = screen.getByTestId('namelabel');
    userEvent.type(nameInput, 'james');

    const emailInput = screen.getByTestId('emaillabel');
    userEvent.type(emailInput, 'james@example.com');

    const phoneInput = screen.getByTestId('phonelabel');
    userEvent.type(phoneInput, '1234567890');

    const addressInput = screen.getByTestId('addresslabel');
    userEvent.type(addressInput, '123 Main St');

    const dobInput = screen.getByTestId('doblabel');
    userEvent.type(dobInput, '19901-01');

    const saveButton = screen.getByTestId('save');
    fireEvent.click(saveButton);

    expect(saveButton).toHaveClass('cursor-not-allowed');

    const errorMessage = screen.getByTestId('errormessage');
    expect(errorMessage).toBeInTheDocument();
    expect(errorMessage.textContent).toEqual('Date of Birth should be of format YYYY-MM-DD');
  });
});
