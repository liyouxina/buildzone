package enviroment

import "github.com/liyouxina/buildzone/server/interfaces"

var GlobalEnviromentContext = Enviroment{}

type Enviroment struct {
	cache interfaces.Cache
}
