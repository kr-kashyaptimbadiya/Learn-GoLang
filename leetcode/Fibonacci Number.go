func fib(n int) int {
    // if n==0{
    //     return 0
    // }
    // if n==1{
    //     return 1
    // }
    // return fib(n-1) + fib(n-2)
    a := -1
    b := 1
    for i :=0 ; i < n+1 ; i++ {
        a , b = b , a + b
    }
    return b
}
