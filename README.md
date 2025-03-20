
A example of a web APP that use Gin framework.


## Get started

install swag
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

download dependency
```bash
go mod download
```

start service
```bash
make run-dev
```
## Architecture

![image](resources/images/Architecture.png)


### Project Directory

```text
├── api                     # web api
│   ├── auth                # connected with authentication, including User Login
│   │   ├── route.go        # authentication route 
│   │   ├── view.go         # authentication handles and swag documentation
│   │   ├── view_test.go    # web api tests
│   │   └── viewset.go      # authentication views
│   ├── instructor
│   │   ├── router.go
│   │   ├── view.go
│   │   ├── view_test.go
│   │   └── viewset.go
│   ├── middleware          # Gin Middleware
│   │   ├── auth.go         # Authentication Middleware
│   │   ├── content.go      # Core module for api viewset
│   │   └── cors.go         # CORS support Middleware
│   └── route.go            # Gin Routes
├── models                  # model structures
│   └── instructor.go
├── main.go
├── service                 # Vunction modules
│   ├── Instructor.go
│   ├── Instructor_test.go
│   ├── auth.go
│   └── auth_test.go
└── types                   # Various structures
    ├── auth.go
    ├── common.go           # common functions
    └── instructor.go
```

### Endpoints
```text
GET    /docs/*any                       # Swagger Documentation
            Provides API-related Swagger documentation, allowing users to browse and test the API online.
POST   /api/v1/auth/login               # User Login
            Authenticates user credentials and returns an access token for subsequent API requests.
GET    /api/v1/instructors              # List All Instructors
            Retrieves a list of all instructors, with support for pagination and filtering.
POST   /api/v1/instructors              # Create a New Instructor
            Adds a new instructor and returns the details of the created instructor.
GET    /api/v1/instructors/:id          # Get Instructor by ID
            Retrieves detailed information about a specific instructor based on the given ID.
PUT    /api/v1/instructors/:id          # Update Instructor Information
            Updates the details of a specific instructor using the provided ID. Supports both partial and full updates.
DELETE /api/v1/instructors/:id          # Delete an Instructor
            Removes an instructor based on the given ID. This action is irreversible.
```

### View or handle
There are two primary functions: one is the Swagger documentation, and the other is the API processing function.
```go
// @Summary	login
// @Description	login
// @Accept	json
// @Produce	json
// @Tags	auth
// @Param   form	body	types.LoginRequest	true	"form""
// @Success 200 {object} models.Instructor "successful"
// @Router /auth/login [POST]
func LoginHandle(ctx *gin.Context) {
	NewAuthViewset(ctx).Login()
}
```

### Viewset
A ViewSet is a collection of APIs that enable adding, deleting, updating, and retrieving resources, such as user data.
```go
type IUser interface {
	GetInstance()
	ListInstances()
	CreateInstance()
	UpdateInstance()
	DeleteInstance()
}


type UserViewset struct {
    s service.IInstructor
}


func (v InstructorViewset) ListInstances() {
    data, err := v.s.ListAllInstances()
    if err != nil {
        v.ResponseError(err)
        return
    }
    v.ResponseOK(data)
}

func (v InstructorViewset) getPathIdParameter() (int, error) {
    id := v.Param("id")
    instanceId, err := strconv.Atoi(id)
    if err != nil {
        return instanceId, errors.New("Instructor ID Error")
    }
    return instanceId, nil
}

func (v InstructorViewset) GetInstance() {
    id, err := v.getPathIdParameter()
    if err != nil {
        v.ResponseError(err)
        return
    }
    obj, err := v.s.GetInstance(id)
    if err != nil {
        v.ResponseError(err)
        return
    }
    v.ResponseOK(obj)
}

func (v InstructorViewset) CreateInstance() {
    var form types.CreateInstructorRequest
    if err := v.Bind(&form); err != nil {
        v.ResponseError(err)
        return
    }
    if err := v.s.CreateInstance(form.Instructor); err != nil {
        v.ResponseError(err)
        return
    }
    v.ResponseOK(nil)
}

func (v InstructorViewset) UpdateInstance() {
    var (
        id   int
        form types.UpdateInstructorRequest
        err  error
    )
    if id, err = v.getPathIdParameter(); err != nil {
        v.ResponseError(err)
        return
    }
    
    if err = v.Bind(&form); err != nil {
        v.ResponseError(err)
        return
    }
    
    if err = v.s.UpdateInstance(id, form); err != nil {
        v.ResponseError(err)
        return
    }
    v.ResponseOK(nil)
}

func (v InstructorViewset) DeleteInstance() {
    id, err := v.getPathIdParameter()
    if err != nil {
        v.ResponseError(err)
        return
    }
    if err = v.s.DeleteInstance(id); err != nil {
        v.ResponseError(err)
        return
    }
    v.ResponseOK(nil)
}

```

### Service

The Service layer implements business logic functions, including database operations such as inserting, deleting, updating, and querying data.

```go
type IUser Interface {
	ListUsers() ([]*models.User, error)
	GetUserById(id int) (*models.User, error)
	DeleteUserById(id int) error
	CreateUser(form types.CreateUserRequest) (*models.User, error)
	UpdateUser(id int, form types.UpdateUserRequest) (*models.User, error)
}
```