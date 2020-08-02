package judge

type Python struct{}

func (p Python) needEditCode() bool {
	return false
}

func (p Python) EditCode(code, name string) (string, error) {
	return code, nil
}

func (p Python) getSourceSuffix() string {
	return ".py"
}

func (p Python) getTargetSuffix() string {
	return ".py"
}

func (p Python) getSPJRunCmd(target, input, expOutput, realOutput string) string {
	return "python3 " + target + " " + input + " " + expOutput + " " + realOutput
}

func (p Python) getSPJCmpCmd(source, target string) string {
	return ""
}

func (p Python) needCompile() bool {
	return false
}

func (p Python) getLangName() string {
	return "Python"
}

func (p Python) getCmpCmd(source, target string) string {
	return ""
}

func (p Python) getRunCmd(target string) string {
	return "python3 " + target
}
