package judge

import (
	"errors"
	"github.com/afanke/OJO/utils/log"
	"regexp"
	"strings"
)

type Java struct{}

func (j Java) needEditCode() bool {
	return true
}

func (j Java) EditCode(code, name string) (string, error) {
	reg := regexp.MustCompile(`class ([0-9A-Za-z_$]+?)[ \n]*?[{]`)
	if reg == nil {
		log.Error("failed to regexp")
		return "", errors.New("Java regexp error, please contact administrator\n")
	}
	s := reg.FindAllString(code, 1)
	if len(s) == 0 {
		log.Error("failed to regexp")
		return "", errors.New("Illegal java class name\n")
	}
	code = strings.Replace(code, s[0], "class "+name+" {", 1)
	log.Debug("java code:\n%v", code)
	return code, nil
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
