# **weatherAPI - Погода из API**
Программа, которая получает данные о погоде из внешнего API.

⚠️ Важно: Этот проект создан в образовательных целях. Он предназначен для изучения архитектуры и принципов написания микросервисов без UI-окружения.
В коде могут быть ошибки, так как автор проекта находится на этапе обучения бэкенд-разработке.


## Технологии
* Golang (version go1.22.2)
* Redis
* Makefile
* UNIX-подобная операционная система (Linux, macOS)
  
## Установка
Клонируйте и соберите проект:

```
git clone https://github.com/JacketMiaimi/weatherApi.git
cd weatherApi
make run
```
Также надо установить Redis. <br>
<a href="https://timeweb.cloud/tutorials/redis/ustanovka-i-nastrojka-redis-dlya-raznyh-os">Гайд</a> 

## Настройка
В **configs.yaml** можно поменять:

* **env** - место окружение <br>
* **http_server** - настройка сервера <br>
* **address** - задаем адресс<br>
* **timeout** - время ожидание<br>
* **idle_timeout** - ожидание бездействия<br> 

В файле **.env** нужно указать: <br>

* **PATH_CONFIG** - путь к **.yaml** конфигу <br>
* **PASSWORD_REDIS** - пароль от redis

## Как пользоваться 
Чтобы получить погоду, используйте формат запроса:
Для этого нужно указать: <br>

```
address/get/{city}
```
<br>

Пример
```
localhost:8080/get/Moscow
```
