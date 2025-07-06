package main

import (
    "fmt"
    "strings"
    "log"
)

type ValidationError struct {
    Field string
    Value interface{}
    Msg   string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for field '%s' with value '%v': %s", 
        e.Field, e.Value, e.Msg)
}

func validateGmail(gmail string) error {
    if !strings.Contains(gmail, "@") {
        return ValidationError{
            Field: "gmail",
            Value: gmail,
            Msg:   "must contain @ symbol",
        }
    }
    return nil
}

func main() {
    err := validateGmail("test")
    if err != nil {
        log.Println(err)
    }
}
