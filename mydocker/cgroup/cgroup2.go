package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

// 这是 systemd 用于管理系统服务的 cgroup 层级
// 它可以用于限制和监控系统服务的资源使用（如 CPU、内存、I/O 等）
// 每个系统服务都有一个对应的 cgroup 目录，包含资源控制文件
const cgroupServiceHierarchyMount = "/sys/fs/cgroup/system.slice"

func main() {
	if os.Args[0] == "/proc/self/exe" {
		fmt.Printf("current pid %d\n", syscall.Getpid())

		cmd := exec.Command("sh", "-c", "stress -m 1 --vm-bytes 200m --vm-keep")
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		check(err)
	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	check(err)

	// 得到clone出来进程映像在外部命名空间的pid
	fmt.Printf("%v\n", cmd.Process.Pid)

	// 在系统默认创建的Hierarchy上创建对应于系统服务的cgroup
	err = os.MkdirAll(path.Join(cgroupServiceHierarchyMount, "test-limit-memory"), 0755)
	check(err)

	err = os.WriteFile(path.Join(cgroupServiceHierarchyMount, "test-limit-memory", "cgroup.procs"),
		[]byte(strconv.Itoa(cmd.Process.Pid)), 0644)
	check(err)

	err = os.WriteFile(path.Join(cgroupServiceHierarchyMount, "test-limit-memory", "memory.max"),
		[]byte("100m"), 0644)
	check(err)

	err = cmd.Wait()
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
