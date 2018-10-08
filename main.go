package main

import (
	"net/http"
  "strconv"
  "fmt"
  "flag"
  "os"
  

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// UserList main storage for users
type UserList []*User

// User ...
type User struct {
	ID       int       `json:"id" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Workouts []Workout `json:"workouts"`
}

// Workout ...
type Workout struct {
	ID         int        `json:"id" binding:"required"`
	UserID     int        `json:"userId" binding:"required"`
	Activities []Activity `json:"activities"`
}

// Activity ...
type Activity struct {
	ID          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Weight      int    `json:"weight" binding:"required"`
	Repetitions int    `json:"repetitions" binding:"required"`
	Sets        int    `json:"sets" binding:"required"`
}

func (workout *Workout) addActivity(activity Activity) {
	workout.Activities = append(workout.Activities, activity)
}

func (user *User) addWorkout(workout Workout) {
	user.Workouts = append(user.Workouts, workout)
}

func (users UserList) getUserByID(id int) (*User, error) {
	for index := 0; index < len(users); index++ {
    if users[index].ID == id {
      return users[index], nil
    }
  }
  return nil, fmt.Errorf("userid not found")
}

var activity = Activity{1, "Barbell", 10, 8, 4}
var workout = Workout{1, 1, []Activity{activity}}
var user = User{1, "username", []Workout{workout}}
var users = UserList{&user}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n", )
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	flag.Parse()
}

func main() {

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	if gin.Mode() == gin.DebugMode {
		glog.Info("Debug mode activated, serving static content from /static")
		router.Use(static.Serve("/static", static.LocalFile("./static", true)))
	}

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/workouts", GetWorkouts)
		api.POST("/workouts", NewWorkout)
		api.GET("/users/:id", GetUser)
		api.POST("/users", NewUser)
	}

	// Start and run the server
	router.Run(":3000")
}

// NewUser Create a new user
func NewUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, users)
}

// GetUser Get existing user
func GetUser(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  if userid, err := strconv.Atoi(c.Param("id")); err == nil {
    user, err := users.getUserByID(userid)
    glog.Info("userid: ", userid, user, err)
    if err != nil {
      c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
      return
    }
		c.JSON(http.StatusOK, user)
	} else {
		// User ID is invalid
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// NewWorkout create a new workout
func NewWorkout(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Workouts not supported yet.",
	})
}

// GetWorkouts ...
func GetWorkouts(c *gin.Context) {

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, users)
}
