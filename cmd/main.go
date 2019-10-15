package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/tarm/serial"
)

func main() {

	config := &serial.Config{
		Name:        "/dev/ttyUSB0",
		Baud:        9600,
		ReadTimeout: 1,
		Size:        8,
	}

	s, err := serial.OpenPort(config)
	if err != nil {
		fmt.Printf("Error open serial port %v", err)
		log.Fatalf("can't open serial port %v", err)
	}

	rand.Seed(42)
	event := []string{
		"ПОЖАР           ",
		"ВНИМАНИЕ        ",
		"ТРЕВОГА         ",
		"ТИХАЯ ТРЕВОГА   ",
		"ТРЕВОГА ВХОДА   ",
		"ОБРЫВ ШС        ",
		"КОРОТКОЕ ЗАМЫКАН",
		"НЕИСПРАВНОСТЬ   ",
		"ВОССТАНОВЛЕНИЕ  ",
		"ОШИБКА ПАРАМ.ШС ",
		"ОТКЛЮЧЕН ШС     ",
		"В0ССТ.30НЫ      ",
		"НЕВЗЯТИЕ        ",
		"СБРОС ТРЕВОГИ ШС",
		"ВЗЯТ ШС         ",
		"СНЯТ ШС         ",
		"ВЗЛОМ КОРПУСА   ",
		"ВОССТ.КОРПУСА   ",
		"СРАБОТКА ДАТЧИКА",
		"НЕОБХ.ОБСЛ      ",
		"ПОВЫШ.ТЕМПЕР.   ",
		"ПОНИЖ.ТЕМПЕР.   ",
		"НОРМА ТЕМПЕР.   ",
		"НЕИСП.ТЕРМОМЕТРА",
		"АВАРИЯ ПИТАНИЯ  ",
		"ВОССТ.ПИТАНИЯ   ",
		"АВАРИЯ БАТАРЕИ  ",
		"ВОССТ.БАТАРЕИ   ",
		"АВАРИЯ 220В     ",
		"ВОССТ.220В      ",
		"СБРОС ПРИБОРА   ",
		"ПОТЕРЯН ПРИБОР  ",
		"ОБНАРУЖЕН ПРИБОР",
		"ОТКЛ.ВЕТВИ RS485",
		"ВСТ.ВЕТВИ RS485 ",
		"КЗ ДПЛС         ",
		"ВОССТАНОВЛ.ДПЛС ",
		"НАРУШ.ТЕХНОЛ.ШС ",
		"ВОССТ.ТЕХНОЛ.ШС ",
		"НЕНОРМАШС       ",
		"ВОССТАНОВЛ.ШС   ",
		"РАЗДЕЛ ВЗЯТ     ",
		"РАЗДЕЛ СНЯТ     ",
		"ЗАПРОС ВЗЯТИЯ   ",
		"ЗАПРОС СНЯТИЯ   ",
		"ИДЕНТИФИКАЦИЯ ХО",
		"ОБРЫВ ВЫХОДА    ",
		"КЗ ВЫХОДА       ",
		"ОТКЛ.ВЫХОДА     ",
		"ВОССТ.ВЫХОДА    ",
		"УСПЕШНЫЙ ЗАПУСК ",
		"НЕУСП.ЗАПУСК    ",
		"АВТОМАТИКА ВЫКЛ.",
		"АВТОМАТИКА ВКЛ. ",
		"ПУСКАСПТ        ",
		"ЗАДЕРЖКА ЗАПУСКА",
		"БЛОКИР.ПУСКА    ",
		"АВАРИЙНЫЙ ПУСК  ",
		"СРАБАТЫВАНИЕ СДУ",
		"ОТКАЗ СДУ       ",
		"ПРОХОД          ",
		"ДОСТУП ОТКЛОНЕН ",
		"ДОСТУП ПРЕДОСТАВ",
		"ДОСТУП ЗАПРЕЩЕН ",
		"ДОСТУП ЗАКРЫТ   ",
		"ДОСТУП ОТКРЫТ   ",
		"ДОСТУП ВОССТАНОВ",
		"ДВЕРЬ ЗАБЛОКИР. ",
		"ДВЕРЬ РАЗБЛОКИР.",
		"ДВЕРЬ ВЗЛОМАНА  ",
		"ПРОГРАММИРОВАНИЕ",
		"НЕИСПР.ТЕЛ.ЛИНИИ",
		"ЖУРНАЛ ЗАПОЛНЕН ",
		"ЖУРНАЛ ПЕРЕПОЛН.",
		"ЗАПУСК ТЕСТА    ",
		"РЕАКЦИЯ         ",
		"ВКЛЮЧЕНИЕ ПУЛЬТА",
		"ВКЛ. ПРИНТЕРА   ",
		"ВЫКЛ. ПРИНТЕРА  ",
		"ИЗМ.ДАТЫ        ",
		"ИЗМ.ВРЕМЕНИ     ",
		"ДАТА            ",
		"ОТМЕТКА ВРЕМЕНИ ",
		"ПОДКЛЮЧЕН ВЫХОД ",
		"ОТКЛЮЧЕН ВЫХОД  ",
		"ОТКЛЮЧЕН        ",
		"ПОДКЛЮЧЕН       ",
		"ВКЛ.ТЕСТ ИЗВЕЩ. ",
		"ВЫКЛ.ТЕСТ ИЗВЕЩ.",
		"ТЕСТ ИЗВЕЩАТЕЛЯ ",
		"НЕИСП.ТЕРМОМЕТРА",
		"ВКЛ. НАСОСА     ",
		"ВЫКЛ. НАСОСА    ",
		"АКБ РАЗРЯЖЕНА   ",
		"ОШИБКА ТЕСТА АКБ",
		"ПЕРЕГРУЗКА РИП  ",
		"УСТР. ПЕРЕГР.РИП",
		"АВАРИЯ ДПЛС     ",
		"ОШИБКА В ОТВЕТЕ ",
		"НЕУСТ. СВЯЗЬ    ",
		"ВОССТ. ТЕЛ.ЛИНИИ",
		"ОШИБКА ТЕСТА    ",
		"ПУСКАУП         ",
		"ТУШЕНИЕ         ",
		"НЕУДАЧНЫЙ ПУСК  ",
		"ОТМЕНА ПУСКА    ",
		"ЗАДЕРЖКА ВЗЯТИЯ ",
		"ВКЛ. КОНТРОЛЬ ШС",
		"НАРУШ.ТЕХНОЛ.ШС ",
		"ПУСК РО         ",
		"ОТМЕНА ПУСКА РО ",
		"РАБОЧ.СОСТОЯНИЕ ",
		"ИСХОДН.СОСТОЯНИЕ",
		"ОШИБ. УПРАВЛЕНИЯ",
		"ОШИБ. УПРАВЛЕНИЯ",
		"ОШИБКА КОНТРОЛЯ ",
		"ПОВЫШ. УРОВНЯ   ",
		"ПОНИЖ. УРОВНЯ   ",
		"НОРМА УРОВНЯ    ",
		"АВ.ПОВЫШ.УРОВНЯ ",
		"АВ.ПОНИЖ.У РОВНЯ",
		"НЕИСПР. ЗУ РИП  ",
		"ВОССТ. ЗУ РИП   ",
		"ОТКЛЮЧЕН РИП    ",
		"ВКЛЮЧЕН РИП     ",
		"НЕТ СВЯЗИ ДПЛС1 ",
		"НЕТ СВЯЗИ ДПЛС2 ",
		"УСТ.СВЯЗЬ ДПЛС1 ",
		"УСТ.СВЯЗЬ ДПЛС2 ",
		"НЕИЗВ.УСТР0ЙСТВ0",
		"ВОССТ. ВЫХОДА   ",
		"РЕЛЕ ВКЛЮЧЕНО   ",
		"РЕЛЕ ВКЛ.ПРЕРЫВ.",
		"ОТМЕТКА НАРЯДА  ",
	}

	rand.Seed(42)
	section := []string{
		"ДИП             ",
		"ИПР/ип212-45(2) ",
		"ип-212-45(1)    ",
		"ПС              ",
		"АР-2            ",
		"ОС              ",
		"ИК              ",
		"ШИК             ",
	}

	for {
		d := rand.Intn(10)
		t := time.Now().Format("01.02 15:04:05")
		e := event[rand.Intn(len(event))]
		sec := section[rand.Intn(len(section))]
		b := []byte(fmt.Sprintf("|%s|%s|  %v    |   %v|%s|№ ПАРОЛЯ: 7     |\r\n", t, e, d, d, sec))
		_, err := s.Write(b)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("|%s|%s|  %v    |   %v|%s|№ ПАРОЛЯ: 7     |\r\n", t, e, d, d, sec)
		time.Sleep(time.Second * 10)
	}
}