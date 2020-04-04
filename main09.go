package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type route struct {
	from     string
	to       string
	distance int
}

type path struct {
	locations []string
	distance  int
}

type byLocations []path

func (a byLocations) Len() int { return len(a) }
func (a byLocations) Less(i, j int) bool {
	return strings.Join(a[i].locations, ":") < strings.Join(a[j].locations, ":")
}
func (a byLocations) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type byDistance []path

func (a byDistance) Len() int { return len(a) }
func (a byDistance) Less(i, j int) bool {
	return a[i].distance < a[j].distance
}
func (a byDistance) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func main09() {
	f, err := os.Open("input-09.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	inputRoutes := []route{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		newRoute, err := parseInput09(s.Text())
		if err != nil {
			log.Fatal(err)
		}
		inputRoutes = append(inputRoutes, newRoute)
	}

	paths := getPaths09(inputRoutes)

	sort.Sort(byDistance(paths))
	shortestPath := paths[0]
	log.Printf("path: %v distance: %d\n", shortestPath.locations, shortestPath.distance)

	sort.Sort(sort.Reverse(byDistance(paths)))
	longestPath := paths[0]
	log.Printf("path: %v distance: %d\n", longestPath.locations, longestPath.distance)
}

func parseInput09(input string) (route, error) {
	from, to := "", ""
	distance := 0

	count, err := fmt.Sscanf(input, "%s to %s = %d", &from, &to, &distance)
	if err != nil || count != 3 {
		log.Fatalf("Error scanning %s.", input)
		return route{}, err
	}
	return route{from, to, distance}, nil
}

func reverseRoute09(input route) route {
	return route{from: input.to, to: input.from, distance: input.distance}
}

func lookupLocation09(locations []route, location string) []route {
	result := []route{}
	for _, currentRoute := range locations {
		if location == currentRoute.from {
			result = append(result, currentRoute)
		}
	}
	return result
}

func process09(locations []route, currentPath path) []path {
	result := []path{}
	place := currentPath.locations[len(currentPath.locations)-1]
	nextPlaces := lookupLocation09(locations, place)
	for _, nextPlace := range nextPlaces {
		if !stringSliceContains(currentPath.locations, nextPlace.to) {
			newPath := path{distance: currentPath.distance + nextPlace.distance}
			newPath.locations = append([]string(nil), currentPath.locations...)
			newPath.locations = append(newPath.locations, nextPlace.to)
			result = append(result, newPath)
		}
	}
	return result
}

func stringSliceContains(values []string, lookup string) bool {
	for _, value := range values {
		if value == lookup {
			return true
		}
	}
	return false
}

func getPaths09(inputRoutes []route) []path {
	routes := inputRoutes
	for _, rt := range inputRoutes {
		routes = append(routes, reverseRoute09(rt))
	}

	placeCount := getPlaceCount09(inputRoutes)

	result := []path{}
	stack := getStartPaths09(routes)

	for len(stack) > 0 {
		index := len(stack) - 1
		poppedPath := stack[index]
		stack = stack[:index]

		if len(poppedPath.locations) == placeCount {
			result = append(result, poppedPath)
		}

		newPaths := process09(routes, poppedPath)
		stack = append(stack, newPaths...)
	}

	return result
}

func getStartPaths09(routes []route) []path {
	result := []path{}
	for _, currentRoute := range routes {
		result = append(result, path{locations: []string{currentRoute.from, currentRoute.to}, distance: currentRoute.distance})
	}
	return result
}

func getPlaceCount09(routes []route) int {
	places := []string{}

	for _, rt := range routes {
		if !stringSliceContains(places, rt.from) {
			places = append(places, rt.from)
		}
		if !stringSliceContains(places, rt.to) {
			places = append(places, rt.to)
		}
	}

	return len(places)
}
