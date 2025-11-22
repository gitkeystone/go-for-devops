package main

const (
	TaskStatusUnaccepted = iota // 未接取
	TaskStatusInProgress        // 进行中
	TaskStatusCompleted         // 已完成
	TaskStatusRewarded          // 已领取
)

type Task struct {
	ID         int
	Name       string
	Desc       string
	Status     int
	PreTaskID  int
	NextTaskID []int
}

type TaskNode struct {
	Task *Task
	Prev *TaskNode
	Next *TaskNode
}

type TaskManager struct {
	head  *TaskNode
	tail  *TaskNode
	count int
}

func main() {

}
