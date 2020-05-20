package operator

import "github.com/afanke/OJO/JudgeServer/dto"

type Operator interface {
	Operate(form *dto.JudgeForm)
	afterRun(form *dto.JudgeForm)
	judge(form *dto.JudgeForm)
	run(form *dto.JudgeForm)
	beforeRun(form *dto.JudgeForm)
}
