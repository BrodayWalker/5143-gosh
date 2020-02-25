ning
func History(args []string) {
	for i := range commandList {
		fmt.Println(i, commandList[i])