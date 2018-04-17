package main
import (
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
	"fmt"
)

type Point struct {
	x uint
	y  uint
}
type Edge struct {
	start Point
	end  Point
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {


	edges := readInputsFromFile("4by4.map")
	fmt.Printf("%d %d %d %d\n", edges[0].start.x, edges[0].start.y, edges[0].end.x, edges[0].end.y)

}

func readInputsFromFile(filename string) ([]Edge) {
	var points []Point
	var edges []Edge

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numOfPoints, numOfEdges := splitCommaSepNumbers(scanner.Text())


	for i := uint(0); i < numOfPoints; i++ {
		scanner.Scan()
		x, y := splitCommaSepNumbers(scanner.Text())
		newPoint := Point{x: x, y: y}
		points = append(points, newPoint)
	}

	if len(points) != int(numOfPoints) {
		log.Fatalf("given number of points %d not equal to observed number of points %d", numOfPoints, len(points))
	}
	for i := uint(0); i < numOfEdges; i++ {
		scanner.Scan()
		startPointIndex, endPointIndex := splitCommaSepNumbers(scanner.Text())
		if startPointIndex >= numOfPoints || endPointIndex >= numOfPoints {
			log.Fatalf("%d or %d greater than or equal to length of points %d", startPointIndex, endPointIndex, len(points))
		} else {
			newEdge := Edge{start: points[startPointIndex], end: points[endPointIndex]}
			edges = append(edges, newEdge)
		}

	}
	if len(edges) != int(numOfEdges) {
		log.Fatalf("given number of edges %d not equal to observed number of points %d", numOfEdges, len(edges))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return edges
}

func splitCommaSepNumbers(line string) (uint, uint) {
	numbers := strings.Split(line,",")
	if len(numbers) != 2 {
		log.Fatal("Line does not have format <number1>, <number2> ")
	}
	numberOne, err := strconv.Atoi(numbers[0])
	if err != nil {
		log.Fatal("Line does not have format <number1>, <number2> ")
	}
	numberTwo, err := strconv.Atoi(numbers[1])
	if err != nil {
		log.Fatal("Line does not have format <number1>, <number2> ")
	}
	return uint(numberOne), uint(numberTwo)
}

// for Michael Erberich
func planPath(startEdgeIndex uint, endEdgeIndex uint, edges [ ]Edge)  ([]Edge) {
	var path []Edge
	return path
}

// for Austin
func findClosestEdgeToPoint(randomPoint Point, edges Edge[]) (Point, Edge) {
	var closestEdge Edge
	var closestPointOnEdge Point
	return closestPointOnEdge, closestEdge
}
