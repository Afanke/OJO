package judge

import "strings"

type Java struct{}

func (j Java) needEditCode() bool {
	return true
}

func (j Java) EditCode(code string) string {
	return code
}

func (j Java) getCmpCmd(source, target string) string {
	return "javac " + source
}

func (j Java) getRunCmd(target string) string {
	return "java " + strings.Replace(target, ".class", "", 1)
}

func (j Java) needCompile() bool {
	return true
}

func (j Java) getSourceSuffix() string {
	return ".java"
}

func (j Java) getTargetSuffix() string {
	return ".java"
}

func (j Java) getLangName() string {
	return "Java"
}

func (j Java) getSPJCmpCmd(source, target string) string {
	return ""
}

func (j Java) getSPJRunCmd(target, input, expOutput, realOutput string) string {
	return "java " + target + " " + input + " " + expOutput + " " + realOutput
}
