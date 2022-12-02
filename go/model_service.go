/*
 * Open Service Broker API
 *
 * The Open Service Broker API defines an HTTP(S) interface between Platforms and Service Brokers.
 *
 * API version: master - might contain changes that are not yet released
 * Contact: open-service-broker-api@googlegroups.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Service struct {

	Bindable bool `json:"bindable"`

	BindingRotatable bool `json:"binding_rotatable,omitempty"`

	DashboardClient *DashboardClient `json:"dashboard_client,omitempty"`

	Description string `json:"description"`

	Id string `json:"id"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Name string `json:"name"`

	PlanUpdateable bool `json:"plan_updateable,omitempty"`

	Plans []Plan `json:"plans"`

	Requires []string `json:"requires,omitempty"`

	Tags []string `json:"tags,omitempty"`
}