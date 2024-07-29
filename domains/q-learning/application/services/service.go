package services

import "math/rand"

func GenerateSignalMap(length int) (mapping map[int64]int64, buyArr, sellArr []int64) {
	mapping = make(map[int64]int64)
	for i := 0; i < length; i++ {
		key := int64(rand.Float64() * 1000000)
		mapping[key] = 1
		buyArr = append(buyArr, key)
	}
	for i := 0; i < length; i++ {
		key := int64(rand.Float64() * 1000000)
		mapping[key] = -1
		sellArr = append(sellArr, key)

	}
	return
}
