package main

import (
	"gen"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"context"
)

func main() {
	fmt.Println("Hello World!")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := gen.NewMetadataStoreServiceClient(conn)

	//Put dataset artifact
	datasetType := gen.ArtifactType{}
	dataName := "Dataset"
	datasetType.Name = &dataName
	properties := make(map[string]gen.PropertyType)
	datasetType.Properties = properties
	datasetType.Properties["day"] = gen.PropertyType_INT
	datasetType.Properties["split"] = gen.PropertyType_STRING

	request := gen.PutArtifactTypeRequest{}
	request.ArtifactType = &datasetType

	response, err := c.PutArtifactType(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling PutArtifactType: %s", err)
	}
	log.Printf("Request Success(?) Response from server, TypeId: %s", response.GetTypeId())

	//Put model artifact
	modelType := gen.ArtifactType{}
	modelName := "Model"
	modelType.Name = &modelName
	properties = make(map[string]gen.PropertyType)
	modelType.Properties = properties
	modelType.Properties["version"] = gen.PropertyType_INT
	modelType.Properties["name"] = gen.PropertyType_STRING

	request = gen.PutArtifactTypeRequest{}
	request.ArtifactType = &modelType

	response, err = c.PutArtifactType(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling PutArtifactType: %s", err)
	}
	log.Printf("Request Success(?) Response from server, TypeId: %s", response.GetTypeId())
/*	Python code to be translated to golang
	from grpc import insecure_channel
	from ml_metadata.proto import metadata_store_pb2
	from ml_metadata.proto import metadata_store_service_pb2
	from ml_metadata.proto import metadata_store_service_pb2_grpc

	channel = insecure_channel('localhost:8080')
	stub = metadata_store_service_pb2_grpc.MetadataStoreServiceStub(channel)

	# Create ArtifactTypes, e.g., Data and Model
	data_type = metadata_store_pb2.ArtifactType()
	data_type.name = "DataSet"
	data_type.properties["day"] = metadata_store_pb2.INT
	data_type.properties["split"] = metadata_store_pb2.STRING

	request = metadata_store_service_pb2.PutArtifactTypeRequest()
	request.all_fields_match = True
	request.artifact_type.CopyFrom(data_type)
	stub.PutArtifactType(request)

	model_type = metadata_store_pb2.ArtifactType()
	model_type.name = "SavedModel"
	model_type.properties["version"] = metadata_store_pb2.INT
	model_type.properties["name"] = metadata_store_pb2.STRING

	request.artifact_type.CopyFrom(model_type)
	stub.PutArtifactType(request)
*/
}
