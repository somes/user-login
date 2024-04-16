import axios from "axios";
import {_getToken, _removeToken} from "../utils";

const request = axios.create({
  baseURL: "http://localhost:8000",
  timeout: 5000
});

// 请求拦截器
request.interceptors.request.use((config) => {
  // 注入 token
  const token = _getToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
}, (error) => {
  return Promise.reject(error);
});

// 响应拦截器
request.interceptors.response.use((response) => {
  return response.data;
}, (error) => {
  switch (error.response.status) {
    case 401:
      _removeToken();
      window.location.reload();
      break;
    default:
      break;
  }
  return Promise.reject(error);
});

export {request};
