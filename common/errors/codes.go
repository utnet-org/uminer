package errors

const (
	/* 10000~11000 基础错误*/
	ErrorUnknown                   = 10001 // 未知错误
	ErrorStructCopy                = 10002 // 结构体拷贝失败
	ErrorInvalidRequestParameter   = 10003 // 参数错误
	ErrorVersionInvalid            = 10004 // 版本字符串非法
	ErrorZipFailed                 = 10005 // 压缩失败
	ErrorUnzipFailed               = 10006 // 解压失败
	ErrorUrlParseFailed            = 10007 // url解析失败
	ErrorParseDurationFailed       = 10008 // 时间解析失败
	ErrorEncryptPasswordFailed     = 10009 // 密码加密失败
	ErrorDirNotExisted             = 10010 // 文件夹不存在
	ErrorOperateDirFailed          = 10011 // 操作文件夹失败
	ErrorOperateFileFailed         = 10012 // 操作文件失败
	ErrorConfigValueValidateFailed = 10013 // 校验配置值失败
	ErrorFormatParseFailed         = 10014 // 格式解析失败
	ErrorInternal                  = 10015 // 内部错误
	// http请求相关错误
	ErrorHttpNewRequest     = 10020 // 获取http request实体失败
	ErrorHttpDoRequest      = 10021 // http请求失败
	ErrorHttpReadBody       = 10022 // http读取回包失败
	ErrorJsonMarshal        = 10023 // json序列化失败
	ErrorJsonUnmarshal      = 10024 // json反序列化失败
	ErrorHttpWriteFailed    = 10025 // http写入失败
	ErrorHttpBindFormFailed = 10026 // http绑定form失败
	// minio操作相关错误
	ErrorMinioBucketInitFailed         = 10030 // minio初始化失败
	ErrorMinioBucketExisted            = 10031 // 桶已存在
	ErrorMinioBucketNotExist           = 10032 // 桶不存在
	ErrorMinioCheckBucketExistFailed   = 10033 // 检查桶存在失败
	ErrorMinioMakeBucketFailed         = 10034 // 创建桶失败
	ErrorMinioDeleteBucketFailed       = 10035 // 删除桶失败
	ErrorMinioPresignedPutObjectFailed = 10036 // 生成上传临时url失败
	ErrorMinioPresignedGetObjectFailed = 10037 // 生成下载临时url失败
	ErrorMinioListObjectFailed         = 10038 // 查看对象列表失败
	ErrorMinioCheckObjectExistFailed   = 10039 // 查看对象失败
	ErrorMinioCheckObjectNotExisted    = 10040 // 对象不存在
	ErrorMinioCreateAccountFailed      = 10041 // 创建用户失败
	ErrorMinioOperationFailed          = 10042 // 操作失败
	ErrorMinioRemoveObjectFailed       = 10043 // 删除对象失败

	// db操作相关错误
	ErrorDBInitFailed        = 10050 // db初始化失败
	ErrorDBFindFailed        = 10051 // db列表查询失败
	ErrorDBFirstFailed       = 10052 // db单条查询失败
	ErrorDBCountFailed       = 10053 // db查询列表条数失败
	ErrorDBFindEmpty         = 10054 // db查询为空
	ErrorDBSelectParamsEmpty = 10055 // db查询条件为空
	ErrorDBUpdateFailed      = 10056 // db更新失败
	ErrorDBCreateFailed      = 10057 // db插入失败
	ErrorDBDeleteFailed      = 10058 // db删除失败
	ErrorDBPrimaryKeyEmpty   = 10059 // db操作缺少主键
	// k8s操作相关错误
	ErrorK8sCreateServiceFailed = 10070 // k8s创建service失败
	ErrorK8sDeleteServiceFailed = 10071 // k8s删除service失败
	ErrorK8sCreateIngressFailed = 10072 // k8s创建ingress失败
	ErrorK8sDeleteIngressFailed = 10073 // k8s删除ingress失败
	ErrorK8sDeletePVFailed      = 10074 // k8s删除PV失败
	ErrorK8sDeletePVCFailed     = 10075 // k8s删除PVC失败
	ErrorK8sDeleteSecretFailed  = 10076 // k8s删除Secret失败
	ErrorFluidInitFailed        = 10077 // Fluid初始化失败
	// Harbor操作相关错误
	ErrorHarborProjectExists       = 10080 // harbor项目已存在
	ErrorHarborCreateProjectFailed = 10081 // harbor创建项目失败
	ErrorHarborCheckProjectFailed  = 10082 // harbor查询项目失败
	// redis操作相关错误
	ErroRedisParseUrlFailed    = 10100 // redis解析url失败
	ErroRedisHSetFailed        = 10101 // redisHSet失败
	ErroRedisHGetFailed        = 10102 // redisHGet失败
	ErroRedisHDelFailed        = 10103 // redisHDel失败
	ErrorRedisLockObtainFailed = 10104 // redis锁获取失败
	// influxdb操作相关错误
	ErroInfluxdbInitFailed  = 10200 // influxdb初始化失败
	ErroInfluxdbFindFailed  = 10201 // influxdb列表查询失败
	ErroInfluxdbWriteFailed = 10202 // influxdb插入失败

	/* 11001~12000 资源管理错误*/
	ErrorDeleteResourcePool   = 11001 // 删除资源池失败
	ErrorCreateResourcePool   = 11002 // 创建资源池失败
	ErrorListResourcePool     = 11003 // 查询资源池失败
	ErrorGetResourcePool      = 11004 // 获取资源池信息失败
	ErrorUpdateResourcePool   = 11005 // 更新资源池信息失败
	ErrorListResource         = 11006 // 获取资源列表失败
	ErrorCreateResource       = 11007 // 创建资源失败
	ErrorDeleteResource       = 11008 // 删除资源失败
	ErrorUpdateResource       = 11009 // 更新资源失败
	ErrorBindingNode          = 11010 // 绑定节点失败
	ErrorCreateResourceSpec   = 11011 // 创建资源规格失败
	ErrorListResourceSpec     = 11012 // 获取资源规格列表失败
	ErrorDeleteResourceSpec   = 11013 // 删除资源规格失败
	ErrorGetResourceSpec      = 11014 // 获取资源规格信息失败
	ErrorResourceDiscovery    = 11015 // 资源发现失败
	ErrorResourceNotExist     = 11016 // 资源名字不存在(创建资源规格)
	ErrorResourceSpecNotExist = 11017 // 资源规格不存在
	ErrorResourceExist        = 11018 // 资源名已存在(创建自定义资源)
	ErrorListNode             = 11019 // 获取节点列表失败

	/* 12001~13000 算法管理错误*/
	ErrorFindAlgorithmVersionAccessMaxIdFailed = 12001 // 查找最新公共算法版本失败
	ErrorFindAlgorithmVersionMaxIdFailed       = 12002 // 查找最新算法版本失败
	ErrorAlgorithmNotMy                        = 12003 // 删除算法不是我的
	ErrorAlgorithmRepeat                       = 12004 // 新增算法重复
	ErrorAlgorithmVersionRepeat                = 12005 // 新增算法版本重复
	ErrorFindAlgorithmAuthWrong                = 12006 // 查找的算法不在权限范围内
	ErrorAlgorithmVersionFileNotFound          = 12007 // 算法版本文件不存在
	ErrorAlgorithmVersionFileExisted           = 12008 // 算法版本文件已存在
	ErrorAlgorithmAccessVersionExisted         = 12009 // 公共算法版本已存在
	ErrorAlgorithmVersionFileNotReady          = 12010 // 旧版本算法文件状态未就绪
	ErrorAlgorithmVersionUploadAuth            = 12011 // 无权提交此算法版本文件

	/* 13001~14000 镜像管理错误*/
	ErrorImageStatusMakeError      = 13001 // 制作镜像状态异常
	ErrorImageSourceTypeToUpload   = 13002 // 上传错误镜像来源类型
	ErrorImageExisted              = 13003 // 镜像已存在
	ErrorImageNotExist             = 13004 // 镜像不存在
	ErrorImageOpForbidden          = 13005 // 无权限镜像操作
	ErrorImageSourceType           = 13006 // 不支持的镜像来源
	ErrorImageContainerCommitError = 13007 // 容器提交操作失败
	ErrorImagePushFailed           = 13008 // 推送镜像操作失败

	/* 14001~15000 开发管理错误*/
	ErrorNotebookStatusForbidden          = 14001 // notebook状态不允许操作
	ErrorNotebookImageNoPermission        = 14002 // notebook使用镜像不在权限范围内
	ErrorNotebookAlgorithmNoPermission    = 14003 // notebook使用算法不在权限范围内
	ErrorNotebookBalanceNotEnough         = 14004 // 余额不足
	ErrorNotebookParseResourceSpecFailed  = 14005 // 解析资源规格失败
	ErrorNotebookImageStatusForbidden     = 14006 // notebook使用镜像状态不允许操作
	ErrorNotebookAlgorithmStatusForbidden = 14007 // notebook使用算法状态不允许操作
	ErrorNotebookRepeat                   = 14008 // notebook重复
	ErrorNotebookDatasetNoPermission      = 14009 // notebook使用数据集不在权限范围内
	ErrorNotebookDatasetStatusForbidden   = 14010 // notebook使用数据集状态不允许操作
	ErrorNotebookRepeatedToSave           = 14011 // notebook保存中，禁止重复保存
	ErrorNotebookNoFoundRuntimeContainer  = 14012 // notebook未找到运行中的容器
	ErrorNotebookResourcePoolForbidden    = 14013 // notebook使用资源池不允许操作
	ErrorNotebookMountExternalForbidden   = 14014 // 用户不允许外部挂载

	/* 15001~16000 训练管理错误*/
	ErrorTrainImageForbidden         = 15001 // 训练使用镜像不在权限范围内
	ErrorTrainDataSetForbidden       = 15002 // 训练使用数据集不在权限范围内
	ErrorTrainAlgorithmForbidden     = 15003 // 训练使用算法不在权限范围内
	ErrorPipelineDoRequest           = 15004 // 训练pipeline请求失败
	ErrorDeleteJobRequest            = 15005 // 删除任务失败, 只能删除终结状态下的任务
	ErrorTrainBalanceNotEnough       = 15006 // 余额不足
	ErrorStopTerminatedJob           = 15007 // 任务已终止，无需停止
	ErrorJobImageStatusForbidden     = 15008 // job使用镜像,状态不允许操作
	ErrorJobAlgorithmStatusForbidden = 15009 // job使用算法,状态不允许操作
	ErrorJobDatasetStatusForbidden   = 15010 // job使用数据集,状态不允许操作
	ErrorJobUniqueIndexConflict      = 15011 // job任务重名
	ErrorRepeatJobConfigName         = 15022 // 分布式子任务重名
	ErrorTrainResourcePoolForbidden  = 15023 // 训练使用资源池不允许操作
	ErrorTrainMountExternalForbidden = 15024 // 用户不允许外部挂载

	/* 16001~17000 用户管理错误*/
	// 用户登录
	ErrorNotAuthorized            = 16001 // 用户未登录
	ErrorAuthenticationForbidden  = 16002 // 用户状态不在权限范围内
	ErrorCreateTokenFailed        = 16003 // 创建token失败
	ErrorParseTokenFailed         = 16004 // 解析token失败
	ErrorTokenInvalid             = 16005 // token无效
	ErrorSessionIdNotFound        = 16006 // sessionId为空
	ErrorUserNoAuthSession        = 16007 // session不存在
	ErrorUserGetAuthSessionFailed = 16008 // session获取失败
	ErrorAuthenticationFailed     = 16009 // 认证失败
	ErrorTokenRenew               = 16010 // 认证失败,token被重置

	// 用户管理
	ErrorUserAccountNotExisted            = 16020 // 用户不存在
	ErrorUserAccountExisted               = 16021 // 用户已存在
	ErrorUserIdNotRight                   = 16022 // 用户id错误
	ErrorWorkSpaceExisted                 = 16023 // 空间已存在
	ErrorWorkSpaceNotExist                = 16024 // 空间不存在
	ErrorUserWorkSpaceNoPermission        = 16025 // 用户无空间权限
	ErrorWorkSpaceResourcePoolBound       = 16026 // 空间与资源池已绑定
	ErrorUserConfigKeyNotExist            = 16027 // 配置项不存在
	ErrorUserAccountBinded                = 16028 // 账号已绑定
	ErrorUserChangeMinioUsernameForbidden = 16029 // 不能修改minio用户名
	ErrorUserMinioUsernameNotExist        = 16030 // minio用户不存在

	/* 17001~18000 计费管理错误*/
	ErrorBillingObtainLockFailed = 17001 // 获取锁失败
	ErrorBillingStatusForbidden  = 17002 // 状态不允许操作

	/* 18001~19000 模型管理错误*/
	ErrorModelNoPermission              = 18001 // 没权限操作模型
	ErrorModelRepeat                    = 18002 // 模型重复
	ErrorModelVersionRepeat             = 18003 // 模型版本重复
	ErrorModelVersionFileNotFound       = 18004 // 模型版本文件不存在
	ErrorModelVersionFileExisted        = 18005 // 模型版本文件已存在
	ErrorModelCreateAlgorithmNotExisted = 18006 // 创建模型时归属算法不存在

	/* 19001~19500 数据集管理错误*/
	ErrorDatasetFileNotFound    = 19001 // 数据集文件不存在
	ErrorDatasetAlreadyShared   = 19002 // 数据集已分享
	ErrorDatasetNoPermission    = 19003 // 没有权限操作
	ErrorDatasetRepeat          = 19004 // 数据集重复
	ErrorDatasetStatusForbidden = 19005 // 状态不允许操作
	ErrorDatasetisBeingUsing    = 19006 //数据集正在被使用
	ErrorDatasetCacheExist      = 19007 //缓存已存在
	OnlySupportCacheUserDataset = 19008 //只支持缓存用户数据集
	ErrorDatasetCacheNotExist   = 19009 //缓存不存在

	/* 19501~20000 标签管理错误*/
	ErrorLableRefered   = 19501 // 标签被引用，不能删除
	ErrorLableRepeated  = 19502 // 标签重复
	ErrorLableIllegal   = 19503 // 标签不合法
	ErrorLableNotModify = 19504 // 预置标签不可更改

	/* 20001-21000 第三方平台管理错误*/
	ErrorPlatformNameRepeat              = 20001 // 平台名称重复
	ErrorPlatformStorageConfigNameRepeat = 20002 // 平台存储配置名称重复
	ErrorPlatformBatchGetPlatform        = 20003 // 批量获取平台信息错误
	ErrorPlatformConfigValueWrong        = 20004 // 配置值不正确
	ErrorPlatformConfigKeyNotExist       = 20005 // 配置项不存在
	ErrorPlatformRequestFail             = 20006 // http请求失败

	/* 21001-22000 云际错误*/
	ErrorJointCloudRequestFailed = 21001 // 云际请求失败
	ErrorJointCloudNoPermission  = 21002 // 无权限访问

	/* 25001~25000 训练管理错误*/
	ErrorModelDeployForbidden             = 25001 // 部署使用计算框架非TF或者PT
	ErrorModelDeployFailed                = 25002 // 创建模型部署服务失败
	ErrorModelDeployDeleteFailed          = 25003 // 删除模型部署服务失败
	ErrorModelInferRequest                = 25004 // 模型部署服务请求失败
	ErrorModelAuthFailed                  = 25005 // 模型权限校验失败
	ErrorModelDeployResourcePoolForbidden = 25006 // 模型部署使用资源池不允许操作

	/* 23001-24000 prometheus操作相关错误*/
	ErrorPrometheusQueryFailed = 23001 // 查询错误
)

