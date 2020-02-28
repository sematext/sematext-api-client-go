package api

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// SematextMonitorCreate TODO Doc Comment
func SematextMonitorCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Client)
	plans := map[string]int64{"basic": 0, "standard": 1, "pro": 2, "enterprise": 3}
	billingPlan := d.Get("billing_code").(string)
	createAppInfo := &CreateAppInfo{}
	createAppInfo.AppType = d.Get("app_type").(string)
	createAppInfo.DiscountCode = d.Get("discount_code").(string)
	createAppInfo.InitialPlanID = plans[billingPlan]
	//createAppInfo.MetaData = &AppMetaData{} // not used for this resource
	createAppInfo.Name = d.Get("label").(string)

	app, err := createAppInfo.Create(client)
	if err != nil {
		return err
	}

	d.SetId(app.ID)
	d.Set("token", app.Token)
	// TODO - Check rest of App values to see if any relevant to inject into state.

	return nil

}

// SematextMonitorRead TODO Doc Comment
func SematextMonitorRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Client)
	id := d.Id()
	var app *App
	var err error
	app, err = app.Retrieve(id, client)
	if err != nil {
		return err
	}
	plans := map[int64]string{0: "basic", 1: "standard", 2: "pro", 3: "enterprise"}
	d.Set("name", app.Name)
	d.Set("description", app.Description)
	d.Set("billing_plan", plans[app.Plan.ID])
	//d.Set("discount_code",) // TODO - JIRA discount_code seems to not be stored in the App.
	return nil
}

// SematextMonitorUpdate TODO Doc Comment
func SematextMonitorUpdate(d *schema.ResourceData, meta interface{}) error {

	// TODO - how to do a plan update?

	client := meta.(*Client)
	id := d.Id()
	dto := &Dto{} // TODO - switch to UpdateAppInfo

	dto.Description = d.Get("description").(string)
	dto.IgnorePercentage = d.Get("ignore_percentage").(int)
	dto.MaxEvents = d.Get("max_events").(int)
	dto.MaxLimitMB = d.Get("max_limit_mb").(int)
	dto.Name = d.Get("name").(string) // TODO name vs resource name/label in terraform
	dto.Sampling = d.Get("sampling").(bool)
	dto.SamplingPercentage = d.Get("sampling_percentage").(int)
	dto.Staggering = d.Get("staggering").(bool)
	dto.Status = d.Get("status").(string) // TODO Status should be reset as part of delete?

	if dto.IsValid() == false {
		return fmt.Errorf("Coding error - dto non-valid") // TODO Check correctness.
	}

	_, err := dto.Update(id, client, *dto)
	if err != nil {
		return err
	}

	return nil

}

// SematextMonitorDelete TODO Doc Comment
func SematextMonitorDelete(d *schema.ResourceData, meta interface{}) error { // TODO Check default is respected

	client := meta.(*Client)
	id := d.Id()
	var app App
	err := app.Delete(id, client)
	if err != nil {
		return err
	}

	return nil
}

// SematextMonitorExists TODO Doc Comment
func SematextMonitorExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {

	// TODO Consider necessity for an app edit-version to catch edit-version mis-match back into state.
	client := meta.(*Client)
	id := d.Id()
	var app App
	b, e = app.Exists(id, client)
	return b, e
}

// SematextMonitorImport TODO Doc Comment
func SematextMonitorImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	// TODO Decide if necessary
	return nil, nil

}
