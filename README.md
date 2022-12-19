
# Arent Test Backend Project - Minh Vo
## Summarize Table

```golang
// Management Users in Applications
type User {
  ID                  uint     `json:"id" gorm:"primary_key"`
  Email               string   `json:"email"`
  Password                string   `json:"password"`
  Name                    string   `json:"name"`
  AchievementWeightFrom   float32  `json:"achievement_weight_from"`
  AchievementWeightTo     float32  `json:"achievement_weight_to"`
}

// Management Sessions in Applications ( e.g: Morning, Lunch, Dinner, Snack, ...)
type Session {
  ID     uint    `json:"id" gorm:"primary_key"`
  Name   string  `json:"name"`
}

// Management dishes in Applications ( e.g: Beefsteak, Rice, Noodle, ...)
// Each dish will have a fixed calorie level that is used to calculate
// the amount of calories the user takes in in  a day or anything else related.
type Dish struct {
  ID       uint    `json:"id" gorm:"primary_key"`
  Name     string  `json:"name"`
  Calories float32 `json:"calories"`
}

// Management meal history of user in Applications ( e.g: Beefsteak, Rice, Noodle, ...)
// It belongs to an User and a Session
// It have one Dish (or can be have many dishes in a meal in a day)
// currently i design it have one dish in a meal

type MealHistory struct {
  ID           uint       `json:"id" gorm:"primary_key"`
  Date         time.Time  `json:"date"`
  SessionID      uint       `json:"session_id"`
  DishID           uint       `json:"dish_id"`
  UserID           uint       `json:"user_id"`
}

// Management Body Record of Users
type BodyRecord struct {
  ID       uint64         `json:"id" gorm:"primary_key"`
  Date     time.Time      `json:"date"`
  Weight   float32        `json:"weight"`
  Height   float32        `json:"height"`
  UserID   uint64         `json:"user_id"`
}

// Management Diary of User
// This record will tracking Exercise that user have trained in day and help we calculate calories loss by exercise
type Diary struct {
  ID                      uint64              `json:"id" gorm:"primary_key"`
  Date                    time.Time           `json:"date"`
  Note                    string              `json:"note"`
  UserID                  uint64              `json:"user_id"`
  User                    User
  ExerciseHistories       []ExerciseHistory
}

// This Model used to save the exercises that the user often uses
type UserExercise struct {
  ID             uint64         `json:"id" gorm:"primary_key"`
  UserID         uint64         `json:"user_id" gorm:"index:,unique,composite:user_exercises_id"`
  ExerciseID     uint64         `json:"exercise_id" gorm:"index:,unique,composite:user_exercises_id"`
}

// List of Exercise in Application
type Exercise struct {
  ID             uint64    `json:"id" gorm:"primary_key"`
  Name           string    `json:"name"`
  CaloriesBurned float32   `json:"calories_burned"`
}

// This model used to track the exercise that user already train in a day
type ExerciseHistory struct {
  ID               uint64       `json:"id" gorm:"primary_key"`
  DiaryID          uint64       `json:"diary_id"`
  ExerciseID       uint64       `json:"exercise_id"`
  Exercise         Exercise
}

// This model used to management the health posts
type Post struct {
  ID               uint64       `json:"id" gorm:"primary_key"`
  DiaryID          uint64       `json:"diary_id"`
  ExerciseID       uint64       `json:"exercise_id"`
  Exercise         Exercise
}
```
## API Documents

