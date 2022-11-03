import request from "@/utils/request";

// Query donation list (you can query everything, or you can also query according to the initiative)
export function queryDonatingList(data) {
  return request({
    url: "/queryDonatingList",
    method: "post",
    data,
  });
}

// Inquire the donation (given by the gift) (the gift) (the gift) according to the gift of the gift (the gift)
export function queryDonatingListByGrantee(data) {
  return request({
    url: "/queryDonatingListByGrantee",
    method: "post",
    data,
  });
}

// Update the donation status (confirmed the gift and cancel) STATUS value is to complete "Done" and cancel "Cancelled"
export function updateDonating(data) {
  return request({
    url: "/updateDonating",
    method: "post",
    data,
  });
}

// Initiate a donation
export function createDonating(data) {
  return request({
    url: "/createDonating",
    method: "post",
    data,
  });
}
