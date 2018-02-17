package main

import "fmt"

func main() {
    var n, T int
    fmt.Scanf("%d %d\n", &n, &T)
    //fmt.Printf("%d", n)
    t := make([]int, n)
    p := make([]float32, n)
    for i := 0; i < n; i++ {     
        fmt.Scanf("%d", &t[i]);       
    }
    fmt.Scanf("\n");                              
    for i := 0; i < n; i++ {     
        fmt.Scanf("%f", &p[i]);  
              p[i] = 1 - p[i]
    }               

    dp := make([]float32, T + 1)
    pr := make([]int, T + 1)

    for i := 0; i <= T; i++ {
        dp[i] = 1
        pr[i] = -2
    }
    pr[0] = -1        

    for item := 0; item < n; item++ {
        for i := T; i - t[item] >= 0; i-- {
            j := i - t[item]
            if pr[j] != -2 {
                if dp[j] * p[item] < dp[i] {
                   dp[i] = dp[j] * p[item]
                    pr[i] = item
                }  
           }                         
        }
    }

    var best float32
    var pos int
    best = 1           
    pos = 0

    for i := 0; i <= T; i++ {
        if best > dp[i] {
            best = dp[i]
            pos = i
        }                       
    }

    answer := make([]int, 0, n)
    for ;pos > 0; pos = pos - t[pr[pos]] {
        answer = append(answer, pr[pos])
    }

    fmt.Printf("%d\n", len(answer));
    for i := 0; i < len(answer); i++ {
        fmt.Printf("%d ", answer[i] + 1)
    }
}