package main

import (
	"bufio"
)

type Buffer struct {
	cmd    []byte
	param  []bool
	cindex int
}

const (
	S2_COMMAND State = iota
	S2_PARAM
	S2_LABEL
)

type Parser2 struct {
	Buffer
	state
	Program
}

func (p *Parser2) AddCommand(cmd Command) {
	p.Program = append(p.Program, cmd)
	p.cindex = 0
}

func (p *parser2) parseCommand(b byte) {
	switch b {
	case S:
		p.cmd[p.cindex] = 0x53
	case T:
		p.cmd[p.cindex] = 0x54
	case L:
		p.cmd[p.cindex] = 0x4c
	}
	p.cindex++

	switch string(p.cmd) {
	case "SS":
		p.AddCommand(&CommandStackPush{})
		p.state = S2_PARAM
	}
}

func (p *Parser2) Parse(r io.Reader) {
	reader := bufio.NewReader(r)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			return
		}
		switch p.state {
		case S2_COMMAND:
			p.parseCommand(b)
		}
	}
}
