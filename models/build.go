//
package models

import "time"

//
type (

	//
	Build struct {
		CreatedAt        time.Time  `json:"created_at"`        //
		ID               string     `json:"build_id"`          // NOTE: when pulling from the database, this will be organized alphabetically as ID
		Logs             []BuildLog `json:"logs"`              //
		MetaData         string     `json:"meta_data"`         // return data from [Packer]
		State            string     `json:"state"`             // incomplete, complete, failed
		Status           string     `json:"status"`            // installing template, copying template, transforming template, running packer, recording results, cleaning up
		TemplateID       string     `json:"template_id"`       // ID of base template
		TransformPayload string     `json:"transform_payload"` // payload that will get passed into the transform_script from template
		UpdatedAt        time.Time  `json:"updated_at"`        //
	}

	//
	BuildLog struct {
		BuildID   string    `json:"build_id"`     //
		CreatedAt time.Time `json:"created_at"`   //
		ID        string    `json:"build_log_id"` // NOTE: when pulling from the database, this will be organized alphabetically as ID
		Message   string    `json:"message"`      //
	}
)
