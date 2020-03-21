package operator

import "github.com/afanke/OJO/JudgeServer/dto"

type Operator interface {
	Operate(form *dto.OperationForm)
	afterRun(form *dto.OperationForm)
	judge(form *dto.OperationForm)
	run(form *dto.OperationForm)
	beforeRun(form *dto.OperationForm)
}
