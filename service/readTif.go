package service

import (
	"GdalProject/entity"
	"GdalProject/global"
	"fmt"
	"github.com/lukeroth/gdal"
)

func ReadTif(name string) {
	prefix := global.FileSetting.Location
	postfix := ".tif"
	filename := prefix + name + postfix
	//打卡文件
	dataset, _ := gdal.Open(filename, gdal.ReadOnly)
	//获取图像尺寸
	nXsize := dataset.RasterXSize()
	nYsize := dataset.RasterYSize()

	testGeoTransform := dataset.GeoTransform()
	//trans[6]  数组adfGeoTransform保存的是仿射变换中的一些参数，分别含义见下
	/*
		trans[0]  左上角x坐标
		trans[1]  东西方向分辨率
		trans[2]  旋转角度, 0表示图像 "北方朝上"
		trans[3]  左上角y坐标
		trans[4]  旋转角度, 0表示图像 "北方朝上"
		trans[5]  南北方向分辨率
	*/
	fmt.Println("开始读取新地图数据，起始点经纬度信息如下：")
	fmt.Printf("%+v\n", testGeoTransform[0])
	fmt.Printf("%+v\n", testGeoTransform[3])
	//申请缓存
	//buff := make([]float64,nYsize*nXsize)
	buffer := make([]float64, 1)
	//计算高程数据
	raster := dataset.RasterBand(1)
	////读取高程数据
	//raster.IO(gdal.Read,0,0,nXsize,nYsize,buff,nXsize,nYsize,0,0)
	high := float64(0)
	points := []*entity.Earth_tif{}
	for i := 0; i < nXsize; i++ {
		for j := 0; j < nYsize; j++ {
			//计算经纬度
			lon := testGeoTransform[0] + float64(i)*testGeoTransform[1] + float64(j)*testGeoTransform[2]
			lat := testGeoTransform[3] + float64(i)*testGeoTransform[4] + float64(j)*testGeoTransform[5]
			raster.IO(gdal.Read, i, j, 1, 1, buffer, 1, 1, 0, 0)
			alt := buffer[0]
			if high < buffer[0] {
				high = buffer[0]
			}
			point := entity.Earth_tif{}
			point.Lon = lon
			point.Lat = lat
			point.Alt = alt
			points = append(points, &point)
		}
	}
	BatchInsert(points, 10000)
	fmt.Println("highest alt is :", high)
	fmt.Printf("End program\n")
}
