import axios from "axios";
import { MessageBox, Message } from "element-ui";

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 5000,
});

service.interceptors.response.use(
  (response) => {
    const res = response.data;
    if (res.code !== 200) {
      MessageBox.alert("The server is small", "error", {
        confirmButtonText: "Sure",
        type: "warning",
      });
      return Promise.reject(new Error(res.msg || "Error"));
    } else {
      return res.data;
    }
  },
  (error) => {
    if (error.response === undefined) {
      Message({
        message: "Request failed " + error.message,
        type: "error",
        duration: 5 * 1000,
      });
      return Promise.reject(error);
    } else {
      Message({
        message: "fail " + error.response.data.data,
        type: "error",
        duration: 5 * 1000,
      });
      return Promise.reject(error.response);
    }
  }
);

export default service;
