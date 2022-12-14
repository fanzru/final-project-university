
import { useForm } from 'react-hook-form';
import { toast } from 'react-toastify';
import { axiosInstanceFast } from '../../lib/axiosfast';
import {useState} from 'react';
import {useEffect} from 'react';
import {axiosInstance} from '../../lib/axios';
import router from 'next/router';

type SummaryType = {
  text: string;
  max_length: number;
};


const Login = () => {
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

      })
      .catch((err) => {
        console.log(err);
        router.push("/login")
      });
  };
  useEffect(() => {
    getProfil();
  }, []);
  const [summaryResult, setSummaryResult] = useState("");
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SummaryType>();

  const onSubmit = handleSubmit(async (data) => {
   
    try {
      toast.dismiss();
      const dataSummary = await toast.promise(
        axiosInstanceFast.post('/summarize/', {
          ...data,
        }),
        {
          pending: 'Loading..',
          success: 'Summarize Success!',
          error: 'Summarize Failed!',
        }
      );

     
      setSummaryResult(dataSummary.data.result)
     
    } catch (e) {
      console.log(e);
    }
  });

  return (
    <div className=" flex justify-center">
      <div className="w-full grid md:grid-cols-2">
        <div className="mt-5 flex justify-center px-5 ">
          <div className="max-w-[1000px] w-full min-h-[450px] mt-20 mb-20 rounded-lg ">
            <form onSubmit={onSubmit} className="form-control w-full border-3 border-white px-10">
              <h1 className="text-center mb-5 mt-5 font-xl">
                Text
              </h1>
              <label className="label">
                <span className="label-text font-bold">Text</span>
                {
                  errors.text && (
                    <span className='label-text-alt text-error'>
                      {errors.text?.message}
                    </span>
                  )
                }
              </label>
              <textarea
                placeholder="Type here"
                className="textarea input input-bordered w-full dark:bg-dark-500 h-[300px]" 
                {...register('text', {
                  required: 'text is Required',
                })}
              />
              <label className="label">
                <span className="label-text font-bold">Max Length</span>
                {errors.max_length && (
                    <span className='label-text-alt text-error'>
                      {errors.max_length?.message}
                    </span>
                  )}
              </label>
              <input 
                type="number" 
                placeholder="Default value 512"
                className="input input-bordered w-full dark:bg-dark-500" 
                id="myInput" 
                {...register('max_length', {
                  required: 'max_length is Required',
                })}
              />
              <button type='submit' className="btn btn-active dark:btn-primary mt-5">Submit</button>
            </form>  
          </div>
        </div>
        <div className="w-full mt-5 px-5 flex justify-center">
          <div className="max-w-[1000px] w-full min-h-[450px] mt-20 mb-20 rounded-lg">
            <div className="form-control w-full border-3 border-white px-10">
              <h1 className="text-center mb-5 mt-5 font-xl">
                Result
              </h1>
              <label className="label">
                <span className="label-text font-bold">Result Summary</span>
              </label>
              <div 
         
                className="textarea  input-bordered w-full dark:bg-dark-500 h-[300px]" 
                
              >
                {summaryResult}
              </div>
            </div>  
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;
