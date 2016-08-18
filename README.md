ebay-go
=======
Another golang ebay API, mostly for searching.

Installation
------------
`go get github.com/ChrisKaufmann/ebay-go`

Usage
-----
```
eb := ebay.New("my application ID")
items, err := eb.Search("Something")
if err != nil {return err}
```
Convenience:
```
cheapest := LowestPrice(itemlist)       // returns the item with the lowest price
c2 := LowestPricePlusShipping(itemlist) //returns cheapest with shipping
e := EndingSoonest(itemlist)            //ending soonest
```

Example
-------
<pre>
package main
import (
	"github.com/ChrisKaufmann/ebay-go"
	"fmt"
)
const ebay_application_id = "My_Application_ID"
var (
	eb *ebay.EBay
)


func init() {
	eb = ebay.New(ebay_application_id)
}

func main() {
    search_string := "Atari%202600%20combat"
    // Optionally, set a category id
    // eb.CategoryID=1122334455
    itemlist, err  := eb.Search(search_string)
    if err != nil {
        fmt.Printf("eb.Search(%s): %s", search_string, err)
        return
    }
    for _, i := range itemlist {
        fmt.Printf("ID: %s\n", i.ID)                        //"322119314985"
        fmt.Printf("Location: %s\n", i.Location)            //"Minneapolis,MN,USA"
        fmt.Printf("Url: %s\n", i.Url)                      //"http://www.ebay.com/itm/3-D-Tic-Tac-Toe-Atari-2600-1980-CARTRIDGE-ONLY-CLEAN-TESTED-/322119314985"
        fmt.Printf("ImageUrl: %s\n", i.ImageUrl)            //"http://thumbs2.ebaystatic.com/m/msVJvxYHgLg8C29GpD6-ZfQ/140.jpg"
        fmt.Printf("Title: %s\n", i.Title)                  //"3-D Tic-Tac-Toe (Atari 2600, 1980) CARTRIDGE ONLY! CLEAN & TESTED! "
        fmt.Printf("Price: %v\n", i.Price)                  //5.95
        fmt.Printf("BuyItNowPrice: %v\n", i.BuyItNowPrice)  //0.00
        fmt.Printf("BuyItNow: %v\n", i.BuyItNow)            //false
        fmt.Printf("ShippingPrice: %v\n", i.ShippingPrice)  //0.0
        fmt.Printf("FreeShipping: %v\n", i.FreeShipping)    //true
        fmt.Printf("StartTime: %s\n", i.StartTime)          //"2016-05-23T20:41:25.000Z"
        fmt.Printf("EndTime: %s\n", i.EndTime)              //"2016-06-22T20:41:25.000Z"
        fmt.Printf("BestOffer: %s\n", i.BestOffer)          //false
    }
    ci := ebay.LowestPrice(itemlist)
    fmt.Printf("Cheapest ID and price: %s, %s\n", ci.ID, ci.Price)
    
    cips := ebay.LowestPricePlusShipping(itemlist)
    fmt.Printf("Cheapest price+shipping: id: %s, price: %v, shipping: %v, shippingtype: %s\n", cips.ID, cips.Price, cips.ShippingPrice, cips.ShippingType)
    
    es := ebay.EndingSoonest(itemlist)
    fmt.Printf("Ending soonest: ID: %s, Endtime: %s\n", es.ID, es.EndTime)
    
    //find some completed auctions
    completedlist, err := eb.SearchCompleted("antique rocking horse")
    if err != nil {
        fmt.Printf("eb.SearchCompleted(...): %s", err)
        return
    }
    // All the sorting, etc work the same on completed as on regular auctions
    ci = ebay.LowestPricePlusShipping(completedlist)
}
</pre>