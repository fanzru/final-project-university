import clsx from 'clsx';
import { useRouter } from 'next/router';
import {axiosInstance} from '../../lib/axios';
import DrawerToggler from './components/drawertoggler';
import DarkMode from './components/darkmode';
import Navbar from './components/navbar';
import Button from './components/button';
import { HiOutlineLogout } from 'react-icons/hi';
import {useEffect} from 'react';
import {toast} from 'react-toastify';
const Header = () => {
  const router = useRouter();
  const { pathname } = useRouter();
  const isError =
    pathname === '/_error' || pathname === '/_offline' || pathname === '/404';

  if (isError) {
    return null;
  }
  var token_access : any
  if (typeof window !== 'undefined') {
    token_access = localStorage.getItem('token');
  }
  const getProfil = () => {
    axiosInstance
      .get('/accounts/profile', {
        headers: {
          Authorization: `Bearer ${token_access}`,
        },
      })
      .then((res) => {
       
      })
      .catch((err) => {
        console.log(err);
        router.push("/")
      });
  };
  useEffect(() => {
    if (typeof window !== 'undefined') {
      var token = localStorage.getItem('token');
      if (token === "undefined") {
        toast.info("Session not found!")
        router.push("/login")
       
      }
    }
    getProfil();
  }, []);
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
       
        <div className="flex">
          {
            token_access !== null ? 
            <Button
            className={clsx(
              'text-lg rounded-md h-10 w-10',
              'bg-primary-50 dark:bg-dark-500/50 mr-2'
            )}
            onClick={() => {
              if (typeof window !== 'undefined') {
                token_access = localStorage.removeItem('token');
                router.push("/login")
              }
              
            }}
            aria-label='Darkmode Button'
          >
            
            <HiOutlineLogout className={clsx('text-red-400 dark:text-red-400')}/>
          </Button>: <></>
          }
          
          <DarkMode />
        </div>
      </div>
    </header>
  );
};

export default Header;
