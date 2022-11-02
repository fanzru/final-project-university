import axios from 'axios';

export const axiosInstance = axios.create({
  baseURL: 'https://api.skripsi.fanzru.dev',
});
