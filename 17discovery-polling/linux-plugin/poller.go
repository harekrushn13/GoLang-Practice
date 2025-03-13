package linux_plugin

import (
	"bytes"
	"discoverypolling/utils/logger"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
	"time"
)

type MetricResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Error string `json:"error"`
}

type OverallResult struct {
	TimeStamp string         `json:"timestamp"`
	Metrics   []MetricResult `json:"metrics"`
}

var allMetricsCmd = `echo -e "` +
	`Overall Memory Free (bytes): $(free -b | awk '/Mem:/ {print $4}')\n` +
	`Overall Memory Used (bytes): $(free -b | awk '/Mem:/ {print $3}')\n` +
	`Installed Memory (bytes): $(free -b | awk '/Mem:/ {print $2}')\n` +
	`Available Memory (bytes): $(free -b | awk '/Mem:/ {print $7}')\n` +
	`Cache Memory (bytes): $(free -b | awk '/Mem:/ {print $6}')\n` +
	`Buffer Memory (bytes): $(free -b | awk '/Mem:/ {print $6}')\n` +
	`Overall Memory Used (%): $(free -b | awk '/Mem:/ {print ($3/$2)*100}')\n` +
	`Overall Memory Free (%): $(free -b | awk '/Mem:/ {print ($4/$2)*100}')\n` +
	`Swap Memory Free (bytes): $(free -b | awk '/Swap:/ {print $4}')\n` +
	`Swap Memory Used (bytes): $(free -b | awk '/Swap:/ {print $3}')\n` +
	`Swap Memory Used (%): $(free -b | awk '/Swap:/ {if ($2 > 0) print ($3/$2)*100; else print 0}')\n` +
	`Swap Memory Free (%): $(free -b | awk '/Swap:/ {if ($2 > 0) print ($4/$2)*100; else print 0}')\n` +
	`Load Avg (1 min): $(uptime | awk '{print $8}' | tr -d ',')\n` +
	`Load Avg (5 min): $(uptime | awk '{print $9}' | tr -d ',')\n` +
	`Load Avg (15 min): $(uptime | awk '{print $10}' | tr -d ',')\n` +
	`CPU Cores: $(nproc)\n` +
	`CPU Usage (%): $(top -bn1 | awk '/%Cpu/ {print 100 - $8}')\n` +
	`CPU Kernel Usage (%): $(cat /proc/stat | awk '/cpu / {print ($2+$4)*100/($2+$4+$5)}')\n` +
	`CPU Idle (%): $(cat /proc/stat | awk '/cpu / {print $5*100/($2+$4+$5)}')\n` +
	`CPU Interrupts (%): $(cat /proc/stat | awk '/cpu / {print $7*100/($2+$4+$5)}')\n` +
	`CPU I/O Wait (%): $(iostat -c | awk 'NR==4 {print $4}')\n` +
	`Disk Capacity (bytes): $(df -B1 / | awk 'NR==2 {print $2}')\n` +
	`Disk Free (bytes): $(df -B1 / | awk 'NR==2 {print $4}')\n` +
	`Disk Used (bytes): $(df -B1 / | awk 'NR==2 {print $3}')\n` +
	`Disk Free (%): $(df -h / | awk 'NR==2 {print $4}' | tr -d '%')\n` +
	`Disk Used (%): $(df -h / | awk 'NR==2 {print $5}' | tr -d '%')\n` +
	`TCP Connections: $(ss -t | wc -l)\n` +
	`UDP Connections: $(ss -u | wc -l)\n` +
	`Network Error Packets: $(cat /proc/net/dev | awk 'NR>2 {sum += $4+$12} END {print sum}')\n` +
	`Running Processes: $(ps -e | wc -l)\n` +
	`Blocked Processes: $(ps -e -o stat | grep 'D' | wc -l)\n` +
	`Threads Count: $(cat /proc/stat | awk '/processes/ {print $2}')\n` +
	`OS Name: $(uname -s)\n` +
	`OS Version: $(lsb_release -rs 2>/dev/null || echo 'N/A')\n` +
	`System Name: $(hostname)\n` +
	`System Uptime: $(uptime -p | cut -d' ' -f2-)\n` +
	`Uptime in Seconds: $(cat /proc/uptime | awk '{print $1}')\n` +
	`Context Switches: $(cat /proc/stat | awk '/ctxt/ {print $2}')"`

