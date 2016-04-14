package entry

const (
	boolFalse byte = 0x00
	boolTrue  byte = 0x01
)

// Boolean Entry
type Boolean struct {
	Base
}

// BuildFromItems builds a boolean entry using the provided parameters
// func BuildFromItems(name string, id [2]byte, sequence [2]byte, persist byte, value byte) *Boolean {
// 	return &Boolean{
// 		Base{
// 			eName:  name,
// 			eType:  typeBoolean,
// 			eID:    id,
// 			eSeq:   sequence,
// 			eFlag:  persist,
// 			eValue: []byte{value},
// 		},
// 	}
// }
