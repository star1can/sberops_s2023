# HW3: MTLS, Egress and K8S

<img width="470" alt="Снимок экрана 2023-05-25 в 01 10 21" src="https://github.com/star1can/sberops_s2023/assets/45429125/755df194-d2e4-4b4d-9d31-1f139e3e0609">

# Описание

Основу данного проекта представляют два сервиса: `nginx` и `pokemon-service`:

`nginx` - выступает в роли pass_proxy и просто пропускает через себя трафик, направляя на нужные внутренние ручки

`pokemon-service` - полноценный сервис, написаный на Golang. Делает запрос за пределы кластера и возвращает список всех покемонов.

На входе в кластер стоит ingress-gateway, на выходе - egress-gateway, что позволяет полностью контролировать входящий и исходящий трафик.

# Доступные ручки
На текущий момент есть лишь одна ручка:
- `/pokemons/all` - возвращает список всех доступных покемонов

# Инструкция по запуску

## Пререквизиты 

1. Склонируйте репозиторий
2. Запустите Docker
3. При необходимости создайте minikube кластер при помощи скрипта `kube/start_minikube.sh`

## Инструкция

### Запуск
1. Перейдите в папку `kube`
2. Запустите скрипт `run.sh` и следуйте инструкциями. ВНИМАНИЕ: ПО ЗАВЕРШЕНИЮ СКРИПТА ВКЛАДКА ТЕРМИНАЛА БУДЕТ ЗАБЛОКИРОВАНА!
3. В новой вкладке терминала выполните скрипт `get_ip.sh`
4. Скопируйте IP-шник из колонки **EXTERNAL-IP** 
5. В браузере выполните запрос по адресу `<IP_FROM_PUNKT_4>/pokemons/all`

Если все сделано правильно, Вы увидите список доступных покемонов:

<img width="177" alt="Снимок экрана 2023-05-25 в 01 09 54" src="https://github.com/star1can/sberops_s2023/assets/45429125/d6697abf-909a-4db1-b939-35916b3a0f19">

### Очистка
В папке `kube`:
- Для очистки рабочего пространства выполните скрипт `clear_workspace.sh`;
- Для удаление кластера minikube, созданного скриптом ранее, можно воспользоваться скриптом `delete_minikube.sh`.
