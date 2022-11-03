package model

// Selling Sales offer
// Need to determine whether Objectofsale belongs to Seller
// Buyer initially empty
// Seller and Objectofsale together as the composite key,Guarantee that all sales initiated in the name can be queried by Seller
type Selling struct {
	ObjectOfSale  string  `json:"objectOfSale"`  //Sales target (Realestateid, real estate is being sold)
	Seller        string  `json:"seller"`        //Initiative, seller, seller(Seller Accountid)
	Buyer         string  `json:"buyer"`         //Participate in sellers, buyers(Buyer Acountid)
	Price         float64 `json:"price"`         //price
	CreateTime    string  `json:"createTime"`    //Creation time
	SalePeriod    int     `json:"salePeriod"`    //The validity period of smart contracts(The unit is the sky)
	SellingStatus string  `json:"sellingStatus"` //Sales status
}

// SellingStatusConstant Sales statuss statuss status
var SellingStatusConstant = func() map[string]string {
	return map[string]string{
		"saleStart": "Sale",      //Sales statussWaiting for buyers to patronizeWaiting for buyers to patronizeWaiting for buyers to patronize
		"cancelled": "Cancelled", //The seller cancels the sales or buyer's refund operation, causing cancellation the sales or buyer's refund operation, causing cancellation the sales or buyer's refund operation, causing cancellation
		"expired":   "expired",   //Sales period expireperiod expireperiod expire
		"delivery":  "Delivery",  //Buyers buy and paybBeing waiting for the seller to confirm the receiving statuseIf the seller fails to confirm the receipt, the buyer can cancel and refund the refundseller to confirm the receiving statuseIf the seller fails to confirm the receipt, the buyer can cancel and refund the refundseller to confirm the receiving status,If the seller fails to confirm the receipt, the buyer can cancel and refund the refund
		"done":      "Finish",    //The seller confirms the receipt of funds, the transaction is completednfirms the receipt of funds, the transaction is completednfirms the receipt of funds, the transaction is completed
	}
}

// Donating Donation offertion offertion offer
// Need to determine whether ObjectOfdonating belongs to DONOR
// You need to specify the gift from Granteee, and wait for the receiver to agree to receive
type Donating struct {
	ObjectOfDonating string `json:"objectOfDonating"` //Donation(Realestateid, real estate that is donating)
	Donor            string `json:"donor"`            //Donors(Donor Acountid)
	Grantee          string `json:"grantee"`          //Recipient(ACCOUNTID)
	CreateTime       string `json:"createTime"`       //Creation time
	DonatingStatus   string `json:"donatingStatus"`   //Donation state
}

// DonatingStatusConstant Donation state
var DonatingStatusConstant = func() map[string]string {
	return map[string]string{
		"donatingStart": "Donation",  //The donor initiates the donation contract, waiting for the receiver to confirm the gift
		"cancelled":     "Cancelled", //The donor cancels the donation or the gift before the receiver is confirmed to cancel the receiver
		"done":          "Finish",    //The gift is confirmed to receive and the transaction is completed
	}
}
