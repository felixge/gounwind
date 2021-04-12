package gounwind

import (
	"debug/dwarf"
	"debug/macho"
	"fmt"
	"os"
	"unsafe"
)

var dwarfData *dwarf.Data

func init() {
	file, err := macho.Open(os.Args[0])
	if err != nil {
		panic(err)
	}
	dwarfData, err = file.DWARF()
	if err != nil {
		panic(err)
	}
}

//go:noinline
func DWARFCallers(skip int, pcs []uintptr) int {
	pc := getpc(0)
	fmt.Printf("pc: %d\n", pc)
	r := dwarfData.Reader()
	cu, err := r.SeekPC(pc)
	if err != nil {
		panic(err)
	}
	lr, err := dwarfData.LineReader(cu)
	if err != nil {
		panic(err)
	}
	var e dwarf.LineEntry
	if err := lr.SeekPC(pc, &e); err != nil {
		panic(err)
	}
	fmt.Printf("addr: %d\n", e.Address)
	fmt.Printf("%#v\n", e)
	//if entry.Children {
	//for {
	//fmt.Printf("== entry: %s\n", entry.Tag)
	//for _, f := range entry.Field {
	//fmt.Printf("%s: %v\n", f.Attr, f.Val)
	//}
	//entry, err = r.Next()
	//if err != nil {
	//panic(err)
	//}
	//if entry.Tag == 0 {
	//break
	//}
	//}
	//}
	//e := json.NewEncoder(os.Stdout)
	//e.SetIndent("", "  ")
	//e.Encode(entry)
	return 0
}

//go:noinline
//go:nosplit
func getpc(dummyarg int) uint64 {
	fp := uintptr(unsafe.Pointer(&dummyarg)) - 16
	return uint64(deref(fp + 8))
}
