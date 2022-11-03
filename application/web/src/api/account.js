import request from "@/utils/request";

// Get the list of character selection of the login interface
export function queryAccountList() {
  return request({
    url: "/queryAccountList",
    method: "post",
  });
}

// Log in
export function login(data) {
  return request({
    url: "/queryAccountList",
    method: "post",
    data,
  });
}