**1.  User**
> Get all users
> `GET` /api/users
```json
`curl "GET /api/users"
{
  "result": [
    {
      "id": 1,
      "email": "man@gmail.com",
      "password": "12345",
      "name": "Man",
      "achievement_weight_from": 60,
      "achievement_weight_to": 75
    }, {...}
  ]
  "status": true
}
```
> Get by id
> `GET` /api/users/:id
```json
`curl "GET /api/users/1"
{
  "result": {
      "id": 1,
      "email": "man@gmail.com",
      "password": "12345",
      "name": "Man",
      "achievement_weight_from": 60,
      "achievement_weight_to": 75
    }
  "status": true
}
```
> Create user
> `POST` /api/users
```json
`curl "POST /api/users"
params:
{
   "email": "man@gmail.com",
   "password": "12345",
   "name": "Man",
   "achievement_weight_from": 60,
   "achievement_weight_to": 75
 }
 Response:
{
  "result": {
      "id": 1,
      "email": "man@gmail.com",
      "password": "12345",
      "name": "Man",
      "achievement_weight_from": 60,
      "achievement_weight_to": 75
    }
  "status": true
}
```
> Update user
> `PATCH` /api/users
```json
`curl "PATCH /api/users"
params:
{
   "email": "man@gmail.com",
   "password": "12345",
   "name": "Man",
   "achievement_weight_from": 60,
   "achievement_weight_to": 75
 }
 response:
{
  "result": {
      "id": 1,
      "email": "man@gmail.com",
      "password": "12345",
      "name": "Man",
      "achievement_weight_from": 60,
      "achievement_weight_to": 75
    }
  "status": true
}
```
> Delete user
> `DELETE` /api/users
```json
`curl "DELETE /api/users"
{
  status: true
}
```
**2. Meal History**
> Get all Meal Histories (this will get all Meal History by user token)
> `GET` /api/meal_histories
```json
e.g: curl "GET /api/meal_histories".
{
  "result": [
      {
            "id": 1,
            "date": "2022-12-09T12:00:00Z",
            "session_id": 1,
            "dish_id": 12,
            "user_id": 1,
            "Session": {
                "id": 1,
                "name": "Morning"
            },
            "Dish": {
                "id": 12,
                "name": "Small Seafood soup",
                "calories": 80
            },
            "User": {
                "id": 1,
                "email": "man@gmail.com",
                "password": "12345",
                "name": "Man",
                "achievement_weight_from": 60,
                "achievement_weight_to": 75
            }
        },{...}
    ],
    "status": true,
    "total": 28,
    }
}
```
> Get all Meal Histories By Session ( reference with user token )
> `GET` /api/sessions/:id/meal_histories
```json
e.g: curl "GET /api/sessions/2/meal_histories".
{
  "result": [
      {
            "id": 1,
            "date": "2022-12-09T12:00:00Z",
            "session_id": 2,
            "dish_id": 12,
            "user_id": 1,
            "Session": {
                "id": 1,
                "name": "Morning"
            },
            "Dish": {
                "id": 12,
                "name": "Small Seafood soup",
                "calories": 80
            },
            "User": {
                "id": 1,
                "email": "man@gmail.com",
                "password": "12345",
                "name": "Man",
                "achievement_weight_from": 60,
                "achievement_weight_to": 75
            }
        }
    ],
    "status": true,
    "total": 28,
    }
}
```
> Get User Achievement Rate
> `GET` /meal_histories/:user_id/achievement_rate
```json
e.g: curl "GET /api/achievement_rate"
{
  "result": {
    "achievement_rate": 75,
    "date": "2022-12-12T00:19:38.155958+07:00"
  },
  "status": true
}
```

> Get User Body Fat Percent Graph
> `GET` /api/body_fat_percent_graph
```json
e.g: curl "GET /api/body_fat_percent_graph"
{
  "result": [
    {
      "date": "2022-12-09T12:00:00Z",
      "total_calories": 80
    },{
      "date": "2022-12-10T12:00:00Z",
      "total_calories": 290
    },{
      "date": "2022-12-11T07:00:00Z",
      "total_calories": 740
    }
  ],
  "status": true
}
```
> Create Meal History
> `POST` /api/meal_histories
```json
e.g: curl "POST /api/meal_histories".
params:
{
  "date": "2022-12-09T12:00:00Z",
  "session_id": 1,
  "dish_id": 12,
  "user_id": 2
}
{
  "result": {
        "id": 21,
        "date": "2022-12-09T12:00:00Z",
        "session_id": 1,
        "dish_id": 12,
        "user_id": 1,
        "Session": {...},
        "Dish": {...},
        "User": {...}
     },
    "status": true
    }
}
```
> Update Meal History
> `PATCH` /api/meal_histories/:id
```json
e.g: curl "PATCH /api/meal_histories".
params:
{
  "date": "2022-12-09T12:00:00Z",
  "session_id": 56,
  "dish_id": 15,
  "user_id": 3
}
{
  "result": {
        "id": 21,
        "date": "2022-12-09T12:00:00Z",
        "session_id": 56,
        "dish_id": 15,
        "user_id": 3,
        "Session": {...},
        "Dish": {...},
        "User": {...}
     },
    "status": true
    }
}
```
> Delete Meal History
> `DELETE` /api/meal_histories/:id
```json
e.g: curl "DELETE /api/meal_histories/23".
{
  "status": true
}
```
**3. Dish**
> Get all dishes
> `GET` /api/dishes
```json
`curl "GET /api/dishes"
{
  "result": [
    {
      "id": 1,
      "name": "Beefsteak",
      "calories": 180
    }, {...}
  ]
  "status": true
}
```
> Search dish by id or name
> `GET` /api/dishes/:id
```json
`curl "GET /dishes/beefsteak"
{
  "result": [
    {
      "id": 1,
      "name": "Beefsteak",
      "calories": 180
    }, {...}
  ]
  "status": true
}
```
**3. Session**
> Get all sessions
> `GET` /sessions
```json
`curl "GET /sessions"
{
  "result": [
    {
      "id": 1,
      "name": "Morning",
    }, {...}
  ]
  "status": true
}
```
> Search session by id or name
> `GET` /sessions/:id
```json
`curl "GET /sessions/morning"
{
  "result": [
    {
      "id": 1,
      "name": "Morning",
    }, {...}
  ]
  "status": true
}
```
> Search user exercise
> `GET` /api/user_exercises
```json
`curl "GET /api/user_exercises"
{
  "result": [
    {
      "user_id": 1,
      "exercise_id": 2,
    }, {...}
  ]
  "status": true
}
```

