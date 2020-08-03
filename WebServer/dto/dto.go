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
type UpdateForm struct {
	Password string `json:"password"`
	New      string `json:"new"`
	Id       int64  `json:"id"`
}

type RegisterForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}
type ProblemForm struct {
	Mine       bool   `json:"mine" db:"mine"`
	Cid        int64  `json:"cid" db:"cid"`
	Page       int    `json:"page" db:"page"`
	Offset     int    `json:"offset" db:"offset"`
	Limit      int    `json:"limit" db:"limit"`
	Difficulty string `json:"difficulty" db:"difficulty"`
	Keywords   string `json:"keywords" db:"keywords"`
}
type RankForm struct {
	Page   int `json:"page" db:"page"`
	Offset int `json:"offset" db:"offset"`
	Limit  int `json:"limit" db:"limit"`
}
type UserForm struct {
	Page     int    `json:"page" db:"page"`
	Offset   int    `json:"offset" db:"offset"`
	Limit    int    `json:"limit" db:"limit"`
	Keywords string `json:"keywords" db:"keywords"`
	Type     int    `json:"type" db:"type"`
}
type UserBrief struct {
	Id            int64  `json:"id" db:"id"`
	Type          int    `json:"type" db:"type"`
	Enabled       bool   `json:"enabled" db:"enabled"`
	Username      string `json:"username" db:"username"`
	Email         string `json:"email" db:"email"`
	CreateTime    string `json:"createTime" db:"create_time"`
	LastLoginTime string `json:"lastLoginTime" db:"last_login_time"`
	IconPath      string `json:"iconPath" db:"icon_path"`
	RealName      string `json:"realName" db:"real_name"`
}
type UserDetail struct {
	Id        int64  `json:"id" db:"id"`
	Type      int    `json:"type" db:"type"`
	Username  string `json:"username" db:"username"`
	Email     string `json:"email" db:"email"`
	RealName  string `json:"realName" db:"real_name"`
	Signature string `json:"signature" db:"signature"`
	School    string `json:"school" db:"school"`
	Blog      string `json:"blog" db:"blog"`
	IconPath  string `json:"iconPath" db:"icon_path"`
	Major     string `json:"major" db:"major"`
	Github    string `json:"github" db:"github"`
}