type codeMsg struct {
	codeType CodeType
	msg      string
}

var codeMsgMap = map[int]codeMsg{
	/* 10000~11000 基础错误*/
	ErrorUnknown:                   {codeType: Unknown, msg: "unknown"},
	ErrorStructCopy:                {codeType: DataLoss, msg: "struct copy failed"},
	ErrorInvalidRequestParameter:   {codeType: InvalidArgument, msg: "invalid request parameter"},
	ErrorVersionInvalid:            {codeType: InvalidArgument, msg: "version invalid"},
	ErrorZipFailed:                 {codeType: DataLoss, msg: "zip failed"},
	ErrorUnzipFailed:               {codeType: DataLoss, msg: "unzip failed"},
	ErrorUrlParseFailed:            {codeType: InvalidArgument, msg: "domain or url parase failed"},
	ErrorParseDurationFailed:       {codeType: Internal, msg: "time.Duration parse failed"},
	ErrorEncryptPasswordFailed:     {codeType: Internal, msg: "password encrypt failed"},
	ErrorDirNotExisted:             {codeType: Internal, msg: "dir not existed"},
	ErrorOperateDirFailed:          {codeType: Internal, msg: "dir operate failed"},
	ErrorOperateFileFailed:         {codeType: Internal, msg: "file operate failed"},
	ErrorConfigValueValidateFailed: {codeType: Internal, msg: "config value validate failed"},
	ErrorFormatParseFailed:         {codeType: Internal, msg: "format parse failed"},
	ErrorInternal:                  {codeType: Internal, msg: "internal error"},
	// http请求相关错误
	ErrorHttpNewRequest:     {codeType: Internal, msg: "http new request failed"},
	ErrorHttpDoRequest:      {codeType: Internal, msg: "http do request failed"},
	ErrorHttpReadBody:       {codeType: Internal, msg: "http read body failed"},
	ErrorJsonMarshal:        {codeType: Internal, msg: "json marshal failed"},
	ErrorJsonUnmarshal:      {codeType: Internal, msg: "json unmarshal failed"},
	ErrorHttpWriteFailed:    {codeType: Internal, msg: "http write failed"},
	ErrorHttpBindFormFailed: {codeType: Internal, msg: "http bind form failed"},
	// minio 操作相关错误
	ErrorMinioBucketInitFailed:         {codeType: Internal, msg: "minio init failed"},
	ErrorMinioBucketExisted:            {codeType: AlreadyExists, msg: "minio bucket existed"},
	ErrorMinioBucketNotExist:           {codeType: NotFound, msg: "minio bucket not exist"},
	ErrorMinioCheckBucketExistFailed:   {codeType: Internal, msg: "minio check bucket exist failed"},
	ErrorMinioMakeBucketFailed:         {codeType: Internal, msg: "minio create bucket failed"},
	ErrorMinioDeleteBucketFailed:       {codeType: Internal, msg: "minio delete bucket failed"},
	ErrorMinioPresignedPutObjectFailed: {codeType: Internal, msg: "minio gen upload url failed"},
	ErrorMinioPresignedGetObjectFailed: {codeType: Internal, msg: "minio gen download url failed"},
	ErrorMinioListObjectFailed:         {codeType: Internal, msg: "minio list object failed"},
	ErrorMinioCheckObjectExistFailed:   {codeType: Internal, msg: "minio check object failed"},
	ErrorMinioCheckObjectNotExisted:    {codeType: NotFound, msg: "minio object not exist"},
	ErrorMinioCreateAccountFailed:      {codeType: Internal, msg: "minio create account failed"},
	ErrorMinioOperationFailed:          {codeType: Internal, msg: "minio operation failed"},

	// db 操作相关错误
	ErrorDBInitFailed:        {codeType: Internal, msg: "db init failed"},
	ErrorDBFindFailed:        {codeType: Internal, msg: "db find failed"},
	ErrorDBFirstFailed:       {codeType: Internal, msg: "db first failed"},
	ErrorDBCountFailed:       {codeType: Internal, msg: "db count failed"},
	ErrorDBFindEmpty:         {codeType: NotFound, msg: "db find empty"},
	ErrorDBSelectParamsEmpty: {codeType: InvalidArgument, msg: "db select params empty"},
	ErrorDBUpdateFailed:      {codeType: Internal, msg: "db update failed"},
	ErrorDBCreateFailed:      {codeType: Internal, msg: "db create failed"},
	ErrorDBDeleteFailed:      {codeType: Internal, msg: "db delete delete failed"},
	ErrorDBPrimaryKeyEmpty:   {codeType: InvalidArgument, msg: "db operate without primary key"},
	// k8s操作相关错误
	ErrorK8sCreateServiceFailed: {codeType: Internal, msg: "create k8s service failed"},
	ErrorK8sDeleteServiceFailed: {codeType: Internal, msg: "delete k8s service failed"},
	ErrorK8sCreateIngressFailed: {codeType: Internal, msg: "create k8s ingress failed"},
	ErrorK8sDeleteIngressFailed: {codeType: Internal, msg: "delete k8s ingress failed"},
	// Harbor操作相关错误
	ErrorHarborProjectExists:       {codeType: AlreadyExists, msg: "harbor project exists"},
	ErrorHarborCreateProjectFailed: {codeType: Internal, msg: "create harbor project failed"},
	ErrorHarborCheckProjectFailed:  {codeType: Internal, msg: "check harbor project failed"},
	// redis操作相关错误
	ErroRedisParseUrlFailed:    {codeType: Internal, msg: "redis parse url failed"},
	ErroRedisHSetFailed:        {codeType: Internal, msg: "redis HSet failed"},
	ErroRedisHGetFailed:        {codeType: Internal, msg: "redis HGet failed"},
	ErroRedisHDelFailed:        {codeType: Internal, msg: "redis HDel failed"},
	ErrorRedisLockObtainFailed: {codeType: Internal, msg: "redis lock obtain failed"},

	/* 11001~12000 资源管理错误*/
	ErrorDeleteResourcePool:   {codeType: Internal, msg: "delete ResourcePool failed"},
	ErrorCreateResourcePool:   {codeType: Internal, msg: "create ResourcePool failed"},
	ErrorListResourcePool:     {codeType: Internal, msg: "list ResourcePool failed"},
	ErrorGetResourcePool:      {codeType: Internal, msg: "get ResourcePool failed"},
	ErrorUpdateResourcePool:   {codeType: Internal, msg: "update ResourcePool failed"},
	ErrorListResource:         {codeType: Internal, msg: "list Resource failed"},
	ErrorCreateResource:       {codeType: Internal, msg: "create Resource failed"},
	ErrorDeleteResource:       {codeType: Internal, msg: "delete Resource failed"},
	ErrorUpdateResource:       {codeType: Internal, msg: "update Resource failed"},
	ErrorBindingNode:          {codeType: Internal, msg: "binding node in cluster failed"},
	ErrorCreateResourceSpec:   {codeType: Internal, msg: "create ResourceSpec failed"},
	ErrorListResourceSpec:     {codeType: Internal, msg: "list ResourceSpec failed"},
	ErrorDeleteResourceSpec:   {codeType: Internal, msg: "delete ResourceSpec failed"},
	ErrorGetResourceSpec:      {codeType: Internal, msg: "get ResourceSpec failed"},
	ErrorResourceDiscovery:    {codeType: Internal, msg: "discovery Resource failed"},
	ErrorResourceNotExist:     {codeType: NotFound, msg: "resource not found"},
	ErrorResourceSpecNotExist: {codeType: NotFound, msg: "resourcespec not found"},
	ErrorResourceExist:        {codeType: AlreadyExists, msg: "resources existed"},

	/* 12001~13000 算法管理错误*/
	ErrorFindAlgorithmVersionAccessMaxIdFailed: {codeType: Internal, msg: "find AlgorithmVersionAccess MaxId failed"},
	ErrorFindAlgorithmVersionMaxIdFailed:       {codeType: Internal, msg: "find AlgorithmVersion MaxId failed"},
	ErrorAlgorithmNotMy:                        {codeType: NotFound, msg: "Algorithm not my"},
	ErrorAlgorithmRepeat:                       {codeType: AlreadyExists, msg: "Algorithm Repeat"},
	ErrorAlgorithmVersionRepeat:                {codeType: AlreadyExists, msg: "AlgorithmVersion Repeat"},
	ErrorFindAlgorithmAuthWrong:                {codeType: NotFound, msg: "Algorithm Auth Wrong"},
	ErrorAlgorithmVersionFileNotFound:          {codeType: NotFound, msg: "AlgorithmVersion File NotFound"},
	ErrorAlgorithmVersionFileNotReady:          {codeType: AlreadyExists, msg: "AlgorithmVersion File NotReady"},
	ErrorAlgorithmVersionFileExisted:           {codeType: AlreadyExists, msg: "AlgorithmVersion FileExists"},
	ErrorAlgorithmAccessVersionExisted:         {codeType: AlreadyExists, msg: "AlgorithmAccessVersion Exists"},
	ErrorAlgorithmVersionUploadAuth:            {codeType: InvalidArgument, msg: "AlgorithmVersionUpload Auth Wrong"},

	/* 13001~14000 镜像管理错误*/
	ErrorImageStatusMakeError:      {codeType: OutOfRange, msg: "make image status error"},
	ErrorImageSourceTypeToUpload:   {codeType: OutOfRange, msg: "error image source to upload"},
	ErrorImageExisted:              {codeType: AlreadyExists, msg: "image exists"},
	ErrorImageNotExist:             {codeType: NotFound, msg: "image not exists"},
	ErrorImageOpForbidden:          {codeType: PermissionDenied, msg: "image operation forbidden"},
	ErrorImageSourceType:           {codeType: Unimplemented, msg: "error image source type"},
	ErrorImageContainerCommitError: {codeType: Unimplemented, msg: "error commit container"},
	ErrorImagePushFailed:           {codeType: OutOfRange, msg: "error push image"},

	/* 14001~15000 开发管理错误*/
	ErrorNotebookStatusForbidden:          {codeType: OutOfRange, msg: "status forbidden"},
	ErrorNotebookImageNoPermission:        {codeType: PermissionDenied, msg: "image no permission"},
	ErrorNotebookAlgorithmNoPermission:    {codeType: PermissionDenied, msg: "algorithm no permission"},
	ErrorNotebookBalanceNotEnough:         {codeType: OutOfRange, msg: "balance not enough"},
	ErrorNotebookParseResourceSpecFailed:  {codeType: Internal, msg: "parse resource spec failed"},
	ErrorNotebookImageStatusForbidden:     {codeType: OutOfRange, msg: "image status forbidden"},
	ErrorNotebookAlgorithmStatusForbidden: {codeType: OutOfRange, msg: "algorithm status forbidden"},
	ErrorNotebookRepeat:                   {codeType: AlreadyExists, msg: "notebook repeat"},
	ErrorNotebookDatasetNoPermission:      {codeType: PermissionDenied, msg: "dataset no permission"},
	ErrorNotebookDatasetStatusForbidden:   {codeType: OutOfRange, msg: "dataset status forbidden"},
	ErrorNotebookRepeatedToSave:           {codeType: AlreadyExists, msg: "repeated notebook saveing"},
	ErrorNotebookNoFoundRuntimeContainer:  {codeType: NotFound, msg: "no found runtime container"},
	ErrorNotebookResourcePoolForbidden:    {codeType: OutOfRange, msg: "resource pool forbidden"},
	ErrorNotebookMountExternalForbidden:   {codeType: OutOfRange, msg: "mount external forbidden"},

	/* 15001~16000 训练管理错误*/
	ErrorTrainImageForbidden:         {codeType: PermissionDenied, msg: "image Auth forbidden"},
	ErrorTrainDataSetForbidden:       {codeType: PermissionDenied, msg: "dataset Auth forbidden"},
	ErrorTrainAlgorithmForbidden:     {codeType: PermissionDenied, msg: "algorithm Auth forbidden"},
	ErrorPipelineDoRequest:           {codeType: Internal, msg: "do pipeline request failed"},
	ErrorDeleteJobRequest:            {codeType: Internal, msg: "delete running job request failed"},
	ErrorTrainBalanceNotEnough:       {codeType: OutOfRange, msg: "balance not enough"},
	ErrorStopTerminatedJob:           {codeType: Unimplemented, msg: "do stop terminated job failed"},
	ErrorJobImageStatusForbidden:     {codeType: OutOfRange, msg: "image status forbidden"},
	ErrorJobAlgorithmStatusForbidden: {codeType: OutOfRange, msg: "algorithm status forbidden"},
	ErrorJobDatasetStatusForbidden:   {codeType: OutOfRange, msg: "dataset status forbidden"},
	ErrorJobUniqueIndexConflict:      {codeType: InvalidArgument, msg: "repeated job name causes conflict unique key"},
	ErrorRepeatJobConfigName:         {codeType: InvalidArgument, msg: "repeated config name in distributed job"},
	ErrorTrainResourcePoolForbidden:  {codeType: InvalidArgument, msg: "train resource pool forbidden"},
	ErrorTrainMountExternalForbidden: {codeType: OutOfRange, msg: "mount external forbidden"},

	/* 16001~17000 用户管理错误*/
	// 用户登录相关错误
	ErrorNotAuthorized:            {codeType: InvalidArgument, msg: "not authorized"},
	ErrorAuthenticationForbidden:  {codeType: PermissionDenied, msg: "user account authentication forbidden"},
	ErrorCreateTokenFailed:        {codeType: Internal, msg: "create token failed"},
	ErrorParseTokenFailed:         {codeType: OutOfRange, msg: "parse token failed"},
	ErrorTokenInvalid:             {codeType: InvalidArgument, msg: "token invalid"},
	ErrorSessionIdNotFound:        {codeType: InvalidArgument, msg: "sessionId not found"},
	ErrorUserNoAuthSession:        {codeType: Unauthorized, msg: "user no auth session"},
	ErrorUserGetAuthSessionFailed: {codeType: Unauthorized, msg: "get session failed"},
	ErrorAuthenticationFailed:     {codeType: Unauthorized, msg: "user account authentication failed"},
	ErrorTokenRenew:               {codeType: InvalidArgument, msg: "token renew"},
	// 用户管理
	ErrorUserAccountNotExisted:            {codeType: NotFound, msg: "user account not exists"},
	ErrorUserAccountExisted:               {codeType: AlreadyExists, msg: "user account exists"},
	ErrorUserAccountBinded:                {codeType: AlreadyExists, msg: "user account exists"},
	ErrorUserIdNotRight:                   {codeType: InvalidArgument, msg: "userid not valid"},
	ErrorWorkSpaceExisted:                 {codeType: AlreadyExists, msg: "workspace existed"},
	ErrorWorkSpaceNotExist:                {codeType: NotFound, msg: "workspace not existed"},
	ErrorUserWorkSpaceNoPermission:        {codeType: PermissionDenied, msg: "user workspace permission deny"},
	ErrorWorkSpaceResourcePoolBound:       {codeType: ResourceExhausted, msg: "workspace and resource pool had bind"},
	ErrorUserConfigKeyNotExist:            {codeType: InvalidArgument, msg: "user config key not exist"},
	ErrorUserChangeMinioUsernameForbidden: {codeType: InvalidArgument, msg: "change minio username forbidden"},
	ErrorUserMinioUsernameNotExist:        {codeType: InvalidArgument, msg: "minio username not exist"},

	/* 17001~18000 机时管理错误*/
	ErrorBillingObtainLockFailed: {codeType: Internal, msg: "billing obtain lock failed"},
	ErrorBillingStatusForbidden:  {codeType: OutOfRange, msg: "status forbidden"},

	/* 18001~19000 模型管理错误*/
	ErrorModelNoPermission:              {codeType: PermissionDenied, msg: "model no permission to operate"},
	ErrorModelRepeat:                    {codeType: AlreadyExists, msg: "model repeat, one algorithm version just one model"},
	ErrorModelVersionRepeat:             {codeType: AlreadyExists, msg: "model version repeat"},
	ErrorModelVersionFileNotFound:       {codeType: NotFound, msg: "model version file not existed"},
	ErrorModelVersionFileExisted:        {codeType: AlreadyExists, msg: "model version file existed"},
	ErrorModelCreateAlgorithmNotExisted: {codeType: NotFound, msg: "model create algorithm not found"},

	/* 19001~20000 数据集管理错误*/
	ErrorDatasetFileNotFound:    {codeType: NotFound, msg: "dataset file not found"},
	ErrorDatasetAlreadyShared:   {codeType: OutOfRange, msg: "dataset already shared"},
	ErrorDatasetNoPermission:    {codeType: PermissionDenied, msg: "no permission"},
	ErrorDatasetRepeat:          {codeType: AlreadyExists, msg: "dataset repeat"},
	ErrorDatasetStatusForbidden: {codeType: OutOfRange, msg: "status forbidden"},
	ErrorDatasetisBeingUsing:    {codeType: PermissionDenied, msg: "dataset is being used"},
	ErrorDatasetCacheExist:      {codeType: AlreadyExists, msg: "dataset cache exist"},
	OnlySupportCacheUserDataset: {codeType: OutOfRange, msg: "only support cache user dataset "},
	ErrorDatasetCacheNotExist:   {codeType: OutOfRange, msg: "dataset cache not exist "},

	/* 19501~20000 标签管理错误*/
	ErrorLableRefered:   {codeType: Unimplemented, msg: "lable refered"},
	ErrorLableRepeated:  {codeType: AlreadyExists, msg: "lable repeated"},
	ErrorLableIllegal:   {codeType: InvalidArgument, msg: "lable illegal"},
	ErrorLableNotModify: {codeType: OutOfRange, msg: "lable not modify"},

	/* 20001-21000 第三方平台管理错误*/
	ErrorPlatformNameRepeat:              {codeType: AlreadyExists, msg: "platform existed"},
	ErrorPlatformStorageConfigNameRepeat: {codeType: AlreadyExists, msg: "platform storage config existed"},
	ErrorPlatformConfigValueWrong:        {codeType: InvalidArgument, msg: "platform config value wrong"},
	ErrorPlatformConfigKeyNotExist:       {codeType: InvalidArgument, msg: "platform config key not exist"},

	/* 21001-22000 云际请求错误*/
	ErrorJointCloudRequestFailed: {codeType: Internal, msg: "joint cloud request failed"},
	ErrorJointCloudNoPermission:  {codeType: PermissionDenied, msg: "no permission"},

	/* 25001~26000 数据集管理错误*/
	ErrorModelAuthFailed:                  {codeType: Internal, msg: "model can`t access"},
	ErrorModelDeployResourcePoolForbidden: {codeType: Internal, msg: "resource pool forbidden"},

	/* 23001-24000 prometheus操作相关错误*/
	ErrorPrometheusQueryFailed: {codeType: Internal, msg: "prometheus query failed"}, // 查询错误
}
