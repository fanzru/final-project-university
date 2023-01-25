import Table from './components/table'
import {axiosInstance} from '../../lib/axios';
import {exportData} from '../../lib/exportData';
import {useState,useEffect} from 'react';
import {Profile} from "../../types/profile"
import { useRouter } from 'next/router';
import { toast } from 'react-toastify';
const Profile = () => {
  const router = useRouter();
  if (typeof window !== 'undefined') {
    var token = localStorage.getItem('token');
  }
  const [dataUser, setDataUser] = useState<Profile>();
  if (typeof window !== 'undefined') {
    var token = localStorage.getItem('token');
  }
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

  const getCsv = (id: number) => {
    axiosInstance
      .get(`/grobid/detail-paper-csv/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((res) => {
        exportData(res.data.value, 'label_' + res.data.data.paper_detail.paper_name);
      })
      .catch((err) => {
        if (err.response.data.code === 400)
          toast.error(err.response.data.message);
        else toast.error('Download Error!');
      });
  };
 

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