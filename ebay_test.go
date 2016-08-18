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
	assert.Equal(t, i.ShippingType, "Free", "ShippingType")
	assert.True(t, i.FreeShipping, "Free shipping")
	//3rd one has bestoffer true
	assert.True(t, il[3].BestOffer, "Best Offer")

	//  the 6th one has non-free shipping :)
	assert.Equal(t, 3.95, il[6].ShippingPrice, "ShippingPrice")
	assert.False(t, il[6].FreeShipping, "Free Shipping")
	assert.Equal(t, il[6].ShippingType, "FlatDomesticCalculatedInternational", "ShippingType")

	il, err = e.ParseJSON(fail_json)
	assert.Nil(t, err, "GetItemsFromJSON")
	assert.Equal(t, 0, len(il))

	il, err = e.ParseJSON(empty_json)
	assert.Nil(t, err, "GetItemsFromJSON")
	assert.Equal(t, 0, len(il))
}
func TestEBay_ParseCompletedJSON(t *testing.T) {
	print("TestParseCompletedJSON\n")
	eb := New("12345")
	il, err := eb.ParseCompletedJSON(test_completed_json)
	assert.Nil(t, err, "ParseCompletdJSON()")
	assert.Equal(t, 13, len(il))
}
func TestLowestPrice(t *testing.T) {
	print("TestLowestPrice\n")
	e := New("12345")
	il, err := e.ParseJSON(test_json)
	assert.Nil(t, err, "e.ParseJson(test_json)")
	lp := LowestPrice(il)
	assert.Equal(t, 0.22, lp.Price, "lp.Price")
}
func TestEndingSoonest(t *testing.T) {
	print("TestEndingSoonest\n")
	e := New("12345")
	il, err := e.ParseJSON(test_json)
	assert.Nil(t, err, "e.ParseJson(test_json)")
	i := EndingSoonest(il)
	assert.Equal(t, "2016-06-16T22:09:55.000Z", i.EndTime, "EndingSoonest")
}
func TestLowestPricePlusShipping(t *testing.T) {
	print("LowestPricePlusShipping\n")
	e := New("12345")
	il, err := e.ParseJSON(test_json)
	assert.Nil(t, err, "e.ParseJson(test_json)")
	i := LowestPricePlusShipping(il)
	assert.Equal(t, 4.45, i.ShippingPrice+i.Price, "Price+shipping")
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
		"listingType":["FixedPrice"],
		"gift":["false"]}
		],
	"returnsAccepted":["false"],
	"condition":[{"conditionId":["6000"],"conditionDisplayName":["Acceptable"]}],
	"isMultiVariationListing":["false"],
	"topRatedListing":["false"]
	},
{
	"itemId":["262453127453"],
	"title":["3-D Tic-Tac-Toe (Atari 2600, 1980)"],
	"globalId":["EBAY-US"],
	"primaryCategory":[
		{"categoryId":["139973"],
		"categoryName":["Video Games"]}
		],
	"galleryURL":["http:\/\/thumbs2.ebaystatic.com\/m\/mz8ERy2_0hUY_S9iCcLXnqg\/140.jpg"],
	"viewItemURL":["http:\/\/www.ebay.com\/itm\/3-D-Tic-Tac-Toe-Atari-2600-1980-\/262453127453"],
	"productId":[{"@type":"ReferenceID","__value__":"170274227"}],
	"paymentMethod":["PayPal"],
	"autoPay":["true"],
	"postalCode":["42066"],
	"location":["Mayfield,KY,USA"],
	"country":["US"],
	"shippingInfo":[
		{"shippingServiceCost":[{"@currencyId":"USD","__value__":"0.0"}],
		"shippingType":["FlatDomesticCalculatedInternational"],
		"shipToLocations":["US","CA","GB","AU","AT","BE","FR","DE","IT","JP","ES","TW","NL","CN","HK","MX","DK","RO","SK","BG","CZ","FI","HU","LV","LT","MT","EE","GR","PT","CY","SI","SE","KR","ID","TH","IE","PL","RU","IL","NZ"],
		"expeditedShipping":["false"],
		"oneDayShippingAvailable":["false"],
		"handlingTime":["1"]}
		],
	"sellingStatus":
		[{"currentPrice":[{"@currencyId":"USD","__value__":"5.99"}],
		"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"5.99"}],
		"sellingState":["Active"],
		"timeLeft":["P8DT22H31M10S"]}
		],
	"listingInfo":[
		{"bestOfferEnabled":["true"],
		"buyItNowAvailable":["false"],
		"startTime":["2016-05-25T01:19:35.000Z"],
		"endTime":["2016-06-24T01:19:35.000Z"],
		"listingType":["StoreInventory"],
		"gift":["false"]}
		],
	"returnsAccepted":["true"],
	"condition":[{"conditionId":["6000"],"conditionDisplayName":["Acceptable"]}],
	"isMultiVariationListing":["false"],
	"topRatedListing":["true"]
	},
{
	"itemId":["281574339391"],
	"title":["3D Tic-Tac-Toe for Atari 2600 - Cartridge Only"],
	"globalId":["EBAY-US"],
	"primaryCategory":[{"categoryId":["139973"],"categoryName":["Video Games"]}],
	"galleryURL":["http:\/\/thumbs4.ebaystatic.com\/m\/mL4auqg2YnuxiaA-KpLSupw\/140.jpg"],
	"viewItemURL":["http:\/\/www.ebay.com\/itm\/3D-Tic-Tac-Toe-Atari-2600-Cartridge-Only-\/281574339391"],
	"productId":[{"@type":"ReferenceID","__value__":"170274227"}],
	"paymentMethod":["PayPal"],
	"autoPay":["true"],
	"postalCode":["91942"],
	"location":["La Mesa,CA,USA"],
	"country":["US"],
	"shippingInfo":[
		{"shippingServiceCost":[{"@currencyId":"USD","__value__":"0.0"}],
		"shippingType":["FlatDomesticCalculatedInternational"],
		"shipToLocations":["US","CA","GB","AU","AT","BE","FR","DE","IT","JP","ES","TW","NL","CN","HK","MX","DK","RO","SK","BG","CZ","FI","HU","LV","LT","MT","EE","GR","PT","CY","SI","SE","KR","ID","TH","IE","PL","RU","IL","NZ"],
		"expeditedShipping":["true"],
		"oneDayShippingAvailable":["false"],
		"handlingTime":["1"]}
		],
	"sellingStatus":[
		{"currentPrice":[{"@currencyId":"USD","__value__":"4.45"}],
		"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"4.45"}],
		"sellingState":["Active"],
		"timeLeft":["P1DT21H2M34S"]}
		],
	"listingInfo":[
		{"bestOfferEnabled":["false"],
		"buyItNowAvailable":["false"],
		"startTime":["2015-01-23T23:45:59.000Z"],
		"endTime":["2016-06-16T23:50:59.000Z"],
		"listingType":["StoreInventory"],
		"gift":["false"]}
		],
	"returnsAccepted":["true"],
		"condition":[{"conditionId":["5000"],
		"conditionDisplayName":["Good"]}
		],
	"isMultiVariationListing":["false"],
	"topRatedListing":["true"]
},
{
	"itemId":["182165392581"],"title":["3-D Tic-Tac-Toe, Loose Cartridge, Atari 2600"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["139973"],"categoryName":["Video Games"]}],"galleryURL":["http:\/\/thumbs2.ebaystatic.com\/m\/mgS5lkEX4caMAGFoE5Pcruw\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/3-D-Tic-Tac-Toe-Loose-Cartridge-Atari-2600-\/182165392581"],"productId":[{"@type":"ReferenceID","__value__":"170274227"}],"paymentMethod":["PayPal"],"autoPay":["false"],"postalCode":["07647"],"location":["Northvale,NJ,USA"],"country":["US"],
	"shippingInfo":[
		{"shippingType":["Calculated"],
		"shipToLocations":["Worldwide"],
		"expeditedShipping":["false"],
		"oneDayShippingAvailable":["false"],
		"handlingTime":["1"]}
		],
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

