package utils

import (
  "strings"
)

func CreateTitle(count int, title string) string {
  const SPACE_BETWEEN = 2
  lineCount := (count - SPACE_BETWEEN - len(title))/2
  return CreateLine(lineCount, "=") +
          " " + strings.ToUpper(title) + 
          " " + CreateLine(lineCount, "=")
}