import axios from 'axios'
import Constant from '../Constant';
import router from '../router/index';
import { store } from "../store";

export default (URL) => {

  const InstanceAxios = axios.create({
    baseURL: URL
  })

  if (localStorage.getItem('tokenAdmin')) {
    const token = localStorage.getItem('tokenAdmin');
    InstanceAxios.defaults.headers.common.Authorization = `Bearer ${token}`
  } else {
    const token = localStorage.getItem('token');
    InstanceAxios.defaults.headers.common.Authorization = `Bearer ${token}`
  }

  InstanceAxios.interceptors.response.use(
    (response) => response,
    (error) => {
      if (error.response.status === 401 && localStorage.getItem('token')) {
        store.dispatch('user/' + Constant.LOGOUT);
        router.push({ name: "login" });
      }

      return Promise.reject(error)
    }
  )


  return InstanceAxios
}