package userfunction

import (
	ofctx "github.com/OpenFunction/functions-framework-go/openfunction-context"
	"log"
)

func BindingsOutput(ctx *ofctx.OpenFunctionContext, in []byte) ofctx.RetValue {
	var greeting []byte
	if in != nil {
		log.Printf("binding - Data: %s", in)
		greeting = in
	} else {
		log.Print("binding - Data: Received")
		greeting = []byte("Hello")
	}

	err := ctx.SendTo(greeting, "echo")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return ctx.ReturnWithInternalError()
	}
	return ctx.ReturnWithSuccess()
}
