// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PlanInputPlanChargesChargeModelEnum - Charge model type
type PlanInputPlanChargesChargeModelEnum string

const (
	PlanInputPlanChargesChargeModelEnumStandard   PlanInputPlanChargesChargeModelEnum = "standard"
	PlanInputPlanChargesChargeModelEnumGraduated  PlanInputPlanChargesChargeModelEnum = "graduated"
	PlanInputPlanChargesChargeModelEnumPackage    PlanInputPlanChargesChargeModelEnum = "package"
	PlanInputPlanChargesChargeModelEnumPercentage PlanInputPlanChargesChargeModelEnum = "percentage"
	PlanInputPlanChargesChargeModelEnumVolume     PlanInputPlanChargesChargeModelEnum = "volume"
)

func (e PlanInputPlanChargesChargeModelEnum) ToPointer() *PlanInputPlanChargesChargeModelEnum {
	return &e
}

func (e *PlanInputPlanChargesChargeModelEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "standard":
		fallthrough
	case "graduated":
		fallthrough
	case "package":
		fallthrough
	case "percentage":
		fallthrough
	case "volume":
		*e = PlanInputPlanChargesChargeModelEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlanInputPlanChargesChargeModelEnum: %v", v)
	}
}

type PlanInputPlanChargesGroupProperties struct {
	GroupID string                 `json:"group_id"`
	Values  map[string]interface{} `json:"values"`
}

type PlanInputPlanCharges struct {
	BillableMetricID *string `json:"billable_metric_id,omitempty"`
	// Charge model type
	ChargeModel     *PlanInputPlanChargesChargeModelEnum  `json:"charge_model,omitempty"`
	GroupProperties []PlanInputPlanChargesGroupProperties `json:"group_properties,omitempty"`
	ID              *string                               `json:"id,omitempty"`
	Instant         *bool                                 `json:"instant,omitempty"`
	Properties      map[string]interface{}                `json:"properties,omitempty"`
}

// PlanInputPlanIntervalEnum - Plan interval
type PlanInputPlanIntervalEnum string

const (
	PlanInputPlanIntervalEnumWeekly  PlanInputPlanIntervalEnum = "weekly"
	PlanInputPlanIntervalEnumMonthly PlanInputPlanIntervalEnum = "monthly"
	PlanInputPlanIntervalEnumYearly  PlanInputPlanIntervalEnum = "yearly"
)

func (e PlanInputPlanIntervalEnum) ToPointer() *PlanInputPlanIntervalEnum {
	return &e
}

func (e *PlanInputPlanIntervalEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "weekly":
		fallthrough
	case "monthly":
		fallthrough
	case "yearly":
		*e = PlanInputPlanIntervalEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlanInputPlanIntervalEnum: %v", v)
	}
}

type PlanInputPlan struct {
	AmountCents        *int64                 `json:"amount_cents,omitempty"`
	AmountCurrency     *string                `json:"amount_currency,omitempty"`
	BillChargesMonthly *bool                  `json:"bill_charges_monthly,omitempty"`
	Charges            []PlanInputPlanCharges `json:"charges,omitempty"`
	Code               *string                `json:"code,omitempty"`
	Description        *string                `json:"description,omitempty"`
	// Plan interval
	Interval     *PlanInputPlanIntervalEnum `json:"interval,omitempty"`
	Name         *string                    `json:"name,omitempty"`
	PayInAdvance *bool                      `json:"pay_in_advance,omitempty"`
	TrialPeriod  *float64                   `json:"trial_period,omitempty"`
}

// PlanInput - Plan payload
type PlanInput struct {
	Plan PlanInputPlan `json:"plan"`
}