type UserDetail2 struct {
	Id        int64  `json:"id" db:"id"`
	Type      int    `json:"type" db:"type"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Email     string `json:"email" db:"email"`
	RealName  string `json:"realName" db:"real_name"`
	Signature string `json:"signature" db:"signature"`
	IconPath  string `json:"iconPath" db:"icon_path"`
	School    string `json:"school" db:"school"`
	Blog      string `json:"blog" db:"blog"`
	Major     string `json:"major" db:"major"`
	Github    string `json:"github" db:"github"`
}

type JudgeServer struct {
	Enabled bool   `json:"enabled" db:"enabled"`
	Status  bool   `json:"status" db:"status"`
	Id      int64  `json:"id" db:"id"`
	Port    int    `json:"port" db:"port"`
	Weight  int    `json:"weight" db:"weight"`
	Name    string `json:"name" db:"name"`
	Address string `json:"address" db:"address"`
}

type PracticeForm struct {
	Page       int    `json:"page" db:"page"`
	Tid        int64  ` json:"tid" db:"tid"`
	Offset     int    `json:"offset" db:"offset"`
	Limit      int    `json:"limit" db:"limit"`
	Difficulty string `json:"difficulty" db:"difficulty"`
	Keywords   string `json:"keywords" db:"keywords"`
}
type PracticeBrief struct {
	Id    int64  `json:"id" db:"id"`
	Cid   int64  `json:"cid" db:"cid"`
	Ref   string `json:"ref" db:"ref"`
	Title string `json:"title" db:"title"`
	// Description string             `json:"description" db:"description"`
	Difficulty string             `json:"difficulty" db:"difficulty"`
	Tags       []TagBrief         `json:"tags"`
	Statistic  *PracticeStatistic `json:"statistic"`
}
type Practice struct {
	Visible           bool   `json:"visible" db:"visible"`
	Id                int64  `json:"id" db:"id"`
	Cid               int64  `json:"cid" db:"cid"`
	MemoryLimit       int    `json:"memoryLimit" db:"memory_limit"`
	CpuTimeLimit      int    `json:"cpuTimeLimit" db:"cpu_time_limit"`
	RealTimeLimit     int    `json:"realTimeLimit" db:"real_time_limit"`
	Ref               string `json:"ref" db:"ref"`
	Hint              string `json:"hint" db:"hint"`
	Title             string `json:"title" db:"title"`
	Source            string `json:"source" db:"source"`
	CreateTime        string `json:"createTime" db:"create_time"`
	Difficulty        string `json:"difficulty" db:"difficulty"`
	Description       string `json:"description" db:"description"`
	LastUpdateTime    string `json:"lastUpdateTime" db:"last_update_time"`
	InputDescription  string `json:"inputDescription" db:"input_description"`
	OutputDescription string `json:"outputDescription" db:"output_description"`
	// ---------------------------------------------------
	CreatorName string             `json:"creatorName"`
	Languages   []Language         `json:"language"`
	Samples     []ProblemSample    `json:"sample"`
	Tags        []TagBrief         `json:"tag"`
	Statistic   *PracticeStatistic `json:"statistic"`
}
type Id struct {
	Id int64 `json:"id"`
}
type Id2 struct {
	Cid int64 `json:"cid"`
	Pid int64 `json:"pid"`
}
type Id3 struct {
	Id int64 `json:"id"`
}
type Id4 struct {
	Cid int64 `json:"cid"`
	Pid int64 `json:"pid"`
}
type PracticeSubStat struct {
	Id          int64  `json:"id" db:"id"`
	Uid         int64  `json:"uid" db:"uid"`
	Pid         int64  `json:"pid" db:"pid"`
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
	// Mode           int // 1:OI 2:ACM
	PcId int64 // Problem case Id
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
	Cid      int64  `json:"cid" db:"cid"` // Contest Id
	Sid      int64  `json:"sid" db:"sid"` // Submission Id
	Uid      int64  `json:"uid" db:"uid"` // User Id
	Pid      int64  `json:"pid" db:"pid"` // Problem Id
	Language string `json:"language" db:"language"`
	Code     string `json:"code" db:"code"`
}
type ContestBrief struct {
	Id        int64  `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Rule      string `json:"rule" db:"rule"`
	StartTime string `json:"startTime" db:"start_time"`
	EndTime   string `json:"endTime" db:"end_time"`
	Now       string `json:"now"`
}
type ContestDetail struct {
	Id          int64  `json:"id" db:"id"`
	Cid         int64  `json:"cid" db:"cid"`
	Punish      int    `json:"punishTime" db:"punish"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Rule        string `json:"rule" db:"rule"`
	StartTime   string `json:"startTime" db:"start_time"`
	EndTime     string `json:"endTime" db:"end_time"`
	Now         string `json:"now"`
	CreatorName string `json:"creatorName" db:"creatorName"`
}
type ContestForm struct {
	Mine     bool   `json:"mine"`
	Cid      int64  `json:"cid" db:"cid"`
	Page     int    `json:"page" db:"page"`
	Rule     string `json:"rule" db:"rule"`
	Status   int    `json:"status" db:"status"`
	Keywords string `json:"keywords" db:"keywords"`
	Offset   int    `json:"offset" db:"offset"`
	Limit    int    `json:"limit" db:"limit"`
}
type ContestQualifyForm struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}
type CtsPbBrief struct {
	Id        int64             `json:"id" db:"id"`
	Ref       string            `json:"ref" db:"ref"`
	Title     string            `json:"title" db:"title"`
	Statistic *ContestStatistic `json:"statistic"`
}
type ContestProblem struct {
	Id                int64  `json:"id" db:"id"`
	Cid               int64  `json:"cid" db:"cid"`
	MemoryLimit       int    `json:"memoryLimit" db:"memory_limit"`
	CpuTimeLimit      int    `json:"cpuTimeLimit" db:"cpu_time_limit"`
	RealTimeLimit     int    `json:"realTimeLimit" db:"real_time_limit"`
	Ref               string `json:"ref" db:"ref"`
	Hint              string `json:"hint" db:"hint"`
	Title             string `json:"title" db:"title"`
	Source            string `json:"source" db:"source"`
	CreateTime        string `json:"createTime" db:"create_time"`
	Difficulty        string `json:"difficulty" db:"difficulty"`
	Description       string `json:"description" db:"description"`
	LastUpdateTime    string `json:"lastUpdateTime" db:"last_update_time"`
	InputDescription  string `json:"inputDescription" db:"input_description"`
	OutputDescription string `json:"outputDescription" db:"output_description"`
	// ---------------------------------------------------
	CreatorName string            `json:"creatorName"`
	Languages   []Language        `json:"language"`
	Samples     []ProblemSample   `json:"sample"`
	Tags        []TagBrief        `json:"tag"`
	Statistic   *ContestStatistic `json:"statistic"`
}

type ContestCaseResult struct {
	CpuTime     int    `json:"cpuTime" db:"cpu_time"`
	Csmid       int64  `json:"csmid" db:"csmid"`
	ErrorOutput string `json:"errorOutput" db:"error_output"`
	Flag        string `json:"flag" db:"flag"`
	Id          int64  `json:"id" db:"id"`
	Pcaseid     int64  `json:"pcaseid" db:"pcaseid"`
	RealMemory  int    `json:"realMemory" db:"real_memory"`
	RealOutput  string `json:"realOutput" db:"real_output"`
	RealTime    int    `json:"realTime" db:"real_time"`
	Score       int    `json:"score" db:"score"`
	Uid         int64  `json:"uid" db:"uid"`
}
type ContestSubStat struct {
	Id          int64  `json:"id" db:"id"`
	Uid         int64  `json:"uid" db:"uid"`
	Cid         int64  `json:"cid" db:"cid"`
	Pid         int64  `json:"pid" db:"pid"`
	TotalScore  int    `json:"totalScore" db:"total_score"`
	Language    string `json:"language" db:"language"`
	Status      string `json:"status" db:"status"`
	SubmitTime  string `json:"submitTime" db:"submit_time"`
	Code        string `json:"code" db:"code"`
	ProblemName string `json:"problemName" db:"problem_name"`
	Username    string `json:"username" db:"username"`
}
type OIRank struct {
	Cid            int64      `json:"cid" db:"cid"`
	Uid            int64      `json:"uid" db:"uid"`
	Username       string     `json:"username" db:"username"`
	TotalScore     int        `json:"totalScore" db:"total_score"`
	LastSubmitTime string     `json:"lastSubmitTime" db:"last_submit_time"`
	OIDetail       []OIDetail `json:"OIDetail" db:"oi_detail"`
}
type OIDetail struct {
	Pid      int64 `json:"pid" db:"pid"`
	MaxScore int   `json:"maxScore" db:"max_score"`
}

type ACMRank struct {
	Id        int         `json:"id" db:"id"`
	Cid       int64       `json:"cid" db:"cid"`
	Uid       int64       `json:"uid" db:"uid"`
	Total     int         `json:"total" db:"total"`
	AC        int         `json:"ac" db:"ac"`
	TotalTime int         `json:"totalTime" db:"total_time"`
	Username  string      `json:"username" db:"username"`
	ACMDetail []ACMDetail `json:"ACMDetail" db:"acm_detail"`
}

type ACMRank2 struct {
	Uid       int64  `json:"uid" db:"uid"`
	Total     int    `json:"total" db:"total"`
	AC        int    `json:"ac" db:"ac"`
	Username  string `json:"username" db:"username"`
	Signature string `json:"signature" db:"signature"`
}

type ACMDetail struct {
	Id             int64 `json:"id" db:"id"`
	Cid            int64 `json:"cid" db:"cid"`
	Uid            int64 `json:"uid" db:"uid"`
	Pid            int64 `json:"pid" db:"pid"`
	LastSubmitTime int64 `json:"lastSubmitTime" db:"last_submit_time"`
	Total          int   `json:"total" db:"total"`
	AC             bool  `json:"ac" db:"ac"`
	FirstAC        bool  `json:"firstAC" db:"first_ac"`
}
type UserStatistic struct {
	AC         int   `json:"ac"`
	Submission int   `json:"submission"`
	Score      int   `json:"score"`
	SolvedList []int `json:"solvedList"`
}

// ------------------------------------------------------------------------------
type ContestSubmission struct {
	Cid        int64  `json:"cid" db:"cid"`
	Id         int64  `json:"id" db:"id"`
	Pid        int64  `json:"pid" db:"pid"`
	Code       string `json:"code" db:"code"`
	TotalScore int    `json:"totalScore" db:"total_score"`
	Uid        int64  `json:"uid" db:"uid"`
	Language   string `json:"language" db:"language"`
	Status     string `json:"status" db:"status"`
	SubmitTime string `json:"submitTime" db:"submit_time"`
}
type Administrator struct {
	Id            int64  `json:"id" db:"id"`
	Name          string `json:"name" db:"name"`
	Account       string `json:"account" db:"account"`
	Password      string `json:"password" db:"password"`
	CreateTime    string `json:"createTime" db:"create_time"`
	LastLoginTime string `json:"lastLoginTime" db:"last_login_time"`
}
type Contest struct {
	Visible        bool             `json:"visible" db:"visible"`
	Id             int64            `json:"id" db:"id"`
	Cid            int64            `json:"cid" db:"cid"`
	Punish         int              `json:"punish" db:"punish"`
	SubmitLimit    int              `json:"submitLimit" db:"submit_limit"`
	Title          string           `json:"title" db:"title"`
	Description    string           `json:"description" db:"description"`
	Rule           string           `json:"rule" db:"rule"`
	StartTime      string           `json:"startTime" db:"start_time"`
	Now            string           `json:"now"`
	EndTime        string           `json:"endTime" db:"end_time"`
	CreateTime     string           `json:"createTime" db:"create_time"`
	LastUpdateTime string           `json:"lastUpdateTime" db:"last_update_time"`
	CreatorName    string           `json:"creatorName"`
	Password       string           `json:"password" db:"password"`
	IPLimit        []ContestIPLimit `json:"IPLimit"`
}
type ContestIPLimit struct {
	Id      int    `json:"id" db:"id"`
	Cid     int    `json:"cid" db:"cid"`
	Address string `json:"address" db:"address"`
}
type ContestStatistic struct {
	Pid   int64 `json:"pid" db:"pid"`
	Total int   `json:"total" db:"total"`
	Ac    int   `json:"ac" db:"ac"`
	Ce    int   `json:"ce" db:"ce"`
	Cid   int   `json:"cid" db:"cid"`
	Id    int   `json:"id" db:"id"`
	Mle   int   `json:"mle" db:"mle"`
	Ole   int   `json:"ole" db:"ole"`
	Re    int   `json:"re" db:"re"`
	Tle   int   `json:"tle" db:"tle"`
	Wa    int   `json:"wa" db:"wa"`
}
type Language struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
type PracticeCaseResult struct {
	Id          int64  `json:"id" db:"id"`
	Psmid       int64  `json:"psmid" db:"psmid"`
	Pcaseid     int64  `json:"pcaseid" db:"pcaseid"`
	Uid         int64  `json:"uid" db:"uid"`
	Flag        string `json:"flag" db:"flag"`
	CpuTime     int    `json:"cpuTime" db:"cpu_time"`
	RealTime    int    `json:"realTime" db:"real_time"`
	RealMemory  int    `json:"realMemory" db:"real_memory"`
	RealOutput  string `json:"realOutput" db:"real_output"`
	ErrorOutput string `json:"errorOutput" db:"error_output"`
	Score       int    `json:"score" db:"score"`
}
type PracticeStatistic struct {
	Id    int64 `json:"id" db:"id"`
	Pbid  int64 `json:"pbid" db:"pbid"`
	Total int   `json:"total" db:"total"`
	Ac    int   `json:"ac" db:"ac"`
	Wa    int   `json:"wa" db:"wa"`
	Re    int   `json:"re" db:"re"`
	Tle   int   `json:"tle" db:"tle"`
	Mle   int   `json:"mle" db:"mle"`
	Ce    int   `json:"ce" db:"ce"`
	Ole   int   `json:"ole" db:"ole"`
}
type PracticeSubmission struct {
	Id         int64  `json:"id" db:"id"`
	Uid        int64  `json:"uid" db:"uid"`
	Pid        int64  `json:"pid" db:"pid"`
	Language   string `json:"language" db:"language"`
	Status     string `json:"status" db:"status"`
	TotalScore int    `json:"totalScore" db:"total_score"`
	SubmitTime string `json:"submitTime" db:"submit_time"`
	Code       string `json:"code" db:"code"`
}
type ProblemBrief struct {
	Shared         bool       `json:"shared" db:"shared"`
	Visible        bool       `json:"visible" db:"visible"`
	Id             int64      `json:"id" db:"id"`
	Cid            int64      `json:"cid" db:"cid"`
	CreatorName    string     `json:"creatorName"`
	Ref            string     `json:"ref" db:"ref"`
	Title          string     `json:"title" db:"title"`
	Difficulty     string     `json:"difficulty" db:"difficulty"`
	CreateTime     string     `json:"createTime" db:"create_time"`
	LastUpdateTime string     `json:"lastUpdateTime" db:"last_update_time"`
	Tags           []TagBrief `json:"tags"`
}
type Problem struct {
	Id                int64           `json:"id" db:"id"`
	Cid               int64           `json:"cid" db:"cid"`
	Ref               string          `json:"ref" db:"ref"`
	Title             string          `json:"title" db:"title"`
	Description       string          `json:"description" db:"description"`
	InputDescription  string          `json:"inputDescription" db:"input_description"`
	OutputDescription string          `json:"outputDescription" db:"output_description"`
	Hint              string          `json:"hint" db:"hint"`
	CreateTime        string          `json:"createTime" db:"create_time"`
	LastUpdateTime    string          `json:"lastUpdateTime" db:"last_update_time"`
	Difficulty        string          `json:"difficulty" db:"difficulty"`
	Source            string          `json:"source" db:"source"`
	Visible           bool            `json:"visible" db:"visible"`
	UseSPJ            bool            `json:"useSPJ" db:"use_spj"`
	Shared            bool            `json:"shared" db:"shared"`
	SPJ               SPJ             `json:"spj" db:"spj"`
	Limit             []ProblemLimit  `json:"limit" db:"limit"`
	Template          []Template      `json:"template" db:"template"`
	ProblemCase       []ProblemCase   `json:"problemCase" db:"problem_case"`
	Language          []Language      `json:"language" db:"language"`
	Sample            []ProblemSample `json:"sample" db:"sample"`
	Tag               []TagBrief      `json:"tag"`
}

type SPJ struct {
	Id   int64  `json:"id" db:"id"`
	Pid  int64  `json:"pid" db:"pid"`
	Lid  int64  `json:"lid" db:"lid"`
	Code string `json:"code" db:"code"`
}

type Template struct {
	Id      int64  `json:"id" db:"id"`
	Pid     int64  `json:"pid" db:"pid"`
	Lid     int64  `json:"lid" db:"lid"`
	Prepend string `json:"prepend" db:"prepend"`
	Content string `json:"content" db:"content"`
	Append  string `json:"append" db:"append"`
}

type ProblemLimit struct {
	Id          int `json:"id" db:"id"`
	Pid         int `json:"pid" db:"pid"`
	Lid         int `json:"lid" db:"lid"`
	MaxCpuTime  int `json:"maxCpuTime" db:"max_cpu_time"`
	MaxRealTime int `json:"maxRealTime" db:"max_real_time"`
	MaxMemory   int `json:"maxMemory" db:"max_memory"`
	CompMp      int `json:"compMp" db:"comp_mp"`
	SPJMp       int `json:"SPJMp" db:"spj_mp"`
}

type ProblemCase struct {
	Id     int64  `json:"id" db:"id"`
	Pid    int64  `json:"pid" db:"pid"`
	Input  string `json:"input" db:"input"`
	Output string `json:"output" db:"output"`
	Score  int    `json:"score" db:"score"`
}
type ProblemLanguage struct {
	Id  int64 `json:"id" db:"id"`
	Pid int64 `json:"pid" db:"pid"`
	Lid int64 `json:"lid" db:"lid"`
}
type ProblemSample struct {
	Id     int64  `json:"id" db:"id"`
	Pid    int64  `json:"pid" db:"pid"`
	Input  string `json:"input" db:"input"`
	Output string `json:"output" db:"output"`
}
type ProblemTag struct {
	Id  int64 `json:"id" db:"id"`
	Tid int64 `json:"tid" db:"tid"`
	Pid int64 `json:"pid" db:"pid"`
}
type Progress struct {
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Progress int    `json:"progress" db:"progress"`
}

type TagForm struct {
	Mine     bool   `json:"mine"`
	Cid      int64  `json:"cid" db:"cid"`
	Page     int    `json:"page" db:"page"`
	Offset   int    `json:"offset" db:"offset"`
	Limit    int    `json:"limit" db:"limit"`
	Keywords string `json:"keywords" db:"keywords"`
}

type Tag struct {
	Id             int64  `json:"id" db:"id"`
	Cid            int64  `json:"cid" db:"cid"`
	Visible        bool   `json:"visible" db:"visible"`
	Shared         bool   `json:"shared" db:"shared"`
	Name           string `json:"name" db:"name"`
	CreatorName    string `json:"creatorName"`
	CreateTime     string `json:"createTime" db:"create_time"`
	LastUpdateTime string `json:"lastUpdateTime" db:"last_update_time"`
}

type TagBrief struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Username struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

type UsernameAndSig struct {
	Id        int64  `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Signature string `json:"signature" db:"signature"`
}

