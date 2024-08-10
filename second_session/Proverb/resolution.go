package proverb

import (
	"fmt"
	"strings"
)

func main() {
	items := []string{"میخ", "کفش", "اسب", "سوار", "پیام", "نبرد", "پادشاهی"}

	proverb := CreateProverb(items)

	fmt.Println(proverb)
}

func CreateProverb(items []string) string {
	var result []string

	for i := 0; i < len(items)-1; i++ {
		line := fmt.Sprintf("به دلیل نداشتن %s، %s گم شد.", items[i], items[i+1])
		result = append(result, line)
	}

	lastLine := fmt.Sprintf("و همه اینها به خاطر نیاز به یک %s.", items[0])
	result = append(result, lastLine)

	return strings.Join(result, "\n")
}
