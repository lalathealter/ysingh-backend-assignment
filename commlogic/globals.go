package commlogic

type WarehousesColl map[int]int

var GetWarehousesColl = func() func() WarehousesColl {
	wares := WarehousesColl{}
	return func() WarehousesColl {
		return wares
	}
}()

type GlobalCatalog map[string]string

var GetCatalog = func() func() GlobalCatalog {
	cat := GlobalCatalog{}
	return func() GlobalCatalog {
		return cat
	}
}()
