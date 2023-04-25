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

	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":8080", nil)
}
