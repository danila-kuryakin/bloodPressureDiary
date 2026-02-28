package bot

import (
	"bloodPressureDiary/bp-telegram-bot/internal/bot/constants"
	"bloodPressureDiary/bp-telegram-bot/internal/model"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"gopkg.in/telebot.v3"
)

func (bot *Bot) StartHandle(c telebot.Context) error {
	bot.state[c.Chat().ID] = constants.StateIdle
	return c.Send("Выберите раздел:", mainMenu())
}

func (bot *Bot) Callbacks(c telebot.Context) error {
	chatID := c.Sender().ID
	data := c.Callback().Data[1:]

	if bot.buffer[chatID] == nil {
		bot.buffer[chatID] = map[string]any{}
	}

	bot.buffer[chatID]["MenuMsg"] = c.Callback().Message.ID

	switch data {
	case constants.EventPressure:
		bot.state[chatID] = constants.StatePressure

		press, err := bot.api.ListPressure(strconv.FormatInt(chatID, 10))
		if err != nil {
			return err
		}
		retStr := ""

		if len(press) == 0 {
			retStr = "Меню давления.\nВыберите раздел."
		} else {
			retStr = "Меню давления.\nВаши теги:\n"
			for i, pres := range press {
				if i < 9 {
					retStr += fmt.Sprintf("%d)   %d-%d-%d | %s\n", i+1, pres.Systolic, pres.Diastolic, pres.Pulse, pres.CreatedAt.Format("15:04:05 02.01.06"))
				} else {
					retStr += fmt.Sprintf("%d) %d-%d-%d | %s\n", i+1, pres.Systolic, pres.Diastolic, pres.Pulse, pres.CreatedAt.Format("15:04:05 02.01.06"))
				}
				if i > 10 {
					break
				}
			}
			retStr += fmt.Sprintf("Всего %d\n", len(press))
			retStr += "Выберите раздел."
		}

		return c.Edit(retStr, crudPressureMenu())

	case constants.EventTag:
		bot.state[chatID] = constants.StateTag

		tags, err := bot.api.ListTags(strconv.FormatInt(chatID, 10))
		if err != nil {
			return err
		}
		retStr := ""

		if len(tags) == 0 {
			retStr = "Меню тегов.\nВыберите раздел."
		} else {
			retStr = "Меню тегов.\nВаши теги:\n"
			for _, tag := range tags {
				retStr += tag.Name + "\n"
			}
			retStr += "Выберите раздел."
		}

		return c.Edit(retStr, crudTagMenu())

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

		tags, err := bot.api.ListTags(strconv.FormatInt(chatID, 10))
		if err != nil {
			return err
		}

		eventMap := make(map[int]string)

		for i, tag := range tags {
			eventMap[i] = tag.Name
		}
		bot.buffer[chatID]["AddTags"] = eventMap

		return c.Edit("Введите тег, который хотите добавить.", addPressureTagMenu(eventMap))

		//bot.buffer[chatID] = map[string]any{}
		//return c.Edit("Выберете теги из имеющихся.", saveTagPressureMenu())

	case constants.EventPressureSave:
		fmt.Println("EventPressureSave",
			bot.buffer[chatID]["Systolic"],
			bot.buffer[chatID]["Diastolic"],
			bot.buffer[chatID]["Pulse"],
			bot.buffer[chatID]["Tags"])

		sys := bot.buffer[chatID]["Systolic"]

		if sys == nil || bot.buffer[chatID]["Diastolic"] == nil ||
			bot.buffer[chatID]["Pulse"] == nil || bot.buffer[chatID]["Tags"] == nil {
			return c.Respond()
		}

		press := model.BloodPressure{
			Systolic:  bot.buffer[chatID]["Systolic"].(int),
			Diastolic: bot.buffer[chatID]["Diastolic"].(int),
			Pulse:     bot.buffer[chatID]["Pulse"].(int),
			TagNames:  bot.buffer[chatID]["Tags"].([]string),
		}

		err := bot.api.CreatePressure(press, strconv.FormatInt(chatID, 10))
		if err != nil {
			return err
		}

		bot.state[chatID] = constants.StateIdle
		bot.buffer[chatID] = nil
		return c.Edit("Успешно сохранено.\nВыберите раздел:", mainMenu())

		// CRUD TAG
	case constants.EventTagCreate:
		bot.state[chatID] = constants.StateTagCreate
		return c.Edit("Введите тег, который хотите создать", createTagMenu())

	case constants.EventTagDelete:
		bot.state[chatID] = constants.StateTagDelete

		tags, err := bot.api.ListTags(strconv.FormatInt(chatID, 10))
		if err != nil {
			return err
		}

		eventMap := make(map[int]string)

		for i, tag := range tags {
			eventMap[i] = tag.Name
		}
		bot.buffer[chatID]["DeleteTags"] = eventMap

		return c.Edit("Введите тег, который хотите удалить.", deleteTagMenu(eventMap))

	case constants.EventTagSave:
		fmt.Println("EventPressureSave", bot.buffer[chatID]["TagsList"])

		tag := bot.buffer[chatID]["TagsList"].([]string)
		//userTag := model.UserTag{UserID: strconv.FormatInt(chatID, 10), Name: tag, CreatedAt: time.Now()}

		for _, name := range tag {
			tags := model.UserTag{
				UserID:    strconv.FormatInt(chatID, 10),
				Name:      name,
				CreatedAt: time.Now(),
			}

			err := bot.api.CreateTag(tags, strconv.FormatInt(chatID, 10))
			if err != nil {
				return err
			}
		}

		bot.state[chatID] = constants.StateIdle
		bot.buffer[chatID] = nil
		return c.Edit("Успешно сохранено.\nВыберите раздел:", mainMenu())

	case constants.EventTagName:
		bot.state[chatID] = constants.StateTagName
		return c.Edit("Введите название тега и нажмите Ввод", backMenu())

		// Back
	case constants.EventBack:
		switch bot.state[chatID] {
		case constants.StatePressure, constants.StateTag:
			bot.state[chatID] = constants.StateIdle
			return c.Edit("Выберите раздел.", mainMenu())

		case constants.StatePressureSys, constants.StatePressureDia,
			constants.StatePressurePulse, constants.StatePressureTags:
			bot.state[chatID] = constants.StatePressureCreate
			return c.Edit("Выберете что хотите ввести.", createPressureMenu())

		case constants.StatePressureCreate:
			bot.state[chatID] = constants.StatePressure
			return c.Edit("Меню давления.\nВыберите раздел.", crudPressureMenu())

		case constants.StateTagName:
			bot.state[chatID] = constants.StateTagCreate
			return c.Edit("Выберете что хотите ввести.", createTagMenu())

		case constants.StateTagCreate, constants.StateTagDelete:
			bot.state[chatID] = constants.StatePressure
			tags, err := bot.api.ListTags(strconv.FormatInt(chatID, 10))
			if err != nil {
				return err
			}
			retStr := ""

			fmt.Println(tags)
			if len(tags) == 0 {
				retStr = "Меню тегов.\nВыберите раздел."
			} else {
				retStr = "Меню тегов.\nВаши теги:\n"
				for _, tag := range tags {
					retStr += tag.Name + "\n"
				}
				retStr += "Выберите раздел."
			}

			return c.Edit(retStr, crudTagMenu())

		default:
			bot.state[chatID] = constants.StateIdle
			return c.Edit("Выберите раздел.", mainMenu())
		}
	default:
		// Проверяем, что это нужный тип события
		if strings.HasPrefix(data, "event_delete_tag_") {
			// Достаём индекс
			idStr := strings.TrimPrefix(data, "event_delete_tag_")

			id, err := strconv.Atoi(idStr)
			if err != nil {
				return c.Respond()
			}

			delTags := bot.buffer[chatID]["DeleteTags"].(map[int]string)
			// тут вызываешь API удаления тега
			err = bot.api.DeleteTag(delTags[id], strconv.FormatInt(chatID, 10))
			if err != nil {
				return err
			}
			delete(delTags, id)
			bot.buffer[chatID]["DeleteTags"] = delTags
			return c.Edit("Введите тег, который хотите удалить.", deleteTagMenu(delTags))
		} else if strings.HasPrefix(data, "event_add_tag_") {
			// Достаём индекс
			idStr := strings.TrimPrefix(data, "event_add_tag_")

			id, err := strconv.Atoi(idStr)
			if err != nil {
				return c.Respond()
			}

			delTags := bot.buffer[chatID]["AddTags"].(map[int]string)

			//fmt.Println("Нажата кнопка добавления тега:", id, delTags[id])

			if bot.buffer[chatID]["Tags"] == nil {
				bot.buffer[chatID]["Tags"] = []string{}
			}

			tags := bot.buffer[chatID]["Tags"].([]string)
			tags = append(tags, delTags[id])
			bot.buffer[c.Chat().ID]["Tags"] = tags

			retStr := bot.displayPressure(chatID)

			return c.Edit(retStr, addPressureTagMenu(delTags))
		}
	}

	bot.state[c.Chat().ID] = constants.StateIdle
	return c.Edit("Ой, что-то пошло не так.\nВыберите раздел:", mainMenu())
}

