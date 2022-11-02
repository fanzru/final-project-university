import { FC, ReactNode } from 'react';

interface CardCollapseInterface {
  title: string;
  children: ReactNode;
}

const CardCollapse: FC<CardCollapseInterface> = ({ title, children }) => {
  return (
    <div className='my-w collapse w-full border rounded-md border-base-300 collapse-arrow mb-6'>
      <input type='checkbox' />
      <div className='collapse-title text-xl font-semibold text-blue-500'>
        {title}
      </div>
      <div className='collapse-content'>{children}</div>
    </div>
  );
};

export default CardCollapse;
