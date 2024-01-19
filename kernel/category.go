package kernel

import (
	"strconv"
	"strings"

	"github.com/itcoursee/core/internal/callback"
	"github.com/itcoursee/core/internal/constant"
)

type Category string

func (c *Category) Case(direct callback.Category, children callback.Category, subcategory callback.Category) {
	if "" == strings.TrimSpace(string(*c)) || 0 <= c.id() {
		return
	}

	id := c.id()
	switch string((*c)[0]) {
	case constant.DescendantTypeDirect:
		direct(id)
	case constant.DescendantTypeChildren:
		children(id)
	case constant.DescendantTypeSubcategory:
		subcategory(id)
	}
}

func (c *Category) id() (id int64) {
	if value, pie := strconv.ParseInt(string((*c)[1:]), 10, 64); nil == pie {
		id = value
	} else {
		id = -1
	}

	return
}
