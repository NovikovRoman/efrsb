# EFRSB

> Библиотека для работы с сервисом получения сведений из ЕФРСБ.

[Документация v1.0.0 15.04.2024](https://fedresurs.ru/helps/bankrupt/Service_rest_1.0.pdf)

## Установка

```shell
go get github.com/NovikovRoman/efrsb
```

## Использование

### Авторизация

```go
import (
    …
    "github.com/NovikovRoman/efrsb"
    …
)

func main() {
    ctx := context.Background()
    // Для production efrsb.New(login, password, Prod())
    // Для development efrsb.New(login, password, Dev())
    // По-умолчанию production
    client := efrsb.New(login, password, Dev())

    var err error
    if err = client.Auth(ctx, cfg); err != nil {
        panic(err)
    }
    …
}
```

### Проверка токена авторизации

```go
…
var ok bool
if ok, err = client.IsActiveToken(); err != nil {
    panic(err)
}

fmt.Printf("Ключ активен: %t\n", ok)

var exp time.Time
if exp, err = client.TokenExpirationTime(); err != nil {
    panic(err)
}
fmt.Printf("Дата окончания действия ключа: %s\n", exp)
…
```

### Обновить токен авторизации

```go
if err = client.RefreshToken(); err != nil {
    panic(err)
}
```

### Поиск банкротов

```go
filter := efrsb.BankruptFilter{
    Type: TypePerson,
    Name: "Иванов",
}

results, err := client.Bankrupts(ctx, filter, offset, limit)
if err != nil {
    panic(err)
}

fmt.Printf("Всего: %d\n", results.Total)
```

### Поиск сообщений

```go
filter := efrsb.MessageFilter{
    BankruptGuid: []string{
        "a79f9366-32f4-ef38-b8b4-22253ffd47a9",
        "c8796d66-2a15-a47a-23a4-22824c0160e2",
    },
}

results, err := client.Messages(ctx, filter, offset, limit)
if err != nil {
    panic(err)
}

fmt.Printf("Всего: %d\n", results.Total)
```

### Получение сообщения

```go
message, err := client.Message(ctx, "deea9d05-9b04-44f5-9f55-64ef53108021")
```

### Получение архива файлов сообщения

```go
b, err := client.MessageFiles(ctx, "deea9d05-9b04-44f5-9f55-64ef53108021", true)
if err != nil {
    panic(err)
}
_ = os.WriteFile("files.zip", b, 0644)
```

### Получение списка связанных сообщений

```go
results, err := client.LinkedMessages(ctx, "deea9d05-9b04-44f5-9f55-64ef53108021")
```

### Поиск отчетов

```go
filter := efrsb.ReportFilter{
    Guid: []string{
        "38cd692f-ec25-4a37-af0b-3bb6e213f809",
    },
}

results, err := client.Reports(ctx, filter, offset, limit)
if err != nil {
    panic(err)
}

fmt.Printf("Всего: %d\n", results.Total)
```

### Получение отчета

```go
message, err := client.Report(ctx, "deea9d05-9b04-44f5-9f55-64ef53108021")
```

### Получение архива файлов отчета

```go
b, err := client.ResportFiles(ctx, "deea9d05-9b04-44f5-9f55-64ef53108021", true)
if err != nil {
    panic(err)
}
_ = os.WriteFile("files.zip", b, 0644)
```

### Получение списка связанных отчетов

```go
results, err := client.LinkedReports(ctx, "deea9d05-9b04-44f5-9f55-64ef53108021")
```
