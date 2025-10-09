package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Using iota to replace indexes / sizes with names in array operations
const (
	healthy = iota
	degraded
	unhealthy
	statusCount
)

// Below variable created as working with files has not been covered yet
// Server health check results (timestamp, serverName, status, responseTimeMs)
var (
	healthChecks = []string{
		"2025-10-06T08:00:00,web-prod-01,healthy,45",
		"2025-10-06T08:00:00,web-prod-02,healthy,52",
		"2025-10-06T08:00:00,db-prod-01,healthy,120",
		"2025-10-06T08:00:00,cache-prod-01,healthy,8",
		"2025-10-06T08:05:00,web-prod-01,healthy,48",
		"2025-10-06T08:05:00,web-prod-02,degraded,450",
		"2025-10-06T08:05:00,db-prod-01,healthy,115",
		"2025-10-06T08:05:00,cache-prod-01,healthy,12",
		"2025-10-06T08:10:00,web-prod-01,healthy,44",
		"2025-10-06T08:10:00,web-prod-02,unhealthy,980",
		"2025-10-06T08:10:00,db-prod-01,degraded,340",
		"2025-10-06T08:10:00,cache-prod-01,healthy,9",
		"2025-10-06T08:15:00,web-prod-01,degraded,220",
		"2025-10-06T08:15:00,web-prod-02,unhealthy,1200",
		"2025-10-06T08:15:00,db-prod-01,unhealthy,890",
		"2025-10-06T08:15:00,cache-prod-01,healthy,11",
	}

	noOfHealthchecks   = len(healthChecks)
	hostStatusCount    [statusCount]int
	avgResponseTimesMs [statusCount]int

	// Known size so using make() to optimize array size
	timeStamps      = make([]string, 0, noOfHealthchecks)
	serverNames     = make([]string, 0, noOfHealthchecks)
	statuses        = make([]string, 0, noOfHealthchecks)
	responseTimesMs = make([]int, 0, noOfHealthchecks)

	//Unknown size
	abnormalIndexes = []int{}
)

func main() {
	fmt.Printf("=== Server Health Summary ===\n\n")

	fmt.Printf("Analyzing logs...\n\n")
	hostStatusCount, avgResponseTimesMs = analyzeLog(healthChecks)

	fmt.Printf("Total Health Checks: %v\n\n", noOfHealthchecks)

	fmt.Println("Status Breakdown:")
	fmt.Println("Healthy:", hostStatusCount[healthy], "| Average response time [ms]:", avgResponseTimesMs[healthy])
	fmt.Println("Degraded:", hostStatusCount[degraded], "| Average response time [ms]:", avgResponseTimesMs[degraded])
	fmt.Println("Unhealthy:", hostStatusCount[unhealthy], "| Average response time [ms]:", avgResponseTimesMs[unhealthy])
	fmt.Println()

	fmt.Println("Servers requiring attention:")
	for _, index := range abnormalIndexes {
		fmt.Println("-", serverNames[index], "(", statuses[index], "-", responseTimesMs[index], "ms response)")
	}
}

// abnormalIndexes is used to stored positions of servers which are not healthy
func analyzeLog(logs []string) (hostStatusCount, avgResponseTimesMs [3]int) {
	for i, logLine := range logs {
		fillSlices(logLine)
		// Initially I put this as separate function - but it required another iteration run through logs
		// It was cleaner but less optimized so I decided to put it there
		switch statuses[i] {
		case "healthy":
			hostStatusCount[healthy]++
			avgResponseTimesMs[healthy] += responseTimesMs[i]
		case "degraded":
			hostStatusCount[degraded]++
			abnormalIndexes = append(abnormalIndexes, i)
			avgResponseTimesMs[degraded] += responseTimesMs[i]
		case "unhealthy":
			hostStatusCount[unhealthy]++
			abnormalIndexes = append(abnormalIndexes, i)
			avgResponseTimesMs[unhealthy] += responseTimesMs[i]
		}
	}
	for i := range avgResponseTimesMs {
		avgResponseTimesMs[i] = int(avgResponseTimesMs[i] / hostStatusCount[i])
	}
	return
}

func fillSlices(logLine string) {
	timeStamp, serverName, status, responseTimeMs := analyzeLine(logLine)
	timeStamps = append(timeStamps, timeStamp)
	serverNames = append(serverNames, serverName)
	statuses = append(statuses, status)
	responseTimesMs = append(responseTimesMs, responseTimeMs)
}

func analyzeLine(logLine string) (timeStamp, serverName, serverStatus string, responseTimeMs int) {
	splitLog := strings.Split(logLine, ",")
	timeStamp = splitLog[healthy]
	serverName = splitLog[degraded]
	serverStatus = splitLog[unhealthy]
	responseTimeMs, err := strconv.Atoi(splitLog[3])
	if err != nil {
		fmt.Println("[ERROR] The last item is not a number!")
		fmt.Println(err)
		return
	}
	return
}
