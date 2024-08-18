Dependency Injection

Dependency Injection is a design pattern in object-oriented programming that helps us manage dependencies between objects. In this pattern, instead of the class or data structure directly creating its own dependencies, these dependencies are injected into them from the outside. In other words, the target object receives its dependencies through the constructor or methods.

Suppose we have a service that needs to connect to the database. 
At first, your code might look like this:

package main


type Database struct {}

func (db *Database) Query() string {
 return "Data from database"
}

type Service struct {
 db *Database
}

func NewService() *Service {
 return &Service{db: &Database{}}
}

func (s *Service) GetData() string {
 return s.db.Query()
}

func main() {
 service := NewService()
 fmt.Println(service.GetData())
}

In this code, Service structure uses Database directly. This has created a direct dependency between Service and Database.

Now using Dependency Injection:


package main


type Database struct {}

func (db *Database) Query() string {
 return "Data from database"
}

type Service struct {
 db DatabaseInterface
}

type DatabaseInterface interface {
 Query() string
}

func NewService(db DatabaseInterface) *Service {
 return &Service{db:db}
}

func (s *Service) GetData() string {
 return s.db.Query()
}

func main() {
 db := &Database{}
 service := NewService(db)
 fmt.Println(service.GetData())
}


Here, Service instead of directly creating an instance of Database, gets it from outside. This makes dependencies easier to test and replace.

Features:
Direct dependency reduction: Using DI, dependencies are explicitly injected into objects, making them easier to test and modify.
- Ease of writing tests: by separating the dependencies from the class or data structure, Mocks and Stubs can be easily used in tests.
- Increased flexibility: It becomes easier to change and replace dependencies over time.


Suppose we have a small program in which there is a ``UserService'' and this service needs a database to be able to store and retrieve user information.

1. Without dependency injection:

type Database struct {
 // Database methods
}

type UserService struct {
 db *Database
}

func NewUserService() *UserService {
 db := &Database{}
 return &UserService{db:db}
}


In this case, UserService creates the database itself and has a direct dependency on the database. This makes it difficult to change the database or test the UserService without a real database.

2. With dependency injection:


type Database struct {
 // Database methods
}

type UserService struct {
 db *Database
}

func NewUserService(db *Database) *UserService {
 return &UserService{db:db}
}


Here, UserService takes the database as input. That is, when you create a new UserService, you can give it any database (or even a mock for testing).