import axios from 'axios';

export const axiosInstanceFast = axios.create({
  baseURL: 'https://fastapi.skripsi.fanzru.dev',
});
