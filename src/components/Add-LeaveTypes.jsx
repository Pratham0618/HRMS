import React, { useState } from 'react';
import * as Dialog from '@radix-ui/react-dialog';
import { Cross2Icon } from '@radix-ui/react-icons';

const LeaveTypesAdd = () => {
  const [leavetypename, setLeaveTypeName] = useState('');
  const [code, setCode] = useState('');
  const [errorMessage, setErrorMessage] = useState('');

  const isValidLeaveTypeName = /^\S+[A-Za-z ]+$/.test(leavetypename);
  const isValidCode = /^\d+$/.test(code);


  const isFormValid = () => {
    return (
      leavetypename !== '' &&
      code !== '' &&
      isValidLeaveTypeName && isValidCode
    );

  };

  const isFilled = () => {
    return (
      leavetypename !== '' &&
      code !== ''
    );

  };

  const handleSave = () => {
    if (!isFilled()) {
      setErrorMessage('Please fill in all asterix fields.');
      return;
    }
    if (!isValidLeaveTypeName) {
      setErrorMessage('Leave Type Name should be in text format');
        return;
    }
    if (!isValidCode) {
      setErrorMessage('Code should be numeric');
        return;
    }
    
    const leavetypeData = {
      code,
      leavetypename,
    };

    console.log('Leave Type Data:', leavetypeData);

    setLeaveTypeName('');
    setCode('');
    setErrorMessage('');
  };

  return (
    <Dialog.Root>
      <Dialog.Trigger asChild>
        <button className="text-indigo11 shadow-blueA4 hover:bg-violet3 inline-flex h-[35px] items-center justify-center rounded-[4px] bg-white px-[15px] font-medium leading-none shadow-[0_2px_10px] focus:shadow-[0_0_0_2px] focus:shadow-blue focus:outline-none" data-testid='addleavetypes'>
          Add Leave Type
        </button>
      </Dialog.Trigger>
      <Dialog.Portal>
        <Dialog.Overlay className="bg-blackA6 data-[state=open]:animate-overlayShow fixed inset-0" />
        <Dialog.Content className="data-[state=open]:animate-contentShow fixed top-[50%] left-[50%] max-h-[85vh] w-[90vw] max-w-[450px] translate-x-[-50%] translate-y-[-50%] rounded-[6px] bg-white p-[25px] shadow-[hsl(206_22%_7%_/_35%)_0px_10px_38px_-10px,_hsl(206_22%_7%_/_20%)_0px_10px_20px_-15px] focus:outline-none z-[100]">
          <Dialog.Title className="text-violet12 m-0 text-[17px] font-medium" data-testid='addleavetypestitle'>
            Add Leave Type
          </Dialog.Title>
          <Dialog.Description className="text-violet12 mt-[10px] mb-5 text-[15px] leading-normal">
            Add leave type details. Click save when you're done.
          </Dialog.Description>
          {errorMessage && (
            <div className="mb-[15px] text-red-600 text-[15px]" data-testid='errormessage'>
              {errorMessage}
            </div>
          )}
          <fieldset className="mb-[15px] flex items-center gap-5">
            <label className="text-indigo11 w-[90px] text-right text-[15px]">
              Code<span className="text-red-600">*</span>
            </label>
            <input
              className="text-indigo11 shadow-indigo7 focus:shadow-indigo8 inline-flex h-[35px] w-full flex-1 items-center justify-center rounded-[4px] px-[10px] text-[15px] leading-none shadow-[0_0_0_1px] outline-none focus:shadow-[0_0_0_2px]"
              id="code"
              value={code}
              onChange={(e) => setCode(e.target.value)}
              data-testid='codelabel'
            />
          </fieldset>
          <fieldset className="mb-[15px] flex items-center gap-5">
            <label className="text-indigo11 w-[90px] text-right text-[14px]">
              Leave Type Name<span className="text-red-600">*</span>
            </label>
            <input
              className="text-indigo11 shadow-indigo7 focus:shadow-indigo8 inline-flex h-[35px] w-full flex-1 items-center justify-center rounded-[4px] px-[10px] text-[15px] leading-none shadow-[0_0_0_1px] outline-none focus:shadow-[0_0_0_2px]"
              id="leavetypename"
              value={leavetypename}
              onChange={(e) => setLeaveTypeName(e.target.value)}
              data-testid='leavetypenamelabel'
            />
          </fieldset>
          <div className="mt-[25px] flex justify-end">
            <button
              className={`bg-green4 text-green11 hover:bg-green5 focus:shadow-green7 inline-flex h-[35px] items-center justify-center rounded-[4px] px-[15px] font-medium leading-none focus:shadow-[0_0_0_2px] focus:outline-none ${!isFormValid() && 'opacity-50 cursor-not-allowed'}`}
              onClick={handleSave} data-testid='save'
              //disabled={!isFormValid()}
            >
              Save
            </button>
          </div>
          <Dialog.Close asChild>
            <button
              className="text-indigo11 hover:bg-indigo4 focus:shadow-indigo7 absolute top-[10px] right-[10px] inline-flex h-[25px] w-[25px] appearance-none items-center justify-center rounded-full focus:shadow-[0_0_0_2px] focus:outline-none"
              aria-label="Close" data-testid='close'
            >
              <Cross2Icon />
            </button>
          </Dialog.Close>
        </Dialog.Content>
      </Dialog.Portal>
    </Dialog.Root>
  );
};

export default LeaveTypesAdd;
