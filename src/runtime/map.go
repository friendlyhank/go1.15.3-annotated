package runtime

const(
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits
)

type hmap struct {
}

type mapextra struct{}


type bmap struct{
	tophash [bucketCnt]uint8
}
