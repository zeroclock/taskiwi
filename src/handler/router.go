package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"taskiwi/config"
	"taskiwi/model"
	"taskiwi/utils"
	"taskiwi/validation"

	"github.com/labstack/echo"
)

const layout = "2006-01-02 15:04"

// received date's format
const layout_req = "2006-01-02"

func InitRouting(e *echo.Echo) {
	e.GET("/", indexHandler)
	e.GET("/all", allTaskHandler)
	e.GET("/allTags", allTagsHandler)
	e.POST("/aggregateTasks", aggregateTasks)
}

type Employee struct {
	Name string
}

func indexHandler(c echo.Context) error {
	f, err := os.Open("./web/taskiwi/build/index.html")
	if err != nil {
		log.Println("[WARNING] Failed to load index.html")
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("[WARNING] Failed to read from file buffer")
	}

	return c.HTML(http.StatusOK, string(b))
}

func allTaskHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, config.GlobalConf.CData)
}

func allTagsHandler(c echo.Context) error {
	var tags []string
	for _, v := range *config.GlobalConf.CData {
		tags = append(tags, v.Tags...)
	}

	return c.JSON(http.StatusOK, utils.Unique(tags))
}

func aggregateTasks(c echo.Context) (err error) {
	searchCondition := new(validation.AggregateCondition)
	if err = c.Bind(searchCondition); err != nil {
		return
	}
	if err = c.Validate(searchCondition); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var workTimes model.WorkTimes
	m := map[string]float64{}
	var total float64 = 0.0

	for _, tag := range searchCondition.Tags {
		for _, v := range *config.GlobalConf.CData {
			conditionStart, err := time.Parse(layout_req, searchCondition.Start)
			if err != nil {
				log.Println(err)
			}
			conditionEnd, err := time.Parse(layout_req, searchCondition.End)
			if err != nil {
				log.Println(err)
			}
			if hasTag(&v, tag) && checkRange(&v, conditionStart, conditionEnd) {
				start, err := time.Parse(layout, v.Start)
				if err != nil {
					log.Println(err)
				}
				end, err := time.Parse(layout, v.End)
				if err != nil {
					log.Println(err)
				}
				sub := end.Sub(start).Minutes()
				m[tag] += sub
				total += sub
			} else {
				m[tag] += 0
			}
		}
	}

	for k, v := range m {
		workTimes = append(workTimes, &model.WorkTime{
			Tag:     k,
			Time:    fmt.Sprint(v),
			Percent: fmt.Sprint(v / total),
		})
	}

	return c.JSON(http.StatusOK, workTimes)
}

func hasTag(task *model.ClockData, tag string) bool {
	for _, v := range task.Tags {
		if v == tag {
			return true
		}
	}
	return false
}

func checkRange(task *model.ClockData, start time.Time, end time.Time) bool {
	clockedTimeStart, err := time.Parse(layout, task.Start)
	hour := time.Duration(clockedTimeStart.Hour())
	minute := time.Duration(clockedTimeStart.Minute())
	second := time.Duration(clockedTimeStart.Second())
	clockedTimeStart = clockedTimeStart.Add(-hour * time.Hour)
	clockedTimeStart = clockedTimeStart.Add(-minute * time.Minute)
	clockedTimeStart = clockedTimeStart.Add(-second * time.Second)
	if err != nil {
		return false
	}
	// it is expected that start and end are the same date
	return !start.After(clockedTimeStart) && !end.Before(clockedTimeStart)
}
