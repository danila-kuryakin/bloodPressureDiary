package bot

import (
	"bp-telegram-bot/internal/bot/constants"
	"bp-telegram-bot/internal/model"
	"fmt"
	"log"
	"sort"
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

		pressures, err := bot.api.ListPressure(strconv.FormatInt(chatID, 10))
		if err != nil {
			return err
		}

		keys := make([]int, 0, 10)
		pressuresList := map[int]model.BloodPressure{}
		lenKeys := 10
		if len(pressures) < 10 {
			lenKeys = len(pressures)
			keys = make([]int, 0, lenKeys)
		}
		for i := range lenKeys {
			keys = append(keys, i)
			pressuresList[i] = pressures[i]
		}
		bot.buffer[chatID]["PressureList"] = pressuresList

		pressureStr := bot.pressureList(pressuresList, keys)

		return c.Edit(pressureStr, crudPressureMenu())
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
	case constants.EventPressureDelete:
		pressureListMap := bot.buffer[chatID]["PressureList"].(map[int]model.BloodPressure)

		keys := make([]int, 0, len(pressureListMap))
		for k := range pressureListMap {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		bot.state[chatID] = constants.StatePressureDelete

		pressureStr := bot.pressureList(pressureListMap, keys)
		return c.Edit(pressureStr, deletePressureMenu(keys))
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

		return c.Edit("Введите тег, который хотите добавить.", addPressureTagMenu(eventMap, false))
	case constants.EventPressureSave:
		log.Println("chatID:",
			chatID,
			"| EventPressureSave",
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
		if bot.buffer[chatID]["TagsList"] == nil {
			return c.Respond()
		}
		tag := bot.buffer[chatID]["TagsList"].([]string)

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
	case constants.EventEmpty:
		return c.Respond()
		// Back
	case constants.EventBack:
		switch bot.state[chatID] {
		case constants.StatePressure, constants.StateTag:
			bot.state[chatID] = constants.StateIdle
			return c.Edit("Выберите раздел.", mainMenu())

		case constants.StatePressureSys, constants.StatePressureDia,
			constants.StatePressurePulse, constants.StatePressureTags:
			bot.state[chatID] = constants.StatePressureCreate
			retStr := bot.displayPressure(chatID)
			return c.Edit(retStr, createPressureMenu())

		case constants.StatePressureCreate:
			pressureKeys := bot.buffer[chatID]["PressureKeys"].([]int)
			pressureListMap := bot.buffer[chatID]["PressureList"].(map[int]model.BloodPressure)

			bot.state[chatID] = constants.StatePressure

			fmt.Println(len(pressureListMap))

			pressureStr := bot.pressureList(pressureListMap, pressureKeys)
			return c.Edit(pressureStr, crudPressureMenu())

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
		}
		if strings.HasPrefix(data, "event_add_pressure_tag_") {
			// Достаём индекс
			idStr := strings.TrimPrefix(data, "event_add_pressure_tag_")

			id, err := strconv.Atoi(idStr)
			if err != nil {
				return c.Respond()
			}

			eventMap := bot.buffer[chatID]["AddTags"].(map[int]string)

			if bot.buffer[chatID]["Tags"] == nil {
				bot.buffer[chatID]["Tags"] = []string{}
			}

			tags := bot.buffer[chatID]["Tags"].([]string)
			tags = append(tags, eventMap[id])
			bot.buffer[c.Chat().ID]["Tags"] = tags

			if bot.state[chatID] == constants.StatePressureTagsFast {
				retStr := constants.StrFastAddingPressure + constants.StrTags + bot.displayPressure(chatID)
				return c.Edit(retStr, addPressureTagMenu(eventMap, true))

			} else {
				retStr := bot.displayPressure(chatID)
				return c.Edit(retStr, addPressureTagMenu(eventMap, false))
			}
		}
		if strings.HasPrefix(data, "event_delete_pressure_") {
			// Достаём индекс
			idStr := strings.TrimPrefix(data, "event_delete_pressure_")

			id, err := strconv.Atoi(idStr)
			if err != nil {
				return c.Respond()
			}

			// тут вызываешь API удаления тега
			if bot.api.DeletePressure(id, strconv.FormatInt(chatID, 10)) != nil {
				return err
			}

			pressureList := bot.buffer[chatID]["PressureList"].(map[int]model.BloodPressure)

			delete(pressureList, id)
			bot.buffer[chatID]["PressureList"] = pressureList

			keys := make([]int, 0, len(pressureList))
			for k := range pressureList {
				keys = append(keys, k)
			}
			sort.Ints(keys)

			pressureStr := bot.pressureList(pressureList, keys)
			return c.Edit(pressureStr, deletePressureMenu(keys))
		} else {
			return c.Respond()
		}
	}
}

func (bot *Bot) pressureList(press map[int]model.BloodPressure, keys []int) string {
	retStr := ""
	lenPress := len(press)
	if lenPress == 0 {
		retStr = "Меню давления.\nВыберите раздел."
	} else {
		retStr = "Меню давления.\nВаши измерения:\n"
		for _, key := range keys {

			if key < 9 {
				retStr += fmt.Sprintf("%d)   %d-%d-%d | %s | ", key+1, press[key].Systolic, press[key].Diastolic, press[key].Pulse, press[key].CreatedAt.Format("15:04:05 02.01.06"))
			} else {
				retStr += fmt.Sprintf("%d) %d-%d-%d | %s | ", key+1, press[key].Systolic, press[key].Diastolic, press[key].Pulse, press[key].CreatedAt.Format("15:04:05 02.01.06"))
			}

			if press[key].TagNames != nil {
				retStr += "\t"
				for j, tag := range press[key].TagNames {
					if j < len(press[key].TagNames)-1 {
						retStr += fmt.Sprintf("%s, ", tag)
					} else {
						retStr += fmt.Sprintf("%s\n", tag)
					}
				}
			}
		}
		if len(press) <= constants.MAX_VIEW_PRESSURE {
			retStr += fmt.Sprintf("Показано %d из %d\n", len(press), len(press))
		} else {
			retStr += fmt.Sprintf("Показано %d из %d\n", constants.MAX_VIEW_PRESSURE, len(press))
		}
		retStr += "Выберите раздел."
	}
	return retStr
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
	switch bot.state[chatID] {
	case constants.StatePressure, constants.StatePressureCreate:
		return bot.addPressureInBufferFast(
			c,
			msg,
			constants.Systolic,
			constants.StrErrPressureSys,
			constants.StatePressureDiaFast,
			constants.StrFastAddingPressure+constants.StrDia,
			backMenu())

	case constants.StatePressureDiaFast:
		return bot.addPressureInBufferFast(
			c,
			msg,
			constants.Diastolic,
			constants.StrErrPressureDia,
			constants.StatePressurePulseFast,
			constants.StrFastAddingPressure+constants.StrPulse,
			backMenu())

	case constants.StatePressurePulseFast:
		tags, err := bot.api.ListTags(strconv.FormatInt(chatID, 10))
		if err != nil {
			return err
		}

		eventMap := make(map[int]string)

		for i, tag := range tags {
			eventMap[i] = tag.Name
		}
		bot.buffer[chatID]["AddTags"] = eventMap
		return bot.addPressureInBufferFast(
			c,
			msg,
			constants.Pulse,
			constants.StrErrPressurePulse,
			constants.StatePressureTagsFast,
			constants.StrFastAddingPressure+constants.StrTags,
			addPressureTagMenu(eventMap, true))

	case constants.StatePressureTagsFast:
		return c.Respond()

	case constants.StatePressureSys:
		return bot.addPressureInBufferFast(
			c,
			msg,
			constants.Systolic,
			constants.StrErrPressureSys,
			constants.StatePressureCreate,
			"",
			createPressureMenu())

	case constants.StatePressureDia:
		return bot.addPressureInBufferFast(
			c,
			msg,
			constants.Diastolic,
			constants.StrErrPressureDia,
			constants.StatePressureCreate,
			"",
			createPressureMenu())

	case constants.StatePressurePulse:
		return bot.addPressureInBufferFast(
			c,
			msg,
			constants.Pulse,
			constants.StrErrPressurePulse,
			constants.StatePressureCreate,
			"",
			createPressureMenu())

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

func (bot *Bot) addPressureInBuffer(c telebot.Context, msg, pressureType, invalidString string) error {
	chatID := c.Sender().ID
	msgInt, err := strconv.Atoi(msg)
	if err != nil {
		return err
	}

	isValid := IsTwoOrThreeDigits(msg)
	if isValid {
		bot.buffer[chatID][pressureType] = msgInt
		bot.state[chatID] = constants.StatePressureCreate
	}

	bot.DeleteLastUserMsg(c)
	msgID := bot.buffer[chatID]["MenuMsg"].(int)
	editMsg := &telebot.Message{
		ID:   msgID,
		Chat: c.Chat(),
	}

	if isValid {
		retStr := bot.displayPressure(chatID)
		_, err = bot.tb.Edit(editMsg, retStr, createPressureMenu())
		if err != nil {
			return err
		}
		return nil
	}
	_, err = bot.tb.Edit(editMsg, invalidString, backMenu())
	if err != nil {
		return err
	}
	return nil
}

func (bot *Bot) addPressureInBufferFast(c telebot.Context, msg, pressureType, invalidString string, nextState constants.State, fastStrMode string, successMenu interface{}) error {
	chatID := c.Sender().ID
	msgInt, err := strconv.Atoi(msg)
	if err != nil {
		return err
	}

	isValid := IsTwoOrThreeDigits(msg)
	if isValid {
		bot.buffer[chatID][pressureType] = msgInt
		bot.state[chatID] = nextState
	}

	bot.DeleteLastUserMsg(c)
	msgID := bot.buffer[chatID]["MenuMsg"].(int)
	editMsg := &telebot.Message{
		ID:   msgID,
		Chat: c.Chat(),
	}

	retStr := ""

	if fastStrMode == "" {
		retStr += bot.displayPressure(chatID)
	} else {
		retStr += fastStrMode
		retStr += bot.displayPressure(chatID)
	}
	if isValid {
		_, err = bot.tb.Edit(editMsg, retStr, successMenu)
		if err != nil {
			return err
		}
		return nil
	} else {
		_, err = bot.tb.Edit(editMsg, invalidString+retStr, backMenu())
		if err != nil {
			return err
		}
		return nil
	}
}

func (bot *Bot) DeleteLastUserMsg(c telebot.Context) {
	//fmt.Println("DeleteLastUserMsg", c.Message().ID)
	err := c.Delete()
	if err != nil {
		log.Printf("Error on DeleteLastUserMsg: %v", err)
	}
}

func (bot *Bot) InitRoutes() {
	bot.tb.Handle("/start", bot.StartHandle)
	bot.tb.Handle(telebot.OnCallback, bot.Callbacks)
	bot.tb.Handle(telebot.OnText, bot.Messages)
}
