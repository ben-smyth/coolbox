package json2yaml

import (
	"testing"
)

func TestConvertJson2Yaml(t *testing.T) {
	tests := []struct {
		name     string
		jsonStr  string
		expected string
	}{
		{
			name: "Simple JSON Object",
			jsonStr: `{
				"name": "John Doe",
				"age": 30,
				"isEmployee": true,
				"department": "Engineering"
			}`,
			expected: `age: 30
department: Engineering
isEmployee: true
name: John Doe
`,
		},
		{
			name: "Nested JSON Object",
			jsonStr: `{
				"name": "Jane Smith",
				"age": 28,
				"address": {
					"street": "123 Elm St",
					"city": "Springfield",
					"zip": "12345"
				},
				"isEmployee": false
			}`,
			expected: `address:
  city: Springfield
  street: 123 Elm St
  zip: "12345"
age: 28
isEmployee: false
name: Jane Smith
`,
		},
		{
			name: "JSON with Array",
			jsonStr: `{
				"company": "TechCorp",
				"employees": [
					{
						"name": "Alice",
						"role": "Developer"
					},
					{
						"name": "Bob",
						"role": "Designer"
					}
				],
				"totalEmployees": 2
			}`,
			expected: `company: TechCorp
employees:
  - name: Alice
    role: Developer
  - name: Bob
    role: Designer
totalEmployees: 2
`,
		},
		{
			name: "JSON with Mixed Data Types",
			jsonStr: `{
				"project": "Data Migration",
				"budget": 5000.75,
				"completed": false,
				"tasks": [
					{
						"id": 1,
						"title": "Initial Setup",
						"status": "done"
					},
					{
						"id": 2,
						"title": "Data Transfer",
						"status": "in progress"
					},
					{
						"id": 3,
						"title": "Verification",
						"status": "pending"
					}
				],
				"startDate": "2024-01-15",
				"endDate": null
			}`,
			expected: `budget: 5000.75
completed: false
endDate: null
project: Data Migration
startDate: "2024-01-15"
tasks:
  - id: 1
    status: done
    title: Initial Setup
  - id: 2
    status: in progress
    title: Data Transfer
  - id: 3
    status: pending
    title: Verification
`,
		},
		{
			name: "Deeply Nested JSON",
			jsonStr: `{
				"organization": {
					"name": "Global Solutions",
					"departments": {
						"IT": {
							"manager": "Emily Clarke",
							"team": [
								{
									"name": "Alice",
									"role": "System Administrator"
								},
								{
									"name": "Bob",
									"role": "Network Engineer"
								}
							]
						},
						"HR": {
							"manager": "David Smith",
							"team": [
								{
									"name": "Charlie",
									"role": "Recruiter"
								}
							]
						}
					}
				}
			}`,
			expected: `organization:
  departments:
    HR:
      manager: David Smith
      team:
        - name: Charlie
          role: Recruiter
    IT:
      manager: Emily Clarke
      team:
        - name: Alice
          role: System Administrator
        - name: Bob
          role: Network Engineer
  name: Global Solutions
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yamlStr, err := ConvertJson2Yaml(tt.jsonStr)
			if err != nil {
				t.Errorf("ConvertJson2Yaml() error = %v", err)
				return
			}

			if yamlStr != tt.expected {
				t.Errorf("ConvertJson2Yaml() got = %v, want %v", yamlStr, tt.expected)
			}
		})
	}
}