> Bookmark exercise
> `POST` /api/user_exercises
```json
`curl "POST /api/user_exercises"
{
  "result": {
    "user_id": 1,
    "exercise_id": 2,
  },
  "status": true
}
```
> Search body records
> `GET` /api/body_records
```json
`curl "GET /api/body_records"
{
  "result": [
    {
      "date": "2022-12-09T12:00:00Z",
      "weight": 70,
      "height": 180,
      "user_id": 2,
    }, {...}
  ]
  "status": true
}
```
> Create Body Records
> `POST` /api/body_records
```json
`curl "POST /api/body_records"
{
  "date": "2022-12-09T12:00:00Z",
  "weight": 70,
  "height": 180,
}
```
> Search Diaries
> `GET` /api/diaries
```json
`curl "GET /api/diaries"
{
  "result": [
    {
      "date": "2022-12-09T12:00:00Z",
      "note": "Diary Note",
      "user_id": 2,
      "exercise_histories": [{ "exercise_id": 1, "diary_id": 2 }],
    }, {...}
  ],
  "status": true
}
```
> GET Diary
> `GET` /api/diaries/:id
```json
`curl "GET /api/diaries/:id"
{
  "result": {
    "date": "2022-12-09T12:00:00Z",
    "note": "Diary Note",
    "user_id": 2,
    "exercise_histories": [{ "exercise_id": 1, "diary_id": 2 }],
  },
  "status": true
}
```

> Create Diary
> `POST` /api/diaries
```json
`curl "GET /api/diaries/:id"
{
  "date": "2022-12-09T12:00:00Z",
  "note": "Diary Note",
  "user_id": 2
}
```

## How it works?

**Date/Achievement rate Section**
This section will be use `Get User Achievement Rate`
The system will be calculate based on Meal History data

**Weight/Body Fat Percentage Graph**
This section will be use `Get User Body Fat Percent Graph`
The system response a list of body fat of user based on Meal History data

**Button to transition to input**
With each input will be use with API `Get Session by Name` corresponding

**Meal history**
This section will be use API `GET /api/meal_histories` or `GET /sessions/:id/meal_histories` to get data corresponding with requirement

**Exercise record**
This section will be use API `GET /api/user_exercises` to get data corresponding with requirement

**Health Post**
This section will be use API `GET /api/posts` or `GET /api/posts/:id` to get data corresponding with requirement

