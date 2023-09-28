func isPerfectSquare(num int) bool {
    i := 0
    for i*i <num {i++}
    if i*i == num {return true}
    return false
}