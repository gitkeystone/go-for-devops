package main

import "fmt"

type TaskType int

const (
	MainTask TaskType = iota
	SideTask
)

type Task struct {
	ID          int
	Name        string
	Description string
	TaskType    TaskType
	Completed   bool
}

type TaskList struct {
	Tasks []Task
}

func (tl *TaskList) AddTask(task Task) {
	tl.Tasks = append(tl.Tasks, task)
}

func (tl *TaskList) CompleteTask(taskID int) {
	for i, task := range tl.Tasks {
		if task.ID == taskID {
			tl.Tasks[i].Completed = true
			fmt.Printf("任务 %s 已完成\n", task.Name)
			return
		}
	}
	fmt.Printf("未找到ID为 %d 的任务\n", taskID)
}

func (tl *TaskList) ShowTask() {
	fmt.Println("任务列表：")
	for _, task := range tl.Tasks {
		status := "未完成"
		if task.Completed {
			status = "已完成"
		}
		taskType := "主线任务"
		if task.TaskType == SideTask {
			taskType = "支线任务"
		}
		fmt.Printf("- %s (%s) - %s: %s\n", task.Name, taskType, status, task.Description)
	}
}

func main() {
	taskList := TaskList{}
	mainTask := Task{
		ID:          1,
		Name:        "击败哥布林",
		Description: "前往哥布林洞穴，击败10只哥布林",
		TaskType:    MainTask,
		Completed:   false,
	}
	taskList.AddTask(mainTask)
	sideTask := Task{
		ID:          2,
		Name:        "收集草药",
		Description: "在森林中收集5株草药",
		TaskType:    SideTask,
		Completed:   false,
	}
	taskList.AddTask(sideTask)

	taskList.ShowTask()
	taskList.CompleteTask(1)
	taskList.ShowTask()
}
