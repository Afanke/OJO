package dto

type Res struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}
type RegisterForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}
type PracticeForm struct {
	Page       int    `json:"page" db:"page"`
	Tid        int    ` json:"tid" db:"tid"`
	Offset     int    `json:"offset" db:"offset"`
	Limit      int    `json:"limit" db:"limit"`
	Difficulty string `json:"difficulty" db:"difficulty"`
	Keywords   string `json:"keywords" db:"keywords"`
}
type PracticeBrief struct {
	Id          int                `json:"id" db:"id"`
	Cid         int                `json:"cid" db:"cid"`
	Ref         string             `json:"ref" db:"ref"`
	Title       string             `json:"title" db:"title"`
	Description string             `json:"description" db:"description"`
	Difficulty  string             `json:"difficulty" db:"difficulty"`
	Tags        []Tags             `json:"tags"`
	Statistic   *PracticeStatistic `json:"statistic"`
}
type Practice struct {
	Show              bool   `json:"show" db:"show"`
	Id                int    `json:"id" db:"id"`
	Cid               int    `json:"cid" db:"cid"`
	MemoryLimit       int    `json:"memoryLimit" db:"memory_limit"`
	CpuTimeLimit      int    `json:"cpuTimeLimit" db:"cpu_time_limit"`
	RealTimeLimit     int    `json:"realTimeLimit" db:"real_time_limit"`
	Ref               string `json:"ref" db:"ref"`
	Hint              string `json:"hint" db:"hint"`
	Title             string `json:"title" db:"title"`
	IoMode            string `json:"ioMode" db:"io_mode"`
	Source            string `json:"source" db:"source"`
	CreateTime        string `json:"createTime" db:"create_time"`
	Difficulty        string `json:"difficulty" db:"difficulty"`
	Description       string `json:"description" db:"description"`
	LastUpdateTime    string `json:"lastUpdateTime" db:"last_update_time"`
	InputDescription  string `json:"inputDescription" db:"input_description"`
	OutputDescription string `json:"outputDescription" db:"output_description"`
	//---------------------------------------------------
	CreatorName string             `json:"creatorName"`
	Languages   []Language         `json:"languages"`
	Samples     []Sample           `json:"samples"`
	Tags        []Tags             `json:"tags"`
	Statistic   *PracticeStatistic `json:"statistic"`
}
type Id struct {
	Id int `json:"id"`
}
type Id2 struct {
	Cid int `json:"cid"`
	Pid int `json:"pid"`
}
type PracticeSubStat struct {
	Id          int    `json:"id" db:"id"`
	Uid         int    `json:"uid" db:"uid"`
	Pid         int    `json:"pid" db:"pid"`
	TotalScore  int    `json:"totalScore" db:"total_score"`
	Language    string `json:"language" db:"language"`
	Status      string `json:"status" db:"status"`
	SubmitTime  string `json:"submitTime" db:"submit_time"`
	Code        string `json:"code" db:"code"`
	ProblemName string `json:"problemName" db:"problem_name"`
	Username    string `json:"username" db:"username"`
}
type OperationForm struct {
	MaxCpuTime     int // 2 seconds
	ActualCpuTime  int
	MaxRealTime    int // 4 seconds
	ActualRealTime int
	MaxMemory      int // 2097152 bytes
	RealMemory     int
	Score          int
	PcId           int // Problem case Id
	// SmId           int // Submission Id
	// Cid            int // Contest Id
	Language     string
	FilePath     string
	CmdLine      string // use to start program or see version
	Code         string
	Input        string
	ExpectOutput string
	RealOutput   string
	ErrorOutput  string
	Flag         string
}
type SubmitForm struct {
	Cid      int    `json:"cid" db:"cid"` // Contest Id
	Sid      int    `json:"sid" db:"sid"` // Submission Id
	Uid      int    `json:"uid" db:"uid"` // User Id
	Pid      int    `json:"pid" db:"pid"` // Problem Id
	Language string `json:"language" db:"language"`
	Code     string `json:"code" db:"code"`
}
type ContestBrief struct {
	Id        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Rule      string `json:"rule" db:"rule"`
	StartTime string `json:"startTime" db:"start_time"`
	EndTime   string `json:"endTime" db:"end_time"`
	Now       string `json:"now"`
}
type ContestDetail struct {
	Id          int    `json:"id" db:"id"`
	Cid         int    `json:"cid" db:"cid"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Rule        string `json:"rule" db:"rule"`
	StartTime   string `json:"startTime" db:"start_time"`
	EndTime     string `json:"endTime" db:"end_time"`
	Now         string `json:"now"`
	CreatorName string `json:"creatorName" db:"creatorName"`
}
type ContestForm struct {
	Cid      int    `json:"cid" db:"cid"`
	Page     int    `json:"page" db:"page"`
	Rule     string `json:"rule" db:"rule"`
	Status   int    `json:"status" db:"status"`
	Keywords string `json:"keywords" db:"keywords"`
	Offset   int    `json:"offset" db:"offset"`
	Limit    int    `json:"limit" db:"limit"`
}
type ContestQualifyForm struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}
type CtsPbBrief struct {
	Id        int               `json:"id" db:"id"`
	Ref       string            `json:"ref" db:"ref"`
	Title     string            `json:"title" db:"title"`
	Statistic *ContestStatistic `json:"statistic"`
}
type ContestProblem struct {
	Id                int    `json:"id" db:"id"`
	Cid               int    `json:"cid" db:"cid"`
	MemoryLimit       int    `json:"memoryLimit" db:"memory_limit"`
	CpuTimeLimit      int    `json:"cpuTimeLimit" db:"cpu_time_limit"`
	RealTimeLimit     int    `json:"realTimeLimit" db:"real_time_limit"`
	Ref               string `json:"ref" db:"ref"`
	Hint              string `json:"hint" db:"hint"`
	Title             string `json:"title" db:"title"`
	IoMode            string `json:"ioMode" db:"io_mode"`
	Source            string `json:"source" db:"source"`
	CreateTime        string `json:"createTime" db:"create_time"`
	Difficulty        string `json:"difficulty" db:"difficulty"`
	Description       string `json:"description" db:"description"`
	LastUpdateTime    string `json:"lastUpdateTime" db:"last_update_time"`
	InputDescription  string `json:"inputDescription" db:"input_description"`
	OutputDescription string `json:"outputDescription" db:"output_description"`
	// ---------------------------------------------------
	CreatorName string            `json:"creatorName"`
	Languages   []Language        `json:"languages"`
	Samples     []Sample          `json:"samples"`
	Tags        []Tags            `json:"tags"`
	Statistic   *ContestStatistic `json:"statistic"`
}
type ContestCaseResult struct {
	CpuTime     int    `json:"cpuTime" db:"cpu_time"`
	Csmid       int    `json:"csmid" db:"csmid"`
	ErrorOutput string `json:"errorOutput" db:"error_output"`
	Flag        string `json:"flag" db:"flag"`
	Id          int    `json:"id" db:"id"`
	Pcaseid     int    `json:"pcaseid" db:"pcaseid"`
	RealMemory  int    `json:"realMemory" db:"real_memory"`
	RealOutput  string `json:"realOutput" db:"real_output"`
	RealTime    int    `json:"realTime" db:"real_time"`
	Score       int    `json:"score" db:"score"`
	Uid         int    `json:"uid" db:"uid"`
}
type ContestSubStat struct {
	Id          int    `json:"id" db:"id"`
	Uid         int    `json:"uid" db:"uid"`
	Cid         int    `json:"cid" db:"cid"`
	Pid         int    `json:"pid" db:"pid"`
	TotalScore  int    `json:"totalScore" db:"total_score"`
	Language    string `json:"language" db:"language"`
	Status      string `json:"status" db:"status"`
	SubmitTime  string `json:"submitTime" db:"submit_time"`
	Code        string `json:"code" db:"code"`
	ProblemName string `json:"problemName" db:"problem_name"`
	Username    string `json:"username" db:"username"`
}
type OIRank struct {
	Cid            int        `json:"cid" db:"cid"`
	Uid            int        `json:"uid" db:"uid"`
	Username       string     `json:"username" db:"username"`
	TotalScore     int        `json:"totalScore" db:"total_score"`
	LastSubmitTime string     `json:"lastSubmitTime" db:"last_submit_time"`
	OIDetail       []OIDetail `json:"OIDetail" db:"oi_detail"`
}
type OIDetail struct {
	Pid      int `json:"pid" db:"pid"`
	MaxScore int `json:"maxScore" db:"max_score"`
}

// ------------------------------------------------------------------------------
type ContestSubmission struct {
	Cid        int    `json:"cid" db:"cid"`
	Id         int    `json:"id" db:"id"`
	Pid        int    `json:"pid" db:"pid"`
	Code       string `json:"code" db:"code"`
	TotalScore int    `json:"totalScore" db:"total_score"`
	Uid        int    `json:"uid" db:"uid"`
	Language   string `json:"language" db:"language"`
	Status     string `json:"status" db:"status"`
	SubmitTime string `json:"submitTime" db:"submit_time"`
}
type Administrator struct {
	Id            int    `json:"id" db:"id"`
	Name          string `json:"name" db:"name"`
	Account       string `json:"account" db:"account"`
	Password      string `json:"password" db:"password"`
	CreateTime    string `json:"createTime" db:"create_time"`
	LastLoginTime string `json:"lastLoginTime" db:"last_login_time"`
}
type Contest struct {
	Id             int    `json:"id" db:"id"`
	Title          string `json:"title" db:"title"`
	Description    string `json:"description" db:"description"`
	Rule           string `json:"rule" db:"rule"`
	StartTime      string `json:"startTime" db:"start_time"`
	EndTime        string `json:"endTime" db:"end_time"`
	CreateTime     string `json:"createTime" db:"create_time"`
	LastUpdateTime string `json:"lastUpdateTime" db:"last_update_time"`
	Cid            int    `json:"cid" db:"cid"`
	Password       string `json:"password" db:"password"`
}
type ContestStatistic struct {
	Total int `json:"total" db:"total"`
	Ac    int `json:"ac" db:"ac"`
	Ce    int `json:"ce" db:"ce"`
	Cid   int `json:"cid" db:"cid"`
	Id    int `json:"id" db:"id"`
	Mle   int `json:"mle" db:"mle"`
	Ole   int `json:"ole" db:"ole"`
	Pid   int `json:"pid" db:"pid"`
	Re    int `json:"re" db:"re"`
	Tle   int `json:"tle" db:"tle"`
	Wa    int `json:"wa" db:"wa"`
}
type Language struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
type PracticeCaseResult struct {
	Id          int    `json:"id" db:"id"`
	Psmid       int    `json:"psmid" db:"psmid"`
	Pcaseid     int    `json:"pcaseid" db:"pcaseid"`
	Uid         int    `json:"uid" db:"uid"`
	Flag        string `json:"flag" db:"flag"`
	CpuTime     int    `json:"cpuTime" db:"cpu_time"`
	RealTime    int    `json:"realTime" db:"real_time"`
	RealMemory  int    `json:"realMemory" db:"real_memory"`
	RealOutput  string `json:"realOutput" db:"real_output"`
	ErrorOutput string `json:"errorOutput" db:"error_output"`
	Score       int    `json:"score" db:"score"`
}
type PracticeStatistic struct {
	Id    int `json:"id" db:"id"`
	Pbid  int `json:"pbid" db:"pbid"`
	Total int `json:"total" db:"total"`
	Ac    int `json:"ac" db:"ac"`
	Wa    int `json:"wa" db:"wa"`
	Re    int `json:"re" db:"re"`
	Tle   int `json:"tle" db:"tle"`
	Mle   int `json:"mle" db:"mle"`
	Ce    int `json:"ce" db:"ce"`
	Ole   int `json:"ole" db:"ole"`
}
type PracticeSubmission struct {
	Id         int    `json:"id" db:"id"`
	Uid        int    `json:"uid" db:"uid"`
	Pid        int    `json:"pid" db:"pid"`
	Language   string `json:"language" db:"language"`
	Status     string `json:"status" db:"status"`
	TotalScore int    `json:"totalScore" db:"total_score"`
	SubmitTime string `json:"submitTime" db:"submit_time"`
	Code       string `json:"code" db:"code"`
}
type Problem struct {
	Id                int    `json:"id" db:"id"`
	Cid               int    `json:"cid" db:"cid"`
	Ref               string `json:"ref" db:"ref"`
	Title             string `json:"title" db:"title"`
	Description       string `json:"description" db:"description"`
	InputDescription  string `json:"inputDescription" db:"input_description"`
	OutputDescription string `json:"outputDescription" db:"output_description"`
	Hint              string `json:"hint" db:"hint"`
	CreateTime        string `json:"createTime" db:"create_time"`
	LastUpdateTime    string `json:"lastUpdateTime" db:"last_update_time"`
	CpuTimeLimit      int    `json:"cpuTimeLimit" db:"cpu_time_limit"`
	MemoryLimit       int    `json:"memoryLimit" db:"memory_limit"`
	IoMode            string `json:"ioMode" db:"io_mode"`
	Difficulty        string `json:"difficulty" db:"difficulty"`
	RealTimeLimit     int    `json:"realTimeLimit" db:"real_time_limit"`
	Source            string `json:"source" db:"source"`
	Show              bool   `json:"show" db:"show"`
}
type ProblemCase struct {
	Id     int    `json:"id" db:"id"`
	Pid    int    `json:"pid" db:"pid"`
	Input  string `json:"input" db:"input"`
	Output string `json:"output" db:"output"`
	Score  int    `json:"score" db:"score"`
}
type ProblemLanguage struct {
	Id  int `json:"id" db:"id"`
	Pid int `json:"pid" db:"pid"`
	Lid int `json:"lid" db:"lid"`
}
type ProblemSample struct {
	Id  int `json:"id" db:"id"`
	Pid int `json:"pid" db:"pid"`
	Sid int `json:"sid" db:"sid"`
}
type ProblemTags struct {
	Id  int `json:"id" db:"id"`
	Tid int `json:"tid" db:"tid"`
	Pid int `json:"pid" db:"pid"`
}
type Progress struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Progress int    `json:"progress" db:"progress"`
}
type Sample struct {
	Id     int    `json:"id" db:"id"`
	Input  string `json:"input" db:"input"`
	Output string `json:"output" db:"output"`
}
type Tags struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
type User struct {
	Id            int    `json:"id" db:"id"`
	Username      string `json:"username" db:"username"`
	Password      string `json:"password" db:"password"`
	Email         string `json:"email" db:"email"`
	CreateTime    string `json:"createTime" db:"create_time"`
	LastLoginTime string `json:"lastLoginTime" db:"last_login_time"`
}
