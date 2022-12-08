import clsx from 'clsx';
import Link from 'next/link';
import { useRouter } from 'next/router'
import {useEffect} from 'react';
import { HiOutlineLogout } from 'react-icons/hi';
export const Routes = [
  {
    link: '/',
    name: 'Home',
  },
  {
    link: '/feature',
    name: 'Feature',
  },
  {
    link: '/login',
    name: 'Login',
  },
  {
    link: '/register',
    name: 'Register',
  },
  {
    link: '/submit',
    name: 'Submit',
  },
  {
    link: '/profile',
    name: 'Profile',
  }
];

const Navbar = () => {
  const { pathname } = useRouter();
  let accessToken: any
  if (typeof window !== 'undefined') {
    accessToken = localStorage.getItem('token');
  }
  

  return (
    <nav className={clsx('items-center -ml-3.5 hidden md:flex')}>
        {Routes.map((route) => {
          if ((route.name === "Profile" || route.name === "Submit") && (accessToken === null)) {
            return <></>
          }
          if (((route.name === "Login" || (route.name === "Register")) && (accessToken !== null))){
            return <></>
          }
          return (
            <Link key={route.name} href={route.link}>
              <a
                className={clsx(
                  'py-2 px-4 rounded-md transition-all font-semibold relative inline-flex',
                  'hover:bg-primary-100 dark:hover:bg-dark-500/50',
                  route.link === pathname
                    ? 'text-primary-700 dark:text-primary-400'
                    : 'text-gray-700 dark:text-gray-300'
                )}
              >
                {route.name}
              </a>
            </Link>
          );
        })
      }
    </nav>
  );
};

export default Navbar;
