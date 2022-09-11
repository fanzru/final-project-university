import clsx from 'clsx';
import React, { createElement } from 'react';

const Button: React.FunctionComponent<
  React.DetailedHTMLProps<
    React.ButtonHTMLAttributes<HTMLButtonElement>,
    HTMLButtonElement
  >
> = ({ children, className, ...props }) => {
  return createElement(
    'button',
    {
      ...props,
      className: clsx('inline-flex items-center justify-center', className),
    },
    children
  );
};

export default Button;
