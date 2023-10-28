package models

import TagDTO "api/routes/modules/tags/dtos"

type ReqCheckpointxTags struct {
	ID      uint                  `json:"id_checkpoint"`
	TagList []TagDTO.SimpleTagDTO `json:"tag_list"`
}