type UserToken struct {
	Enabled  bool   `json:"enabled" db:"enabled"`
	Id       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Type     int    `json:"type" db:"type"`
}
type User struct {
	Id            int64  `json:"id" db:"id"`
	Username      string `json:"username" db:"username"`
	Password      string `json:"password" db:"password"`
	Email         string `json:"email" db:"email"`
	CreateTime    string `json:"createTime" db:"create_time"`
	LastLoginTime string `json:"lastLoginTime" db:"last_login_time"`
	Type          int    `json:"type" db:"type"`
	Enabled       bool   `json:"enabled" db:"enabled"`
	IconPath      string `json:"iconPath" db:"icon_path"`
	RealName      string `json:"realName" db:"real_name"`
	Signature     string `json:"signature" db:"signature"`
	School        string `json:"school" db:"school"`
	Blog          string `json:"blog" db:"blog"`
	Major         string `json:"major" db:"major"`
	Github        string `json:"github" db:"github"`
}
type SystemConfig struct {
	Server        string `json:"server" db:"server"`
	Port          int    `json:"port" db:"port"`
	Email         string `json:"email" db:"email"`
	Password      string `json:"password" db:"password"`
	Name          string `json:"name" db:"name"`
	Footer        string `json:"footer" db:"footer"`
	AllowRegister bool   `json:"allowRegister" db:"allow_register"`
}

