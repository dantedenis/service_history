# service_history  

Данный сервис периодически опрашивает сервис генерации и сохраняет полученные значение в базу данных.  
Имеется АПИ-метод(GET /health) для проверки статуса работы сервиса, возвращает 200 - если сервис доступен и работает.  

Для получения значений в необходимый промежуток времени для конкретной валютной паре реализован RPC метод  
`internal/proto/server.proto`

Запуск с помощью docker-compose:  
    `make create_network` - если не настроен, необходим для взаимодействия сервисов  
    `make run` - старт сервиса  
    `make restart` - рестарт сервиса  
    `make kill` - стоп сервиса и удаление контейнеров  
    `make logs` - вывод логов  


Запуск тестов:  
    `make test` - тесты  
    `make coverage` - генерация отчета о покрытии  
    `make lint` - проверка линтером  

proto-generate:  
    `protoc -I internal/app/proto server.proto --go-grpc_out=internal/app --go_out=internal/app --go-grpc_opt=require_unimplemented_servers=false`

mocken:  
    `mockgen -source={PATH_INTERFACE} -destination={PATH_DEST}`