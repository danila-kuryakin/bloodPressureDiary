package bot

import (
	"bloodPressureDiary/bp-telegram-bot/internal/bot/constants"

	"gopkg.in/telebot.v3"
)

func mainMenu() *telebot.ReplyMarkup {
	m := &telebot.ReplyMarkup{}
	m.Inline(
		m.Row(
			m.Data(constants.NamePressure, constants.EventPressure),
			m.Data(constants.NameTag, constants.EventTag),
		),
	)
	return m
}

func crudPressureMenu() *telebot.ReplyMarkup {
	m := &telebot.ReplyMarkup{}
	m.Inline(
		m.Row(
			m.Data(constants.NameCreate, constants.EventPressureCreate),
		),
		m.Row(
			m.Data(constants.NameBack, constants.EventBack),
		),
	)
	return m
}

func createPressureMenu() *telebot.ReplyMarkup {
	m := &telebot.ReplyMarkup{}
	m.Inline(
		m.Row(
			m.Data(constants.NamePressureSys, constants.EventPressureSys),
			m.Data(constants.NamePressureDia, constants.EventPressureDia),
		),
		m.Row(
			m.Data(constants.NamePressurePulse, constants.EventPressurePulse),
			m.Data(constants.NamePressureTags, constants.EventPressureTags),
		),
		m.Row(
			m.Data(constants.NameSave, constants.EventPressureSave),
		),
		m.Row(
			m.Data(constants.NameBack, constants.EventBack),
		),
	)
	return m
}

func saveTagPressureMenu() *telebot.ReplyMarkup {
	m := &telebot.ReplyMarkup{}
	m.Inline(
		m.Row(
			m.Data(constants.NameBack, constants.EventBack),
		),
	)
	return m
}

func crudTagMenu() *telebot.ReplyMarkup {
	m := &telebot.ReplyMarkup{}
	m.Inline(
		m.Row(
			m.Data(constants.NameCreate, constants.EventTagCreate),
			m.Data(constants.NameDelete, constants.EventTagDelete),
		),
		m.Row(
			m.Data(constants.NameSave, constants.EventTagSave),
			m.Data(constants.NameBack, constants.EventBack),
		),
	)
	return m
}

func createTagMenu() *telebot.ReplyMarkup {
	m := &telebot.ReplyMarkup{}
	m.Inline(
		m.Row(
			m.Data(constants.NameBack, constants.EventBack),
		),
	)
	return m
}

func deleteTagMenu() *telebot.ReplyMarkup {
	m := &telebot.ReplyMarkup{}
	m.Inline(
		m.Row(
			m.Data(constants.NameBack, constants.EventBack),
		),
	)
	return m
}

func backMenu() *telebot.ReplyMarkup {
	m := &telebot.ReplyMarkup{}
	m.Inline(
		m.Row(
			m.Data(constants.NameBack, constants.EventBack),
		),
	)
	return m
}
