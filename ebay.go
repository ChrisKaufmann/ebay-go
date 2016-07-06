package ebay

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type EBay struct {
	ApplicationId string
	URL           string
	EntriesLength int
}

type Item struct {
	ID            string
	Location      string
	Url           string
	ImageUrl      string
	Title         string
	Price         float64
	BuyItNowPrice float64
	BuyItNow      bool
	ShippingPrice float64
	FreeShipping  bool
	ShippingType  string
	StartTime     string
	EndTime       string
	BestOffer     bool
}

func New(application_id string) *EBay {
	e := EBay{}
	e.ApplicationId = application_id
	e.URL = "http://svcs.ebay.com/services/search/FindingService/v1"
	e.EntriesLength = 100
	return &e
}
func (e *EBay) Search(searchstring string) (il []Item, err error) {
	searchstring = strings.Replace(searchstring, " ", "%20", -1)
	x, err := e.GetResponse(searchstring)
	if err != nil {
		glog.Errorf("e.GetResponse(%s): %s", searchstring, err)
		return il, err
	}
	il, err = e.ParseJSON(x)
	if err != nil {
		glog.Errorf("ebay.GetResponse(%s): %s", searchstring, err)
	}
	return il, err
}
func (e *EBay) ParseJSON(x string) (il []Item, err error) {
	type KWS struct {
		Resp []struct {
			Ack          []string `json:"ack"`
			SearchResult []struct {
				Count string `json:"@count"`
				Item  []struct {
					ItemID   []string `json:"itemId"`
					Title    []string `json:"title"`
					Url      []string `json:"viewItemURL"`
					ImageURL []string `json:"galleryURL"`
					Location []string `json:"location"`
					Shipping []struct {
						ShippingType []string `json:"shippingType"`
						ShippingCost []struct {
							Value string `json:"__value__"`
						} `json:"shippingServiceCost"`
					} `json:"shippingInfo"`
					Price []struct {
						CurrentPrice []struct {
							Value string `json:"__value__"`
						} `json:"currentPrice"`
					} `json:"sellingStatus"`
					ListingInfo []struct {
						BuyNowEnabled []string `json:"buyItNowAvailable"`
						StartTime     []string `json:"startTime"`
						EndTime       []string `json:"endTime"`
						BestOffer     []string `json:"bestOfferEnabled"`
					} `json:"listingInfo"`
				} `json:"item"`
			} `json:"searchResult"`
		} `json:"findItemsByKeywordsResponse"`
	}

	var g KWS
	err = json.Unmarshal([]byte(x), &g)
	if err != nil {
		glog.Errorf("json.Unmarshal(see below, &f): %s\n%s", err, x)
		return il, err
	}
	if g.Resp[0].Ack[0] != "Success" {
		return il, err
	}
	rc, err := strconv.Atoi(g.Resp[0].SearchResult[0].Count)
	if err != nil || rc < 1 {
		return il, err
	}
	for _, i := range g.Resp[0].SearchResult[0].Item {
		var ni Item
		ni.Price, err = strconv.ParseFloat(i.Price[0].CurrentPrice[0].Value, 64)
		if err != nil {
			glog.Errorf("parsefloat(%s,64): %s", i.Price[0].CurrentPrice[0].Value, err)
			return il, err
		}
		if len(i.Shipping[0].ShippingCost) > 0 {
			ni.ShippingPrice, err = strconv.ParseFloat(i.Shipping[0].ShippingCost[0].Value, 64)
			if err != nil {
				glog.Errorf("Parsefloat(%s, 64): %s", i.Shipping[0].ShippingCost[0].Value, err)
			}
			ni.FreeShipping = false
		}
		if i.Shipping[0].ShippingType[0] == "Free" {
			ni.FreeShipping = true
		}
		ni.ShippingType = i.Shipping[0].ShippingType[0]
		if i.ListingInfo[0].BuyNowEnabled[0] == "true" {
			ni.BuyItNow = true
		} else {
			ni.BuyItNow = false
		}
		if i.ListingInfo[0].BestOffer[0] == "true" {
			ni.BestOffer = true
		} else {
			ni.BestOffer = false
		}
		ni.ID = i.ItemID[0]
		ni.Title = i.Title[0]
		ni.Url = i.Url[0]
		ni.Location = i.Location[0]
		ni.StartTime = i.ListingInfo[0].StartTime[0]
		ni.EndTime = i.ListingInfo[0].EndTime[0]
		if len(i.ImageURL) > 0 {
			ni.ImageUrl = i.ImageURL[0]
		}
		il = append(il, ni)
	}
	return il, err
}
func (e *EBay) GetResponse(searchstring string) (x string, err error) {
	url := fmt.Sprintf("%s?OPERATION-NAME=findItemsByKeywords&SERVICE-VERSION=1.0.0&SECURITY-APPNAME=%s&RESPONSE-DATA-FORMAT=JSON&REST-PAYLOAD&keywords=%s&GLOBAL-ID=EBAY-US&paginationInput.entriesPerPage=%v", e.URL, e.ApplicationId, searchstring, e.EntriesLength)
	resp, err := http.Get(url)
	if err != nil {
		glog.Errorf("http.Get(%s): %s", url, err)
		return x, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorf("ioutil.ReadAll(%s): %s", resp.Body, err)
		return x, err
	}
	x = fmt.Sprintf("%s", body)
	return x, err
}

func (i Item) String() string {
	return fmt.Sprintf("Url: %s\nTitle: %s\nPrice: %v\n", i.Url, i.Title, i.Price)
}

func LowestPrice(il []Item) Item {
	sort.Sort(ByPrice(il))
	return (il[0])
}
func LowestPricePlusShipping(il []Item) Item {
	sort.Sort(ByPricePlusShipping(il))
	if len(il) > 0 {
		return (il[0])
	}
	var i Item
	return i
}
func EndingSoonest(il []Item) Item {
	sort.Sort(EndingSooner(il))
	for _, i := range il {
		fmt.Printf("id: %s, price: %v, ship: %v, shiptype: %s\n", i.ID, i.Price, i.ShippingPrice, i.ShippingType)
	}
	return (il[0])
}

type ByPrice []Item

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }

type ByPricePlusShipping []Item

func (a ByPricePlusShipping) Len() int      { return len(a) }
func (a ByPricePlusShipping) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPricePlusShipping) Less(i, j int) bool {
	if a[i].ShippingType == "Calculated" {
		return false
	}
	return (a[i].Price + a[i].ShippingPrice) < (a[j].Price + a[j].ShippingPrice)
}

type EndingSooner []Item

func (a EndingSooner) Len() int           { return len(a) }
func (a EndingSooner) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a EndingSooner) Less(i, j int) bool { return a[i].EndTime < a[j].EndTime }
