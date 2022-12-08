import React from 'react';

const QuickTo = () => {
  return (
    <>
      <ul className='list-decimal px-2 ml-2'>
        <li>
          {"Mark the checkbox if you feel the sentence above the checkbox is important and will be part of the pdf summary that you have input."}
        </li>
        <li>
          {"Uncheck the checkbox if you feel the sentence is not important and will not be part of the pdf summary that you have input."}
        </li>
        <li>
          {"Press the save button to save your current progress and make it a draft."}
        </li>
        <li>
          {"Press the submit button to complete your annotation process."}
        </li>
      </ul>
    </>
  );
};

export default QuickTo;
