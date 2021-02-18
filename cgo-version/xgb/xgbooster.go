package xgb

// #include <xgboost/c_api.h>
// #include <stdlib.h>
// #include <stdio.h>
import "C"
import (
	"runtime"
	"unsafe"
	"reflect"
)

type XGBooster struct {
	handle C.BoosterHandle
}

func XGBoosterCreate() (*XGBooster, error) {
	var ptr *C.DMatrixHandle
	var out C.BoosterHandle
	res := C.XGBoosterCreate(ptr, 0, &out)
	if err := checkError(res); err != nil {
		return nil, err
	}

	booster := &XGBooster{handle: out}
	runtime.SetFinalizer(booster, boosterFinalizer)

	return booster, nil
}

func (booster *XGBooster) Predict(mat *XGDMatrix) ([]float32, error) {
	var outLen C.bst_ulong
	var outResult *C.float

	optionMask := 0 // normal prediction
	nTreeLimit := 0 // all tress
	training := 0 // not training

	res := C.XGBoosterPredict(booster.handle, mat.handle, C.int(optionMask), C.uint(nTreeLimit), C.int(training), &outLen, &outResult)
	if err := checkError(res); err != nil {
		return nil, err
	}

	var list []float32
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&list)))
	sliceHeader.Cap = int(outLen)
	sliceHeader.Len = int(outLen)
	sliceHeader.Data = uintptr(unsafe.Pointer(outResult))

	runtime.KeepAlive(mat)

	return list, nil
}

func (booster *XGBooster) LoadModel(filePath string) error {
	cfilePath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cfilePath))

	return checkError(C.XGBoosterLoadModel(booster.handle, cfilePath))
}

func boosterFinalizer(booster *XGBooster) {
	C.XGBoosterFree(booster.handle)
}