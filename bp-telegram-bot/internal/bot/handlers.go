package bot

import (
	"bloodPressureDiary/bp-telegram-bot/internal/bot/constants"
	"fmt"
	"log"

	"gopkg.in/telebot.v3"
)

func (bot *Bot) StartHandle(c telebot.Context) error {
	bot.state[c.Chat().ID] = constants.StateIdle
	return c.Send("Выберите раздел:", mainMenu())
}

func (bot *Bot) Callbacks(c telebot.Context) error {
	chatID := c.Sender().ID
	data := c.Callback().Data[1:]

	if bot.buffer[c.Chat().ID] == nil {
		bot.buffer[c.Chat().ID] = map[string]any{}
	}

	bot.buffer[c.Chat().ID]["MenuMsg"] = c.Callback().Message.ID

	log.Println(chatID, c.Callback().Message.ID)

	switch data {
	case constants.EventPressure:
		bot.state[chatID] = constants.StatePressure
		return c.Edit("Меню давления.", crudPressureMenu())

	case constants.EventTag:
		bot.state[chatID] = constants.StateTag
		return c.Edit("Меню тегов.", crudTagMenu())

		// CRUD Pressure
	case constants.EventPressureCreate:
		bot.state[chatID] = constants.StatePressureCreate
		return c.Edit("Выберете что хотите ввести.", createPressureMenu())

		// ADD Pressure
	case constants.EventPressureSys:
		bot.state[chatID] = constants.StatePressureSys
		return c.Edit("Введите верхнее давление и нажмите Ввод.", backMenu())

	case constants.EventPressureDia:
		bot.state[chatID] = constants.StatePressureDia
		return c.Edit("Введите нижнее давление и нажмите Ввод.", backMenu())

	case constants.EventPressurePulse:
		bot.state[chatID] = constants.StatePressurePulse
		return c.Edit("Введите пульс и нажмите Ввод.", backMenu())

	case constants.EventPressureTags:
		bot.state[chatID] = constants.StatePressureTags
		//bot.buffer[c.Chat().ID] = map[string]any{}
		return c.Edit("Выберете теги из имеющихся.", saveTagPressureMenu())

	case constants.EventPressureSave:
		fmt.Println("EventPressureSave", bot.buffer[c.Chat().ID]["Systolic"],
			bot.buffer[c.Chat().ID]["Diastolic"],
			bot.buffer[c.Chat().ID]["Pulse"],
			bot.buffer[c.Chat().ID]["Tags"])

		bot.state[chatID] = constants.StateIdle
		bot.buffer[c.Chat().ID] = nil
		return c.Edit("Успешно сохранено.\nВыберите раздел:", mainMenu())

		// CRUD TAG
	case constants.EventTagCreate:
		bot.state[chatID] = constants.StateTagCreate
		return c.Edit("Введите тег, который хотите создать", createTagMenu())

	case constants.EventTagDelete:
		bot.state[chatID] = constants.StateTagDelete
		return c.Edit("Введите тег, который хотите удалить.", deleteTagMenu())

	case constants.EventTagSave:
		fmt.Println("EventPressureSave", bot.buffer[c.Chat().ID]["TagsList"])

		bot.state[chatID] = constants.StateIdle
		bot.buffer[c.Chat().ID] = nil
		return c.Edit("Успешно сохранено.\nВыберите раздел:", mainMenu())

	case constants.EventTagName:
		bot.state[chatID] = constants.StateTagName
		return c.Edit("Введите название тега и нажмите Ввод", backMenu())

		// Back
	case constants.EventBack:
		switch bot.state[chatID] {
		case constants.StatePressure, constants.StateTag:
			bot.state[chatID] = constants.StateIdle
			return c.Edit("Выберите раздел:", mainMenu())

		case constants.StatePressureSys, constants.StatePressureDia,
			constants.StatePressurePulse, constants.StatePressureTags:
			bot.state[chatID] = constants.StatePressureCreate
			return c.Edit("Выберете что хотите ввести.", createPressureMenu())

		case constants.StatePressureCreate:
			bot.state[chatID] = constants.StatePressure
			return c.Edit("Выберете что хотите ввести.", crudPressureMenu())

		case constants.StateTagName:
			bot.state[chatID] = constants.StateTagCreate
			return c.Edit("Выберете что хотите ввести.", createTagMenu())

		case constants.StateTagCreate, constants.StateTagDelete:
			bot.state[chatID] = constants.StatePressure
			return c.Edit("Выберите раздел:", crudTagMenu())

		default:
			bot.state[chatID] = constants.StateIdle
			return c.Edit("Выберите раздел:", mainMenu())
		}
	}

	bot.state[c.Chat().ID] = constants.StateIdle
	return c.Edit("Ой, что-то пошло не так.\nВыберите раздел:", mainMenu())
}

