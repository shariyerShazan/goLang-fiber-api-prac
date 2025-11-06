package main

import (
	// "errors"
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/valyala/fasthttp/reuseport"
)

func main() {
	app := InitApp()
	log.Fatal(app.Listen(":3000"))
}


func InitApp() *fiber.App {
    app := fiber.New()

	// server main path get
    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })


// version controll with api v1,v2 group
	apiV1 := app.Group("/v1")
	apiV1.Get("/" , func (c *fiber.Ctx) error {
		return c.SendString("Version 1 is running")
	})

	apiv2 := app.Group("/v2")
	apiv2.Get("/" , func (c *fiber.Ctx) error {
		return c.SendString("Version 2 is runnign")
	})

    
	// params get
    app.Get("user/:userId?" , func (c *fiber.Ctx) error {
          if c.Params("userId") != "" {
            return c.SendString("User ID is = " + c.Params("userId"))
          } else {
			return c.Status(fiber.StatusBadRequest).SendString("User ID is missing")
		  }
	  })     

	//  api path get
	app.Get("/api/*", func (c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
	} )


	// statuc
	app.Static("/html" , "./public" ) // app.Static("api path" , "folder path" )
	// and it's return the html file of that folder path



	// goFiber Route 
	app.Route("/students" , func(router fiber.Router) {

		// get all
		router.Get("/" , func (c *fiber.Ctx) error {
			return c.SendString("All students here")
		})

		// post
		router.Post("/" , func (c *fiber.Ctx) error {
			return c.SendString("Student data posted")
		} )

		// get by id
		router.Get("/:studentId" , func (c *fiber.Ctx) error {
			if c.Params("studentId") != "" {
				return c.SendString("Here is student id: " + c.Params("studentId"))
			} else {
				return c.Status(fiber.ErrNotFound.Code).SendString("studentId not found")
			}
		} )

		// patch
		router.Patch("/:studentId" , func (c *fiber.Ctx) error {
			if c.Params("studentId") != "" {
				return c.SendString("Updated student data of StudentId no: " + c.Params("studentId"))
			} else {
				return c.Status(fiber.ErrNotFound.Code).SendString("studentId not found")
			}
		})

		// delete 
		router.Delete("/:studentId" , func (c *fiber.Ctx) error {
			if c.Params("studentId") != "" {
				return c.SendString("Deleted student data of StudentId no: " + c.Params("studentId"))
			} else {
				return c.Status(fiber.ErrNotFound.Code).SendString("studentId not found")
			}
		})
	})


	// Wrong api to shutDown 
	app.Get("/shutdown", func(c *fiber.Ctx) error {
		if c.IP() != "127.0.0.1" {
			return c.Status(fiber.StatusForbidden).SendString("Access denied!")
		}
		go func() {
			_ = app.Shutdown()
		}()
		return c.SendString("Server shutting down...")
	})
	

	// token 
	app.Get("/check-token" , func (c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.SendStatus(fiber.StatusUnauthorized) //! OK
			// return c.SendStatus(fiber.StatusBadGateway) //! FAIL
		} else {
			return c.SendStatus(fiber.StatusOK)
		}
	})

	return app
}
//
