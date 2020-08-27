package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/zeroclock/taskiwi/config"
	"github.com/zeroclock/taskiwi/model"
	"github.com/zeroclock/taskiwi/utils"
	"github.com/zeroclock/taskiwi/validation"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo) {
	e.GET("/", indexHandler)
	e.POST("/allTags", allTagsHandler)
	e.POST("/aggregateTasks", aggregateTasks)
	e.POST("/taskByDate", taskByDateHandler)
}

func indexHandler(c echo.Context) error {
	f, err := Assets.Open("/web/taskiwi/build/index.html")
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

func allTagsHandler(c echo.Context) (err error) {
	searchCondition := new(validation.TagByDateCondition)
	if err = c.Bind(searchCondition); err != nil {
		return
	}
	if err = c.Validate(searchCondition); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	clockDatas, err := QueryTasks(*config.GlobalConf.CData, []string{}, searchCondition.Start, searchCondition.End)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	tags := GetTagsFromClockDatas(clockDatas)

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
	taskDatas, err := QueryTasks(*config.GlobalConf.CData, []string{}, searchCondition.Date, searchCondition.Date)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal server error")
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

	clockDatas, err := QueryTasks(*config.GlobalConf.CData, []string{}, searchCondition.Start, searchCondition.End)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	workTimes, taskDatas, err := AggregateClockDatasByTags(clockDatas, searchCondition.Tags)

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
