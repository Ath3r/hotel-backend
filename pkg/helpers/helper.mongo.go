package helpers

import "go.mongodb.org/mongo-driver/bson/primitive"

func ToObjectId(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}
