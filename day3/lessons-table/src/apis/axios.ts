import axios from "axios";

export function myAxios(axiosConfig: any) {
  const service = axios.create({
    // baseURL: "https://mock.apifox.cn/m1/1379345-0-default",
    baseURL: "http://localhost:8888",
    timeout: 10000,
  });
  return service(axiosConfig);
}
