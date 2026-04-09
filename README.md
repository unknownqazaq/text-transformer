# Text Transformer

Консольное приложение на Go для постобработки текста по набору встроенных правил.

Программа читает текст из входного файла, применяет команды-теги внутри текста и сохраняет результат в выходной файл.

## Возможности

- Конвертация числа перед тегом:
	- `(hex)` -> из шестнадцатеричной системы в десятичную
	- `(bin)` -> из двоичной системы в десятичную
- Изменение регистра слов перед тегом:
	- `(up)` -> верхний регистр для 1 слова
	- `(low)` -> нижний регистр для 1 слова
	- `(cap)` -> капитализация для 1 слова
	- `(up, N)` -> верхний регистр для последних N слов
	- `(low, N)` -> нижний регистр для последних N слов
	- `(cap, N)` -> капитализация для последних N слов
- Нормализация пунктуации:
	- удаление лишнего пробела перед `.,!?:;`
	- добавление пробела после знака препинания, если он нужен
- Нормализация одинарных кавычек:
	- `' text '` -> `'text'`
- Автозамена артикля:
	- `a` -> `an`, `A` -> `An` перед гласными и `h`

## Требования

- Go 1.25+

## Запуск

1. Подготовьте входной файл с текстом, например `sample.txt`.
2. Выполните команду:

	 go run . sample.txt output.txt

3. Результат будет записан в `output.txt`.

Если аргументы переданы неверно, приложение выводит подсказку:

использование: go run . input.txt output.txt

## Пример

Вход (`sample.txt`):

1E (hex) files were added. 10 (bin) Years it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

Выход (`output.txt`):

30 files were added. 2 Years It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.

## Как устроен проект

- `main.go`:
	- проверка аргументов CLI
	- чтение входного файла
	- запуск обработки
	- запись результата
- `internal/reader.go`: чтение текстового файла
- `internal/processor.go`: основной пайплайн обработки и разбор тегов
- `internal/rules.go`: правила преобразования (регистр, числа, пунктуация, кавычки, артикли)
- `internal/writer.go`: запись результата в файл

## Важная деталь по формату тегов

Текст разбирается по словам через разделение по пробелам. Для тегов с числом ожидается формат в два токена:

- `(up, 3)`
- `(low, 2)`
- `(cap, 5)`

То есть запятая идет вместе с тегом, а число с закрывающей скобкой - отдельным словом.
