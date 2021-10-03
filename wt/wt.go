package wt

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func W(args ...interface{}) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	if len(args) == 1 {
		arg := args[0]
		if b, ok := arg.(bool); ok {
			if b {
				fmt.Fprintln(w, "Yes")
			} else {
				fmt.Fprintln(w, "No")
			}
			return
		}
		if v := reflect.ValueOf(arg); v.Kind() == reflect.Slice {
			for i := 0; i < v.Len(); i++ {
				if i > 0 {
					fmt.Fprint(w, " ")
				}
				fmt.Fprint(w, v.Index(i))
			}
			fmt.Fprintln(w)
			return
		}
	}
	fmt.Fprintln(w, args...)
}
