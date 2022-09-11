import useDrawer from '../../../hooks/useDrawer';
import clsx from 'clsx';
import React from 'react';
import { HiMenuAlt2, HiOutlineX } from 'react-icons/hi';
import Drawer from './drawer';
import Button from './button';

const DrawerToggler = () => {
  const { isOpen, changeDrawer } = useDrawer();

  return (
    <>
      <Button
        className={clsx(
          'h-10 w-10 -ml-3.5',
          'text-lg rounded-lg',
          'md:hidden block'
        )}
        onClick={changeDrawer}
      >
        {!isOpen ? <HiMenuAlt2 /> : <HiOutlineX />}
      </Button>
      {isOpen && <Drawer />}
    </>
  );
};

export default DrawerToggler;
