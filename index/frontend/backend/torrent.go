package backend

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koromaru-tracker/koromaru/index"
	"github.com/koromaru-tracker/koromaru/index/model"
	"gorm.io/gorm"
)

func TorrentView(c *fiber.Ctx) error {
	db := index.GetLocal[*gorm.DB](c, "db")
	var torrent model.Torrent

	if err := db.Where("id = ?", c.Params("id")).First(&torrent).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Render("backend/torrent/view", fiber.Map{
		"Title":   "View Torrent",
		"Torrent": torrent,
	})
}

func TorrentCreate(c *fiber.Ctx) error {
	db := index.GetLocal[*gorm.DB](c, "db")
	var torrent model.Torrent

	if err := c.BodyParser(&torrent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := torrent.Create(db); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Redirect("/backend/torrent/list")
}

func TorrentCreateView(c *fiber.Ctx) error {
	return c.Render("backend/torrent/create", fiber.Map{
		"Title": "Create Torrent",
	})
}

func TorrentList(c *fiber.Ctx) error {
	db := index.GetLocal[*gorm.DB](c, "db")
	var torrents []model.Torrent

	db.Find(&torrents)

	return c.Render("backend/torrent/list", fiber.Map{
		"Title":    "List Torrent",
		"Torrents": torrents,
	})
}
