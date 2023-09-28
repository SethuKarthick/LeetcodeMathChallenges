func arrangeCoins(n int) int {
    // n = 5
    // row := 1
    // coins := 1
    // for ; coins <= n && n-coins > 0; row++{
    //     n = n-coins
    // }
    // return row
    // row := 0
    // for coins:=1; coins <= n && n- coins > 0; coins*=row{
    //     n = n-coins
    //     row+=1
        // coins = 1*row
    // }
    // return row-1
    idx := 0
    for true{
        if n < 0 {
            return idx -1
        }
        idx++
        n = n - idx 
    }
    return idx-1
}