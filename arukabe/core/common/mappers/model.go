package mappers

import (
	modelsv1 "arukabe/gen/connect/aruka/models/v1"
	"arukabe/gen/sqlc"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelStatusToDB(status modelsv1.ModelStatus) sqlc.ModelStatus {
	switch status {
	case modelsv1.ModelStatus_MODEL_STATUS_UNSPECIFIED:
		return sqlc.ModelStatusACTIVE
	case modelsv1.ModelStatus_MODEL_STATUS_ACTIVE:
		return sqlc.ModelStatusACTIVE
	case modelsv1.ModelStatus_MODEL_STATUS_INACTIVE:
		return sqlc.ModelStatusINACTIVE
	case modelsv1.ModelStatus_MODEL_STATUS_DEPRECATED:
		return sqlc.ModelStatusDEPRECATED
	default:
		return sqlc.ModelStatusACTIVE
	}
}

func ModelStatusToPB(status sqlc.ModelStatus) modelsv1.ModelStatus {
	switch status {
	case sqlc.ModelStatusACTIVE:
		return modelsv1.ModelStatus_MODEL_STATUS_ACTIVE
	case sqlc.ModelStatusINACTIVE:
		return modelsv1.ModelStatus_MODEL_STATUS_INACTIVE
	case sqlc.ModelStatusDEPRECATED:
		return modelsv1.ModelStatus_MODEL_STATUS_DEPRECATED
	default:
		return modelsv1.ModelStatus_MODEL_STATUS_ACTIVE
	}
}

func ModelToPB(model sqlc.Model) *modelsv1.Model {
	return &modelsv1.Model{
		Id:              model.ID.String(),
		Name:            model.Name,
		ModelIdentifier: model.ModelIdentifier,
		CreatedAt:       timestamppb.New(model.CreatedAt),
		Status:          ModelStatusToPB(model.Status),
	}
}

func ModelsToPB(models []sqlc.Model) []*modelsv1.Model {
	result := make([]*modelsv1.Model, len(models))
	for i := range models {
		result[i] = ModelToPB(models[i])
	}
	return result
}
