package outburst

import (
	"bufio"
	"fmt"
	"os"
)

func (l Line) writelog(logText ...string) {
	f, err := os.OpenFile(l.File, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Could not create to logfile Error: ", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, v := range logText {
		_, err := fmt.Fprintln(w, v)
		if err != nil {
			fmt.Println("Could not write to logfile Error: ", err)
		}
	}
	w.Flush()
}
