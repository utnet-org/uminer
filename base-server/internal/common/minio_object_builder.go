package common

import (
	"fmt"
)

const (
	DATASET_FOLDER   string = "dataSets"  // 数据集存储文件夹
	CODE_FOLDER      string = "codes"     // 代码存储文件夹
	MODEL_FOLDER     string = "models"    // 模型存储文件夹
	TRAIN_JOB_FOLDER string = "trainJobs" // 训练任务存储文件夹

	UPLOAD_IMAGE_FOLDER   string = "images/swap/upload"   // 镜像上传临时文件夹
	UPLOAD_DATASET_FOLDER string = "dataSets/swap/upload" // 数据集上传临时文件夹
	UPLOAD_CODE_FOLDER    string = "codes/swap/upload"    // 代码上传临时文件夹
	UPLOAD_MODEL_FOLDER   string = "models/swap/upload"   // 模型上传临时文件夹

	DOWNLOAD_DATASET_FOLDER string = "dataSets/swap/download" // 数据集下载临时文件夹
	DOWNLOAD_CODE_FOLDER    string = "codes/swap/download"    // 代码下载临时文件夹
	DOWNLOAD_MODEL_FOLDER   string = "models/swap/download"   // 模型下载临时文件夹

	PREAB_FOLDER string = "global" // 预置目录

	BUCKET string = "spider" // 桶目录

	USERHOME string = "userhome"
)

// 桶
func GetMinioBucket() string {
	return BUCKET
}

/***** 上传 *****/
// 上传模型对象:models/swap/upload/modelId/fileName
func GetMinioUploadModelObject(modelId string, version string, fileName string) string {
	return fmt.Sprintf("%s/%s/%s/%s", UPLOAD_MODEL_FOLDER, modelId, version, fileName)
}

// 上传算法对象名：codes/swap/upload/algorithmId/version/fileName
func GetMinioUploadCodeObject(algorithmId string, version string, fileName string) string {
	return fmt.Sprintf("%s/%s/%s/%s", UPLOAD_CODE_FOLDER, algorithmId, version, fileName)
}

// 上传数据集对象名：dataSets/swap/upload/dataSetId/version/fileName
func GetMinioUploadDataSetObject(dataSetId string, version string, fileName string) string {
	return fmt.Sprintf("%s/%s/%s/%s", UPLOAD_DATASET_FOLDER, dataSetId, version, fileName)
}

// 上传镜像对象名：images/swap/images/imageId/version/fileName
func GetMinioUploadImageObject(imageId string, version string, fileName string) string {
	return fmt.Sprintf("%s/%s/%s/%s", UPLOAD_IMAGE_FOLDER, imageId, version, fileName)
}

/***** 下载 *****/
// 下载模型对象:models/swap/download/modelId/version/fileName
func GetMinioDownloadModelObject(modelId string, version string, fileName string) string {
	return fmt.Sprintf("%s/%s/%s/%s", DOWNLOAD_MODEL_FOLDER, modelId, version, fileName)
}

// 下载模型对象:models/swap/download/modelId/version
func GetMinioDownloadModelVersionObject(modelId string, version string) string {
	return fmt.Sprintf("%s/%s/%s", DOWNLOAD_MODEL_FOLDER, modelId, version)
}

// 下载模型对象:models/swap/download/modelId
func GetMinioDownloadModelPathObject(modelId string) string {
	return fmt.Sprintf("%s/%s", DOWNLOAD_MODEL_FOLDER, modelId)
}

// 下载算法对象名：codes/swap/download/algorithmId/version/fileName
func GetMinioDownloadCodeObject(algorithmId string, version string, fileName string) string {
	return fmt.Sprintf("%s/%s/%s/%s", DOWNLOAD_CODE_FOLDER, algorithmId, version, fileName)
}

// 下载算法对象名：codes/swap/download/algorithmId/version
func GetMinioDownloadCodeVersionObject(algorithmId string, version string) string {
	return fmt.Sprintf("%s/%s/%s", DOWNLOAD_CODE_FOLDER, algorithmId, version)
}

