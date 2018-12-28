package main

type rPlugin int

type CallMessage struct {
	Func    string
	Cookies string
}

func (rp rPlugin) RunPlugin(jArgs string) string {
	return "hello" + jArgs
}

// exported as symbol named "Greeter"
var ARPlugin rPlugin
