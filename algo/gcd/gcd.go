package main

func GCD(m, n int) int {
	r := m % n
	for r != 0 {
		m, n = n, r
		r = m % n
	}
	return n
}

func gcd(m, n int) int {
	for n != 0 {
		m, n = n, m%n
	}
	return m
}
