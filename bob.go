package bob

import "strings"

type Bot struct {
  name string
}

func (bob Bot) Response (input string) (answer string) {
  message := strings.Trim(input, " ")
  if isEmpty(message) {
    answer = "Fine"
  } else if isQuestion(message) {
    answer = "Sure"
  } else if isScream(message) {
    answer = "Whoa, chill out !"
  } else {
    answer = "Whatever"
  }
  return answer
}

func isEmpty (message string) bool {
  return message == ""
}

func isQuestion (message string) bool {
  return strings.HasSuffix(message, "?")
}

func isScream (message string) bool {
  return message == strings.ToUpper(message) && message != strings.ToLower(message)
}
  