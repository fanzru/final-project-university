export const variants = {
  hidden: {
    y: 15,
    opacity: 0,
  },
  enter: { opacity: 1, x: 0, y: 0 },
  exit: {
    y: 0,
    opacity: 0,
    transition: {
      type: 'tween',
      duration: 0.25,
    },
  },
};
