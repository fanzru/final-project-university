import useDrawer from '../../../hooks/useDrawer';
import clsx from 'clsx';
import { useRouter } from 'next/router';
import UnstyledLink from '../../unstyledlink/index';
import { Routes } from './navbar';
import {useEffect} from 'react';
const Drawer = () => {
  const { pathname } = useRouter();
  const { changeDrawer } = useDrawer();
  let accessToken: any
  useEffect(() => {
    if (typeof window !== 'undefined') {
      accessToken = localStorage.getItem('token');
    }
  },[])
 
  return (
      <aside
        className={clsx(
          'fixed md:hidden w-screen h-screen',
          'dark:bg-dark-900 bg-gray-50',
          'left-0 top-20 z-50'
        )}
      >
        <div className='layout flex flex-col'>
        {
          
          Routes.map((route) => {
            if (route.name === "Profile" && (accessToken === null)) {
              return <></>
            }
            if (((route.name === "Login" || (route.name === "Register")) && (accessToken !== null))){
              return <></>
            }
            return (
              <UnstyledLink
                href={route.link}
                key={route.link}
                onClick={changeDrawer}
                className={clsx(
                  'py-4 border-b-[1.6px] font-medium',
                  pathname === route.link
                    ? 'border-primary-600 dark:border-primary-500'
                    : 'text-gray-700 dark:text-gray-300 border-dark-100 dark:border-dark-500'
                )}
              >
                {route.name}
              </UnstyledLink>
            );
          })
        }
        </div>
      </aside>
  );
};

export default Drawer;
