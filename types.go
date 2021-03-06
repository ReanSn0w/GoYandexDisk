package goyandexdisk

import (
	"time"
)

//Link Объект содержит URL для запроса метаданных ресурса.
type Link struct {
	Href      string `json:"href"`
	Method    string `json:"method"`
	Templated bool   `json:"templated"`
}

//GetOperationToken Вернет токен операции для проверки статуса
func (l Link) GetOperationToken() string {
	if len(l.Href) < 49 {
		return ""
	}
	return l.Href[48:]
}

//Resource Описание ресурса, мета-информация о файле или папке. Включается в ответ на запрос метаинформации.
type Resource struct {
	PublicKey        string            `json:"publik_key"`
	Embeded          ResourceList      `json:"_embedded"`
	Name             string            `json:"name"`
	Created          time.Time         `json:"created"`
	CustomProperties map[string]string `json:"custom_properties"`
	PublicURL        string            `json:"public_url"`
	OriginPATH       string            `json:"origin_path"`
	Modified         string            `json:"modified"`
	Path             string            `json:"path"`
	Md5              string            `json:"md_5"`
	Type             string            `json:"type"`
	MimeType         string            `json:"mime_type"`
	Size             int               `json:"size"`
}

// PublicResource Описание публичного ресурса
type PublicResource Resource

//ResourceList Список ресурсов, содержащихся в папке. Содержит объекты Resource и свойства списка.
type ResourceList struct {
	Sort      string     `json:"sort"`
	PublicKey string     `json:"public_key"`
	Items     []Resource `json:"items"`
	Path      string     `json:"path"`
	Limit     int        `json:"limit"`
	Offset    int        `json:"offset"`
	Total     int        `json:"total"`
}

//FileResourceList Плоский список всех файлов на Диске в алфавитном порядке.
type FileResourceList struct {
	Items  []Resource `json:"items"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
}

//LastUploadedResourceList Список последних добавленных на Диск файлов, отсортированных по дате загрузки (от поздних к ранним).
type LastUploadedResourceList struct {
	Items []Resource `json:"items"`
	Limit int        `json:"limit"`
}

//PublicResourcesList Список опубликованных файлов на Диске.
type PublicResourcesList struct {
	Items  []Resource `json:"items"`
	Type   string     `json:"type"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
}

type systemFolders struct {
	Applications string `json:"applications"`
	Downloads    string `json:"downloads"`
}

//Disk Данные о свободном и занятом пространстве на Диске
type Disk struct {
	TrashSize     int           `json:"trash_size"`
	TotalSpace    int           `json:"total_space"`
	UserSpace     int           `json:"user_space"`
	SystemFolders systemFolders `json:"system_folders"`
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
	Error       string `json:"error"`
}

//RequestStatus описивает состояние запроса
type RequestStatus struct {
	Status string `json:"status"`
}
