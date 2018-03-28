type Runtime struct {
	stdin     *bufio.Reader
	stdout    io.Writer
	stack     Stack
	commands  Program
	heap      map[int]int
	labels    map[int]int
	callstack []int
	index     int
}
