# Custom Terraform Provider Sample  

This repository demonstrates how to generate a custom Terraform provider using the OpenAPI Provider Spec Generator and the Framework Code Generator. Follow the steps below to create your own custom provider, complete with resources and data sources.

---

## Prerequisites  

- Go (>= 1.22.10)  
- Terraform CLI  
- Access to an OpenAPI specification file (e.g., `openapi.json`)  

---

## Steps to Generate the Custom Provider  

### 1. **Generate the Provider Spec**  

The first step is to generate a Provider Code Specification from your OpenAPI spec. This serves as the foundation for the custom Terraform provider.  

```bash  
# Generate the Provider Code Spec  
go run github.com/hashicorp/terraform-plugin-codegen-openapi/cmd/tfplugingen-openapi@latest generate \
  --config generator_config.yml \
  --output ./provider_spec.json \
  ./openapi.json  
```

- **Input:**  
  - `openapi.json` (your OpenAPI specification).  
  - `generator_config.yml` (configuration file for the spec generator).  
- **Output:**  
  - `provider_spec.json` (Provider Code Spec).  

---

### 2. **Generate Provider Models**  

With the `provider_spec.json` ready, use the Framework Code Generator to create the foundational provider models.  

```bash  
go run github.com/hashicorp/terraform-plugin-codegen-framework/cmd/tfplugingen-framework@latest generate all \
  --input ./provider_spec.json \
  --output ./generated  
```

- **Input:**  
  - `provider_spec.json`.  
- **Output:**  
  - Generated provider models under the `generated` directory.  

---

### 3. **Add Data Sources**  

You can scaffold new data sources for your provider.

```bash  
# Ensure the `api_provider` directory exists  
mkdir -p api_provider  

# List of data sources to generate
data_sources=("orders" "order" "products" "product" "customers" "customer")

# Loop through each data source in the list
for data_source in "${data_sources[@]}"
do
  # Run the go command to scaffold the data source for each entry in the list
  go run github.com/hashicorp/terraform-plugin-codegen-framework/cmd/tfplugingen-framework@latest scaffold data-source \
    --name "$data_source" \
    --output-dir api_provider \
    --package api_provider
done
```

- **Output:**  
  - Data source scaffolding in the `api_provider` directory.  

---

### 4. **Add Resources**  

Similarly, scaffold new resources for your provider.

```bash  
# List of resources to generate
resources=("orders" "order" "products" "product" "customers" "customer")

# Loop through each resource in the list
for resource in "${resources[@]}"
do
  # Run the go command to scaffold the resource for each entry in the list
  go run github.com/hashicorp/terraform-plugin-codegen-framework/cmd/tfplugingen-framework@latest scaffold resource \
    --name "$resource" \
    --output-dir api_provider \
    --package api_provider
done
```

- **Output:**  
  - Resource scaffolding in the `api_provider` directory.  

---
---

## Next Steps

> **_NOTE:_** The `api_provider/provider.go` file was manually created to define the core logic for API interaction, including configuration and resource registration. It serves as the backbone of the provider, ensuring seamless communication with your API. Customize this file to fit your API’s specific requirements.

1. Implement the logic in the generated scaffolding to interact with your API.  
2. Compile and test your custom provider.  
3. Package and distribute your provider using a registry for use with Terraform.  