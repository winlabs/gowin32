/*
 * Copyright (c) 2014 MongoDB, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the license is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gowin32

import (
	"github.com/winlabs/gowin32/wrappers"

	"syscall"
)

type ResourceType *uint16

func CustomResourceType(resourceTypeName string) ResourceType {
	return ResourceType(syscall.StringToUTF16Ptr(resourceTypeName))
}

var (
	ResourceTypeCursor        ResourceType = wrappers.RT_CURSOR
	ResourceTypeBitmap        ResourceType = wrappers.RT_BITMAP
	ResourceTypeIcon          ResourceType = wrappers.RT_ICON
	ResourceTypeMenu          ResourceType = wrappers.RT_MENU
	ResourceTypeDialog        ResourceType = wrappers.RT_DIALOG
	ResourceTypeString        ResourceType = wrappers.RT_STRING
	ResourceTypeFontDir       ResourceType = wrappers.RT_FONTDIR
	ResourceTypeFont          ResourceType = wrappers.RT_FONT
	ResourceTypeAccelerator   ResourceType = wrappers.RT_ACCELERATOR
	ResourceTypeRCData        ResourceType = wrappers.RT_RCDATA
	ResourceTypeMessageTable  ResourceType = wrappers.RT_MESSAGETABLE
	ResourceTypeGroupCursor   ResourceType = wrappers.RT_GROUP_CURSOR
	ResourceTypeGroupIcon     ResourceType = wrappers.RT_GROUP_ICON
	ResourceTypeVersion       ResourceType = wrappers.RT_VERSION
	ResourceTypeDialogInclude ResourceType = wrappers.RT_DLGINCLUDE
	ResourceTypePlugPlay      ResourceType = wrappers.RT_PLUGPLAY
	ResourceTypeVxD           ResourceType = wrappers.RT_VXD
	ResourceTypeAniCursor     ResourceType = wrappers.RT_ANICURSOR
	ResourceTypeAniIcon       ResourceType = wrappers.RT_ANIICON
	ResourceTypeHTML          ResourceType = wrappers.RT_HTML
	ResourceTypeManifest      ResourceType = wrappers.RT_MANIFEST
)

type ResourceId *uint16

func IntResourceId(resourceId uint16) ResourceId {
	return ResourceId(wrappers.MakeIntResource(resourceId))
}

func StringResourceId(resourceId string) ResourceId {
	return ResourceId(syscall.StringToUTF16Ptr(resourceId))
}

type ResourceUpdate struct {
	handle syscall.Handle
}

func NewResourceUpdate(fileName string, deleteExistingResources bool) (*ResourceUpdate, error) {
	hUpdate, err := wrappers.BeginUpdateResource(syscall.StringToUTF16Ptr(fileName), deleteExistingResources)
	if err != nil {
		return nil, err
	}
	return &ResourceUpdate{handle: hUpdate}, nil
}

func (self *ResourceUpdate) Close() error {
	if self.handle != 0 {
		if err := wrappers.EndUpdateResource(self.handle, true); err != nil {
			return err
		}
		self.handle = 0
	}
	return nil
}

func (self *ResourceUpdate) Save() error {
	if err := wrappers.EndUpdateResource(self.handle, false); err != nil {
		return err
	}
	self.handle = 0
	return nil
}

func (self *ResourceUpdate) Update(resourceType ResourceType, resourceId ResourceId, language Language, data []byte) error {
	return wrappers.UpdateResource(
		self.handle,
		(*uint16)(resourceType),
		(*uint16)(resourceId),
		uint16(language),
		&data[0],
		uint32(len(data)))
}

func (self *ResourceUpdate) Delete(resourceType ResourceType, resourceId ResourceId, language Language) error {
	return wrappers.UpdateResource(
		self.handle,
		(*uint16)(resourceType),
		(*uint16)(resourceId),
		uint16(language),
		nil,
		0)
}
