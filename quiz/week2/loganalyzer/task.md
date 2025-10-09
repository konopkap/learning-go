# Week 2 Bonus Challenge - Server Health Monitor

## **Scenario:**

You're a DevOps engineer managing a fleet of servers. You receive health check results throughout the day and need to process them to identify problems and generate a summary report.

---

## **Sample Data:**

```go
// Server health check results (timestamp, serverName, status, responseTimeMs)
healthChecks := []string{
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
```

---

## **Your Task:**

Write a program that demonstrates Week 2 concepts by processing this data:

### **Requirements:**

1. **Pre-allocated Slices** (demonstrate `make()` with capacity)
   - Create slices to store different categories of results

2. **Slice Operations** (append, slicing, deletion)
   - Filter health checks by status
   - Extract only unhealthy servers

3. **Control Flow** (loops, conditionals)
   - Loop through all health checks
   - Use if/else to categorize by status
   - Count occurrences of each status

4. **Type Conversion** (string to int for response times)
   - Parse the response time from string
   - Calculate averages or find slow responses

### **Suggested Output:**

```
=== Server Health Summary ===

Total Health Checks: 16

Status Breakdown:
- Healthy: 9
- Degraded: 4
- Unhealthy: 3

Servers Requiring Attention:
- web-prod-02 (unhealthy - 1200ms response)
- db-prod-01 (unhealthy - 890ms response)
- web-prod-01 (degraded - 220ms response)
- db-prod-01 (degraded - 340ms response)

Average Response Times:
- Healthy servers: 52ms
- Degraded servers: 337ms
- Unhealthy servers: 1023ms
```

---

## **Bonus Points For:**

- Using constants with `iota` for status levels
- Pre-allocating slices with appropriate capacity
- Clean separation of concerns (multiple functions)
- Handling edge cases (empty data, invalid format)
- Professional output formatting

---

## **Concepts to Demonstrate:**

✅ **From Week 1:**
- Constants and `iota`
- Type conversion
- Printf formatting

✅ **From Week 2:**
- Control flow (for, if/else, switch)
- Slice operations (append, filtering)
- Pre-allocation with `make()`
- Slice len/cap management

---

## **Hints:**

**Parsing a health check line:**
```go
// Split by comma to get fields
// fields[0] = timestamp
// fields[1] = serverName
// fields[2] = status
// fields[3] = responseTime (as string, need to convert to int)
```

**You might need string operations** (not covered in Week 2, so here's a hint):
```go
import "strings"
fields := strings.Split(healthCheck, ",")  // Splits by comma
```

**For string to int conversion:**
```go
import "strconv"
responseTime, err := strconv.Atoi(fields[3])
if err != nil {
    // Handle error
}
```

---

## **Alternative Simpler Scenario (if the above is too complex):**

### **Deployment Task Manager**

```go
// List of deployment tasks with duplicates (tasks get repeated across environments)
tasks := []string{
    "deploy-frontend",
    "run-migrations",
    "deploy-backend",
    "restart-services",
    "deploy-frontend",  // duplicate
    "verify-health",
    "run-migrations",   // duplicate
    "deploy-backend",   // duplicate
    "send-notification",
}
```

**Requirements:**
1. Remove duplicate tasks (Week 2 - you already did this!)
2. Pre-allocate slice for processed tasks
3. Count how many duplicates were found
4. Print task list with numbering
5. Use constants for task categories (if you want to categorize them)

---

**Choose whichever scenario interests you more!** The health monitor is more realistic but complex. The task manager is simpler and focuses purely on slice operations.

**Good luck!** 🚀