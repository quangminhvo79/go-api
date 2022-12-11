
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

```
## API Documents

**1.  User**
> Get all users
> `GET` /users
```json
`curl "GET /users"
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
> `GET` /users/:id
```json
`curl "GET /users/1"
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
> `POST` /users
```json
`curl "POST /users"
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
> `PATCH` /users
```json
`curl "PATCH /users"
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
> `DELETE` /users
```json
`curl "PATCH /users"
{
  status: true
}
```
**2. Meal History**
> Get all Meal Histories by User
> `GET` /meal_histories/:user_id
```json
e.g: curl "GET /meal_histories/1".
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
> Get all Meal Histories by User By Session
> `GET` /meal_histories/:user_id/sessions/:session_id
```json
e.g: curl "GET /meal_histories/1/sessions/2".
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
e.g: curl "GET /meal_histories/23/achievement_rate"
{
  "result": {
    "achievement_rate": 75,
    "date": "2022-12-12T00:19:38.155958+07:00"
  },
  "status": true
}
```

> Get User Body Fat Percent Graph
> `GET` /meal_histories/:user_id/body_fat_percent_graph
```json
e.g: curl "GET /meal_histories/23/achievement_rate"
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
> `POST` /meal_histories
```json
e.g: curl "POST /meal_histories".
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
> `PATCH` /meal_histories/:id
```json
e.g: curl "PATCH /meal_histories".
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
> `DELETE` /meal_histories/:id
```json
e.g: curl "DELETE /meal_histories/23".
{
  "status": true
}
```
**3. Dish**
> Get all dishes
> `GET` /dishes
```json
`curl "GET /dishes"
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
> Search dish by name
> `GET` /dishes/:name
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
> Search session by name
> `GET` /sessions/:name
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
This section will be use API `Get Meal History by Users` or `Get Meal History By Session` to get data corresponding with requirement
