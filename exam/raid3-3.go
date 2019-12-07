package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
)

func getDemention(a []byte, x *int, y *int) {
    count := 0
    for _, v := range a {
        if v == '\n' {
            *y++
            *x = count
        } else if *x == 0 {
            count++
        }
    }
}

func itoa(a int) string {
    digit := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}	sort.Sort(sort.StringSlice(arrOfRaids))
    temp := []string{}
    res := ""
    for a > 0 {
        temp = append(temp, digit[a%10])
        a /= 10
    }
    for i := len(temp) - 1; i >= 0; i-- {
        res += temp[i]
    }
    return res
}

func main() {
    arg := bufio.NewReader(os.Stdin)
    var input []byte
    for {
        line, err := arg.ReadBytes('\n')
        if err == nil {
            input = append(input, line...)
        } else {
            break
        }	sort.Sort(sort.StringSlice(arrOfRaids))
    }

    var x, y int
    getDemention(input, &x, &y)

    if x == 0 || y == 0 {
        fmt.Println("Not a Raid function")
        return
    }
    command := []string{"raid1a", "raid1b", "raid1c", "raid1d", "raid1e"}

    isStarted := false
    flag := false

    for _, v := range command {
        cmd := exec.Command("./"+v, itoa(x), itoa(y))
        stdin, _ := cmd.StdinPipe()

        go func() {
            defer stdin.Close()
        }()

        out, _ := cmd.CombinedOutput()
        match := true
        for i := 0; i < len(out); i++ {
            if input[i] != out[i] {
                match = false
            }
        }

        if isStarted && match {
            fmt.Printf(" %s ", "||")
        }
        if match {
            fmt.Printf("[%s] [%d] [%d]", v, x, y)
            isStarted = true
            flag = true
        }
    }
    if flag == false {
        fmt.Println("Not a Raid function")
        return
    }

	fmt.Println()
	
	func main() {

		arg := os.Args[1:]
	
		v := []byte(arg[0])  //3 51
		v2 := []byte(arg[1]) //6
	
		x := int(v[0] - 48) //3
		y := int(v2[0] - 48)
		Raid1a(x, y)
	}
	