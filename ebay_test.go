package ebay

import (
	"fmt"
	"github.com/stvp/assert"
	"testing"
)

func TestEBay_ParseJSON(t *testing.T) {
	fmt.Printf("TestParseJSON\n")
	e := New("12345")
	il, err := e.ParseJSON(test_json)
	assert.Nil(t, err, "GetItemsFromJSON")
	assert.Equal(t, 9, len(il))
	i := il[0]
	assert.Equal(t, "322119314985", i.ID, "ID")
	assert.Equal(t, "Minneapolis,MN,USA", i.Location, "Location")
	assert.Equal(t, "3-D Tic-Tac-Toe (Atari 2600, 1980) CARTRIDGE ONLY! CLEAN & TESTED! ", i.Title, "item.Title")
	assert.Equal(t, 5.95, i.Price, "item.Price")
	assert.Equal(t, "http://www.ebay.com/itm/3-D-Tic-Tac-Toe-Atari-2600-1980-CARTRIDGE-ONLY-CLEAN-TESTED-/322119314985", i.Url, "item.URL")
	assert.Equal(t, "http://thumbs2.ebaystatic.com/m/msVJvxYHgLg8C29GpD6-ZfQ/140.jpg", i.ImageUrl, "Image URL")
	assert.Equal(t, 0.0, i.ShippingPrice, "Shipping Price")
	assert.Equal(t, "2016-05-23T20:41:25.000Z", i.StartTime, "Start Time")
	assert.Equal(t, "2016-06-22T20:41:25.000Z", i.EndTime, "End Time")
	assert.False(t, i.BuyItNow, "Buy it Now")
	assert.False(t, i.BestOffer, "Best Offer")
	assert.True(t, i.FreeShipping, "Free shipping")
	//3rd one has bestoffer true
	assert.True(t, il[3].BestOffer, "Best Offer")
	// only the 6th one has non-free shipping :)
	assert.Equal(t, 3.95, il[6].ShippingPrice, "ShippingPrice")
	assert.False(t, il[6].FreeShipping, "Free Shipping")


	il, err = e.ParseJSON(fail_json)
	assert.Nil(t, err, "GetItemsFromJSON")
	assert.Equal(t, 0, len(il))


	il, err = e.ParseJSON(empty_json)
	assert.Nil(t, err, "GetItemsFromJSON")
	assert.Equal(t, 0, len(il))
}
func TestLowestPrice(t *testing.T) {
	print("TestLowestPrice\n")
	e := New("12345")
	il, err := e.ParseJSON(test_json)
	assert.Nil(t, err, "e.ParseJson(test_json)")
	lp := LowestPrice(il)
	assert.Equal(t, 0.22, lp.Price, "lp.Price")
}

