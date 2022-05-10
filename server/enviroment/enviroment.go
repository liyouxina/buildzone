package enviroment

import (
	"github.com/liyouxina/buildzone/server/interfaces"
)

type Enviroment struct {
	cache interfaces.Cache
}

func (e *Enviroment) GetCache() interfaces.Cache {
	return e.cache
}

func (e *Enviroment) SetCache(cache interfaces.Cache) {
	e.cache = cache
}
