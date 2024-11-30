# HOW TO USE
1. Create Database on MySQL/MariaDB
2. Setup your Database Config & Apps Config on Config.yml
3. Do Migrate using this command
``
go run . migrate
``
4. Run Apps with `` go run . serve``



# Sample Api Request

``{
	"info": {
		"_postman_id": "32445938-3703-4d79-956f-f86416863c89",
		"name": "Interview",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "27416807",
		"_collection_link": "https://interstellar-water-4440.postman.co/workspace/Dev~fb6c679f-508f-44e0-bbba-589566e5fe2d/collection/27416807-32445938-3703-4d79-956f-f86416863c89?action=share&source=collection_link&creator=27416807"
	},
	"item": [
		{
			"name": "WidaTech",
			"item": [
				{
					"name": "Check Request",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "{{url}}",
							"host": [
								"{{url}}"
							]
						}
					},
					"response": []
				},
				{
					"name": "Create",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"invoice_no\": \"INV-2\",\r\n    \"date\": \"10/11/2024\",\r\n    \"customer_name\": \"Dyeva\",\r\n    \"sales_person_name\": \"Deva\",\r\n    \"payment_type\": \"CASH\",\r\n    \"notes\": \"\",\r\n    \"list_of_product\": [\r\n        {\r\n            \"item_name\": \"C\",\r\n            \"quantity\": 1,\r\n            \"total_cogs\": 100,\r\n            \"total_price\": 100\r\n        }\r\n    ]\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{url}}/create",
							"host": [
								"{{url}}"
							],
							"path": [
								"create"
							]
						}
					},
					"response": []
				},
				{
					"name": "Update",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"invoice_no\": \"INV-1\",\r\n    \"date\": \"10/11/2024\",\r\n    \"customer_name\": \"Dyeva\",\r\n    \"sales_person_name\": \"Deva\",\r\n    \"payment_type\": \"CREDIT\",\r\n    \"notes\": \"\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{url}}/update",
							"host": [
								"{{url}}"
							],
							"path": [
								"update"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete",
					"request": {
						"method": "DELETE",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"invoice_no\": \"INV-3\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{url}}/delete",
							"host": [
								"{{url}}"
							],
							"path": [
								"delete"
							]
						}
					},
					"response": []
				},
				{
					"name": "Import",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "formdata",
							"formdata": [
								{
									"key": "file",
									"type": "file",
									"src": "/D:/Pribadi/int/wida tech/InvoiceImport-2.xlsx"
								}
							]
						},
						"url": {
							"raw": "{{url}}/import",
							"host": [
								"{{url}}"
							],
							"path": [
								"import"
							]
						}
					},
					"response": []
				},
				{
					"name": "Read",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"date\": \"10/11/2024\",\r\n    \"size\": 10,\r\n    \"page\": 1\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{url}}/read",
							"host": [
								"{{url}}"
							],
							"path": [
								"read"
							]
						}
					},
					"response": []
				}
			]
		}
	]
}``