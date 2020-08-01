package dto

type JudgeForm struct {
	UseSPJ      bool       `json:"useSPJ"`
	MaxCpuTime  int        `json:"maxCpuTime"`  // 2000 ms
	MaxRealTime int        `json:"maxRealTime"` // 4000 ms
	MaxMemory   int        `json:"maxMemory"`   // 2048*2048kb
	TotalScore  int        `json:"totalScore"`
	CompMp      int        `json:"compMp"` // time multiple of compile
	SPJMp       int        `json:"SPJMp"`  // time multiple of special judge
	Id          int64      `json:"id"`
	Lid         int64      `json:"lid"`
	SPJLid      int64      `json:"SPJLid"`
	SPJCode     string     `json:"SPJCode"`
	Code        string     `json:"code"`
	Flag        string     `json:"flag"`
	ErrorMsg    string     `json:"errorMsg"`
	TestCase    []TestCase `json:"testCase"`
}

type TestCase struct {
	Flag           string `json:"flag"`
	Input          string `json:"input"`
	ExpectedOutput string `json:"expectedOutput"`
	RealOutput     string `json:"realOutput"`
	ErrorOutput    string `json:"errorOutput"`
	SPJOutput      string `json:"SPJOutput"`
	SPJErrorOutput string `json:"SPJErrorOutput"`
	ActualCpuTime  int    `json:"actualCpuTime"`
	ActualRealTime int    `json:"actualRealTime"`
	RealMemory     int    `json:"realMemory"`
	Score          int    `json:"score"`
	Id             int64  `json:"id"` // Problem case Id
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
