package Convert

import (
    "fmt"
    "strconv"
)

func ToBoolean(value interface{}) bool {
    switch t := value.(type) {
    case bool:
        return t
    }

    panic(fmt.Sprintf("ToBoolean %+v %T", value, value))
}

func ToInt(value interface{}) int{
    switch t := value.(type) {
    case string:
        i, err := strconv.Atoi(t)
        if err != nil {
            panic(err)
        }

        return i

    case int:
        return t
    case bool:
        if t {
            return 1
        }

        return 0
    }

    panic(fmt.Sprintf("ToInt %+v %T", value, value))
}

func ToString(value interface{}) string {
    switch t := value.(type) {
    case string:
        return t
    case int:
        return strconv.Itoa(t)
    case bool:
        return strconv.FormatBool(t)
    }

    panic(fmt.Sprintf("ToString %+v %T", value, value))
}
