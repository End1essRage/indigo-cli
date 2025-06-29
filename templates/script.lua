-- Основная функция обработки (вызывается автоматически)
local function handle()
    -- 1. Пример логирования
    log("Скрипт запущен для пользователя: " .. ctx.user.name)
    
    -- 2. Пример работы с секретами
    -- local api_key = reveal("API_KEY")
    -- log("Получен секрет: " .. api_key)
    
    -- 3. Пример работы с хранилищем
    --[[
    local user_data = {
        name = ctx.user.name,
        last_active = os.date("%Y-%m-%d")
    }
    storage_save("users", ctx.user.id, user_data)
    
    local loaded_data = storage_load("users", ctx.user.id)
    log("Данные пользователя: " .. loaded_data.name)
    
    local items = storage_load_array("products", "list")
    log("Загружено элементов: " .. #items)
    --]]
    
    -- 4. Пример работы с кэшем
    --[[
    cache_set("last_action_" .. ctx.user.id, "login")
    local action = cache_get("last_action_" .. ctx.user.id)
    log("Последнее действие: " .. action)
    --]]
    
    -- 5. Пример отправки сообщения
    send_message(ctx.chat_id, "Привет, " .. ctx.user.name .. "! Добро пожаловать!")
    
    -- 6. Пример отправки клавиатуры
    --[[
    local menu = {
        {"Кнопка 1", "Кнопка 2"},
        {"Меню"}
    }
    send_keyboard(ctx.chat_id, "Выберите действие:", menu)
    --]]
    
    -- 7. Пример HTTP-запроса
    --[[
    local headers = {
        ["Authorization"] = "Bearer " .. reveal("API_TOKEN"),
        ["Content-Type"] = "application/json"
    }
    
    -- GET-запрос
    local res_get = http_get("https://api.example.com/data", headers)
    log("GET статус: " .. res_get.status)
    log("GET ответ: " .. res_get.body)
    
    -- POST-запрос
    local data = {
        user = ctx.user.id,
        action = "start"
    }
    local res_post = http_post("https://api.example.com/log", json_encode(data), headers)
    log("POST статус: " .. res_post.status)
    --]]
    
    -- 8. Пример работы с JSON
    --[[
    local json_str = json_encode({key = "value", arr = {1, 2, 3}})
    log("JSON строка: " .. json_str)
    --]]
end

-- Вызов основной функции
handle()