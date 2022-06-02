package solver

import (
	"fmt"
	"sync"
)

const ThreadCount = 2

var (
	ClassDirpath = map[int]string{
		ClassConvex:  "./convex/",
		ClassConcave: "./concave/",
		ClassOther:   "./other/",
	}
)

type Solver struct{}

func (s *Solver) Solve(path string) (err error) {
	filenames, err := ListFiles(path)
	if err != nil {
		fmt.Println("list files error: ", err)
		return err
	}

	if len(filenames) == 0 {
		return
	}

	threadFileLen := len(filenames) / ThreadCount
	wg := &sync.WaitGroup{}
	wg.Add(threadFileLen)
	for i := 0; i < ThreadCount; i++ {
		start := i * threadFileLen
		end := (i + 1) * threadFileLen

		if start > len(filenames) {
			continue
		}
		if end > len(filenames) {
			end = len(filenames)
		}

		go solve(wg, path, filenames[start:end])
	}

	wg.Wait()
	return
}

func solve(wg *sync.WaitGroup, dirName string, filenames []string) {
	defer wg.Done()

	for _, file := range filenames {
		filedata, err := ReadFile(dirName + file)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		pointOrder := ParseFigurePoints(filedata[0])
		points := []Point{}
		for _, data := range filedata[1:] {
			point := ParsePoint(data)
			points = append(points, point)
		}

		destPath := ClassDirpath[NewFigure(pointOrder, points).DetectClassType()]
		if err = MoveFile(dirName+file, destPath+file); err != nil {
			fmt.Println("move file failed, error: ", err)
			return
		}
	}
}
