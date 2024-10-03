package main

import (
	"fmt"
	"math"
	"sort"
)

type task map[string]float64

type Node struct {
	Char  string
	Prob  float64
	Left  *Node
	Right *Node
	Code  string
}

func buildHuffmanTree(probabilities map[string]float64) *Node {
	var nodes []*Node
	for char, prob := range probabilities {
		nodes = append(nodes, &Node{Char: char, Prob: prob})
	}
	for len(nodes) > 1 {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Prob < nodes[j].Prob
		})
		left := nodes[0]
		right := nodes[1]
		newNode := &Node{
			Prob:  left.Prob + right.Prob,
			Left:  left,
			Right: right,
		}
		nodes = append(nodes[2:], newNode)
	}
	return nodes[0]
}

func generateHuffmanCodes(root *Node, code string, codes map[string]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		root.Code = code
		codes[root.Char] = code
		return
	}
	generateHuffmanCodes(root.Left, code+"0", codes)
	generateHuffmanCodes(root.Right, code+"1", codes)
}

func calculateEntropy(probabilities map[string]float64) float64 {
	entropy := 0.0
	for _, prob := range probabilities {
		if prob > 0 {
			entropy += -prob * math.Log2(prob)
		}
	}
	return entropy
}

func calculateAverageLength(codes map[string]string, probabilities map[string]float64) float64 {
	averageLength := 0.0
	for char, code := range codes {
		averageLength += float64(len(code)) * probabilities[char]
	}
	return averageLength
}

func encodeText(stringToEncode string, huffmanCodes map[string]string) (encodedSequence string) {
	for _, char := range stringToEncode {
		encodedSequence += huffmanCodes[string(char)]
	}
	return
}

func main() {

	tasks := []map[string]task{}
	probabilities1 := map[string]float64{
		"P1": 0.35,
		"P2": 0.15,
		"P3": 0.06,
		"P4": 0.02,
		"P5": 0.03,
		"P6": 0.08,
		"P7": 0.02,
		"P8": 0.07,
		"P9": 0.22,
	}

	probabilities2_text := "ACBABAAAADDAADBAADAADBAAAAAADA"
	probabilities2 := map[string]float64{
		"A": 0.65,
		"B": 0.15,
		"C": 0.06,
		"D": 0.14,
	}

	tasks = append(tasks, map[string]task{"": probabilities1})
	tasks = append(tasks, map[string]task{probabilities2_text: probabilities2})

	for _, t := range tasks {
		println("---------------------NEW TASK (VARIANT 7)------------------------")
		for k, v := range t {
			printResult(v, k)
		}
		println("-----------------------------------------------------------------")
	}
}

func printResult(probabilities map[string]float64, encodeString string) {
	root := buildHuffmanTree(probabilities)
	codes := make(map[string]string)
	generateHuffmanCodes(root, "", codes)

	fmt.Println("Huffman Codes:")
	for char, code := range codes {
		fmt.Printf("%s: %s\n", char, code)
	}

	if encodeString != "" {
		fmt.Println("Encoded Text:", encodeText(encodeString, codes))
	}

	entropy := calculateEntropy(probabilities)
	fmt.Printf("Entropy: %.4f bit\n", calculateEntropy(probabilities))

	averageLength := calculateAverageLength(codes, probabilities)
	fmt.Printf("Average Length of the code: %.4f\n", calculateAverageLength(codes, probabilities))

	efficiency := entropy / averageLength
	fmt.Printf("Efficiency of the code: %.4f\n", efficiency)
}