func (bot *Bot) Messages(c telebot.Context) error {
	chatID := c.Sender().ID
	msg := c.Message().Text

	log.Println("Messages", chatID, bot.state[chatID], bot.buffer[c.Chat().ID])

	if bot.buffer[c.Chat().ID] == nil {
		bot.buffer[c.Chat().ID] = map[string]any{}
	}

	log.Println(msg)
	switch bot.state[chatID] {
	case constants.StatePressureSys:
		bot.buffer[c.Chat().ID]["Systolic"] = msg
		bot.state[chatID] = constants.StatePressureCreate
		bot.DeleteLastUserMsg(c)
		msgID := bot.buffer[c.Chat().ID]["MenuMsg"].(int)

		editMsg := &telebot.Message{
			ID:   msgID,
			Chat: c.Chat(),
		}
		_, err := bot.tb.Edit(editMsg, "Выберете что хотите ввести.", createPressureMenu())
		if err != nil {
			return err
		}
		return nil

	case constants.StatePressureDia:
		bot.buffer[c.Chat().ID]["Diastolic"] = msg
		bot.state[chatID] = constants.StatePressureCreate
		bot.DeleteLastUserMsg(c)
		msgID := bot.buffer[c.Chat().ID]["MenuMsg"].(int)

		editMsg := &telebot.Message{
			ID:   msgID,
			Chat: c.Chat(),
		}
		_, err := bot.tb.Edit(editMsg, "Выберете что хотите ввести.", createPressureMenu())
		if err != nil {
			return err
		}
		return nil

	case constants.StatePressurePulse:
		bot.buffer[c.Chat().ID]["Pulse"] = msg
		bot.state[chatID] = constants.StatePressureCreate
		bot.DeleteLastUserMsg(c)
		msgID := bot.buffer[c.Chat().ID]["MenuMsg"].(int)

		editMsg := &telebot.Message{
			ID:   msgID,
			Chat: c.Chat(),
		}
		_, err := bot.tb.Edit(editMsg, "Выберете что хотите ввести.", createPressureMenu())
		if err != nil {
			return err
		}
		return nil

	case constants.StatePressureTags:
		bot.buffer[c.Chat().ID]["Tags"] = msg
		bot.state[chatID] = constants.StatePressureCreate
		bot.DeleteLastUserMsg(c)
		msgID := bot.buffer[c.Chat().ID]["MenuMsg"].(int)

		editMsg := &telebot.Message{
			ID:   msgID,
			Chat: c.Chat(),
		}
		_, err := bot.tb.Edit(editMsg, "Выберете что хотите ввести.", createPressureMenu())
		if err != nil {
			return err
		}
		return nil

	case constants.StateTagCreate:
		bot.buffer[c.Chat().ID]["TagsList"] = msg
		bot.state[chatID] = constants.StateTagCreate
		bot.DeleteLastUserMsg(c)
		msgID := bot.buffer[c.Chat().ID]["MenuMsg"].(int)

		editMsg := &telebot.Message{
			ID:   msgID,
			Chat: c.Chat(),
		}
		_, err := bot.tb.Edit(editMsg, "Выберете что хотите ввести.", createTagMenu())
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (bot *Bot) DeleteLastUserMsg(c telebot.Context) {
	fmt.Println("DeleteLastUserMsg", c.Message().ID)
	err := c.Delete()
	if err != nil {
		log.Println(err)
	}
}

func (bot *Bot) InitRoutes() {
	bot.tb.Handle("/start", bot.StartHandle)
	bot.tb.Handle(telebot.OnCallback, bot.Callbacks)
	bot.tb.Handle(telebot.OnText, bot.Messages)
}
