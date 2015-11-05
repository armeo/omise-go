/*
!!! DO NOT MODIFY !!!

autogenerated
 src: gen_list_job.tmpl
 job: &main.GenListJob{Name:"Customer"}
  on: Thu Nov 05 15:22:29 +0700 2015
  by: chakrit
*/
package omise

// CustomerList represents the list structure returned by Omise's REST API that contains
// Customer struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type CustomerList struct {
	List
	Data []*Customer `json:"data"`
}

// Find finds and returns Customer with the given id. Returns nil if not found.
func (list *CustomerList) Find(id string) *Customer {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}