const test_completed_json = `
{"findCompletedItemsResponse":[
	{"ack":["Success"],
	"version":["1.13.0"],
	"timestamp":["2016-08-16T21:17:32.658Z"],
	"searchResult":[
		{"@count":"100","item":[
			{"itemId":["201633751337"],
				"title":["My Little Pony Applejack "],
				"globalId":["EBAY-US"],
				"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],
				"galleryURL":["http:\/\/thumbs2.ebaystatic.com\/m\/m_V8gHg7O74GMeigUTNCw0A\/140.jpg"],
				"viewItemURL":["http:\/\/www.ebay.com\/itm\/My-Little-Pony-Applejack-\/201633751337"],
				"paymentMethod":["PayPal"],
				"autoPay":["false"],
				"postalCode":["85710"],
				"location":["Tucson,AZ,USA"],
				"country":["US"],
				"shippingInfo":[{"shippingServiceCost":
					[{"@currencyId":"USD","__value__":"3.99"}],"shippingType":["Flat"],"shipToLocations":["US"],
					"expeditedShipping":["false"],"oneDayShippingAvailable":["false"],"handlingTime":["3"]}],
				"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"2.5"}],
					"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"2.5"}],
					"sellingState":["EndedWithSales"]}],
				"listingInfo":[
					{"bestOfferEnabled":["true"],"buyItNowAvailable":["false"],
					"startTime":["2016-07-27T23:25:50.000Z"],
					"endTime":["2016-08-16T20:41:21.000Z"],
					"listingType":["FixedPrice"],"gift":["false"]}],
				"returnsAccepted":["true"],
				"condition":[{"conditionId":["3000"],"conditionDisplayName":["Used"]}],
				"isMultiVariationListing":["false"],
				"topRatedListing":["false"]},
			{"itemId":["282120880472"],"title":["MY LITTLE PONY APPLEJACK Build A Bear Plush Stuffed Animal Toy Yellow Horse 15\""],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],"galleryURL":["http:\/\/thumbs1.ebaystatic.com\/m\/mcvxNxPiRQwKinc1V9U8PlA\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/MY-LITTLE-PONY-APPLEJACK-Build-Bear-Plush-Stuffed-Animal-Toy-Yellow-Horse-15-\/282120880472"],"paymentMethod":["PayPal"],"autoPay":["false"],"postalCode":["16438"],"location":["Union City,PA,USA"],"country":["US"],"shippingInfo":[{"shippingType":["Calculated"],"shipToLocations":["US","CA","GB","AU","AT","BE","FR","DE","IT","JP","ES","TW","NL","CN","HK","MX","DK","RO","SK","BG","CZ","FI","HU","LV","LT","MT","EE","GR","PT","CY","SI","SE","KR","ID","TH","IE","PL","RU","IL","NZ"],"expeditedShipping":["true"],"oneDayShippingAvailable":["false"],"handlingTime":["1"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"19.0"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"19.0"}],"sellingState":["EndedWithSales"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2016-08-02T15:02:31.000Z"],"endTime":["2016-08-16T19:02:36.000Z"],"listingType":["StoreInventory"],"gift":["false"]}],"returnsAccepted":["true"],"condition":[{"conditionId":["3000"],"conditionDisplayName":["Used"]}],"isMultiVariationListing":["false"],"topRatedListing":["true"]},
			{"itemId":["131902982734"],"title":["VAULTED ! Applejack my little pony From Funko ! A Collector !"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],"galleryURL":["http:\/\/thumbs3.ebaystatic.com\/m\/mVT_kbsvCNT93rorKUVJ03w\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/VAULTED-Applejack-my-little-pony-Funko-Collector-\/131902982734"],"paymentMethod":["PayPal"],"autoPay":["false"],"location":["USA"],"country":["US"],"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"7.0"}],"shippingType":["Flat"],"shipToLocations":["US"],"expeditedShipping":["true"],"oneDayShippingAvailable":["false"],"handlingTime":["2"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"25.0"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"25.0"}],"sellingState":["EndedWithSales"]}],"listingInfo":[{"bestOfferEnabled":["true"],"buyItNowAvailable":["false"],"startTime":["2016-08-08T20:30:55.000Z"],"endTime":["2016-08-16T06:06:35.000Z"],"listingType":["FixedPrice"],"gift":["false"]}],"returnsAccepted":["false"],"condition":[{"conditionId":["1000"],"conditionDisplayName":["New"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]},
			{"itemId":["231308645380"],"title":["New My Little Pony MLP plush Doll Stuffed Toy Figure Applejack Orange 7\" 18 cm "],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],"galleryURL":["http:\/\/thumbs1.ebaystatic.com\/m\/m06Ab1K5595ppqzAxMSwdlg\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/New-My-Little-Pony-MLP-plush-Doll-Stuffed-Toy-Figure-Applejack-Orange-7-18-cm-\/231308645380"],"paymentMethod":["PayPal"],"autoPay":["true"],"location":["Hong Kong"],"country":["HK"],"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"0.0"}],"shippingType":["Free"],"shipToLocations":["Worldwide"],"expeditedShipping":["false"],"oneDayShippingAvailable":["false"],"handlingTime":["1"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"7.99"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"7.99"}],"sellingState":["EndedWithSales"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2014-08-15T17:29:20.000Z"],"endTime":["2016-08-16T03:46:06.000Z"],"listingType":["FixedPrice"],"gift":["false"]}],"returnsAccepted":["true"],"condition":[{"conditionId":["1000"],"conditionDisplayName":["New"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]},
			{"itemId":["311673900981"],"title":["My Little Pony Applejack Plush Doll 12'' POPL9040"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],"galleryURL":["http:\/\/thumbs2.ebaystatic.com\/m\/mjaeozJic52RCOZHXWiBo_w\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/My-Little-Pony-Applejack-Plush-Doll-12-POPL9040-\/311673900981"],"paymentMethod":["PayPal"],"autoPay":["false"],"location":["China"],"country":["CN"],"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"9.9"}],"shippingType":["Flat"],"shipToLocations":["Worldwide"],"expeditedShipping":["false"],"oneDayShippingAvailable":["false"],"handlingTime":["5"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"9.99"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"9.99"}],"bidCount":["1"],"sellingState":["EndedWithSales"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2016-08-11T02:05:53.000Z"],"endTime":["2016-08-16T02:05:53.000Z"],"listingType":["Auction"],"gift":["false"]}],"returnsAccepted":["true"],"condition":[{"conditionId":["1000"],"conditionDisplayName":["New"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]},
			{"itemId":["381649132124"],"title":["G2-8 My Little Pony Friendship Explore Equestria Applejack & Fluttershy Lot of 3"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],"galleryURL":["http:\/\/thumbs1.ebaystatic.com\/m\/mqr9YUW1MDIFjambpspnCdg\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/G2-8-My-Little-Pony-Friendship-Explore-Equestria-Applejack-Fluttershy-Lot-3-\/381649132124"],"paymentMethod":["PayPal"],"autoPay":["false"],"postalCode":["45663"],"location":["West Portsmouth,OH,USA"],"country":["US"],"shippingInfo":[{"shippingType":["Calculated"],"shipToLocations":["Worldwide"],"expeditedShipping":["true"],"oneDayShippingAvailable":["false"],"handlingTime":["1"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"7.49"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"7.49"}],"sellingState":["EndedWithSales"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2016-05-26T01:12:36.000Z"],"endTime":["2016-08-16T00:29:37.000Z"],"listingType":["StoreInventory"],"gift":["false"]}],"returnsAccepted":["true"],"condition":[{"conditionId":["1000"],"conditionDisplayName":["New"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]},
			{"itemId":["142076134817"],"title":["Applejack 2.25\" button My Little Pony Equestria Girls Rainbow Rocks"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],"galleryURL":["http:\/\/thumbs2.ebaystatic.com\/m\/mFG9_3xdLD6u48k2NgZVszA\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/Applejack-2-25-button-My-Little-Pony-Equestria-Girls-Rainbow-Rocks-\/142076134817"],"paymentMethod":["PayPal"],"autoPay":["false"],"postalCode":["19311"],"location":["Avondale,PA,USA"],"country":["US"],"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"3.08"}],"shippingType":["Flat"],"shipToLocations":["Worldwide"],"expeditedShipping":["true"],"oneDayShippingAvailable":["false"],"handlingTime":["3"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"2.0"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"2.0"}],"sellingState":["EndedWithSales"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2016-08-05T05:23:35.000Z"],"endTime":["2016-08-15T22:03:55.000Z"],"listingType":["StoreInventory"],"gift":["false"]}],"returnsAccepted":["false"],"condition":[{"conditionId":["1000"],"conditionDisplayName":["New"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]},
			{"itemId":["252501345702"],"title":["Applejack CUSTOM Fantasy Armor My Little Pony Game of Thrones Lord of the Rings"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],"galleryURL":["http:\/\/thumbs3.ebaystatic.com\/m\/ma7b3uWspK2nkbWhm_ACy9w\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/Applejack-CUSTOM-Fantasy-Armor-My-Little-Pony-Game-Thrones-Lord-Rings-\/252501345702"],"paymentMethod":["PayPal"],"autoPay":["false"],"postalCode":["04276"],"location":["Rumford,ME,USA"],"country":["US"],"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"0.0"}],"shippingType":["Free"],"shipToLocations":["US"],"expeditedShipping":["false"],"oneDayShippingAvailable":["false"],"handlingTime":["1"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"15.0"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"15.0"}],"sellingState":["EndedWithSales"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2016-08-15T17:36:27.000Z"],"endTime":["2016-08-15T21:52:31.000Z"],"listingType":["FixedPrice"],"gift":["false"]}],"returnsAccepted":["false"],"condition":[{"conditionId":["1000"],"conditionDisplayName":["New"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]},
			{"itemId":["162165020242"],"title":["My Little Pony APPLEJACK Plush Doll 12inches"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],"galleryURL":["http:\/\/thumbs3.ebaystatic.com\/m\/mMdthrGD5w2s_TZed0F4OZw\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/My-Little-Pony-APPLEJACK-Plush-Doll-12inches-\/162165020242"],"paymentMethod":["PayPal"],"autoPay":["false"],"location":["China"],"country":["CN"],"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"12.8"}],"shippingType":["Flat"],"shipToLocations":["Worldwide"],"expeditedShipping":["false"],"oneDayShippingAvailable":["false"],"handlingTime":["1"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"0.8"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"0.8"}],"bidCount":["1"],"sellingState":["EndedWithSales"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2016-08-10T17:23:53.000Z"],"endTime":["2016-08-15T17:23:53.000Z"],"listingType":["Auction"],"gift":["false"]}],"returnsAccepted":["true"],"condition":[{"conditionId":["1000"],"conditionDisplayName":["New"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]},
			{"itemId":["142081364412"],"title":["4\" TY Beanie Babies Sparkle AppleJack Pendant My Little Pony Plush Stuffed Toys"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],"galleryURL":["http:\/\/thumbs1.ebaystatic.com\/m\/m3WLoCTqSk-TYiDitp4q2rQ\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/4-TY-Beanie-Babies-Sparkle-AppleJack-Pendant-My-Little-Pony-Plush-Stuffed-Toys-\/142081364412"],"paymentMethod":["PayPal"],"autoPay":["false"],"postalCode":["510700"],"location":["China"],"country":["CN"],"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"0.0"}],"shippingType":["Free"],"shipToLocations":["Worldwide"],"expeditedShipping":["false"],"oneDayShippingAvailable":["false"],"handlingTime":["1"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"3.25"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"3.25"}],"bidCount":["9"],"sellingState":["EndedWithSales"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2016-08-10T17:08:27.000Z"],"endTime":["2016-08-15T17:08:27.000Z"],"listingType":["Auction"],"gift":["false"]}],"returnsAccepted":["true"],"condition":[{"conditionId":["1000"],"conditionDisplayName":["New"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]},
			{"itemId":["182235456000"],"title":["MY LITTLE PONY, YELLOW AND BLOND, APPLEJACK FRIENDSHIP FIGURINE, HASBRO 2010"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],"galleryURL":["http:\/\/thumbs1.ebaystatic.com\/m\/msn3sREqAUFC5gRXxvHgPRA\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/MY-LITTLE-PONY-YELLOW-AND-BLOND-APPLEJACK-FRIENDSHIP-FIGURINE-HASBRO-2010-\/182235456000"],"paymentMethod":["PayPal"],"autoPay":["false"],"postalCode":["01532"],"location":["Northborough,MA,USA"],"country":["US"],"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"2.65"}],"shippingType":["Flat"],"shipToLocations":["US"],"expeditedShipping":["false"],"oneDayShippingAvailable":["false"],"handlingTime":["4"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"0.5"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"0.5"}],"bidCount":["1"],"sellingState":["EndedWithSales"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2016-08-08T10:22:06.000Z"],"endTime":["2016-08-15T10:22:06.000Z"],"listingType":["Auction"],"gift":["false"]}],"returnsAccepted":["false"],"condition":[{"conditionId":["3000"],"conditionDisplayName":["Used"]}],"isMultiVariationListing":["false"],"topRatedListing":["false"]},
			{"itemId":["131902272271"],"title":["Funko My Little Pony Vinyl Figure: APPLEJACK - New"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],"galleryURL":["http:\/\/thumbs4.ebaystatic.com\/m\/mPErmF94Va7zb1MnpC4ozpQ\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/Funko-My-Little-Pony-Vinyl-Figure-APPLEJACK-New-\/131902272271"],"paymentMethod":["PayPal"],"autoPay":["false"],"postalCode":["24521"],"location":["Amherst,VA,USA"],"country":["US"],"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"7.99"}],"shippingType":["FlatDomesticCalculatedInternational"],"shipToLocations":["Worldwide"],"expeditedShipping":["true"],"oneDayShippingAvailable":["false"],"handlingTime":["1"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"22.5"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"22.5"}],"bidCount":["23"],"sellingState":["EndedWithSales"]}],"listingInfo":[{"bestOfferEnabled":["false"],"buyItNowAvailable":["false"],"startTime":["2016-08-08T04:48:53.000Z"],"endTime":["2016-08-15T04:48:53.000Z"],"listingType":["Auction"],"gift":["false"]}],"returnsAccepted":["true"],"condition":[{"conditionId":["1000"],"conditionDisplayName":["New"]}],"isMultiVariationListing":["false"],"topRatedListing":["true"]},
			{"itemId":["222206930786"],"title":["My Little Pony Lot of 4 Rainbowfied Twilight Applejack Pinky Pie Fluttershy"],"globalId":["EBAY-US"],"primaryCategory":[{"categoryId":["47228"],"categoryName":["1990-Now"]}],"galleryURL":["http:\/\/thumbs3.ebaystatic.com\/m\/myKse5v7j5dUUmsdbmYK-ZQ\/140.jpg"],"viewItemURL":["http:\/\/www.ebay.com\/itm\/My-Little-Pony-Lot-4-Rainbowfied-Twilight-Applejack-Pinky-Pie-Fluttershy-\/222206930786"],"paymentMethod":["PayPal"],"autoPay":["false"],"postalCode":["21012"],"location":["Arnold,MD,USA"],"country":["US"],"shippingInfo":[{"shippingServiceCost":[{"@currencyId":"USD","__value__":"3.0"}],"shippingType":["Flat"],"shipToLocations":["US"],"expeditedShipping":["false"],"oneDayShippingAvailable":["false"],"handlingTime":["1"]}],"sellingStatus":[{"currentPrice":[{"@currencyId":"USD","__value__":"4.0"}],"convertedCurrentPrice":[{"@currencyId":"USD","__value__":"4.0"}],"sellingState":["EndedWithSales"]}],"listingInfo":[{"bestOfferEnabled":["true"],"buyItNowAvailable":["false"],"startTime":["2016-08-01T21:04:40.000Z"],"endTime":["2016-08-05T03:22:08.000Z"],"listingType":["FixedPrice"],"gift":["false"]}],"returnsAccepted":["true"],"condition":[{"conditionId":["3000"],"conditionDisplayName":["Used"]}],"isMultiVariationListing":["false"],"topRatedListing":["true"]}
			]
		}
	],
	"paginationOutput":[{"pageNumber":["1"],"entriesPerPage":["100"],"totalPages":["6"],"totalEntries":["557"]}]}]}`
