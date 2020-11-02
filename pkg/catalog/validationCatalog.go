package catalog

import (
	"sync"

	"gitlab.com/rbell/gospecexpress/pkg/internal/catalog"

	"gitlab.com/rbell/gospecexpress/pkg/interfaces"
)

var instance interfaces.Cataloger
var instanceOnce = &sync.Once{}

// ValidationCatalog gets the singleton instance of the Cataloger
func ValidationCatalog() interfaces.Cataloger {
	instanceOnce.Do(func() {
		instance = catalog.NewDefaultCatalog()
	})

	return instance
}
