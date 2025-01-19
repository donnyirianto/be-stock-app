package utils

import (
	"fmt"

	"dario.cat/mergo"
)

func MergeByKdcab(base, additional []map[string]interface{}) []map[string]interface{} {
	for i, baseItem := range base {
		for _, addItem := range additional {
			if baseItem["kdcab"] == addItem["kdcab"] {
				if err := mergo.Merge(&base[i], addItem, mergo.WithOverride); err != nil {
					fmt.Println("Error merging maps:", err)
				}
			}
		}
	}
	return base
}

func MergeByToko(base, additional []map[string]interface{}) []map[string]interface{} {
	for i, baseItem := range base {
		for _, addItem := range additional {
			if baseItem["toko"] == addItem["toko"] {
				if err := mergo.Merge(&base[i], addItem, mergo.WithOverride); err != nil {
					fmt.Println("Error merging maps:", err)
				}
			}
		}
	}
	return base
}
func MergeByCatcod(base, additional []map[string]interface{}) []map[string]interface{} {
	for i, baseItem := range base {
		for _, addItem := range additional {
			if baseItem["cat_cod"] == addItem["cat_cod"] {
				if err := mergo.Merge(&base[i], addItem, mergo.WithOverride); err != nil {
					fmt.Println("Error merging maps:", err)
				}
			}
		}
	}
	return base
}

func MergeByPlu(base, additional []map[string]interface{}) []map[string]interface{} {
	for i, baseItem := range base {
		for _, addItem := range additional {
			if baseItem["prdcd"] == addItem["prdcd"] {
				if err := mergo.Merge(&base[i], addItem, mergo.WithOverride); err != nil {
					fmt.Println("Error merging maps:", err)
				}
			}
		}
	}
	return base
}

func MergeByDep(base, additional []map[string]interface{}) []map[string]interface{} {
	for i, baseItem := range base {
		for _, addItem := range additional {
			if baseItem["departement"] == addItem["departement"] {
				if err := mergo.Merge(&base[i], addItem, mergo.WithOverride); err != nil {
					fmt.Println("Error merging maps:", err)
				}
			}
		}
	}
	return base
}
