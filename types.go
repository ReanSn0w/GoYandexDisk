package GoYandexDisk

import (
	"time"
)

//Link Объект содержит URL для запроса метаданных ресурса.
type Link struct{
	Href string `json:"href"`
	Method string `json:"method"`
	Templated string `json:"templated"`
}

//Resource Описание ресурса, мета-информация о файле или папке. Включается в ответ на запрос метаинформации.
type Resource struct {
	Publik_key string `json:"publik_key"`
	Embeded interface{} `json:"_embeded"`
	Name string `json:"name"`
	Created time.Time `json:"created"`
	CustomProperties map[string]string `json:"custom_properties"`
	PublicURL string `json:"public_url"`
	OriginPATH string `json:"origin_path"`
	Modified string `json:"modified"`
	Path string `json:"path"`
	Md5 string `json:"md_5"`
	Type string `json:"type"`
	Mime_type string `json:"mime_type"`
	Size int `json:"size"`
}

//ResourceList Список ресурсов, содержащихся в папке. Содержит объекты Resource и свойства списка.
type ResourceList struct {
	Sort string `json:"sort"`
	Public_key string `json:"public_key"`
	Items []Resource `json:"items"`
	Path string `json:"path"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	Total int `json:"total"`
}

//FileResourceList Плоский список всех файлов на Диске в алфавитном порядке.
type FileResourceList struct {
	Items []Resource `json:"items"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

//LastUploadedResourceList Список последних добавленных на Диск файлов, отсортированных по дате загрузки (от поздних к ранним).
type LastUploadedResourceList struct {
	Items []Resource `json:"items"`
	Limit int `json:"limit"`
}

//PublicResourcesList Список опубликованных файлов на Диске.
type PublicResourcesList struct {
	Items []Resource `json:"items"`
	Type string `json:"type"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

type systemFolders struct {
	Applications string `json:"applications"`
	Downloads string `json:"downloads"`
}

//DiskДанные о свободном и занятом пространстве на Диске
type Disk struct {
	Trash_size int `json:"trash_size"`
	Total_space int `json:"total_space"`
	User_space int `json:"user_space"`
	System_folders systemFolders `json:"system_folders"`
}

//Operation Статус операции.
//Операции запускаются, когда вы копируете, перемещаете или удаляете непустые папки.
//URL для запроса статуса возвращается в ответ на такие запросы.
type Operation struct {
	Status string `json:"status"`
}

//RequestError Объект для описания ошибки
type RequestError struct {
	Description string `json:"description"`
	Error string `json:"error"`
}