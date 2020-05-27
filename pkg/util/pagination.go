// 编写分页页码的获取方法

package util

import (
	"gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	if page, _ := com.StrTo(c.Query("page")).Int(); page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}


