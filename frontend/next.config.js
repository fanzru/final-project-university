/** @type {import('next').NextConfig} */
module.exports = {
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: 'http://localhost:3000/:path*'
       //destination: 'https://is3.cloudhost.id/:path*',
      },
    ]
  },
  reactStrictMode: true,
};
