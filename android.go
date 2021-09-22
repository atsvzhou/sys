package sys

import "github.com/lunny/axmlParser"

// GetApkVersionCode 获取 apk 版本
func GetApkVersionCode(filePath string) (string, error) {
	listener := new(axmlParser.AppNameListener)
	_, err := axmlParser.ParseApk(filePath, listener)
	versionCode := listener.VersionCode
	return versionCode, err
}
