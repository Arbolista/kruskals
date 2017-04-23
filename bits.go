package kruskals

// Sets the bit at pos in the integer n.
func setBit(n int, pos int) int {
    n |= (1 << uint(pos))
    return n
}

// Clears the bit at pos in n.
func clearBit(n int, pos int) int {
    mask := ^(1 << uint(pos))
    n &= mask
    return n
}

func hasBit(n int, pos int) bool {
    val := n & (1 << uint(pos))
    return (val > 0)
}

func FlipBit(n int, pos int) int {
  if hasBit(n, pos) {
    return clearBit(n, pos)
  }
  return setBit(n, pos)
}
