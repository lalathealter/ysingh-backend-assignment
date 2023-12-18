package commlogic

type Warehouse struct {
	ItemsCount int
	Limit      int
	Storage    map[string]int
}

func (w *Warehouse) Store(SKU string, QTY int) int {
	oldAmount := w.Storage[SKU]
	w.Storage[SKU] += QTY
	if w.Limit != 0 && w.ItemsCount+QTY > w.Limit {
		w.Storage[SKU] -= ((w.ItemsCount + QTY) - w.Limit)
	} else if w.Storage[SKU] < 0 {
		delete(w.Storage, SKU)
	}

	diff := w.Storage[SKU] - oldAmount
	w.ItemsCount += diff
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func ProduceWarehouse(cap int) *Warehouse {
	if cap < 0 {
		cap = 0
	}
	return &Warehouse{0, cap, map[string]int{}}
}

type WarehousesColl map[int]*Warehouse

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
