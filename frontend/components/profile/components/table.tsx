import Link from 'next/link';
import {PapersUsers} from "../../../types/profile"
import { FC } from 'react';
import {axiosInstance} from '../../../lib/axios';
import { AiFillEdit,AiOutlineCloudDownload } from 'react-icons/ai';
import {exportData} from '../../../lib/exportData';
import { toast } from 'react-toastify';
interface TableInterface {
  data: PapersUsers[] | any;
}
const Table:FC<TableInterface> = ({data})=> {
  if (typeof window !== 'undefined') {
    var token = localStorage.getItem('token');
  }
  const getCsv = (id: number) => {
    axiosInstance
      .get(`/grobid/detail-paper-csv/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((res) => {
        console.log(res.data.data)
        exportData(res.data.data.sentences_label, 'label_' + res.data.data.paper_detail.paper_name);
      })
      .catch((err) => {
        console.log(err);
        toast.error('Download Error!');
      });
  };
  return (
    <div>
      <div className="overflow-x-auto max-h-[300px]">
        <table className="table w-full">
          <thead>
            <tr>
              <th></th>
              <th>Paper Title</th>
              <th>Features</th>
            </tr>
          </thead>
          <tbody>
            {
              data?.map((v: PapersUsers,index: number)=>{
                return (
                  <>
                    <tr >
                      <th>{index+1}</th>
                      <td>{v.paper_name}</td>
                      <td>
                        <div className='flex gap-2 items-center'>
                          <Link href={v.link_pdf}>
                            <a
                              className='btn btn-primary  btn-xs gap-1'
                              target={'_blank'}
                            >
                              <svg
                                xmlns='http://www.w3.org/2000/svg'
                                className='h-4 w-4'
                                fill='none'
                                viewBox='0 0 24 24'
                                stroke='currentColor'
                                strokeWidth='2'
                              >
                                <path
                                  strokeLinecap='round'
                                  strokeLinejoin='round'
                                  d='M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z'
                                />
                              </svg>
                              show pdf
                            </a>
                          </Link>
                          <Link href={`/annotation?paper_id=${v.id}`}>
                            <a
                              className='btn btn-primary btn-xs gap-1 text-white'
                            >
                              <AiFillEdit/>
                              Edit
                            </a>
                          </Link>
                          <button
                            onClick={()=>{getCsv(v.id)}}
                          >
                            <a
                              className='btn btn-primary btn-xs gap-1 text-white'
                            >
                              <AiOutlineCloudDownload/>
                              DOWNLOAD
                            </a>
                          </button>
                          <div>
                            {
                              !v.is_done? 
                              <div className="badge bg-gray-500 border-gray-500 gap-2 font-bold text-xs">
                                DRAFT
                              </div>: <></>
                            }
                          </div>
                          
                        </div>
                        
                      </td>
                    </tr>
                  </>
                );
              })
            }
            
            
          </tbody>
        </table>
      </div>
    </div>
  )
}

export default Table;