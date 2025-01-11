// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_customer

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CustomerDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"customer": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"customer_id": schema.StringAttribute{
						Computed: true,
					},
					"email": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
				},
				CustomType: CustomerType{
					ObjectType: types.ObjectType{
						AttrTypes: CustomerValue{}.AttributeTypes(ctx),
					},
				},
				Computed: true,
			},
			"id": schema.StringAttribute{
				Required:            true,
				Description:         "Id of the customer to be retrieved.",
				MarkdownDescription: "Id of the customer to be retrieved.",
			},
		},
	}
}

type CustomerModel struct {
	Customer CustomerValue `tfsdk:"customer"`
	Id       types.String  `tfsdk:"id"`
}

var _ basetypes.ObjectTypable = CustomerType{}

type CustomerType struct {
	basetypes.ObjectType
}

func (t CustomerType) Equal(o attr.Type) bool {
	other, ok := o.(CustomerType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t CustomerType) String() string {
	return "CustomerType"
}

func (t CustomerType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	customerIdAttribute, ok := attributes["customer_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`customer_id is missing from object`)

		return nil, diags
	}

	customerIdVal, ok := customerIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`customer_id expected to be basetypes.StringValue, was: %T`, customerIdAttribute))
	}

	emailAttribute, ok := attributes["email"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`email is missing from object`)

		return nil, diags
	}

	emailVal, ok := emailAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`email expected to be basetypes.StringValue, was: %T`, emailAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return CustomerValue{
		CustomerId: customerIdVal,
		Email:      emailVal,
		Name:       nameVal,
		state:      attr.ValueStateKnown,
	}, diags
}

func NewCustomerValueNull() CustomerValue {
	return CustomerValue{
		state: attr.ValueStateNull,
	}
}

func NewCustomerValueUnknown() CustomerValue {
	return CustomerValue{
		state: attr.ValueStateUnknown,
	}
}

func NewCustomerValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (CustomerValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing CustomerValue Attribute Value",
				"While creating a CustomerValue value, a missing attribute value was detected. "+
					"A CustomerValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CustomerValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid CustomerValue Attribute Type",
				"While creating a CustomerValue value, an invalid attribute value was detected. "+
					"A CustomerValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CustomerValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("CustomerValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra CustomerValue Attribute Value",
				"While creating a CustomerValue value, an extra attribute value was detected. "+
					"A CustomerValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra CustomerValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewCustomerValueUnknown(), diags
	}

	customerIdAttribute, ok := attributes["customer_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`customer_id is missing from object`)

		return NewCustomerValueUnknown(), diags
	}

	customerIdVal, ok := customerIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`customer_id expected to be basetypes.StringValue, was: %T`, customerIdAttribute))
	}

	emailAttribute, ok := attributes["email"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`email is missing from object`)

		return NewCustomerValueUnknown(), diags
	}

	emailVal, ok := emailAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`email expected to be basetypes.StringValue, was: %T`, emailAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewCustomerValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	if diags.HasError() {
		return NewCustomerValueUnknown(), diags
	}

	return CustomerValue{
		CustomerId: customerIdVal,
		Email:      emailVal,
		Name:       nameVal,
		state:      attr.ValueStateKnown,
	}, diags
}

func NewCustomerValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) CustomerValue {
	object, diags := NewCustomerValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewCustomerValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t CustomerType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewCustomerValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewCustomerValueUnknown(), nil
	}

	if in.IsNull() {
		return NewCustomerValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewCustomerValueMust(CustomerValue{}.AttributeTypes(ctx), attributes), nil
}

func (t CustomerType) ValueType(ctx context.Context) attr.Value {
	return CustomerValue{}
}

var _ basetypes.ObjectValuable = CustomerValue{}

type CustomerValue struct {
	CustomerId basetypes.StringValue `tfsdk:"customer_id"`
	Email      basetypes.StringValue `tfsdk:"email"`
	Name       basetypes.StringValue `tfsdk:"name"`
	state      attr.ValueState
}

func (v CustomerValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["customer_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["email"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.CustomerId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["customer_id"] = val

		val, err = v.Email.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["email"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v CustomerValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v CustomerValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v CustomerValue) String() string {
	return "CustomerValue"
}

func (v CustomerValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"customer_id": basetypes.StringType{},
		"email":       basetypes.StringType{},
		"name":        basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"customer_id": v.CustomerId,
			"email":       v.Email,
			"name":        v.Name,
		})

	return objVal, diags
}

func (v CustomerValue) Equal(o attr.Value) bool {
	other, ok := o.(CustomerValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.CustomerId.Equal(other.CustomerId) {
		return false
	}

	if !v.Email.Equal(other.Email) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	return true
}

func (v CustomerValue) Type(ctx context.Context) attr.Type {
	return CustomerType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v CustomerValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"customer_id": basetypes.StringType{},
		"email":       basetypes.StringType{},
		"name":        basetypes.StringType{},
	}
}