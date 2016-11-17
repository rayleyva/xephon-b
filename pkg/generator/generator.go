package generator

import (
	"errors"

	"github.com/xephonhq/xephon-b/pkg/common"
)

// ErrEndOfPoints is end of points
var ErrEndOfPoints = errors.New("EOP")

// IntPointGenerator generate integer value
// TODO: may change to some interface that support built-in range operator, or may just add a Channel?
// TODO: may support channel
type IntPointGenerator interface {
	Next() (p common.IntPoint, err error)
	IsValid() bool
}

// DoublePointGenerator generate double value
type DoublePointGenerator interface {
	Next() (p common.DoublePoint, err error)
	IsValid() bool
}