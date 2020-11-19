package runtime


// Needs to be in sync with ../cmd/link/internal/ld/decodesym.go:/^func.commonsize,
// ../cmd/compile/internal/gc/reflect.go:/^func.dcommontype and
// ../reflect/type.go:/^type.rtype.
// ../internal/reflectlite/type.go:/^type.rtype.
type _type struct {

}

type typeOff int32

type imethod struct{}

type interfacetype struct{}

type maptype struct{}

type arraytype struct{

}

type chantype struct{}

type slicetype struct{}

type functype struct{}

type ptrtype struct{}

type structtype struct{}
