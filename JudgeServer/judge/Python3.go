package judge

type Python3 struct{}

func (p Python3) needEditCode() bool {
	return false
}

func (p Python3) EditCode(code string) string {
	return code
}

func (p Python3) getSourceSuffix() string {
	return ".py"
}

func (p Python3) getTargetSuffix() string {
	return ".py"
}

func (p Python3) getSPJRunCmd(target, input, expOutput, realOutput string) string {
	return "python3 " + target + " " + input + " " + expOutput + " " + realOutput
}

func (p Python3) getSPJCmpCmd(source, target string) string {
	return ""
}

func (p Python3) needCompile() bool {
	return false
}

func (p Python3) getLangName() string {
	return "Python3"
}

func (p Python3) getCmpCmd(source, target string) string {
	return ""
}

func (p Python3) getRunCmd(target string) string {
	return "python3 " + target
}
