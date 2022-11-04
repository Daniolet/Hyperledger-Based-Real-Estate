interface Queryacc extends Queryaccdata {
  code: number;

  msg: string;
}

interface Queryaccdata {
  data: [{ accountId: string; balance: number; userName: string }];
}

interface QueryRealEstate {
  encumbrance: Boolean;
  livingSpace: number;
  proprietor: string;
  realEstateId: string;
  totalArea: number;
}

interface QuerySelling {
  buyer: string;
  createTime: string;
  objectOfSale: string;
  price: number;
  salePeriod: number;
  seller: string;
  sellingStatus: string;
}

interface QueryDonating {
  createTime: string;
  donatingStatus: string;
  donor: string;
  grantee: string;
  objectOfDonating: string;
}
