
import {axiosInstance} from '../../lib/axios';
import {useState,useEffect} from 'react';
import {DetailPaper,Body} from "../../types/detailpaper"
import {FC} from 'react';
import QuickTo from './components/quick_to';
import Guidelines from './components/guidelines';
import CardCollapse from './components/card_collapse';
import { toast } from 'react-toastify';
import '@react-pdf-viewer/default-layout/lib/styles/index.css';
import { SpecialZoomLevel, Viewer } from '@react-pdf-viewer/core';
import { defaultLayoutPlugin } from '@react-pdf-viewer/default-layout';
import clsx from 'clsx';
import { useRouter } from 'next/router';
import { totalmem } from 'os';

type AnnotationInterface = {
  paper_id: any;
}

const Annotation:FC<AnnotationInterface>  = ({paper_id}) => {
  if (typeof window !== 'undefined') {
    var token = localStorage.getItem('token');
  }
  const router = useRouter();
  const [isShowPdf, setIsShowPDF] = useState<boolean>(false);
  const [page, setPage] = useState<number>(1);
  const [dataDetailPaper, setDataDetailPaper] = useState<DetailPaper>();
  const defaultLayoutPluginInstance = defaultLayoutPlugin();

  const onSubmit = (async (data: any) => {
    try {
      const res = await toast.promise(
        axiosInstance.post(
          `/grobid/paper-edit/${false}`,
          dataDetailPaper,
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        ),
        {
          pending: 'Loading..',
          success: 'Saved Success!',
          error: 'Save Failed!',
        }
      );
      router.push(`/annotation?paper_id=${dataDetailPaper?.paper_id}`)
    } catch (err) {
      console.log("Saved Paper : ",err)
    }
  });

  const onSubmitLast = (async (data: any) => {
    try {
      const res = await toast.promise(
        axiosInstance.post(
          `/grobid/paper-edit/${true}`,
          dataDetailPaper,
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        ),
        {
          pending: 'Loading..',
          success: 'Submit Success!',
          error: 'Submit Failed!',
        }
      );
      router.push(`/profile`)
    } catch (err) {
      console.log("Submit Paper : ",err)
    }
  });

  const getDetailPaper = () => {
    axiosInstance
    .get(`/grobid/detail-paper/${paper_id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
    .then((res) => {
      setDataDetailPaper(res.data.data);
    })
    .catch((err) => {
      console.log(err);
    })
    
    
  };

  useEffect(() => {
    getDetailPaper();
  }, [paper_id]);

  return(
    <div className=" flex justify-center">
      <div className="max-w-[2000px] mt-20 flex px-10">
        {
          isShowPdf? 
            <div className="border-2 w-1/2 max-h-[1000px]">
              {
                dataDetailPaper?.link_pdf?
                <Viewer
                  fileUrl={dataDetailPaper?.link_pdf}
                  plugins={[defaultLayoutPluginInstance]}
                  defaultScale={SpecialZoomLevel.PageFit}
                />
                :
                <></>
              }
            </div>
          : 
            <></>
        }
        <div className={clsx(isShowPdf? "w-1/2 ml-2" : "")}>
          <div>
            <button className="btn btn-primary w-full mb-5" onClick={() => {
              setIsShowPDF(!isShowPdf)
              }}>
              {
                isShowPdf ? "Hide PDF" : "Show PDF"
              }
            </button>
          </div>
          <div>
            <button className="btn btn-primary w-full mb-5" onClick={() => {
              onSubmit(dataDetailPaper)
              }}>
              {
                "Save"
              }
            </button>
          </div>
          <CardCollapse title='Guidelines'>
            <QuickTo/>
          </CardCollapse>
          <div className="border-2 p-5 rounded-lg max-h-[880px] overflow-auto">
            {
              dataDetailPaper?.body.map((dataDetailPaper: Body) =>{
                if (dataDetailPaper.head_key === page) {
                  return (
                    <div>
                      <div className="font-bold">
                        {
                          dataDetailPaper.head
                        }
                        
                      </div>
                      <div>
                        {
                          dataDetailPaper.sentences.map((sent)=> {
                            return (
                              <div className="border-2 p-5 mt-4">
                                <div className=" ">
                                  {
                                    sent.text
                                  }
                                </div>
                                <div className="form-control flex-row mt-2">
                                  <label className="label cursor-pointer">
                                    <input  
                                      type="checkbox"
                                      className="checkbox dark:checkbox-primary mr-2"
                                      defaultChecked={sent.is_important}
                                      onChange={(e)=>{
                                        e.defaultPrevented
                                        sent.is_important = e.target.checked
                                      }}
                                    />
                                    <div className="label-text">Is Important</div> 
                                  </label>
                                </div>
                              </div>
                            )
                          })
                        }
                      </div>
                    </div>
                  )
                }
              })
            }
          </div>
          <div className="flex justify-between">
            <div className="mt-10 mb-2">
              <button 
                className="btn btn-primary"
                onClick={()=>{
                  if (page >= 2){
                    setPage(page - 1);
                  }
                }}
              >
                Prev
              </button>
            </div>
            <div className="mt-10 mb-2">
              <button 
                className="btn btn-primary"
                onClick={()=>{
                  if (page < (dataDetailPaper?.len_head === null ? 1 : dataDetailPaper!.body.length) ){
                    setPage(page + 1);
                  }
                }}
              >
                Next
              </button>
            </div>
            
          </div>
          <button 
            className="btn btn-primary w-full mb-10"
            onClick={() => {
              onSubmitLast(dataDetailPaper)
            }}
          
          >
              Submit
          </button>
        </div>
      </div>
    </div>
  )
}


export default Annotation;