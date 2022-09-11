/** @type {import('tailwindcss').Config} */
const { fontFamily } = require('tailwindcss/defaultTheme');
const colors = require('tailwindcss/colors');
module.exports = {
  content: [
    './pages/**/*.{js,ts,jsx,tsx}',
    './components/**/*.{js,ts,jsx,tsx}',
  ],
  darkMode: 'class',
  theme: {
    fontFamily: {
      inter: ['Inter', 'sans-serif'],
    },
    extend: {
      fontFamily: {
        sans: ['Inter', ...fontFamily.sans],
      },
      colors: {
        primary: colors.indigo,
        secondary: colors.pink,
        ternary: colors.purple,
        dark: {
          100: '#A6A7AB',
          200: '#909296',
          300: '#5C5F66',
          400: '#373A40',
          500: '#2C2E33',
          600: '#25262B',
          700: '#1A1B1E',
          800: '#141517',
          900: '#101113',
        },
      },
    },
  },
  plugins: [],
}