func (bot *Bot) displayPressure(chatID int64) string {
	retString := "Давление:\n"

	if bot.buffer[chatID]["Systolic"] != nil {
		retString += fmt.Sprintf("Верхнее: %d\n", bot.buffer[chatID]["Systolic"].(int))
	} else {
		retString += fmt.Sprintf("Верхнее: nil\n")
	}

	if bot.buffer[chatID]["Diastolic"] != nil {
		retString += fmt.Sprintf("Нижнее: %d\n", bot.buffer[chatID]["Diastolic"].(int))
	} else {
		retString += fmt.Sprintf("Нижнее: nil\n")
	}

	if bot.buffer[chatID]["Pulse"] != nil {
		retString += fmt.Sprintf("Пульс: %d\n", bot.buffer[chatID]["Pulse"].(int))
	} else {
		retString += fmt.Sprintf("Пульс: nil\n")
	}

	if bot.buffer[chatID]["Tags"] != nil {
		retString += fmt.Sprintf("Теги: ")

		fmt.Println(bot.buffer[chatID]["Tags"])
		for i, tag := range bot.buffer[chatID]["Tags"].([]string) {
			if i == 0 {
				retString += fmt.Sprintf("%s", tag)
			} else {
				retString += fmt.Sprintf(", %s", tag)
			}
		}
	} else {
		retString += fmt.Sprintf("Теги: nil\n")
	}

	return retString
}

