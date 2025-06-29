go install github.com/end1essrage/indigo-core/indigo-cli@latest

Создать новый проект
indigo-cli new --cache=redis --storage=mongo --output=./my-bot

Параметры:
   --cache     : Тип кэша (redis/memory)
   --storage   : Тип хранилища (mongo/file)
   --output (-o): Путь для создания проекта
