import request from "@/utils/request";

// Query sales(You can query everything, or you can inquire according to the initiator)(Initiative)
export function querySellingList(data) {
  return request({
    url: "/querySellingList",
    method: "post",
    data,
  });
}

// According to participating sellers, buyers(Buyer Acountid)Query sales(attended)
export function querySellingListByBuyer(data) {
  return request({
    url: "/querySellingListByBuyer",
    method: "post",
    data,
  });
}

// 买家购买
export function createSellingByBuy(data) {
  return request({
    url: "/createSellingByBuy",
    method: "post",
    data,
  });
}

// 更新销售状态（买家确认、买卖家取消）Status取值为 完成"done"、取消"cancelled" 当处于销售中状态，卖家要取消时，buyer为""空
export function updateSelling(data) {
  return request({
    url: "/updateSelling",
    method: "post",
    data,
  });
}

// 发起销售
export function createSelling(data) {
  return request({
    url: "/createSelling",
    method: "post",
    data,
  });
}
