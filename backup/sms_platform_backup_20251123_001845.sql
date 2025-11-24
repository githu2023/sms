-- MySQL dump 10.13  Distrib 9.5.0, for macos14.8 (arm64)
--
-- Host: 127.0.0.1    Database: sms_platform
-- ------------------------------------------------------
-- Server version	9.5.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN= 0;

--
-- GTID state at the beginning of the backup 
--

-- SET @@GLOBAL.GTID_PURGED=/*!80000 '+'*/ 'c41549ee-60d2-11eb-8bc7-b9cbd0affad8:1-234';

--
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=952 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (939,'p','888','/api/createApi','POST','','',''),(938,'p','888','/api/deleteApi','POST','','',''),(933,'p','888','/api/deleteApisByIds','DELETE','','',''),(930,'p','888','/api/enterSyncApi','POST','','',''),(935,'p','888','/api/getAllApis','POST','','',''),(934,'p','888','/api/getApiById','POST','','',''),(931,'p','888','/api/getApiGroups','GET','','',''),(936,'p','888','/api/getApiList','POST','','',''),(929,'p','888','/api/ignoreApi','POST','','',''),(932,'p','888','/api/syncApi','GET','','',''),(937,'p','888','/api/updateApi','POST','','',''),(825,'p','888','/attachmentCategory/addCategory','POST','','',''),(824,'p','888','/attachmentCategory/deleteCategory','POST','','',''),(826,'p','888','/attachmentCategory/getCategoryList','GET','','',''),(928,'p','888','/authority/copyAuthority','POST','','',''),(927,'p','888','/authority/createAuthority','POST','','',''),(926,'p','888','/authority/deleteAuthority','POST','','',''),(924,'p','888','/authority/getAuthorityList','POST','','',''),(923,'p','888','/authority/setDataAuthority','POST','','',''),(925,'p','888','/authority/updateAuthority','PUT','','',''),(849,'p','888','/authorityBtn/canRemoveAuthorityBtn','POST','','',''),(850,'p','888','/authorityBtn/getAuthorityBtn','POST','','',''),(851,'p','888','/authorityBtn/setAuthorityBtn','POST','','',''),(876,'p','888','/autoCode/addFunc','POST','','',''),(884,'p','888','/autoCode/createPackage','POST','','',''),(892,'p','888','/autoCode/createTemp','POST','','',''),(881,'p','888','/autoCode/delPackage','POST','','',''),(877,'p','888','/autoCode/delSysHistory','POST','','',''),(890,'p','888','/autoCode/getColumn','GET','','',''),(894,'p','888','/autoCode/getDB','GET','','',''),(880,'p','888','/autoCode/getMeta','POST','','',''),(882,'p','888','/autoCode/getPackage','POST','','',''),(878,'p','888','/autoCode/getSysHistory','POST','','',''),(893,'p','888','/autoCode/getTables','GET','','',''),(883,'p','888','/autoCode/getTemplates','GET','','',''),(889,'p','888','/autoCode/installPlugin','POST','','',''),(887,'p','888','/autoCode/mcp','POST','','',''),(885,'p','888','/autoCode/mcpList','POST','','',''),(886,'p','888','/autoCode/mcpTest','POST','','',''),(891,'p','888','/autoCode/preview','POST','','',''),(888,'p','888','/autoCode/pubPlug','POST','','',''),(879,'p','888','/autoCode/rollback','POST','','',''),(921,'p','888','/casbin/getPolicyPathByAuthorityId','POST','','',''),(922,'p','888','/casbin/updateCasbin','POST','','',''),(897,'p','888','/customer/customer','DELETE','','',''),(896,'p','888','/customer/customer','GET','','',''),(898,'p','888','/customer/customer','POST','','',''),(899,'p','888','/customer/customer','PUT','','',''),(895,'p','888','/customer/customerList','GET','','',''),(853,'p','888','/email/emailTest','POST','','',''),(852,'p','888','/email/sendEmail','POST','','',''),(910,'p','888','/fileUploadAndDownload/breakpointContinue','POST','','',''),(909,'p','888','/fileUploadAndDownload/breakpointContinueFinish','POST','','',''),(906,'p','888','/fileUploadAndDownload/deleteFile','POST','','',''),(905,'p','888','/fileUploadAndDownload/editFileName','POST','','',''),(911,'p','888','/fileUploadAndDownload/findFile','GET','','',''),(904,'p','888','/fileUploadAndDownload/getFileList','POST','','',''),(903,'p','888','/fileUploadAndDownload/importURL','POST','','',''),(908,'p','888','/fileUploadAndDownload/removeChunk','POST','','',''),(907,'p','888','/fileUploadAndDownload/upload','POST','','',''),(839,'p','888','/info/createInfo','POST','','',''),(838,'p','888','/info/deleteInfo','DELETE','','',''),(837,'p','888','/info/deleteInfoByIds','DELETE','','',''),(835,'p','888','/info/findInfo','GET','','',''),(834,'p','888','/info/getInfoList','GET','','',''),(836,'p','888','/info/updateInfo','PUT','','',''),(951,'p','888','/jwt/jsonInBlacklist','POST','','',''),(920,'p','888','/menu/addBaseMenu','POST','','',''),(912,'p','888','/menu/addMenuAuthority','POST','','',''),(918,'p','888','/menu/deleteBaseMenu','POST','','',''),(916,'p','888','/menu/getBaseMenuById','POST','','',''),(914,'p','888','/menu/getBaseMenuTree','POST','','',''),(919,'p','888','/menu/getMenu','POST','','',''),(913,'p','888','/menu/getMenuAuthority','POST','','',''),(915,'p','888','/menu/getMenuList','POST','','',''),(917,'p','888','/menu/updateBaseMenu','POST','','',''),(855,'p','888','/simpleUploader/checkFileMd5','GET','','',''),(854,'p','888','/simpleUploader/mergeFileMd5','GET','','',''),(856,'p','888','/simpleUploader/upload','POST','','',''),(816,'p','888','/smsApiLogs/createSmsApiLogs','POST','','',''),(815,'p','888','/smsApiLogs/deleteSmsApiLogs','DELETE','','',''),(814,'p','888','/smsApiLogs/deleteSmsApiLogsByIds','DELETE','','',''),(812,'p','888','/smsApiLogs/findSmsApiLogs','GET','','',''),(811,'p','888','/smsApiLogs/getSmsApiLogsList','GET','','',''),(813,'p','888','/smsApiLogs/updateSmsApiLogs','PUT','','',''),(780,'p','888','/smsCustomers/createSmsCustomers','POST','','',''),(774,'p','888','/smsCustomers/creditDebit','POST','','',''),(779,'p','888','/smsCustomers/deleteSmsCustomers','DELETE','','',''),(778,'p','888','/smsCustomers/deleteSmsCustomersByIds','DELETE','','',''),(776,'p','888','/smsCustomers/findSmsCustomers','GET','','',''),(775,'p','888','/smsCustomers/getSmsCustomersList','GET','','',''),(777,'p','888','/smsCustomers/updateSmsCustomers','PUT','','',''),(804,'p','888','/smsIpWhitelist/createSmsIpWhitelist','POST','','',''),(803,'p','888','/smsIpWhitelist/deleteSmsIpWhitelist','DELETE','','',''),(802,'p','888','/smsIpWhitelist/deleteSmsIpWhitelistByIds','DELETE','','',''),(800,'p','888','/smsIpWhitelist/findSmsIpWhitelist','GET','','',''),(799,'p','888','/smsIpWhitelist/getSmsIpWhitelistList','GET','','',''),(801,'p','888','/smsIpWhitelist/updateSmsIpWhitelist','PUT','','',''),(798,'p','888','/smsPhoneAssignments/createSmsPhoneAssignments','POST','','',''),(797,'p','888','/smsPhoneAssignments/deleteSmsPhoneAssignments','DELETE','','',''),(796,'p','888','/smsPhoneAssignments/deleteSmsPhoneAssignmentsByIds','DELETE','','',''),(794,'p','888','/smsPhoneAssignments/findSmsPhoneAssignments','GET','','',''),(793,'p','888','/smsPhoneAssignments/getSmsPhoneAssignmentsList','GET','','',''),(795,'p','888','/smsPhoneAssignments/updateSmsPhoneAssignments','PUT','','',''),(767,'p','888','/smsPlatformBusinessTypes/createSmsPlatformBusinessTypes','POST','','',''),(766,'p','888','/smsPlatformBusinessTypes/deleteSmsPlatformBusinessTypes','DELETE','','',''),(765,'p','888','/smsPlatformBusinessTypes/deleteSmsPlatformBusinessTypesByIds','DELETE','','',''),(763,'p','888','/smsPlatformBusinessTypes/findSmsPlatformBusinessTypes','GET','','',''),(762,'p','888','/smsPlatformBusinessTypes/getSmsPlatformBusinessTypesList','GET','','',''),(764,'p','888','/smsPlatformBusinessTypes/updateSmsPlatformBusinessTypes','PUT','','',''),(761,'p','888','/smsPlatformProviderBusinessMapping/createSmsPlatformProviderBusinessMapping','POST','','',''),(760,'p','888','/smsPlatformProviderBusinessMapping/deleteSmsPlatformProviderBusinessMapping','DELETE','','',''),(759,'p','888','/smsPlatformProviderBusinessMapping/deleteSmsPlatformProviderBusinessMappingByIds','DELETE','','',''),(757,'p','888','/smsPlatformProviderBusinessMapping/findSmsPlatformProviderBusinessMapping','GET','','',''),(756,'p','888','/smsPlatformProviderBusinessMapping/getSmsPlatformProviderBusinessMappingList','GET','','',''),(758,'p','888','/smsPlatformProviderBusinessMapping/updateSmsPlatformProviderBusinessMapping','PUT','','',''),(792,'p','888','/smsProviders/createSmsProviders','POST','','',''),(791,'p','888','/smsProviders/deleteSmsProviders','DELETE','','',''),(790,'p','888','/smsProviders/deleteSmsProvidersByIds','DELETE','','',''),(788,'p','888','/smsProviders/findSmsProviders','GET','','',''),(787,'p','888','/smsProviders/getSmsProvidersList','GET','','',''),(789,'p','888','/smsProviders/updateSmsProviders','PUT','','',''),(773,'p','888','/smsProvidersBusinessTypes/createSmsProvidersBusinessTypes','POST','','',''),(772,'p','888','/smsProvidersBusinessTypes/deleteSmsProvidersBusinessTypes','DELETE','','',''),(771,'p','888','/smsProvidersBusinessTypes/deleteSmsProvidersBusinessTypesByIds','DELETE','','',''),(769,'p','888','/smsProvidersBusinessTypes/findSmsProvidersBusinessTypes','GET','','',''),(768,'p','888','/smsProvidersBusinessTypes/getSmsProvidersBusinessTypesList','GET','','',''),(770,'p','888','/smsProvidersBusinessTypes/updateSmsProvidersBusinessTypes','PUT','','',''),(786,'p','888','/smsTransactions/createSmsTransactions','POST','','',''),(785,'p','888','/smsTransactions/deleteSmsTransactions','DELETE','','',''),(784,'p','888','/smsTransactions/deleteSmsTransactionsByIds','DELETE','','',''),(782,'p','888','/smsTransactions/findSmsTransactions','GET','','',''),(781,'p','888','/smsTransactions/getSmsTransactionsList','GET','','',''),(783,'p','888','/smsTransactions/updateSmsTransactions','PUT','','',''),(866,'p','888','/sysDictionary/createSysDictionary','POST','','',''),(865,'p','888','/sysDictionary/deleteSysDictionary','DELETE','','',''),(863,'p','888','/sysDictionary/findSysDictionary','GET','','',''),(862,'p','888','/sysDictionary/getSysDictionaryList','GET','','',''),(864,'p','888','/sysDictionary/updateSysDictionary','PUT','','',''),(874,'p','888','/sysDictionaryDetail/createSysDictionaryDetail','POST','','',''),(873,'p','888','/sysDictionaryDetail/deleteSysDictionaryDetail','DELETE','','',''),(872,'p','888','/sysDictionaryDetail/findSysDictionaryDetail','GET','','',''),(868,'p','888','/sysDictionaryDetail/getDictionaryDetailsByParent','GET','','',''),(867,'p','888','/sysDictionaryDetail/getDictionaryPath','GET','','',''),(870,'p','888','/sysDictionaryDetail/getDictionaryTreeList','GET','','',''),(869,'p','888','/sysDictionaryDetail/getDictionaryTreeListByType','GET','','',''),(871,'p','888','/sysDictionaryDetail/getSysDictionaryDetailList','GET','','',''),(875,'p','888','/sysDictionaryDetail/updateSysDictionaryDetail','PUT','','',''),(848,'p','888','/sysExportTemplate/createSysExportTemplate','POST','','',''),(847,'p','888','/sysExportTemplate/deleteSysExportTemplate','DELETE','','',''),(846,'p','888','/sysExportTemplate/deleteSysExportTemplateByIds','DELETE','','',''),(842,'p','888','/sysExportTemplate/exportExcel','GET','','',''),(841,'p','888','/sysExportTemplate/exportTemplate','GET','','',''),(844,'p','888','/sysExportTemplate/findSysExportTemplate','GET','','',''),(843,'p','888','/sysExportTemplate/getSysExportTemplateList','GET','','',''),(840,'p','888','/sysExportTemplate/importExcel','POST','','',''),(845,'p','888','/sysExportTemplate/updateSysExportTemplate','PUT','','',''),(861,'p','888','/sysOperationRecord/createSysOperationRecord','POST','','',''),(858,'p','888','/sysOperationRecord/deleteSysOperationRecord','DELETE','','',''),(857,'p','888','/sysOperationRecord/deleteSysOperationRecordByIds','DELETE','','',''),(860,'p','888','/sysOperationRecord/findSysOperationRecord','GET','','',''),(859,'p','888','/sysOperationRecord/getSysOperationRecordList','GET','','',''),(833,'p','888','/sysParams/createSysParams','POST','','',''),(832,'p','888','/sysParams/deleteSysParams','DELETE','','',''),(831,'p','888','/sysParams/deleteSysParamsByIds','DELETE','','',''),(829,'p','888','/sysParams/findSysParams','GET','','',''),(827,'p','888','/sysParams/getSysParam','GET','','',''),(828,'p','888','/sysParams/getSysParamsList','GET','','',''),(830,'p','888','/sysParams/updateSysParams','PUT','','',''),(818,'p','888','/sysVersion/deleteSysVersion','DELETE','','',''),(817,'p','888','/sysVersion/deleteSysVersionByIds','DELETE','','',''),(821,'p','888','/sysVersion/downloadVersionJson','GET','','',''),(820,'p','888','/sysVersion/exportVersion','POST','','',''),(823,'p','888','/sysVersion/findSysVersion','GET','','',''),(822,'p','888','/sysVersion/getSysVersionList','GET','','',''),(819,'p','888','/sysVersion/importVersion','POST','','',''),(902,'p','888','/system/getServerInfo','POST','','',''),(901,'p','888','/system/getSystemConfig','POST','','',''),(900,'p','888','/system/setSystemConfig','POST','','',''),(949,'p','888','/user/admin_register','POST','','',''),(943,'p','888','/user/changePassword','POST','','',''),(950,'p','888','/user/deleteUser','DELETE','','',''),(945,'p','888','/user/getUserInfo','GET','','',''),(948,'p','888','/user/getUserList','POST','','',''),(941,'p','888','/user/resetPassword','POST','','',''),(946,'p','888','/user/setSelfInfo','PUT','','',''),(940,'p','888','/user/setSelfSetting','PUT','','',''),(944,'p','888','/user/setUserAuthorities','POST','','',''),(942,'p','888','/user/setUserAuthority','POST','','',''),(947,'p','888','/user/setUserInfo','PUT','','',''),(139,'p','8881','/api/createApi','POST','','',''),(142,'p','8881','/api/deleteApi','POST','','',''),(144,'p','8881','/api/getAllApis','POST','','',''),(141,'p','8881','/api/getApiById','POST','','',''),(140,'p','8881','/api/getApiList','POST','','',''),(143,'p','8881','/api/updateApi','POST','','',''),(145,'p','8881','/authority/createAuthority','POST','','',''),(146,'p','8881','/authority/deleteAuthority','POST','','',''),(147,'p','8881','/authority/getAuthorityList','POST','','',''),(148,'p','8881','/authority/setDataAuthority','POST','','',''),(167,'p','8881','/casbin/getPolicyPathByAuthorityId','POST','','',''),(166,'p','8881','/casbin/updateCasbin','POST','','',''),(173,'p','8881','/customer/customer','DELETE','','',''),(174,'p','8881','/customer/customer','GET','','',''),(171,'p','8881','/customer/customer','POST','','',''),(172,'p','8881','/customer/customer','PUT','','',''),(175,'p','8881','/customer/customerList','GET','','',''),(163,'p','8881','/fileUploadAndDownload/deleteFile','POST','','',''),(164,'p','8881','/fileUploadAndDownload/editFileName','POST','','',''),(162,'p','8881','/fileUploadAndDownload/getFileList','POST','','',''),(165,'p','8881','/fileUploadAndDownload/importURL','POST','','',''),(161,'p','8881','/fileUploadAndDownload/upload','POST','','',''),(168,'p','8881','/jwt/jsonInBlacklist','POST','','',''),(151,'p','8881','/menu/addBaseMenu','POST','','',''),(153,'p','8881','/menu/addMenuAuthority','POST','','',''),(155,'p','8881','/menu/deleteBaseMenu','POST','','',''),(157,'p','8881','/menu/getBaseMenuById','POST','','',''),(152,'p','8881','/menu/getBaseMenuTree','POST','','',''),(149,'p','8881','/menu/getMenu','POST','','',''),(154,'p','8881','/menu/getMenuAuthority','POST','','',''),(150,'p','8881','/menu/getMenuList','POST','','',''),(156,'p','8881','/menu/updateBaseMenu','POST','','',''),(169,'p','8881','/system/getSystemConfig','POST','','',''),(170,'p','8881','/system/setSystemConfig','POST','','',''),(138,'p','8881','/user/admin_register','POST','','',''),(158,'p','8881','/user/changePassword','POST','','',''),(176,'p','8881','/user/getUserInfo','GET','','',''),(159,'p','8881','/user/getUserList','POST','','',''),(160,'p','8881','/user/setUserAuthority','POST','','',''),(178,'p','9528','/api/createApi','POST','','',''),(181,'p','9528','/api/deleteApi','POST','','',''),(183,'p','9528','/api/getAllApis','POST','','',''),(180,'p','9528','/api/getApiById','POST','','',''),(179,'p','9528','/api/getApiList','POST','','',''),(182,'p','9528','/api/updateApi','POST','','',''),(184,'p','9528','/authority/createAuthority','POST','','',''),(185,'p','9528','/authority/deleteAuthority','POST','','',''),(186,'p','9528','/authority/getAuthorityList','POST','','',''),(187,'p','9528','/authority/setDataAuthority','POST','','',''),(215,'p','9528','/autoCode/createTemp','POST','','',''),(206,'p','9528','/casbin/getPolicyPathByAuthorityId','POST','','',''),(205,'p','9528','/casbin/updateCasbin','POST','','',''),(213,'p','9528','/customer/customer','DELETE','','',''),(211,'p','9528','/customer/customer','GET','','',''),(212,'p','9528','/customer/customer','POST','','',''),(210,'p','9528','/customer/customer','PUT','','',''),(214,'p','9528','/customer/customerList','GET','','',''),(202,'p','9528','/fileUploadAndDownload/deleteFile','POST','','',''),(203,'p','9528','/fileUploadAndDownload/editFileName','POST','','',''),(201,'p','9528','/fileUploadAndDownload/getFileList','POST','','',''),(204,'p','9528','/fileUploadAndDownload/importURL','POST','','',''),(200,'p','9528','/fileUploadAndDownload/upload','POST','','',''),(207,'p','9528','/jwt/jsonInBlacklist','POST','','',''),(190,'p','9528','/menu/addBaseMenu','POST','','',''),(192,'p','9528','/menu/addMenuAuthority','POST','','',''),(194,'p','9528','/menu/deleteBaseMenu','POST','','',''),(196,'p','9528','/menu/getBaseMenuById','POST','','',''),(191,'p','9528','/menu/getBaseMenuTree','POST','','',''),(188,'p','9528','/menu/getMenu','POST','','',''),(193,'p','9528','/menu/getMenuAuthority','POST','','',''),(189,'p','9528','/menu/getMenuList','POST','','',''),(195,'p','9528','/menu/updateBaseMenu','POST','','',''),(208,'p','9528','/system/getSystemConfig','POST','','',''),(209,'p','9528','/system/setSystemConfig','POST','','',''),(177,'p','9528','/user/admin_register','POST','','',''),(197,'p','9528','/user/changePassword','POST','','',''),(216,'p','9528','/user/getUserInfo','GET','','',''),(198,'p','9528','/user/getUserList','POST','','',''),(199,'p','9528','/user/setUserAuthority','POST','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_attachment_category`
--

DROP TABLE IF EXISTS `exa_attachment_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_attachment_category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '分类名称',
  `pid` bigint DEFAULT '0' COMMENT '父节点ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_attachment_category_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_attachment_category`
--

LOCK TABLES `exa_attachment_category` WRITE;
/*!40000 ALTER TABLE `exa_attachment_category` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_attachment_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_customers`
--

DROP TABLE IF EXISTS `exa_customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_customers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint unsigned DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint unsigned DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_customers`
--

LOCK TABLES `exa_customers` WRITE;
/*!40000 ALTER TABLE `exa_customers` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_file_chunks`
--

DROP TABLE IF EXISTS `exa_file_chunks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_file_chunks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `exa_file_id` bigint unsigned DEFAULT NULL,
  `file_chunk_number` bigint DEFAULT NULL,
  `file_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_file_chunks`
--

LOCK TABLES `exa_file_chunks` WRITE;
/*!40000 ALTER TABLE `exa_file_chunks` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_file_chunks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_file_upload_and_downloads`
--

DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_file_upload_and_downloads` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '文件名',
  `class_id` bigint DEFAULT '0' COMMENT '分类id',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_file_upload_and_downloads`
--

LOCK TABLES `exa_file_upload_and_downloads` WRITE;
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` DISABLE KEYS */;
INSERT INTO `exa_file_upload_and_downloads` VALUES (1,'2025-11-21 17:18:36.120','2025-11-21 17:18:36.120',NULL,'10.png',0,'https://qmplusimg.henrongyi.top/gvalogo.png','png','158787308910.png'),(2,'2025-11-21 17:18:36.120','2025-11-21 17:18:36.120',NULL,'logo.png',0,'https://qmplusimg.henrongyi.top/1576554439myAvatar.png','png','1587973709logo.png');
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exa_files`
--

DROP TABLE IF EXISTS `exa_files`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_files` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `chunk_total` bigint DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_files_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exa_files`
--

LOCK TABLES `exa_files` WRITE;
/*!40000 ALTER TABLE `exa_files` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_files` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gva_announcements_info`
--

DROP TABLE IF EXISTS `gva_announcements_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gva_announcements_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '公告标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '公告内容',
  `user_id` bigint DEFAULT NULL COMMENT '发布者',
  `attachments` json DEFAULT NULL COMMENT '相关附件',
  PRIMARY KEY (`id`),
  KEY `idx_gva_announcements_info_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gva_announcements_info`
--

LOCK TABLES `gva_announcements_info` WRITE;
/*!40000 ALTER TABLE `gva_announcements_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `gva_announcements_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jwt_blacklists`
--

DROP TABLE IF EXISTS `jwt_blacklists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jwt_blacklists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jwt_blacklists`
--

LOCK TABLES `jwt_blacklists` WRITE;
/*!40000 ALTER TABLE `jwt_blacklists` DISABLE KEYS */;
/*!40000 ALTER TABLE `jwt_blacklists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sms_api_logs`
--

DROP TABLE IF EXISTS `sms_api_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sms_api_logs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `customer_id` bigint DEFAULT NULL COMMENT '客户ID',
  `request_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '请求来源IP',
  `request_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '请求的API路径',
  `request_body` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求体内容',
  `response_code` int DEFAULT NULL COMMENT 'HTTP响应状态码',
  `duration_ms` int DEFAULT NULL COMMENT '请求处理耗时(毫秒)',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_customer_id` (`customer_id`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_sms_api_logs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='API请求日志表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sms_api_logs`
--

LOCK TABLES `sms_api_logs` WRITE;
/*!40000 ALTER TABLE `sms_api_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `sms_api_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sms_customer_business_config`
--

DROP TABLE IF EXISTS `sms_customer_business_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sms_customer_business_config` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_id` bigint NOT NULL COMMENT '商户ID',
  `platform_business_type_id` bigint NOT NULL COMMENT '平台业务类型ID',
  `business_code` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '业务编码',
  `business_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '业务名称',
  `weight` int NOT NULL DEFAULT '1' COMMENT '权重（用于随机选择，权重越高被选中概率越大）',
  `status` tinyint(1) DEFAULT '1' COMMENT '是否启用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_customer_business` (`customer_id`,`platform_business_type_id`),
  KEY `idx_customer_id` (`customer_id`),
  KEY `idx_business_code` (`business_code`),
  KEY `idx_sms_customer_business_config_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='商户业务配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sms_customer_business_config`
--

LOCK TABLES `sms_customer_business_config` WRITE;
/*!40000 ALTER TABLE `sms_customer_business_config` DISABLE KEYS */;
/*!40000 ALTER TABLE `sms_customer_business_config` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sms_customers`
--

DROP TABLE IF EXISTS `sms_customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sms_customers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `merchant_name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '商户名称',
  `merchant_no` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '商户号',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '客户端登录用户名',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '客户端登录邮箱',
  `password_hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '客户端登录用的密码哈希',
  `api_secret_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用于生成API Token的唯一密钥',
  `balance` decimal(10,2) DEFAULT '0.00' COMMENT '客户余额',
  `parent_id` bigint DEFAULT NULL COMMENT '上级商户ID',
  `frozen_amount` decimal(10,2) DEFAULT '0.00' COMMENT '冻结金额',
  `status` tinyint(1) DEFAULT NULL COMMENT '客户状态 (1:正常, 0:冻结)',
  `registration_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '注册时的IP地址',
  `last_login_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '最后一次登录的IP地址',
  `last_login_at` datetime(3) DEFAULT NULL COMMENT '最后一次登录的时间',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `remark` varchar(500) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_sms_customers_merchant_no` (`merchant_no`),
  KEY `idx_sms_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='客户信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sms_customers`
--

LOCK TABLES `sms_customers` WRITE;
/*!40000 ALTER TABLE `sms_customers` DISABLE KEYS */;
INSERT INTO `sms_customers` VALUES (1,'12121','121212','test','111@qq.com','$2a$10$xPBwRscV2SEnj38K5tzd0uFy5JKeQs8W6VkWGvG2fqt6c4Ho..gGy','676494626b8738264235e69a781bbac50e8be469fba02ea51bc0c48142274711',1003.00,NULL,0.00,1,'127.0.0.1',NULL,NULL,'2025-11-21 17:53:45.730','2025-11-22 23:51:53.667',NULL,NULL);
/*!40000 ALTER TABLE `sms_customers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sms_ip_whitelist`
--

DROP TABLE IF EXISTS `sms_ip_whitelist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sms_ip_whitelist` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `customer_id` bigint NOT NULL COMMENT '客户ID',
  `ip_address` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '白名单IP或IP段',
  `notes` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注, 例如 ',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_customer_ip` (`customer_id`,`ip_address`),
  KEY `idx_sms_ip_whitelist_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='API IP白名单表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sms_ip_whitelist`
--

LOCK TABLES `sms_ip_whitelist` WRITE;
/*!40000 ALTER TABLE `sms_ip_whitelist` DISABLE KEYS */;
/*!40000 ALTER TABLE `sms_ip_whitelist` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sms_phone_assignments`
--

DROP TABLE IF EXISTS `sms_phone_assignments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sms_phone_assignments` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `business_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '业务名称',
  `business_code` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '业务编码',
  `merchant_no` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商户号',
  `merchant_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '商户名称',
  `phone_number` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '获取到的手机号',
  `verification_code` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '获取到的验证码',
  `fetch_count` int DEFAULT '0' COMMENT '获取验证码次数',
  `status` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT 'pending' COMMENT '状态 (pending:待取码, completed:已完成, expired:已过期, failed:失败)',
  `provider_cost` decimal(10,4) DEFAULT NULL COMMENT '渠道成本',
  `merchant_fee` decimal(10,4) DEFAULT NULL COMMENT '商户费用',
  `profit` decimal(10,4) DEFAULT NULL COMMENT '利润',
  `remark` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `customer_id` bigint DEFAULT NULL COMMENT '客户ID, 关联到sms_customers.id',
  `provider_id` bigint DEFAULT NULL COMMENT '服务商ID, 关联到sms_providers.id',
  `platform_business_type_id` bigint DEFAULT NULL COMMENT '平台业务类型ID',
  PRIMARY KEY (`id`),
  KEY `idx_customer_id` (`customer_id`),
  KEY `idx_provider_id` (`provider_id`),
  KEY `idx_merchant_no` (`merchant_no`),
  KEY `idx_business_code` (`business_code`),
  KEY `idx_phone_number` (`phone_number`),
  KEY `idx_status` (`status`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='号码记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sms_phone_assignments`
--

LOCK TABLES `sms_phone_assignments` WRITE;
/*!40000 ALTER TABLE `sms_phone_assignments` DISABLE KEYS */;
/*!40000 ALTER TABLE `sms_phone_assignments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sms_platform_business_types`
--

DROP TABLE IF EXISTS `sms_platform_business_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sms_platform_business_types` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '平台业务名称',
  `code` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '平台业务编码',
  `description` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '业务描述',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `remark` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`),
  KEY `idx_sms_platform_business_types_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='平台业务类型表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sms_platform_business_types`
--

LOCK TABLES `sms_platform_business_types` WRITE;
/*!40000 ALTER TABLE `sms_platform_business_types` DISABLE KEYS */;
INSERT INTO `sms_platform_business_types` VALUES (1,'2025-11-22 22:56:06.878','2025-11-22 22:56:06.878',NULL,'微信','wx','sdfadf',1,'asdfadf'),(2,'2025-11-22 23:49:33.409','2025-11-22 23:49:33.409',NULL,'qq','qq','1qq',0,'');
/*!40000 ALTER TABLE `sms_platform_business_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sms_platform_provider_business_mapping`
--

DROP TABLE IF EXISTS `sms_platform_provider_business_mapping`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sms_platform_provider_business_mapping` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `platform_business_type_id` bigint DEFAULT NULL COMMENT '平台业务ID（关联sms_platform_business_types表的ID）',
  `platform_business_code` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '平台业务编码',
  `provider_business_type_id` bigint DEFAULT NULL COMMENT '三方业务ID（关联sms_providers_business_types表的ID）',
  `provider_code` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '三方编码',
  `business_code` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '三方业务编码',
  `weight` int DEFAULT NULL COMMENT '权重（用于随机选择，权重越高被选中概率越大）',
  `status` tinyint(1) DEFAULT NULL COMMENT '是否启用该映射',
  `remark` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_sms_platform_provider_business_mapping_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sms_platform_provider_business_mapping`
--

LOCK TABLES `sms_platform_provider_business_mapping` WRITE;
/*!40000 ALTER TABLE `sms_platform_provider_business_mapping` DISABLE KEYS */;
INSERT INTO `sms_platform_provider_business_mapping` VALUES (1,'2025-11-22 23:04:02.834','2025-11-22 23:04:02.834',NULL,1,'wx',1,'test','qq',11,1,'1212');
/*!40000 ALTER TABLE `sms_platform_provider_business_mapping` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sms_providers`
--

DROP TABLE IF EXISTS `sms_providers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sms_providers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '三方名称',
  `code` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '三方编码',
  `api_gateway` text COLLATE utf8mb4_bin COMMENT '三方API网关地址',
  `merchant_id` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '三方商户号',
  `merchant_key` text COLLATE utf8mb4_bin COMMENT '三方商户key',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `remark` varchar(500) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `api_config` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '服务商的API配置 (如URL, key等)',
  `is_enabled` tinyint(1) DEFAULT NULL COMMENT '是否启用该服务商',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sms_providers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='第三方服务商表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sms_providers`
--

LOCK TABLES `sms_providers` WRITE;
/*!40000 ALTER TABLE `sms_providers` DISABLE KEYS */;
INSERT INTO `sms_providers` VALUES (1,'测试','test','https://www.google.com','merch_no','12121212',1,'imya ',NULL,NULL,'2025-11-22 22:42:36.184','2025-11-22 22:42:36.184',NULL),(2,'test2','test','http://www.baidu.com','1121212','123456121212',1,'12121',NULL,NULL,'2025-11-22 22:55:27.876','2025-11-22 22:55:27.876',NULL);
/*!40000 ALTER TABLE `sms_providers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sms_providers_business_types`
--

DROP TABLE IF EXISTS `sms_providers_business_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sms_providers_business_types` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `provider_id` int NOT NULL COMMENT '三方渠道ID（关联sms_providers表的ID）',
  `provider_code` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '三方编码',
  `business_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '业务名称',
  `business_code` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '业务编码',
  `price` decimal(10,4) DEFAULT '0.0000' COMMENT '该渠道该业务的价格',
  `status` tinyint(1) DEFAULT '1' COMMENT '该渠道是否支持该业务',
  `remark` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_provider_business` (`provider_id`),
  KEY `idx_provider_code` (`provider_code`),
  KEY `idx_business_code` (`business_code`),
  CONSTRAINT `fk_provider_id` FOREIGN KEY (`provider_id`) REFERENCES `sms_providers` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='三方渠道与业务关系管理表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sms_providers_business_types`
--

LOCK TABLES `sms_providers_business_types` WRITE;
/*!40000 ALTER TABLE `sms_providers_business_types` DISABLE KEYS */;
INSERT INTO `sms_providers_business_types` VALUES (1,'2025-11-22 22:54:59.221','2025-11-22 22:54:59.221',NULL,1,'test','测试业务','qq',1.0000,1,'1212'),(2,'2025-11-22 22:55:44.529','2025-11-22 22:55:44.529',NULL,2,'test','tmg ','qq',10.0000,1,'1212');
/*!40000 ALTER TABLE `sms_providers_business_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sms_transactions`
--

DROP TABLE IF EXISTS `sms_transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sms_transactions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `customer_id` bigint NOT NULL COMMENT '客户ID',
  `amount` float DEFAULT NULL COMMENT '变动金额 (正数为充值, 负数为消费)',
  `balance_before` float DEFAULT NULL COMMENT '变动前余额',
  `balance_after` float DEFAULT NULL COMMENT '变动后余额',
  `type` varchar(10) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '交易类型 (1:充值, 2:拉号码, 3:拉号-回退, 4:上分, 5:下分)',
  `reference_id` bigint DEFAULT NULL COMMENT '关联的业务ID, 例如sms_phone_assignments.id',
  `notes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_customer_id` (`customer_id`),
  KEY `idx_sms_transactions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='客户余额交易记录表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sms_transactions`
--

LOCK TABLES `sms_transactions` WRITE;
/*!40000 ALTER TABLE `sms_transactions` DISABLE KEYS */;
INSERT INTO `sms_transactions` VALUES (1,1,1,1002,1003,'1',NULL,'1212','2025-11-21 18:26:06.224','2025-11-21 18:26:06.224',NULL),(2,1,1,1003,1004,'1',NULL,NULL,'2025-11-21 23:13:08.475','2025-11-21 23:13:08.475',NULL),(3,1,-1,1004,1003,'2',NULL,NULL,'2025-11-21 23:44:31.853','2025-11-21 23:44:31.853',NULL),(4,1,-10,1003,993,'6',NULL,NULL,'2025-11-22 23:46:26.418','2025-11-22 23:46:26.418',NULL),(5,1,10,993,1003,'7',NULL,NULL,'2025-11-22 23:48:46.519','2025-11-22 23:48:46.519',NULL);
/*!40000 ALTER TABLE `sms_transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_apis`
--

DROP TABLE IF EXISTS `sys_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=197 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_apis`
--

LOCK TABLES `sys_apis` WRITE;
/*!40000 ALTER TABLE `sys_apis` DISABLE KEYS */;
INSERT INTO `sys_apis` VALUES (1,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/jwt/jsonInBlacklist','jwt加入黑名单(退出，必选)','jwt','POST'),(2,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/user/deleteUser','删除用户','系统用户','DELETE'),(3,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/user/admin_register','用户注册','系统用户','POST'),(4,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/user/getUserList','获取用户列表','系统用户','POST'),(5,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/user/setUserInfo','设置用户信息','系统用户','PUT'),(6,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/user/setSelfInfo','设置自身信息(必选)','系统用户','PUT'),(7,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/user/getUserInfo','获取自身信息(必选)','系统用户','GET'),(8,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/user/setUserAuthorities','设置权限组','系统用户','POST'),(9,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/user/changePassword','修改密码（建议选择)','系统用户','POST'),(10,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/user/setUserAuthority','修改用户角色(必选)','系统用户','POST'),(11,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/user/resetPassword','重置用户密码','系统用户','POST'),(12,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/user/setSelfSetting','用户界面配置','系统用户','PUT'),(13,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/api/createApi','创建api','api','POST'),(14,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/api/deleteApi','删除Api','api','POST'),(15,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/api/updateApi','更新Api','api','POST'),(16,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/api/getApiList','获取api列表','api','POST'),(17,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/api/getAllApis','获取所有api','api','POST'),(18,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/api/getApiById','获取api详细信息','api','POST'),(19,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/api/deleteApisByIds','批量删除api','api','DELETE'),(20,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/api/syncApi','获取待同步API','api','GET'),(21,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/api/getApiGroups','获取路由组','api','GET'),(22,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/api/enterSyncApi','确认同步API','api','POST'),(23,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/api/ignoreApi','忽略API','api','POST'),(24,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/authority/copyAuthority','拷贝角色','角色','POST'),(25,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/authority/createAuthority','创建角色','角色','POST'),(26,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/authority/deleteAuthority','删除角色','角色','POST'),(27,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/authority/updateAuthority','更新角色信息','角色','PUT'),(28,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/authority/getAuthorityList','获取角色列表','角色','POST'),(29,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/authority/setDataAuthority','设置角色资源权限','角色','POST'),(30,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/casbin/updateCasbin','更改角色api权限','casbin','POST'),(31,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/casbin/getPolicyPathByAuthorityId','获取权限列表','casbin','POST'),(32,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/menu/addBaseMenu','新增菜单','菜单','POST'),(33,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/menu/getMenu','获取菜单树(必选)','菜单','POST'),(34,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/menu/deleteBaseMenu','删除菜单','菜单','POST'),(35,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/menu/updateBaseMenu','更新菜单','菜单','POST'),(36,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/menu/getBaseMenuById','根据id获取菜单','菜单','POST'),(37,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/menu/getMenuList','分页获取基础menu列表','菜单','POST'),(38,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/menu/getBaseMenuTree','获取用户动态路由','菜单','POST'),(39,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/menu/getMenuAuthority','获取指定角色menu','菜单','POST'),(40,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/menu/addMenuAuthority','增加menu和角色关联关系','菜单','POST'),(41,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/fileUploadAndDownload/findFile','寻找目标文件（秒传）','分片上传','GET'),(42,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/fileUploadAndDownload/breakpointContinue','断点续传','分片上传','POST'),(43,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/fileUploadAndDownload/breakpointContinueFinish','断点续传完成','分片上传','POST'),(44,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/fileUploadAndDownload/removeChunk','上传完成移除文件','分片上传','POST'),(45,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/fileUploadAndDownload/upload','文件上传（建议选择）','文件上传与下载','POST'),(46,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/fileUploadAndDownload/deleteFile','删除文件','文件上传与下载','POST'),(47,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/fileUploadAndDownload/editFileName','文件名或者备注编辑','文件上传与下载','POST'),(48,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/fileUploadAndDownload/getFileList','获取上传文件列表','文件上传与下载','POST'),(49,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/fileUploadAndDownload/importURL','导入URL','文件上传与下载','POST'),(50,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/system/getServerInfo','获取服务器信息','系统服务','POST'),(51,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/system/getSystemConfig','获取配置文件内容','系统服务','POST'),(52,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/system/setSystemConfig','设置配置文件内容','系统服务','POST'),(53,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/customer/customer','更新客户','客户','PUT'),(54,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/customer/customer','创建客户','客户','POST'),(55,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/customer/customer','删除客户','客户','DELETE'),(56,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/customer/customer','获取单一客户','客户','GET'),(57,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/customer/customerList','获取客户列表','客户','GET'),(58,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/getDB','获取所有数据库','代码生成器','GET'),(59,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/getTables','获取数据库表','代码生成器','GET'),(60,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/createTemp','自动化代码','代码生成器','POST'),(61,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/preview','预览自动化代码','代码生成器','POST'),(62,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/getColumn','获取所选table的所有字段','代码生成器','GET'),(63,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/installPlugin','安装插件','代码生成器','POST'),(64,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/pubPlug','打包插件','代码生成器','POST'),(65,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/mcp','自动生成 MCP Tool 模板','代码生成器','POST'),(66,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/mcpTest','MCP Tool 测试','代码生成器','POST'),(67,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/mcpList','获取 MCP ToolList','代码生成器','POST'),(68,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/createPackage','配置模板','模板配置','POST'),(69,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/getTemplates','获取模板文件','模板配置','GET'),(70,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/getPackage','获取所有模板','模板配置','POST'),(71,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/delPackage','删除模板','模板配置','POST'),(72,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/getMeta','获取meta信息','代码生成器历史','POST'),(73,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/rollback','回滚自动生成代码','代码生成器历史','POST'),(74,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/getSysHistory','查询回滚记录','代码生成器历史','POST'),(75,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/delSysHistory','删除回滚记录','代码生成器历史','POST'),(76,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/autoCode/addFunc','增加模板方法','代码生成器历史','POST'),(77,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionaryDetail/updateSysDictionaryDetail','更新字典内容','系统字典详情','PUT'),(78,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionaryDetail/createSysDictionaryDetail','新增字典内容','系统字典详情','POST'),(79,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionaryDetail/deleteSysDictionaryDetail','删除字典内容','系统字典详情','DELETE'),(80,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionaryDetail/findSysDictionaryDetail','根据ID获取字典内容','系统字典详情','GET'),(81,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionaryDetail/getSysDictionaryDetailList','获取字典内容列表','系统字典详情','GET'),(82,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionaryDetail/getDictionaryTreeList','获取字典数列表','系统字典详情','GET'),(83,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionaryDetail/getDictionaryTreeListByType','根据分类获取字典数列表','系统字典详情','GET'),(84,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionaryDetail/getDictionaryDetailsByParent','根据父级ID获取字典详情','系统字典详情','GET'),(85,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionaryDetail/getDictionaryPath','获取字典详情的完整路径','系统字典详情','GET'),(86,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionary/createSysDictionary','新增字典','系统字典','POST'),(87,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionary/deleteSysDictionary','删除字典','系统字典','DELETE'),(88,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionary/updateSysDictionary','更新字典','系统字典','PUT'),(89,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionary/findSysDictionary','根据ID获取字典（建议选择）','系统字典','GET'),(90,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysDictionary/getSysDictionaryList','获取字典列表','系统字典','GET'),(91,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysOperationRecord/createSysOperationRecord','新增操作记录','操作记录','POST'),(92,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysOperationRecord/findSysOperationRecord','根据ID获取操作记录','操作记录','GET'),(93,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysOperationRecord/getSysOperationRecordList','获取操作记录列表','操作记录','GET'),(94,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysOperationRecord/deleteSysOperationRecord','删除操作记录','操作记录','DELETE'),(95,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysOperationRecord/deleteSysOperationRecordByIds','批量删除操作历史','操作记录','DELETE'),(96,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/simpleUploader/upload','插件版分片上传','断点续传(插件版)','POST'),(97,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/simpleUploader/checkFileMd5','文件完整度验证','断点续传(插件版)','GET'),(98,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/simpleUploader/mergeFileMd5','上传完成合并文件','断点续传(插件版)','GET'),(99,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/email/emailTest','发送测试邮件','email','POST'),(100,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/email/sendEmail','发送邮件','email','POST'),(101,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/authorityBtn/setAuthorityBtn','设置按钮权限','按钮权限','POST'),(102,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/authorityBtn/getAuthorityBtn','获取已有按钮权限','按钮权限','POST'),(103,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/authorityBtn/canRemoveAuthorityBtn','删除按钮','按钮权限','POST'),(104,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysExportTemplate/createSysExportTemplate','新增导出模板','导出模板','POST'),(105,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysExportTemplate/deleteSysExportTemplate','删除导出模板','导出模板','DELETE'),(106,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysExportTemplate/deleteSysExportTemplateByIds','批量删除导出模板','导出模板','DELETE'),(107,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysExportTemplate/updateSysExportTemplate','更新导出模板','导出模板','PUT'),(108,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysExportTemplate/findSysExportTemplate','根据ID获取导出模板','导出模板','GET'),(109,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysExportTemplate/getSysExportTemplateList','获取导出模板列表','导出模板','GET'),(110,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysExportTemplate/exportExcel','导出Excel','导出模板','GET'),(111,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysExportTemplate/exportTemplate','下载模板','导出模板','GET'),(112,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysExportTemplate/importExcel','导入Excel','导出模板','POST'),(113,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/info/createInfo','新建公告','公告','POST'),(114,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/info/deleteInfo','删除公告','公告','DELETE'),(115,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/info/deleteInfoByIds','批量删除公告','公告','DELETE'),(116,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/info/updateInfo','更新公告','公告','PUT'),(117,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/info/findInfo','根据ID获取公告','公告','GET'),(118,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/info/getInfoList','获取公告列表','公告','GET'),(119,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysParams/createSysParams','新建参数','参数管理','POST'),(120,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysParams/deleteSysParams','删除参数','参数管理','DELETE'),(121,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysParams/deleteSysParamsByIds','批量删除参数','参数管理','DELETE'),(122,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysParams/updateSysParams','更新参数','参数管理','PUT'),(123,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysParams/findSysParams','根据ID获取参数','参数管理','GET'),(124,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysParams/getSysParamsList','获取参数列表','参数管理','GET'),(125,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysParams/getSysParam','获取参数列表','参数管理','GET'),(126,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/attachmentCategory/getCategoryList','分类列表','媒体库分类','GET'),(127,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/attachmentCategory/addCategory','添加/编辑分类','媒体库分类','POST'),(128,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/attachmentCategory/deleteCategory','删除分类','媒体库分类','POST'),(129,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysVersion/findSysVersion','获取单一版本','版本控制','GET'),(130,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysVersion/getSysVersionList','获取版本列表','版本控制','GET'),(131,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysVersion/downloadVersionJson','下载版本json','版本控制','GET'),(132,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysVersion/exportVersion','创建版本','版本控制','POST'),(133,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysVersion/importVersion','同步版本','版本控制','POST'),(134,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysVersion/deleteSysVersion','删除版本','版本控制','DELETE'),(135,'2025-11-21 17:18:35.985','2025-11-21 17:18:35.985',NULL,'/sysVersion/deleteSysVersionByIds','批量删除版本','版本控制','DELETE'),(136,'2025-11-21 17:28:07.207','2025-11-21 17:28:07.207',NULL,'/smsApiLogs/createSmsApiLogs','新增访问日志','访问日志','POST'),(137,'2025-11-21 17:28:07.208','2025-11-21 17:28:07.208',NULL,'/smsApiLogs/deleteSmsApiLogs','删除访问日志','访问日志','DELETE'),(138,'2025-11-21 17:28:07.210','2025-11-21 17:28:07.210',NULL,'/smsApiLogs/deleteSmsApiLogsByIds','批量删除访问日志','访问日志','DELETE'),(139,'2025-11-21 17:28:07.211','2025-11-21 17:28:07.211',NULL,'/smsApiLogs/updateSmsApiLogs','更新访问日志','访问日志','PUT'),(140,'2025-11-21 17:28:07.212','2025-11-21 17:28:07.212',NULL,'/smsApiLogs/findSmsApiLogs','根据ID获取访问日志','访问日志','GET'),(141,'2025-11-21 17:28:07.212','2025-11-21 17:28:07.212',NULL,'/smsApiLogs/getSmsApiLogsList','获取访问日志列表','访问日志','GET'),(148,'2025-11-21 17:30:54.231','2025-11-21 17:30:54.231',NULL,'/smsCustomers/createSmsCustomers','新增商户','商户','POST'),(149,'2025-11-21 17:30:54.232','2025-11-21 17:30:54.232',NULL,'/smsCustomers/deleteSmsCustomers','删除商户','商户','DELETE'),(150,'2025-11-21 17:30:54.233','2025-11-21 17:30:54.233',NULL,'/smsCustomers/deleteSmsCustomersByIds','批量删除商户','商户','DELETE'),(151,'2025-11-21 17:30:54.234','2025-11-21 17:30:54.234',NULL,'/smsCustomers/updateSmsCustomers','更新商户','商户','PUT'),(152,'2025-11-21 17:30:54.234','2025-11-21 17:30:54.234',NULL,'/smsCustomers/findSmsCustomers','根据ID获取商户','商户','GET'),(153,'2025-11-21 17:30:54.235','2025-11-21 17:30:54.235',NULL,'/smsCustomers/getSmsCustomersList','获取商户列表','商户','GET'),(154,'2025-11-21 17:31:47.479','2025-11-21 17:31:47.479',NULL,'/smsIpWhitelist/createSmsIpWhitelist','新增白名单','白名单','POST'),(155,'2025-11-21 17:31:47.481','2025-11-21 17:31:47.481',NULL,'/smsIpWhitelist/deleteSmsIpWhitelist','删除白名单','白名单','DELETE'),(156,'2025-11-21 17:31:47.482','2025-11-21 17:31:47.482',NULL,'/smsIpWhitelist/deleteSmsIpWhitelistByIds','批量删除白名单','白名单','DELETE'),(157,'2025-11-21 17:31:47.483','2025-11-21 17:31:47.483',NULL,'/smsIpWhitelist/updateSmsIpWhitelist','更新白名单','白名单','PUT'),(158,'2025-11-21 17:31:47.483','2025-11-21 17:31:47.483',NULL,'/smsIpWhitelist/findSmsIpWhitelist','根据ID获取白名单','白名单','GET'),(159,'2025-11-21 17:31:47.484','2025-11-21 17:31:47.484',NULL,'/smsIpWhitelist/getSmsIpWhitelistList','获取白名单列表','白名单','GET'),(160,'2025-11-21 17:33:01.910','2025-11-21 17:33:01.910',NULL,'/smsPhoneAssignments/createSmsPhoneAssignments','新增号码记录','号码记录','POST'),(161,'2025-11-21 17:33:01.911','2025-11-21 17:33:01.911',NULL,'/smsPhoneAssignments/deleteSmsPhoneAssignments','删除号码记录','号码记录','DELETE'),(162,'2025-11-21 17:33:01.912','2025-11-21 17:33:01.912',NULL,'/smsPhoneAssignments/deleteSmsPhoneAssignmentsByIds','批量删除号码记录','号码记录','DELETE'),(163,'2025-11-21 17:33:01.912','2025-11-21 17:33:01.912',NULL,'/smsPhoneAssignments/updateSmsPhoneAssignments','更新号码记录','号码记录','PUT'),(164,'2025-11-21 17:33:01.913','2025-11-21 17:33:01.913',NULL,'/smsPhoneAssignments/findSmsPhoneAssignments','根据ID获取号码记录','号码记录','GET'),(165,'2025-11-21 17:33:01.913','2025-11-21 17:33:01.913',NULL,'/smsPhoneAssignments/getSmsPhoneAssignmentsList','获取号码记录列表','号码记录','GET'),(166,'2025-11-21 17:35:33.939','2025-11-21 17:35:33.939',NULL,'/smsProviders/createSmsProviders','新增服务端','服务端','POST'),(167,'2025-11-21 17:35:33.941','2025-11-21 17:35:33.941',NULL,'/smsProviders/deleteSmsProviders','删除服务端','服务端','DELETE'),(168,'2025-11-21 17:35:33.941','2025-11-21 17:35:33.941',NULL,'/smsProviders/deleteSmsProvidersByIds','批量删除服务端','服务端','DELETE'),(169,'2025-11-21 17:35:33.942','2025-11-21 17:35:33.942',NULL,'/smsProviders/updateSmsProviders','更新服务端','服务端','PUT'),(170,'2025-11-21 17:35:33.943','2025-11-21 17:35:33.943',NULL,'/smsProviders/findSmsProviders','根据ID获取服务端','服务端','GET'),(171,'2025-11-21 17:35:33.944','2025-11-21 17:35:33.944',NULL,'/smsProviders/getSmsProvidersList','获取服务端列表','服务端','GET'),(172,'2025-11-21 17:36:33.849','2025-11-21 17:36:33.849',NULL,'/smsTransactions/createSmsTransactions','新增交易记录','交易记录','POST'),(173,'2025-11-21 17:36:33.850','2025-11-21 17:36:33.850',NULL,'/smsTransactions/deleteSmsTransactions','删除交易记录','交易记录','DELETE'),(174,'2025-11-21 17:36:33.851','2025-11-21 17:36:33.851',NULL,'/smsTransactions/deleteSmsTransactionsByIds','批量删除交易记录','交易记录','DELETE'),(175,'2025-11-21 17:36:33.852','2025-11-21 17:36:33.852',NULL,'/smsTransactions/updateSmsTransactions','更新交易记录','交易记录','PUT'),(176,'2025-11-21 17:36:33.853','2025-11-21 17:36:33.853',NULL,'/smsTransactions/findSmsTransactions','根据ID获取交易记录','交易记录','GET'),(177,'2025-11-21 17:36:33.853','2025-11-21 17:36:33.853',NULL,'/smsTransactions/getSmsTransactionsList','获取交易记录列表','交易记录','GET'),(178,'2025-11-21 18:25:47.888','2025-11-21 18:25:47.888',NULL,'/smsCustomers/creditDebit','上下分','商户','POST'),(179,'2025-11-22 00:10:43.882','2025-11-22 00:10:43.882',NULL,'/smsProvidersBusinessTypes/createSmsProvidersBusinessTypes','新增三方业务','三方业务','POST'),(180,'2025-11-22 00:10:43.884','2025-11-22 00:10:43.884',NULL,'/smsProvidersBusinessTypes/deleteSmsProvidersBusinessTypes','删除三方业务','三方业务','DELETE'),(181,'2025-11-22 00:10:43.886','2025-11-22 00:10:43.886',NULL,'/smsProvidersBusinessTypes/deleteSmsProvidersBusinessTypesByIds','批量删除三方业务','三方业务','DELETE'),(182,'2025-11-22 00:10:43.887','2025-11-22 00:10:43.887',NULL,'/smsProvidersBusinessTypes/updateSmsProvidersBusinessTypes','更新三方业务','三方业务','PUT'),(183,'2025-11-22 00:10:43.889','2025-11-22 00:10:43.889',NULL,'/smsProvidersBusinessTypes/findSmsProvidersBusinessTypes','根据ID获取三方业务','三方业务','GET'),(184,'2025-11-22 00:10:43.890','2025-11-22 00:10:43.890',NULL,'/smsProvidersBusinessTypes/getSmsProvidersBusinessTypesList','获取三方业务列表','三方业务','GET'),(185,'2025-11-22 00:29:33.707','2025-11-22 00:29:33.707',NULL,'/smsPlatformBusinessTypes/createSmsPlatformBusinessTypes','新增平台业务','平台业务','POST'),(186,'2025-11-22 00:29:33.708','2025-11-22 00:29:33.708',NULL,'/smsPlatformBusinessTypes/deleteSmsPlatformBusinessTypes','删除平台业务','平台业务','DELETE'),(187,'2025-11-22 00:29:33.709','2025-11-22 00:29:33.709',NULL,'/smsPlatformBusinessTypes/deleteSmsPlatformBusinessTypesByIds','批量删除平台业务','平台业务','DELETE'),(188,'2025-11-22 00:29:33.711','2025-11-22 00:29:33.711',NULL,'/smsPlatformBusinessTypes/updateSmsPlatformBusinessTypes','更新平台业务','平台业务','PUT'),(189,'2025-11-22 00:29:33.712','2025-11-22 00:29:33.712',NULL,'/smsPlatformBusinessTypes/findSmsPlatformBusinessTypes','根据ID获取平台业务','平台业务','GET'),(190,'2025-11-22 00:29:33.713','2025-11-22 00:29:33.713',NULL,'/smsPlatformBusinessTypes/getSmsPlatformBusinessTypesList','获取平台业务列表','平台业务','GET'),(191,'2025-11-22 00:30:37.805','2025-11-22 00:30:37.805',NULL,'/smsPlatformProviderBusinessMapping/createSmsPlatformProviderBusinessMapping','新增平台子业务','平台子业务','POST'),(192,'2025-11-22 00:30:37.807','2025-11-22 00:30:37.807',NULL,'/smsPlatformProviderBusinessMapping/deleteSmsPlatformProviderBusinessMapping','删除平台子业务','平台子业务','DELETE'),(193,'2025-11-22 00:30:37.808','2025-11-22 00:30:37.808',NULL,'/smsPlatformProviderBusinessMapping/deleteSmsPlatformProviderBusinessMappingByIds','批量删除平台子业务','平台子业务','DELETE'),(194,'2025-11-22 00:30:37.809','2025-11-22 00:30:37.809',NULL,'/smsPlatformProviderBusinessMapping/updateSmsPlatformProviderBusinessMapping','更新平台子业务','平台子业务','PUT'),(195,'2025-11-22 00:30:37.810','2025-11-22 00:30:37.810',NULL,'/smsPlatformProviderBusinessMapping/findSmsPlatformProviderBusinessMapping','根据ID获取平台子业务','平台子业务','GET'),(196,'2025-11-22 00:30:37.812','2025-11-22 00:30:37.812',NULL,'/smsPlatformProviderBusinessMapping/getSmsPlatformProviderBusinessMappingList','获取平台子业务列表','平台子业务','GET');
/*!40000 ALTER TABLE `sys_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authorities`
--

DROP TABLE IF EXISTS `sys_authorities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `uni_sys_authorities_authority_id` (`authority_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9529 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authorities`
--

LOCK TABLES `sys_authorities` WRITE;
/*!40000 ALTER TABLE `sys_authorities` DISABLE KEYS */;
INSERT INTO `sys_authorities` VALUES ('2025-11-21 17:18:35.992','2025-11-22 00:36:04.133',NULL,888,'普通用户',0,'dashboard'),('2025-11-21 17:18:35.992','2025-11-21 17:18:36.118',NULL,8881,'普通用户子角色',888,'dashboard'),('2025-11-21 17:18:35.992','2025-11-21 17:18:36.117',NULL,9528,'测试角色',0,'dashboard');
/*!40000 ALTER TABLE `sys_authorities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authority_btns`
--

DROP TABLE IF EXISTS `sys_authority_btns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authority_btns` (
  `authority_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authority_btns`
--

LOCK TABLES `sys_authority_btns` WRITE;
/*!40000 ALTER TABLE `sys_authority_btns` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_authority_btns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_authority_menus`
--

DROP TABLE IF EXISTS `sys_authority_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_authority_menus`
--

LOCK TABLES `sys_authority_menus` WRITE;
/*!40000 ALTER TABLE `sys_authority_menus` DISABLE KEYS */;
INSERT INTO `sys_authority_menus` VALUES (1,888),(1,8881),(1,9528),(2,888),(2,8881),(2,9528),(3,888),(3,8881),(4,888),(4,8881),(4,9528),(5,888),(5,8881),(6,888),(6,8881),(7,888),(7,8881),(8,888),(8,8881),(8,9528),(9,888),(9,8881),(10,888),(11,888),(12,888),(13,888),(14,888),(15,888),(16,888),(17,888),(17,8881),(18,888),(18,8881),(19,888),(19,8881),(20,888),(20,8881),(21,888),(21,8881),(22,888),(22,8881),(23,888),(23,8881),(24,888),(24,8881),(25,888),(25,8881),(26,888),(26,8881),(27,888),(27,8881),(28,888),(28,8881),(29,888),(29,8881),(30,888),(30,8881),(31,888),(32,888),(33,888),(34,888),(35,888),(36,888),(37,888),(38,888),(39,888),(40,888),(41,888),(42,888),(43,888),(44,888),(45,888);
/*!40000 ALTER TABLE `sys_authority_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_auto_code_histories`
--

DROP TABLE IF EXISTS `sys_auto_code_histories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_auto_code_histories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '表名',
  `package` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '模块名/插件名',
  `request` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '前端传入的结构化信息',
  `struct_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '结构体名称',
  `abbreviation` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '结构体名称缩写',
  `business_db` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '业务库',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'Struct中文名称',
  `templates` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '模板信息',
  `Injections` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '注入路径',
  `flag` bigint DEFAULT NULL COMMENT '[0:创建,1:回滚]',
  `api_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'api表注册内容',
  `menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `export_template_id` bigint unsigned DEFAULT NULL COMMENT '导出模板ID',
  `package_id` bigint unsigned DEFAULT NULL COMMENT '包ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_auto_code_histories`
--

LOCK TABLES `sys_auto_code_histories` WRITE;
/*!40000 ALTER TABLE `sys_auto_code_histories` DISABLE KEYS */;
INSERT INTO `sys_auto_code_histories` VALUES (1,'2025-11-21 17:28:07.232','2025-11-21 17:28:07.232',NULL,'sms_api_logs','sms','{\"package\":\"sms\",\"tableName\":\"sms_api_logs\",\"businessDB\":\"\",\"structName\":\"SmsApiLogs\",\"packageName\":\"smsApiLogs\",\"description\":\"访问日志\",\"abbreviation\":\"smsApiLogs\",\"humpPackageName\":\"sms_api_logs\",\"gvaModel\":true,\"autoMigrate\":true,\"autoCreateResource\":false,\"autoCreateApiToSql\":true,\"autoCreateMenuToSql\":true,\"autoCreateBtnAuth\":false,\"onlyTemplate\":false,\"isTree\":false,\"treeJson\":\"\",\"isAdd\":false,\"fields\":[{\"fieldName\":\"CustomerId\",\"fieldDesc\":\"客户ID\",\"fieldType\":\"int\",\"fieldJson\":\"customerId\",\"dataTypeLong\":\"19\",\"comment\":\"客户ID\",\"columnName\":\"customer_id\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"RequestIp\",\"fieldDesc\":\"请求来源IP\",\"fieldType\":\"string\",\"fieldJson\":\"requestIp\",\"dataTypeLong\":\"45\",\"comment\":\"请求来源IP\",\"columnName\":\"request_ip\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"RequestPath\",\"fieldDesc\":\"请求的API路径\",\"fieldType\":\"string\",\"fieldJson\":\"requestPath\",\"dataTypeLong\":\"255\",\"comment\":\"请求的API路径\",\"columnName\":\"request_path\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"RequestBody\",\"fieldDesc\":\"请求体内容\",\"fieldType\":\"string\",\"fieldJson\":\"requestBody\",\"dataTypeLong\":\"\",\"comment\":\"请求体内容\",\"columnName\":\"request_body\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"ResponseCode\",\"fieldDesc\":\"HTTP响应状态码\",\"fieldType\":\"int\",\"fieldJson\":\"responseCode\",\"dataTypeLong\":\"10\",\"comment\":\"HTTP响应状态码\",\"columnName\":\"response_code\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"DurationMs\",\"fieldDesc\":\"请求处理耗时(毫秒)\",\"fieldType\":\"int\",\"fieldJson\":\"durationMs\",\"dataTypeLong\":\"10\",\"comment\":\"请求处理耗时(毫秒)\",\"columnName\":\"duration_ms\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"}],\"generateWeb\":true,\"generateServer\":true,\"primaryField\":{\"fieldName\":\"ID\",\"fieldDesc\":\"ID\",\"fieldType\":\"uint\",\"fieldJson\":\"ID\",\"dataTypeLong\":\"20\",\"comment\":\"主键ID\",\"columnName\":\"id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":false,\"table\":false,\"desc\":false,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":false,\"sort\":false,\"primaryKey\":false,\"dataSource\":null,\"checkDataSource\":false,\"fieldIndexType\":\"\"}}','SmsApiLogs','SmsApiLogs','','访问日志','{\"resource/plugin/server/api/api.go.tpl\":\"plugin/sms/api/sms_api_logs.go\",\"resource/plugin/server/model/model.go.tpl\":\"plugin/sms/model/sms_api_logs.go\",\"resource/plugin/server/model/request/request.go.tpl\":\"plugin/sms/model/request/sms_api_logs.go\",\"resource/plugin/server/router/router.go.tpl\":\"plugin/sms/router/sms_api_logs.go\",\"resource/plugin/server/service/service.go.tpl\":\"plugin/sms/service/sms_api_logs.go\",\"resource/plugin/web/api/api.js.tpl\":\"plugin/sms/api/smsApiLogs.js\",\"resource/plugin/web/form/form.vue.tpl\":\"plugin/sms/form/smsApiLogs.vue\",\"resource/plugin/web/view/view.vue.tpl\":\"plugin/sms/view/smsApiLogs.vue\"}','{\"PluginApiEnter\":\"{\\\"Type\\\":\\\"PluginApiEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/api/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/api/enter.go\\\",\\\"StructName\\\":\\\"SmsApiLogs\\\",\\\"StructCamelName\\\":\\\"smsApiLogs\\\",\\\"ModuleName\\\":\\\"serviceSmsApiLogs\\\",\\\"GroupName\\\":\\\"Service\\\",\\\"PackageName\\\":\\\"service\\\",\\\"ServiceName\\\":\\\"SmsApiLogs\\\"}\",\"PluginGen\":\"{\\\"Type\\\":\\\"PluginGen\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/gen/gen.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/gen/gen.go\\\",\\\"StructName\\\":\\\"SmsApiLogs\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeGorm\":\"{\\\"Type\\\":\\\"PluginInitializeGorm\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/gorm.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/gorm.go\\\",\\\"StructName\\\":\\\"SmsApiLogs\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeRouter\":\"{\\\"Type\\\":\\\"PluginInitializeRouter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/router.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/router\\\\\\\"\\\",\\\"ImportGlobalPath\\\":\\\"\\\",\\\"ImportMiddlewarePath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/router.go\\\",\\\"AppName\\\":\\\"Router\\\",\\\"GroupName\\\":\\\"SmsApiLogs\\\",\\\"PackageName\\\":\\\"router\\\",\\\"FunctionName\\\":\\\"Init\\\",\\\"LeftRouterGroupName\\\":\\\"public\\\",\\\"RightRouterGroupName\\\":\\\"private\\\"}\",\"PluginRouterEnter\":\"{\\\"Type\\\":\\\"PluginRouterEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/router/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/router/enter.go\\\",\\\"StructName\\\":\\\"SmsApiLogs\\\",\\\"StructCamelName\\\":\\\"smsApiLogs\\\",\\\"ModuleName\\\":\\\"apiSmsApiLogs\\\",\\\"GroupName\\\":\\\"Api\\\",\\\"PackageName\\\":\\\"api\\\",\\\"ServiceName\\\":\\\"SmsApiLogs\\\"}\",\"PluginServiceEnter\":\"{\\\"Type\\\":\\\"PluginServiceEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/service/enter.go\\\",\\\"ImportPath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/service/enter.go\\\",\\\"StructName\\\":\\\"SmsApiLogs\\\",\\\"StructCamelName\\\":\\\"smsApiLogs\\\",\\\"ModuleName\\\":\\\"\\\",\\\"GroupName\\\":\\\"\\\",\\\"PackageName\\\":\\\"\\\",\\\"ServiceName\\\":\\\"\\\"}\"}',0,'[136,137,138,139,140,141]',36,0,0),(2,'2025-11-21 17:29:31.658','2025-11-21 17:29:31.658',NULL,'sms_business_types','sms','{\"package\":\"sms\",\"tableName\":\"sms_business_types\",\"businessDB\":\"\",\"structName\":\"SmsBusinessTypes\",\"packageName\":\"smsBusinessTypes\",\"description\":\"业务类型\",\"abbreviation\":\"smsBusinessTypes\",\"humpPackageName\":\"sms_business_types\",\"gvaModel\":false,\"autoMigrate\":true,\"autoCreateResource\":false,\"autoCreateApiToSql\":true,\"autoCreateMenuToSql\":true,\"autoCreateBtnAuth\":false,\"onlyTemplate\":false,\"isTree\":false,\"treeJson\":\"\",\"isAdd\":false,\"fields\":[{\"fieldName\":\"Id\",\"fieldDesc\":\"id字段\",\"fieldType\":\"int\",\"fieldJson\":\"id\",\"dataTypeLong\":\"10\",\"comment\":\"\",\"columnName\":\"id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":true,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Name\",\"fieldDesc\":\"业务名称, 例如 \\\"腾讯QQ\\\"\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"255\",\"comment\":\"业务名称, 例如 \\\"腾讯QQ\\\"\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Code\",\"fieldDesc\":\"业务代码, 例如 \\\"qq\\\"\",\"fieldType\":\"string\",\"fieldJson\":\"code\",\"dataTypeLong\":\"50\",\"comment\":\"业务代码, 例如 \\\"qq\\\"\",\"columnName\":\"code\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"IsEnabled\",\"fieldDesc\":\"是否开放该业务\",\"fieldType\":\"bool\",\"fieldJson\":\"isEnabled\",\"dataTypeLong\":\"\",\"comment\":\"是否开放该业务\",\"columnName\":\"is_enabled\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"}],\"generateWeb\":true,\"generateServer\":true,\"primaryField\":{\"fieldName\":\"Id\",\"fieldDesc\":\"id字段\",\"fieldType\":\"int\",\"fieldJson\":\"id\",\"dataTypeLong\":\"10\",\"comment\":\"\",\"columnName\":\"id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":true,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"}}','SmsBusinessTypes','SmsBusinessTypes','','业务类型','{\"resource/plugin/server/api/api.go.tpl\":\"plugin/sms/api/sms_business_types.go\",\"resource/plugin/server/model/model.go.tpl\":\"plugin/sms/model/sms_business_types.go\",\"resource/plugin/server/model/request/request.go.tpl\":\"plugin/sms/model/request/sms_business_types.go\",\"resource/plugin/server/router/router.go.tpl\":\"plugin/sms/router/sms_business_types.go\",\"resource/plugin/server/service/service.go.tpl\":\"plugin/sms/service/sms_business_types.go\",\"resource/plugin/web/api/api.js.tpl\":\"plugin/sms/api/smsBusinessTypes.js\",\"resource/plugin/web/form/form.vue.tpl\":\"plugin/sms/form/smsBusinessTypes.vue\",\"resource/plugin/web/view/view.vue.tpl\":\"plugin/sms/view/smsBusinessTypes.vue\"}','{\"PluginApiEnter\":\"{\\\"Type\\\":\\\"PluginApiEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/api/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/api/enter.go\\\",\\\"StructName\\\":\\\"SmsBusinessTypes\\\",\\\"StructCamelName\\\":\\\"smsBusinessTypes\\\",\\\"ModuleName\\\":\\\"serviceSmsBusinessTypes\\\",\\\"GroupName\\\":\\\"Service\\\",\\\"PackageName\\\":\\\"service\\\",\\\"ServiceName\\\":\\\"SmsBusinessTypes\\\"}\",\"PluginGen\":\"{\\\"Type\\\":\\\"PluginGen\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/gen/gen.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/gen/gen.go\\\",\\\"StructName\\\":\\\"SmsBusinessTypes\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeGorm\":\"{\\\"Type\\\":\\\"PluginInitializeGorm\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/gorm.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/gorm.go\\\",\\\"StructName\\\":\\\"SmsBusinessTypes\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeRouter\":\"{\\\"Type\\\":\\\"PluginInitializeRouter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/router.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/router\\\\\\\"\\\",\\\"ImportGlobalPath\\\":\\\"\\\",\\\"ImportMiddlewarePath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/router.go\\\",\\\"AppName\\\":\\\"Router\\\",\\\"GroupName\\\":\\\"SmsBusinessTypes\\\",\\\"PackageName\\\":\\\"router\\\",\\\"FunctionName\\\":\\\"Init\\\",\\\"LeftRouterGroupName\\\":\\\"public\\\",\\\"RightRouterGroupName\\\":\\\"private\\\"}\",\"PluginRouterEnter\":\"{\\\"Type\\\":\\\"PluginRouterEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/router/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/router/enter.go\\\",\\\"StructName\\\":\\\"SmsBusinessTypes\\\",\\\"StructCamelName\\\":\\\"smsBusinessTypes\\\",\\\"ModuleName\\\":\\\"apiSmsBusinessTypes\\\",\\\"GroupName\\\":\\\"Api\\\",\\\"PackageName\\\":\\\"api\\\",\\\"ServiceName\\\":\\\"SmsBusinessTypes\\\"}\",\"PluginServiceEnter\":\"{\\\"Type\\\":\\\"PluginServiceEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/service/enter.go\\\",\\\"ImportPath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/service/enter.go\\\",\\\"StructName\\\":\\\"SmsBusinessTypes\\\",\\\"StructCamelName\\\":\\\"smsBusinessTypes\\\",\\\"ModuleName\\\":\\\"\\\",\\\"GroupName\\\":\\\"\\\",\\\"PackageName\\\":\\\"\\\",\\\"ServiceName\\\":\\\"\\\"}\"}',0,'[142,143,144,145,146,147]',37,0,0),(3,'2025-11-21 17:30:54.259','2025-11-21 17:30:54.259',NULL,'sms_customers','sms','{\"package\":\"sms\",\"tableName\":\"sms_customers\",\"businessDB\":\"\",\"structName\":\"SmsCustomers\",\"packageName\":\"smsCustomers\",\"description\":\"商户\",\"abbreviation\":\"smsCustomers\",\"humpPackageName\":\"sms_customers\",\"gvaModel\":true,\"autoMigrate\":true,\"autoCreateResource\":false,\"autoCreateApiToSql\":true,\"autoCreateMenuToSql\":true,\"autoCreateBtnAuth\":false,\"onlyTemplate\":false,\"isTree\":false,\"treeJson\":\"\",\"isAdd\":false,\"fields\":[{\"fieldName\":\"Username\",\"fieldDesc\":\"客户端登录用户名\",\"fieldType\":\"string\",\"fieldJson\":\"username\",\"dataTypeLong\":\"255\",\"comment\":\"客户端登录用户名\",\"columnName\":\"username\",\"fieldSearchType\":\"LIKE\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Email\",\"fieldDesc\":\"客户端登录邮箱\",\"fieldType\":\"string\",\"fieldJson\":\"email\",\"dataTypeLong\":\"255\",\"comment\":\"客户端登录邮箱\",\"columnName\":\"email\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"PasswordHash\",\"fieldDesc\":\"客户端登录用的密码哈希\",\"fieldType\":\"string\",\"fieldJson\":\"passwordHash\",\"dataTypeLong\":\"255\",\"comment\":\"客户端登录用的密码哈希\",\"columnName\":\"password_hash\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"ApiSecretKey\",\"fieldDesc\":\"用于生成API Token的唯一密钥\",\"fieldType\":\"string\",\"fieldJson\":\"apiSecretKey\",\"dataTypeLong\":\"255\",\"comment\":\"用于生成API Token的唯一密钥\",\"columnName\":\"api_secret_key\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Balance\",\"fieldDesc\":\"客户余额\",\"fieldType\":\"float64\",\"fieldJson\":\"balance\",\"dataTypeLong\":\"10\",\"comment\":\"客户余额\",\"columnName\":\"balance\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Status\",\"fieldDesc\":\"客户状态 (1:正常, 2:冻结, 0:已删除)\",\"fieldType\":\"bool\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"客户状态 (1:正常, 2:冻结, 0:已删除)\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"RegistrationIp\",\"fieldDesc\":\"注册时的IP地址\",\"fieldType\":\"string\",\"fieldJson\":\"registrationIp\",\"dataTypeLong\":\"45\",\"comment\":\"注册时的IP地址\",\"columnName\":\"registration_ip\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"LastLoginIp\",\"fieldDesc\":\"最后一次登录的IP地址\",\"fieldType\":\"string\",\"fieldJson\":\"lastLoginIp\",\"dataTypeLong\":\"45\",\"comment\":\"最后一次登录的IP地址\",\"columnName\":\"last_login_ip\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"LastLoginAt\",\"fieldDesc\":\"最后一次登录的时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"lastLoginAt\",\"dataTypeLong\":\"\",\"comment\":\"最后一次登录的时间\",\"columnName\":\"last_login_at\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"}],\"generateWeb\":true,\"generateServer\":true,\"primaryField\":{\"fieldName\":\"ID\",\"fieldDesc\":\"ID\",\"fieldType\":\"uint\",\"fieldJson\":\"ID\",\"dataTypeLong\":\"20\",\"comment\":\"主键ID\",\"columnName\":\"id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":false,\"table\":false,\"desc\":false,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":false,\"sort\":false,\"primaryKey\":false,\"dataSource\":null,\"checkDataSource\":false,\"fieldIndexType\":\"\"}}','SmsCustomers','SmsCustomers','','商户','{\"resource/plugin/server/api/api.go.tpl\":\"plugin/sms/api/sms_customers.go\",\"resource/plugin/server/model/model.go.tpl\":\"plugin/sms/model/sms_customers.go\",\"resource/plugin/server/model/request/request.go.tpl\":\"plugin/sms/model/request/sms_customers.go\",\"resource/plugin/server/router/router.go.tpl\":\"plugin/sms/router/sms_customers.go\",\"resource/plugin/server/service/service.go.tpl\":\"plugin/sms/service/sms_customers.go\",\"resource/plugin/web/api/api.js.tpl\":\"plugin/sms/api/smsCustomers.js\",\"resource/plugin/web/form/form.vue.tpl\":\"plugin/sms/form/smsCustomers.vue\",\"resource/plugin/web/view/view.vue.tpl\":\"plugin/sms/view/smsCustomers.vue\"}','{\"PluginApiEnter\":\"{\\\"Type\\\":\\\"PluginApiEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/api/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/api/enter.go\\\",\\\"StructName\\\":\\\"SmsCustomers\\\",\\\"StructCamelName\\\":\\\"smsCustomers\\\",\\\"ModuleName\\\":\\\"serviceSmsCustomers\\\",\\\"GroupName\\\":\\\"Service\\\",\\\"PackageName\\\":\\\"service\\\",\\\"ServiceName\\\":\\\"SmsCustomers\\\"}\",\"PluginGen\":\"{\\\"Type\\\":\\\"PluginGen\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/gen/gen.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/gen/gen.go\\\",\\\"StructName\\\":\\\"SmsCustomers\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeGorm\":\"{\\\"Type\\\":\\\"PluginInitializeGorm\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/gorm.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/gorm.go\\\",\\\"StructName\\\":\\\"SmsCustomers\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeRouter\":\"{\\\"Type\\\":\\\"PluginInitializeRouter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/router.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/router\\\\\\\"\\\",\\\"ImportGlobalPath\\\":\\\"\\\",\\\"ImportMiddlewarePath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/router.go\\\",\\\"AppName\\\":\\\"Router\\\",\\\"GroupName\\\":\\\"SmsCustomers\\\",\\\"PackageName\\\":\\\"router\\\",\\\"FunctionName\\\":\\\"Init\\\",\\\"LeftRouterGroupName\\\":\\\"public\\\",\\\"RightRouterGroupName\\\":\\\"private\\\"}\",\"PluginRouterEnter\":\"{\\\"Type\\\":\\\"PluginRouterEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/router/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/router/enter.go\\\",\\\"StructName\\\":\\\"SmsCustomers\\\",\\\"StructCamelName\\\":\\\"smsCustomers\\\",\\\"ModuleName\\\":\\\"apiSmsCustomers\\\",\\\"GroupName\\\":\\\"Api\\\",\\\"PackageName\\\":\\\"api\\\",\\\"ServiceName\\\":\\\"SmsCustomers\\\"}\",\"PluginServiceEnter\":\"{\\\"Type\\\":\\\"PluginServiceEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/service/enter.go\\\",\\\"ImportPath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/service/enter.go\\\",\\\"StructName\\\":\\\"SmsCustomers\\\",\\\"StructCamelName\\\":\\\"smsCustomers\\\",\\\"ModuleName\\\":\\\"\\\",\\\"GroupName\\\":\\\"\\\",\\\"PackageName\\\":\\\"\\\",\\\"ServiceName\\\":\\\"\\\"}\"}',0,'[148,149,150,151,152,153]',38,0,0),(4,'2025-11-21 17:31:47.508','2025-11-21 17:31:47.508',NULL,'sms_ip_whitelist','sms','{\"package\":\"sms\",\"tableName\":\"sms_ip_whitelist\",\"businessDB\":\"\",\"structName\":\"SmsIpWhitelist\",\"packageName\":\"smsIpWhitelist\",\"description\":\"白名单\",\"abbreviation\":\"smsIpWhitelist\",\"humpPackageName\":\"sms_ip_whitelist\",\"gvaModel\":true,\"autoMigrate\":true,\"autoCreateResource\":false,\"autoCreateApiToSql\":true,\"autoCreateMenuToSql\":true,\"autoCreateBtnAuth\":false,\"onlyTemplate\":false,\"isTree\":false,\"treeJson\":\"\",\"isAdd\":false,\"fields\":[{\"fieldName\":\"CustomerId\",\"fieldDesc\":\"客户ID\",\"fieldType\":\"int\",\"fieldJson\":\"customerId\",\"dataTypeLong\":\"19\",\"comment\":\"客户ID\",\"columnName\":\"customer_id\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"IpAddress\",\"fieldDesc\":\"白名单IP或IP段\",\"fieldType\":\"string\",\"fieldJson\":\"ipAddress\",\"dataTypeLong\":\"45\",\"comment\":\"白名单IP或IP段\",\"columnName\":\"ip_address\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Notes\",\"fieldDesc\":\"备注, 例如 \\\"办公室IP\\\"\",\"fieldType\":\"string\",\"fieldJson\":\"notes\",\"dataTypeLong\":\"255\",\"comment\":\"备注, 例如 \\\"办公室IP\\\"\",\"columnName\":\"notes\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"}],\"generateWeb\":true,\"generateServer\":true,\"primaryField\":{\"fieldName\":\"ID\",\"fieldDesc\":\"ID\",\"fieldType\":\"uint\",\"fieldJson\":\"ID\",\"dataTypeLong\":\"20\",\"comment\":\"主键ID\",\"columnName\":\"id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":false,\"table\":false,\"desc\":false,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":false,\"sort\":false,\"primaryKey\":false,\"dataSource\":null,\"checkDataSource\":false,\"fieldIndexType\":\"\"}}','SmsIpWhitelist','SmsIpWhitelist','','白名单','{\"resource/plugin/server/api/api.go.tpl\":\"plugin/sms/api/sms_ip_whitelist.go\",\"resource/plugin/server/model/model.go.tpl\":\"plugin/sms/model/sms_ip_whitelist.go\",\"resource/plugin/server/model/request/request.go.tpl\":\"plugin/sms/model/request/sms_ip_whitelist.go\",\"resource/plugin/server/router/router.go.tpl\":\"plugin/sms/router/sms_ip_whitelist.go\",\"resource/plugin/server/service/service.go.tpl\":\"plugin/sms/service/sms_ip_whitelist.go\",\"resource/plugin/web/api/api.js.tpl\":\"plugin/sms/api/smsIpWhitelist.js\",\"resource/plugin/web/form/form.vue.tpl\":\"plugin/sms/form/smsIpWhitelist.vue\",\"resource/plugin/web/view/view.vue.tpl\":\"plugin/sms/view/smsIpWhitelist.vue\"}','{\"PluginApiEnter\":\"{\\\"Type\\\":\\\"PluginApiEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/api/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/api/enter.go\\\",\\\"StructName\\\":\\\"SmsIpWhitelist\\\",\\\"StructCamelName\\\":\\\"smsIpWhitelist\\\",\\\"ModuleName\\\":\\\"serviceSmsIpWhitelist\\\",\\\"GroupName\\\":\\\"Service\\\",\\\"PackageName\\\":\\\"service\\\",\\\"ServiceName\\\":\\\"SmsIpWhitelist\\\"}\",\"PluginGen\":\"{\\\"Type\\\":\\\"PluginGen\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/gen/gen.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/gen/gen.go\\\",\\\"StructName\\\":\\\"SmsIpWhitelist\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeGorm\":\"{\\\"Type\\\":\\\"PluginInitializeGorm\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/gorm.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/gorm.go\\\",\\\"StructName\\\":\\\"SmsIpWhitelist\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeRouter\":\"{\\\"Type\\\":\\\"PluginInitializeRouter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/router.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/router\\\\\\\"\\\",\\\"ImportGlobalPath\\\":\\\"\\\",\\\"ImportMiddlewarePath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/router.go\\\",\\\"AppName\\\":\\\"Router\\\",\\\"GroupName\\\":\\\"SmsIpWhitelist\\\",\\\"PackageName\\\":\\\"router\\\",\\\"FunctionName\\\":\\\"Init\\\",\\\"LeftRouterGroupName\\\":\\\"public\\\",\\\"RightRouterGroupName\\\":\\\"private\\\"}\",\"PluginRouterEnter\":\"{\\\"Type\\\":\\\"PluginRouterEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/router/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/router/enter.go\\\",\\\"StructName\\\":\\\"SmsIpWhitelist\\\",\\\"StructCamelName\\\":\\\"smsIpWhitelist\\\",\\\"ModuleName\\\":\\\"apiSmsIpWhitelist\\\",\\\"GroupName\\\":\\\"Api\\\",\\\"PackageName\\\":\\\"api\\\",\\\"ServiceName\\\":\\\"SmsIpWhitelist\\\"}\",\"PluginServiceEnter\":\"{\\\"Type\\\":\\\"PluginServiceEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/service/enter.go\\\",\\\"ImportPath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/service/enter.go\\\",\\\"StructName\\\":\\\"SmsIpWhitelist\\\",\\\"StructCamelName\\\":\\\"smsIpWhitelist\\\",\\\"ModuleName\\\":\\\"\\\",\\\"GroupName\\\":\\\"\\\",\\\"PackageName\\\":\\\"\\\",\\\"ServiceName\\\":\\\"\\\"}\"}',0,'[154,155,156,157,158,159]',39,0,0),(5,'2025-11-21 17:33:01.957','2025-11-21 17:33:01.957',NULL,'sms_phone_assignments','sms','{\"package\":\"sms\",\"tableName\":\"sms_phone_assignments\",\"businessDB\":\"\",\"structName\":\"SmsPhoneAssignments\",\"packageName\":\"smsPhoneAssignments\",\"description\":\"号码记录\",\"abbreviation\":\"smsPhoneAssignments\",\"humpPackageName\":\"sms_phone_assignments\",\"gvaModel\":true,\"autoMigrate\":true,\"autoCreateResource\":false,\"autoCreateApiToSql\":true,\"autoCreateMenuToSql\":true,\"autoCreateBtnAuth\":false,\"onlyTemplate\":false,\"isTree\":false,\"treeJson\":\"\",\"isAdd\":false,\"fields\":[{\"fieldName\":\"CustomerId\",\"fieldDesc\":\"客户ID, 关联到sms_customers.id\",\"fieldType\":\"int\",\"fieldJson\":\"customerId\",\"dataTypeLong\":\"19\",\"comment\":\"客户ID, 关联到sms_customers.id\",\"columnName\":\"customer_id\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"ProviderId\",\"fieldDesc\":\"服务商ID, 关联到sms_providers.id\",\"fieldType\":\"int\",\"fieldJson\":\"providerId\",\"dataTypeLong\":\"10\",\"comment\":\"服务商ID, 关联到sms_providers.id\",\"columnName\":\"provider_id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"BusinessTypeId\",\"fieldDesc\":\"业务类型ID, 关联到sms_business_types.id\",\"fieldType\":\"int\",\"fieldJson\":\"businessTypeId\",\"dataTypeLong\":\"10\",\"comment\":\"业务类型ID, 关联到sms_business_types.id\",\"columnName\":\"business_type_id\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"CardType\",\"fieldDesc\":\"卡类型 (例如: physical, virtual)\",\"fieldType\":\"string\",\"fieldJson\":\"cardType\",\"dataTypeLong\":\"50\",\"comment\":\"卡类型 (例如: physical, virtual)\",\"columnName\":\"card_type\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"PhoneNumber\",\"fieldDesc\":\"获取到的手机号\",\"fieldType\":\"string\",\"fieldJson\":\"phoneNumber\",\"dataTypeLong\":\"50\",\"comment\":\"获取到的手机号\",\"columnName\":\"phone_number\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"VerificationCode\",\"fieldDesc\":\"获取到的验证码\",\"fieldType\":\"string\",\"fieldJson\":\"verificationCode\",\"dataTypeLong\":\"50\",\"comment\":\"获取到的验证码\",\"columnName\":\"verification_code\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Cost\",\"fieldDesc\":\"本次操作的费用\",\"fieldType\":\"float64\",\"fieldJson\":\"cost\",\"dataTypeLong\":\"10\",\"comment\":\"本次操作的费用\",\"columnName\":\"cost\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态 (1:待取码, 2:已完成, 3:已过期, 4:失败)\",\"fieldType\":\"bool\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"状态 (1:待取码, 2:已完成, 3:已过期, 4:失败)\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"ExpiresAt\",\"fieldDesc\":\"手机号锁定的过期时间\",\"fieldType\":\"time.Time\",\"fieldJson\":\"expiresAt\",\"dataTypeLong\":\"\",\"comment\":\"手机号锁定的过期时间\",\"columnName\":\"expires_at\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"}],\"generateWeb\":true,\"generateServer\":true,\"primaryField\":{\"fieldName\":\"ID\",\"fieldDesc\":\"ID\",\"fieldType\":\"uint\",\"fieldJson\":\"ID\",\"dataTypeLong\":\"20\",\"comment\":\"主键ID\",\"columnName\":\"id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":false,\"table\":false,\"desc\":false,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":false,\"sort\":false,\"primaryKey\":false,\"dataSource\":null,\"checkDataSource\":false,\"fieldIndexType\":\"\"}}','SmsPhoneAssignments','SmsPhoneAssignments','','号码记录','{\"resource/plugin/server/api/api.go.tpl\":\"plugin/sms/api/sms_phone_assignments.go\",\"resource/plugin/server/model/model.go.tpl\":\"plugin/sms/model/sms_phone_assignments.go\",\"resource/plugin/server/model/request/request.go.tpl\":\"plugin/sms/model/request/sms_phone_assignments.go\",\"resource/plugin/server/router/router.go.tpl\":\"plugin/sms/router/sms_phone_assignments.go\",\"resource/plugin/server/service/service.go.tpl\":\"plugin/sms/service/sms_phone_assignments.go\",\"resource/plugin/web/api/api.js.tpl\":\"plugin/sms/api/smsPhoneAssignments.js\",\"resource/plugin/web/form/form.vue.tpl\":\"plugin/sms/form/smsPhoneAssignments.vue\",\"resource/plugin/web/view/view.vue.tpl\":\"plugin/sms/view/smsPhoneAssignments.vue\"}','{\"PluginApiEnter\":\"{\\\"Type\\\":\\\"PluginApiEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/api/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/api/enter.go\\\",\\\"StructName\\\":\\\"SmsPhoneAssignments\\\",\\\"StructCamelName\\\":\\\"smsPhoneAssignments\\\",\\\"ModuleName\\\":\\\"serviceSmsPhoneAssignments\\\",\\\"GroupName\\\":\\\"Service\\\",\\\"PackageName\\\":\\\"service\\\",\\\"ServiceName\\\":\\\"SmsPhoneAssignments\\\"}\",\"PluginGen\":\"{\\\"Type\\\":\\\"PluginGen\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/gen/gen.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/gen/gen.go\\\",\\\"StructName\\\":\\\"SmsPhoneAssignments\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeGorm\":\"{\\\"Type\\\":\\\"PluginInitializeGorm\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/gorm.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/gorm.go\\\",\\\"StructName\\\":\\\"SmsPhoneAssignments\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeRouter\":\"{\\\"Type\\\":\\\"PluginInitializeRouter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/router.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/router\\\\\\\"\\\",\\\"ImportGlobalPath\\\":\\\"\\\",\\\"ImportMiddlewarePath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/router.go\\\",\\\"AppName\\\":\\\"Router\\\",\\\"GroupName\\\":\\\"SmsPhoneAssignments\\\",\\\"PackageName\\\":\\\"router\\\",\\\"FunctionName\\\":\\\"Init\\\",\\\"LeftRouterGroupName\\\":\\\"public\\\",\\\"RightRouterGroupName\\\":\\\"private\\\"}\",\"PluginRouterEnter\":\"{\\\"Type\\\":\\\"PluginRouterEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/router/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/router/enter.go\\\",\\\"StructName\\\":\\\"SmsPhoneAssignments\\\",\\\"StructCamelName\\\":\\\"smsPhoneAssignments\\\",\\\"ModuleName\\\":\\\"apiSmsPhoneAssignments\\\",\\\"GroupName\\\":\\\"Api\\\",\\\"PackageName\\\":\\\"api\\\",\\\"ServiceName\\\":\\\"SmsPhoneAssignments\\\"}\",\"PluginServiceEnter\":\"{\\\"Type\\\":\\\"PluginServiceEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/service/enter.go\\\",\\\"ImportPath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/service/enter.go\\\",\\\"StructName\\\":\\\"SmsPhoneAssignments\\\",\\\"StructCamelName\\\":\\\"smsPhoneAssignments\\\",\\\"ModuleName\\\":\\\"\\\",\\\"GroupName\\\":\\\"\\\",\\\"PackageName\\\":\\\"\\\",\\\"ServiceName\\\":\\\"\\\"}\"}',0,'[160,161,162,163,164,165]',40,0,0),(6,'2025-11-21 17:35:33.968','2025-11-21 17:35:33.968',NULL,'sms_providers','sms','{\"package\":\"sms\",\"tableName\":\"sms_providers\",\"businessDB\":\"\",\"structName\":\"SmsProviders\",\"packageName\":\"smsProviders\",\"description\":\"服务端\",\"abbreviation\":\"smsProviders\",\"humpPackageName\":\"sms_providers\",\"gvaModel\":true,\"autoMigrate\":true,\"autoCreateResource\":false,\"autoCreateApiToSql\":true,\"autoCreateMenuToSql\":true,\"autoCreateBtnAuth\":false,\"onlyTemplate\":false,\"isTree\":false,\"treeJson\":\"\",\"isAdd\":false,\"fields\":[{\"fieldName\":\"Name\",\"fieldDesc\":\"服务商名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"255\",\"comment\":\"服务商名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"ApiConfig\",\"fieldDesc\":\"服务商的API配置 (如URL, key等)\",\"fieldType\":\"richtext\",\"fieldJson\":\"apiConfig\",\"dataTypeLong\":\"\",\"comment\":\"服务商的API配置 (如URL, key等)\",\"columnName\":\"api_config\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"IsEnabled\",\"fieldDesc\":\"是否启用该服务商\",\"fieldType\":\"bool\",\"fieldJson\":\"isEnabled\",\"dataTypeLong\":\"\",\"comment\":\"是否启用该服务商\",\"columnName\":\"is_enabled\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"}],\"generateWeb\":true,\"generateServer\":true,\"primaryField\":{\"fieldName\":\"ID\",\"fieldDesc\":\"ID\",\"fieldType\":\"uint\",\"fieldJson\":\"ID\",\"dataTypeLong\":\"20\",\"comment\":\"主键ID\",\"columnName\":\"id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":false,\"table\":false,\"desc\":false,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":false,\"sort\":false,\"primaryKey\":false,\"dataSource\":null,\"checkDataSource\":false,\"fieldIndexType\":\"\"}}','SmsProviders','SmsProviders','','服务端','{\"resource/plugin/server/api/api.go.tpl\":\"plugin/sms/api/sms_providers.go\",\"resource/plugin/server/model/model.go.tpl\":\"plugin/sms/model/sms_providers.go\",\"resource/plugin/server/model/request/request.go.tpl\":\"plugin/sms/model/request/sms_providers.go\",\"resource/plugin/server/router/router.go.tpl\":\"plugin/sms/router/sms_providers.go\",\"resource/plugin/server/service/service.go.tpl\":\"plugin/sms/service/sms_providers.go\",\"resource/plugin/web/api/api.js.tpl\":\"plugin/sms/api/smsProviders.js\",\"resource/plugin/web/form/form.vue.tpl\":\"plugin/sms/form/smsProviders.vue\",\"resource/plugin/web/view/view.vue.tpl\":\"plugin/sms/view/smsProviders.vue\"}','{\"PluginApiEnter\":\"{\\\"Type\\\":\\\"PluginApiEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/api/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/api/enter.go\\\",\\\"StructName\\\":\\\"SmsProviders\\\",\\\"StructCamelName\\\":\\\"smsProviders\\\",\\\"ModuleName\\\":\\\"serviceSmsProviders\\\",\\\"GroupName\\\":\\\"Service\\\",\\\"PackageName\\\":\\\"service\\\",\\\"ServiceName\\\":\\\"SmsProviders\\\"}\",\"PluginGen\":\"{\\\"Type\\\":\\\"PluginGen\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/gen/gen.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/gen/gen.go\\\",\\\"StructName\\\":\\\"SmsProviders\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeGorm\":\"{\\\"Type\\\":\\\"PluginInitializeGorm\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/gorm.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/gorm.go\\\",\\\"StructName\\\":\\\"SmsProviders\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeRouter\":\"{\\\"Type\\\":\\\"PluginInitializeRouter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/router.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/router\\\\\\\"\\\",\\\"ImportGlobalPath\\\":\\\"\\\",\\\"ImportMiddlewarePath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/router.go\\\",\\\"AppName\\\":\\\"Router\\\",\\\"GroupName\\\":\\\"SmsProviders\\\",\\\"PackageName\\\":\\\"router\\\",\\\"FunctionName\\\":\\\"Init\\\",\\\"LeftRouterGroupName\\\":\\\"public\\\",\\\"RightRouterGroupName\\\":\\\"private\\\"}\",\"PluginRouterEnter\":\"{\\\"Type\\\":\\\"PluginRouterEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/router/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/router/enter.go\\\",\\\"StructName\\\":\\\"SmsProviders\\\",\\\"StructCamelName\\\":\\\"smsProviders\\\",\\\"ModuleName\\\":\\\"apiSmsProviders\\\",\\\"GroupName\\\":\\\"Api\\\",\\\"PackageName\\\":\\\"api\\\",\\\"ServiceName\\\":\\\"SmsProviders\\\"}\",\"PluginServiceEnter\":\"{\\\"Type\\\":\\\"PluginServiceEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/service/enter.go\\\",\\\"ImportPath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/service/enter.go\\\",\\\"StructName\\\":\\\"SmsProviders\\\",\\\"StructCamelName\\\":\\\"smsProviders\\\",\\\"ModuleName\\\":\\\"\\\",\\\"GroupName\\\":\\\"\\\",\\\"PackageName\\\":\\\"\\\",\\\"ServiceName\\\":\\\"\\\"}\"}',0,'[166,167,168,169,170,171]',41,0,0),(7,'2025-11-21 17:36:33.884','2025-11-21 17:36:33.884',NULL,'sms_transactions','sms','{\"package\":\"sms\",\"tableName\":\"sms_transactions\",\"businessDB\":\"\",\"structName\":\"SmsTransactions\",\"packageName\":\"smsTransactions\",\"description\":\"交易记录\",\"abbreviation\":\"smsTransactions\",\"humpPackageName\":\"sms_transactions\",\"gvaModel\":true,\"autoMigrate\":true,\"autoCreateResource\":false,\"autoCreateApiToSql\":true,\"autoCreateMenuToSql\":true,\"autoCreateBtnAuth\":false,\"onlyTemplate\":false,\"isTree\":false,\"treeJson\":\"\",\"isAdd\":false,\"fields\":[{\"fieldName\":\"CustomerId\",\"fieldDesc\":\"客户ID\",\"fieldType\":\"int\",\"fieldJson\":\"customerId\",\"dataTypeLong\":\"19\",\"comment\":\"客户ID\",\"columnName\":\"customer_id\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Amount\",\"fieldDesc\":\"变动金额 (正数为充值, 负数为消费)\",\"fieldType\":\"float64\",\"fieldJson\":\"amount\",\"dataTypeLong\":\"10\",\"comment\":\"变动金额 (正数为充值, 负数为消费)\",\"columnName\":\"amount\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"BalanceBefore\",\"fieldDesc\":\"变动前余额\",\"fieldType\":\"float64\",\"fieldJson\":\"balanceBefore\",\"dataTypeLong\":\"10\",\"comment\":\"变动前余额\",\"columnName\":\"balance_before\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"BalanceAfter\",\"fieldDesc\":\"变动后余额\",\"fieldType\":\"float64\",\"fieldJson\":\"balanceAfter\",\"dataTypeLong\":\"10\",\"comment\":\"变动后余额\",\"columnName\":\"balance_after\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Type\",\"fieldDesc\":\"交易类型 (1:充值, 2:API消费)\",\"fieldType\":\"bool\",\"fieldJson\":\"type\",\"dataTypeLong\":\"\",\"comment\":\"交易类型 (1:充值, 2:API消费)\",\"columnName\":\"type\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"ReferenceId\",\"fieldDesc\":\"关联的业务ID, 例如sms_phone_assignments.id\",\"fieldType\":\"int\",\"fieldJson\":\"referenceId\",\"dataTypeLong\":\"19\",\"comment\":\"关联的业务ID, 例如sms_phone_assignments.id\",\"columnName\":\"reference_id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Notes\",\"fieldDesc\":\"备注\",\"fieldType\":\"string\",\"fieldJson\":\"notes\",\"dataTypeLong\":\"255\",\"comment\":\"备注\",\"columnName\":\"notes\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"}],\"generateWeb\":true,\"generateServer\":true,\"primaryField\":{\"fieldName\":\"ID\",\"fieldDesc\":\"ID\",\"fieldType\":\"uint\",\"fieldJson\":\"ID\",\"dataTypeLong\":\"20\",\"comment\":\"主键ID\",\"columnName\":\"id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":false,\"table\":false,\"desc\":false,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":false,\"sort\":false,\"primaryKey\":false,\"dataSource\":null,\"checkDataSource\":false,\"fieldIndexType\":\"\"}}','SmsTransactions','SmsTransactions','','交易记录','{\"resource/plugin/server/api/api.go.tpl\":\"plugin/sms/api/sms_transactions.go\",\"resource/plugin/server/model/model.go.tpl\":\"plugin/sms/model/sms_transactions.go\",\"resource/plugin/server/model/request/request.go.tpl\":\"plugin/sms/model/request/sms_transactions.go\",\"resource/plugin/server/router/router.go.tpl\":\"plugin/sms/router/sms_transactions.go\",\"resource/plugin/server/service/service.go.tpl\":\"plugin/sms/service/sms_transactions.go\",\"resource/plugin/web/api/api.js.tpl\":\"plugin/sms/api/smsTransactions.js\",\"resource/plugin/web/form/form.vue.tpl\":\"plugin/sms/form/smsTransactions.vue\",\"resource/plugin/web/view/view.vue.tpl\":\"plugin/sms/view/smsTransactions.vue\"}','{\"PluginApiEnter\":\"{\\\"Type\\\":\\\"PluginApiEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/api/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/api/enter.go\\\",\\\"StructName\\\":\\\"SmsTransactions\\\",\\\"StructCamelName\\\":\\\"smsTransactions\\\",\\\"ModuleName\\\":\\\"serviceSmsTransactions\\\",\\\"GroupName\\\":\\\"Service\\\",\\\"PackageName\\\":\\\"service\\\",\\\"ServiceName\\\":\\\"SmsTransactions\\\"}\",\"PluginGen\":\"{\\\"Type\\\":\\\"PluginGen\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/gen/gen.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/gen/gen.go\\\",\\\"StructName\\\":\\\"SmsTransactions\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeGorm\":\"{\\\"Type\\\":\\\"PluginInitializeGorm\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/gorm.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/gorm.go\\\",\\\"StructName\\\":\\\"SmsTransactions\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeRouter\":\"{\\\"Type\\\":\\\"PluginInitializeRouter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/initialize/router.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/router\\\\\\\"\\\",\\\"ImportGlobalPath\\\":\\\"\\\",\\\"ImportMiddlewarePath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/router.go\\\",\\\"AppName\\\":\\\"Router\\\",\\\"GroupName\\\":\\\"SmsTransactions\\\",\\\"PackageName\\\":\\\"router\\\",\\\"FunctionName\\\":\\\"Init\\\",\\\"LeftRouterGroupName\\\":\\\"public\\\",\\\"RightRouterGroupName\\\":\\\"private\\\"}\",\"PluginRouterEnter\":\"{\\\"Type\\\":\\\"PluginRouterEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/router/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/router/enter.go\\\",\\\"StructName\\\":\\\"SmsTransactions\\\",\\\"StructCamelName\\\":\\\"smsTransactions\\\",\\\"ModuleName\\\":\\\"apiSmsTransactions\\\",\\\"GroupName\\\":\\\"Api\\\",\\\"PackageName\\\":\\\"api\\\",\\\"ServiceName\\\":\\\"SmsTransactions\\\"}\",\"PluginServiceEnter\":\"{\\\"Type\\\":\\\"PluginServiceEnter\\\",\\\"Path\\\":\\\"/Users/jarvis/work/tools/sms-platform/admin-server/server/plugin/sms/service/enter.go\\\",\\\"ImportPath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/service/enter.go\\\",\\\"StructName\\\":\\\"SmsTransactions\\\",\\\"StructCamelName\\\":\\\"smsTransactions\\\",\\\"ModuleName\\\":\\\"\\\",\\\"GroupName\\\":\\\"\\\",\\\"PackageName\\\":\\\"\\\",\\\"ServiceName\\\":\\\"\\\"}\"}',0,'[172,173,174,175,176,177]',42,0,0),(8,'2025-11-22 00:10:43.894','2025-11-22 00:10:43.894',NULL,'sms_providers_business_types','sms','{\"package\":\"sms\",\"tableName\":\"sms_providers_business_types\",\"businessDB\":\"\",\"structName\":\"SmsProvidersBusinessTypes\",\"packageName\":\"smsProvidersBusinessTypes\",\"description\":\"三方业务\",\"abbreviation\":\"smsProvidersBusinessTypes\",\"humpPackageName\":\"sms_providers_business_types\",\"gvaModel\":true,\"autoMigrate\":true,\"autoCreateResource\":false,\"autoCreateApiToSql\":true,\"autoCreateMenuToSql\":true,\"autoCreateBtnAuth\":false,\"onlyTemplate\":false,\"isTree\":false,\"treeJson\":\"\",\"isAdd\":false,\"fields\":[{\"fieldName\":\"ProviderId\",\"fieldDesc\":\"三方渠道ID\",\"fieldType\":\"int\",\"fieldJson\":\"providerId\",\"dataTypeLong\":\"10\",\"comment\":\"三方渠道ID（关联sms_providers表的ID）\",\"columnName\":\"provider_id\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"ProviderCode\",\"fieldDesc\":\"三方编码\",\"fieldType\":\"string\",\"fieldJson\":\"providerCode\",\"dataTypeLong\":\"50\",\"comment\":\"三方编码\",\"columnName\":\"provider_code\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"BusinessTypeId\",\"fieldDesc\":\"业务ID\",\"fieldType\":\"int\",\"fieldJson\":\"businessTypeId\",\"dataTypeLong\":\"10\",\"comment\":\"业务ID（关联sms_business_types表的ID）\",\"columnName\":\"business_type_id\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"BusinessCode\",\"fieldDesc\":\"业务编码\",\"fieldType\":\"string\",\"fieldJson\":\"businessCode\",\"dataTypeLong\":\"50\",\"comment\":\"业务编码\",\"columnName\":\"business_code\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Price\",\"fieldDesc\":\"价格\",\"fieldType\":\"float64\",\"fieldJson\":\"price\",\"dataTypeLong\":\"10\",\"comment\":\"该渠道该业务的价格\",\"columnName\":\"price\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"bool\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"该渠道是否支持该业务\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Remark\",\"fieldDesc\":\"备注\",\"fieldType\":\"string\",\"fieldJson\":\"remark\",\"dataTypeLong\":\"500\",\"comment\":\"备注\",\"columnName\":\"remark\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"}],\"generateWeb\":true,\"generateServer\":true,\"primaryField\":{\"fieldName\":\"ID\",\"fieldDesc\":\"ID\",\"fieldType\":\"uint\",\"fieldJson\":\"ID\",\"dataTypeLong\":\"20\",\"comment\":\"主键ID\",\"columnName\":\"id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":false,\"table\":false,\"desc\":false,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":false,\"sort\":false,\"primaryKey\":false,\"dataSource\":null,\"checkDataSource\":false,\"fieldIndexType\":\"\"}}','SmsProvidersBusinessTypes','SmsProvidersBusinessTypes','','三方业务','{\"resource/plugin/server/api/api.go.tpl\":\"plugin/sms/api/sms_providers_business_types.go\",\"resource/plugin/server/model/model.go.tpl\":\"plugin/sms/model/sms_providers_business_types.go\",\"resource/plugin/server/model/request/request.go.tpl\":\"plugin/sms/model/request/sms_providers_business_types.go\",\"resource/plugin/server/router/router.go.tpl\":\"plugin/sms/router/sms_providers_business_types.go\",\"resource/plugin/server/service/service.go.tpl\":\"plugin/sms/service/sms_providers_business_types.go\",\"resource/plugin/web/api/api.js.tpl\":\"plugin/sms/api/smsProvidersBusinessTypes.js\",\"resource/plugin/web/form/form.vue.tpl\":\"plugin/sms/form/smsProvidersBusinessTypes.vue\",\"resource/plugin/web/view/view.vue.tpl\":\"plugin/sms/view/smsProvidersBusinessTypes.vue\"}','{\"PluginApiEnter\":\"{\\\"Type\\\":\\\"PluginApiEnter\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/api/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/api/enter.go\\\",\\\"StructName\\\":\\\"SmsProvidersBusinessTypes\\\",\\\"StructCamelName\\\":\\\"smsProvidersBusinessTypes\\\",\\\"ModuleName\\\":\\\"serviceSmsProvidersBusinessTypes\\\",\\\"GroupName\\\":\\\"Service\\\",\\\"PackageName\\\":\\\"service\\\",\\\"ServiceName\\\":\\\"SmsProvidersBusinessTypes\\\"}\",\"PluginGen\":\"{\\\"Type\\\":\\\"PluginGen\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/gen/gen.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/gen/gen.go\\\",\\\"StructName\\\":\\\"SmsProvidersBusinessTypes\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeGorm\":\"{\\\"Type\\\":\\\"PluginInitializeGorm\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/initialize/gorm.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/gorm.go\\\",\\\"StructName\\\":\\\"SmsProvidersBusinessTypes\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeRouter\":\"{\\\"Type\\\":\\\"PluginInitializeRouter\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/initialize/router.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/router\\\\\\\"\\\",\\\"ImportGlobalPath\\\":\\\"\\\",\\\"ImportMiddlewarePath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/router.go\\\",\\\"AppName\\\":\\\"Router\\\",\\\"GroupName\\\":\\\"SmsProvidersBusinessTypes\\\",\\\"PackageName\\\":\\\"router\\\",\\\"FunctionName\\\":\\\"Init\\\",\\\"LeftRouterGroupName\\\":\\\"public\\\",\\\"RightRouterGroupName\\\":\\\"private\\\"}\",\"PluginRouterEnter\":\"{\\\"Type\\\":\\\"PluginRouterEnter\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/router/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/router/enter.go\\\",\\\"StructName\\\":\\\"SmsProvidersBusinessTypes\\\",\\\"StructCamelName\\\":\\\"smsProvidersBusinessTypes\\\",\\\"ModuleName\\\":\\\"apiSmsProvidersBusinessTypes\\\",\\\"GroupName\\\":\\\"Api\\\",\\\"PackageName\\\":\\\"api\\\",\\\"ServiceName\\\":\\\"SmsProvidersBusinessTypes\\\"}\",\"PluginServiceEnter\":\"{\\\"Type\\\":\\\"PluginServiceEnter\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/service/enter.go\\\",\\\"ImportPath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/service/enter.go\\\",\\\"StructName\\\":\\\"SmsProvidersBusinessTypes\\\",\\\"StructCamelName\\\":\\\"smsProvidersBusinessTypes\\\",\\\"ModuleName\\\":\\\"\\\",\\\"GroupName\\\":\\\"\\\",\\\"PackageName\\\":\\\"\\\",\\\"ServiceName\\\":\\\"\\\"}\"}',0,'[179,180,181,182,183,184]',43,0,0),(9,'2025-11-22 00:29:33.718','2025-11-22 00:29:33.718',NULL,'sms_platform_business_types','sms','{\"package\":\"sms\",\"tableName\":\"sms_platform_business_types\",\"businessDB\":\"\",\"structName\":\"SmsPlatformBusinessTypes\",\"packageName\":\"smsPlatformBusinessTypes\",\"description\":\"平台业务\",\"abbreviation\":\"smsPlatformBusinessTypes\",\"humpPackageName\":\"sms_platform_business_types\",\"gvaModel\":true,\"autoMigrate\":true,\"autoCreateResource\":false,\"autoCreateApiToSql\":true,\"autoCreateMenuToSql\":true,\"autoCreateBtnAuth\":false,\"onlyTemplate\":false,\"isTree\":false,\"treeJson\":\"\",\"isAdd\":false,\"fields\":[{\"fieldName\":\"Name\",\"fieldDesc\":\"平台业务名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"255\",\"comment\":\"平台业务名称\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Code\",\"fieldDesc\":\"平台业务编码\",\"fieldType\":\"string\",\"fieldJson\":\"code\",\"dataTypeLong\":\"50\",\"comment\":\"平台业务编码\",\"columnName\":\"code\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Description\",\"fieldDesc\":\"业务描述\",\"fieldType\":\"string\",\"fieldJson\":\"description\",\"dataTypeLong\":\"500\",\"comment\":\"业务描述\",\"columnName\":\"description\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Status\",\"fieldDesc\":\"启用状态\",\"fieldType\":\"bool\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"启用状态\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Remark\",\"fieldDesc\":\"备注\",\"fieldType\":\"string\",\"fieldJson\":\"remark\",\"dataTypeLong\":\"500\",\"comment\":\"备注\",\"columnName\":\"remark\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"}],\"generateWeb\":true,\"generateServer\":true,\"primaryField\":{\"fieldName\":\"ID\",\"fieldDesc\":\"ID\",\"fieldType\":\"uint\",\"fieldJson\":\"ID\",\"dataTypeLong\":\"20\",\"comment\":\"主键ID\",\"columnName\":\"id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":false,\"table\":false,\"desc\":false,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":false,\"sort\":false,\"primaryKey\":false,\"dataSource\":null,\"checkDataSource\":false,\"fieldIndexType\":\"\"}}','SmsPlatformBusinessTypes','SmsPlatformBusinessTypes','','平台业务','{\"resource/plugin/server/api/api.go.tpl\":\"plugin/sms/api/sms_platform_business_types.go\",\"resource/plugin/server/model/model.go.tpl\":\"plugin/sms/model/sms_platform_business_types.go\",\"resource/plugin/server/model/request/request.go.tpl\":\"plugin/sms/model/request/sms_platform_business_types.go\",\"resource/plugin/server/router/router.go.tpl\":\"plugin/sms/router/sms_platform_business_types.go\",\"resource/plugin/server/service/service.go.tpl\":\"plugin/sms/service/sms_platform_business_types.go\",\"resource/plugin/web/api/api.js.tpl\":\"plugin/sms/api/smsPlatformBusinessTypes.js\",\"resource/plugin/web/form/form.vue.tpl\":\"plugin/sms/form/smsPlatformBusinessTypes.vue\",\"resource/plugin/web/view/view.vue.tpl\":\"plugin/sms/view/smsPlatformBusinessTypes.vue\"}','{\"PluginApiEnter\":\"{\\\"Type\\\":\\\"PluginApiEnter\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/api/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/api/enter.go\\\",\\\"StructName\\\":\\\"SmsPlatformBusinessTypes\\\",\\\"StructCamelName\\\":\\\"smsPlatformBusinessTypes\\\",\\\"ModuleName\\\":\\\"serviceSmsPlatformBusinessTypes\\\",\\\"GroupName\\\":\\\"Service\\\",\\\"PackageName\\\":\\\"service\\\",\\\"ServiceName\\\":\\\"SmsPlatformBusinessTypes\\\"}\",\"PluginGen\":\"{\\\"Type\\\":\\\"PluginGen\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/gen/gen.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/gen/gen.go\\\",\\\"StructName\\\":\\\"SmsPlatformBusinessTypes\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeGorm\":\"{\\\"Type\\\":\\\"PluginInitializeGorm\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/initialize/gorm.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/gorm.go\\\",\\\"StructName\\\":\\\"SmsPlatformBusinessTypes\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeRouter\":\"{\\\"Type\\\":\\\"PluginInitializeRouter\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/initialize/router.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/router\\\\\\\"\\\",\\\"ImportGlobalPath\\\":\\\"\\\",\\\"ImportMiddlewarePath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/router.go\\\",\\\"AppName\\\":\\\"Router\\\",\\\"GroupName\\\":\\\"SmsPlatformBusinessTypes\\\",\\\"PackageName\\\":\\\"router\\\",\\\"FunctionName\\\":\\\"Init\\\",\\\"LeftRouterGroupName\\\":\\\"public\\\",\\\"RightRouterGroupName\\\":\\\"private\\\"}\",\"PluginRouterEnter\":\"{\\\"Type\\\":\\\"PluginRouterEnter\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/router/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/router/enter.go\\\",\\\"StructName\\\":\\\"SmsPlatformBusinessTypes\\\",\\\"StructCamelName\\\":\\\"smsPlatformBusinessTypes\\\",\\\"ModuleName\\\":\\\"apiSmsPlatformBusinessTypes\\\",\\\"GroupName\\\":\\\"Api\\\",\\\"PackageName\\\":\\\"api\\\",\\\"ServiceName\\\":\\\"SmsPlatformBusinessTypes\\\"}\",\"PluginServiceEnter\":\"{\\\"Type\\\":\\\"PluginServiceEnter\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/service/enter.go\\\",\\\"ImportPath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/service/enter.go\\\",\\\"StructName\\\":\\\"SmsPlatformBusinessTypes\\\",\\\"StructCamelName\\\":\\\"smsPlatformBusinessTypes\\\",\\\"ModuleName\\\":\\\"\\\",\\\"GroupName\\\":\\\"\\\",\\\"PackageName\\\":\\\"\\\",\\\"ServiceName\\\":\\\"\\\"}\"}',0,'[185,186,187,188,189,190]',44,0,0),(10,'2025-11-22 00:30:37.817','2025-11-22 00:30:37.817',NULL,'sms_platform_provider_business_mapping','sms','{\"package\":\"sms\",\"tableName\":\"sms_platform_provider_business_mapping\",\"businessDB\":\"\",\"structName\":\"SmsPlatformProviderBusinessMapping\",\"packageName\":\"smsPlatformProviderBusinessMapping\",\"description\":\"平台子业务\",\"abbreviation\":\"smsPlatformProviderBusinessMapping\",\"humpPackageName\":\"sms_platform_provider_business_mapping\",\"gvaModel\":true,\"autoMigrate\":true,\"autoCreateResource\":false,\"autoCreateApiToSql\":true,\"autoCreateMenuToSql\":true,\"autoCreateBtnAuth\":false,\"onlyTemplate\":false,\"isTree\":false,\"treeJson\":\"\",\"isAdd\":false,\"fields\":[{\"fieldName\":\"PlatformBusinessTypeId\",\"fieldDesc\":\"平台业务ID\",\"fieldType\":\"int\",\"fieldJson\":\"platformBusinessTypeId\",\"dataTypeLong\":\"19\",\"comment\":\"平台业务ID（关联sms_platform_business_types表的ID）\",\"columnName\":\"platform_business_type_id\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"PlatformBusinessCode\",\"fieldDesc\":\"平台业务编码\",\"fieldType\":\"string\",\"fieldJson\":\"platformBusinessCode\",\"dataTypeLong\":\"50\",\"comment\":\"平台业务编码\",\"columnName\":\"platform_business_code\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"ProviderBusinessTypeId\",\"fieldDesc\":\"三方业务ID\",\"fieldType\":\"int\",\"fieldJson\":\"providerBusinessTypeId\",\"dataTypeLong\":\"19\",\"comment\":\"三方业务ID（关联sms_providers_business_types表的ID）\",\"columnName\":\"provider_business_type_id\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"ProviderCode\",\"fieldDesc\":\"三方编码\",\"fieldType\":\"string\",\"fieldJson\":\"providerCode\",\"dataTypeLong\":\"50\",\"comment\":\"三方编码\",\"columnName\":\"provider_code\",\"fieldSearchType\":\"=\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"BusinessCode\",\"fieldDesc\":\"三方业务编码\",\"fieldType\":\"string\",\"fieldJson\":\"businessCode\",\"dataTypeLong\":\"50\",\"comment\":\"三方业务编码\",\"columnName\":\"business_code\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Weight\",\"fieldDesc\":\"权重\",\"fieldType\":\"int\",\"fieldJson\":\"weight\",\"dataTypeLong\":\"10\",\"comment\":\"权重（用于随机选择，权重越高被选中概率越大）\",\"columnName\":\"weight\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Status\",\"fieldDesc\":\"是否启用该映射\",\"fieldType\":\"bool\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"是否启用该映射\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"},{\"fieldName\":\"Remark\",\"fieldDesc\":\"备注\",\"fieldType\":\"string\",\"fieldJson\":\"remark\",\"dataTypeLong\":\"500\",\"comment\":\"备注\",\"columnName\":\"remark\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":true,\"table\":true,\"desc\":true,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":true,\"sort\":false,\"primaryKey\":false,\"dataSource\":{\"dbName\":\"\",\"table\":\"\",\"label\":\"\",\"value\":\"\",\"association\":1,\"hasDeletedAt\":false},\"checkDataSource\":false,\"fieldIndexType\":\"\"}],\"generateWeb\":true,\"generateServer\":true,\"primaryField\":{\"fieldName\":\"ID\",\"fieldDesc\":\"ID\",\"fieldType\":\"uint\",\"fieldJson\":\"ID\",\"dataTypeLong\":\"20\",\"comment\":\"主键ID\",\"columnName\":\"id\",\"fieldSearchType\":\"\",\"fieldSearchHide\":false,\"dictType\":\"\",\"form\":false,\"table\":false,\"desc\":false,\"excel\":false,\"require\":false,\"defaultValue\":\"\",\"errorText\":\"\",\"clearable\":false,\"sort\":false,\"primaryKey\":false,\"dataSource\":null,\"checkDataSource\":false,\"fieldIndexType\":\"\"}}','SmsPlatformProviderBusinessMapping','SmsPlatformProviderBusinessMapping','','平台子业务','{\"resource/plugin/server/api/api.go.tpl\":\"plugin/sms/api/sms_platform_provider_business_mapping.go\",\"resource/plugin/server/model/model.go.tpl\":\"plugin/sms/model/sms_platform_provider_business_mapping.go\",\"resource/plugin/server/model/request/request.go.tpl\":\"plugin/sms/model/request/sms_platform_provider_business_mapping.go\",\"resource/plugin/server/router/router.go.tpl\":\"plugin/sms/router/sms_platform_provider_business_mapping.go\",\"resource/plugin/server/service/service.go.tpl\":\"plugin/sms/service/sms_platform_provider_business_mapping.go\",\"resource/plugin/web/api/api.js.tpl\":\"plugin/sms/api/smsPlatformProviderBusinessMapping.js\",\"resource/plugin/web/form/form.vue.tpl\":\"plugin/sms/form/smsPlatformProviderBusinessMapping.vue\",\"resource/plugin/web/view/view.vue.tpl\":\"plugin/sms/view/smsPlatformProviderBusinessMapping.vue\"}','{\"PluginApiEnter\":\"{\\\"Type\\\":\\\"PluginApiEnter\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/api/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/service\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/api/enter.go\\\",\\\"StructName\\\":\\\"SmsPlatformProviderBusinessMapping\\\",\\\"StructCamelName\\\":\\\"smsPlatformProviderBusinessMapping\\\",\\\"ModuleName\\\":\\\"serviceSmsPlatformProviderBusinessMapping\\\",\\\"GroupName\\\":\\\"Service\\\",\\\"PackageName\\\":\\\"service\\\",\\\"ServiceName\\\":\\\"SmsPlatformProviderBusinessMapping\\\"}\",\"PluginGen\":\"{\\\"Type\\\":\\\"PluginGen\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/gen/gen.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/gen/gen.go\\\",\\\"StructName\\\":\\\"SmsPlatformProviderBusinessMapping\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeGorm\":\"{\\\"Type\\\":\\\"PluginInitializeGorm\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/initialize/gorm.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/model\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/gorm.go\\\",\\\"StructName\\\":\\\"SmsPlatformProviderBusinessMapping\\\",\\\"PackageName\\\":\\\"model\\\",\\\"IsNew\\\":true}\",\"PluginInitializeRouter\":\"{\\\"Type\\\":\\\"PluginInitializeRouter\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/initialize/router.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/router\\\\\\\"\\\",\\\"ImportGlobalPath\\\":\\\"\\\",\\\"ImportMiddlewarePath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/initialize/router.go\\\",\\\"AppName\\\":\\\"Router\\\",\\\"GroupName\\\":\\\"SmsPlatformProviderBusinessMapping\\\",\\\"PackageName\\\":\\\"router\\\",\\\"FunctionName\\\":\\\"Init\\\",\\\"LeftRouterGroupName\\\":\\\"public\\\",\\\"RightRouterGroupName\\\":\\\"private\\\"}\",\"PluginRouterEnter\":\"{\\\"Type\\\":\\\"PluginRouterEnter\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/router/enter.go\\\",\\\"ImportPath\\\":\\\"\\\\\\\"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/api\\\\\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/router/enter.go\\\",\\\"StructName\\\":\\\"SmsPlatformProviderBusinessMapping\\\",\\\"StructCamelName\\\":\\\"smsPlatformProviderBusinessMapping\\\",\\\"ModuleName\\\":\\\"apiSmsPlatformProviderBusinessMapping\\\",\\\"GroupName\\\":\\\"Api\\\",\\\"PackageName\\\":\\\"api\\\",\\\"ServiceName\\\":\\\"SmsPlatformProviderBusinessMapping\\\"}\",\"PluginServiceEnter\":\"{\\\"Type\\\":\\\"PluginServiceEnter\\\",\\\"Path\\\":\\\"/Users/nl/work/code/zhy/sms-platform/admin-server/server/plugin/sms/service/enter.go\\\",\\\"ImportPath\\\":\\\"\\\",\\\"RelativePath\\\":\\\"plugin/sms/service/enter.go\\\",\\\"StructName\\\":\\\"SmsPlatformProviderBusinessMapping\\\",\\\"StructCamelName\\\":\\\"smsPlatformProviderBusinessMapping\\\",\\\"ModuleName\\\":\\\"\\\",\\\"GroupName\\\":\\\"\\\",\\\"PackageName\\\":\\\"\\\",\\\"ServiceName\\\":\\\"\\\"}\"}',0,'[191,192,193,194,195,196]',45,0,0);
/*!40000 ALTER TABLE `sys_auto_code_histories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_auto_code_packages`
--

DROP TABLE IF EXISTS `sys_auto_code_packages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_auto_code_packages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '描述',
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '展示名',
  `template` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '模版',
  `package_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '包名',
  `module` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_packages_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_auto_code_packages`
--

LOCK TABLES `sys_auto_code_packages` WRITE;
/*!40000 ALTER TABLE `sys_auto_code_packages` DISABLE KEYS */;
INSERT INTO `sys_auto_code_packages` VALUES (1,'2025-11-21 17:24:06.449','2025-11-21 17:24:06.449',NULL,'系统自动读取example包','example包','package','example','github.com/flipped-aurora/gin-vue-admin/server'),(2,'2025-11-21 17:24:06.449','2025-11-21 17:24:06.449',NULL,'系统自动读取system包','system包','package','system','github.com/flipped-aurora/gin-vue-admin/server'),(3,'2025-11-21 17:24:06.449','2025-11-21 17:24:06.449',NULL,'系统自动读取announcement插件，使用前请确认是否为v2版本插件','announcement插件','plugin','announcement','github.com/flipped-aurora/gin-vue-admin/server'),(4,'2025-11-21 17:24:06.449','2025-11-21 17:24:06.449',NULL,'系统自动读取，但是缺少 initialize、plugin 结构，不建议自动化和mcp使用','email插件','plugin','email','github.com/flipped-aurora/gin-vue-admin/server'),(5,'2025-11-21 17:24:06.449','2025-11-21 17:24:06.449',NULL,'系统自动读取，但是缺少 config、initialize、plugin、router、service、api 结构，不建议自动化和mcp使用','plugin-tool插件','plugin','plugin-tool','github.com/flipped-aurora/gin-vue-admin/server'),(6,'2025-11-21 17:27:35.883','2025-11-21 17:27:35.883',NULL,'管理sms等东西','','plugin','sms','github.com/flipped-aurora/gin-vue-admin/server');
/*!40000 ALTER TABLE `sys_auto_code_packages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menu_btns`
--

DROP TABLE IF EXISTS `sys_base_menu_btns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menu_btns` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menu_btns`
--

LOCK TABLES `sys_base_menu_btns` WRITE;
/*!40000 ALTER TABLE `sys_base_menu_btns` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_btns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menu_parameters`
--

DROP TABLE IF EXISTS `sys_base_menu_parameters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menu_parameters`
--

LOCK TABLES `sys_base_menu_parameters` WRITE;
/*!40000 ALTER TABLE `sys_base_menu_parameters` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_parameters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_base_menus`
--

DROP TABLE IF EXISTS `sys_base_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint unsigned DEFAULT NULL,
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '高亮菜单',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '是否缓存',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '是否是基础路由（开发中）',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '菜单名',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '菜单图标',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '自动关闭tab',
  `transition_type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '路由切换动画',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_base_menus`
--

LOCK TABLES `sys_base_menus` WRITE;
/*!40000 ALTER TABLE `sys_base_menus` DISABLE KEYS */;
INSERT INTO `sys_base_menus` VALUES (1,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,0,0,'dashboard','dashboard',0,'view/dashboard/index.vue',1,'',0,0,'仪表盘','odometer',0,''),(2,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,0,0,'about','about',0,'view/about/index.vue',9,'',0,0,'关于我们','info-filled',0,''),(3,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,0,0,'admin','superAdmin',0,'view/superAdmin/index.vue',3,'',0,0,'超级管理员','user',0,''),(4,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,0,0,'person','person',1,'view/person/person.vue',4,'',0,0,'个人信息','message',0,''),(5,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,0,0,'example','example',0,'view/example/index.vue',7,'',0,0,'示例文件','management',0,''),(6,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,0,0,'systemTools','systemTools',0,'view/systemTools/index.vue',5,'',0,0,'系统工具','tools',0,''),(7,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,0,0,'https://www.gin-vue-admin.com','https://www.gin-vue-admin.com',0,'/',0,'',0,0,'官方网站','customer-gva',0,''),(8,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,0,0,'state','state',0,'view/system/state.vue',8,'',0,0,'服务器状态','cloudy',0,''),(9,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,0,0,'plugin','plugin',0,'view/routerHolder.vue',6,'',0,0,'插件系统','cherry',0,''),(10,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,3,'authority','authority',0,'view/superAdmin/authority/authority.vue',1,'',0,0,'角色管理','avatar',0,''),(11,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,3,'menu','menu',0,'view/superAdmin/menu/menu.vue',2,'',1,0,'菜单管理','tickets',0,''),(12,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,3,'api','api',0,'view/superAdmin/api/api.vue',3,'',1,0,'api管理','platform',0,''),(13,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,3,'user','user',0,'view/superAdmin/user/user.vue',4,'',0,0,'用户管理','coordinate',0,''),(14,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,3,'dictionary','dictionary',0,'view/superAdmin/dictionary/sysDictionary.vue',5,'',0,0,'字典管理','notebook',0,''),(15,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,3,'operation','operation',0,'view/superAdmin/operation/sysOperationRecord.vue',6,'',0,0,'操作历史','pie-chart',0,''),(16,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,3,'sysParams','sysParams',0,'view/superAdmin/params/sysParams.vue',7,'',0,0,'参数管理','compass',0,''),(17,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,5,'upload','upload',0,'view/example/upload/upload.vue',5,'',0,0,'媒体库（上传下载）','upload',0,''),(18,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,5,'breakpoint','breakpoint',0,'view/example/breakpoint/breakpoint.vue',6,'',0,0,'断点续传','upload-filled',0,''),(19,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,5,'customer','customer',0,'view/example/customer/customer.vue',7,'',0,0,'客户列表（资源示例）','avatar',0,''),(20,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,6,'autoCode','autoCode',0,'view/systemTools/autoCode/index.vue',1,'',1,0,'代码生成器','cpu',0,''),(21,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,6,'formCreate','formCreate',0,'view/systemTools/formCreate/index.vue',3,'',1,0,'表单生成器','magic-stick',0,''),(22,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,6,'system','system',0,'view/systemTools/system/system.vue',4,'',0,0,'系统配置','operation',0,''),(23,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,6,'autoCodeAdmin','autoCodeAdmin',0,'view/systemTools/autoCodeAdmin/index.vue',2,'',0,0,'自动化代码管理','magic-stick',0,''),(24,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,6,'autoCodeEdit/:id','autoCodeEdit',1,'view/systemTools/autoCode/index.vue',0,'',0,0,'自动化代码-${id}','magic-stick',0,''),(25,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,6,'autoPkg','autoPkg',0,'view/systemTools/autoPkg/autoPkg.vue',0,'',0,0,'模板配置','folder',0,''),(26,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,6,'exportTemplate','exportTemplate',0,'view/systemTools/exportTemplate/exportTemplate.vue',5,'',0,0,'导出模板','reading',0,''),(27,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,6,'picture','picture',0,'view/systemTools/autoCode/picture.vue',6,'',0,0,'AI页面绘制','picture-filled',0,''),(28,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,6,'mcpTool','mcpTool',0,'view/systemTools/autoCode/mcp.vue',7,'',0,0,'Mcp Tools模板','magnet',0,''),(29,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,6,'mcpTest','mcpTest',0,'view/systemTools/autoCode/mcpTest.vue',7,'',0,0,'Mcp Tools测试','partly-cloudy',0,''),(30,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,6,'sysVersion','sysVersion',0,'view/systemTools/version/version.vue',8,'',0,0,'版本管理','server',0,''),(31,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,9,'https://plugin.gin-vue-admin.com/','https://plugin.gin-vue-admin.com/',0,'https://plugin.gin-vue-admin.com/',0,'',0,0,'插件市场','shop',0,''),(32,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,9,'installPlugin','installPlugin',0,'view/systemTools/installPlugin/index.vue',1,'',0,0,'插件安装','box',0,''),(33,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,9,'pubPlug','pubPlug',0,'view/systemTools/pubPlug/pubPlug.vue',3,'',0,0,'打包插件','files',0,''),(34,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,9,'plugin-email','plugin-email',0,'plugin/email/view/index.vue',4,'',0,0,'邮件插件','message',0,''),(35,'2025-11-21 17:18:36.005','2025-11-21 17:18:36.005',NULL,1,9,'anInfo','anInfo',0,'plugin/announcement/view/info.vue',5,'',0,0,'公告管理[示例]','scaleToOriginal',0,''),(36,'2025-11-21 17:28:07.226','2025-11-22 00:14:02.147',NULL,0,0,'smsApiLogs','smsApiLogs',0,'plugin/sms/view/smsApiLogs.vue',0,'',0,0,'访问日志','alarm-clock',0,''),(38,'2025-11-21 17:30:54.258','2025-11-23 00:02:47.299',NULL,0,0,'smsCustomers','smsCustomers',0,'plugin/sms/view/smsCustomers.vue',0,'',0,0,'商户管理','dish',0,''),(39,'2025-11-21 17:31:47.503','2025-11-21 17:37:52.381',NULL,0,0,'smsIpWhitelist','smsIpWhitelist',0,'plugin/sms/view/smsIpWhitelist.vue',0,'',0,0,'白名单','flag',0,''),(40,'2025-11-21 17:33:01.945','2025-11-23 00:01:46.581',NULL,0,0,'smsPhoneAssignments','smsPhoneAssignments',0,'plugin/sms/view/smsPhoneAssignments.vue',0,'',0,0,'订单管理','film',0,''),(41,'2025-11-21 17:35:33.963','2025-11-22 23:58:53.486',NULL,0,0,'smsProviders','smsProviders',0,'plugin/sms/view/smsProviders.vue',0,'',0,0,'三方渠道','coffee',0,''),(42,'2025-11-21 17:36:33.878','2025-11-22 23:58:44.958',NULL,0,0,'smsTransactions','smsTransactions',0,'plugin/sms/view/smsTransactions.vue',0,'',0,0,'余额变动记录','chicken',0,''),(43,'2025-11-22 00:10:43.892','2025-11-22 23:59:03.263',NULL,0,0,'smsProvidersBusinessTypes','smsProvidersBusinessTypes',0,'plugin/sms/view/smsProvidersBusinessTypes.vue',0,'',0,0,'三方渠道子业务','alarm-clock',0,''),(44,'2025-11-22 00:29:33.717','2025-11-22 23:58:03.646',NULL,0,0,'smsPlatformBusinessTypes','smsPlatformBusinessTypes',0,'plugin/sms/view/smsPlatformBusinessTypes.vue',0,'',0,0,'平台业务','add-location',0,''),(45,'2025-11-22 00:30:37.816','2025-11-22 23:57:49.770',NULL,0,0,'smsPlatformProviderBusinessMapping','smsPlatformProviderBusinessMapping',0,'plugin/sms/view/smsPlatformProviderBusinessMapping.vue',0,'',0,0,'平台子业务','apple',0,'');
/*!40000 ALTER TABLE `sys_base_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_data_authority_id`
--

DROP TABLE IF EXISTS `sys_data_authority_id`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_data_authority_id`
--

LOCK TABLES `sys_data_authority_id` WRITE;
/*!40000 ALTER TABLE `sys_data_authority_id` DISABLE KEYS */;
INSERT INTO `sys_data_authority_id` VALUES (888,888),(888,8881),(888,9528),(9528,8881),(9528,9528);
/*!40000 ALTER TABLE `sys_data_authority_id` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dictionaries`
--

DROP TABLE IF EXISTS `sys_dictionaries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dictionaries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '描述',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父级字典ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dictionaries`
--

LOCK TABLES `sys_dictionaries` WRITE;
/*!40000 ALTER TABLE `sys_dictionaries` DISABLE KEYS */;
INSERT INTO `sys_dictionaries` VALUES (1,'2025-11-21 17:18:35.995','2025-11-21 17:18:35.997',NULL,'性别','gender',1,'性别字典',NULL),(2,'2025-11-21 17:18:35.995','2025-11-21 17:18:35.999',NULL,'数据库int类型','int',1,'int类型对应的数据库类型',NULL),(3,'2025-11-21 17:18:35.995','2025-11-21 17:18:36.000',NULL,'数据库时间日期类型','time.Time',1,'数据库时间日期类型',NULL),(4,'2025-11-21 17:18:35.995','2025-11-21 17:18:36.001',NULL,'数据库浮点型','float64',1,'数据库浮点型',NULL),(5,'2025-11-21 17:18:35.995','2025-11-21 17:18:36.002',NULL,'数据库字符串','string',1,'数据库字符串',NULL),(6,'2025-11-21 17:18:35.995','2025-11-21 17:18:36.003',NULL,'数据库bool类型','bool',1,'数据库bool类型',NULL),(7,'2025-11-21 18:15:56.877','2025-11-21 23:15:01.850',NULL,'变动类型','transaction_type',1,'金额操作变动类型',NULL);
/*!40000 ALTER TABLE `sys_dictionaries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dictionary_details`
--

DROP TABLE IF EXISTS `sys_dictionary_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '展示值',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint unsigned DEFAULT NULL COMMENT '关联标记',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父级字典详情ID',
  `level` bigint DEFAULT NULL COMMENT '层级深度',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '层级路径',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dictionary_details`
--

LOCK TABLES `sys_dictionary_details` WRITE;
/*!40000 ALTER TABLE `sys_dictionary_details` DISABLE KEYS */;
INSERT INTO `sys_dictionary_details` VALUES (1,'2025-11-21 17:18:35.997','2025-11-21 17:18:35.997',NULL,'男','1','',1,1,1,NULL,0,''),(2,'2025-11-21 17:18:35.997','2025-11-21 17:18:35.997',NULL,'女','2','',1,2,1,NULL,0,''),(3,'2025-11-21 17:18:35.999','2025-11-21 17:18:35.999',NULL,'smallint','1','mysql',1,1,2,NULL,0,''),(4,'2025-11-21 17:18:35.999','2025-11-21 17:18:35.999',NULL,'mediumint','2','mysql',1,2,2,NULL,0,''),(5,'2025-11-21 17:18:35.999','2025-11-21 17:18:35.999',NULL,'int','3','mysql',1,3,2,NULL,0,''),(6,'2025-11-21 17:18:35.999','2025-11-21 17:18:35.999',NULL,'bigint','4','mysql',1,4,2,NULL,0,''),(7,'2025-11-21 17:18:35.999','2025-11-21 17:18:35.999',NULL,'int2','5','pgsql',1,5,2,NULL,0,''),(8,'2025-11-21 17:18:35.999','2025-11-21 17:18:35.999',NULL,'int4','6','pgsql',1,6,2,NULL,0,''),(9,'2025-11-21 17:18:35.999','2025-11-21 17:18:35.999',NULL,'int6','7','pgsql',1,7,2,NULL,0,''),(10,'2025-11-21 17:18:35.999','2025-11-21 17:18:35.999',NULL,'int8','8','pgsql',1,8,2,NULL,0,''),(11,'2025-11-21 17:18:36.000','2025-11-21 17:18:36.000',NULL,'date','0','mysql',1,0,3,NULL,0,''),(12,'2025-11-21 17:18:36.000','2025-11-21 17:18:36.000',NULL,'time','1','mysql',1,1,3,NULL,0,''),(13,'2025-11-21 17:18:36.000','2025-11-21 17:18:36.000',NULL,'year','2','mysql',1,2,3,NULL,0,''),(14,'2025-11-21 17:18:36.000','2025-11-21 17:18:36.000',NULL,'datetime','3','mysql',1,3,3,NULL,0,''),(15,'2025-11-21 17:18:36.000','2025-11-21 17:18:36.000',NULL,'timestamp','5','mysql',1,5,3,NULL,0,''),(16,'2025-11-21 17:18:36.000','2025-11-21 17:18:36.000',NULL,'timestamptz','6','pgsql',1,5,3,NULL,0,''),(17,'2025-11-21 17:18:36.001','2025-11-21 17:18:36.001',NULL,'float','0','mysql',1,0,4,NULL,0,''),(18,'2025-11-21 17:18:36.001','2025-11-21 17:18:36.001',NULL,'double','1','mysql',1,1,4,NULL,0,''),(19,'2025-11-21 17:18:36.001','2025-11-21 17:18:36.001',NULL,'decimal','2','mysql',1,2,4,NULL,0,''),(20,'2025-11-21 17:18:36.001','2025-11-21 17:18:36.001',NULL,'numeric','3','pgsql',1,3,4,NULL,0,''),(21,'2025-11-21 17:18:36.001','2025-11-21 17:18:36.001',NULL,'smallserial','4','pgsql',1,4,4,NULL,0,''),(22,'2025-11-21 17:18:36.002','2025-11-21 17:18:36.002',NULL,'char','0','mysql',1,0,5,NULL,0,''),(23,'2025-11-21 17:18:36.002','2025-11-21 17:18:36.002',NULL,'varchar','1','mysql',1,1,5,NULL,0,''),(24,'2025-11-21 17:18:36.002','2025-11-21 17:18:36.002',NULL,'tinyblob','2','mysql',1,2,5,NULL,0,''),(25,'2025-11-21 17:18:36.002','2025-11-21 17:18:36.002',NULL,'tinytext','3','mysql',1,3,5,NULL,0,''),(26,'2025-11-21 17:18:36.002','2025-11-21 17:18:36.002',NULL,'text','4','mysql',1,4,5,NULL,0,''),(27,'2025-11-21 17:18:36.002','2025-11-21 17:18:36.002',NULL,'blob','5','mysql',1,5,5,NULL,0,''),(28,'2025-11-21 17:18:36.002','2025-11-21 17:18:36.002',NULL,'mediumblob','6','mysql',1,6,5,NULL,0,''),(29,'2025-11-21 17:18:36.002','2025-11-21 17:18:36.002',NULL,'mediumtext','7','mysql',1,7,5,NULL,0,''),(30,'2025-11-21 17:18:36.002','2025-11-21 17:18:36.002',NULL,'longblob','8','mysql',1,8,5,NULL,0,''),(31,'2025-11-21 17:18:36.002','2025-11-21 17:18:36.002',NULL,'longtext','9','mysql',1,9,5,NULL,0,''),(32,'2025-11-21 17:18:36.003','2025-11-21 17:18:36.003',NULL,'tinyint','1','mysql',1,0,6,NULL,0,''),(33,'2025-11-21 17:18:36.003','2025-11-21 17:18:36.003',NULL,'bool','2','pgsql',1,0,6,NULL,0,''),(34,'2025-11-21 18:17:03.069','2025-11-21 18:17:03.069',NULL,'充值','1','',1,1,7,NULL,0,''),(35,'2025-11-21 18:17:48.752','2025-11-21 18:17:48.752',NULL,'拉号码','2','',1,1,7,NULL,0,''),(36,'2025-11-21 18:17:59.291','2025-11-21 18:17:59.291',NULL,'拉号-回退','3','',1,1,7,NULL,0,''),(37,'2025-11-21 18:18:11.811','2025-11-21 18:18:11.811',NULL,'上分','4','',1,4,7,NULL,0,''),(38,'2025-11-21 18:18:27.079','2025-11-21 18:18:27.079',NULL,'下分','5','',1,5,7,NULL,0,''),(39,'2025-11-22 23:47:44.832','2025-11-22 23:47:44.832',NULL,'冻结金额','6','',1,6,7,NULL,0,''),(40,'2025-11-22 23:47:58.713','2025-11-22 23:47:58.713',NULL,'冻结金额返回','7','',1,7,7,NULL,0,'');
/*!40000 ALTER TABLE `sys_dictionary_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_export_template_condition`
--

DROP TABLE IF EXISTS `sys_export_template_condition`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_template_condition` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '模板标识',
  `from` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '条件取的key',
  `column` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '作为查询条件的字段',
  `operator` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '操作符',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_template_condition_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_export_template_condition`
--

LOCK TABLES `sys_export_template_condition` WRITE;
/*!40000 ALTER TABLE `sys_export_template_condition` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_export_template_condition` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_export_template_join`
--

DROP TABLE IF EXISTS `sys_export_template_join`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_template_join` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '模板标识',
  `joins` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '关联',
  `table` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '关联表',
  `on` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '关联条件',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_template_join_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_export_template_join`
--

LOCK TABLES `sys_export_template_join` WRITE;
/*!40000 ALTER TABLE `sys_export_template_join` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_export_template_join` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_export_templates`
--

DROP TABLE IF EXISTS `sys_export_templates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_templates` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `db_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '数据库名称',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '模板名称',
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '表名称',
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '模板标识',
  `template_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin,
  `limit` bigint DEFAULT NULL COMMENT '导出限制',
  `order` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_templates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_export_templates`
--

LOCK TABLES `sys_export_templates` WRITE;
/*!40000 ALTER TABLE `sys_export_templates` DISABLE KEYS */;
INSERT INTO `sys_export_templates` VALUES (1,'2025-11-21 17:18:36.114','2025-11-21 17:18:36.114',NULL,'','api','sys_apis','api','{\n\"path\":\"路径\",\n\"method\":\"方法（大写）\",\n\"description\":\"方法介绍\",\n\"api_group\":\"方法分组\"\n}',NULL,'');
/*!40000 ALTER TABLE `sys_export_templates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_ignore_apis`
--

DROP TABLE IF EXISTS `sys_ignore_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_ignore_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'api路径',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_ignore_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_ignore_apis`
--

LOCK TABLES `sys_ignore_apis` WRITE;
/*!40000 ALTER TABLE `sys_ignore_apis` DISABLE KEYS */;
INSERT INTO `sys_ignore_apis` VALUES (1,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/swagger/*any','GET'),(2,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/api/freshCasbin','GET'),(3,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/uploads/file/*filepath','GET'),(4,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/health','GET'),(5,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/uploads/file/*filepath','HEAD'),(6,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/autoCode/llmAuto','POST'),(7,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/system/reloadSystem','POST'),(8,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/base/login','POST'),(9,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/base/captcha','POST'),(10,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/init/initdb','POST'),(11,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/init/checkdb','POST'),(12,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/info/getInfoDataSource','GET'),(13,'2025-11-21 17:18:35.988','2025-11-21 17:18:35.988',NULL,'/info/getInfoPublic','GET');
/*!40000 ALTER TABLE `sys_ignore_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_operation_records`
--

DROP TABLE IF EXISTS `sys_operation_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_operation_records` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求路径',
  `status` bigint DEFAULT NULL COMMENT '请求状态',
  `latency` bigint DEFAULT NULL COMMENT '延迟',
  `agent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '响应Body',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_operation_records`
--

LOCK TABLES `sys_operation_records` WRITE;
/*!40000 ALTER TABLE `sys_operation_records` DISABLE KEYS */;
INSERT INTO `sys_operation_records` VALUES (1,'2025-11-21 17:37:08.455','2025-11-21 17:37:08.455',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,14321208,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":42,\"CreatedAt\":\"2025-11-21T17:36:33.878+08:00\",\"UpdatedAt\":\"2025-11-21T17:36:33.878+08:00\",\"parentId\":0,\"path\":\"smsTransactions\",\"name\":\"smsTransactions\",\"hidden\":false,\"component\":\"plugin/sms/view/smsTransactions.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"交易记录\",\"icon\":\"chicken\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(2,'2025-11-21 17:37:30.805','2025-11-21 17:37:30.805',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,19047500,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":41,\"CreatedAt\":\"2025-11-21T17:35:33.963+08:00\",\"UpdatedAt\":\"2025-11-21T17:35:33.963+08:00\",\"parentId\":0,\"path\":\"smsProviders\",\"name\":\"smsProviders\",\"hidden\":false,\"component\":\"plugin/sms/view/smsProviders.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务商\",\"icon\":\"coffee\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(3,'2025-11-21 17:37:46.125','2025-11-21 17:37:46.125',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,13915208,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":40,\"CreatedAt\":\"2025-11-21T17:33:01.945+08:00\",\"UpdatedAt\":\"2025-11-21T17:33:01.945+08:00\",\"parentId\":0,\"path\":\"smsPhoneAssignments\",\"name\":\"smsPhoneAssignments\",\"hidden\":false,\"component\":\"plugin/sms/view/smsPhoneAssignments.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"号码记录\",\"icon\":\"film\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(4,'2025-11-21 17:37:52.394','2025-11-21 17:37:52.394',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,14709250,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":39,\"CreatedAt\":\"2025-11-21T17:31:47.503+08:00\",\"UpdatedAt\":\"2025-11-21T17:31:47.503+08:00\",\"parentId\":0,\"path\":\"smsIpWhitelist\",\"name\":\"smsIpWhitelist\",\"hidden\":false,\"component\":\"plugin/sms/view/smsIpWhitelist.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"白名单\",\"icon\":\"flag\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(5,'2025-11-21 17:38:15.845','2025-11-21 17:38:15.845',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,19856250,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":38,\"CreatedAt\":\"2025-11-21T17:30:54.258+08:00\",\"UpdatedAt\":\"2025-11-21T17:30:54.258+08:00\",\"parentId\":0,\"path\":\"smsCustomers\",\"name\":\"smsCustomers\",\"hidden\":false,\"component\":\"plugin/sms/view/smsCustomers.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"商户\",\"icon\":\"dish\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(6,'2025-11-21 17:38:26.866','2025-11-21 17:38:26.866',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,13669167,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":37,\"CreatedAt\":\"2025-11-21T17:29:31.657+08:00\",\"UpdatedAt\":\"2025-11-21T17:29:31.657+08:00\",\"parentId\":0,\"path\":\"smsBusinessTypes\",\"name\":\"smsBusinessTypes\",\"hidden\":false,\"component\":\"plugin/sms/view/smsBusinessTypes.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"业务类型\",\"icon\":\"help\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(7,'2025-11-21 17:38:36.116','2025-11-21 17:38:36.116',NULL,'127.0.0.1','GET','/api/getApiGroups',200,617000,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"apiGroupMap\":{\"api\":\"api\",\"attachmentCategory\":\"媒体库分类\",\"authority\":\"角色\",\"authorityBtn\":\"按钮权限\",\"autoCode\":\"代码生成器历史\",\"casbin\":\"casbin\",\"customer\":\"客户\",\"email\":\"email\",\"fileUploadAndDownload\":\"文件上传与下载\",\"info\":\"公告\",\"jwt\":\"jwt\",\"menu\":\"菜单\",\"simpleUploader\":\"断点续传(插件版)\",\"smsApiLogs\":\"访问日志\",\"smsBusinessTypes\":\"业务类型\",\"smsCustomers\":\"商户\",\"smsIpWhitelist\":\"白名单\",\"smsPhoneAssignments\":\"号码记录\",\"smsProviders\":\"服务端\",\"smsTransactions\":\"交易记录\",\"sysDictionary\":\"系统字典\",\"sysDictionaryDetail\":\"系统字典详情\",\"sysExportTemplate\":\"导出模板\",\"sysOperationRecord\":\"操作记录\",\"sysParams\":\"参数管理\",\"sysVersion\":\"版本控制\",\"system\":\"系统服务\",\"user\":\"系统用户\"},\"groups\":[\"jwt\",\"系统用户\",\"api\",\"角色\",\"casbin\",\"菜单\",\"分片上传\",\"文件上传与下载\",\"系统服务\",\"客户\",\"代码生成器\",\"模板配置\",\"代码生成器历史\",\"系统字典详情\",\"系统字典\",\"操作记录\",\"断点续传(插件版)\",\"email\",\"按钮权限\",\"导出模板\",\"公告\",\"参数管理\",\"媒体库分类\",\"版本控制\",\"访问日志\",\"业务类型\",\"商户\",\"白名单\",\"号码记录\",\"服务端\",\"交易记录\"]},\"msg\":\"成功\"}',1),(8,'2025-11-21 17:38:55.900','2025-11-21 17:38:55.900',NULL,'127.0.0.1','POST','/menu/addMenuAuthority',200,16824875,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','[超出记录长度]','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),(9,'2025-11-21 17:39:12.738','2025-11-21 17:39:12.738',NULL,'127.0.0.1','POST','/casbin/updateCasbin',200,21102959,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','[超出记录长度]','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(10,'2025-11-21 17:45:08.395','2025-11-21 17:45:08.395',NULL,'127.0.0.1','POST','/smsBusinessTypes/createSmsBusinessTypes',200,1801084,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"id\":1,\"name\":\"qq\",\"code\":\"qq\",\"isEnabled\":true}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(11,'2025-11-21 17:45:19.974','2025-11-21 17:45:19.974',NULL,'127.0.0.1','POST','/smsBusinessTypes/createSmsBusinessTypes',200,1897209,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"id\":2,\"name\":\"wx\",\"code\":\"wx\",\"isEnabled\":true}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(12,'2025-11-21 17:53:45.732','2025-11-21 17:53:45.732',NULL,'127.0.0.1','POST','/smsCustomers/createSmsCustomers',200,57217666,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"username\":\"test\",\"email\":\"111@qq.com\",\"password\":\"1234566\",\"passwordHash\":\"\",\"apiSecretKey\":\"\",\"balance\":1000,\"status\":true}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(13,'2025-11-21 18:15:56.900','2025-11-21 18:15:56.900',NULL,'127.0.0.1','POST','/sysDictionary/createSysDictionary',200,24340708,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"name\":\"变动类型\",\"type\":\"transatction_type\",\"status\":true,\"desc\":\"金额操作变动类型\",\"parentID\":null}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(14,'2025-11-21 18:17:03.081','2025-11-21 18:17:03.081',NULL,'127.0.0.1','POST','/sysDictionaryDetail/createSysDictionaryDetail',200,12917375,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"label\":\"充值\",\"value\":\"1\",\"status\":true,\"sort\":1,\"parentID\":null,\"sysDictionaryID\":7}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(15,'2025-11-21 18:17:48.770','2025-11-21 18:17:48.770',NULL,'127.0.0.1','POST','/sysDictionaryDetail/createSysDictionaryDetail',200,18631333,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"label\":\"拉号码\",\"value\":\"2\",\"status\":true,\"sort\":1,\"parentID\":null,\"sysDictionaryID\":7}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(16,'2025-11-21 18:17:59.304','2025-11-21 18:17:59.304',NULL,'127.0.0.1','POST','/sysDictionaryDetail/createSysDictionaryDetail',200,12625833,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"label\":\"拉号-回退\",\"value\":\"3\",\"status\":true,\"sort\":1,\"parentID\":null,\"sysDictionaryID\":7}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(17,'2025-11-21 18:18:11.823','2025-11-21 18:18:11.823',NULL,'127.0.0.1','POST','/sysDictionaryDetail/createSysDictionaryDetail',200,12597500,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"label\":\"上分\",\"value\":\"4\",\"status\":true,\"sort\":4,\"parentID\":null,\"sysDictionaryID\":7}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(18,'2025-11-21 18:18:27.102','2025-11-21 18:18:27.102',NULL,'127.0.0.1','POST','/sysDictionaryDetail/createSysDictionaryDetail',200,23981917,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"label\":\"下分\",\"value\":\"5\",\"status\":true,\"sort\":5,\"parentID\":null,\"sysDictionaryID\":7}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(19,'2025-11-21 18:19:49.455','2025-11-21 18:19:49.455',NULL,'127.0.0.1','PUT','/smsCustomers/updateSmsCustomers',200,12615250,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":1,\"CreatedAt\":\"2025-11-21T17:53:45.73+08:00\",\"UpdatedAt\":\"2025-11-21T17:53:45.73+08:00\",\"username\":\"test\",\"email\":\"111@qq.com\",\"passwordHash\":\"$2a$10$xPBwRscV2SEnj38K5tzd0uFy5JKeQs8W6VkWGvG2fqt6c4Ho..gGy\",\"apiSecretKey\":\"676494626b8738264235e69a781bbac50e8be469fba02ea51bc0c48142274711\",\"balance\":1002,\"status\":true,\"registrationIp\":\"127.0.0.1\",\"lastLoginIp\":null,\"lastLoginAt\":null}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(20,'2025-11-21 18:25:06.120','2025-11-21 18:25:06.120',NULL,'127.0.0.1','GET','/api/getApiGroups',200,605000,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"apiGroupMap\":{\"api\":\"api\",\"attachmentCategory\":\"媒体库分类\",\"authority\":\"角色\",\"authorityBtn\":\"按钮权限\",\"autoCode\":\"代码生成器历史\",\"casbin\":\"casbin\",\"customer\":\"客户\",\"email\":\"email\",\"fileUploadAndDownload\":\"文件上传与下载\",\"info\":\"公告\",\"jwt\":\"jwt\",\"menu\":\"菜单\",\"simpleUploader\":\"断点续传(插件版)\",\"smsApiLogs\":\"访问日志\",\"smsBusinessTypes\":\"业务类型\",\"smsCustomers\":\"商户\",\"smsIpWhitelist\":\"白名单\",\"smsPhoneAssignments\":\"号码记录\",\"smsProviders\":\"服务端\",\"smsTransactions\":\"交易记录\",\"sysDictionary\":\"系统字典\",\"sysDictionaryDetail\":\"系统字典详情\",\"sysExportTemplate\":\"导出模板\",\"sysOperationRecord\":\"操作记录\",\"sysParams\":\"参数管理\",\"sysVersion\":\"版本控制\",\"system\":\"系统服务\",\"user\":\"系统用户\"},\"groups\":[\"jwt\",\"系统用户\",\"api\",\"角色\",\"casbin\",\"菜单\",\"分片上传\",\"文件上传与下载\",\"系统服务\",\"客户\",\"代码生成器\",\"模板配置\",\"代码生成器历史\",\"系统字典详情\",\"系统字典\",\"操作记录\",\"断点续传(插件版)\",\"email\",\"按钮权限\",\"导出模板\",\"公告\",\"参数管理\",\"媒体库分类\",\"版本控制\",\"访问日志\",\"业务类型\",\"商户\",\"白名单\",\"号码记录\",\"服务端\",\"交易记录\"]},\"msg\":\"成功\"}',1),(21,'2025-11-21 18:25:47.905','2025-11-21 18:25:47.905',NULL,'127.0.0.1','POST','/api/createApi',200,18359541,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"path\":\"/smsCustomers/creditDebit\",\"apiGroup\":\"商户\",\"method\":\"POST\",\"description\":\"上下分\"}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(22,'2025-11-21 18:25:47.921','2025-11-21 18:25:47.921',NULL,'127.0.0.1','GET','/api/getApiGroups',200,758500,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"apiGroupMap\":{\"api\":\"api\",\"attachmentCategory\":\"媒体库分类\",\"authority\":\"角色\",\"authorityBtn\":\"按钮权限\",\"autoCode\":\"代码生成器历史\",\"casbin\":\"casbin\",\"customer\":\"客户\",\"email\":\"email\",\"fileUploadAndDownload\":\"文件上传与下载\",\"info\":\"公告\",\"jwt\":\"jwt\",\"menu\":\"菜单\",\"simpleUploader\":\"断点续传(插件版)\",\"smsApiLogs\":\"访问日志\",\"smsBusinessTypes\":\"业务类型\",\"smsCustomers\":\"商户\",\"smsIpWhitelist\":\"白名单\",\"smsPhoneAssignments\":\"号码记录\",\"smsProviders\":\"服务端\",\"smsTransactions\":\"交易记录\",\"sysDictionary\":\"系统字典\",\"sysDictionaryDetail\":\"系统字典详情\",\"sysExportTemplate\":\"导出模板\",\"sysOperationRecord\":\"操作记录\",\"sysParams\":\"参数管理\",\"sysVersion\":\"版本控制\",\"system\":\"系统服务\",\"user\":\"系统用户\"},\"groups\":[\"jwt\",\"系统用户\",\"api\",\"角色\",\"casbin\",\"菜单\",\"分片上传\",\"文件上传与下载\",\"系统服务\",\"客户\",\"代码生成器\",\"模板配置\",\"代码生成器历史\",\"系统字典详情\",\"系统字典\",\"操作记录\",\"断点续传(插件版)\",\"email\",\"按钮权限\",\"导出模板\",\"公告\",\"参数管理\",\"媒体库分类\",\"版本控制\",\"访问日志\",\"业务类型\",\"商户\",\"白名单\",\"号码记录\",\"服务端\",\"交易记录\"]},\"msg\":\"成功\"}',1),(23,'2025-11-21 18:25:55.523','2025-11-21 18:25:55.523',NULL,'127.0.0.1','POST','/casbin/updateCasbin',200,20154750,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','[超出记录长度]','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(24,'2025-11-21 18:26:06.242','2025-11-21 18:26:06.242',NULL,'127.0.0.1','POST','/smsCustomers/creditDebit',200,18942584,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"customerId\":1,\"amount\":1,\"type\":1,\"notes\":\"1212\"}','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),(25,'2025-11-21 23:13:08.479','2025-11-21 23:13:08.479',NULL,'127.0.0.1','POST','/smsCustomers/creditDebit',200,5944291,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"customerId\":1,\"amount\":1,\"type\":1}','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),(26,'2025-11-21 23:15:01.854','2025-11-21 23:15:01.854',NULL,'127.0.0.1','PUT','/sysDictionary/updateSysDictionary',200,6315334,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','[超出记录长度]','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(27,'2025-11-21 23:38:20.224','2025-11-21 23:38:20.224',NULL,'127.0.0.1','POST','/smsCustomers/creditDebit',200,51166,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"customerId\":1,\"amount\":1,\"type\":2}','{\"code\":7,\"data\":{},\"msg\":\"json: cannot unmarshal number into Go struct field CreditDebitSmsCustomersReq.type of type string\"}',1),(28,'2025-11-21 23:44:31.858','2025-11-21 23:44:31.858',NULL,'127.0.0.1','POST','/smsCustomers/creditDebit',200,9888333,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"customerId\":1,\"amount\":1,\"type\":\"2\"}','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),(29,'2025-11-22 00:13:37.359','2025-11-22 00:13:37.359',NULL,'127.0.0.1','POST','/menu/addMenuAuthority',200,11382667,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','[超出记录长度]','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),(30,'2025-11-22 00:13:39.501','2025-11-22 00:13:39.501',NULL,'127.0.0.1','POST','/casbin/updateCasbin',200,21061958,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','[超出记录长度]','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(31,'2025-11-22 00:13:52.475','2025-11-22 00:13:52.475',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,3838917,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":43,\"CreatedAt\":\"2025-11-22T00:10:43.892+08:00\",\"UpdatedAt\":\"2025-11-22T00:10:43.892+08:00\",\"parentId\":0,\"path\":\"smsProvidersBusinessTypes\",\"name\":\"smsProvidersBusinessTypes\",\"hidden\":false,\"component\":\"plugin/sms/view/smsProvidersBusinessTypes.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"三方业务\",\"icon\":\"alarm-clock\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(32,'2025-11-22 00:14:02.149','2025-11-22 00:14:02.149',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,3744292,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":36,\"CreatedAt\":\"2025-11-21T17:28:07.226+08:00\",\"UpdatedAt\":\"2025-11-21T17:28:07.226+08:00\",\"parentId\":0,\"path\":\"smsApiLogs\",\"name\":\"smsApiLogs\",\"hidden\":false,\"component\":\"plugin/sms/view/smsApiLogs.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"访问日志\",\"icon\":\"alarm-clock\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(33,'2025-11-22 00:35:32.994','2025-11-22 00:35:32.994',NULL,'127.0.0.1','POST','/menu/addMenuAuthority',200,12542333,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','[超出记录长度]','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),(34,'2025-11-22 00:36:04.140','2025-11-22 00:36:04.140',NULL,'127.0.0.1','POST','/menu/addMenuAuthority',200,11436625,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','[超出记录长度]','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),(35,'2025-11-22 00:36:51.830','2025-11-22 00:36:51.830',NULL,'127.0.0.1','POST','/casbin/updateCasbin',200,22561750,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','[超出记录长度]','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(36,'2025-11-22 22:42:36.189','2025-11-22 22:42:36.189',NULL,'127.0.0.1','POST','/smsProviders/createSmsProviders',200,5365625,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"name\":\"测试\",\"code\":\"test\",\"apiGateway\":\"https://www.google.com\",\"merchantId\":\"merch_no\",\"merchantKey\":\"12121212\",\"status\":true,\"remark\":\"imya \"}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(37,'2025-11-22 22:54:35.970','2025-11-22 22:54:35.970',NULL,'127.0.0.1','POST','/smsProvidersBusinessTypes/createSmsProvidersBusinessTypes',200,7171416,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"providerId\":1,\"providerCode\":\"test\",\"businessName\":\"测试业务\",\"businessCode\":\"qq\",\"price\":1,\"status\":true,\"remark\":\"1212\"}','{\"code\":7,\"data\":{},\"msg\":\"创建失败:Error 1054 (42S22): Unknown column \'business_type_id\' in \'field list\'\"}',1),(38,'2025-11-22 22:54:59.224','2025-11-22 22:54:59.224',NULL,'127.0.0.1','POST','/smsProvidersBusinessTypes/createSmsProvidersBusinessTypes',200,3297916,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"providerId\":1,\"providerCode\":\"test\",\"businessName\":\"测试业务\",\"businessCode\":\"qq\",\"price\":1,\"status\":true,\"remark\":\"1212\"}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(39,'2025-11-22 22:55:27.879','2025-11-22 22:55:27.879',NULL,'127.0.0.1','POST','/smsProviders/createSmsProviders',200,2412375,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"name\":\"test2\",\"code\":\"test\",\"apiGateway\":\"http://www.baidu.com\",\"merchantId\":\"1121212\",\"merchantKey\":\"123456121212\",\"status\":true,\"remark\":\"12121\"}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(40,'2025-11-22 22:55:44.533','2025-11-22 22:55:44.533',NULL,'127.0.0.1','POST','/smsProvidersBusinessTypes/createSmsProvidersBusinessTypes',200,4161458,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"providerId\":2,\"providerCode\":\"test\",\"businessName\":\"tmg \",\"businessCode\":\"qq\",\"price\":10,\"status\":true,\"remark\":\"1212\"}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(41,'2025-11-22 22:56:06.880','2025-11-22 22:56:06.880',NULL,'127.0.0.1','POST','/smsPlatformBusinessTypes/createSmsPlatformBusinessTypes',200,2424292,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"name\":\"微信\",\"code\":\"wx\",\"description\":\"sdfadf\",\"status\":true,\"remark\":\"asdfadf\"}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(42,'2025-11-22 23:04:02.838','2025-11-22 23:04:02.838',NULL,'127.0.0.1','POST','/smsPlatformProviderBusinessMapping/createSmsPlatformProviderBusinessMapping',200,4245667,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"platformBusinessTypeId\":1,\"platformBusinessCode\":\"wx\",\"providerBusinessTypeId\":1,\"providerCode\":\"test\",\"businessCode\":\"qq\",\"weight\":11,\"status\":true,\"remark\":\"1212\"}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(43,'2025-11-22 23:46:26.420','2025-11-22 23:46:26.420',NULL,'127.0.0.1','POST','/smsCustomers/creditDebit',200,6936750,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"customerId\":1,\"amount\":10,\"type\":\"6\"}','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),(44,'2025-11-22 23:47:44.836','2025-11-22 23:47:44.836',NULL,'127.0.0.1','POST','/sysDictionaryDetail/createSysDictionaryDetail',200,4217917,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"label\":\"冻结金额\",\"value\":\"6\",\"status\":true,\"sort\":6,\"parentID\":null,\"sysDictionaryID\":7}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(45,'2025-11-22 23:47:58.715','2025-11-22 23:47:58.715',NULL,'127.0.0.1','POST','/sysDictionaryDetail/createSysDictionaryDetail',200,3067583,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"label\":\"冻结金额返回\",\"value\":\"7\",\"status\":true,\"sort\":7,\"parentID\":null,\"sysDictionaryID\":7}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(46,'2025-11-22 23:48:46.524','2025-11-22 23:48:46.524',NULL,'127.0.0.1','POST','/smsCustomers/creditDebit',200,7205667,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"customerId\":1,\"amount\":10,\"type\":\"7\"}','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),(47,'2025-11-22 23:49:33.413','2025-11-22 23:49:33.413',NULL,'127.0.0.1','POST','/smsPlatformBusinessTypes/createSmsPlatformBusinessTypes',200,3945375,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"name\":\"qq\",\"code\":\"qq\",\"description\":\"1qq\",\"status\":false,\"remark\":\"\"}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),(48,'2025-11-22 23:51:53.672','2025-11-22 23:51:53.672',NULL,'127.0.0.1','PUT','/smsCustomers/updateSmsCustomers',200,4633875,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":1,\"CreatedAt\":\"2025-11-21T17:53:45.73+08:00\",\"UpdatedAt\":\"2025-11-22T23:48:46.518+08:00\",\"merchantName\":\"12121\",\"merchantNo\":\"121212\",\"username\":\"test\",\"email\":\"111@qq.com\",\"passwordHash\":\"$2a$10$xPBwRscV2SEnj38K5tzd0uFy5JKeQs8W6VkWGvG2fqt6c4Ho..gGy\",\"apiSecretKey\":\"676494626b8738264235e69a781bbac50e8be469fba02ea51bc0c48142274711\",\"balance\":1003,\"parentId\":null,\"frozenAmount\":0,\"status\":true,\"registrationIp\":\"127.0.0.1\",\"lastLoginIp\":null,\"lastLoginAt\":null,\"remark\":\"\"}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(49,'2025-11-22 23:57:49.774','2025-11-22 23:57:49.774',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,6146875,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":45,\"CreatedAt\":\"2025-11-22T00:30:37.816+08:00\",\"UpdatedAt\":\"2025-11-22T00:30:37.816+08:00\",\"parentId\":0,\"path\":\"smsPlatformProviderBusinessMapping\",\"name\":\"smsPlatformProviderBusinessMapping\",\"hidden\":false,\"component\":\"plugin/sms/view/smsPlatformProviderBusinessMapping.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"平台子业务\",\"icon\":\"apple\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(50,'2025-11-22 23:58:03.649','2025-11-22 23:58:03.649',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,4796875,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":44,\"CreatedAt\":\"2025-11-22T00:29:33.717+08:00\",\"UpdatedAt\":\"2025-11-22T00:29:33.717+08:00\",\"parentId\":0,\"path\":\"smsPlatformBusinessTypes\",\"name\":\"smsPlatformBusinessTypes\",\"hidden\":false,\"component\":\"plugin/sms/view/smsPlatformBusinessTypes.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"平台业务\",\"icon\":\"add-location\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(51,'2025-11-22 23:58:44.960','2025-11-22 23:58:44.960',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,4085292,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":42,\"CreatedAt\":\"2025-11-21T17:36:33.878+08:00\",\"UpdatedAt\":\"2025-11-21T17:37:08.442+08:00\",\"parentId\":0,\"path\":\"smsTransactions\",\"name\":\"smsTransactions\",\"hidden\":false,\"component\":\"plugin/sms/view/smsTransactions.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"余额变动记录\",\"icon\":\"chicken\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(52,'2025-11-22 23:58:53.489','2025-11-22 23:58:53.489',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,4847500,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":41,\"CreatedAt\":\"2025-11-21T17:35:33.963+08:00\",\"UpdatedAt\":\"2025-11-21T17:37:30.787+08:00\",\"parentId\":0,\"path\":\"smsProviders\",\"name\":\"smsProviders\",\"hidden\":false,\"component\":\"plugin/sms/view/smsProviders.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"三方渠道\",\"icon\":\"coffee\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(53,'2025-11-22 23:59:03.266','2025-11-22 23:59:03.266',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,4093500,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":43,\"CreatedAt\":\"2025-11-22T00:10:43.892+08:00\",\"UpdatedAt\":\"2025-11-22T00:13:52.474+08:00\",\"parentId\":0,\"path\":\"smsProvidersBusinessTypes\",\"name\":\"smsProvidersBusinessTypes\",\"hidden\":false,\"component\":\"plugin/sms/view/smsProvidersBusinessTypes.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"三方渠道子业务\",\"icon\":\"alarm-clock\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(54,'2025-11-23 00:01:46.585','2025-11-23 00:01:46.585',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,5202500,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":40,\"CreatedAt\":\"2025-11-21T17:33:01.945+08:00\",\"UpdatedAt\":\"2025-11-21T17:37:46.112+08:00\",\"parentId\":0,\"path\":\"smsPhoneAssignments\",\"name\":\"smsPhoneAssignments\",\"hidden\":false,\"component\":\"plugin/sms/view/smsPhoneAssignments.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单管理\",\"icon\":\"film\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(55,'2025-11-23 00:02:00.953','2025-11-23 00:02:00.953',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,4792583,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":38,\"CreatedAt\":\"2025-11-21T17:30:54.258+08:00\",\"UpdatedAt\":\"2025-11-21T17:38:15.827+08:00\",\"parentId\":0,\"path\":\"smsCustomers\",\"name\":\"smsCustomers\",\"hidden\":false,\"component\":\"plugin/sms/view/smsCustomers.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"商户管理\",\"icon\":\"dish\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),(56,'2025-11-23 00:02:47.301','2025-11-23 00:02:47.301',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,3763833,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36','','{\"ID\":38,\"CreatedAt\":\"2025-11-21T17:30:54.258+08:00\",\"UpdatedAt\":\"2025-11-23T00:02:00.95+08:00\",\"parentId\":0,\"path\":\"smsCustomers\",\"name\":\"smsCustomers\",\"hidden\":false,\"component\":\"plugin/sms/view/smsCustomers.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"商户管理\",\"icon\":\"dish\",\"closeTab\":false,\"transitionType\":\"\"},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1);
/*!40000 ALTER TABLE `sys_operation_records` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_params`
--

DROP TABLE IF EXISTS `sys_params`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_params` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '参数名称',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '参数键',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '参数值',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '参数说明',
  PRIMARY KEY (`id`),
  KEY `idx_sys_params_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_params`
--

LOCK TABLES `sys_params` WRITE;
/*!40000 ALTER TABLE `sys_params` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_params` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_authority`
--

DROP TABLE IF EXISTS `sys_user_authority`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_authority`
--

LOCK TABLES `sys_user_authority` WRITE;
/*!40000 ALTER TABLE `sys_user_authority` DISABLE KEYS */;
INSERT INTO `sys_user_authority` VALUES (1,888),(1,8881),(1,9528),(2,888);
/*!40000 ALTER TABLE `sys_user_authority` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_users`
--

DROP TABLE IF EXISTS `sys_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '系统用户' COMMENT '用户昵称',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `authority_id` bigint unsigned DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  `origin_setting` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '配置',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`),
  KEY `idx_sys_users_uuid` (`uuid`),
  KEY `idx_sys_users_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_users`
--

LOCK TABLES `sys_users` WRITE;
/*!40000 ALTER TABLE `sys_users` DISABLE KEYS */;
INSERT INTO `sys_users` VALUES (1,'2025-11-21 17:18:36.110','2025-11-21 17:18:36.111',NULL,'e1e7bbad-2bfc-4337-b013-038ea13858d6','admin','$2a$10$hRo7VwZRiftlYOUSdynVueELxKPr2Q4v3zWVLtDi1lH/BXtFzZBPa','Mr.奇淼','https://qmplusimg.henrongyi.top/gva_header.jpg',888,'17611111111','333333333@qq.com',1,NULL),(2,'2025-11-21 17:18:36.110','2025-11-21 17:18:36.112',NULL,'072d7153-5df9-49d4-9177-38c0256bdbd0','a303176530','$2a$10$GdovHFs3Tr65rBEIFNHj0OEQ2Od9EG8jUc4IMnqIURNLFuhSMsrv6','用户1','https://qmplusimg.henrongyi.top/1572075907logo.png',9528,'17611111111','333333333@qq.com',1,NULL);
/*!40000 ALTER TABLE `sys_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_versions`
--

DROP TABLE IF EXISTS `sys_versions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_versions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `version_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '版本名称',
  `version_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '版本号',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '版本描述',
  `version_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '版本数据JSON',
  PRIMARY KEY (`id`),
  KEY `idx_sys_versions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_versions`
--

LOCK TABLES `sys_versions` WRITE;
/*!40000 ALTER TABLE `sys_versions` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_versions` ENABLE KEYS */;
UNLOCK TABLES;
SET @@SESSION.SQL_LOG_BIN = @MYSQLDUMP_TEMP_LOG_BIN;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-11-23  0:18:46
