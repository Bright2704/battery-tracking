package models




// type Vision struct {
// 	Vision string `json:"related" bson: "related"`
// 	Related Related `json:"related" bson: "related"`
// }

type Vision struct {
	Related string `json:"related" bson: "related"`
	Stage_1 Stage_1 `json:"stage_1" bson: "stage_1"`
	Stage_2 Stage_2 `json:"stage_2" bson: "stage_2"`
	Stage_3 Stage_3 `json:"stage_3" bson: "stage_3"`
	Stage_4 Stage_4 `json:"stage_4" bson: "stage_4"`
	Stage_5 Stage_5 `json:"stage_5" bson: "stage_5"`
}

type Stage_1 struct {
	Process_name string `json:"process_name" bson: "process_name"`
	description string `json:"description" bson: "description"`
	dutation_in_hours int`json:"dutation_in_hours" bson: "dutation_in_hours"`
	responsible_department string `json:"responsible_department" bson: "responsible_department"`
	process_in bool `json:"process_in" bson: "process_in"`
	process_out bool `json:"process_out" bson: "process_out"`
}

type Stage_2 struct {
	Process_name string `json:"process_name" bson: "process_name"`
	description string `json:"description" bson: "description"`
	dutation_in_hours int`json:"dutation_in_hours" bson: "dutation_in_hours"`
	responsible_department string `json:"responsible_department" bson: "responsible_department"`
	process_in bool `json:"process_in" bson: "process_in"`
	process_out bool `json:"process_out" bson: "process_out"`
}

type Stage_3 struct {
	Process_name string `json:"process_name" bson: "process_name"`
	description string `json:"description" bson: "description"`
	dutation_in_hours int`json:"dutation_in_hours" bson: "dutation_in_hours"`
	responsible_department string `json:"responsible_department" bson: "responsible_department"`
	process_in bool `json:"process_in" bson: "process_in"`
	process_out bool `json:"process_out" bson: "process_out"`
}

type Stage_4 struct {
	Process_name string `json:"process_name" bson: "process_name"`
	description string `json:"description" bson: "description"`
	dutation_in_hours int`json:"dutation_in_hours" bson: "dutation_in_hours"`
	responsible_department string `json:"responsible_department" bson: "responsible_department"`
	process_in bool `json:"process_in" bson: "process_in"`
	process_out bool `json:"process_out" bson: "process_out"`
}

type Stage_5 struct {
	Process_name string `json:"process_name" bson: "process_name"`
	description string `json:"description" bson: "description"`
	dutation_in_hours int`json:"dutation_in_hours" bson: "dutation_in_hours"`
	responsible_department string `json:"responsible_department" bson: "responsible_department"`
	process_in bool `json:"process_in" bson: "process_in"`
	process_out bool `json:"process_out" bson: "process_out"`
}

