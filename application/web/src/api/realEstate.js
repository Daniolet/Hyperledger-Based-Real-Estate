import request from "@/utils/request";

// New real estate(administrator)
export function createRealEstate(data) {
  return request({
    url: "/createRealEstate",
    method: "post",
    data,
  });
}

// Get real estate information(Sky JSON{}You can query all, specify the protrietor to query the designated property under the name of the owner)
export function queryRealEstateList(data) {
  return request({
    url: "/queryRealEstateList",
    method: "post",
    data,
  });
}
