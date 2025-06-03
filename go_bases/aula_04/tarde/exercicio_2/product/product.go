package product

type ProductType int

const (
	Small = iota
	Medium
	Large
)

type Item struct {
	productType  ProductType
	value        float64
	tax          int64
	stockTax     int64
	shipmentCost float64
}

func NewItem(pType ProductType, value float64) Item {
	var item Item
	switch pType {
	case Small:
		item = Item{
			productType: pType,
			value:       value,
		}
	case Medium:
		item = Item{
			productType: pType,
			value:       value,
			tax:         3,
			stockTax:    3,
		}
	case Large:
		item = Item{
			productType:  pType,
			value:        value,
			stockTax:     6,
			shipmentCost: 2500,
		}
	}
	return item
}

func (i Item) Price() float64 {
	return i.value + (i.value * float64(i.tax) / 100) + (i.value * float64(i.stockTax) / 100) + i.shipmentCost
}

type Product interface {
	Price() float64
}
