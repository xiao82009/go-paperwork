
package main
 
import (
	bufio
	fmt
	os
)
 
type LnFile struct {
	Count     int
	Filenames []string
}
 
func main() {
	counts = make(map[string]LnFile)
	files = os.Args[1]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg = range files {
			f, err = os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, dup2%vn, err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n = range counts {
		if n.Count  1 {
			fmt.Printf(%d %vn%sn, n.Count, n.Filenames, line)
		}
	}
}
 
func countLines(f os.File, counts map[string]LnFile) {
	input = bufio.NewScanner(f)
	for input.Scan() {
		key = input.Text()
		_, ok = counts[key]
		if ok {
			counts[key].Count++
			counts[key].Filenames = append(counts[key].Filenames, f.Name())
		} else {
			counts[key] = new(LnFile)
			counts[key].Count = 1
			counts[key].Filenames = append(counts[key].Filenames, f.Name())
		}
	}
}
--------------------- 
作者：金戋 
来源：CSDN 
原文：httpsblog.csdn.netq1576962841articledetails85162346 
版权声明：本文为博主原创文章，转载请附上博文链接！s