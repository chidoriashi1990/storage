package scales

const (
	_ = iota
	// BYTE ..
	BYTE float32 = 1 << (10 * iota)
	// KBYTE ..
	KBYTE float32 = 1 << (10 * iota)
	// MBYTE ..
	MBYTE float32 = 1 << (10 * iota)
)

// ToConvertBytes is a function to convert bytes
func ToConvertBytes(byteVal int64) (converted map[string]interface{}) {
	var byteFloat float32
	var b int64
	var kb float32
	var mb float32
	var gb float32

	byteFloat = float32(byteVal)

	b = byteVal
	kb = byteFloat / BYTE
	mb = byteFloat / KBYTE
	gb = byteFloat / MBYTE

	// debug
	// fmt.Printf("kb: %g, mb: %g, gb: %g", kb, mb, gb)

	converted = map[string]interface{}{
		"b":  b,
		"kb": kb,
		"mb": mb,
		"gb": gb,
	}
	return
}
