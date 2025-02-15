package userfunction

import (
	ofctx "github.com/OpenFunction/functions-framework-go/openfunction-context"
	"log"
)

func BindingsNoOutput(ctx *ofctx.OpenFunctionContext, in []byte) ofctx.RetValue {
	if in != nil {
		log.Printf("binding - Data: %s", in)
	} else {
		log.Print("binding - Data: Received")
	}
	return ctx.ReturnWithSuccess()
}
