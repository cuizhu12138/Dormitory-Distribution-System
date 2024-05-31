// src/utils/http.js

import axios from 'axios';
import router from '@/router';
import { requestUrl } from './contants';

let config = {
  // baseURL: process.env.baseURL || process.env.apiUrl || ""
  // timeout: 60 * 1000, // Timeout
  // withCredentials: true, // Check cross-site Access-Control
};

// 创建 axios 实例
const request = axios.create(config);

// 请求拦截器
request.interceptors.request.use(
  config => {
    if (["/login", "/register"].includes(config.url)) {
      config.url = requestUrl + config.url;
     return config  // 登录注册接口不需要token
    }
    config.url = requestUrl + config.url;
    // 在发送请求之前做些什么
    const token = localStorage.getItem('token');
    // const token = store.getters['auth/token']; // 从 Vuex 获取 token
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`; // 让每个请求携带自定义 token
    }
    else {
      router.push({ path: '/login' });
      return Promise.reject(new Error('Token is null, please login again'));
    }
    return config;
  },
  error => {
    // 对请求错误做些什么
    console.error('Request error: ', error);
    return Promise.reject(error);
  }
);

// 响应拦截器
request.interceptors.response.use(
  response => {
    // 对响应数据做点什么
    return response.data;
  },
  error => {
    // 对响应错误做点什么
    console.error('Response error: ', error);
    if (error.response && error.response.status === 401) {
      // 如果响应状态码为 401，跳转到登录页
      router.push({ name: 'login' });
    }
    return Promise.reject(error);
  }
);

export default request;
