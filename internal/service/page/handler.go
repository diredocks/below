package page

import (
	"below/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTeapot).
		SendString("I'm not a Tea pot!")
}

func GetPage(c *fiber.Ctx) error {
	q := c.Locals("validatedBody").(*service.ReqPage)
	res, err := QueryPageDB(q)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to get page",
				"msg":   err.Error(),
			})
	}
	return c.JSON(res)
}

func DelPage(c *fiber.Ctx) error {
	q := c.Locals("validatedBody").(*service.ReqPage)
	err := DelPageDB(q)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to delete page",
				"msg":   err.Error(),
			})
	}
	return c.JSON(fiber.Map{
		"success": "deleted page",
	})
}

func DelSite(c *fiber.Ctx) error {
	q := c.Locals("validatedBody").(*service.ReqSite)
	err := DelSiteDB(q)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to delete site",
				"msg":   err.Error(),
			})
	}
	return c.JSON(fiber.Map{
		"success": "deleted site",
	})
}

func AddOrUpdate(c *fiber.Ctx) error {
	q := c.Locals("validatedBody").(*service.ReqSiteMap)
	data, err := FetchSitemap(q.SiteMap)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to fetch sitemap",
				"msg":   err.Error(),
			})
	}

	site, pages, err := ParseSitemap(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to parse sitemap",
				"msg":   err.Error(),
			})
	}
	site.SiteMap = q.SiteMap // Fill Sitemap URL

	affected, err := InsertPagesDB(site, pages)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to insert sitemap into database",
				"msg":   err.Error(),
			})
	}

	return c.JSON(fiber.Map{
		"success":  "updated sitemap",
		"affected": affected,
	})
}

/*
func Add(c *fiber.Ctx) error {
	p := c.Locals("validatedBody").(*service.Page)

	if err := InsertDB(p); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to add page",
				"msg":   err.Error(),
			})
	}

	return c.Status(fiber.StatusCreated).
		JSON(fiber.Map{
			"success": "page added",
			"id":      p.ID,
		})
}

func Del(c *fiber.Ctx) error {
	p := c.Locals("validatedBody").(*service.Page)

	if err := DelDB(p); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to delete page",
				"msg":   err.Error(),
			})
	}

	return c.JSON(fiber.Map{
		"success": "deleted page",
	})
}
*/