// 下载算法对象名：codes/swap/download/algorithmId
func GetMinioDownloadCodePathObject(algorithmId string) string {
	return fmt.Sprintf("%s/%s", DOWNLOAD_CODE_FOLDER, algorithmId)
}

// 下载数据集对象名：dataSets/swap/download/dataSetId/version/fileName
func GetMinioDownloadDataSetObject(dataSetId string, version string, fileName string) string {
	return fmt.Sprintf("%s/%s/%s/%s", DOWNLOAD_DATASET_FOLDER, dataSetId, version, fileName)
}

/**************************************
		管理端(预置)相关目录获取
**************************************/

// 模型对象:models/global/modelId/version
func GetMinioPreModelObject(modelId string, version string) string {
	return fmt.Sprintf("%s/%s/%s/%s", MODEL_FOLDER, PREAB_FOLDER, modelId, version)
}

// 模型对象:models/global/modelId
func GetMinioPreModelPathObject(modelId string) string {
	return fmt.Sprintf("%s/%s/%s", MODEL_FOLDER, PREAB_FOLDER, modelId)
}

// 算法对象:codes/global/algorithmId/version
func GetMinioPreCodeObject(algorithmId string, version string) string {
	return fmt.Sprintf("%s/%s/%s/%s", CODE_FOLDER, PREAB_FOLDER, algorithmId, version)
}

// 算法对象:codes/global/algorithmId
func GetMinioPreCodePathObject(algorithmId string) string {
	return fmt.Sprintf("%s/%s/%s", CODE_FOLDER, PREAB_FOLDER, algorithmId)
}

// 数据集对象:datasets/global/dataSetId/version
func GetMinioPreDataSetObject(dataSetId string, version string) string {
	return fmt.Sprintf("%s/%s/%s/%s", DATASET_FOLDER, PREAB_FOLDER, dataSetId, version)
}

// 数据集对象:datasets/global/dataSetId
func GetMinioPreDataSetPathObject(dataSetId string) string {
	return fmt.Sprintf("%s/%s/%s", DATASET_FOLDER, PREAB_FOLDER, dataSetId)
}

/**************************************
		用户端相关目录获取
**************************************/

// 模型对象:models/spaceId/userId/modelId/version
func GetMinioModelObject(spaceId string, userId string, modelId string, version string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", MODEL_FOLDER, spaceId, userId, modelId, version)
}

// 模型对象:models/spaceId/userId/modelId/version
func GetMinioModelPathObject(spaceId string, userId string, modelId string) string {
	return fmt.Sprintf("%s/%s/%s/%s", MODEL_FOLDER, spaceId, userId, modelId)
}

// 算法对象:codes/spaceId/userId/algorithmId/version
func GetMinioCodeObject(spaceId string, userId string, algorithmId string, version string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", CODE_FOLDER, spaceId, userId, algorithmId, version)
}

// 算法对象:codes/spaceId/userId/algorithmId
func GetMinioCodePathObject(spaceId string, userId string, algorithmId string) string {
	return fmt.Sprintf("%s/%s/%s/%s", CODE_FOLDER, spaceId, userId, algorithmId)
}

// 数据集对象:dataSets/spaceId/userId/dataSetId/version
func GetMinioDataSetObject(spaceId string, userId string, dataSetId string, version string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", DATASET_FOLDER, spaceId, userId, dataSetId, version)
}

// 数据集对象:dataSets/spaceId/userId/dataSetId
func GetMinioDataSetPathObject(spaceId string, userId string, dataSetId string) string {
	return fmt.Sprintf("%s/%s/%s/%s", DATASET_FOLDER, spaceId, userId, dataSetId)
}

// 训练任务对象:trainJobs/spaceId/userId/trainJobId
func GetMinioTrainJobObject(spaceId string, userId string, trainJobId string) string {
	return fmt.Sprintf("%s/%s/%s/%s", TRAIN_JOB_FOLDER, spaceId, userId, trainJobId)
}

func GetUserBucket(userId string) string {
	return fmt.Sprintf("%s", userId)
}

func GetUserHomeObject() string {
	return USERHOME
}

func GetUserHomePath(userId string) string {
	return fmt.Sprintf("%s/%s", userId, USERHOME)
}
