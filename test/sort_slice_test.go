package test

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

type banner struct {
	DirectUrl string
	ImageUrl  string
	Sort      int
}

func TestSortSlice(t *testing.T) {
	banner1 := banner{
		DirectUrl: "123",
		ImageUrl:  "456",
		Sort:      1,
	}

	banner2 := banner{
		DirectUrl: "1237",
		ImageUrl:  "4569",
		Sort:      2,
	}

	banner3 := banner{
		DirectUrl: "1235",
		ImageUrl:  "4560",
		Sort:      3,
	}

	var banners []banner
	banners = append(banners, banner1, banner2, banner3)

	sort.Slice(banners, func(i, j int) bool {
		return banners[i].Sort > banners[j].Sort
	})

	fmt.Printf("sort banners: %+v\n", banners)

	now := time.Now().UnixNano() / 1e6
	fmt.Printf("now: +++ %v\n", now)

	timestamp := fmt.Sprintf("%d", time.Now().Unix()*1000)
	fmt.Printf("timestamp: +++++ %s\n", timestamp)

}
