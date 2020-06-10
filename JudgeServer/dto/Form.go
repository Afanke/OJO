package dto

type JudgeForm struct {
	UseSPJ      bool
	MaxCpuTime  int // 2000 ms
	MaxRealTime int // 4000 ms
	MaxMemory   int // 2048*2048kb
	TotalScore  int
	Id          int64
	Lid         int64
	SPJLid      int64
	SPJCode     string
	Code        string
	Flag        string
	ErrorMsg    string
	TestCase    []TestCase
}

type TestCase struct {
	Flag           string
	Input          string
	ExpectedOutput string
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
	UseSPJ   bool
	FilePath string
	SPJPath  string
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
