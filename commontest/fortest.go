package main

func length(s string) int {
	println("call lengh.")
	return len(s)
}

func fortest(){
	s := "abcd"
	for i,n := 0,length(s);i<n;i++{
		println(i,s[i])
	}
}