interface Queryacc extends Queryaccdata {
  code: number;

  msg: string;
}

interface Queryaccdata {
  data: [{ accountId: string; balance: number; userName: string }];
}
