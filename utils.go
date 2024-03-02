package main

func validGET(args []string) bool {
	return len(args) == 2 && args[0] == GET
}

func validSET(args []string) bool {
	return len(args) == 3 && args[0] == SET
}

func validVIEW(args []string) bool {
	return args[0] == "view" && len(args) == 1
}

func validArgs(args []string) bool {
	return validGET(args) || validSET(args) || validVIEW(args)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