const fail_json = `{"findItemsByKeywordsResponse":[{"ack":["Whoopsie"],"version":["1.13.0"]}]}`
const empty_json = `{"findItemsByKeywordsResponse":[{"ack":["Success"],"version":["1.13.0"],"timestamp":["2016-06-15T02:48:25.565Z"],"searchResult":[{"@count":"0","item":[]}]}]}`
const test_json = `
{
"findItemsByKeywordsResponse":
[
{
"ack":["Success"],
"version":["1.13.0"],
"timestamp":["2016-06-15T02:48:25.565Z"],
"searchResult":
[
{
"@count":"22","item":
[
	{
	"itemId":["322119314985"],
	"title":["3-D Tic-Tac-Toe (Atari 2600, 1980) CARTRIDGE ONLY! CLEAN & TESTED! "],
	"globalId":["EBAY-US"],
	"primaryCategory":[{"categoryId":["139973"],"categoryName":["Video Games"]}],
	"galleryURL":["http:\/\/thumbs2.ebaystatic.com\/m\/msVJvxYHgLg8C29GpD6-ZfQ\/140.jpg"],
	"viewItemURL":["http:\/\/www.ebay.com\/itm\/3-D-Tic-Tac-Toe-Atari-2600-1980-CARTRIDGE-ONLY-CLEAN-TESTED-\/322119314985"],
	"productId":[{"@type":"ReferenceID","__value__":"56230072"}],
	"paymentMethod":["PayPal"],
	"autoPay":["false"],
	"postalCode":["55414"],
	"location":["Minneapolis,MN,USA"],
	"country":["US"],
	"shippingInfo":[
		{"shippingServiceCost":[{"@currencyId":"USD","__value__":"0.0"}],
		"shippingType":["Free"],
		"shipToLocations":["US"],
		"expeditedShipping":["false"],
		"oneDayShippingAvailable":["false"],
		"handlingTime":["1"]}
		],
	"sellingStatus":[
		{"currentPrice":[{"@currencyId":"USD","__value__":"5.95"}],
		"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"5.95"}],
		"sellingState":["Active"],
		"timeLeft":["P7DT17H53M0S"]}
		],
	"listingInfo":[
		{"bestOfferEnabled":["false"],
		"buyItNowAvailable":["false"],
		"startTime":["2016-05-23T20:41:25.000Z"],
		"endTime":["2016-06-22T20:41:25.000Z"],
		"listingType":["FixedPrice"],"gift":["false"]}
		],
	"returnsAccepted":["false"],
	"condition":[{"conditionId":["6000"],"conditionDisplayName":["Acceptable"]}],
	"isMultiVariationListing":["false"],
	"topRatedListing":["false"]
	},
	{
	"itemId":["322119314985"],
	"title":["3-D Tic-Tac-Toe (Atari 2600, 1980) CARTRIDGE ONLY! CLEAN & TESTED! "],
	"globalId":["EBAY-US"],
	"primaryCategory":[{"categoryId":["139973"],"categoryName":["Video Games"]}],
	"viewItemURL":["http:\/\/www.ebay.com\/itm\/3-D-Tic-Tac-Toe-Atari-2600-1980-CARTRIDGE-ONLY-CLEAN-TESTED-\/322119314985"],
	"productId":[{"@type":"ReferenceID","__value__":"56230072"}],
	"paymentMethod":["PayPal"],
	"autoPay":["false"],
	"postalCode":["55414"],
	"location":["Minneapolis,MN,USA"],
	"country":["US"],
	"shippingInfo":[
		{"shippingServiceCost":[{"@currencyId":"USD","__value__":"0.0"}],
		"shippingType":["Free"],
		"shipToLocations":["US"],
		"expeditedShipping":["false"],
		"oneDayShippingAvailable":["false"],
		"handlingTime":["1"]}
		],
	"sellingStatus":[
		{"currentPrice":[{"@currencyId":"USD","__value__":"5.95"}],
		"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"5.95"}],
		"sellingState":["Active"],
		"timeLeft":["P7DT17H53M0S"]}
		],
	"listingInfo":[
		{"bestOfferEnabled":["false"],
		"buyItNowAvailable":["false"],
		"startTime":["2016-05-23T20:41:25.000Z"],
		"endTime":["2016-06-22T20:41:25.000Z"],
		"listingType":["FixedPrice"],
		"gift":["false"]}
		],
	"returnsAccepted":["false"],
	"condition":[{"conditionId":["6000"],"conditionDisplayName":["Acceptable"]}],
	"isMultiVariationListing":["false"],
	"topRatedListing":["false"]
	},
	{
	"itemId":["322119314985"],
	"title":["3-D Tic-Tac-Toe (Atari 2600, 1980) CARTRIDGE ONLY! CLEAN & TESTED! "],
	"globalId":["EBAY-US"],
	"primaryCategory":[{"categoryId":["139973"],"categoryName":["Video Games"]}],
	"galleryURL":["http:\/\/thumbs2.ebaystatic.com\/m\/msVJvxYHgLg8C29GpD6-ZfQ\/140.jpg"],
	"viewItemURL":["http:\/\/www.ebay.com\/itm\/3-D-Tic-Tac-Toe-Atari-2600-1980-CARTRIDGE-ONLY-CLEAN-TESTED-\/322119314985"],
	"productId":[{"@type":"ReferenceID","__value__":"56230072"}],
	"paymentMethod":["PayPal"],
	"autoPay":["false"],
	"postalCode":["55414"],
	"location":["Minneapolis,MN,USA"],
	"country":["US"],
	"shippingInfo":[
		{"shippingServiceCost":[{"@currencyId":"USD","__value__":"0.0"}],
		"shippingType":["Free"],
		"shipToLocations":["US"],
		"expeditedShipping":["false"],
		"oneDayShippingAvailable":["false"],
		"handlingTime":["1"]}
		],
	"sellingStatus":[
		{"currentPrice":[{"@currencyId":"USD","__value__":"5.95"}],
		"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"5.95"}],
		"sellingState":["Active"],
		"timeLeft":["P7DT17H53M0S"]}
		],
	"listingInfo":[
		{"bestOfferEnabled":["false"],
		"buyItNowAvailable":["false"],
		"startTime":["2016-05-23T20:41:25.000Z"],
		"endTime":["2016-06-22T20:41:25.000Z"],
		"listingType":["FixedPrice"],"gift":["false"]}
		],
	"returnsAccepted":["false"],
	"condition":[{"conditionId":["6000"],"conditionDisplayName":["Acceptable"]}],
	"isMultiVariationListing":["false"],
	"topRatedListing":["false"]
	},
{"itemId":["262453127453"],"title":["3-D Tic-Tac-Toe (Atari 2600, 1980)"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["139973"],"categoryName":["Video Games"]}],"galleryURL":["http:\/\/thumbs2.ebaystatic.com\/m\/mz8ERy2_0hUY_S9iCcLXnqg\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/3-D-Tic-Tac-Toe-Atari-2600-1980-\/262453127453"],"productId":[{"@type":"ReferenceID","__value__":"170274227"}],"paymentMethod":["PayPal"],"autoPay":["true"],"postalCode":["42066"],"location":["Mayfield,KY,USA"],"country":["US"],"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"0.0"}],"shippingType":["FlatDomesticCalculatedInternational"],"shipToLocations":["US","CA","GB","AU","AT","BE","FR","DE","IT","JP","ES","TW","NL","CN","HK","MX","DK","RO","SK","BG","CZ","FI","HU","LV","LT","MT","EE","GR","PT","CY","SI","SE","KR","ID","TH","IE","PL","RU","IL","NZ"],"expeditedShipping":["false"],"oneDayShippingAvailable":["false"],"handlingTime":["1"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"5.99"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"5.99"}],"sellingState":["Active"],"timeLeft":["P8DT22H31M10S"]}],"listingInfo":[{"bestOfferEnabled":["true"],"buyItNowAvailable":["false"],"startTime":["2016-05-25T01:19:35.000Z"],"endTime":["2016-06-24T01:19:35.000Z"],"listingType":["StoreInventory"],"gift":["false"]}],"returnsAccepted":["true"],"condition":[{"conditionId":["6000"],"conditionDisplayName":["Acceptable"]}],"isMultiVariationListing":["false"],"topRatedListing":["true"]},
{"itemId":["281574339391"],"title":["3D Tic-Tac-Toe for Atari 2600 - Cartridge Only"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["139973"],"categoryName":["Video Games"]}],"galleryURL":["http:\/\/thumbs4.ebaystatic.com\/m\/mL4auqg2YnuxiaA-KpLSupw\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/3D-Tic-Tac-Toe-Atari-2600-Cartridge-Only-\/281574339391"],"productId":[{"@type":"ReferenceID","__value__":"170274227"}],"paymentMethod":["PayPal"],"autoPay":["true"],"postalCode":["91942"],"location":["La Mesa,CA,USA"],"country":["US"],"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"0.0"}],"shippingType":["FlatDomesticCalculatedInternational"],"shipToLocations":["US","CA","GB","AU","AT","BE","FR","DE","IT","JP","ES","TW","NL","CN","HK","MX","DK","RO","SK","BG","CZ","FI","HU","LV","LT","MT","EE","GR","PT","CY","SI","SE","KR","ID","TH","IE","PL","RU","IL","NZ"],"expeditedShipping":["true"],"oneDayShippingAvailable":["false"],"handlingTime":["1"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"4.45"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"4.45"}],"sellingState":["Active"],"timeLeft":["P1DT21H2M34S"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2015-01-23T23:45:59.000Z"],"endTime":["2016-06-16T23:50:59.000Z"],"listingType":["StoreInventory"],"gift":["false"]}],"returnsAccepted":["true"],"condition":[{"conditionId":["5000"],"conditionDisplayName":["Good"]}],"isMultiVariationListing":["false"],"topRatedListing":["true"]},
{"itemId":["182165392581"],"title":["3-D Tic-Tac-Toe, Loose Cartridge, Atari 2600"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["139973"],"categoryName":["Video Games"]}],"galleryURL":["http:\/\/thumbs2.ebaystatic.com\/m\/mgS5lkEX4caMAGFoE5Pcruw\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/3-D-Tic-Tac-Toe-Loose-Cartridge-Atari-2600-\/182165392581"],"productId":[{"@type":"ReferenceID","__value__":"170274227"}],"paymentMethod":["PayPal"],"autoPay":["false"],"postalCode":["07647"],"location":["Northvale,NJ,USA"],"country":["US"],
	"shippingInfo":[{"shippingType":["Calculated"],"shipToLocations":["Worldwide"],"expeditedShipping":["false"],"oneDayShippingAvailable":["false"],"handlingTime":["1"]}],
	"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"0.22"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"0.22"}],"bidCount":["2"],"sellingState":["Active"],"timeLeft":["P1DT19H21M30S"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2016-06-09T22:09:55.000Z"],"endTime":["2016-06-16T22:09:55.000Z"],"listingType":["Auction"],"gift":["false"]}],"returnsAccepted":["true"],"condition":[{"conditionId":["4000"],"conditionDisplayName":["Very Good"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]},
{"itemId":["391473063984"],"title":["Atari 2600 game 3D Tic Tac Toe  With Instructions Tested and Working"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["139973"],"categoryName":["Video Games"]}],"galleryURL":["http:\/\/thumbs1.ebaystatic.com\/m\/mwVb_2ap65J7PGmbrrWJ08A\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/Atari-2600-game-3D-Tic-Tac-Toe-Instructions-Tested-and-Working-\/391473063984"],"paymentMethod":["PayPal"],"autoPay":["false"],"postalCode":["61201"],"location":["Rock Island,IL,USA"],"country":["US"],
	"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"3.95"}],"shippingType":["FlatDomesticCalculatedInternational"],"shipToLocations":["Worldwide"],"expeditedShipping":["false"],"oneDayShippingAvailable":["false"],"handlingTime":["2"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"5.75"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"5.75"}],"sellingState":["Active"],"timeLeft":["P17DT16H48M40S"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2016-06-02T19:37:05.000Z"],"endTime":["2016-07-02T19:37:05.000Z"],"listingType":["StoreInventory"],"gift":["false"]}],"returnsAccepted":["true"],"condition":[{"conditionId":["5000"],"conditionDisplayName":["Good"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]},
{"itemId":["322143019049"],"title":["3-D Tic-Tac-Toe (Atari 2600, 1980)"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["139973"],"categoryName":["Video Games"]}],"galleryURL":["http:\/\/thumbs2.ebaystatic.com\/m\/m7uq9T-082DAnYLu6XbGGAQ\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/3-D-Tic-Tac-Toe-Atari-2600-1980-\/322143019049"],"productId":[{"@type":"ReferenceID","__value__":"56230072"}],"paymentMethod":["PayPal"],"autoPay":["false"],"postalCode":["92223"],"location":["Beaumont,CA,USA"],"country":["US"],
	"shippingInfo":[{
		"shippingType":["Calculated"],
		"shipToLocations":["US","CA","GB","AU","AT","BE","FR","DE","IT","JP","ES","TW","NL","CN","HK","MX","DK","RO","SK","BG","CZ","FI","HU","LV","LT","MT","EE","GR","PT","CY","SI","SE","KR","ID","TH","IE","PL","RU","IL","NZ"],
		"expeditedShipping":["false"],
		"oneDayShippingAvailable":["false"],
		"handlingTime":["2"]}],
	"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"4.95"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"4.95"}],"bidCount":["0"],"sellingState":["Active"],"timeLeft":["P3DT21H16M26S"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2016-06-09T00:04:51.000Z"],"endTime":["2016-06-19T00:04:51.000Z"],"listingType":["Auction"],"gift":["false"]}],"returnsAccepted":["false"],"condition":[{"conditionId":["5000"],"conditionDisplayName":["Good"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]},
{"itemId":["201574375625"],"title":["3-D Tic-Tac-Toe (Atari 2600, 1980) Cartridge Only"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["139973"],"categoryName":["Video Games"]}],"galleryURL":["http:\/\/thumbs2.ebaystatic.com\/m\/mhzAzqpU7W8utgStBQPDtBQ\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/3-D-Tic-Tac-Toe-Atari-2600-1980-Cartridge-Only-\/201574375625"],"productId":[{"@type":"ReferenceID","__value__":"56230072"}],"paymentMethod":["PayPal"],"autoPay":["true"],"postalCode":["41566"],"location":["Steele,KY,USA"],"country":["US"],
	"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"0.0"}],"shippingType":["FlatDomesticCalculatedInternational"],"shipToLocations":["US","CA","GB","AU","AT","BE","FR","DE","IT","JP","ES","TW","NL","CN","HK","MX","DK","RO","SK","BG","CZ","FI","HU","LV","LT","MT","EE","GR","PT","CY","SI","SE","KR","ID","TH","IE","PL","RU","IL","NZ"],"expeditedShipping":["false"],"oneDayShippingAvailable":["false"],"handlingTime":["2"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"4.99"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"4.99"}],"sellingState":["Active"],"timeLeft":["P15DT17H34M21S"]}],"listingInfo":[{"bestOfferEnabled":["true"],"buyItNowAvailable":["false"],"startTime":["2016-05-01T20:17:46.000Z"],"endTime":["2016-06-30T20:22:46.000Z"],"listingType":["StoreInventory"],"gift":["false"]}],"returnsAccepted":["false"],"condition":[{"conditionId":["5000"],"conditionDisplayName":["Good"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]}]}
],
"paginationOutput":[{"pageNumber":["1"],"entriesPerPage":["100"],"totalPages":["1"],"totalEntries":["22"]
}
],
"itemSearchURL":["http:\/\/www.ebay.com\/sch\/i.html?_nkw=3d+toe+atari+2600&_ddo=1&_ipg=100&_pgn=1"]
}
]
}`
