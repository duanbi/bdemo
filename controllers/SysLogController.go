package controllers

import (
	"bdemo/models"
	"strconv"
)

type SysLogController struct {
	SysBaseController
}

func (c *SysLogController) Prepare() {
	c.SysBaseController.Prepare()
}

func (c *SysLogController) GetSysLogListByPage() {
	where := make(map[string]string)

	where["user_name"] = c.Input().Get("user_name")
	where["start_time"] = c.Input().Get("start_time")
	where["end_time"] = c.Input().Get("end_time")

	pageNum, _ := strconv.Atoi(c.Input().Get("page_num"))
	if pageNum <= 0 {
		pageNum = 1
	}
	roleList, count := models.GetSysLogListByPage(where, pageNum, 11, "log_id desc")
	c.Data["LogList"] = roleList
	c.Data["LogCount"] = count
	c.Data["PageNum"] = pageNum
	c.Data["where"] = where
	c.TplName = "syslog/listSysLog.html"
}

func (c *SysLogController) DeleteSysLog() {
	ids := c.Input().Get("log_ids")
	_, err := models.DeleteSysLog(ids)
	if err != nil {
		c.DisplayJson(0, "修改失败", err.Error())
	}
	c.DisplayJson(1, "删除成功", c.URLFor("SysLogController.GetSysLogListByPage"))
}
