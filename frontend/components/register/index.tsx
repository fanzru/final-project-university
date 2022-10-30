
import { useForm } from 'react-hook-form';
import Link from 'next/link'
import { toast } from 'react-toastify';
import { axiosInstance } from '../../lib/axios';
import { useRouter } from 'next/router';

type RegisterType = {
  name: string;
  email: string;
  password: string;
};


const Register = () => {
  const router = useRouter();
  const showPassword = () => {
    var x = (document.getElementById("myInput") as HTMLInputElement);
    if (x.type === "password") {
      x.type = "text";
    } else {
      x.type = "password";
    }
  }
  
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterType>();
  
  const onSubmit = handleSubmit(async (data) => {
    try {
      toast.dismiss();
      const auth = await toast.promise(
        axiosInstance.post('/accounts/register', {
          ...data,
        }),
        {
          pending: 'Loading..',
          success: 'Login Success!',
          error: 'Login Failed!',
        }
      );

      if (auth?.data.message==="success_created") {
        router.push('/login');
      }

     
    } catch (e) {
      toast.info("Login Failed!")
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
              Register
            </h1>
            <label className="label">
              <span className="label-text font-bold">Name</span>
              {errors.name && (
                  <span className='label-text-alt text-error'>
                    {errors.name?.message}
                  </span>
                )}
            </label>
            <input 
              type="text"
              placeholder="Type here"
              className="input input-bordered w-full dark:bg-dark-500"
              {...register('name', {
                required: 'Name is Required',
              })}
            />
            <label className="label">
              <span className="label-text font-bold">Email</span>
              {errors.email && (
                  <span className='label-text-alt text-error'>
                    {errors.email?.message}
                  </span>
                )}
            </label>
            <input 
              type="text"
              placeholder="Type here"
              className="input input-bordered w-full dark:bg-dark-500"
              {...register('email', {
                required: 'Email is Required',
                pattern: /^\S+@\S+$/i,
              })}
            />
            <label className="label">
              <span className="label-text font-bold">Password</span>
              {errors.password && (
                  <span className='label-text-alt text-error'>
                    {errors.password?.message}
                  </span>
                )}
            </label>
            <input 
              type="password" 
              placeholder="Type here" 
              className="input input-bordered w-full dark:bg-dark-500" 
              id="myInput" 
              {...register('password', {
                required: 'Password is Required',
                min: 8,
              })}
            />
            <div className="form-control flex-row mt-2">
              <label className="label cursor-pointer">
                <input  
                  type="checkbox"
                  className="checkbox dark:checkbox-primary mr-2"
                  onClick={showPassword}
                 />
                <div className="label-text">Show Password</div> 
              </label>
            </div>
            <button type='submit' className="btn btn-active dark:btn-primary">Button</button>
            <div className="mt-5">
              {`Already have an account? `}
                <span className="mr-3">
                  <Link href={`/login`}>
                  <a className="underline font-bold">Login</a>
                  </Link>
                </span> 
            </div>
          </form>  
        </div>
      </div>
      </div>
    </div>
  );
};

export default Register;