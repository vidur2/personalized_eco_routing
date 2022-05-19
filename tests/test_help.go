package tests

/*
#cgo LDFLAGS: /Users/vidurmodgil/Desktop/ProgrammingProjects/line_integrals_fuel_efficiency/lib/libregression.a -ldl
#include "/Users/vidurmodgil/Desktop/ProgrammingProjects/line_integrals_fuel_efficiency/lib/regression.h"
*/
import "C"

import (
	"encoding/json"
	"fmt"
	datahandler "line_integrals_fuel_efficiency/dataHandler"
)

func testRegression() {
	coord := []float32{5., 1.}
	coords := make([][]float32, 0)
	coords = append(coords, coord)
	coord = []float32{2., 3.}
	coords = append(coords, coord)
	coord = []float32{1., 1.}
	coords = append(coords, coord)
	fmt.Println(coords)
	datapoints := datahandler.RegressionInputs{
		XValues: coords,
		YValues: []float32{1.0, 2.0, 0.},
	}

	final, _ := json.Marshal(datapoints)

	model := C.train_regression(C.CString(string(final)))

	fmt.Println(C.GoString(model))

	output := C.predict_regression(C.CString("[[3.1, 1.2]]"), model)

	fmt.Println(C.GoString(output))
}
