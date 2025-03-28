package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	file, err := os.Create("lotto_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()

	headers := []string{"期号", "星期", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "和值", "跨度", "区间比", "奇偶比"}
	err = writer.Write(headers)
	if err != nil {
		log.Fatal(err)
	}

	c.OnHTML("table tbody tr", func(e *colly.HTMLElement) {
		columns := make([]string, 0)
		// columns = append(columns, e.ChildText(fmt.Sprintf("td:nth-child(%d)", 1)))
		for i := 1; i < 62; i++ {
			if InArray([]int{2, 4, 17, 30, 42, 60}, i) {
				continue
			}
			class := e.ChildAttr(fmt.Sprintf("td:nth-child(%d)", i), "class")
			if strings.Index(class, "chartball") >= 0 {
				columns = append(columns, e.ChildText(fmt.Sprintf("td:nth-child(%d)", i)))
			} else {
				if i > 2 {
					columns = append(columns, "-")
				} else {
					columns = append(columns, e.ChildText(fmt.Sprintf("td:nth-child(%d)", i)))
				}

			}
		}

		err := writer.Write(columns)
		if err != nil {
			log.Fatal(err)
		}
	})

	err = c.Visit("https://lotto.sina.cn/trend/qxc_qlc_proxy.d.html?lottoType=dlt&actionType=chzs&0_ala_h5baidu&_headline=baidu_ala&type=300")
	if err != nil {
		log.Fatal(err)
	}
}

func InArray(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
