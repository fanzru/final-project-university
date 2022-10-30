import Table from './components/table'
import {axiosInstance} from '../../lib/axios';
import {useState,useEffect} from 'react';
import {Profile} from "../../types/profile"
const Profile = () => {
  if (typeof window !== 'undefined') {
    var token = localStorage.getItem('token');
  }
  const [dataUser, setDataUser] = useState<Profile>();
  const getProfil = () => {
    axiosInstance
      .get('/accounts/profile', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((res) => {
        setDataUser(res.data.data);
      })
      .catch((err) => {
        console.log(err);
      });
  };
  useEffect(() => {
  

    getProfil();
  }, []);

  return(
    <div className=" flex justify-center">
      <div className="max-w-[1000px] w-full">
        <div className="mt-20 flex flex-col md:flex-row items-center">
          <div className="max-w-[300px] w-full h-[300px] bg-black text-white flex items-center justify-center font-bold text-[100px] dark:bg-white dark:text-dark-900 text-">
            {dataUser?.name[0].toUpperCase()}
          </div>
          <div className="md:ml-10 text-center md:text-left ">
            <div className="mt-10 text-xl font-bold">
              {dataUser?.name}
            </div>
            <div className="">
              {dataUser?.email}
            </div>
            <div>
              {`Halo kamu, terimakasih sudah menjadi bagian dari penelitian ini ya :)`}
            </div>
          </div>
        </div>
        <div className="mt-5 px-5 md:px-0">
          <Table data={dataUser?.papers_users}/>
        </div>
      </div>
    </div>
  )
}


export default Profile;