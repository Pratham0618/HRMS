import React from 'react';
import * as Popover from '@radix-ui/react-popover';
import { BellIcon, Cross2Icon } from '@radix-ui/react-icons';

const PopoverDemo = () => (
  <Popover.Root>
    <Popover.Trigger asChild>
      <button
        className="rounded-full w-[35px] h-[35px] inline-flex items-center justify-center text-indigo11 bg-white shadow-[0_2px_10px] shadow-blackA4 hover:bg-indigo3 focus:shadow-[0_0_0_2px] focus:shadow-black cursor-default outline-none"
        aria-label="Notifications"
      >
        <BellIcon />
      </button>
    </Popover.Trigger>
    <Popover.Portal>
      <Popover.Content
        className="rounded p-5 w-[260px] bg-white shadow-[0_10px_38px_-10px_hsla(206,22%,7%,.35),0_10px_20px_-15px_hsla(206,22%,7%,.2)] focus:shadow-[0_10px_38px_-10px_hsla(206,22%,7%,.35),0_10px_20px_-15px_hsla(206,22%,7%,.2),0_0_0_2px_theme(colors.indigo7)] will-change-[transform,opacity] data-[state=open]:data-[side=top]:animate-slideDownAndFade data-[state=open]:data-[side=right]:animate-slideLeftAndFade data-[state=open]:data-[side=bottom]:animate-slideUpAndFade data-[state=open]:data-[side=left]:animate-slideRightAndFade"
        sideOffset={5}
      >
        <div className="flex flex-col gap-2.5">
          <p className="text-mauve12 text-[15px] leading-[19px] font-medium mb-2.5">Notifications</p>
        </div>
        <Popover.Close
          className="rounded-full h-[25px] w-[25px] inline-flex items-center justify-center text-indigo11 absolute top-[5px] right-[5px] hover:bg-indigo4 focus:shadow-[0_0_0_2px] focus:shadow-indigo7 outline-none cursor-default"
          aria-label="Close"
        >
          <Cross2Icon />
        </Popover.Close>
        <Popover.Arrow className="fill-white" />
      </Popover.Content>
    </Popover.Portal>
  </Popover.Root>
);

export default PopoverDemo;