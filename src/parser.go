package main

import "strings"

type JClass struct {
    Name   string
    Method string
    Body   string
}

func ParseJavaLike(code string) JClass {
    lines := strings.Split(code, "\n")
    cls := JClass{}

    for _, l := range lines {
        l = strings.TrimSpace(l)
        if strings.HasPrefix(l, "class ") {
            cls.Name = strings.TrimSuffix(strings.TrimPrefix(l, "class "), " {")
        }
        if strings.Contains(l, "static void main") {
            cls.Method = "main"
        }
        if strings.Contains(l, "print(") {
            inside := l[strings.Index(l, "(")+1 : strings.LastIndex(l, ")")]
            cls.Body = inside
        }
    }

    return cls
}
