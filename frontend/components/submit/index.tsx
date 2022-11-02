
import { useForm } from 'react-hook-form';
import Link from 'next/link'
import { toast } from 'react-toastify';
import { axiosInstance } from '../../lib/axios';
import { useRouter } from 'next/router';
import {useEffect,useState} from 'react';
import { FileUploader } from 'react-drag-drop-files';
import axios from 'axios';
type SubmitPaperType = {
  pdf_title: string;
};


const SubmitPaper = () => {
  const router = useRouter();
  const fileTypes = ['CSV', 'PDF'];
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SubmitPaperType>();
  if (typeof window !== 'undefined') {
    var token = localStorage.getItem('token');
  }
  const [isSetFile, SetIsSetFile] = useState(false);
  const [file, setFile] = useState<File>();
  const handleChange = (file: any) => {
    SetIsSetFile(true);
    setFile(file);
  };
  
  const onSubmit = handleSubmit(async (data) => {
    try {
      toast.dismiss();
      let bodyFormData = new FormData();
      bodyFormData.append('pdf_name', data.pdf_title);
      bodyFormData.append('pdf_file',file!)
      console.log("pdf_name : ", data.pdf_title);
      const res = await toast.promise(
        axios({
          method: 'POST',
          url: 'http://localhost:8888/grobid/pdf-to-tei',
          data: bodyFormData,
          headers: { Authorization: `Bearer ${token}` },
        }),
        {
          pending: 'Loading..',
          success: 'Upload PDF Success!',
          error: 'Upload PDF Failed!',
        }
      );
    
      if (res?.data.message==="success_ok") {
        router.push(`/annotation?paper_id=${res?.data.data.paper_id}`);
      }

     
    } catch (e) {
      console.log(e);
    }
  });

  return (
    <div className=" flex justify-center">
      <div className="w-[1000px]">
      <div className="mt-20 flex justify-center px-10">
        <div className="w-[500px] h-[550px]  mt-20 mb-20 rounded-lg shadow-md shadow-md dark:shadow-indigo-500/50">
          <form onSubmit={onSubmit} className="form-control w-full border-3 border-white px-10">
            <h1 className="text-center mb-10 mt-5 font-xl">
              Upload Your Paper
            </h1>
            <label className="label">
              <span className="label-text font-bold">Paper Titel</span>
              {errors.pdf_title && (
                  <span className='label-text-alt text-error'>
                    {errors.pdf_title?.message}
                  </span>
                )}
            </label>
            <input 
              type="text"
              placeholder="Type here"
              className="input input-bordered w-full dark:bg-dark-500"
              {...register('pdf_title', {
                required: 'Paper title is Required',
              })}
            />
            
            
            
            <label className="label">
              <span className="label-text font-bold">Please choose PDF article</span>
            </label>
            <FileUploader
                handleChange={handleChange}
                name='file'
                types={fileTypes}
              >
                <button
                  onClick={(e) => e.preventDefault()}
                  className='h-[200px] border-2 w-[100%] rounded-md border-dashed opacity-60'
                >
                  {isSetFile ? (
                    <>
                      <p>Paper Added</p>
                      <p className='font-medium'>{file?.name}</p>
                    </>
                  ) : (
                    <p>Upload Here</p>
                  )}
                </button>
              </FileUploader>
            <button className="btn btn-active dark:btn-primary mt-5">Upload</button>
         
          </form>  
        </div>
      </div>
      </div>
    </div>
  );
};

export default SubmitPaper;
