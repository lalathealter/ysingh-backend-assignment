package commlogic

type Warehouse struct {
	Limit   int
	Storage map[string]int
}

func (w Warehouse) Store(SKU string, QTY int) int {
	oldAmount := w.Storage[SKU]
	w.Storage[SKU] += QTY
	if w.Limit != 0 && w.Storage[SKU] > w.Limit {
		w.Storage[SKU] = w.Limit
	} else if w.Storage[SKU] < 0 {
		w.Storage[SKU] = 0
	}

	diff := oldAmount - w.Storage[SKU]
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func ProduceWarehouse(cap int) Warehouse {
	if cap < 0 {
		cap = 0
	}
	return Warehouse{cap, map[string]int{}}
}

type WarehousesColl map[int]Warehouse

var GetWarehousesColl = func() func() WarehousesColl {
	wares := WarehousesColl{}
	return func() WarehousesColl {
		return wares
	}
}()

type Catalog map[string]string

var GetCatalog = func() func() Catalog {
	cat := Catalog{}
	return func() Catalog {
		return cat
	}
}()
