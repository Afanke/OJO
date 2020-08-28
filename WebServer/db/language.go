package db

type Language struct{}

var lang Language

func (l Language) getName(lid int64) string {
	switch lid {
	case 1:
		return "C"
	case 2:
		return "Cpp"
	case 3:
		return "Java"
	case 4:
		return "Python"
	case 5:
		return "Go"
	default:
		return "C"
	}
}
