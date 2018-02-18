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

    dp := make([][]float32, n + 1)
    pr := make([][]int, n + 1)

    for i := 0; i <= n; i++ {
        dp[i] = make([]float32, T + 1)
        pr[i] = make([]int, T + 1)
    }
         
    for i := 0; i <= n; i++ {
        for j := 0; j <= T; j++ {
            dp[i][j] = 1.0
            pr[i][j] = -2 
        }
    }
    pr[0][0] = -1    

    for item := 0; item < n; item++ {
        for j := 0; j + t[item] <= T; j++ {
            i := j + t[item]
            if pr[item][j] != -2 {            
                if dp[item + 1][i] >= dp[item][j] * p[item] {
                   dp[item + 1][i] = dp[item][j] * p[item]
                   pr[item + 1][i] = item
                }  
           }                         
        }
        for i := T; i >= 0; i-- {
            if pr[item][i] != -2 {
                if dp[item + 1][i] >= dp[item][i] {
                   dp[item + 1][i] = dp[item][i]
                   pr[item + 1][i] = -1        
                }
            }
        }
    }                                         

    var best float32
    var pos int
    best = 1           
    pos = 0

    for i := 0; i <= T; i++ {
        if best > dp[n][i] {
            best = dp[n][i]
            pos = i
        }                       
    }

    answer := make([]int, 0, n)
    for i := n; i > 0; i-- {
        if pr[i][pos] != -1 {
            answer = append(answer, pr[i][pos])
            pos -= t[pr[i][pos]]
        }
    }
    fmt.Printf("%f\n", 1 - best);                                            
    fmt.Printf("%d\n", len(answer));
    for i := 0; i < len(answer); i++ {
        fmt.Printf("%d ", answer[i] + 1)
    }
}                      
