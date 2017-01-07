package main

import (
	"net/http"

	"fmt"

	"math"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Static("/css", "css")

	r.GET("/data", calculate)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

type Input struct {
	Salary int64 `form:"salary"`
	Goal   int64 `form:"goal"`
}

type ResultRow struct {
	Percent      int64
	ResultIncome int64
	Time         string
}

func prepareTime(months int64) string {
	if months < 12 {
		return fmt.Sprintf("%v months", months)
	} else if months < 24 {
		return fmt.Sprintf("%v year, %v months", months/12, months%12)
	} else {
		return fmt.Sprintf("%v years, %v months", months/12, months%12)
	}
}

const (
	maxPercent = 80
	minPercent = 50
	step       = 5
)

func calculate(c *gin.Context) {
	f := new(Input)
	err := c.Bind(f)
	if err != nil {
		return
	}
	if f.Salary <= 0 || f.Goal <= 0 {
		return
	}
	response := []ResultRow{}
	for base := int64(maxPercent); base >= minPercent; base -= step {
		income := int64(float64(f.Salary) * float64(base) / 100.0)

		months := int64(math.Ceil(float64(f.Goal) / float64(income)))

		response = append(response, ResultRow{
			Percent:      base,
			ResultIncome: income,
			Time:         prepareTime(months),
		})
	}
	c.HTML(http.StatusOK, "index.html", response)
}
