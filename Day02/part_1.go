package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1Answer() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the input file", err.Error())
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		// Reading each line from the input file
		line := scanner.Text()
		lineInfo := strings.Split(line, ":")
		gameInfo := lineInfo[0]
		cubesInfo := lineInfo[1]
		setInfo := strings.Split(cubesInfo, ";")
		isValid := true
		for _, info := range setInfo {
			cubes := strings.Split(info, ",")
			for _, cube := range cubes {
				cube = strings.TrimSpace(cube)
				cubeInfo := strings.Split(cube, " ")
				numOfCube := cubeInfo[0]
				colOfCube := cubeInfo[1]
				numOfCubeInt, err := strconv.Atoi(numOfCube)
				if err != nil {
					fmt.Println("Error converting numOfCube string to int", err.Error())
					os.Exit(1)
				}
				if (colOfCube == "blue" && numOfCubeInt > 14) ||
					(colOfCube == "green" && numOfCubeInt > 13) ||
					(colOfCube == "red" && numOfCubeInt > 12) {
					isValid = false
				}

			}

		}
		if !isValid {
			continue
		}
		gameId := strings.Split(gameInfo, " ")[1]
		gameIdInt, err := strconv.Atoi(gameId)
		if err != nil {
			fmt.Println("Error converting string to int", err.Error())
			os.Exit(1)
		}
		result += gameIdInt
	}
	fmt.Println("Result --> ", result)
}
