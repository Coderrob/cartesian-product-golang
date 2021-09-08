package main

func combineSet(setA [][]string, setB []string) [][]string{
	var result [][]string
	if len(setA) == 1 {
		temp := setA[0]
		setA = make([][]string, len(temp))
		for index := 0; index < len(temp); index++ {
			setA[index] = []string{temp[index]}
		}
	}
	for indexA := 0; indexA < len(setA); indexA++ {
		for indexB := 0; indexB < len(setB); indexB++ {
			var temp []string
			temp = append(temp, setA[indexA]...)
			temp = append(temp, setB[indexB])
			result = append(result, temp)
			temp = nil
		}
	}
	return result
}

func Cartesian(list [][]string) [][]string {
	count := len(list)
	var temp [][]string
	temp = append(temp, list[0])
	for index := 1; index < count; index++ {
		temp = combineSet(temp, list[index])
	}
	return temp
}