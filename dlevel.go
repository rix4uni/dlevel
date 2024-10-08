package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

// prints the version message
const version = "0.0.2"

func printVersion() {
	fmt.Printf("Current dlevel version %s\n", version)
}

// Prints the Colorful banner
func printBanner() {
	banner := `
       __ __                   __
  ____/ // /___  _   __ ___   / /
 / __  // // _ \| | / // _ \ / / 
/ /_/ // //  __/| |/ //  __// /  
\__,_//_/ \___/ |___/ \___//_/   
`
fmt.Printf("%s\n%70s\n\n", banner, "Current dlevel version "+version)
}

func main() {
	// Define the --level, --max-level, --min-level, --until-count, and --until-level flags
	level := pflag.String("level", "", "Specify the subdomain levels (comma-separated) to filter")
	maxLevel := pflag.Bool("max-level", false, "Print subdomains from max dots to min dots")
	minLevel := pflag.Bool("min-level", false, "Print subdomains from min dots to max dots")
	untilCount := pflag.Int("until-count", -1, "Stop after printing this many lines")
	untilLevel := pflag.Int("until-level", -1, "Stop after reaching this level (dot count)")
	silent := pflag.Bool("silent", false, "silent mode.")
	version := pflag.Bool("version", false, "Print the version of the tool and exit.")
	pflag.Parse()

	// Print version and exit if -version flag is provided
	if *version {
		printBanner()
		printVersion()
		return
	}

	// Don't Print banner if -silnet flag is provided
	if !*silent {
		printBanner()
	}

	// Ensure that only one of --level, --max-level, or --min-level can be used at a time
	if (*level != "" && *maxLevel) || (*level != "" && *minLevel) || (*maxLevel && *minLevel) {
		fmt.Fprintln(os.Stderr, "Error: --level, --max-level, and --min-level cannot be used together")
		os.Exit(1)
	}

	// Ensure --until-level is only used with --max-level or --min-level
	if *untilLevel > 0 && !(*maxLevel || *minLevel) {
		fmt.Fprintln(os.Stderr, "Error: --until-level can only be used with --max-level or --min-level")
		os.Exit(1)
	}

	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var domains []string
	for scanner.Scan() {
		domain := scanner.Text()
		domains = append(domains, domain)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		return
	}

	lineCount := 0 // Keep track of the number of lines printed

	if *level != "" {
		// Handle --level: parse levels and filter based on them
		levels := parseLevels(*level)
		groupedDomains := groupDomainsByLevel(domains, levels)
		printGroupedDomainsInOrder(groupedDomains, levels, &lineCount, *untilCount)
	} else if *maxLevel {
		// Handle --max-level
		sortDomainsByDotsDescending(domains)
		for _, domain := range domains {
			dotCount := strings.Count(domain, ".")
			fmt.Println(domain)
			lineCount++
			if *untilCount > 0 && lineCount >= *untilCount {
				break
			}
			if *untilLevel > 0 && dotCount == *untilLevel {
				break
			}
			// Stop if we passed the highest level and --until-level doesn't exist
			if *untilLevel > 0 && dotCount < *untilLevel {
				break
			}
		}
	} else if *minLevel {
		// Handle --min-level
		sortDomainsByDotsAscending(domains)
		for _, domain := range domains {
			dotCount := strings.Count(domain, ".")
			fmt.Println(domain)
			lineCount++
			if *untilCount > 0 && lineCount >= *untilCount {
				break
			}
			if *untilLevel > 0 && dotCount == *untilLevel {
				break
			}
			// Stop if we passed the lowest level and --until-level doesn't exist
			if *untilLevel > 0 && dotCount > *untilLevel {
				break
			}
		}
	} else {
		// Default case: just print all domains
		for _, domain := range domains {
			fmt.Println(domain)
			lineCount++
			if *untilCount > 0 && lineCount >= *untilCount {
				break
			}
		}
	}
}

// parseLevels parses the --level flag and returns a slice of levels as integers
// It preserves the order in which levels are provided.
func parseLevels(levels string) []int {
	levelStrings := strings.Split(levels, ",")
	var result []int
	for _, levelStr := range levelStrings {
		levelStr = strings.TrimSpace(levelStr)
		level, err := strconv.Atoi(levelStr)
		if err == nil {
			result = append(result, level)
		} else {
			fmt.Fprintf(os.Stderr, "Invalid level: %v\n", levelStr)
			os.Exit(1)
		}
	}
	// Do not sort levels, keep the order they were provided in.
	return result
}

// groupDomainsByLevel groups domains by their dot levels
func groupDomainsByLevel(domains []string, levels []int) map[int][]string {
	grouped := make(map[int][]string)
	for _, domain := range domains {
		dotCount := strings.Count(domain, ".")
		for _, level := range levels {
			if dotCount == level {
				grouped[level] = append(grouped[level], domain)
				break
			}
		}
	}
	return grouped
}

// printGroupedDomainsInOrder prints the grouped domains in the order of the specified levels
// Stops after printing the specified number of lines if --until-count is used
func printGroupedDomainsInOrder(grouped map[int][]string, levels []int, lineCount *int, untilCount int) {
	for _, level := range levels {
		if domains, exists := grouped[level]; exists {
			for _, domain := range domains {
				fmt.Println(domain)
				*lineCount++
				if untilCount > 0 && *lineCount >= untilCount {
					return
				}
			}
		}
	}
}

// sortDomainsByDotsDescending sorts domains from the maximum to the minimum number of dots
func sortDomainsByDotsDescending(domains []string) {
	sort.Slice(domains, func(i, j int) bool {
		return strings.Count(domains[i], ".") > strings.Count(domains[j], ".")
	})
}

// sortDomainsByDotsAscending sorts domains from the minimum to the maximum number of dots
func sortDomainsByDotsAscending(domains []string) {
	sort.Slice(domains, func(i, j int) bool {
		return strings.Count(domains[i], ".") < strings.Count(domains[j], ".")
	})
}
