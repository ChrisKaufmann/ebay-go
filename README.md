Another golang ebay API, mostly for searching.

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
    search_string = "Atari%202600%20combat"
    itemlist, err  := eb.Search(search_string)
    if err != nil {
        fmt.Printf("eb.Search(%s): %s", search_string, err)
        return
    }
    for _, i := range itemlist {
        fmt.Printf("ID: %s\n", i.ID)
    }
    ci := eb.LowestPrice(itemlist)
    fmt.Printf("Cheapest ID and price: %s, %s\n", ci.ID, ci.Price)
    
}
</pre>