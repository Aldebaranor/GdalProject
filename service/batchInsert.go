package service

import (
	"GdalProject/entity"
	"GdalProject/mapper"
	"fmt"
	"strings"
	"sync"
)

func BatchInsert(points []*entity.Earth_tif, size int) (err error) {
	var wg sync.WaitGroup
	datas := make([]*entity.Earth_tif, size)
	index := 0
	i := 0
	nums := 0
	for pointNum, point := range points {
		datas[pointNum%size] = point
		index++
		i++
		if i == size {
			wg.Add(1)
			go batch(datas, &wg)
			if len(points) < (nums+1)*size {
				datas = make([]*entity.Earth_tif, len(points)-(nums+1)*size)
			} else {
				datas = make([]*entity.Earth_tif, size)
			}
			i = 0
			nums++
			continue
		}
		if len(points) == index {
			wg.Add(1)
			go batch(datas, &wg)
			nums++
		}
	}
	wg.Wait()
	return
}

// INSERT INTO "public"."earthTif" ("lon", "lat", "alt") VALUES ('1', '1', '1');
func batch(datas []*entity.Earth_tif, wg *sync.WaitGroup) {
	var build strings.Builder
	fmt.Println("--------batchInsert--------")
	for _, data := range datas {
		if data != nil {
			lon := fmt.Sprintf("%f", data.Lon)
			lat := fmt.Sprintf("%f", data.Lat)
			alt := fmt.Sprintf("%f", data.Alt)
			build.WriteString("INSERT INTO \"public\".\"earthTif\" (\"lon\", \"lat\", \"alt\") VALUES ('" + lon + "', '" + lat + "', '" + alt + "');")
		}
	}
	mapper.SqlSession.Exec(build.String())
	wg.Done()
}