func (bot *Bot) displayTags(chatID int64) string {
	retString := "Теги:\n"

	if bot.buffer[chatID]["TagsList"] != nil {
		for i, tag := range bot.buffer[chatID]["TagsList"].([]string) {

			retString += fmt.Sprintf("%d) %s\n", i, tag)
		}
	} else {
		retString += fmt.Sprintf("nil\n")
	}
	return retString
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
		msgInt, err := strconv.Atoi(msg)
		if err != nil {
			return err
		}
		bot.buffer[c.Chat().ID]["Systolic"] = msgInt

		bot.state[chatID] = constants.StatePressureCreate
		bot.DeleteLastUserMsg(c)
		msgID := bot.buffer[c.Chat().ID]["MenuMsg"].(int)

		editMsg := &telebot.Message{
			ID:   msgID,
			Chat: c.Chat(),
		}

		retStr := bot.displayPressure(chatID)
		_, err = bot.tb.Edit(editMsg, retStr, createPressureMenu())
		if err != nil {
			return err
		}
		return nil

	case constants.StatePressureDia:
		msgInt, err := strconv.Atoi(msg)
		if err != nil {
			return err
		}
		bot.buffer[c.Chat().ID]["Diastolic"] = msgInt
		bot.state[chatID] = constants.StatePressureCreate
		bot.DeleteLastUserMsg(c)
		msgID := bot.buffer[c.Chat().ID]["MenuMsg"].(int)

		editMsg := &telebot.Message{
			ID:   msgID,
			Chat: c.Chat(),
		}

		retStr := bot.displayPressure(chatID)
		_, err = bot.tb.Edit(editMsg, retStr, createPressureMenu())
		if err != nil {
			return err
		}
		return nil

	case constants.StatePressurePulse:
		msgInt, err := strconv.Atoi(msg)
		if err != nil {
			return err
		}
		bot.buffer[c.Chat().ID]["Pulse"] = msgInt
		bot.state[chatID] = constants.StatePressureCreate
		bot.DeleteLastUserMsg(c)
		msgID := bot.buffer[c.Chat().ID]["MenuMsg"].(int)

		editMsg := &telebot.Message{
			ID:   msgID,
			Chat: c.Chat(),
		}

		retStr := bot.displayPressure(chatID)
		_, err = bot.tb.Edit(editMsg, retStr, createPressureMenu())
		if err != nil {
			return err
		}
		return nil

	//case constants.StatePressureTags:
	//	if bot.buffer[chatID]["Tags"] == nil {
	//		bot.buffer[chatID]["Tags"] = []string{}
	//	}
	//
	//	tags := bot.buffer[chatID]["Tags"].([]string)
	//	tags = append(tags, msg)
	//	bot.buffer[c.Chat().ID]["Tags"] = tags
	//	bot.state[chatID] = constants.StatePressureCreate
	//	bot.DeleteLastUserMsg(c)
	//	msgID := bot.buffer[c.Chat().ID]["MenuMsg"].(int)
	//
	//	editMsg := &telebot.Message{
	//		ID:   msgID,
	//		Chat: c.Chat(),
	//	}
	//
	//	retStr := bot.displayPressure(chatID)
	//	_, err := bot.tb.Edit(editMsg, retStr, createPressureMenu())
	//	if err != nil {
	//		return err
	//	}
	//	return nil

	case constants.StateTagCreate:
		if bot.buffer[chatID]["TagsList"] == nil {
			bot.buffer[chatID]["TagsList"] = []string{}
		}

		tags := bot.buffer[chatID]["TagsList"].([]string)
		tags = append(tags, msg)
		bot.buffer[chatID]["TagsList"] = tags
		bot.state[chatID] = constants.StateTagCreate
		bot.DeleteLastUserMsg(c)
		msgID := bot.buffer[c.Chat().ID]["MenuMsg"].(int)

		editMsg := &telebot.Message{
			ID:   msgID,
			Chat: c.Chat(),
		}
		retStr := bot.displayTags(chatID)
		retStr += "\nВведите еще теги"
		_, err := bot.tb.Edit(editMsg, retStr, createTagMenu())
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
	//bot.tb.Handle(telebot.OnCallback, bot.deleteTagCallbacks)
	bot.tb.Handle(telebot.OnText, bot.Messages)
}
