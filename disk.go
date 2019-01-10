package GoYandexDisk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const baseURL = "https://cloud-api.yandex.net/v1/disk/"

//YandexDisk Структура для работы с диском
type YandexDisk struct {
	accessToken string
}

func ConnectDisk(token string) YandexDisk {
	return YandexDisk{token}
}

func (yd *YandexDisk) DiskStatus() (Disk, error) {
	body, err := yd.createRequest(
		"",
		"GET",
		"",
		nil,
	)
	if err != nil {
		return Disk{}, err
	}

	var result Disk
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

//Get Получение ресурса списка ресурсов
func (yd YandexDisk) Get(path string, limit int, offset int) (Resource, error) {
	body, err := yd.createRequest(
		"resources",
		"GET",
		createParamsMap(map[string]string{
			"path": path,
		}),
		nil,
	)
	if err != nil {
		return Resource{}, err
	}

	var result Resource
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

////FlatList Получение плоского списка файлов
//func (yd YandexDisk)FlatList(limit int, mediatype []string, offset int) (FileResourceList, error) {
//	//TODO добавить в запрос mediatype
//	link := "resources/files?limit=" + string(limit) + "&offset=" + string(offset)
//	body, err := yd.createRequest(link, "GET")
//	if err != nil {
//		return FileResourceList{}, err
//	}
//
//	var result FileResourceList
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}
//
////LastUploaded получение списка последних загруженных файлов
//func (yd YandexDisk)LastUploaded(limit int, mediatype []string) (LastUploadedResourceList, error) {
//	//TODO добавить в запрос mediatype
//	link := "resources/files?limit=" + string(limit)
//	body, err := yd.createRequest(link, "GET")
//	if err != nil {
//		return LastUploadedResourceList{}, err
//	}
//
//	var result LastUploadedResourceList
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}
//
////UploadFile
////func (yd YandexDisk)UploadFile() {}

//UploadFileFromNet Загрузка файла на диск по ссылке из интернета
func (yd YandexDisk) UploadFileFromNet(url string, path string, fields []string, disableRedirects bool) (Link, error) {
	body, err := yd.createRequest(
		"resources/upload",
		"POST",
		createParamsMap(map[string]string{
			"url":  url,
			"path": path,
		}),
		nil,
	)
	if err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

////DownloadFile вернет объект Link со ссылкой для загрузки
//func (yd YandexDisk)DownloadFile(path string) (Link, error){
//	link := "resources/download?path=" + path
//	body, err := yd.createRequest(link, "GET")
//	if err != nil {
//		return Link{}, err
//	}
//
//	var result Link
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}
//
////CopyResource копирование ресурса
//func (yd YandexDisk)CopyResource(from string, path string, owerwrite bool) (Link, error){
//	//TODO дописать
//	link := "resources/copy?from=" + from + "&path=" + path
//	body, err := yd.createRequest(link, "GET")
//	if err != nil {
//		return Link{}, err
//	}
//
//	var result Link
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}
//
////ReplaceResource перемещение ресурса
//func (yd YandexDisk)ReplaceResource(from string, path string, owerwrite bool) (Link, error){
//	link := "resources/move?from=" + from + "&path=" + path
//	body, err := yd.createRequest(link, "GET")
//	if err != nil {
//		return Link{}, err
//	}
//
//	var result Link
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}
//
////DeleteResource удаление ресурса
//func (yd YandexDisk)DeleteResource(path string, permanently bool) (Link, error){
//	//TODO дописать
//	link := "resources?path=" + path
//	body, err := yd.createRequest(link, "DELETE")
//	if err != nil {
//		return Link{}, err
//	}
//
//	var result Link
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}
//
////CreateFolder создает папке на диске по указанному пути
//func (yd YandexDisk)CreateFolder(path string) (Link, error){
//	link := "resources/move?path=" + path
//	body, err := yd.createRequest(link, "PUT")
//	if err != nil {
//		return Link{}, err
//	}
//
//	var result Link
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}