type Announcement struct {
	Visible        bool   `json:"visible" db:"visible"`
	Id             int64  `json:"id" db:"id"`
	Cid            int64  `json:"cid" db:"cid"`
	CreateTime     string `json:"createTime" db:"create_time"`
	LastUpdateTime string `json:"lastUpdateTime" db:"last_update_time"`
	CreatorName    string `json:"creatorName"`
	Title          string `json:"title" db:"title"`
	Content        string `json:"content" db:"content"`
}

type AnnouncementForm struct {
	Mine     bool   `json:"mine" db:"mine"`
	Cid      int64  `json:"cid" db:"cid"`
	Page     int    `json:"page" db:"page"`
	Offset   int    `json:"offset" db:"offset"`
	Limit    int    `json:"limit" db:"limit"`
	Keywords string `json:"keywords" db:"keywords"`
}
type TodayCount struct {
	Hour  string `json:"hour" db:"hour"`
	Count int    `json:"count" db:"count"`
}

type WeekCount struct {
	Today    string     `json:"today"`
	DayCount []DayCount `json:"dayCount"`
}

type MonthCount struct {
	Today    string     `json:"today"`
	DayCount []DayCount `json:"dayCount"`
}

type DayCount struct {
	Day   string `json:"day" db:"day"`
	Count int    `json:"count" db:"count"`
}

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
