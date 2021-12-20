package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
)

// damn I'm just learning now I should be using camel case for go 
type bitString string

type packet struct {
	bstr bitString // whole packet string
	//	len int // packet len 
	ver int // version
	tid int // type id
	// for literal values
	val int // the value
	// for operator packets
	ltid byte // length type id. '1' or '0'
	len int // the length of either bits of subpackets, or num subpackets
	//	mode bool // mode of the length type id. false for length, true for num packets
	subp []packet // subpackets
}
// damn...

func (p packet) print(prefix string) {
	if len(p.bstr) > 30 {
		fmt.Println(prefix, "bitstring:", p.bstr[:25], "...", p.bstr[len(p.bstr)-5:])
	} else {
		fmt.Println(prefix, "bitstring:", p.bstr)
	}

	if p.tid == 4 {
		fmt.Println(prefix, "literal value")
		fmt.Println(prefix, "version:", p.ver, "type id:", p.tid)
		fmt.Println(prefix, "value:", p.val)
	} else {
		fmt.Println(prefix, "operator packet")
		fmt.Println(prefix, "version:", p.ver, "type id:", p.tid)
		if p.ltid == '0' {
			fmt.Println(prefix, "total subpacket bitlength:", p.len)
		} else {
			fmt.Println(prefix, "number of subpackets:", p.len)
		}
		for _, p := range p.subp {
			p.print(prefix + "    ")
		}
	}
}

// parses the literal value
// returns the value, offset
func (bs bitString) parseLiteralValue() (int, int) {
	idx := 0
	val := 0
	done := false
	for ! done {
		for i := 1 ; i < 5 ; i ++ {
			val += int(bs[idx + i] - '0')
			val = val << 1
		}
		if bs[idx] == '0' {
			done = true
		}
		idx += 5
	}
	return val>>1, idx
}

func (bs bitString) parsePacket() packet {
	var out packet

	out.ver = (bs[:3]).toInt()
	out.tid = (bs[3:6]).toInt()
	// fmt.Println("parsing:", bs)
	//	fmt.Println("got header! ver:", out.ver, "tid:", out.tid)

	if out.tid == 4 {
		//	fmt.Println("got literal value!")
		val, len := (bs[6:]).parseLiteralValue()
		// fmt.Println("val:", val, "len:", len+6)
		out.val = val
		out.bstr = bs[:6+len]
	} else {
		out.ltid = bs[6]
		if out.ltid == '0' {
			out.len = (bs[7:22]).toInt()
			
			//	fmt.Println("got operator! ltid:", out.ltid, "len:", out.len)
			off := 0 // end of length header, start of packets
			for off != out.len {
				//	fmt.Println("off:",off)
				sub := (bs[22+off:]).parsePacket()
				off += len(sub.bstr)
				out.subp = append(out.subp, sub)
			}
			out.bstr = bs[:22+off]
			return out
		} else {
			out.len = (bs[7:18]).toInt()
			//	fmt.Println("got operator! ltid:", out.ltid, "len:", out.len)
			off := 0
			for i := 0 ; i < out.len ; i ++ {
				sub := (bs[18+off:]).parsePacket()
				off += len(sub.bstr)
				out.subp = append(out.subp, sub)
			}
			out.bstr = bs[:18+off]
			return out
		}
	}
	
	return out
}


func (bs bitString) toInt() int {
	var out int
	for _, c := range bs {
		out += int(c - '0')
		out = out << 1
	}
	return out >> 1
}

// this would literally be easier with a bit switch
// statement with 16 cases
func fromHex(hex string) bitString {
	var out bitString
	for _, c := range hex {
		var n int
		if c >= '0' && c <= '9' {
			n = int(c - '0')
		} else {
			n = int(c - 'A') + 10
		}
		for i := 0 ; i < 4 ; i ++ {
			if n & 8 == 8 {
				out += "1"
			} else {
				out += "0"
			}
			n = n << 1
		}
	}
	// it would literally be shorter
	return out 
}

func read_input() bitString {
	var output string
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = scanner.Text()
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fromHex(output)
}

func (p packet) calculate() int {	
	switch p.tid {
	case 0:
		sum := 0
		for _, sub := range p.subp {
			sum += sub.calculate()
		}
		return sum
	case 1:
		prod := 1
		for _, sub := range p.subp {
			prod *= sub.calculate()
		}
		return prod
	case 2:
		min := -1
		for _, sub := range p.subp {
			calc := sub.calculate()
			if min == -1 || calc < min {
				min = calc
			}
		}
		return min
	case 3:
		max := 0
		for _, sub := range p.subp {
			calc := sub.calculate()
			if calc > max {
				max = calc
			}
		}
		return max
	case 4:
		return p.val
	case 5:
		c1 := p.subp[0].calculate()
		c2 := p.subp[1].calculate()
		if c1 > c2 {
			return 1
		} else {
			return 0
		}
	case 6:
		c1 := p.subp[0].calculate()
		c2 := p.subp[1].calculate()
		if c1 < c2 {
			return 1
		} else {
			return 0
		}
	case 7:
		c1 := p.subp[0].calculate()
		c2 := p.subp[1].calculate()
		if c1 == c2 {
			return 1
		} else {
			return 0
		}
	}
	return -1
}

func solve_part2(p packet) int {
	return p.calculate()
}

func (p packet) sumVersions() int {
	sum := p.ver
	for _, sub := range p.subp {
		sum += sub.sumVersions()
	}
	return sum
}

func solve_part1(p packet) int {
	return p.sumVersions()
}

func main() {
	hex_packet := read_input()
	p := hex_packet.parsePacket()
	p.print("")
	//	fmt.Println(hex_packet)
	fmt.Println("part1: ", solve_part1(p))
	fmt.Println("part2: ", solve_part2(p))
}
