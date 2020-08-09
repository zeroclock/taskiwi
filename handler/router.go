package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/zeroclock/taskiwi/config"
	"github.com/zeroclock/taskiwi/model"
	"github.com/zeroclock/taskiwi/utils"
	"github.com/zeroclock/taskiwi/validation"

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
	e.POST("/taskByDate", taskByDateHandler)
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

func taskByDateHandler(c echo.Context) (err error) {
	searchCondition := new(validation.TaskByDateCondition)
	if err = c.Bind(searchCondition); err != nil {
		return
	}
	if err = c.Validate(searchCondition); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	conditionDate, err := time.Parse(layout_req, searchCondition.Date)
	if err != nil {
		log.Println(err)
	}
	taskDatas := []model.ClockData{}
	for _, v := range *config.GlobalConf.CData {
		clockedDate, err := time.Parse(layout, v.Start)
		if err != nil {
			log.Println(err)
		}
		clockedDate = utils.TruncateDateHMS(clockedDate)
		if conditionDate.Equal(clockedDate) {
			taskDatas = append(taskDatas, v)
		}
	}
	sort.SliceStable(taskDatas, func(i, j int) bool {
		startA, _ := time.Parse(layout, taskDatas[i].Start)
		startB, _ := time.Parse(layout, taskDatas[j].Start)
		return startA.Before(startB)
	})
	return c.JSON(http.StatusOK, taskDatas)
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
	var taskDatas []model.ClockData
	m := map[string]float64{}
	var total float64 = 0.0

	conditionStart, err := time.Parse(layout_req, searchCondition.Start)
	if err != nil {
		log.Println(err)
	}
	conditionEnd, err := time.Parse(layout_req, searchCondition.End)
	if err != nil {
		log.Println(err)
	}

	for i, tag := range searchCondition.Tags {
		for _, v := range *config.GlobalConf.CData {
			if checkRange(&v, conditionStart, conditionEnd) {
				if hasTag(&v, tag) {
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
				if i == 0 {
					taskDatas = append(taskDatas, v)
				}
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

	return c.JSON(http.StatusOK, &model.Response{
		WorkTimes:  workTimes,
		ClockDatas: taskDatas,
	})
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
	clockedTimeStart = utils.TruncateDateHMS(clockedTimeStart)
	if err != nil {
		return false
	}
	// it is expected that start and end are the same date
	return !start.After(clockedTimeStart) && !end.Before(clockedTimeStart)
}
