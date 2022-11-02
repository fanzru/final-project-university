import React from 'react';

const QuickTo = () => {
  return (
    <>
      <ul className='list-decimal px-2 ml-2'>
        <li>
          The first step:{' '}
          <span className='font-semibold'>
            check if the section heading doesn't belong to one of the paper
            sections or subsections
          </span>
          : mark the checkbox{' '}
          <span className='font-semibold'>"Wrong extracted heading"</span>. If
          the section is wrong one, you can choose whether to complete
          annotation or jump directly to step 3.
        </li>
        <li>
          For each sentence selected (highlighled) from the paper section below,
          choose one action of the following:
          <ul className='list-decimal ml-5'>
            <li>
              <span className='font-semibold'>
                If the label/annotation of the sentence marked by the bullet
                point is correct
              </span>
              : mark the checkbox{' '}
              <span className='font-semibold'>"Correct label"</span>
            </li>
            <li>
              <span className='font-semibold'>
                If there is something wrong with the extraction of the sentence
                (for example: the sentence doesn't belong to the context of the
                section):
              </span>
              select the radio button{' '}
              <span className='font-semibold'>"Wrong extraction"</span>{' '}
            </li>
            <li>
              <span className='font-semibold'>
                If the label/annotation of the sentence marked by the bullet
                point is not correct:
              </span>{' '}
              choose the collect label from the bullet points
            </li>
          </ul>
        </li>
        <li>
          Press "Next section" to proceed in the next section with selected
          sentences from the paper
        </li>
        <li>
          Repeat the process until sections with selected sentences are done.
        </li>
      </ul>
      <ul className='list-disc mt-4 px-2 ml-2'>
        <li>
          The description of the sentences labels and guidelines are described
          in the "Guidelines section" below
        </li>
        <li>
          You can download the sentences selected with there annotation and your
          correct annotation using the "Download progress" button in the top
          left corner
        </li>
        <li>
          If you left the session the progress will not be recorded and you will
          need to upload the paper from the start.
        </li>
      </ul>
    </>
  );
};

export default QuickTo;
