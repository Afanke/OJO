package judge

type Cpp struct{}

func (c Cpp) needEditCode() bool {
	return false
}

func (c Cpp) EditCode(code string) string {
	return code
}

func (c Cpp) getCmpCmd(source, target string) string {
	return "g++ " + source + " -o " + target
}

func (c Cpp) getRunCmd(target string) string {
	return "./" + target
}

func (c Cpp) needCompile() bool {
	return true
}

func (c Cpp) getSourceSuffix() string {
	return ".cpp"
}

func (c Cpp) getTargetSuffix() string {
	return ""
}

func (c Cpp) getLangName() string {
	return "Cpp"
}

func (c Cpp) getSPJCmpCmd(source, target string) string {
	return "g++ " + source + " -o " + target
}

func (c Cpp) getSPJRunCmd(target, input, expOutput, realOutput string) string {
	return "./" + target + " " + input + " " + expOutput + " " + realOutput
}