//PublishResource публикация ресурса
func (yd YandexDisk) PublishResource(path string) (Link, error) {
	link := "resources/publish" + "?path=" + path
	body, err := yd.createRequest(
		link,
		"PUT",
		"",
		nil,
	)
	if err != nil {
		return Link{}, err
	}

	var result Link
	err = json.NewDecoder(body).Decode(&result)
	fmt.Println(result)
	return result, err
}

////UnpublishResource отказаться от публикации ресурса
//func (yd YandexDisk)UnpublishResource(publickey string) (Resource, error){
//	//TODO доделать функцию
//	link := "public/resources" + publickey
//	body, err := yd.createRequest(link, "GET")
//	if err != nil {
//		return Resource{}, err
//	}
//
//	var result Resource
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}
//
////PublicResourceMeta получение мета информации о публичном ресурсе
//func (yd YandexDisk)PublicResourceMeta(publikKey string) (Resource, error){
//	link := "public/resources?public_key=" + publikKey
//	body, err := yd.createRequest(link, "GET")
//	if err != nil {
//		return Resource{}, err
//	}
//
//	var result Resource
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}
//
////DownloadPublicResource получение ссылки на зугрузку публичного ресурса
//func (yd YandexDisk)DownloadPublicResource(publikKey string) (Link, error){
//	link := "public/resources/download?public_key=" + publikKey
//	body, err := yd.createRequest(link, "GET")
//	if err != nil {
//		return Link{}, err
//	}
//
//	var result Link
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}
//
////SavePublicResource сохранение ресурса в папку загрузки
//func (yd YandexDisk)SavePublicResource(publikKey string) (Link, error){
//	link := "resources/download?public_key=" + publikKey
//	body, err := yd.createRequest(link, "POST")
//	if err != nil {
//		return Link{}, err
//	}
//
//	var result Link
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}
//
////PublicResourcesMeta получение списка публичных файлов
//func (yd YandexDisk)PublishResourcesList(limit int, offset int) (PublicResourcesList, error){
//	link := "resources/public?limit=" + string(limit) + "&offset=" + string(offset)
//	body, err := yd.createRequest(link, "GET")
//	if err != nil {
//		return PublicResourcesList{}, err
//	}
//
//	var result PublicResourcesList
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}
//
////CleanTrash Очистка файла из корзины, если path = "" происходит полная очистка
//func (yd YandexDisk)CleanTrash(path string) (Link, error){
//	link := "trash/resources?path=" + path
//	body, err := yd.createRequest(
//		link,
//		"GET",
//		createParamsMap(map[string]string{
//			"path": path,
//		}),
//		nil,
//	)
//	if err != nil {
//		return Link{}, err
//	}
//
//	var result Link
//	err = json.NewDecoder(body).Decode(&result)
//	return result, err
//}
//
////RestoreResource восстановление ранее удаленного ресурса
////func (yd YandexDisk)RestoreResource() {}

//Status возвращает текущий статус выполнения операции
func (yd YandexDisk) Status(id string) (RequestStatus, error) {
	link := "operations/" + id
	body, err := yd.createRequest(
		link,
		"GET",
		createParamsMap(map[string]string{}),
		nil,
	)
	if err != nil {
		return RequestStatus{}, err
	}

	var result RequestStatus
	err = json.NewDecoder(body).Decode(&result)
	return result, err
}

func (ya YandexDisk) createRequest(link, method string, form string, body io.Reader) (io.ReadCloser, error) {
	//Формирование запроса
	client := http.Client{}
	req, err := http.NewRequest(method, baseURL+link+"?"+form, body)
	if form == "" {
		req, err = http.NewRequest(method, baseURL+link, body)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", ya.accessToken)

	//Его выполнение
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.StatusCode)

	return resp.Body, err
}

func createParamsMap(params map[string]string) string {
	values := url.Values{}
	for key, value := range params {
		values.Add(key, value)
	}
	return values.Encode()
}
