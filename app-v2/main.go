package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"

	_ "time/tzdata"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func main() {
	tracer.Start()
	defer tracer.Stop()

	// 環境変数からアプリバージョンを取得
	appVersion := os.Getenv("APP_VERSION")

	// profiling 設定
	err := profiler.Start(
		profiler.WithService("app"),
		profiler.WithEnv("prod"),
		profiler.WithVersion(appVersion),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

	// Http server
	mux := httptrace.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// リクエストのボディを取得します
		w.Write([]byte("Hello World!"))

		// ロジック
		count := calcTargetLogic(appVersion)
		fmt.Println("count: ", count)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func calcTargetLogic(appVersion string) (total int) {
	dummyData, err := read("./data/input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	total = count(dummyData, appVersion)

	return total
}

func read(filename string) ([]int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var dummyData []int
	for _, char := range string(content) {
		value, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, fmt.Errorf("error converting character to int: %v", err)
		}
		if value != 0 && value != 1 {
			return nil, fmt.Errorf("invalid value in file: %d", value)
		}
		dummyData = append(dummyData, value)
	}

	return dummyData, nil
}

func count(dummyData []int, appVersion string) (total int) {
	sort.Ints(dummyData)
	index := sort.SearchInts(dummyData, 1)
	fmt.Println("index: ", index)

	return len(dummyData) - index
}
