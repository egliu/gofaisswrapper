package gfwcpu

/*
#cgo CFLAGS: -I../include
#cgo LDFLAGS: -L../libs/cpu -lfaiss_c -Wl,-rpath=../libs/cpu
#include "Index_c.h"
#include "AutoTune_c.h"
*/
import "C"
import "fmt"

func example() {
	fmt.Println("Generating some data...")
	// d := 128     // dimension
	// nb := 100000 // database size
	// nq := 10000  // nb of queries
	// xb := make([]float32, d*nb)
	// xq := make([]float32, d*nq)
	// for i := 0; i < nb; i++ {
	// 	for j := 0; j < d; j++ {
	// 	}
	// }
	var f *C.FaissIndex
	C.faiss_index_factory(&f, 128, C.CString("Flat"), C.METRIC_L2)
}
