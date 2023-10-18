# user-api crud implementation

Creating an **API** implementing a *users* **CRUD** with golang. I will list all the packages I used to implement the code.

* [godotenv](https://github.com/joho/godotenv)
1.  Dotenv load variables from a .env file into ENV when the environment is bootstrapped.

* [gin-gonic](https://github.com/gin-gonic/gin)
1. Gin is a web framework written in Go. It has a martini-like API with up to 40 times faster performance thanks to httprouter.

>The **router := gin.New()** statement will create a new gin router. Routers can be initialized in two ways, one using *New()* and the other using *Default()*.
>The difference is that *New()* boots a router without any **middleware** while *Default()* boots the router with **logger** and **recovery middleware**

