package api

import v1 "karuhundeveloper.com/gostarterkit/routes/api/v1"

func V1() {
	v1.V1Auth()
	v1.V1Role()
}