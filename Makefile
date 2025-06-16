SWAG_CMD = swag init -d cmd/admin-service/,internal/http/handlers/,internal/http/dto/ --parseInternal -o pkg/docs/

.PHONY: swagger
swagger:
	$(SWAG_CMD)