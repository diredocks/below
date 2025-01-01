package page

import (
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

// Get all Pages from Database
func Index(c *fiber.Ctx) error {
	res, err := GetAllDB(&service.Page{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get pages",
			"msg":   err.Error(),
		})
	}

	return c.JSON(res)
}

// Add a Page to Database
func Add(c *fiber.Ctx) error {
	p := c.Locals("validatedBody").(*service.Page)

	if err := InsertDB(p); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to add page",
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": "page added",
		"id":      p.ID,
	})
}

// Get a Page from Database
func Get(c *fiber.Ctx) error {
	p := c.Locals("validatedBody").(*service.Page)
	res, err := QueryDB(p)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get page",
			"msg":   err.Error(),
		})
	}

	return c.JSON(res)
}

func Del(c *fiber.Ctx) error {
	p := c.Locals("validatedBody").(*service.Page)
	err := DelDB(p)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete page",
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": "deleted page",
	})
}

func Update(c *fiber.Ctx) error {
	sitemapURL := "https://diredocks.github.io/blog/sitemap.xml"
	data, err := FetchSitemap(sitemapURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update sitemap",
			"msg":   err.Error(),
		})
	}
	pages, err := ParseSitemap(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update sitemap",
			"msg":   err.Error(),
		})
	}
	for _, p := range pages {
		InsertDB(&p)
	}
	return c.JSON(fiber.Map{
		"success": "updated sitemap",
	})
}
