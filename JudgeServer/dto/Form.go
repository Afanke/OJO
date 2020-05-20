package dto

type JudgeForm struct {
	UseSPJ      bool
	MaxCpuTime  int // 2 seconds
	MaxRealTime int // 4 seconds
	MaxMemory   int // 2097152 bytes
	TotalScore  int
	Id          int64
	SPJCode     string
	Code        string
	Flag        string
	TestCase    []TestCase
}

type TestCase struct {
	Flag           string
	Input          string
	ExpectOutput   string
	RealOutput     string
	ErrorOutput    string
	SPJOutput      string
	SPJErrorOutput string
	ActualCpuTime  int
	ActualRealTime int
	RealMemory     int
	Score          int
	Id             int64 // Problem case Id
}

type TempStorage struct {
	FilePath string
	SPJPath  string
	CmdLine  string
}

type Res struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
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
