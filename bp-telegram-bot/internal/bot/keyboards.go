package bot

import (
	"bp-telegram-bot/internal/bot/constants"
	"fmt"

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

func deleteTagMenu(tags map[int]string) *telebot.ReplyMarkup {
	m := &telebot.ReplyMarkup{}
	if len(tags) == 0 {
		m.Inline(
			m.Row(
				m.Data(constants.NameBack, constants.EventBack),
			),
		)
		return m
	}
	var buttons []telebot.Btn
	for i := 0; i < len(tags); i += 2 {
		btn := m.Data(tags[i], fmt.Sprintf("event_delete_tag_%d", i))
		buttons = append(buttons, btn)

		if i+1 < len(tags) {
			btn = m.Data(tags[i+1], fmt.Sprintf("event_delete_tag_%d", i+1))
		} else {
			btn = m.Data(" ", "empty")
		}
		buttons = append(buttons, btn)
	}

	var rows []telebot.Row
	const perRow = 2

	for i := 0; i < len(buttons); i += perRow {
		end := i + perRow
		if end > len(buttons) {
			end = len(buttons)
		}

		// берём срез кнопок для текущей строки
		rowButtons := buttons[i:end]

		// создаём строку (Row)
		rows = append(rows, m.Row(rowButtons...))
	}

	// ─── Добавляем кнопку "Назад" в отдельной строке ───
	backBtn := m.Data(constants.NameBack, constants.EventBack)
	rows = append(rows, m.Row(backBtn))

	// ─── Собираем всю клавиатуру ───
	m.Inline(rows...)
	return m
}

func addPressureTagMenu(tags map[int]string) *telebot.ReplyMarkup {
	m := &telebot.ReplyMarkup{}
	if len(tags) == 0 {
		m.Inline(
			m.Row(
				m.Data(constants.NameBack, constants.EventBack),
			),
		)
		return m
	}
	var buttons []telebot.Btn
	for i := 0; i < len(tags); i += 2 {
		btn := m.Data(tags[i], fmt.Sprintf("event_add_tag_%d", i))
		buttons = append(buttons, btn)

		if i+1 < len(tags) {
			btn = m.Data(tags[i+1], fmt.Sprintf("event_add_tag_%d", i+1))
		} else {
			btn = m.Data(" ", "empty")
		}
		buttons = append(buttons, btn)
	}

	var rows []telebot.Row
	const perRow = 2

	for i := 0; i < len(buttons); i += perRow {
		end := i + perRow
		if end > len(buttons) {
			end = len(buttons)
		}

		// берём срез кнопок для текущей строки
		rowButtons := buttons[i:end]

		// создаём строку (Row)
		rows = append(rows, m.Row(rowButtons...))
	}

	// ─── Добавляем кнопку "Назад" в отдельной строке ───
	backBtn := m.Data(constants.NameBack, constants.EventBack)
	rows = append(rows, m.Row(backBtn))

	// ─── Собираем всю клавиатуру ───
	m.Inline(rows...)
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
