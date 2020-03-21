package dto

type OperationForm struct {
	MaxCpuTime     int // 2 seconds
	ActualCpuTime  int
	MaxRealTime    int // 4 seconds
	ActualRealTime int
	MaxMemory      int // 2097152 bytes
	RealMemory     int
	Score          int
	PcId           int // Problem case Id
	Language       string
	FilePath       string
	CmdLine        string // use to start program or see version
	Code           string
	Input          string
	ExpectOutput   string
	RealOutput     string
	ErrorOutput    string
	Flag           string
}

// type RequestForm struct {
// 	MaxCpuTime     int // 2 seconds
// 	MaxRealTime    int // 4 seconds
// 	MaxMemory      int // 2097152 bytes
// 	Score          int
// 	PcId           int // Problem case Id
// 	SmId           int // Submission Id/
// 	Cid            int // Contest Id
// 	Language     string
// 	Code         string
// 	Input        string
// 	ExpectOutput string
// }
//
// type ResponseForm struct {
// 	ActualCpuTime  int
// 	ActualRealTime int
// 	RealMemory     int
// 	Score          int
// 	PcId           int // Problem case Id
// 	SmId           int // Submission Id
// 	Cid            int // Contest Id
// 	RealOutput   string
// 	ErrorOutput  string
// 	Flag         string
// }
