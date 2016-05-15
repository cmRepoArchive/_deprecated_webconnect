package xmessage

import (
	"reflect"
	"strings"
)

type Processor struct {
	Module  string
	PkgPath string
	Name    string
	Func    func([]reflect.Value) []reflect.Value
}

func registerProcessor(p *Processor) {
	Table[p.Module+"."+p.Name] = p
}

/*
	Client use a part of the table's key (usually is ProcessorName) to match a processor.
	Currently, the key is "PkgName.ProcessorName" so team members should use different package names to do a replacement though their package path is not the same.
*/
func matchProcessor(suffix string) {
	matchedList := []string{}
	// key is "PkgName.ProcessorName"
	for key, _ := range Table {
		if strings.HasSuffix(key, suffix) {
			matchedList = append(matchedList, key)
		}
	}
}