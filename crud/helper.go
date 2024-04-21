package crud

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VarFactory struct {
	Value string
}

func (inst *VarFactory) ToInt() int {
	val, _ := strconv.Atoi(inst.Value)
	return val
}

func (inst *VarFactory) ToString() string {
	return fmt.Sprintf("%v", inst.Value)
}

func (inst *VarFactory) ToFloat() float64 {
	val, _ := strconv.ParseFloat(inst.Value, 64)
	return val
}

func BuidFactory(c *gin.Context) map[string]*VarFactory {
	queryParams := c.Request.URL.Query()

	// Extract values from queryParams map
	result := make(map[string]*VarFactory)
	for k, v := range queryParams {
		result[k] = &VarFactory{
			Value: v[0],
		}
	}
	return result

}
