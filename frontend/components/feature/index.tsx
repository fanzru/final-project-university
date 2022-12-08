
import router from 'next/router';
const Feature = () => {
  
  return (
    <>
      <div className="mt-10 flex justify-center min-h-screen">
        <div className="grid grid-cols-2 max-w-[1400px] h-[200px] mt-20 gap-5">

          <div className="card w-96 bg-primary text-primary-content border-2 border-dashed">
            <div className="card-body">
              <h2 className="text-black dark:text-white">Summarizer</h2>
              <p className="text-black dark:text-white ">Artificial intelligence text summarization abstractive or generative tools</p>
              <div className="card-actions justify-end">
                <button 
                  className="btn"
                  onClick={()=> {router.push('/summary')}}
                >Use Now</button>
              </div>
            </div>
          </div>

          <div className="card w-96 bg-primary text-primary-content border-2 border-dashed">
            <div className="card-body">
              <h2 className="text-black dark:text-white">Annotation Tools</h2>
              <p className="text-black dark:text-white ">A tool for annotating in the creation of a summary dataset of a science article</p>
              <div className="card-actions justify-end">
                <button 
                  className="btn"
                  onClick={()=> {router.push('/submit')}}
                >Use Now</button>
              </div>
            </div>
          </div>

        </div>
      </div>
    </>
  );
}

export default Feature;

