package main
import (
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
	"math"
	"Demo2Planner/dijkstra"
	"fmt"
)

type Point struct {
	id uint
	x uint
	y  uint
	adjacent_points []uint
}



func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {


	points := readInputsFromFile("4by4.map")
	graph:=dijkstra.NewGraph()
	for k ,_ := range points {
		graph.AddVertex(int(k))
	}

	for index , point := range points {
		for _, a_point := range point.adjacent_points {
			distance := int64(point.Distance(points[a_point]))
			graph.AddArc(int(index), int(a_point), distance)
		}
	}

	best, err := graph.Shortest(18,26)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)


}

func readInputsFromFile(filename string) (map[uint]Point) {
	points := make(map[uint]Point)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line :=strings.Split(scanner.Text()," ")
		index, _ := strconv.Atoi(line[0])
		x, y := splitCommaSepNumbers(line[1])
		adjacent_points := []uint{}
		for _, point := range line[2:] {
			temp, _ := strconv.Atoi(point)
			adjacent_points = append(adjacent_points, uint(temp))
		}
		points[uint(index)] = Point{x: x, y: y, adjacent_points:adjacent_points}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return points
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


//// TODO
//func findClosestEdgeToPoint(randomPoint Point, edges []Edge) (Point, Edge) {
//	var closestEdge Edge
//	var closestPointOnEdge Point
//	return closestPointOnEdge, closestEdge
//}

func (p Point) Distance(p2 Point) float64 {
	first := math.Pow(float64(p2.x)-float64(p.x), 2)
	second := math.Pow(float64(p2.y)-float64(p.y), 2)
	return math.Sqrt(first + second)
}