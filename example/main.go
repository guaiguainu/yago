package main // import "github.com/hulklab/yago/example"

import (
	"github.com/hulklab/yago"

	_ "github.com/hulklab/yago/example/app/routes/cmdroute"
	_ "github.com/hulklab/yago/example/app/routes/httproute"
	_ "github.com/hulklab/yago/example/app/routes/rpcroute"
	_ "github.com/hulklab/yago/example/app/routes/taskroute"
)

func main() {
	yago.NewCmd().RunCmd()
}
