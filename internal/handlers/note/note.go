package noteHandler

import (
	"rest-api/database"
	"rest-api/internal/model"

	"github.com/gofiber/fiber/v2"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB

	var notes []model.Note

	db.Find(&notes)

	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No notes present",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Notes found",
		"data":    notes,
	})
}

func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	err := c.BodyParser(note)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    nil,
		})
	}

	err = db.Create(&note).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create note",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Note created!",
		"data":    note,
	})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB

	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Note not found",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Note found",
		"data":    note,
	})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"subtitle"`
		Text     string `json:"text"`
	}

	db := database.DB

	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "note not found",
			"data":    nil,
		})
	}

	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)

	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"data":    nil,
		})
	}

	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	db.Save(&note)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Note has been saved",
		"data":    note,
	})
}

func Delete(c *fiber.Ctx) error {
	db := database.DB

	var note model.Note

	id := c.Params("noteId")

	db.Find(&note, "id = ?", id)

	if note.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "note not found",
			"data":    nil,
		})
	}

	err := db.Delete(&note, "id = ?", id).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not delete note",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Note has been deleted",
		"data":    nil,
	})

}
