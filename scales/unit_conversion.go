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
func ToConvertBytes(byteVal float32) (status int, converted map[string]interface{}) {
	var kb float32
	var mb float32
	var gb float32

	kb = byteVal / BYTE
	mb = byteVal / KBYTE
	gb = byteVal / MBYTE

	// debug
	// fmt.Printf("kb: %g, mb: %g, gb: %g", kb, mb, gb)

	converted = map[string]interface{}{
		"kb": kb,
		"mb": mb,
		"gb": gb,
	}
	return
}
