package bob

import "strings"

func Bob (message string) (reponse string) {
    message = strings.Trim(message, " ")

    if  message == ""{
      reponse = "Fine"
    } else if strings.HasSuffix(message, "?"){
      reponse = "Sure"
      } else if message == strings.ToUpper(message)&& message != strings.ToLower(message) {
        reponse = "Whoa, chill out!"
      }else { 
        reponse = "Whatever"
}
  return reponse
}
  