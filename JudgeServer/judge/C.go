package judge

type C struct{}

func (c C) getCmpCmd(source, target string) string {
	return "gcc " + source + " -o " + target
}

func (c C) getRunCmd(target string) string {
	return "./" + target
}

func (c C) needCompile() bool {
	return true
}

func (c C) getSourceSuffix() string {
	return ".c"
}

func (c C) getTargetSuffix() string {
	return ""
}

func (c C) getLangName() string {
	return "C"
}

func (c C) getSPJCmpCmd(source, target string) string {
	return "gcc " + source + " -o " + target
}

func (c C) getSPJRunCmd(target, input, expOutput, realOutput string) string {
	return "./" + target + " " + input + " " + expOutput + " " + realOutput
}
