package model

// Account Account, virtual administrator and several owner accounts
type Account struct {
	AccountId string  `json:"accountId"` //Account ID
	UserName  string  `json:"userName"`  //accountName
	Balance   float64 `json:"balance"`   //Balance
}

// RealEstate Real estate is True as a guarantee for sale, donation or pledge, and False defaults to the default state.
// Only when ENCUMBRANCE is false, it can be launched, donated or pledged
// Proprietor and Realestateid are used as composite keys to ensure
type RealEstate struct {
	RealEstateID string  `json:"realEstateId"` //Real estate ID
	Proprietor   string  `json:"proprietor"`   //Owner (owner) (owner accountidid)
	Encumbrance  bool    `json:"encumbrance"`  //Whether to be a guarantee
	TotalArea    float64 `json:"totalArea"`    //The total area
	LivingSpace  float64 `json:"livingSpace"`  //Living space
}

// Selling Sales offer
// Need to determine whether Objectofsale belongs to Seller
// Buyer initially empty
// Seller and ObjectOfsale together as a composite key to ensure that all sales initiated in the name can be queried by Seller
type Selling struct {
	ObjectOfSale  string  `json:"objectOfSale"`  //Sales target (Realestateid, real estate is being sold)
	Seller        string  `json:"seller"`        //Initiating seller and seller (seller accountidid)
	Buyer         string  `json:"buyer"`         //Participate in seller and buyer (buyer Academidid)
	Price         float64 `json:"price"`         //price
	CreateTime    string  `json:"createTime"`    //Creation time
	SalePeriod    int     `json:"salePeriod"`    //The validity period of smart contracts(The unit is the sky)
	SellingStatus string  `json:"sellingStatus"` //销售状态
}

// SellingStatusConstant Sales status
var SellingStatusConstant = func() map[string]string {
	return map[string]string{
		"saleStart": "Sale",      //Sales status,Waiting for buyers to patronize
		"cancelled": "Cancelled", //The seller cancels the sales or buyer's refund operation, causing cancellation
		"expired":   "expired",   //Sales period expire
		"delivery":  "Delivery",  //Buyers buy and pay,Being waiting for the seller to confirm the receiving status,If the seller fails to confirm the receipt, the buyer can cancel and refund the refund
		"done":      "Finish",    //The seller confirms the receipt of funds, the transaction is completed
	}
}

// SellingBuy Buyers participate in sales
// Sales objects cannot be initiated by buyers
// Buyer and CreateTime as a composite key,Guaranteed to query all the sales in the name through Buyer
type SellingBuy struct {
	Buyer      string  `json:"buyer"`      //Participate in sellers, buyers(Buyer Acountid)
	CreateTime string  `json:"createTime"` //Creation time
	Selling    Selling `json:"selling"`    //Sales targets
}

// Donating Donation offer
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

// DonatingGrantee Forwarding the gift of the gift
type DonatingGrantee struct {
	Grantee    string   `json:"grantee"`    //Recipient(Accountid)
	CreateTime string   `json:"createTime"` //Creation time
	Donating   Donating `json:"donating"`   //Donation
}

const (
	AccountKey         = "account-key"
	RealEstateKey      = "real-estate-key"
	SellingKey         = "selling-key"
	SellingBuyKey      = "selling-buy-key"
	DonatingKey        = "donating-key"
	DonatingGranteeKey = "donating-grantee-key"
)
