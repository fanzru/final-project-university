import clsx from 'clsx';
import Link, { LinkProps } from 'next/link';
import React, { createElement } from 'react';

export interface UnstyledLinkProps extends LinkProps {
  href: string;
  title?: string;
  className?: string;
  children?: React.ReactNode;
}

const UnstyledLink: React.FunctionComponent<UnstyledLinkProps> = ({
  href,
  children,
  ...props
}) => {
  if (href.startsWith('http')) {
    return createElement(
      'a',
      { href, rel: 'noopener noreferrer', target: '_blank', ...props },
      children
    );
  }
  return (
    <Link href={href} scroll={false} {...props}>
      <a
        title={props.title ?? ''}
        onClick={props.onClick}
        className={clsx(props.className)}
      >
        {children}
      </a>
    </Link>
  );
};

export default UnstyledLink;
