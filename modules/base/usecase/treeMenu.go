package usecase

import (
	"sort"

	"github.com/donnyirianto/be-stock-app/modules/base/domain"
)

func (uc *BaseUsecaseImpl) TreeBase(baseItems []*domain.BaseItem) ([]*domain.NewBaseItem, error) {
	// Map item ID ke NewBaseItem untuk pencarian yang lebih cepat
	itemMap := make(map[int16]*domain.NewBaseItem)

	// Membuat item dasar dan memetakan ke itemMap
	for _, item := range baseItems {
		itemMap[item.ID] = &domain.NewBaseItem{
			ID:       item.ID,
			IDMain:   item.IDMain,
			Urut:     item.Urut,
			Label:    item.Nama,
			Href:     item.Link,
			Children: []*domain.NewBaseItem{},
		}
	}

	// Menambahkan children ke parent yang sesuai
	var roots []*domain.NewBaseItem
	for _, item := range baseItems {
		if item.IDMain == 0 {
			// Ini adalah root item
			roots = append(roots, itemMap[item.ID])
		} else {
			// Ini adalah child, menambahkannya ke parent yang sesuai
			parent := itemMap[item.IDMain]
			parent.Children = append(parent.Children, itemMap[item.ID])
		}
	}

	// Menambahkan subitems secara rekursif ke dalam children
	for _, root := range roots {
		uc.addChildrenRecursive(root, itemMap)
	}

	// Sorting setelah rekursi selesai
	uc.sortTree(roots)

	// Mengembalikan roots yang sudah terstruktur dan di-sort
	return roots, nil
}

// Fungsi rekursif untuk menambahkan children
func (uc *BaseUsecaseImpl) addChildrenRecursive(parent *domain.NewBaseItem, itemMap map[int16]*domain.NewBaseItem) {
	// Mencari children untuk parent tertentu
	for _, child := range parent.Children {
		// Memanggil rekursi untuk children
		child.Children = itemMap[child.ID].Children
		uc.addChildrenRecursive(child, itemMap)
	}
}

// Fungsi untuk sorting tree berdasarkan Urut
func (uc *BaseUsecaseImpl) sortTree(items []*domain.NewBaseItem) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Urut < items[j].Urut
	})

	// Sorting children secara rekursif
	for _, item := range items {
		if len(item.Children) > 0 {
			uc.sortTree(item.Children)
		}
	}
}
