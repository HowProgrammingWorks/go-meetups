// errMethod, passUp, wrapErr ...  are just constants
func TheGreatFunction(errMethod int) error {
	if errMethod == ignoring {
		ioutil.WriteFile("test.txt", []byte("Hello, there"), 0644)
		return nil
	}
	err := ioutil.WriteFile("test.txt", []byte("Hello, there"), 0644)
	switch errMethod {
	case passUp:
		return err
	case wrapErr:
		return fmt.Errorf("ioutil.WriteFile(%q): %v", "test.txt", err)
	case fatalErr:
		log.Fatalf("ioutil.WriteFile(%q): %v", "test.txt", err)
	case panicErr:
		panic(fmt.Sprintf("ioutil.WriteFile(%q): %v", "test.txt", err))
	case checkingErr:
		if err == io.EOF {
			logic()
			return nil
		}
	}
	return nil
}
