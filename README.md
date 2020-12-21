Go field selection

Пакет для выбора поле из стрыктуры на освновании списка шаблонов

Пример шаблонов:
```
[]string{"Id", "FirstName", "Age", "Roles.Id","Roles.Images.Url", "Roles.Images.ImageGroups.Id", "Image.Url", "Birthday"}
```

Для запуска:
```
go run main.go | json_pp
```