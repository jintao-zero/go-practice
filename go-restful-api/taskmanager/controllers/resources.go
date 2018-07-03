package controllers

import "go-practice/go-restful-api/taskmanager/models"

type (
    //For Post - /user/register
    UserResource struct {
        Data models.User `json:"data"`
    }

    // For Post - /user/login
    LoginResource struct {
        Data LoginModel `json:"data"`
    }

    // Response for authorized user Post - /user/login
    AuthUserResource struct {
        Data AuthUserModel `json:"data"`
    }

    // Model for authentication
    LoginModel struct {
        Email string `json:"email"`
        Password string `json:"password"`
    }
    // Model for authorized user with access token
    AuthUserModel struct {
        User models.User `json:"user"`
        Token string `json:"token"`
    }

    // For Post/Put - /tasks
    // For Get - /tasks/id
    TaskResource struct {
        Data models.Task `json:"data"`
    }

    // For Get - /tasks
    TasksResource struct {
        Data []models.Task `json:"data"`
    }
)
