package main

import (
	"fmt"
)

// var statusCount int64 = 0

// // model for course
// type Course struct {
// 	ID     string  `json:"id"`
// 	Name   string  `json:"name"`
// 	Price  float64 `json:"price"`
// 	Author *Author `json:"author"`
// }
// type Author struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

// // fake db
// var courses []Course

// // middleware
// func (c *Course) IsEmpty() bool {
// 	return c.ID == "" || c.Name == ""
// }

// func serverHost(w http.ResponseWriter, r *http.Request) {
// 	clientIP, _, _ := net.SplitHostPort(r.RemoteAddr)
// 	fmt.Println("Sever status check", clientIP)
// 	statusCount++
// 	response := fmt.Sprintf(`<h1>Server is up now. statusCount: %d</h1>`, statusCount)
// 	w.Write([]byte(response))
// }

// func getAllCources(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Get all courses", r.Host)
// 	w.Header().Set("Content-Type", "application/json")
// 	if len(courses) == 0 {
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"message": "No courses found",
// 			"success": true,
// 			"data":    nil,
// 		})
// 		return
// 	}
// 	json.NewEncoder(w).Encode((courses))
// }
// func addACourse(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Add a course", r.Body)
// 	w.Header().Set("Content-Type", "application/json")
// 	if r.Body == nil {
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"message": "No body found",
// 			"success": false,
// 			"data":    nil,
// 		})
// 		return
// 	}
// 	var course Course
// 	err := json.NewDecoder(r.Body).Decode(&course)
// 	fmt.Println(course)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"message": err.Error(),
// 			"success": false,
// 			"data":    nil,
// 		})
// 		return

// 	}
// 	if course.IsEmpty() {
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"message": "Course is empty",
// 			"success": false,
// 			"data":    nil,
// 		})
// 		return
// 	}
// 	courses = append(courses, course)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"message": "Course added",
// 		"success": true,
// 		"data":    course,
// 	})

// }

// func updateCourse(w http.ResponseWriter, r *http.Request) {
// 	//fmt.Println("Update a course", ps)
// 	w.Header().Set("Content-Type", "application/json")
// 	if r.Body == nil {
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"message": "No body found",
// 			"success": false,
// 			"data":    nil,
// 		})
// 		return
// 	}
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	fmt.Println(id)
// 	var course Course
// 	err := json.NewDecoder(r.Body).Decode(&course)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"message": err.Error(),
// 			"success": false,
// 			"data":    nil,
// 		})
// 		return
// 	}
// 	if course.IsEmpty() {
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"message": "Course is empty",
// 			"success": false,
// 			"data":    nil,
// 		})
// 		return

// 	}
// 	for i, c := range courses {
// 		if c.ID == id {
// 			courses[i] = course

// 		}
// 	}
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"message": "Course updated",
// 		"success": true,
// 		"data":    course,
// 	})
// }
// func main() {
// 	fmt.Println("====hello Amit===")
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", serverHost).Methods("GET")
// 	r.HandleFunc("/getAllCources", getAllCources).Methods("GET")
// 	r.HandleFunc("/addCourse", addACourse).Methods("POST")
// 	r.HandleFunc("/updateCourse/{id}", updateCourse).Methods("PUT")

// 	err := http.ListenAndServe(":8000", r)
// 	if err != nil {
// 		panic(err)
// 	}

// }

func main() {
	fmt.Println("====hello Amit===")
	//go printName("Amit")
	// i := 10
	// for i < 12 {
	// 	if i == 11 && i < 10 {
	// 		fmt.Println("Amit", i)
	// 		panic("Limit Exceed")
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }
	// var age int64
	// fmt.Println("Entre your age : ")
	// fmt.Scan(&age)
	// switch age {
	// case 18:
	// case 21:
	// 	fmt.Println("You are a child")
	// case 20:
	// 	fmt.Println("You are an adult")
	// default:
	// 	fmt.Println("You are a senior citizen")

	// }

	// var tech = [...]string{"Java", "Node", "Go"}
	// var tech = [3]string{"Java", "Node", "Go"}
	// fmt.Println(tech[0])
	// tech[2] = "Null"
	// fmt.Println(tech)

	//tech := [2]float64{}
	//tech := make([]int, 2, 10)
	//tech := make([]int, 10, 10)
	//tech := [][]int{{1, 2, 3}, {4}}

	//fmt.Print(tech, cap(tech[0]), len(tech[0]))

	//----map
	// a := make(map[string]int64)
	// a["Amit"] = 1
	// a["Amit"] = 2
	// a["jdncjedcbnejcercr"] = 9223372036854775807
	// var name = "````AMit"
	// fmt.Println(name[0])
	// {
	// 	//var name = "Amif"
	// }
	// for i, v := range "am-```nma" {
	// 	fmt.Println(i, v)
	// }
	//fmt.Println(a["Amist"])
	// age := printName(2, 3)()
	// age := printName(2, 3)()
	// age := printName(2, 3)()
	cacheFun := checheTheData()
	fmt.Println(cacheFun(2, 3))
	fmt.Println(cacheFun(1, 3))
	fmt.Println(cacheFun(2, 3))
	fmt.Println(cacheFun(6, 3))
	var a int = 10
	check(a)
	check(a)

}
func checheTheData() func(int, int) int {
	var age = 10
	ma := make(map[int]int)
	return func(key, mul int) int {
		v, ok := ma[key]
		if ok {
			fmt.Println("Cached data")
			return v
		}
		ma[key] = mul * age
		return ma[key]
	}
}

func check(a int) {
	a++
	fmt.Println(a)
}
