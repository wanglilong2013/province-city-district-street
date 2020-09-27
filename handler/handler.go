package handler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"province-city-district-street/conf"
	"province-city-district-street/fetcher"
	"province-city-district-street/model"
	"province-city-district-street/parser"
	"strconv"
	"time"
)

func Handle(resp parser.RespData) error {
	if len(resp.Districts) == 0 {
		fmt.Println("Districts node data is empty")
		return nil
	}
	user := conf.MysqlDbUser
	host := conf.MysqlDbHost
	port := conf.MysqlDbPort
	dbMame := conf.MysqlDbName
	pwd := conf.MysqlDbPwd
	conStr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, port, dbMame)
	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for _, item := range resp.Districts {
		fmt.Println(item.Name)
		for _, district := range item.Districts {
			//省数据保存
			save(db, district, 0)

			keywords := district.AdCode
			provinceId, _ := strconv.Atoi(keywords)
			cityUrl := conf.URL
			cityUrl = fmt.Sprintf(cityUrl, keywords, 3, conf.Key)

			cityData, err := fetcher.Fetch(cityUrl)
			if err != nil {
				panic(err)
			}
			cityResp, err := parser.Parse(cityData)
			for _, cityItem := range cityResp.Districts {
				for _, cityDistrict := range cityItem.Districts {
					//市数据保存
					save(db, cityDistrict, provinceId)
					cityId, _ := strconv.Atoi(cityDistrict.AdCode)

					////区数据保存
					dDistricts := cityDistrict.Districts
					for _, dDistrict := range dDistricts {
						save(db, dDistrict, cityId)
						districtId, _ := strconv.Atoi(dDistrict.AdCode)
						//街道数据保存
						streetDisctricts := dDistrict.Districts
						for _, streetDistrict := range streetDisctricts {
							save(db, streetDistrict, districtId)
						}
					}

				}
			}
		}
	}
	return nil
}

func save(db *gorm.DB, district parser.District, pid int) {
	id, _ := strconv.Atoi(district.AdCode)
	area := &model.Area{
		id,
		district.Name,
		pid,
		district.Level,
		time.Now(),
		time.Now(),
	}
	db.Create(area)
}
