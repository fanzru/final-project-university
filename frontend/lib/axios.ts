import axios from 'axios';

export const axiosInstance = axios.create({
  baseURL: 'https://api.skripsi.fanzru.dev',
  // baseURL: 'http://localhost:8888',
});
