# Habit tracker

This is basicly a TODO app. But it's developed for me, for two reasons, the first one is that I need an app to keep
track of my daily activities, and the second one is that I want to practice my development skills.

In this app you can also add other users routines to your day to day activities. So you can start challenging routines 
with your friends. (e.g. Gym, Study or meditation routines).

You can use a specific hour for your activity or you can create it using a time of day (morning, afternoon, night).

You can share in your profile the routines that you already made.

# Technical information

The app is developed in Java (Spring) and React. 

It uses Spring libraries like: Spring Security, Spring JPA, lombok, and so on.

The app uses PostgreSQL for its database.

#### Backend Entities:

**UserAuth**

* id
* profile_id
* email
* password
* *some spring security options*

**UserProfile**

* id | BIGINT
* auth_id | UUID
* username | STRING
* profile_pic_url | STRING

**Activities**

* id | BIGINT
* userId | BIGINT
* text | STRING
* days | day_of_week (enum)
* specificHour | TIME
* timeZone | time_of_day (enum)
* doneToday | BOOLEAN

**Routines**

* id | BIGINT
* userId | BIGINT
* title | STRING
* ammount_of_subscribers | INT

# Backend Architecture

The folder structure of this app uses the following logic

**entity folder -> component folder**

**e.g:**

* user -> controller -> UserController.java
* activity -> repository -> ActivityRepository.java

# PostgreSQL scripts

#### Endpoints

All REST endpoints of the api. All under "/api".

* POST "/auth/register"
* POST "/auth/login"
* POST "/auth/logout"

* GET "/users"
* GET "/users/me"
* GET "/users/{userId}"
* POST "/users"
* PUT "/users/modify/me"
* DELETE "/users/delete/me"

* GET "/activities/user/me"
* GET "/activities/user/{userId}"
* GET "/activities/routine/{routineId}"
* GET "/activities/{activityId}"
* POST "/activities/create"
* PUT "/activities/modify/{activityId}"
* DELETE "/activities/delete/{activityId}"

* GET "/routines"
* GET "/routines/user/topLikes?quantity=5"
* GET "/routines/user/"
* GET "/routines/user/{userId}"
* POST "/routines/create"
* PUT "/routines/modify/{routineId}"
* DELETE "/routines/delete/{routineId}"
