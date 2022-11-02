import React from 'react';
import { useTheme } from 'next-themes';
import { useEffect, useState } from 'react';
import { HiSun, HiMoon,HiOutlineLogout } from 'react-icons/hi';
import Button from '../components/button';
import clsx from 'clsx';

const DarkMode = () => {
  const [mounted, setMounted] = useState(false);
  const { theme, setTheme } = useTheme();

  useEffect(() => setMounted(true), []);
  if (!mounted) return null;
  return (
    <Button
      className={clsx(
        'text-lg rounded-md h-10 w-10',
        'bg-primary-50 dark:bg-dark-500/50'
      )}
      onClick={() => {
        setTheme(theme === 'light' ? 'dark' : 'light');
      }}
      aria-label='Darkmode Button'
    >
      {theme === 'light' ? (
        <HiMoon className={clsx('text-primary-700')} />
      ) : (
        <HiSun className={clsx('text-yellow-400')} />
      )}
    </Button>
  );
};

export default DarkMode;
