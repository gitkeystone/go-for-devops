/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const (
	TableWidth      = 145
	TableHeight     = 30
	NamespacePrefix = "ns-"
)

// 指定具体节点
var nodeName string

// 节点标签
var selectors = []string{
	"accelerator=huawei-Ascend910",
	"gpu=true",
}

// chipsCmd represents the chips command
var chipsCmd = &cobra.Command{
	Use:   "chips",
	Short: "显示kubernetes集群种的芯片使用情况",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取kubeconfig文件路径
		var kubeconfig string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		} else {
			kubeconfig = os.Getenv("KUBECONFIG")
		}

		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err)
		}

		// 优化配置
		config.QPS = 50
		config.Burst = 100
		config.Timeout = 30 * time.Second

		// 创建Kubernetes客户端
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}

		columns := []table.Column{
			{Title: "Hostname", Width: 25},
			{Title: "IP", Width: 15},
			{Title: "Namespace", Width: 10},
			{Title: "Pod", Width: 40},
			{Title: "Status", Width: 10},
			{Title: "Chip", Width: 25},
			{Title: "Quantity", Width: 10},
		}

		// 所有节点
		allNodes := make([]*Node, 0)
		for _, s := range selectors {
			nodes, err := GetNodesWithChips(clientset, s)
			if err != nil {
				panic(err)
			}
			// sort
			sort.Slice(nodes, func(i, j int) bool {
				return nodes[i].ChipsUsed < nodes[j].ChipsUsed
			})
			allNodes = append(allNodes, nodes...)
		}

		fmt.Println(len(allNodes))
		// 所有行
		rows := make([]table.Row, 0, 10)

		firstNode := true
		oldNodeName := ""
		for _, node := range allNodes {
			// 每个节点是否空出一行
			// 如果是第一个节点, 起始不会空出一行；其他节点开头空一行；
			if !firstNode {
				rows = append(rows, table.Row{})
			} else {
				firstNode = false
			}

			// 对于没有相关pod的节点
			if len(node.Pods) == 0 {
				row := table.Row{node.Hostname, node.IP, "", "", "", "", ""}
				rows = append(rows, row)
				continue
			}

			// 一行记录一个pod
			for _, pod := range node.Pods {
				row := table.Row{
					"",
					"",
					pod.Namespace,
					pod.Name,
					pod.Status,
					pod.ChipResourceName,
					pod.GetChipQuantity(),
				}
				if node.Hostname != oldNodeName {
					row[0] = node.Hostname
					row[1] = node.IP
				}
				oldNodeName = node.Hostname

				rows = append(rows, row)
			}
		}

		// 排序：按节点使用显卡数量从小到大排序

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithFocused(true),
			table.WithHeight(TableHeight),
			table.WithWidth(TableWidth),
		)

		s := table.DefaultStyles()
		s.Header = s.Header.
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Bold(false)
		s.Selected = s.Selected.
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Bold(false)
		t.SetStyles(s)

		m := model{t}
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(chipsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chipsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chipsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	chipsCmd.Flags().StringVarP(&nodeName, "node", "n", "all", "集群节点, 默认所有节点")
	//chipsCmd.Flags().StringVarP(&labelSelector, "selector", "s", "all", "带芯片的节点标签, 默认为NPU&GPU")
}

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return updateTable(m, msg)
}

func (m model) View() string {
	return tableView(m)
}

type Pod struct {
	Namespace        string
	Name             string
	Status           string
	ChipResourceName string
	ChipQuantity     int64
}

func (p *Pod) GetChipQuantity() string {
	if p.ChipQuantity == 0 {
		return ""
	} else {
		return strconv.FormatInt(p.ChipQuantity, 10)
	}
}

type Node struct {
	Hostname  string
	IP        string
	Pods      []*Pod
	ChipsUsed int64
}

func (n *Node) SetChipsUsed() {
	for _, p := range n.Pods {
		n.ChipsUsed += p.ChipQuantity
	}
}

