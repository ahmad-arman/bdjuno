package source

import (
	stakeibctypes "github.com/Stride-Labs/stride/v12/x/stakeibc/types"
)

type Source interface {
	Params(height int64) (stakeibctypes.Params, error)
}
