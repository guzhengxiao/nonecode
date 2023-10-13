package nonecode

import "math/rand"

//	func NewHttpServer(e *gin.Engine) *http_server.HttpServer {
//		return http_server.NewHttpServer(e)
//	}
type Ramdom int

const (
	RamdomUpper Ramdom = iota + 1
	RamdomLower
	RamdomAllString
	RamdomNumber
	RamdomUpperNumber
	RamdomLowerNumber
	RamdomAll
)

func (t Ramdom) def() string {
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower := "abcdefghijklmnopqrstuvwxyz"
	number := "0123456789"
	switch t {
	case RamdomUpper:
		return upper
	case RamdomLower:
		return lower
	case RamdomAllString:
		return upper + lower
	case RamdomNumber:
		return upper + number
	case RamdomLowerNumber:
		return lower + number
	case RamdomAll:
		return upper + lower + number
	default:
		return upper + lower + number
	}
	// return upper + lower + number
}

func (t Ramdom) String(n int) string {
	var letters = []rune(t.def())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
