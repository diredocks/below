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

func GetSite(c *fiber.Ctx) error {
	q := c.Locals("validatedBody").(*service.ReqSite)
	res, err := QuerySiteDB(q)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to get site",
				"msg":   err.Error(),
			})
	}
	return c.JSON(res)
}

func GetAllSite(c *fiber.Ctx) error {
	sites, err := GetAllSiteDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to get all sites",
				"msg":   err.Error(),
			})
	}
	return c.JSON(sites)
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

func AddSite(c *fiber.Ctx) error {
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

func UpdateSite(c *fiber.Ctx) error {
	q_site := c.Locals("validatedBody").(*service.ReqSite)
	q_sitemap, err := QuerySiteDB(q_site)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "failed to get site",
				"msg":   err.Error(),
			})
	}

	data, err := FetchSitemap(q_sitemap.SiteMap)
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
	site.SiteMap = q_sitemap.SiteMap // Fill Sitemap URL

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
