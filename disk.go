package GoYandexDisk

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const baseURL = "https://cloud-api.yandex.net/v1/disk/"

//YandexDisk Структура для работы с диском
type YandexDisk struct {
	accessToken string
}

//ConnectDisk вернет структуру для управления диском
func ConnectDisk(token string) YandexDisk {
	return YandexDisk{token}
}

//DiskStatus Запрос статуса диска
func (yd *YandexDisk) DiskStatus() (Disk, error) {
	body, code, err := yd.createRequest(
		"",
		"GET",
		"",
		nil,
	)
	if err != nil {
		return Disk{}, err
	}

	if err = parseError(code); err != nil {
		return Disk{}, err
	}

	var result Disk
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//Get Получение ресурса списка ресурсов
func (yd YandexDisk) Get(path string, limit int, offset int) (Resource, error) {
	body, code, err := yd.createRequest(
		"resources",
		"GET",
		createQuery(map[string]string{
			"path":   path,
			"limit":  strconv.Itoa(limit),
			"offset": strconv.Itoa(offset),
		}),
		nil,
	)
	if err != nil {
		return Resource{}, err
	}

	if err = parseError(code); err != nil {
		return Resource{}, err
	}

	var result Resource
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//FlatList Получение плоского списка файлов
func (yd YandexDisk) FlatList(limit int, offset int, media_type string) (FileResourceList, error) {
	body, code, err := yd.createRequest(
		"resources/files",
		"GET",
		createQuery(map[string]string{
			"limit":      strconv.Itoa(limit),
			"offset":     strconv.Itoa(offset),
			"media_type": media_type,
		}),
		nil,
	)
	if err != nil {
		return FileResourceList{}, err
	}

	if err = parseError(code); err != nil {
		return FileResourceList{}, err
	}

	var result FileResourceList
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//LastUploaded получение списка последних загруженных файлов
func (yd YandexDisk) LastUploaded(limit int, media_type string) (LastUploadedResourceList, error) {
	body, code, err := yd.createRequest(
		"resources/files",
		"GET",
		createQuery(map[string]string{
			"limit":      strconv.Itoa(limit),
			"media_type": media_type,
		}),
		nil,
	)
	if err != nil {
		return LastUploadedResourceList{}, err
	}

	if err = parseError(code); err != nil {
		return LastUploadedResourceList{}, err
	}

	var result LastUploadedResourceList
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

////UploadFile
////func (yd YandexDisk)UploadFile() {}

//UploadFileFromNet Загрузка файла на диск по ссылке из интернета
func (yd YandexDisk) UploadFileFromNet(url string, path string, disable_redirects bool) (Link, error) {
	body, code, err := yd.createRequest(
		"resources/upload",
		"POST",
		createQuery(map[string]string{
			"url":               url,
			"path":              path,
			"disable_redirects": strconv.FormatBool(disable_redirects),
		}),
		nil,
	)
	if err != nil {
		return Link{}, err
	}

	if err = parseError(code); err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//DownloadFile вернет объект Link со ссылкой для загрузки
func (yd YandexDisk) DownloadFile(path string) (Link, error) {
	body, code, err := yd.createRequest(
		"resources/download",
		"GET",
		createQuery(map[string]string{
			"path": path,
		}),
		nil,
	)
	if err != nil {
		return Link{}, err
	}

	if err = parseError(code); err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//CopyResource копирование ресурса
//Обязательные параметры from, path
func (yd YandexDisk) CopyResource(from string, path string, owerwrite bool, force_async bool) (Link, error) {
	body, code, err := yd.createRequest(
		"resources/move",
		"PATH",
		createQuery(map[string]string{
			"from":        from,
			"path":        path,
			"owerwrite":   strconv.FormatBool(owerwrite),
			"force_async": strconv.FormatBool(force_async),
		}),
		nil,
	)
	if err != nil {
		return Link{}, err
	}

	if err = parseError(code); err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//ReplaceResource перемещение ресурса
//Обязательные параметры from, path
func (yd YandexDisk) ReplaceResource(from string, path string, owerwrite bool, force_async bool) (Link, error) {
	body, code, err := yd.createRequest(
		"resources/move",
		"PATH",
		createQuery(map[string]string{
			"from":        from,
			"path":        path,
			"owerwrite":   strconv.FormatBool(owerwrite),
			"force_async": strconv.FormatBool(force_async),
		}),
		nil,
	)
	if err != nil {
		return Link{}, err
	}

	if err = parseError(code); err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//DeleteResource удаление ресурса
//Обязательные параметры path
func (yd YandexDisk) DeleteResource(path string, permanently bool, force_async bool) (Link, error) {
	body, code, err := yd.createRequest(
		"resource",
		"DELETE",
		createQuery(map[string]string{
			"path":        path,
			"permanently": strconv.FormatBool(permanently),
			"force_async": strconv.FormatBool(force_async),
		}),
		nil,
	)
	if err != nil {
		return Link{}, err
	}

	if err = parseError(code); err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//CreateFolder создает папке на диске по указанному пути
func (yd YandexDisk) CreateFolder(path string) (Link, error) {
	body, code, err := yd.createRequest(
		"resources",
		"PUT",
		createQuery(map[string]string{
			"path": path,
		}),
		nil,
	)
	if err != nil {
		return Link{}, err
	}

	if err = parseError(code); err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//PublishResource публикация ресурса
func (yd YandexDisk) PublishResource(path string) (Link, error) {
	body, code, err := yd.createRequest(
		"resources/publish",
		"PUT",
		createQuery(map[string]string{
			"path": path,
		}),
		nil,
	)
	if err != nil {
		return Link{}, err
	}

	if err = parseError(code); err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	fmt.Println(result)
	return result, err
}

////UnpublishResource отказаться от публикации ресурса
func (yd YandexDisk) UnpublishResource(public_key string) (Link, error) {
	body, code, err := yd.createRequest(
		"resources/unpublish",
		"PUT",
		createQuery(map[string]string{
			"public_key": public_key,
		}),
		nil)
	if err != nil {
		return Link{}, err
	}

	//Обрабатываем статус коды запроса
	if err := parseError(code); err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//PublicResourceMeta получение мета информации о публичном ресурсе
//Обязательные параметры publik_key
func (yd YandexDisk) PublicResourceMeta(publik_key string, limit int, offset int, path string, sort string) (PublicResource, error) {
	body, code, err := yd.createRequest(
		"public/resources",
		"GET",
		createQuery(map[string]string{
			"public_key": publik_key,
			"limit":      strconv.Itoa(limit),
			"offset":     strconv.Itoa(offset),
			"path":       path,
			"sort":       sort,
		}),
		nil,
	)
	if err != nil {
		return PublicResource{}, err
	}

	//Обработка документированных ошибок
	if err := parseError(code); err != nil {
		return PublicResource{}, err
	}

	var result PublicResource
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//DownloadPublicResource получение ссылки на зугрузку публичного ресурса
//Обязательные аргументы public_key
func (yd YandexDisk) DownloadPublicResource(publik_key string, path string) (Link, error) {
	body, code, err := yd.createRequest(
		"public/resources/download",
		"GET",
		createQuery(map[string]string{
			"public_key": publik_key,
			"path":       path,
		}),
		nil)
	if err != nil {
		return Link{}, err
	}

	//Обработка кода ответа от сервера
	if err := parseError(code); err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//SavePublicResource сохранение ресурса в папку загрузки
//Обязательные параметры public_key
func (yd YandexDisk) SavePublicResource(publik_key string, path string, name string, save_path string) (Link, error) {
	//Формирование запроса
	body, code, err := yd.createRequest(
		"public/resources/save-to-disk",
		"POST",
		createQuery(map[string]string{
			"public_key": publik_key,
			"path":       path,
			"name":       name,
			"save_path":  save_path,
		}),
		nil,
	)
	if err != nil {
		return Link{}, err
	}

	//Обработка документированных ошибок
	if err := parseError(code); err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//PublicResourcesMeta получение списка публичных файлов
//Обязательные аргументы public_key
func (yd YandexDisk) PublishResourcesList(public_key string, limit int, offset int, path string, sort string) (PublicResource, error) {
	//Формирование запроса
	body, code, err := yd.createRequest(
		"resources/public",
		"GET",
		createQuery(map[string]string{
			"public_key": public_key,
			"limit":      strconv.Itoa(limit),
			"offset":     strconv.Itoa(offset),
			"path":       path,
			"sort":       sort,
		}),
		nil,
	)
	if err != nil {
		return PublicResource{}, err
	}

	//Обработка документированных ошибок
	if err := parseError(code); err != nil {
		return PublicResource{}, err
	}

	var result PublicResource
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//CleanTrash Очистка файла из корзины, если path = "" происходит полная очистка
func (yd YandexDisk) CleanTrash(path string) (Link, error) {
	//Формирование запроса
	body, code, err := yd.createRequest(
		"trash/resources",
		"DELETE",
		createQuery(map[string]string{
			"path": path,
		}),
		nil,
	)
	if err != nil {
		return Link{}, err
	}

	//Обработка документированных ошибок
	if err := parseError(code); err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//RestoreResource восстановление ранее удаленного ресурса
//Обязательные параметры path
func (yd YandexDisk) RestoreResource(path string, name string, owerwrite bool, force_asunc bool) (Link, error) {
	//Формирование запроса
	body, code, err := yd.createRequest(
		"trash/resources",
		"DELETE",
		createQuery(map[string]string{
			"path":        path,
			"name":        name,
			"owerwrite":   strconv.FormatBool(owerwrite),
			"force_asunc": strconv.FormatBool(force_asunc),
		}),
		nil,
	)
	if err != nil {
		return Link{}, err
	}

	//Обработка документированных ошибок
	if err := parseError(code); err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//Status возвращает текущий статус выполнения операции
func (yd YandexDisk) Status(id string) (RequestStatus, error) {
	//Отправка запроса
	body, num, err := yd.createRequest(
		"operations/"+id,
		"GET",
		createQuery(map[string]string{}),
		nil,
	)
	if err != nil {
		return RequestStatus{}, err
	}

	//Обработка документированных ошибок
	switch num {
	case 400:
		return RequestStatus{}, errors.New("Некорректные данные")
	case 401:
		return RequestStatus{}, errors.New("Не авторизован")
	case 403:
		return RequestStatus{}, errors.New("Не достаточно прав для изменения данных в общей папке.")
	case 404:
		return RequestStatus{}, errors.New("Не удалось найти запрошенный ресурс")
	case 406:
		return RequestStatus{}, errors.New("Ресурс не может быть представлен в запрошенном формате")
	case 429:
		return RequestStatus{}, errors.New("Слишком много запросов")
	case 503:
		return RequestStatus{}, errors.New("Сервис временно не доступен")
	}

	var result RequestStatus
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

func (ya YandexDisk) createRequest(link, method string, form string, body io.Reader) (io.ReadCloser, int, error) {
	//Формирование запроса
	client := http.Client{}
	req, err := http.NewRequest(method, baseURL+link+"?"+form, body)
	if form == "" {
		req, err = http.NewRequest(method, baseURL+link, body)
	}
	if err != nil {
		return nil, 404, err
	}
	req.Header.Add("Authorization", ya.accessToken)

	//Его выполнение
	resp, err := client.Do(req)
	if err != nil {
		return nil, 404, err
	}

	return resp.Body, resp.StatusCode, err
}

func createQuery(params map[string]string) string {
	values := url.Values{}
	for key, value := range params {
		if value == "" || value == "false" || value == "0" {
			continue
		}
		values.Add(key, value)
	}
	return values.Encode()
}

func parseError(code int) error {
	switch code {
	case 400:
		return errors.New("Не корректные данные")
	case 401:
		return errors.New("Не авторизован")
	case 403:
		return errors.New("Доступ запрещен. Возможно, у приложения не достаточно прав")
	case 404:
		return errors.New("Не удалось найти запрошенный ресурс")
	case 406:
		return errors.New("Ресурс не может быть представлен в запрошенном виде")
	case 409:
		return errors.New("Указанного пути не существует")
	case 423:
		return errors.New("Ресурс заблокирован. Возможно над ним выполняется другая операция")
	case 429:
		return errors.New("Слишком много запросов")
	case 503:
		return errors.New("Ресурс временно не доступен")
	case 507:
		return errors.New("Недостаточно свободного места")
	}
	return nil
}
