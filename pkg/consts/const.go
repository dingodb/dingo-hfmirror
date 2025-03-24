//  Copyright (c) 2025 dingodb.com, Inc. All Rights Reserved
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http:www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package consts

import "time"

var RepoTypesMapping = map[string]RepoType{
	"models":   RepoTypeModel,
	"spaces":   RepoTypeSpace,
	"datasets": RepoTypeDataset,
}

// repo类型
type RepoType string

const (
	RepoTypeModel   RepoType = RepoType("model")
	RepoTypeSpace            = RepoType("space")
	RepoTypeDataset          = RepoType("dataset")
)

func (a RepoType) Value() string {
	return string(a)
}

var ApiTimeOut = 15 * time.Second

// 定义常量
const SmallFileSize = 1024 * 1024   // 小文件大小
const SliceBytes = 1024 * 1024 * 1  // 分片大小
const UploadRetryChannelNum = 100   // 上传的重试通道队列大小
const DownloadRetryChannelNum = 100 // 下载的重试通道队列大小
const UploadTimeout = 300           // 上传超时时间，单位秒
const DownloadTimeout = 300         // 上传超时时间，单位秒
const UpGoroutineMaxNumPerFile = 10 // 每个上传文件开启的goroutine最大数量
const DpGoroutineMaxNumPerFile = 10 // 每个下载文件开启的goroutine最大数量

// const BlockSize = 1024 //8 * 1024 *
const BlockSize = 8 * 1024 * 1024 // 8 * 1024 *
