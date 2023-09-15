package date

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type date string

const (
	DD_MM_YY        date = "02-01-2006"
	DATE_MONTH      date = "02 Jan"
	DATE_MONTH_FULL date = "02 January"
)

// GetInputDate reads cli data and returns response in given
// date format
func GetInputDate(format date) (time.Time, error) {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		return time.Time{}, err
	}

	return time.Parse(string(format), input)

}

// DateDiffrence returns diffrence between two dates
func DateDiffrence(t1, t2 time.Time) string {
	duration := t2.Sub(t1)
	result := ""
	years := int(duration.Hours() / 24 / 365.25)
	days := int(duration.Hours()/24) % 365
	minutes := int(duration.Minutes()) % 60
	if years > 0 {
		result = fmt.Sprintf("%d year", years)
		if years > 1 {
			result += "s"
		}
	}
	if days > 0 {
		if len(result) > 0 {
			result += " "
		}
		result += fmt.Sprintf("%d day", days)
		if days > 1 {
			result += "s"
		}
	}
	if minutes > 0 {
		if len(result) > 0 {
			result += " "
		}
		result += fmt.Sprintf("%d minute", minutes)
		if minutes > 1 {
			result += "s"
		}
	}
	return result

}