func GetNodesWithChips(clientset *kubernetes.Clientset, labelSelector string) ([]*Node, error) {
	nodes := make([]*Node, 0)

	// 根据指定标签过滤相应节点
	opt := metav1.ListOptions{
		LabelSelector: labelSelector,
	}
	if nodeName != "all" {
		opt.FieldSelector = "metadata.name=" + nodeName
	}
	nodesWithChips, err := clientset.CoreV1().Nodes().List(context.Background(), opt)
	if err != nil {
		return nil, errors.New("获取NPU节点失败")
	}

	ch := make(chan *Node, 10)
	var wg sync.WaitGroup
	for _, nodeWithChips := range nodesWithChips.Items {
		wg.Add(1)
		go func(nodeWithChips corev1.Node) {
			defer wg.Done()
			node, err := GetChipPodByNode(clientset, nodeWithChips)
			if err != nil {
				panic(err)
			}
			if node != nil {

				ch <- node
			}
		}(nodeWithChips)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for node := range ch {
		nodes = append(nodes, node)
	}

	return nodes, nil
}

var ChipResourceList = []string{
	"huawei.com/Ascend910",
	"nvidia.com/gpu",
}

func GetChipPodByNode(clientset *kubernetes.Clientset, node corev1.Node) (*Node, error) {
	nodeWithChip := new(Node)

	nodeWithChip.Hostname = node.Name

	ipFound := false
	for _, addr := range node.Status.Addresses {
		if addr.Type == corev1.NodeInternalIP {
			nodeWithChip.IP = addr.Address
			ipFound = true
			break
		}
	}
	if !ipFound {
		return nil, fmt.Errorf("节点 %s 没有找到内部IP", node.Name)
	}

	// 使用 fieldSelector 过滤指定节点上的pod
	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{
		FieldSelector: fmt.Sprintf("spec.nodeName=%s", nodeWithChip.Hostname),
	})
	if err != nil {
		return nil, fmt.Errorf("获取节点 %s 上的Pod失败: %v\n", nodeWithChip.Hostname, err)
	}

	ch := make(chan *Pod, 10)

	var wg sync.WaitGroup
	for _, pod := range pods.Items {
		wg.Add(1)
		go func(pod corev1.Pod) {
			defer wg.Done()
			p, err := GetPodInfo(clientset, pod)
			if err != nil {
				panic(err)
			}
			if p != nil {

				ch <- p
			}
		}(pod)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for p := range ch {
		nodeWithChip.Pods = append(nodeWithChip.Pods, p)
	}

	// calculate chips count used per node
	nodeWithChip.SetChipsUsed()

	return nodeWithChip, nil
}

func GetPodInfo(clientset *kubernetes.Clientset, pod corev1.Pod) (*Pod, error) {
	podWithChip := new(Pod)

	if !strings.HasPrefix(pod.Namespace, NamespacePrefix) {
		return nil, nil
	}

	podWithChip.Namespace = pod.Namespace
	podWithChip.Name = pod.Name

	// 获取pod的yaml详情（这里获取pod详情）
	podDetail, err := clientset.CoreV1().Pods(pod.Namespace).Get(context.Background(), pod.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("获取Pod详情失败: %s/%s: %v\n", pod.Namespace, pod.Name, err)
	}

	podWithChip.Status = string(podDetail.Status.Phase)

	for _, chipResouceName := range ChipResourceList {
		if podDetail.Spec.Containers[0].Resources.Limits == nil {
			continue
		}

		for resourceName, quantity := range podDetail.Spec.Containers[0].Resources.Limits {
			if resourceName.String() != chipResouceName {
				continue
			}
			podWithChip.ChipResourceName = resourceName.String()

			// Convert Pod chips limit value to int64
			chipQuantity, ok := quantity.AsInt64()
			if !ok {
				return nil, fmt.Errorf("带芯片的Pod的芯片数量无法转化成数值：%s\n", quantity.String())
			}
			podWithChip.ChipQuantity = chipQuantity

			break
		}
	}
	return podWithChip, nil
}
