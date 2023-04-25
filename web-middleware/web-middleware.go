package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Course struct {
	ClassID    int    `json: "id"`
	Level      string `json: "level"`
	Subject    string `json: "subject"`
	Department string `json: "department"`
}

var courseList []Course

func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJson, err := json.Marshal(courseList)
	switch r.Method {
	case "GET":
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJson)
	case http.MethodPost:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Println(r.Body)
		// w.Write()
	}
}
func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler middleware start")

		handler.ServeHTTP(w, r)
		fmt.Println("middleware finished!")
	})
}
func main() {
	courseJson := `[
			  {
				"classid": 2,
				"level": "101",
				"subject": "Algebra 1",
				"department": "Mathematics"
			  },
			  {
				"classid": 3,
				"level": "102",
				"subject": "Algebra 2",
				"department": "Mathematics",
				"prereqs": ["Algebra 1"]
			  },
			  {
				"classid": 4,
				"level": "110",
				"subject": "Geometry",
				"department": "Mathematics",
				"location": {
					"building": "AM1",
					"room": 111
				}
			  },
			  {
				"classid": 5,
				"level": "111",
				"subject": "Trigonometry",
				"department": "Mathematics",
				"prereqs": ["Algebra 1", "Geometry"],
				"comments": "2nd semester only",
				"location": {
					"building": "AM1",
					"room": 104
				}
			  }
	]`
	err := json.Unmarshal([]byte(courseJson), &courseList)
	if err != nil {
		log.Fatal(err)
	}

	// http.HandleFunc("/course", courseHandler)
	// add middleware
	courseHandlerFunc := http.HandlerFunc(courseHandler)
	http.Handle("/course", middlewareHandler(courseHandlerFunc))
	http.ListenAndServe(":8080", nil)
}
