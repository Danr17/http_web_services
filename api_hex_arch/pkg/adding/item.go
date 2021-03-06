package adding

import "time"

/*
Good struct holds the data of an item type:
ID -- autogenarated
Name -- Goods name
Manufactured -- Goods manufactured date
ExpDate -- Goods validity
ExpOpen -- Goods validity if it is opened
*/
type Good struct {
	Name         string    `json:"name"`
	Manufactured time.Time `json:"manufactured"`
	ExpDate      time.Time `jsnon:"expdate"`
	ExpOpen      int       `json:"expopen"`
}

/*
Item struct holds the data of the instance of the goods:
ID -- autogenerated
Type -- the type of product
IsOpen -- True if the product is opened
Opened -- The date when it was opeened
IsValid -- Is the item still in validity or has expired
*/
type Item struct {
	ID      int `json:"id"`
	Created time.Time
	Good
	IsOpen  bool      `json:"isopen"`
	Opened  time.Time `json:"opened,omitempty"`
	IsValid bool      `json:"isvalid"`
}
