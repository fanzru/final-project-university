import clsx from 'clsx';
import { useRouter } from 'next/router';

import DrawerToggler from './components/drawertoggler';
import DarkMode from './components/darkmode';
import Navbar from './components/navbar';

const Header = () => {
  const { pathname } = useRouter();
  const isError =
    pathname === '/_error' || pathname === '/_offline' || pathname === '/404';

  if (isError) {
    return null;
  }
  return (
    <header
      className={clsx(
        ' inset-0 z-50 h-14 z-40 fixed',
        'bg-zinc-50 dark:bg-dark-900'
      )}
    >
      <div
        className={clsx(
          'layout gap-2 h-full',
          'flex items-center justify-between'
        )}
      >
        <DrawerToggler/>
        <Navbar />
        <DarkMode />
      </div>
    </header>
  );
};

export default Header;
