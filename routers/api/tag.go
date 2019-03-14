package api

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"supremine/lib/e"
	"supremine/models"
	"supremine/lib/util"
	"supremine/lib/setting"
	"net/http"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	code := e.INVALID_PARAMS
	if ! models.ExistTagByName(name) {
		code = e.SUCCESS
		models.AddTag(name, state)
	} else {
		code = e.ERROR_EXIST_TAG
	}


	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

func EditTag(c *gin.Context) {
}

func DeleteTag(c *gin.Context) {
}
