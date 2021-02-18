package xgb

// #include <xgboost/c_api.h>
// #include <stdlib.h>
// #include <stdio.h>
import "C"
import (
	"errors"
	"runtime"
)

type XGDMatrix struct {
	handle C.DMatrixHandle
	cols   int
	rows   int
}

func matrixFinalizer(mat *XGDMatrix) {
	C.XGDMatrixFree(mat.handle)
}

func XGDMatrixCreateFromMat(data []float32, nrows int, ncols int, missing float32) (*XGDMatrix, error) {
	if len(data) != nrows*ncols {
		return nil, errors.New("data length doesn't match given dimensions")
	}

	var out C.DMatrixHandle
	res := C.XGDMatrixCreateFromMat((*C.float)(&data[0]), C.bst_ulong(nrows), C.bst_ulong(ncols), C.float(missing), &out)
	if err := checkError(res); err != nil {
		return nil, err
	}

	matrix := &XGDMatrix{handle: out, rows: nrows, cols: ncols}
	runtime.SetFinalizer(matrix, matrixFinalizer)
	runtime.KeepAlive(data)

	return matrix, nil
}

func checkError(res C.int) error {
	if int(res) != 0 {
		return errors.New(C.GoString(C.XGBGetLastError()))
	}
	return nil
}