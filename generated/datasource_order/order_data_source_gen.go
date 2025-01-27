// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_order

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

func OrderDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:            true,
				Description:         "Id of the order to be examined.",
				MarkdownDescription: "Id of the order to be examined.",
			},
			"order": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"customer_id": schema.StringAttribute{
						Computed: true,
					},
					"order_id": schema.StringAttribute{
						Computed: true,
					},
					"product_id": schema.StringAttribute{
						Computed: true,
					},
				},
				CustomType: OrderType{
					ObjectType: types.ObjectType{
						AttrTypes: OrderValue{}.AttributeTypes(ctx),
					},
				},
				Computed: true,
			},
		},
	}
}

type OrderModel struct {
	Id    types.String `tfsdk:"id"`
	Order OrderValue   `tfsdk:"order"`
}

var _ basetypes.ObjectTypable = OrderType{}

type OrderType struct {
	basetypes.ObjectType
}

func (t OrderType) Equal(o attr.Type) bool {
	other, ok := o.(OrderType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t OrderType) String() string {
	return "OrderType"
}

func (t OrderType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
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

	orderIdAttribute, ok := attributes["order_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`order_id is missing from object`)

		return nil, diags
	}

	orderIdVal, ok := orderIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`order_id expected to be basetypes.StringValue, was: %T`, orderIdAttribute))
	}

	productIdAttribute, ok := attributes["product_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`product_id is missing from object`)

		return nil, diags
	}

	productIdVal, ok := productIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`product_id expected to be basetypes.StringValue, was: %T`, productIdAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return OrderValue{
		CustomerId: customerIdVal,
		OrderId:    orderIdVal,
		ProductId:  productIdVal,
		state:      attr.ValueStateKnown,
	}, diags
}

func NewOrderValueNull() OrderValue {
	return OrderValue{
		state: attr.ValueStateNull,
	}
}

func NewOrderValueUnknown() OrderValue {
	return OrderValue{
		state: attr.ValueStateUnknown,
	}
}

func NewOrderValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (OrderValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing OrderValue Attribute Value",
				"While creating a OrderValue value, a missing attribute value was detected. "+
					"A OrderValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("OrderValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid OrderValue Attribute Type",
				"While creating a OrderValue value, an invalid attribute value was detected. "+
					"A OrderValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("OrderValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("OrderValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra OrderValue Attribute Value",
				"While creating a OrderValue value, an extra attribute value was detected. "+
					"A OrderValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra OrderValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewOrderValueUnknown(), diags
	}

	customerIdAttribute, ok := attributes["customer_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`customer_id is missing from object`)

		return NewOrderValueUnknown(), diags
	}

	customerIdVal, ok := customerIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`customer_id expected to be basetypes.StringValue, was: %T`, customerIdAttribute))
	}

	orderIdAttribute, ok := attributes["order_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`order_id is missing from object`)

		return NewOrderValueUnknown(), diags
	}

	orderIdVal, ok := orderIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`order_id expected to be basetypes.StringValue, was: %T`, orderIdAttribute))
	}

	productIdAttribute, ok := attributes["product_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`product_id is missing from object`)

		return NewOrderValueUnknown(), diags
	}

	productIdVal, ok := productIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`product_id expected to be basetypes.StringValue, was: %T`, productIdAttribute))
	}

	if diags.HasError() {
		return NewOrderValueUnknown(), diags
	}

	return OrderValue{
		CustomerId: customerIdVal,
		OrderId:    orderIdVal,
		ProductId:  productIdVal,
		state:      attr.ValueStateKnown,
	}, diags
}

func NewOrderValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) OrderValue {
	object, diags := NewOrderValue(attributeTypes, attributes)

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

		panic("NewOrderValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t OrderType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewOrderValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewOrderValueUnknown(), nil
	}

	if in.IsNull() {
		return NewOrderValueNull(), nil
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

	return NewOrderValueMust(OrderValue{}.AttributeTypes(ctx), attributes), nil
}

func (t OrderType) ValueType(ctx context.Context) attr.Value {
	return OrderValue{}
}

var _ basetypes.ObjectValuable = OrderValue{}

type OrderValue struct {
	CustomerId basetypes.StringValue `tfsdk:"customer_id"`
	OrderId    basetypes.StringValue `tfsdk:"order_id"`
	ProductId  basetypes.StringValue `tfsdk:"product_id"`
	state      attr.ValueState
}

func (v OrderValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["customer_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["order_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["product_id"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.CustomerId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["customer_id"] = val

		val, err = v.OrderId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["order_id"] = val

		val, err = v.ProductId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["product_id"] = val

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

func (v OrderValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v OrderValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v OrderValue) String() string {
	return "OrderValue"
}

func (v OrderValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"customer_id": basetypes.StringType{},
		"order_id":    basetypes.StringType{},
		"product_id":  basetypes.StringType{},
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
			"order_id":    v.OrderId,
			"product_id":  v.ProductId,
		})

	return objVal, diags
}

func (v OrderValue) Equal(o attr.Value) bool {
	other, ok := o.(OrderValue)

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

	if !v.OrderId.Equal(other.OrderId) {
		return false
	}

	if !v.ProductId.Equal(other.ProductId) {
		return false
	}

	return true
}

func (v OrderValue) Type(ctx context.Context) attr.Type {
	return OrderType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v OrderValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"customer_id": basetypes.StringType{},
		"order_id":    basetypes.StringType{},
		"product_id":  basetypes.StringType{},
	}
}
