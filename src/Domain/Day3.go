package Domain

type Day3OperationType int

const (
	D3OpMul Day3OperationType = 0
	D3OpEn  Day3OperationType = 1
	D3OpDis Day3OperationType = 2
)

type Day3Operation struct {
	Operation Day3OperationType
	Num1      int
	Num2      int
}