func CollectLinuxData(username, password, host, port string) string {

	client, err := connectToSSH(username, password, host, port)

	if err != nil {
		logger.Error("SSH connection failed", err)
		return ""
	}
	defer client.Close()

	logger.Info("SSH connection successful")

	allMetricsOutput, err := executeCommand(client, allMetricsCmd)
	if err != nil {
		logger.Error("Failed to execute combined metrics command", err)
		return ""
	}

	metrics := parseAllMetrics(allMetricsOutput)

	overallResult := OverallResult{
		TimeStamp: time.Now().Format(time.RFC3339),
		Metrics:   metrics,
	}

	jsonData, err := json.MarshalIndent(overallResult, "", "  ")
	if err != nil {
		logger.Error("Failed to marshal JSON", err)
		return ""
	}

	return string(jsonData)
}

func connectToSSH(user, password, host, port string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}
	return ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), config)
}

func executeCommand(client *ssh.Client, cmd string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create SSH session: %v", err)
	}
	defer session.Close()

	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	if err := session.Run(cmd); err != nil {
		return "", fmt.Errorf("command failed: %v | stderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}

func parseAllMetrics(output string) []MetricResult {
	metricMapping := map[string]string{
		"Overall Memory Free (bytes)": "system.overall.memory.free.bytes",
		"Overall Memory Used (bytes)": "system.overall.memory.used.bytes",
		"Installed Memory (bytes)":    "system.memory.installed.bytes",
		"Available Memory (bytes)":    "system.memory.available.bytes",
		"Cache Memory (bytes)":        "system.cache.memory.bytes",
		"Buffer Memory (bytes)":       "system.buffer.memory.bytes",
		"Overall Memory Used (%)":     "system.overall.memory.used.percent",
		"Overall Memory Free (%)":     "system.overall.memory.free.percent",
		"Swap Memory Free (bytes)":    "system.swap.memory.free.bytes",
		"Swap Memory Used (bytes)":    "system.swap.memory.used.bytes",
		"Swap Memory Used (%)":        "system.swap.memory.used.percent",
		"Swap Memory Free (%)":        "system.swap.memory.free.percent",
		"Load Avg (1 min)":            "system.load.avg1.min",
		"Load Avg (5 min)":            "system.load.avg5.min",
		"Load Avg (15 min)":           "system.load.avg15.min",
		"CPU Cores":                   "system.cpu.cores",
		"CPU Usage (%)":               "system.cpu.percent",
		"CPU Kernel Usage (%)":        "system.cpu.kernel.percent",
		"CPU Idle (%)":                "system.cpu.idle.percent",
		"CPU Interrupts (%)":          "system.cpu.interrupt.percent",
		"CPU I/O Wait (%)":            "system.cpu.io.percent",
		"Disk Capacity (bytes)":       "system.disk.capacity.bytes",
		"Disk Free (bytes)":           "system.disk.free.bytes",
		"Disk Used (bytes)":           "system.disk.used.bytes",
		"Disk Free (%)":               "system.disk.free.percent",
		"Disk Used (%)":               "system.disk.used.percent",
		"TCP Connections":             "system.network.tcp.connections",
		"UDP Connections":             "system.network.udp.connections",
		"Network Error Packets":       "system.network.error.packets",
		"Running Processes":           "system.running.processes",
		"Blocked Processes":           "system.blocked.processes",
		"Threads Count":               "system.threads",
		"OS Name":                     "system.os.name",
		"OS Version":                  "system.os.version",
		"System Name":                 "system.name",
		"System Uptime":               "started.time",
		"Uptime in Seconds":           "started.time.seconds",
		"Context Switches":            "system.context.switches.per.sec",
	}

	var metrics []MetricResult
	lines := strings.Split(strings.TrimSpace(output), "\n")

	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) != 2 {
			continue
		}

		metricName, metricValue := parts[0], strings.TrimSpace(parts[1])

		if metricKey, found := metricMapping[metricName]; found {
			metrics = append(metrics, MetricResult{Name: metricKey, Value: metricValue})
		}
	}

	return metrics
}
