# HW2: MTLS, Ingress and K8S

<img width="482" alt="Снимок экрана 2023-05-24 в 00 15 48" src="https://github.com/star1can/sberops_s2023/assets/45429125/959293c0-6d25-490b-b204-4632803ccfa1">


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
3. Вернитесь на один уровень выше
4. Выполните скрипт `test.sh`

Если все сделано правильно, Вы увидите следующее:

<img width="343" alt="Снимок экрана 2023-05-24 в 00 18 34" src="https://github.com/star1can/sberops_s2023/assets/45429125/8d9292eb-3af5-4633-8850-b8a1b4c23988">


### Очистка
В папке `kube`:
- Для очистки рабочего пространства выполните скрипт `clear_workspace.sh`;
- Для удаление кластера minikube, созданного скриптом ранее, можно воспользоваться скриптом `delete_minikube.sh`.
