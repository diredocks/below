package page

import (
	"below/internal/config"
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

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

	if err := DelDB(p); err != nil {
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
	data, err := FetchSitemap(config.Config("SITEMAP_URL"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch sitemap",
			"msg":   err.Error(),
		})
	}
	pages, err := ParseSitemap(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to parse sitemap",
			"msg":   err.Error(),
		})
	}
	affected, err := InsertsDB(pages)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to insert sitemap into database",
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"success":  "updated sitemap",
		"affected": affected,
	})
}
