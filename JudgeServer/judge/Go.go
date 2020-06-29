package judge

type Go struct{}

func (g Go) getCmpCmd(source, target string) string {
	return "go build " + " -o " + target + " " + source
}

func (g Go) getRunCmd(target string) string {
	return "./" + target
}

func (g Go) needCompile() bool {
	return true
}

func (g Go) needEditCode() bool {
	return false
}

func (g Go) EditCode(code, name string) (string, error) {
	return code, nil
}

func (g Go) getSourceSuffix() string {
	return ".go"
}

func (g Go) getTargetSuffix() string {
	return ""
}

func (g Go) getLangName() string {
	return "Go"
}

func (g Go) getSPJCmpCmd(source, target string) string {
	return "go build " + source + " -o " + target
}

func (g Go) getSPJRunCmd(target, input, expOutput, realOutput string) string {
	return "./" + target + " " + input + " " + expOutput + " " + realOutput
}
