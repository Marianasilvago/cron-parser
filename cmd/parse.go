package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse [cron string]",
	Short: "Parses a cron expression and outputs the times at which it runs",
	Long: `The parse command parses a standard cron expression and outputs 
the schedule for each of the five fields: minute, hour, day of month, month, 
and day of week, followed by the command.

Cron expressions consist of 5 fields followed by a command:
  * Minute (0-59)
  * Hour (0-23)
  * Day of month (1-31)
  * Month (1-12)
  * Day of week (0-6 where 0 is Sunday)
  
Example usage:
  parse "*/15 0 1,15 * 1-5 /usr/bin/find"

This will print the expanded schedule for each field of the cron expression.`,

	Example: `  parse "*/15 0 1,15 * 1-5 /usr/bin/find"
  parse "0 12 * * 0 echo 'Run at noon on Sundays'"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cronString := args[0]
		parseCron(cronString)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}

// parseCron parses the provided cron string and prints the expanded fields
func parseCron(cronString string) {
	fields := strings.Fields(cronString)

	// Ensure that there are at least 6 fields (5 for cron, 1 for the command)
	if len(fields) < 6 {
		fmt.Println("Invalid cron string format. Expected 5 fields followed by a command.")
		return
	}

	// The first 5 fields are for the cron expression
	minute, hour, dayOfMonth, month, dayOfWeek := fields[0], fields[1], fields[2], fields[3], fields[4]
	// The rest is the command (join all remaining fields into a single string)
	command := strings.Join(fields[5:], " ")

	fmt.Printf("minute        %s\n", expandField(minute, 0, 59))
	fmt.Printf("hour          %s\n", expandField(hour, 0, 23))
	fmt.Printf("day of month  %s\n", expandField(dayOfMonth, 1, 31))
	fmt.Printf("month         %s\n", expandField(month, 1, 12))
	fmt.Printf("day of week   %s\n", expandField(dayOfWeek, 0, 6))
	fmt.Printf("command       %s\n", command)
}

// expandField is a helper function that expands the cron field (similar to the one discussed earlier)
func expandField(field string, min, max int) string {
	// Implementation for expanding the cron field
	if field == "*" {
		return rangeToString(min, max)
	}
	if strings.Contains(field, "*/") {
		step, _ := strconv.Atoi(strings.TrimPrefix(field, "*/"))
		return stepRangeToString(min, max, step)
	}
	values := []string{}
	for _, part := range strings.Split(field, ",") {
		if strings.Contains(part, "-") {
			bounds := strings.Split(part, "-")
			start, _ := strconv.Atoi(bounds[0])
			end, _ := strconv.Atoi(bounds[1])
			values = append(values, rangeToString(start, end))
		} else {
			values = append(values, part)
		}
	}
	return strings.Join(values, " ")
}

func rangeToString(start, end int) string {
	values := []string{}
	for i := start; i <= end; i++ {
		values = append(values, fmt.Sprintf("%d", i))
	}
	return strings.Join(values, " ")
}

func stepRangeToString(start, end, step int) string {
	values := []string{}
	for i := start; i <= end; i += step {
		values = append(values, fmt.Sprintf("%d", i))
	}
	return strings.Join(values, " ")
}
