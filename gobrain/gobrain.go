package gobrain

import (
	"os"
	"fmt"
	"flag"
	"strings"
	
)


func check_error(e error) {
    if e != nil {
        panic(e)
    }
}


func interpret_file(filename string) string{
	f, err := os.Create("info.log")
    check_error(err)

	defer f.Close()



	data, err := os.ReadFile(filename)
	check_error(err)
	content := strings.Replace(string(data), "\n", "", -1)
	memory := make([]byte, 30)

	var ptr, idx int
	var chr byte
	var result []byte
	var skip bool = false
	loop_start_slice := []int{}
	var skip_start int = 0
	var last_loop_start int = 255

	for idx < len(content) {
		chr = content[idx]
		if chr == 43 && !skip{
			memory[ptr]++
		} else if chr == 45 && !skip{
			memory[ptr]--
		} else if chr == 46 && !skip{
			result = append(result, memory[ptr])
		}	else if chr == 62 && !skip{
			ptr++
			if ptr >= len(memory) {
				memory = append(memory, 0)
			}
		} else if chr == 60 && !skip{
			ptr--
		} else if chr == 91 {
			if memory[ptr] == 0 && !skip {
				skip = true
				skip_start = idx
			}
			loop_start_slice = append(loop_start_slice, idx)
		} else if chr == 93 {
			last_loop_start = loop_start_slice[len(loop_start_slice)-1]
			loop_start_slice = loop_start_slice[0:len(loop_start_slice)-1]
			if memory[ptr] != 0 && !skip {
				idx = last_loop_start
				continue
			} else if last_loop_start == skip_start {
				
				skip = false
			}


		}


		/*
		string1 := fmt.Sprintf("%d %s %d %d %d %t\n", idx, string(chr), memory, ptr, memory[ptr], skip)
		
		_, err := f.WriteString(string1)
		check_error(err)
*/
		idx++
	}
	f.Sync()
	return string(result)
}

func main() {
	method := flag.String("m", "run", "run or build")
	filename := flag.String("f", "", "file to be interpreted/compiled")
	flag.Parse()
	fmt.Println("Method", *method)

	result := interpret_file(*filename)

	fmt.Println(result)

}