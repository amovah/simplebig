package simplebig

import (
	"fmt"
	"math"
	"math/big"
	"simplebig/simpleint"
	"strings"
)

func SimpleIntToFloatString(simpleIntValue simpleint.Int, precision int) string {
	decimals := simpleint.New(10).Pow(simpleint.New(int64(precision)).Mul(simpleint.New(-1)))
	mod := simpleIntValue.Mod(decimals)

	modFloat := new(big.Float)
	modFloat.SetPrec(236)
	modFloat.SetMode(big.ToNearestEven)
	modFloat.SetInt(mod.BigInt())
	modFloat.Mul(modFloat, big.NewFloat(math.Pow(10, float64(precision)*-1)))
	modPartArr := strings.Split(modFloat.String(), ".")

	var modPart string
	if len(modPartArr) < 2 {
		modPart = "0"
	} else {
		modPart = modPartArr[1]
	}

	integer := simpleIntValue.Div(decimals)

	return fmt.Sprintf("%s.%s", integer.String(), modPart)
}
