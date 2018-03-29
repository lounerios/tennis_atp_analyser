package utils

import (
  "strconv"
)

func CheckErr(e error) {
    if (e != nil) {
        panic(e)
    }
}


func GetNumber(s string) int {
    num, err := strconv.Atoi(s)

    if (err != nil) {
        return 0
     }

    return num
}
