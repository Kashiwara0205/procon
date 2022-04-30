package p34

func Slove(a []int, sum int) bool{
	n := len(a)

	if dfs(a, n, sum){
		return true
	}

	return false
}

func dfs(a []int,  n int, sum int) bool{

	if n == 0{
		return sum == 0
	}

	if sum >= a[n-1] && dfs(a, n-1, sum - a[n-1]) {
		return true
	}

	if dfs(a, n-1, sum){
		return true
	}

	return false
}

