package main

import "fmt"

const (
	TaskStatusUnaccepted = iota // 未接取
	TaskStatusInProgress        // 进行中
	TaskStatusCompleted         // 已完成
	TaskStatusRewarded          // 已领取
)

type Task struct {
	ID          int
	Name        string
	Desc        string
	Status      int
	PrevTaskID  int
	NextTaskIDs []int
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

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (m *TaskManager) AddTask(task *Task) {
	node := &TaskNode{Task: task}
	if m.head == nil {
		m.head = node
		m.tail = node
	} else {
		m.tail.Next = node
		node.Prev = m.tail
		m.tail = node
	}
	m.count++
}

func (m *TaskManager) RemoveTaskByID(taskID int) bool {
	current := m.head
	for current != nil {
		if current.Task.ID == taskID {
			// 处理前驱节点
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				m.head = current.Next
			}

			// 处理后继节点
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				m.tail = current.Prev
			}

			m.count--
			return true
		}
		current = current.Next
	}
	return false
}

func (m *TaskManager) FindTaskByID(taskID int) *Task {
	current := m.head
	for current != nil {
		if current.Task.ID == taskID {
			return current.Task
		}
		current = current.Next
	}
	return nil
}

func (m *TaskManager) GetTasksByStatus(status int) []*Task {
	var result []*Task
	current := m.head
	for current != nil {
		if current.Task.Status == status {
			result = append(result, current.Task)
		}
		current = current.Next
	}
	return result
}

func (m *TaskManager) UpdateTaskStatus(taskID int, newStatus int) bool {
	task := m.FindTaskByID(taskID)
	if task == nil {
		return false
	}
	task.Status = newStatus

	if newStatus == TaskStatusCompleted {
		for _, nextID := range task.NextTaskIDs {
			nextTask := m.FindTaskByID(nextID)
			if nextTask != nil && nextTask.Status == TaskStatusUnaccepted {
				nextTask.Status = TaskStatusInProgress
				fmt.Printf("任务[%d]完成，已解锁后续任务[%d]\n", taskID, nextID)
			}
		}
	}

	return false
}

func (m *TaskManager) PrintAllTasks() {
	current := m.head
	fmt.Println("当前任务列表：")
	for current != nil {
		statusStr := ""
		switch current.Task.Status {
		case TaskStatusUnaccepted:
			statusStr = "未接取"
		case TaskStatusInProgress:
			statusStr = "进行中"
		case TaskStatusCompleted:
			statusStr = "已完成"
		case TaskStatusRewarded:
			statusStr = "已领奖"
		}
		fmt.Printf("ID: %d, 名称: %s, 状态: %s, 前置任务: %d\n",
			current.Task.ID, current.Task.Name, statusStr, current.Task.PrevTaskID)
		current = current.Next
	}
}

func main() {
	taskMgr := NewTaskManager()

	task1 := &Task{
		ID:          1,
		Name:        "新手引导",
		Desc:        "与村长对话",
		Status:      TaskStatusInProgress,
		PrevTaskID:  -1,
		NextTaskIDs: []int{2}, // 完成后解锁任务2
	}
	task2 := &Task{
		ID:          2,
		Name:        "清理野怪",
		Desc:        "消灭3只史莱姆",
		Status:      TaskStatusUnaccepted,
		PrevTaskID:  1,
		NextTaskIDs: []int{3},
	}
	task3 := &Task{
		ID:          3,
		Name:        "交付物资",
		Desc:        "将木材交给铁匠",
		Status:      TaskStatusUnaccepted,
		PrevTaskID:  2,
		NextTaskIDs: []int{},
	}

	taskMgr.AddTask(task1)
	taskMgr.AddTask(task2)
	taskMgr.AddTask(task3)

	taskMgr.PrintAllTasks()

	fmt.Println("\n--- 玩家完成任务1 ---")
	taskMgr.UpdateTaskStatus(1, TaskStatusCompleted)
	taskMgr.PrintAllTasks()

	fmt.Println("\n--- 玩家领取任务1奖励 ---")
	taskMgr.RemoveTaskByID(1)
	taskMgr.PrintAllTasks()

	inProgress := taskMgr.GetTasksByStatus(TaskStatusInProgress)
	fmt.Println("\n进行中的任务：")
	for _, t := range inProgress {
		fmt.Printf("ID: %d, 名称: %s\n", t.ID, t.Name) // 仅任务2
	}
}
