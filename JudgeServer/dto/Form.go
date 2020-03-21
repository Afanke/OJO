package dto

type OperationForm struct {
	MaxCpuTime     int //2 seconds
	ActualCpuTime  int
	MaxRealTime    int //4 seconds
	ActualRealTime int
	MaxMemory      int // 2097152 bytes
	RealMemory     int
	Score          int
	PcId           int //Problem case Id
	// SmId           int //Submission Id
	// Cid            int //Contest Id
	Language     string
	FilePath     string
	CmdLine      string //use to start program or see version
	Code         string
	Input        string
	ExpectOutput string
	RealOutput   string
	ErrorOutput  string
	Flag         string
}
