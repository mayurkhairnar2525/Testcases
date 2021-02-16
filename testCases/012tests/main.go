package main


func Ints (Vs ... int ) int {
	return ints (Vs)
	//return Ints(Vs)
}

func ints (Vs [] int) int {
	if len (Vs) ==0{
		return 0
	}
	return ints(Vs [1:]) +Vs[0]
}